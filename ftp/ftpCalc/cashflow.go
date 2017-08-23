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
	"mas/ftp/calcLog"
	"math"
	"strconv"
	"utils"
)

// every cashflow values
// ftp:= sum(pv*t*r)/sum(pv*t)
type cash struct {
	pv float64
	r  float64
	t  float64
}

// acct means account infos
// curve  means curve infos
// ct means curve struct infos
type cashFlow struct {
	acct  acctInfo
	curve []curveInfo
	ct    map[string]curveStruct
	log   calcLog.Calclog
}

func calcFlowStart(acct acctInfo, curve []curveInfo, ct map[string]curveStruct, log calcLog.Calclog) string {
	//new instance for every cashflow calcuate
	r := new(cashFlow)
	//variable acct save one account info
	r.acct = acct
	//variable curve save curveinfo
	r.curve = curve
	//variable ct save curve struct node
	r.ct = ct

	r.log = log

	//variable start_date:
	//when adjustable_type_cd is fixed
	//start_date equals origination_date
	//end_date equals matury_date
	//when adjustable_type_cd is float
	//start_date equals last_reprice_date
	//end_date equals next_reprice_date
	start_date := r.getDate()
	end_date := r.getEndDate()

	// 800 means average capital
	// 700 means average capital plus interest
	// 500 means installment payments due debt
	// 600 means And with this clear

	switch r.acct.amart_type_cd {
	case "900":
		return r.cashFlow_900("", start_date, end_date)
	case "800":
		// get the first payment_date of this reprice duration
		// variable st is the value of first payment_date
		st := r.cashFlow_getTOfFirstDate(start_date, 0)
		return r.cashFlow_800(st, start_date, end_date)
	case "700":
		// get the first payment_date of this reprice duration
		// variable st is the value of first payment_date
		st := r.cashFlow_getTOfFirstDate(start_date, 0)
		return r.cashFlow_700(st, start_date, end_date)
	case "600":
		// get the first payment_date of this reprice duration
		// variable st is the value of first payment_date
		term, mult := r.getTerm()
		return linearInterpolation(start_date, term, mult, r.curve, r.ct)
	case "500":
		// get the first payment_date of this reprice duration
		// variable st is the value of first payment_date
		st := r.cashFlow_getTOfFirstDate(start_date, 0)
		return r.cashFlow_500(st, start_date, end_date)
	case "400":
		// get the first payment_date of this reprice duration
		// variable st is the value of first payment_date
		st := r.cashFlow_getTOfFirstDate(start_date, 0)
		return r.cashFlow_400(st, start_date, end_date)
	case "300":
		// get the first payment_date of this reprice duration
		// variable st is the value of first payment_date
		st := r.cashFlow_getTOfFirstDate(start_date, 0)
		return r.cashFlow_300(st, start_date, end_date)
	}
	return ""
}

func (this *cashFlow) getTerm() (string, string) {
	term := ""
	mult := ""
	switch this.acct.adjustable_type_cd {
	case "0":
		term = this.acct.org_term
		mult = this.acct.org_term_mult
	case "50":
		term = this.acct.org_term
		mult = this.acct.org_term_mult
	case "250":
		term = this.acct.reprice_freq
		mult = this.acct.reprice_freq_mult
	}
	return term, mult
}

// getDate function :
// get the last_reprice_date. if adjustable_type_cd is fixed. return origination_date
// if adjust_type_cd is float , return last_reprice_date
func (this *cashFlow) getDate() string {
	rateType := this.acct.adjustable_type_cd
	date := ""
	switch rateType {
	case "0":
		date = this.acct.origination_date
	case "50":
		date = this.acct.origination_date
	case "250":
		date = this.acct.last_reprice_date
	default:
		date = this.acct.origination_date
	}
	return date
}

// getEndDate function:
// get the next_reprice_date. if adjustable_type_cd is fixed, return matury_date
// if adjustable_type_cd is float, return next_reprice_date.
func (this *cashFlow) getEndDate() string {
	rateType := this.acct.adjustable_type_cd
	date := ""
	switch rateType {
	case "0":
		date = this.acct.maturity_date
	case "50":
		date = this.acct.maturity_date
	case "250":
		date = this.acct.next_reprice_date
	default:
		date = this.acct.maturity_date
	}
	return date
}

