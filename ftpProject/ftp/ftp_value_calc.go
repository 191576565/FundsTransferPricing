package ftp

import (
	"encoding/json"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
	"net/rpc"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/astaxie/beego/config"
)

type FtpValueCalc struct {
	AccNumber         string //  随机
	AsOfRate          string //  x
	BusinessId        string //  y
	OriginalDate      string //  y
	MaturityDate      string //  y
	OrgTerm           string //  maturydate - originaldate
	OrgTermMult       string //  D
	IsoCurrencyCd     string //  y
	AdjustType        string //  y
	OrgParBal         string // y
	OrgPaymentAmt     string // y
	CurNetRate        string //  y
	InterestMode      string // y
	PayInterestMode   string //y
	AmartTypeCd       string //
	PaymentFreq       string //
	PaymentFreqMult   string
	LastPaymentDate   string //
	NextPaymentDate   string
	LastBalacne       string
	LastRepriceDate   string
	NextRepriceDate   string
	Remain_no_pmts_c  string //
	Reprice_freq      string //
	Reprice_freq_mult string //
	FtpRate           string //
	DomainID          string
}
type FtpValueCalcCtl struct {
	RouteControl
}
type FtpValueCalcPage struct {
	RouteControl
}
type ACCTInfo struct {
	Account_number     string
	As_of_date         string
	Busiz_id           string // y
	Origination_date   string // y
	Maturity_date      string // y
	Org_term           string // t
	Org_term_mult      string // t
	Iso_currency_cd    string // y
	Adjustable_type_cd string // y
	Org_par_bal        string // y
	Org_payment_amt    string // y
	Cur_payment        string // x
	Cur_net_rate       string // y
	Accrual_basis_cd   string // y
	Amart_type_cd      string // y
	Pmt_freq           string // y
	Pmt_freq_mult      string // y
	Last_payment_date  string // x
	Next_payment_date  string // y
	Remain_no_pmts_c   string // x
	Reprice_freq       string // y
	Reprice_freq_mult  string // y
	Lrd_balance        string // y = org_par_bal
	Last_reprice_date  string // t
	Next_reprice_date  string // y
	FtpRate            string
	Adjone             string
	Adjtwo             string
	Adjthree           string
	Domain_id          string //
}

type InsideAdjust struct {
	Adjustname string
	Adjustvale string
}
type FtpCalcResult struct {
	ErrorCode  string
	ErrorMsg   string
	FtpValue   string
	Insideinfo []InsideAdjust
}

func (this *FtpValueCalcPage) Get() {
	this.TplName = "mas/ftp/ftp_ValueCalc.tpl"
}

type BusizInfoCalc struct {
	Busiz_id    string
	Busiz_desc  string
	Al_flag     string
	Method_id   string
	Method_desc string
}
type BusizInfoCalcCtl struct {
	RouteControl
}
type RateAdjustType struct {
	Adjustable_type_cd   string
	Adjustable_type_desc string
}
type RateAdjustTypeCtl struct {
	RouteControl
}
type AccrualCdAttr struct {
	Accrual_basis_cd   string
	Accrual_basis_desc string
}
type AccrualCdAttrCtl struct {
	RouteControl
}
type PaymentTypeAttr struct {
	Amrt_type_cd   string
	Amrt_type_desc string
}
type PaymentTypeAttrCtl struct {
	RouteControl
}

