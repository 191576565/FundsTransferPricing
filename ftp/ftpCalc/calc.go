// Copyright 2016 huangzhanwei. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ftpCalc

import (
	"dbobj"
	"errors"
	"fmt"
	"mas/ftp/calcLog"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
	"time"
	"utils"
)

//
// calculate ftp rate
// this struct has 9 member
// first precipitation keeps weight values.
// bm keeps business price method
// cv keeps curve info
// ct keeps curve info struct
// msg is used to send message to other process
// status keeps the calculate process status.
// -1 is exit calculate process.
// Author: hzwy23
// Time: 2016-08-22
//
type CalcFTP struct {
	precipitation map[string][]precipitation
	bm            map[string]busizMethod
	cv            map[string]*curve
	ct            map[string]curveStruct
	adjust        *ftpAdjustment
	msg           string
	lock          *sync.RWMutex
	status        int
	dispatchId    string
	inputSrc      string
	outputSrc     string
	rstCache      chan ftpResult
	inputCache    chan acctInfo
	log           calcLog.Calclog
	wg            *sync.WaitGroup
	RUNTIME_CNT   chan int
	as_of_date    string
	offset        int64
	limit         int64
	domain_id     string
	successCnt    int64
	totalCnt      int64
	errorCnt      int64
	noBusizCnt    int64
}

// Function: 采用流式计算方式,逐步读取账户数据,并计算
//           降低内存使用率
// Author: 黄占威
//stream calculate
func (this *CalcFTP) calcStream(as_of_date string, domain_id string, offset int, limit int) error {
	defer func() {
		if r := recover(); r != nil {
			updateDispatchStatus(this.dispatchId, as_of_date, "2")
			this.msg = "批次运行失败,请检查曲线，金融工具表，规则配置表数据情况"
			this.log.Error("批次运行初始化过程中失败", r)
		}
	}()
	/*
	* 手工开启多个核心,防止低版本gosdk无法自动启动多核
	 */
	var MULTICORE int = runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(MULTICORE)        //running in multicore

	//输入参数校验
	//as_of_date 表示批次日期,该便令一定要是一个正确的日期类型
	//domain_id  表示域名
	//offset     表示偏移量
	//limit      表示查询的最大行数
	//as, err := utils.FormatStringToDate(as_of_date)

	if as, err := utils.FormatStringToDate(as_of_date); err != nil {
		this.log.Error("日期不对，请输入正确的格式：yyyy/mm/ddss或yyyy-mm-dd。并且请检查日期是否符合正确的逻辑")
		this.msg = "日期不对，请输入正确的格式：yyyy/mm/ddss或yyyy-mm-dd。并且请检查日期是否符合正确的逻辑"
		return errors.New("日期不对，请输入正确的格式：yyyy/mm/ddss或yyyy-mm-dd。并且请检查日期是否符合正确的逻辑")
	} else {
		as_of_date = as
	}

	// init as_of_date
	this.as_of_date = as_of_date

	if !utils.ValidWord(domain_id, 1, 30) {
		this.log.Error("域名格式不正确。请输入1-30位英文字母，数字，下滑线组合的字符串")
		return errors.New("域名格式不正确。请输入1-30位英文字母，数字，下滑线组合的字符串")
	}
	if !utils.ValidNumber(strconv.Itoa(offset)) {
		this.log.Error("偏移地址因该为数字，请检查第三个参数是否为数字")
		return errors.New("偏移地址因该为数字，请检查第三个参数是否为数字")
	}
	if !utils.ValidNumber(strconv.Itoa(limit)) {
		this.log.Error("第四个参数应该是数字，请检查参数是否符合要求")
		return errors.New("第四个参数应该是数字，请检查参数是否符合要求")
	}
	this.offset = int64(offset)
	this.limit = int64(limit)
	this.domain_id = domain_id

	// init sync
	this.wg = new(sync.WaitGroup)
	this.wg.Add(5)

	// 清除ftp定价结果集数据,取消自动删除结果集功能
	// 需要手动删除结果集
	// ClearFtpRate()

	fmt.Println("开始计算FTP价格", time.Now().Format("2006-01-02 03:04:05"), this.dispatchId, as_of_date, domain_id)
	//首先获取账户数据

	go GetAcctInfoStream(as_of_date, domain_id, offset, limit, this)
	this.log.Info(this.dispatchId, "-> Getting account info by backgroud process. starting...", as_of_date)
	//go GetAcctInfoStream(as_of_date, domain_id, offset, limit, rst, this.dispatchId, this.inputSrc, this.status, this.lock)

	//获取定价规则信息
	go func() {
		this.bm = GetBusizMethod(domain_id, this)
	}()

	//获取最近几年曲线，
	//每一年按照一个区间进行分区处理
	go func() {
		this.cv = GetCurveVal(domain_id, as_of_date[0:4], this)
	}()

	//获取所有期限点信息
	go func() {
		this.ct = GetCurveStruct(domain_id, this)
	}()

	//获取偿还曲线配置信息
	go func() {
		this.precipitation = getPrecipitation(domain_id, this)
	}()

	go func() {
		this.adjust = ftpAdjustInit(domain_id, this, as_of_date)
	}()

	// waiting for init process.
	//
	this.wg.Wait()

	if this.status != 0 {
		this.log.Error(this.dispatchId, "-> init batch config info failed. please check log info and analysic problem.")
		return errors.New("init rules failed.")
	} else {
		this.log.Info(this.dispatchId, "-> starting calculate ftp.")
	}

	go this.isJack(as_of_date)

	this.RUNTIME_CNT = make(chan int, MULTICORE*10)

	for val := range this.inputCache {
		//this.totalCnt = atomic.AddInt64(&this.totalCnt, 1)
		biz := val.busiz_id

		//fmt.Println("data buffer:", len(this.inputCache), len(RUNTIME_CNT))
		if m, ok := this.bm[biz]; ok {
			this.RUNTIME_CNT <- 1
			go this.calculateByParallel(val, m)
		} else {
			this.log.Warn(biz, "没有匹配上定价单元")
			this.lock.Lock()
			this.noBusizCnt++
			this.lock.Unlock()
			//this.noBusizCnt = atomic.AddInt64(&this.noBusizCnt, 1)
		}
		this.lock.RLock()
		if this.status != 0 {
			this.lock.RUnlock()
			this.log.Debug("batch was stopped", this.dispatchId, as_of_date)
			this.lock.RLock()
			if this.status == 3 {
				this.lock.RUnlock()
				fmt.Println("FTP价格计算完成", time.Now().Format("2006-01-02 03:04:05"))
			}
			return nil
		} else {
			this.lock.RUnlock()
		}
	}
	this.checkResult()
	fmt.Println("FTP价格计算计算完成，完成时间是：", time.Now().Format("2006-01-02 03:04:05"))
	//数据匹配定价方法
	return nil
}