// cashFlow_getTOfFirstDate function:
// this function get 2 args.
// first arg is start_date,second arg is flag.
// when flag == 0, then this function work, or not.
// this funciton return a date, which is the first paymet_date of this reprice duration
func (this *cashFlow) cashFlow_getTOfFirstDate(last_reprice_date string, sflag int) string {

	// variable rst save the next_payment_date
	// variable tmp save the next_payment_date newer than rst
	// if tmp less than next_payment_date of acct. then copy tmp value to rst
	// if tmp granter than next_payment_date of acct. then break this circle.
	rst := this.acct.next_payment_date
	tmp := this.acct.next_payment_date

	// pmt_freq is equal the pay_freq of the acct,
	// common . then pmt_freq_mult is equal "M"
	num, _ := strconv.Atoi(this.acct.pmt_freq)
	pmt_freq_mult := this.acct.pmt_freq_mult

	// when tmp granter than last_reprice_date
	// circle
	for utils.CompareDate(tmp, last_reprice_date) <= 0 {
		rst = tmp
		switch pmt_freq_mult {
		case "D":
			tmp, _ = utils.AddDays(tmp, num*(-1))
		case "M":
			tmp, _ = utils.AddMonths(tmp, num*(-1))
		case "Y":
			tmp, _ = utils.AddMonths(tmp, num*(-12))
		default:
		}
	}

	// modify last_payment_date
	// if last_reprice_date is equal to origination_date
	// if the month of last_payment_date is equal last_reprice_date, last_payment_date add pmt_freq month
	// if the month of last_payment_date is granter than the month of last_reprice_date ,but the day of last_payment_date less than the day of last_reprice_date
	// if sflag is equal to 0 , then last_payment_date add pmt_freq month
	if (this.acct.adjustable_type_cd == "0" || last_reprice_date == this.acct.origination_date) && sflag == 0 {
		switch pmt_freq_mult {
		case "D":
			newrst, err := utils.AddDays(rst, num)
			if err != nil {
				this.log.Error(err)
				return rst
			}
			if utils.CompareDate(newrst, this.acct.next_payment_date) >= 0 {
				return newrst
			} else {
				return rst
			}
		case "M":
			newrst, err := utils.AddMonths(rst, num)
			if err != nil {
				this.log.Error(err)
				return rst
			}
			if utils.CompareDate(newrst, this.acct.next_payment_date) >= 0 {
				return newrst
			} else {
				return rst
			}
		case "Y":
			newrst, err := utils.AddMonths(rst, 12*num)
			if err != nil {
				this.log.Error(err)
				return rst
			}
			if utils.CompareDate(newrst, this.acct.next_payment_date) >= 0 {
				return newrst
			} else {
				return rst
			}
		}
		return rst
	} else {
		return rst
	}
}

//getACCDays function:
//this function get 2 args.first arg is current payment date. second arg is flag ,when flag == 0 ,this funcion work ,or not.
//this funcion return 2 values.
//first return values is days,which is this accrual duation.
func (this *cashFlow) getACCDays(date string, sflag int) (int, string) {
	rst := 0
	next_date := ""
	//将支付频率换成数字行
	num, _ := strconv.Atoi(this.acct.pmt_freq)

	//根据支付频率单位，判断支付周期单位
	switch this.acct.pmt_freq_mult {
	case "D":
		lt, _ := utils.AddMonths(date, (-1)*num)
		rst = utils.DataInterval(lt, date)
		next_date, _ = utils.AddDays(date, num)
	case "M":
		rst, _ = utils.MonthInterval(date, (-1)*num)
		lltt, _ := utils.AddMonths(date, (-2)*num)
		if utils.CompareDate(lltt, this.acct.origination_date) > 0 {
			rst = utils.DataInterval(this.acct.origination_date, date)
		}
		next_date, _ = utils.AddMonths(date, num)
	case "Y":
		rst, _ = utils.MonthInterval(date, (-12)*num)
		lltt, _ := utils.AddMonths(date, (-24)*num)
		if utils.CompareDate(lltt, this.acct.origination_date) > 0 {
			rst = utils.DataInterval(this.acct.origination_date, date)
		}
		next_date, _ = utils.AddMonths(date, 12*num)
	default:
	}
	return rst, next_date
}