func (this *BusizInfoCalcCtl) Get() {

	w := this.Ctx.ResponseWriter
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	sql := FTP_VALUECALC_GET1
	rows, err := dbobj.Default.Query(sql, doName)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询业务单元失败"))
		return
	}
	var one BusizInfoCalc
	var all []BusizInfoCalc
	for rows.Next() {

		err := rows.Scan(&one.Busiz_id, &one.Busiz_desc, &one.Al_flag, &one.Method_id, &one.Method_desc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值业务单元失败"))
			return
		}
		all = append(all, one)
	}
	ojs, err := json.Marshal(all)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *RateAdjustTypeCtl) Get() {
	w := this.Ctx.ResponseWriter
	//	if sys.Privilege.Access(w, r) == false {
	//		w.WriteHeader(http.StatusForbidden)
	//		return
	//	}
	//r.ParseForm()
	//取利率调整方式
	sql := FTP_VALUECALC_GET2
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询利率调整方式失败"))
		return
	}
	var one RateAdjustType
	var all []RateAdjustType
	for rows.Next() {
		err := rows.Scan(&one.Adjustable_type_cd, &one.Adjustable_type_desc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值利率调整方式失败"))
			return
		}
		all = append(all, one)
	}
	ojs, err := json.Marshal(all)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *AccrualCdAttrCtl) Get() {
	w := this.Ctx.ResponseWriter
	//	if sys.Privilege.Access(w, r) == false {
	//		w.WriteHeader(http.StatusForbidden)
	//		return
	//	}
	//r.ParseForm()
	//利息计提
	sql := FTP_VALUECALC_GET3
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询利息计提方式失败"))
		return
	}
	var one AccrualCdAttr
	var all []AccrualCdAttr
	for rows.Next() {
		err = rows.Scan(&one.Accrual_basis_cd, &one.Accrual_basis_desc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值利息计提方式失败"))
			return
		}
		all = append(all, one)
	}
	ojs, err := json.Marshal(all)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *PaymentTypeAttrCtl) Get() {
	//	if sys.Privilege.Access(w, r) == false {
	//		w.WriteHeader(http.StatusForbidden)
	//		return
	//	}
	//r.ParseForm()
	//偿还
	w := this.Ctx.ResponseWriter
	sql := FTP_VALUECALC_GET4
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询利息计提方式失败"))
		return
	}
	var one PaymentTypeAttr
	var all []PaymentTypeAttr
	for rows.Next() {
		err := rows.Scan(&one.Amrt_type_cd, &one.Amrt_type_desc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值利息计提方式失败"))
			return
		}
		all = append(all, one)
	}
	ojs, err := json.Marshal(all)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *FtpValueCalcCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var result FtpCalcResult
	//取数据
	var tmp ACCTInfo
	methodid := r.FormValue("MethodID")

	tmp.Busiz_id = r.FormValue("BusinessId") //业务单元ok
	//fmt.Println(tmp.Busiz_id)
	tmp.Origination_date = r.FormValue("OriginalDate")     //起息日ok
	tmp.Maturity_date = r.FormValue("MaturityDate")        //到期日ok
	tmp.Iso_currency_cd = r.FormValue("IsoCurrencyCd")     //币种ok
	tmp.Adjustable_type_cd = r.FormValue("AdjustType")     //利率调整方式ok
	tmp.Org_par_bal = r.FormValue("OrgParBal")             //票面金额=原始支付金额ok
	tmp.Org_payment_amt = r.FormValue("OrgPaymentAmt")     //首次支付金额ok
	tmp.Cur_net_rate = r.FormValue("CurNetRate")           //执行利率ok
	tmp.Accrual_basis_cd = r.FormValue("InterestMode")     //计息方式ok利息计提
	tmp.Amart_type_cd = r.FormValue("PayInterestMode")     //还款方式ok支付方式
	tmp.Pmt_freq = r.FormValue("PaymentFreq")              //支付频率ok
	tmp.Pmt_freq_mult = r.FormValue("PaymentFreqMult")     //支付频率单位ok
	tmp.Next_payment_date = r.FormValue("NextPaymentDate") //下次支付日=首次ok
	tmp.Lrd_balance = r.FormValue("OrgParBal")             //r.FormValue("LastBalacne")  ok         //下次重定价余额

	tmp.Next_reprice_date = r.FormValue("NextRepriceDate") //下次重定价日期=首次ok
	//tmp.Domain_id = r.FormValue("DomainID")                //与编码
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	tmp.Domain_id = doName
	tmp.Reprice_freq = r.FormValue("RepriceFreq")          //重定价频率ok
	tmp.Reprice_freq_mult = r.FormValue("RepriceFreqMult") //重定价频率单位ok
	tmp.Account_number = "622765"
	tmp.As_of_date = time.Now().Format("2006-01-02")
	//2016.8.19
	if methodid == "101" || methodid == "104" || methodid == "106" {
		if tmp.Origination_date == "" {
			tmp.Origination_date = time.Now().Format("2006-01-02")
		}

	} else {
		if !utils.ValidDate(tmp.Origination_date) {
			result.ErrorCode = "0"
			result.ErrorMsg = "起息日错误"
			ojs, err := json.Marshal(result)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
		if !utils.ValidDate(tmp.Maturity_date) {
			result.ErrorCode = "0"
			result.ErrorMsg = "到期日错误"
			ojs, err := json.Marshal(result)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
		nday := utils.DataInterval(tmp.Origination_date, tmp.Maturity_date)
		if nday < 0 {
			result.ErrorCode = "0"
			result.ErrorMsg = "到期日应大于起息日"
			ojs, err := json.Marshal(result)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
		tmp.Org_term = strconv.Itoa(utils.DataInterval(tmp.Origination_date, tmp.Maturity_date))
		tmp.Org_term_mult = "D"
		//
		num, _ := strconv.Atoi(tmp.Reprice_freq)
		fmt.Println("num", num)
		var targetdate string
		if tmp.Next_reprice_date != "" {
			//fmt.Println("Reprice_freq_mult ", tmp.Reprice_freq_mult)
			switch tmp.Reprice_freq_mult {
			case "D":
				targetdate, _ = utils.AddDays(tmp.Next_reprice_date, num*(-1))
			case "M":
				targetdate, _ = utils.AddMonths(tmp.Next_reprice_date, num*(-1))
			case "Y":
				fmt.Println("come")
				targetdate, _ = utils.AddMonths(tmp.Next_reprice_date, num*(-12))
			default:
			}
			//fmt.Println("Next_reprice_date ", tmp.Next_reprice_date)
			//fmt.Println("targetdate ", targetdate)
			nday := utils.DataInterval(tmp.Origination_date, targetdate)
			fmt.Println(nday)
			if nday >= 0 {
				tmp.Last_reprice_date = targetdate
			} else {
				tmp.Last_reprice_date = tmp.Origination_date
			}
			fmt.Println("shangci chongdingjia :", tmp.Last_reprice_date)
		}
	}
	var args []ACCTInfo
	var reply []ACCTInfo
	//	var ftpsinglecalc ftpCalc.RpcSrvOfFTP

	args = append(args, tmp)
	//
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")

	red, err := config.NewConfig("ini", appConfigPath)
	if err != nil {
		logs.Error("cant not read ./conf/app.conf .please check this file.")
		result.ErrorCode = "0"
		result.ErrorMsg = "读取计算引擎地址出错，请联系管理员检查计算引擎配置文件"
		ojs, err := json.Marshal(result)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	calcip := red.String("Calc.IP")
	//
	conn, err := rpc.DialHTTP("tcp", calcip)
	if err != nil {

		result.ErrorCode = "0"
		result.ErrorMsg = "与引擎服务建立连接失败,请联系管理员检查引擎服务配置和引擎是否启动"
		ojs, err := json.Marshal(result)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	err = conn.Call("RpcSrvOfFTP.FtpTrial", args, &reply)
	if err != nil {
		logs.Error(err)
		result.ErrorCode = "0"
		result.ErrorMsg = "调用引擎服务失败,请联系管理员检查系统日志"
		ojs, err := json.Marshal(result)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	if reply == nil {
		fmt.Println("服务返回值为空")
		result.ErrorCode = "0"
		result.ErrorMsg = "服务返回值为空,请检查参数"
		ojs, err := json.Marshal(result)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return

	}
	//
	opcontent := "单笔试算业务单元编码为:" + tmp.Busiz_id
	this.InsertLogToDB(valuecalcrun, opcontent, myapp)
	//

	result.ErrorCode = "1"
	result.ErrorMsg = "计算成功"
	result.FtpValue = reply[0].FtpRate
	var single InsideAdjust
	//返回数据
	if reply[0].Adjone != "" {
		single.Adjustname = "期限流动性溢价"
		single.Adjustvale = reply[0].Adjone
		//fmt.Println("调节项1", reply[0].Adjone)
		result.Insideinfo = append(result.Insideinfo, single)
	}
	if reply[0].Adjtwo != "" {
		single.Adjustname = "准备金调节"
		single.Adjustvale = reply[0].Adjtwo
		//fmt.Println("调节项2", reply[0].Adjtwo)
		result.Insideinfo = append(result.Insideinfo, single)
	}
	if reply[0].Adjthree != "" {
		single.Adjustname = "司库利润还原调节"
		single.Adjustvale = reply[0].Adjthree
		//fmt.Println("调节项2", reply[0].Adjthree)
		result.Insideinfo = append(result.Insideinfo, single)
	}
	ojs, err := json.Marshal(result)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
