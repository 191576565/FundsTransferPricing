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
	"logs"
	"mas/ftp/calcLog"
	"strconv"
	"sync"
	"utils"
)

type ACCTInfo struct {
	Account_number     string
	As_of_date         string
	Busiz_id           string
	Origination_date   string
	Maturity_date      string
	Org_term           string
	Org_term_mult      string
	Iso_currency_cd    string
	Adjustable_type_cd string
	Org_par_bal        string
	Org_payment_amt    string
	Cur_payment        string
	Cur_net_rate       string
	Accrual_basis_cd   string
	Amart_type_cd      string
	Pmt_freq           string
	Pmt_freq_mult      string
	Last_payment_date  string
	Next_payment_date  string
	Remain_no_pmts_c   string
	Reprice_freq       string
	Reprice_freq_mult  string
	Lrd_balance        string
	Last_reprice_date  string
	Next_reprice_date  string
	FtpRate            string
	Adjone             string
	Adjtwo             string
	Adjthree           string
	Domain_id          string
}

type RpcSrvOfFTP struct {
}

// Author: huangzhnawei
// Time:2016-07-01
// you can call this func by rpc.
// first arg : account info
// second is pointer. it pointer string slice.
// first string is as_of_date of the batch
// second string is domain_id of the batch
// three string is offset
// four string is limit
// five string is batch id
func (t *RpcSrvOfFTP) StartFTPCalc(args []string, reply *string) error {
	defer func() {
		if r := recover(); r != nil {
			*reply = "计算ftp价格失败，请检查账户数据，规则，曲线信息等等"
		}
	}()
	var err error
	as_of_date := args[0]
	domain_id := args[1]
	soffset := args[2]
	slimit := args[3]
	dispatchId := args[4]

	if !utils.ValidDate(as_of_date) {
		*reply = "as_of_date is not a valid date parameter. please check it."
		return errors.New(*reply)
	}

	if !utils.ValidNumber(soffset) {
		*reply = "offset is not a valid number parameter. please check it."
		return errors.New(*reply)
	}
	offset, _ := strconv.Atoi(soffset)

	if !utils.ValidNumber(slimit) {
		*reply = "limit is not a valid number parameter. please check it."
		return errors.New(*reply)
	}
	limit, _ := strconv.Atoi(slimit)

	r := new(CalcFTP)

	//
	// set dispatch id
	r.dispatchId = dispatchId

	// set data input source
	// set data output dest

	r.inputSrc, r.outputSrc, err = initDispatch(dispatchId, domain_id)
	if err != nil {
		logs.Warn("this batch", dispatchId, "input Source or output Source")
		r.inputSrc = "mas_ftp_acct_info"
		r.outputSrc = "FTP_CALC_RESULT"
		return err
	}

	// init dispatch status
	// set batch stauts 1
	//
	updateDispatchStatus(dispatchId, as_of_date, "1")

	// init lock
	//
	r.lock = new(sync.RWMutex)

	// set status 0
	// batch starting
	r.status = 0

	// init statics
	r.successCnt = 0
	r.errorCnt = 0
	r.totalCnt = 0
	r.noBusizCnt = 0

	// inputCache hold input account info
	//
	r.inputCache = make(chan acctInfo, 8192)

	// init chan to recive success data.
	r.rstCache = make(chan ftpResult, 10240)

	// init batch log
	//
	r.log = calcLog.NewBatchLog(dispatchId, as_of_date, 0)
	r.log.Info("start batch: ", dispatchId, as_of_date)

	// start update ftp rate
	go r.updateFtpRate()

	// start get account info
	r.calcStream(as_of_date, domain_id, offset, limit)

	*reply = r.msg

	return nil
}

func (t *RpcSrvOfFTP) FtpTrial(args []ACCTInfo, reply *[]ACCTInfo) error {

	domain_id := "FTP"
	var rst []acctInfo
	for _, val := range args {
		var one acctInfo
		one.account_number = val.Account_number
		one.accrual_basis_cd = val.Accrual_basis_cd
		one.adjustable_type_cd = val.Adjustable_type_cd
		one.amart_type_cd = val.Amart_type_cd
		one.as_of_date = val.As_of_date
		one.busiz_id = val.Busiz_id
		one.cur_net_rate = val.Cur_net_rate
		one.cur_payment = val.Cur_payment
		one.domain_id = val.Domain_id
		one.iso_currency_cd = val.Iso_currency_cd
		one.last_payment_date = val.Last_payment_date
		one.last_reprice_date = val.Last_reprice_date
		one.lrd_balance = val.Lrd_balance
		one.maturity_date = val.Maturity_date
		one.next_payment_date = val.Next_payment_date
		one.next_reprice_date = val.Next_reprice_date
		one.org_par_bal = val.Org_par_bal
		one.org_payment_amt = val.Org_payment_amt
		one.org_term = val.Org_term
		one.org_term_mult = val.Org_term_mult
		one.origination_date = val.Origination_date
		one.pmt_freq = val.Pmt_freq
		one.pmt_freq_mult = val.Pmt_freq_mult
		one.remain_no_pmts_c = val.Remain_no_pmts_c
		one.reprice_freq = val.Reprice_freq
		one.reprice_freq_mult = val.Reprice_freq_mult
		rst = append(rst, one)
		domain_id = val.Domain_id
	}

	r := new(TrialFtpCalc)
	a, err := r.FtpTrial(rst, domain_id)
	if err != nil {
		logs.Error(err)
		return err
	}

	var ret []ACCTInfo
	for _, val := range a {
		var one ACCTInfo
		one.Account_number = val.account_number
		one.Accrual_basis_cd = val.accrual_basis_cd
		one.Adjustable_type_cd = val.adjustable_type_cd
		one.Amart_type_cd = val.amart_type_cd
		one.As_of_date = val.as_of_date
		one.Busiz_id = val.busiz_id
		one.Cur_net_rate = val.cur_net_rate
		one.Cur_payment = val.cur_payment
		one.Domain_id = val.domain_id
		one.FtpRate = val.ftpRate
		one.Adjone = val.adjone
		one.Adjtwo = val.adjtwo
		one.Adjthree = val.adjthree
		one.Iso_currency_cd = val.iso_currency_cd
		one.Last_payment_date = val.last_payment_date
		one.Last_reprice_date = val.last_reprice_date
		one.Lrd_balance = val.lrd_balance
		one.Maturity_date = val.maturity_date
		one.Next_payment_date = val.next_payment_date
		one.Next_reprice_date = val.next_reprice_date
		one.Org_par_bal = val.org_par_bal
		one.Org_payment_amt = val.org_payment_amt
		one.Org_term = val.org_term
		one.Org_term_mult = val.org_term_mult
		one.Origination_date = val.origination_date
		one.Pmt_freq = val.pmt_freq
		one.Pmt_freq_mult = val.pmt_freq_mult
		one.Remain_no_pmts_c = val.remain_no_pmts_c
		one.Reprice_freq = val.reprice_freq
		one.Reprice_freq_mult = val.reprice_freq_mult
		ret = append(ret, one)
	}
	*reply = ret
	return err
}