func (this *cashFlow) cashFlow_900(t_first string, last_reprice_date string, next_reprice_date string) string {

	var one cash
	var rst []cash
	//还款日当年总天数
	var ACC_TOTAL = 0
	//本次还款，计息累计天数
	var ACC_PMT_TOTAL = 0

	//折现天数
	var DISCOUNTING_DAYS = 0
	//折现当年，本年总天数
	var DISCOUNTING_TOTAL_YEAR = 0
	//本次重定价周期内还款次数
	var INDEX = 0

	//获取上次重定价日余额
	var LRD_BALANCE, _ = strconv.ParseFloat(this.acct.lrd_balance, 8)
	//获取本期本金信息
	var ORG_PAYMENT_AMT = float64(0) //strconv.ParseFloat(this.acct.org_payment_amt, 8)
	//获取业务执行利率
	var CUR_NET_RATE, _ = strconv.ParseFloat(this.acct.cur_net_rate, 8)

	CUR_NET_RATE = CUR_NET_RATE / 100
	//客户支付信息
	var CustomPaymentInfo = getCustomPaymentInfo(this.acct)

	var PayTotal = len(CustomPaymentInfo)

	if PayTotal == 0 {
		if this.acct.adjustable_type_cd == "250" {
			org_term := utils.DataInterval(this.acct.last_reprice_date, this.acct.next_reprice_date)
			org_term_mult := "D"
			return linearInterpolation(this.acct.as_of_date, strconv.Itoa(org_term), org_term_mult, this.curve, this.ct)
		} else {
			org_term := utils.DataInterval(this.acct.origination_date, this.acct.maturity_date)
			org_term_mult := "D"
			return linearInterpolation(this.acct.as_of_date, strconv.Itoa(org_term), org_term_mult, this.curve, this.ct)
		}

	}

	//现金流日期
	var CUR_DATE = CustomPaymentInfo[0].pay_date
	//下次还款日
	var NEXT_DATE = CUR_DATE

	//在重定价周期内，使用支付频率，遍历重定价周期内的还款日期
	for i := 1; utils.CompareDate(NEXT_DATE, next_reprice_date) > 0 && PayTotal >= i; i++ {
		//本次重定价周期内，还款次数
		INDEX = i
		//当下次支付日小于下次重定价日时
		//将下次重定价日替换这次的还款日
		//后边根据lastT计算本次还款日的现金流
		CUR_DATE = NEXT_DATE
		//获取本期现金流利息计提天数与下次支付日期
		//ACC_PMT_TOTAL, NEXT_DATE = this.getACCDays(CUR_DATE, 0)

		ACC_PMT_TOTAL, _ = strconv.Atoi(CustomPaymentInfo[i-1].pay_days)

		if PayTotal == i {
		} else {
			NEXT_DATE = CustomPaymentInfo[i].pay_date
		}

		ORG_PAYMENT_AMT, _ = strconv.ParseFloat(CustomPaymentInfo[i-1].pay_amount, 8)

		//获取现金流日，本年计息总天数
		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, CUR_DATE)
		//获取现金流日，本年总天数。主要用于折现计算PV值
		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(CUR_DATE)
		//获取上次重定价日，到本次现金流日的折现天数，主要用于PV计算
		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, CUR_DATE)
		fc := (ORG_PAYMENT_AMT + LRD_BALANCE*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL)))
		pv := fc / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

		LRD_BALANCE = LRD_BALANCE - ORG_PAYMENT_AMT
		if LRD_BALANCE < 0 {
			this.log.Warn(this.acct.account_number, " has no balance. break")
			break
		}

		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)
		this.writeMsg(this.acct.account_number, INDEX, fc, CUR_NET_RATE,
			ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, CUR_DATE, pv, r)

		//保存pv,r,t值到结构体中，用于后续的ftp加权价格计算
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
	}

	//判断下次支付日期，是否大于下次重定价日，如果大于或等于下次重定价日，则将最后一期纳入现金流计算，
	//计算完成之后，退出现金流计算过程，将结果传递到pv集合中，计算加权平均的ftp价格
	if utils.CompareDate(NEXT_DATE, next_reprice_date) <= 0 {

		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(next_reprice_date)

		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, next_reprice_date)

		ACC_PMT_TOTAL = utils.DataInterval(CUR_DATE, next_reprice_date)

		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, next_reprice_date)

		ORG_PAYMENT_AMT, _ = strconv.ParseFloat(CustomPaymentInfo[INDEX].pay_amount, 8)

		pv := float64(0)
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		if next_reprice_date == this.acct.maturity_date {
			fc := (ORG_PAYMENT_AMT + LRD_BALANCE*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL)))
			pv = fc / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, fc, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		} else {
			pv = LRD_BALANCE / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, LRD_BALANCE, CUR_NET_RATE,
				ACC_PMT_TOTAL, 0, 0, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		}
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
		INDEX += 1
	}
	if INDEX == 0 {
		return ""
	}
	pvt := float64(0)
	pvtr := float64(0)
	for _, val := range rst {
		pvt += val.pv * val.t
		pvtr += val.pv * val.r * val.t
	}
	return strconv.FormatFloat(pvtr/pvt, 'f', 6, 32)
}