func (this *CalcFTP) checkResult() {
	this.wg.Wait()
	this.lock.RLock()
	//if this.successCnt+this.errorCnt+this.noBusizCnt != this.totalCnt {
	//	time.Sleep(time.Second * 2)
	//	fmt.Println(this.successCnt, this.errorCnt, this.totalCnt)
	//}
	//fmt.Println(this.successCnt, this.errorCnt, this.totalCnt, this.noBusizCnt)
	if this.errorCnt > 0 {
		err_msg := "错误条数是：" + strconv.Itoa(int(this.errorCnt))
		e := dbobj.Default.Exec("update ftp_dispatch_pro set err_msg = :1,dispatch_status='2' where dispatch_id = :2 and dispatch_date = to_date(:3,'YYYY-MM-DD') and domain_id = :4", err_msg, this.dispatchId, this.as_of_date, this.domain_id)
		if e != nil {
			this.log.Error(e, "写入批次结果信息失败，请联系管理员")
			return
		}
		return
	}
	this.status = 3
	updateDispatchStatus(this.dispatchId, this.as_of_date, "3")
	return
}

func (this *CalcFTP) isJack(as_of_date string) {
	var st string
	for {
		err := dbobj.Default.QueryRow(FTPCALC_ISJACK, this.dispatchId, as_of_date).Scan(&st)
		if err != nil {
			this.log.Debug(this.dispatchId, "-> this is backgroud process. if you want to stop this batch. please insert this batch info to table(ftp_dispatch_pro). and restart this batch.")
			return
		}
		if st == "4" {
			this.lock.RLock()
			this.status = 4
			this.lock.RUnlock()
			this.log.Debug("stop batch.", this.dispatchId, as_of_date)
			return
		}
		if st == "3" {
			this.lock.RLock()
			this.status = 3
			this.lock.RUnlock()
			this.log.Debug("this batch complete")
			return
		}

		this.lock.RLock()
		if this.status != 0 {
			this.lock.RUnlock()
			this.log.Debug("batch listener exit")
			return
		}
		this.lock.RUnlock()

		time.Sleep(time.Second * 2)
	}
}

