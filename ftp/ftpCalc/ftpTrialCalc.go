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
	"errors"
	"mas/ftp/calcLog"
	"runtime"
	"strconv"
	"sync"
	"time"
	"utils"
)

//
//
type TrialFtpCalc struct {
	precipitation map[string][]precipitation
	bm            map[string]busizMethod
	cv            map[string]*curve
	ct            map[string]curveStruct
	adjust        *ftpAdjustment
	log           calcLog.Calclog
	msg           string
	calc          *CalcFTP
}

func (this *TrialFtpCalc) param(at []acctInfo, domain_id string) error {
	/*
	* 手工开启多个核心,防止低版本gosdk无法自动启动多核
	 */
	var MULTICORE int = runtime.NumCPU() //number of core
	runtime.GOMAXPROCS(MULTICORE)        //running in multicore

	if !utils.ValidWord(domain_id, 1, 30) {
		this.calc.log.Error("域名格式不正确。请输入1-30位英文字母，数字，下滑线组合的字符串")
		return errors.New("域名格式不正确。请输入1-30位英文字母，数字，下滑线组合的字符串")
	}

	// init calculate ftp object
	//
	this.calc = new(CalcFTP)
	this.calc.wg = new(sync.WaitGroup)
	this.calc.status = 0

	this.calc.log = calcLog.NewBatchLog("", "", -1)
	this.log = this.calc.log

	this.calc.lock = new(sync.RWMutex)
	this.calc.dispatchId = "ftp trial calculate."
	// init sync
	this.calc.wg.Add(5)

	// 清除ftp定价结果集数据,取消自动删除结果集功能
	// 需要手动删除结果集
	// ClearFtpRate()

	go func() {
		this.calc.bm = GetBusizMethod(domain_id, this.calc)
		this.bm = this.calc.bm
	}()

	//获取最近几年曲线，
	//每一年按照一个区间进行分区处理
	go func() {
		this.calc.cv = GetCurveVal(domain_id, "3000", this.calc)
		this.cv = this.calc.cv
	}()

	//获取所有期限点信息
	go func() {
		this.calc.ct = GetCurveStruct(domain_id, this.calc)
		this.ct = this.calc.ct
	}()

	//获取偿还曲线配置信息
	go func() {
		this.calc.precipitation = getPrecipitation(domain_id, this.calc)
		this.precipitation = this.calc.precipitation
	}()

	go func() {
		this.calc.adjust = ftpAdjustInit(domain_id, this.calc, time.Now().Format("2006-01-02"))
		this.adjust = this.calc.adjust
	}()

	// waiting for init process.
	//
	this.calc.wg.Wait()

	if this.calc.status != 0 {
		this.calc.log.Error("init batch config info failed. please check log info and analysic problem.")
		return errors.New("init rules failed.")
	}
	return nil
}