func (this *cashFlow) cashFlow_800(t_first string, last_reprice_date string, next_reprice_date string) string {

	var one cash
	var rst []cash
	//还款日当年总天数
	var ACC_TOTAL = 0
	//本次还款，计息累计天数
	var ACC_PMT_TOTAL = 0

	//折现天数
	var DISCOUNTING_DAYS = 0
	//折现当年，本年总天数
	var DISCOUNTING_TOTAL_YEAR = 0
	//本次重定价周期内还款次数
	var INDEX = 0
	//现金流日期
	var CUR_DATE = t_first
	//下次还款日
	var NEXT_DATE = t_first
	//获取上次重定价日余额
	var LRD_BALANCE, _ = strconv.ParseFloat(this.acct.lrd_balance, 8)
	//获取本期本金信息
	var ORG_PAYMENT_AMT, _ = strconv.ParseFloat(this.acct.org_payment_amt, 8)
	//获取业务执行利率
	var CUR_NET_RATE, _ = strconv.ParseFloat(this.acct.cur_net_rate, 8)
	// rate transform
	CUR_NET_RATE = CUR_NET_RATE / 100

	//在重定价周期内，使用支付频率，遍历重定价周期内的还款日期
	for i := 1; utils.CompareDate(NEXT_DATE, next_reprice_date) > 0; i++ {
		if LRD_BALANCE-float64(i-1)*ORG_PAYMENT_AMT <= 0 {
			this.log.Warn(this.acct.account_number, " has no balance. break calculate.")
			break
		}
		//本次重定价周期内，还款次数
		INDEX = i
		//当下次支付日小于下次重定价日时
		//将下次重定价日替换这次的还款日
		//后边根据lastT计算本次还款日的现金流
		CUR_DATE = NEXT_DATE
		//获取本期现金流利息计提天数与下次支付日期
		ACC_PMT_TOTAL, NEXT_DATE = this.getACCDays(CUR_DATE, 0)
		//获取现金流日，本年计息总天数
		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, CUR_DATE)
		//获取现金流日，本年总天数。主要用于折现计算PV值
		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(CUR_DATE)
		//获取上次重定价日，到本次现金流日的折现天数，主要用于PV计算
		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, CUR_DATE)

		fc := (ORG_PAYMENT_AMT + (LRD_BALANCE-(float64(INDEX)-1)*ORG_PAYMENT_AMT)*(CUR_NET_RATE)*float64(ACC_PMT_TOTAL)/float64(ACC_TOTAL))
		pv := fc / (math.Pow(1+(CUR_NET_RATE), float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		this.writeMsg(this.acct.account_number, INDEX, fc, CUR_NET_RATE,
			ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, CUR_DATE, pv, r)

		//保存pv,r,t值到结构体中，用于后续的ftp加权价格计算
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
	}

	//判断下次支付日期，是否大于下次重定价日，如果大于或等于下次重定价日，则将最后一期纳入现金流计算，
	//计算完成之后，退出现金流计算过程，将结果传递到pv集合中，计算加权平均的ftp价格
	if utils.CompareDate(NEXT_DATE, next_reprice_date) <= 0 {

		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(next_reprice_date)

		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, next_reprice_date)

		ACC_PMT_TOTAL = utils.DataInterval(CUR_DATE, next_reprice_date)

		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, next_reprice_date)

		pv := float64(0)
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		if next_reprice_date == this.acct.maturity_date {
			b := LRD_BALANCE - float64(INDEX)*ORG_PAYMENT_AMT
			// pv
			// B is balance
			// i is interest of this interest period
			// TDURATION is totla days of this reprice duration
			// T is the total days of this year (next_reprice_date)
			// pv = (B+i)/(1+r/100)^(TDURATION/T)
			fc := b + (LRD_BALANCE-(float64(INDEX))*ORG_PAYMENT_AMT)*(CUR_NET_RATE)*float64(ACC_PMT_TOTAL)/float64(ACC_TOTAL)
			pv = fc / (math.Pow(1+(CUR_NET_RATE), float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, fc, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		} else {
			// pv
			// B is balance
			// p is payment_amt total of this duration
			// TDURATION is total days of this duration
			// T is total days of this year (next_reprice_date)
			// pv = (B-p)/(1+r/100)^(TDURATION/T)
			fc := (LRD_BALANCE - float64(INDEX)*ORG_PAYMENT_AMT)
			pv = fc / math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR))

			this.writeMsg(this.acct.account_number, INDEX+1, fc, CUR_NET_RATE,
				0, 0, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		}
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
		INDEX += 1
	}
	if INDEX == 0 {
		return ""
	}
	pvt := float64(0)
	pvtr := float64(0)
	for _, val := range rst {
		pvt += val.pv * val.t
		pvtr += val.pv * val.r * val.t
	}
	return strconv.FormatFloat(pvtr/pvt, 'f', 6, 32)
}