func (this *CalcFTP) baseCalculateByParallel(val acctInfo, m busizMethod) (string, error) {

	ftpRate := ""
	switch m.Ftp_method_id {
	case "101":
		cm, ok := this.cv[m.Curve_id]
		if ok == false {
			this.log.Error("账户：", val.account_number, "配置中指定的曲线", m.Curve_id, "在曲线信息表中无法获取。请清查曲线是否有值")
			return "", errors.New("无法匹配曲线")
		}
		date := this.getDate(val)
		ci, err := this.matchCurve(date, cm)
		if err != nil {
			this.log.Error(err)
			return "", err
		}
		ftpRate = linearInterpolation(date, m.Term_cd, m.Term_cd_mult, ci, this.ct)
	case "102":
		cm := this.cv[m.Curve_id]
		date := this.getDate(val)
		endDate := this.getEndDate(val)
		ci, err := this.matchCurve(date, cm)
		if err != nil {
			this.log.Error(err)
			return "", err
		}
		termD := utils.DataInterval(date, endDate)
		//
		//基础价格计算
		ftpRate = linearInterpolation(date, strconv.Itoa(termD), "D", ci, this.ct)

	case "103":
		date := this.getDate(val)
		cm := this.cv[m.Curve_id]
		ci, err := this.matchCurve(date, cm)

		if err != nil {
			this.log.Error(err)
			return "", err
		}

		//
		// 基础价格计算
		ftpRate = calcFlowStart(val, ci, this.ct, this.log)

	case "104":
		termWeight := this.precipitation[val.busiz_id]
		ftpR := float64(0)
		date := this.getDate(val)
		cm := this.cv[m.Curve_id]
		ci, err := this.matchCurve(date, cm)
		if err != nil {
			this.log.Error(err.Error(), val.account_number, "This account calculate precipitation failed")
			return "", errors.New("This account calculate precipitation failed")
		}
		for _, tw := range termWeight {
			ftprate := linearInterpolation(val.origination_date, tw.term_cd, tw.term_cd_mult, ci, this.ct)
			ftprat, err := strconv.ParseFloat(ftprate, 8)
			if err != nil {
				this.log.Error(err.Error(), val.account_number, "This account calculate precipitation failed")
				return "", errors.New("This account calculate precipitation failed")
			}
			ftweight, err := strconv.ParseFloat(tw.weight, 8)
			if err != nil {
				this.log.Error(err.Error(), val.account_number, "This account calculate precipitation failed")
				return "", errors.New("This account calculate precipitation failed")
			}
			ftpR += ftprat * ftweight
		}

		//
		// 基础价格计算
		ftpRate = strconv.FormatFloat(ftpR, 'f', 6, 64)

	case "105":
		date := this.getDate(val)
		cm := this.cv[m.Curve_id]
		ci, err := this.matchCurve(date, cm)

		if err != nil {
			this.log.Error(err)
			return "", err
		}
		frd := calcDurationFlowStart(val, ci, this.ct, this.log)

		//
		// 基础价格计算
		ftpRate = linearInterpolation(date, frd, "D", ci, this.ct)

	case "106":
		c, _ := strconv.ParseFloat(val.cur_net_rate, 6)
		p, _ := strconv.ParseFloat(m.Point_val, 6)

		// 基础价格计算
		ftpRate = strconv.FormatFloat(c+p/100, 'f', 6, 32)

	default:
		this.log.Warn(m.Ftp_method_id, "定价方法无效")
		return "", errors.New("没有匹配上定价方法")
	}
	return ftpRate, nil
}

func (this *CalcFTP) adjustCalculateByParallel(val acctInfo, m busizMethod, ftpRate string) map[string]string {
	return this.adjust.calcAdjustRate(this, val, m, ftpRate)
}

func (this *CalcFTP) calculateByParallel(val acctInfo, m busizMethod) {
	//匹配上定价单元
	//根据定价单元查找定价方法
	defer func() {
		<-this.RUNTIME_CNT
		if r := recover(); r != nil {
			this.log.Error("计算错误", r)
			this.lock.Lock()
			this.errorCnt++
			this.lock.Unlock()
			//this.errorCnt = atomic.AddInt64(&this.errorCnt, 1)
		}
	}()

	ftpRate, err := this.baseCalculateByParallel(val, m)
	if err != nil {
		this.log.Error(err)
		this.lock.Lock()
		this.errorCnt++
		this.lock.Unlock()
		//this.errorCnt = atomic.AddInt64(&this.errorCnt, 1)
		return
	}
	adj := this.adjustCalculateByParallel(val, m, ftpRate)
	go this.sendFtpRate(val.domain_id, val.as_of_date, val.account_number, ftpRate, adj)

}

//
// get start date
func (this *CalcFTP) getDate(acct acctInfo) string {
	rateType := acct.adjustable_type_cd
	date := ""
	switch rateType {
	case "0":
		date = acct.origination_date
	case "50":
		date = acct.origination_date
	case "250":
		date = acct.last_reprice_date
	default:
		date = acct.origination_date
	}
	return date
}

// getEndDate function:
// get the next_reprice_date. if adjustable_type_cd is fixed, return matury_date
// if adjustable_type_cd is float, return next_reprice_date.
func (this *CalcFTP) getEndDate(val acctInfo) string {
	rateType := val.adjustable_type_cd
	date := ""
	switch rateType {
	case "0":
		date = val.maturity_date
	case "50":
		date = val.maturity_date
	case "250":
		date = val.next_reprice_date
	default:
		date = val.maturity_date
	}
	return date
}