// Function: FTP试算接口
// Author: 黄占威
// Time: 2016-07-31
//
func (this *TrialFtpCalc) FtpTrial(at []acctInfo, domain_id string) ([]acctInfo, error) {
	defer func() {
		if r := recover(); r != nil {
			this.msg = "批次运行失败"
			this.log.Error("批次运行初始化")
		}
	}()

	// init rules.
	// get curve info, business info,adjustment info.
	err := this.param(at, domain_id)
	if err != nil {
		this.log.Error(err)
		return nil, err
	}

	var rst []acctInfo
	var one acctInfo

	for _, val := range at {
		biz := val.busiz_id
		if m, ok := this.bm[biz]; ok {
			ftpRate, adj, err := this.calculate(val, m)
			if err != nil {
				this.log.Error(err)
				return nil, err
			}
			one.account_number = val.account_number
			one.accrual_basis_cd = val.accrual_basis_cd
			one.adjustable_type_cd = val.adjustable_type_cd
			one.amart_type_cd = val.amart_type_cd
			one.as_of_date = val.as_of_date
			one.busiz_id = val.busiz_id
			one.cur_net_rate = val.cur_net_rate
			one.cur_payment = val.cur_payment
			one.domain_id = val.domain_id
			one.ftpRate = ftpRate
			one.adjone = adj["601"]
			one.adjtwo = adj["604"]
			one.adjthree = adj["603"]
			one.iso_currency_cd = val.iso_currency_cd
			one.last_payment_date = val.last_payment_date
			one.last_reprice_date = val.last_reprice_date
			one.lrd_balance = val.lrd_balance
			one.maturity_date = val.maturity_date
			one.next_payment_date = val.next_payment_date
			one.next_reprice_date = val.next_reprice_date
			one.org_par_bal = val.org_par_bal
			one.org_payment_amt = val.org_payment_amt
			one.org_term = val.org_term
			one.org_term_mult = val.org_term_mult
			one.origination_date = val.origination_date
			one.pmt_freq = val.pmt_freq
			one.pmt_freq_mult = val.pmt_freq_mult
			one.remain_no_pmts_c = val.remain_no_pmts_c
			one.reprice_freq = val.reprice_freq
			one.reprice_freq_mult = val.reprice_freq_mult

			rst = append(rst, one)
		} else {

			one.account_number = val.account_number
			one.accrual_basis_cd = val.accrual_basis_cd
			one.adjustable_type_cd = val.adjustable_type_cd
			one.amart_type_cd = val.amart_type_cd
			one.as_of_date = val.as_of_date
			one.busiz_id = val.busiz_id
			one.cur_net_rate = val.cur_net_rate
			one.cur_payment = val.cur_payment
			one.domain_id = val.domain_id
			one.ftpRate = ""
			one.adjone = ""
			one.adjtwo = ""
			one.adjthree = ""
			one.iso_currency_cd = val.iso_currency_cd
			one.last_payment_date = val.last_payment_date
			one.last_reprice_date = val.last_reprice_date
			one.lrd_balance = val.lrd_balance
			one.maturity_date = val.maturity_date
			one.next_payment_date = val.next_payment_date
			one.next_reprice_date = val.next_reprice_date
			one.org_par_bal = val.org_par_bal
			one.org_payment_amt = val.org_payment_amt
			one.org_term = val.org_term
			one.org_term_mult = val.org_term_mult
			one.origination_date = val.origination_date
			one.pmt_freq = val.pmt_freq
			one.pmt_freq_mult = val.pmt_freq_mult
			one.remain_no_pmts_c = val.remain_no_pmts_c
			one.reprice_freq = val.reprice_freq
			one.reprice_freq_mult = val.reprice_freq_mult

			rst = append(rst, one)
			this.log.Warn(biz, "没有匹配上定价单元")
		}
	}
	return rst, nil
}

func (this *TrialFtpCalc) baseCalculateByParallel(val acctInfo, m busizMethod) (string, error) {

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

func (this *TrialFtpCalc) adjustCalculateByParallel(val acctInfo, m busizMethod, ftpRate string) map[string]string {
	return this.adjust.calcAdjustRate(this.calc, val, m, ftpRate)
}

func (this *TrialFtpCalc) calculate(val acctInfo, m busizMethod) (string, map[string]string, error) {
	//匹配上定价单元
	//根据定价单元查找定价方法

	defer func() {
		if r := recover(); r != nil {
			this.log.Error("FTP定价引擎计算错误", r)
		}
	}()

	ftpRate, err := this.baseCalculateByParallel(val, m)
	if err != nil {
		this.log.Error(err)
		return "", nil, err
	}
	adj := this.adjustCalculateByParallel(val, m, ftpRate)

	return ftpRate, adj, nil

}

func (this *TrialFtpCalc) getDate(acct acctInfo) string {
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
func (this *TrialFtpCalc) getEndDate(val acctInfo) string {
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
func (this *TrialFtpCalc) matchCurve(date string, cm *curve) ([]curveInfo, error) {

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