func (this *cashFlow) cashFlow_700(t_first string, last_reprice_date string, next_reprice_date string) string {

	var one cash
	var rst []cash
	//还款日当年总天数
	var ACC_TOTAL = 0
	//本次还款，计息累计天数
	var ACC_PMT_TOTAL = 0

	//折现天数
	var DISCOUNTING_DAYS = 0
	//折现当年，本年总天数
	var DISCOUNTING_TOTAL_YEAR = 0
	//本次重定价周期内还款次数
	var INDEX = 0
	//现金流日期
	var CUR_DATE = t_first
	//下次还款日
	var NEXT_DATE = t_first
	//获取上次重定价日余额
	var LRD_BALANCE, _ = strconv.ParseFloat(this.acct.lrd_balance, 8)
	//获取本期本金信息
	var ORG_PAYMENT_AMT, _ = strconv.ParseFloat(this.acct.org_payment_amt, 8)
	//获取业务执行利率
	var CUR_NET_RATE, _ = strconv.ParseFloat(this.acct.cur_net_rate, 8)
	CUR_NET_RATE = CUR_NET_RATE / 100
	//在重定价周期内，使用支付频率，遍历重定价周期内的还款日期
	for i := 1; utils.CompareDate(NEXT_DATE, next_reprice_date) > 0; i++ {
		//本次重定价周期内，还款次数
		INDEX = i
		//当下次支付日小于下次重定价日时
		//将下次重定价日替换这次的还款日
		//后边根据lastT计算本次还款日的现金流
		CUR_DATE = NEXT_DATE
		//获取本期现金流利息计提天数与下次支付日期
		ACC_PMT_TOTAL, NEXT_DATE = this.getACCDays(CUR_DATE, 0)
		//获取现金流日，本年计息总天数
		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, CUR_DATE)
		//获取现金流日，本年总天数。主要用于折现计算PV值
		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(CUR_DATE)
		//获取上次重定价日，到本次现金流日的折现天数，主要用于PV计算
		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, CUR_DATE)

		LRD_BALANCE = LRD_BALANCE - ORG_PAYMENT_AMT + LRD_BALANCE*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL))
		if LRD_BALANCE < 0 {
			this.log.Warn("等额本息还款，本金已经为0，停止现金流计算过程")
		}
		pv := ORG_PAYMENT_AMT / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
		//r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", ci, ct)
		//fc := (ORG_PAYMENT_AMT + (LRD_BALANCE-(float64(INDEX)-1)*ORG_PAYMENT_AMT)*(CUR_NET_RATE/100)*float64(ACC_PMT_TOTAL)/float64(ACC_TOTAL))
		//pv := fc / (math.Pow(1+(CUR_NET_RATE/100), float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		this.writeMsg(this.acct.account_number, INDEX, ORG_PAYMENT_AMT, CUR_NET_RATE,
			ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, CUR_DATE, pv, r)

		//保存pv,r,t值到结构体中，用于后续的ftp加权价格计算
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
	}

	//判断下次支付日期，是否大于下次重定价日，如果大于或等于下次重定价日，则将最后一期纳入现金流计算，
	//计算完成之后，退出现金流计算过程，将结果传递到pv集合中，计算加权平均的ftp价格
	if utils.CompareDate(NEXT_DATE, next_reprice_date) <= 0 {

		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(next_reprice_date)

		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, next_reprice_date)

		ACC_PMT_TOTAL = utils.DataInterval(CUR_DATE, next_reprice_date)

		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, next_reprice_date)

		pv := float64(0)
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		if next_reprice_date == this.acct.maturity_date {
			pv = ORG_PAYMENT_AMT / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
			this.writeMsg(this.acct.account_number, INDEX+1, ORG_PAYMENT_AMT, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		} else {
			pv = LRD_BALANCE / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, LRD_BALANCE, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		}
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
		INDEX += 1
	}
	if INDEX == 0 {
		return ""
	}
	pvt := float64(0)
	pvtr := float64(0)
	for _, val := range rst {
		pvt += val.pv * val.t
		pvtr += val.pv * val.r * val.t
	}
	return strconv.FormatFloat(pvtr/pvt, 'f', 6, 32)
}