// 匹配值信息
// 如果指定日期当天没有曲线值，则匹配之前最近的曲线
// 如果指定日期之前已经没有曲线值，则匹配最早的一条曲线
// 遇到年初没有曲线情况时，向前推移至上一年取值
func (this *CalcFTP) matchCurve(date string, cm *curve) ([]curveInfo, error) {
	//判断曲线信息表是否为空
	//如果为空，直接退出，无需继续匹配曲线
	if cm == nil {
		return nil, errors.New("no curve info in mas_curve_info")
	}
	if !utils.ValidDate(date) || !utils.ValidDate(cm.oldestdate) {
		return nil, errors.New("日期格式不符合要求" + date + "," + cm.oldestdate)
	}
	if utils.DataInterval(date, cm.oldestdate) >= 0 {
		if len(cm.oldest_curve) > 0 {
			return cm.oldest_curve, nil
		} else {
			return cm.oldest_curve, errors.New("no curve find in mas_curve_info")
		}
	}
	year := date[0:4]
	var rst []curveInfo
	var readflag = false
	if c, ok := cm.value[year]; ok {
		var hzw = true
		var yph = ""
		for _, val := range c {
			if utils.CompareDate(date, val.as_of_date) > 0 {
				continue
			} else if utils.CompareDate(date, val.as_of_date) == 0 {
				hzw = false
				readflag = true
				rst = append(rst, val)
			} else {
				if hzw == true {
					yph = val.as_of_date
					hzw = false
				}
				if val.as_of_date == yph {
					readflag = true
					rst = append(rst, val)
				} else {
					break
				}
			}
		}
	}
	if readflag == false {
		oldestYear, _ := strconv.Atoi(cm.oldestdate[0:4])
		y, _ := strconv.Atoi(year)
		for oldestYear <= y && readflag == false {
			if c, ok := cm.value[strconv.Itoa(y-1)]; ok {
				var hzw = true
				var yph = ""
				for _, val := range c {
					if hzw == true {
						yph = val.as_of_date
						hzw = false
					}
					if val.as_of_date == yph {
						readflag = true
						rst = append(rst, val)
					} else {
						break
					}
				}
			}
			y = y - 1
		}
	}
	return rst, nil
}

func (this *CalcFTP) sendFtpRate(v1, v2, v3, v4 string, adj map[string]string) {
	var one ftpResult
	one.domain_id = v1
	one.as_of_date = v2
	one.account_number = v3
	one.ftp_rate = v4
	one.ajust_01 = adj["601"]
	one.ajust_02 = adj["604"]
	one.ajust_03 = adj["603"]
	if v4 != "" {
		this.rstCache <- one

		atomic.AddInt64(&this.successCnt, 1)

		//this.successCnt = atomic.AddInt64(&this.successCnt, 1)
	} else {
		atomic.AddInt64(&this.errorCnt, 1)
		//this.errorCnt = atomic.AddInt64(&this.errorCnt, 1)
	}
}

func (this *CalcFTP) updateFtpRate() {
	this.log.Info(this.dispatchId, "-> Dest output table is ", this.outputSrc)
	var index = 0
	var tmp []ftpResult
	var rst []ftpResult
	for {
		select {
		case val := <-this.rstCache:
			rst = append(rst, val)
			if index++; index > 500 {
				index = 0
				go this.insert(rst)
				rst = tmp
			}
		case <-time.After(time.Second * 2):
			if index > 0 {
				index = 0
				go this.insert(rst)
				rst = tmp
			}
			this.lock.RLock()
			if len(this.rstCache) == 0 && this.status != 0 {
				this.lock.RUnlock()
				this.log.Info("批次运行结束，退出数据更新服务进程。")
				fmt.Println("批次运行结束，退出数据更新服务进程。")
				return
			}
			this.lock.RUnlock()
		}
	}
	this.log.Close()
}

func (this *CalcFTP) insert(rst []ftpResult) {
	this.wg.Add(1)
	sql := strings.Replace(FTPCALC_INSERT_RESULT, "HZWY23", this.outputSrc, -1)
	tx, _ := dbobj.Default.Begin()
	for _, val := range rst {
		tx.Exec(sql,
			val.domain_id,
			val.as_of_date,
			val.account_number,
			val.ftp_rate,
			val.ajust_01,
			val.ajust_02,
			val.ajust_03,
			val.ajust_04)
	}
	tx.Commit()
	dbobj.Default.Exec("update ftp_dispatch_pro set cur_rows = :1 where dispatch_id = :2 and domain_id = :3 and dispatch_date = to_date(:4,'YYYY-MM-DD')", this.successCnt, this.dispatchId, this.domain_id, this.as_of_date)
	this.wg.Done()
}