func (this *cashFlow) cashFlow_500(t_first string, last_reprice_date string, next_reprice_date string) string {

	var one cash
	var rst []cash
	//还款日当年总天数
	var ACC_TOTAL = 0
	//本次还款，计息累计天数
	var ACC_PMT_TOTAL = 0

	//折现天数
	var DISCOUNTING_DAYS = 0
	//折现当年，本年总天数
	var DISCOUNTING_TOTAL_YEAR = 0
	//本次重定价周期内还款次数
	var INDEX = 0
	//现金流日期
	var CUR_DATE = t_first
	//下次还款日
	var NEXT_DATE = t_first
	//获取上次重定价日余额
	var LRD_BALANCE, _ = strconv.ParseFloat(this.acct.lrd_balance, 8)
	//获取本期本金信息
	//	var ORG_PAYMENT_AMT, _ = strconv.ParseFloat("0", 8)
	//获取业务执行利率
	var CUR_NET_RATE, _ = strconv.ParseFloat(this.acct.cur_net_rate, 8)
	CUR_NET_RATE = CUR_NET_RATE / 100
	//在重定价周期内，使用支付频率，遍历重定价周期内的还款日期
	for i := 1; utils.CompareDate(NEXT_DATE, next_reprice_date) > 0; i++ {
		//本次重定价周期内，还款次数
		INDEX = i
		//当下次支付日小于下次重定价日时
		//将下次重定价日替换这次的还款日
		//后边根据lastT计算本次还款日的现金流
		CUR_DATE = NEXT_DATE
		//获取本期现金流利息计提天数与下次支付日期
		ACC_PMT_TOTAL, NEXT_DATE = this.getACCDays(CUR_DATE, 0)
		//获取现金流日，本年计息总天数
		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, CUR_DATE)
		//获取现金流日，本年总天数。主要用于折现计算PV值
		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(CUR_DATE)
		//获取上次重定价日，到本次现金流日的折现天数，主要用于PV计算
		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, CUR_DATE)

		fc := LRD_BALANCE * CUR_NET_RATE * float64(ACC_PMT_TOTAL) / float64(ACC_TOTAL)
		pv := fc / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		this.writeMsg(this.acct.account_number, INDEX, fc, CUR_NET_RATE,
			ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, CUR_DATE, pv, r)

		//保存pv,r,t值到结构体中，用于后续的ftp加权价格计算
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
	}

	//判断下次支付日期，是否大于下次重定价日，如果大于或等于下次重定价日，则将最后一期纳入现金流计算，
	//计算完成之后，退出现金流计算过程，将结果传递到pv集合中，计算加权平均的ftp价格
	if utils.CompareDate(NEXT_DATE, next_reprice_date) <= 0 {

		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(next_reprice_date)

		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, next_reprice_date)

		ACC_PMT_TOTAL = utils.DataInterval(CUR_DATE, next_reprice_date)

		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, next_reprice_date)

		pv := float64(0)
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		//		if next_reprice_date == this.acct.maturity_date {
		//			pv = (LRD_BALANCE + ORG_PAYMENT_AMT + (LRD_BALANCE-(float64(INDEX))*ORG_PAYMENT_AMT)*(CUR_NET_RATE/100)*float64(ACC_PMT_TOTAL)/float64(ACC_TOTAL)) / (math.Pow(1+(CUR_NET_RATE/100), float64(TDURATION)/float64(T)))
		//		} else {
		//			pv = (LRD_BALANCE - float64(INDEX)*ORG_PAYMENT_AMT) / math.Pow(1+CUR_NET_RATE/100, float64(TDURATION)/float64(T))
		//		}

		if next_reprice_date == this.acct.maturity_date {
			// pv
			// B is balance
			// i is interest of this interest period
			// TDURATION is totla days of this reprice duration
			// T is the total days of this year (next_reprice_date)
			// pv = (B+i)/(1+r/100)^(TDURATION/T)
			fc := LRD_BALANCE + LRD_BALANCE*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/float64(ACC_TOTAL)
			pv = fc / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, fc, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		} else {
			// pv
			// B is balance
			// p is payment_amt total of this duration
			// TDURATION is total days of this duration
			// T is total days of this year (next_reprice_date)
			// pv = (B-p)/(1+r/100)^(TDURATION/T)
			fc := LRD_BALANCE
			pv = fc / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, fc, CUR_NET_RATE,
				0, 0, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		}
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
		INDEX += 1
	}
	if INDEX == 0 {
		return ""
	}
	pvt := float64(0)
	pvtr := float64(0)
	for _, val := range rst {
		pvt += val.pv * val.t
		pvtr += val.pv * val.r * val.t
	}
	return strconv.FormatFloat(pvtr/pvt, 'f', 6, 32)
}

func (this *cashFlow) cashFlow_400(t_first string, last_reprice_date string, next_reprice_date string) string {

	var one cash
	var rst []cash
	//还款日当年总天数
	var ACC_TOTAL = 0
	//本次还款，计息累计天数
	var ACC_PMT_TOTAL = 0

	//折现天数
	var DISCOUNTING_DAYS = 0
	//折现当年，本年总天数
	var DISCOUNTING_TOTAL_YEAR = 0
	//本次重定价周期内还款次数
	var INDEX = 0
	//现金流日期
	var CUR_DATE = t_first
	//下次还款日
	var NEXT_DATE = t_first
	//获取上次重定价日余额
	var LRD_BALANCE = float64(0)
	//获取本期本金信息
	var ORG_PAYMENT_AMT, _ = strconv.ParseFloat(this.acct.org_payment_amt, 8)
	//获取业务执行利率
	var CUR_NET_RATE, _ = strconv.ParseFloat(this.acct.cur_net_rate, 8)
	CUR_NET_RATE = CUR_NET_RATE / 100
	//在重定价周期内，使用支付频率，遍历重定价周期内的还款日期
	for i := 1; utils.CompareDate(NEXT_DATE, next_reprice_date) > 0; i++ {
		//本次重定价周期内，还款次数
		INDEX = i
		//当下次支付日小于下次重定价日时
		//将下次重定价日替换这次的还款日
		//后边根据lastT计算本次还款日的现金流
		CUR_DATE = NEXT_DATE
		//获取本期现金流利息计提天数与下次支付日期
		ACC_PMT_TOTAL, NEXT_DATE = this.getACCDays(CUR_DATE, 0)
		//获取现金流日，本年计息总天数
		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, CUR_DATE)
		//获取现金流日，本年总天数。主要用于折现计算PV值
		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(CUR_DATE)
		//获取上次重定价日，到本次现金流日的折现天数，主要用于PV计算
		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, CUR_DATE)
		LRD_BALANCE = LRD_BALANCE + ORG_PAYMENT_AMT*float64(INDEX)*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL))
		pv := ORG_PAYMENT_AMT / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
		//r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", ci, ct)
		//fc := (ORG_PAYMENT_AMT + (LRD_BALANCE-(float64(INDEX)-1)*ORG_PAYMENT_AMT)*(CUR_NET_RATE/100)*float64(ACC_PMT_TOTAL)/float64(ACC_TOTAL))
		//pv := fc / (math.Pow(1+(CUR_NET_RATE/100), float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		this.writeMsg(this.acct.account_number, INDEX, ORG_PAYMENT_AMT, CUR_NET_RATE,
			ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, CUR_DATE, pv, r)

		//保存pv,r,t值到结构体中，用于后续的ftp加权价格计算
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
	}

	//判断下次支付日期，是否大于下次重定价日，如果大于或等于下次重定价日，则将最后一期纳入现金流计算，
	//计算完成之后，退出现金流计算过程，将结果传递到pv集合中，计算加权平均的ftp价格
	if utils.CompareDate(NEXT_DATE, next_reprice_date) <= 0 {

		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(next_reprice_date)

		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, next_reprice_date)

		ACC_PMT_TOTAL = utils.DataInterval(CUR_DATE, next_reprice_date)

		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, next_reprice_date)

		pv := float64(0)
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		if next_reprice_date == this.acct.maturity_date {
			ac := ORG_PAYMENT_AMT*float64(INDEX+1)*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL)) + LRD_BALANCE
			pv = (ac) / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, ac, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		} else {
			ac := ORG_PAYMENT_AMT*float64(INDEX+1)*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL)) + LRD_BALANCE
			pv = ac / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, ac, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		}
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
		INDEX += 1
	}
	if INDEX == 0 {
		return ""
	}
	pvt := float64(0)
	pvtr := float64(0)
	for _, val := range rst {
		pvt += val.pv * val.t
		pvtr += val.pv * val.r * val.t
	}
	return strconv.FormatFloat(pvtr/pvt, 'f', 6, 32)
}

func (this *cashFlow) cashFlow_300(t_first string, last_reprice_date string, next_reprice_date string) string {

	var one cash
	var rst []cash
	//还款日当年总天数
	var ACC_TOTAL = 0
	//本次还款，计息累计天数
	var ACC_PMT_TOTAL = 0

	//折现天数
	var DISCOUNTING_DAYS = 0
	//折现当年，本年总天数
	var DISCOUNTING_TOTAL_YEAR = 0
	//本次重定价周期内还款次数
	var INDEX = 0
	//现金流日期
	var CUR_DATE = t_first
	//下次还款日
	var NEXT_DATE = t_first
	//获取上次重定价日余额
	var LRD_BALANCE, _ = strconv.ParseFloat(this.acct.org_par_bal, 8)

	var ACC = float64(0)
	//获取本期本金信息
	var ORG_PAYMENT_AMT, _ = strconv.ParseFloat(this.acct.org_payment_amt, 64)
	//获取业务执行利率
	var CUR_NET_RATE, _ = strconv.ParseFloat(this.acct.cur_net_rate, 64)
	CUR_NET_RATE = CUR_NET_RATE / 100
	//在重定价周期内，使用支付频率，遍历重定价周期内的还款日期
	for i := 1; utils.CompareDate(NEXT_DATE, next_reprice_date) > 0; i++ {
		//本次重定价周期内，还款次数
		INDEX = i
		//当下次支付日小于下次重定价日时
		//将下次重定价日替换这次的还款日
		//后边根据lastT计算本次还款日的现金流
		CUR_DATE = NEXT_DATE
		//获取本期现金流利息计提天数与下次支付日期
		ACC_PMT_TOTAL, NEXT_DATE = this.getACCDays(CUR_DATE, 0)
		//获取现金流日，本年计息总天数
		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, CUR_DATE)
		//获取现金流日，本年总天数。主要用于折现计算PV值
		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(CUR_DATE)
		//获取上次重定价日，到本次现金流日的折现天数，主要用于PV计算
		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, CUR_DATE)

		ACC = ACC + LRD_BALANCE*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL))

		LRD_BALANCE = LRD_BALANCE - ORG_PAYMENT_AMT

		pv := ORG_PAYMENT_AMT / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
		//r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", ci, ct)
		//fc := (ORG_PAYMENT_AMT + (LRD_BALANCE-(float64(INDEX)-1)*ORG_PAYMENT_AMT)*(CUR_NET_RATE/100)*float64(ACC_PMT_TOTAL)/float64(ACC_TOTAL))
		//pv := fc / (math.Pow(1+(CUR_NET_RATE/100), float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		this.writeMsg(this.acct.account_number, INDEX, ORG_PAYMENT_AMT, CUR_NET_RATE,
			ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, CUR_DATE, pv, r)

		//保存pv,r,t值到结构体中，用于后续的ftp加权价格计算
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 32)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
	}

	//判断下次支付日期，是否大于下次重定价日，如果大于或等于下次重定价日，则将最后一期纳入现金流计算，
	//计算完成之后，退出现金流计算过程，将结果传递到pv集合中，计算加权平均的ftp价格
	if utils.CompareDate(NEXT_DATE, next_reprice_date) <= 0 {

		DISCOUNTING_TOTAL_YEAR = this.getTotalDaysOfYear(next_reprice_date)

		DISCOUNTING_DAYS = utils.DataInterval(last_reprice_date, next_reprice_date)

		ACC_PMT_TOTAL = utils.DataInterval(CUR_DATE, next_reprice_date)

		ACC_TOTAL = this.getTotalDaysOfAccrual(this.acct.accrual_basis_cd, next_reprice_date)

		pv := float64(0)

		r := linearInterpolation(last_reprice_date, strconv.Itoa(DISCOUNTING_DAYS), "D", this.curve, this.ct)

		if next_reprice_date == this.acct.maturity_date {
			ac := LRD_BALANCE*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL)) + ACC
			pv = (ac + ORG_PAYMENT_AMT) / (math.Pow(1+CUR_NET_RATE, float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))

			this.writeMsg(this.acct.account_number, INDEX+1, ac+ORG_PAYMENT_AMT, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		} else {
			ac := LRD_BALANCE*CUR_NET_RATE*float64(ACC_PMT_TOTAL)/(float64(ACC_TOTAL)) + ACC
			pv = (ac + ORG_PAYMENT_AMT) / (math.Pow(1+(CUR_NET_RATE), float64(DISCOUNTING_DAYS)/float64(DISCOUNTING_TOTAL_YEAR)))
			this.writeMsg(this.acct.account_number, INDEX+1, ac+ORG_PAYMENT_AMT, CUR_NET_RATE,
				ACC_PMT_TOTAL, ACC_TOTAL, DISCOUNTING_DAYS, DISCOUNTING_TOTAL_YEAR, next_reprice_date, pv, r)

		}
		one.pv = pv
		one.r, _ = strconv.ParseFloat(r, 64)
		one.t = float64(DISCOUNTING_DAYS)
		rst = append(rst, one)
		INDEX += 1
	}
	if INDEX == 0 {
		return ""
	}
	pvt := float64(0)
	pvtr := float64(0)
	for _, val := range rst {
		pvt += val.pv * val.t
		pvtr += val.pv * val.r * val.t
	}
	return strconv.FormatFloat(pvtr/pvt, 'f', 6, 32)
}

func (this *cashFlow) getTotalDaysOfYear(date string) int {
	year, _ := strconv.Atoi(date[0:4])
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return 366
	} else {
		return 365
	}
}

func (this *cashFlow) getTotalDaysOfAccrual(id string, date string) int {
	t := 0
	switch id {
	case "1", "2":
		t = 360
	case "4", "6":
		t = 365
	case "3", "5":
		t = this.getTotalDaysOfYear(date)
	}
	return t
}

func (this *cashFlow) getNextPaymentDate(date string) string {

	next_date := ""
	//将支付频率换成数字行
	num, _ := strconv.Atoi(this.acct.pmt_freq)

	//根据支付频率单位，判断支付周期单位
	switch this.acct.pmt_freq_mult {
	case "D":
		next_date, _ = utils.AddDays(date, num)
	case "M":
		next_date, _ = utils.AddMonths(date, num)
	case "Y":
		next_date, _ = utils.AddMonths(date, 12*num)
	default:
	}
	return next_date
}

func (this *cashFlow) writeMsg(v1 string,
	v2 int, v3 float64, v4 float64,
	v5 int, v6 int, v7 int, v8 int,
	v9 string, v10 float64, v11 string) {
	//go func() {
	this.log.Debug("{\"Account_number\":\"", v1,
		"\",\"Index_num\":\"", v2,
		"\",\"Current_FC\":\"", v3,
		"\",\"Currnet_rate\":\"", v4*100,
		"\",\"Acc_days\":\"", v5,
		"\",\"Acc_total_year\":\"", v6,
		"\",\"Discount_days\":\"", v7,
		"\",\"Discount_total_year\":\"", v8,
		"\",\"Cash_date\":\"", v9[0:10],
		"\",\"Current_PV\":\"", v10,
		"\",\"Current_R\":\"", v11,
		"\"}")
	//}()
}
