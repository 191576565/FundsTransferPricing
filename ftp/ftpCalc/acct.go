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
	"logs"
	"strings"
	"time"
)

type acctInfo struct {
	account_number     string
	as_of_date         string
	busiz_id           string
	origination_date   string
	maturity_date      string
	org_term           string
	org_term_mult      string
	iso_currency_cd    string
	adjustable_type_cd string
	org_par_bal        string
	org_payment_amt    string
	cur_payment        string
	cur_net_rate       string
	accrual_basis_cd   string
	amart_type_cd      string
	pmt_freq           string
	pmt_freq_mult      string
	last_payment_date  string
	next_payment_date  string
	remain_no_pmts_c   string
	reprice_freq       string
	reprice_freq_mult  string
	lrd_balance        string
	last_reprice_date  string
	next_reprice_date  string
	ftpRate            string
	adjone             string
	adjtwo             string
	adjthree           string
	domain_id          string
}

type caltmp struct {
	account_number     string
	as_of_date         string
	busiz_id           []byte
	origination_date   []byte
	maturity_date      []byte
	org_term           []byte
	org_term_mult      []byte
	iso_currency_cd    []byte
	adjustable_type_cd []byte
	org_par_bal        []byte
	org_payment_amt    []byte
	cur_payment        []byte
	cur_net_rate       []byte
	accrual_basis_cd   []byte
	amart_type_cd      []byte
	pmt_freq           []byte
	pmt_freq_mult      []byte
	last_payment_date  []byte
	next_payment_date  []byte
	remain_no_pmts_c   []byte
	reprice_freq       []byte
	reprice_freq_mult  []byte
	lrd_balance        []byte
	last_reprice_date  []byte
	next_reprice_date  []byte
	domain_id          string
}

type calcRst struct {
	account_number string
	as_of_date     string
	ftp_rate       float32
	ajust_01       float32
	ajust_02       float32
}

// Author: huangzhanwei
// Time:2016-07-02
// get account info
func (this *acctInfo) get(as_of_date string, domain_id string, offset int, limit int) ([]acctInfo, error) {

	sql := FTPCALC_ACCT

	rows, err := dbobj.Default.Query(sql, as_of_date, domain_id, offset, offset+limit)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var one acctInfo
	var tmp caltmp
	var rst []acctInfo
	for rows.Next() {
		err := rows.Scan(&tmp.account_number,
			&tmp.as_of_date,
			&tmp.busiz_id,
			&tmp.origination_date,
			&tmp.maturity_date,
			&tmp.org_term,
			&tmp.org_term_mult,
			&tmp.iso_currency_cd,
			&tmp.adjustable_type_cd,
			&tmp.org_par_bal,
			&tmp.org_payment_amt,
			&tmp.cur_payment,
			&tmp.cur_net_rate,
			&tmp.accrual_basis_cd,
			&tmp.amart_type_cd,
			&tmp.pmt_freq,
			&tmp.pmt_freq_mult,
			&tmp.last_payment_date,
			&tmp.next_payment_date,
			&tmp.remain_no_pmts_c,
			&tmp.reprice_freq,
			&tmp.reprice_freq_mult,
			&tmp.lrd_balance,
			&tmp.last_reprice_date,
			&tmp.next_reprice_date,
			&tmp.domain_id)
		if err != nil {
			logs.Error(err)
		}
		one.account_number = tmp.account_number
		one.as_of_date = this.stringToDate(tmp.as_of_date)
		one.busiz_id = string(tmp.busiz_id)
		one.origination_date = this.byteToDate(tmp.origination_date)
		one.maturity_date = this.byteToDate(tmp.maturity_date)
		one.org_term = string(tmp.org_term)
		one.org_term_mult = string(tmp.org_term_mult)
		one.iso_currency_cd = string(tmp.iso_currency_cd)
		one.adjustable_type_cd = string(tmp.adjustable_type_cd)
		one.org_par_bal = string(tmp.org_par_bal)
		one.org_payment_amt = string(tmp.org_payment_amt)
		one.cur_payment = string(tmp.cur_payment)
		one.cur_net_rate = string(tmp.cur_net_rate)
		one.accrual_basis_cd = string(tmp.accrual_basis_cd)
		one.amart_type_cd = string(tmp.amart_type_cd)
		one.pmt_freq = string(tmp.pmt_freq)
		one.pmt_freq_mult = string(tmp.pmt_freq_mult)
		one.last_payment_date = this.byteToDate(tmp.last_payment_date)
		one.next_payment_date = this.byteToDate(tmp.next_payment_date)
		one.remain_no_pmts_c = string(tmp.remain_no_pmts_c)
		one.reprice_freq = string(tmp.reprice_freq)
		one.reprice_freq_mult = string(tmp.reprice_freq_mult)
		one.lrd_balance = string(tmp.lrd_balance)
		one.last_reprice_date = this.byteToDate(tmp.last_reprice_date)
		one.next_reprice_date = this.byteToDate(tmp.next_reprice_date)
		one.domain_id = string(tmp.domain_id)
		rst = append(rst, one)
	}
	return rst, nil
}

// Author: huangzhanwei
// Time: 2016-07-15
// get account info by stream
func (this *acctInfo) getAcctStream(as_of_date string, domain_id string, offset int, limit int, cf *CalcFTP) error {

	/*
	* init total acct
	 */
	ts := `select count(*)  from ( 
	        select account_number
		     ,row_number() over(order by account_number) rk 
		    from  HZWY23 where as_of_date = to_date(:1,'YYYY-MM-DD') and domain_id = :2 
		  ) where rk > :3 and rk <= :4`
	ts = strings.Replace(ts, "HZWY23", cf.inputSrc, -1)
	dbobj.Default.QueryRow(ts, as_of_date, domain_id, offset, offset+limit).Scan(&cf.totalCnt)
	dbobj.Default.Exec("update ftp_dispatch_pro set all_rows = :1 where dispatch_id = :2 and domain_id = :3 and dispatch_date = to_date(:4,'YYYY-MM-DD')", cf.totalCnt, cf.dispatchId, domain_id, as_of_date)

	sql := strings.Replace(FTPCALC_ACCT_STREAM, "HZWY23", cf.inputSrc, -1)

	rows, err := dbobj.Default.Query(sql, as_of_date, domain_id, offset, offset+limit)
	defer rows.Close()
	if err != nil {
		cf.log.Error(err)
		return err
	}
	var one acctInfo
	var tmp caltmp

	for rows.Next() {
		err := rows.Scan(&tmp.account_number,
			&tmp.as_of_date,
			&tmp.busiz_id,
			&tmp.origination_date,
			&tmp.maturity_date,
			&tmp.org_term,
			&tmp.org_term_mult,
			&tmp.iso_currency_cd,
			&tmp.adjustable_type_cd,
			&tmp.org_par_bal,
			&tmp.org_payment_amt,
			&tmp.cur_payment,
			&tmp.cur_net_rate,
			&tmp.accrual_basis_cd,
			&tmp.amart_type_cd,
			&tmp.pmt_freq,
			&tmp.pmt_freq_mult,
			&tmp.last_payment_date,
			&tmp.next_payment_date,
			&tmp.remain_no_pmts_c,
			&tmp.reprice_freq,
			&tmp.reprice_freq_mult,
			&tmp.lrd_balance,
			&tmp.last_reprice_date,
			&tmp.next_reprice_date,
			&tmp.domain_id)
		if err != nil {
			cf.log.Error(err)
		}

		one.account_number = tmp.account_number
		one.as_of_date = this.stringToDate(tmp.as_of_date)
		one.busiz_id = string(tmp.busiz_id)
		one.origination_date = this.byteToDate(tmp.origination_date)
		one.maturity_date = this.byteToDate(tmp.maturity_date)
		one.org_term = string(tmp.org_term)
		one.org_term_mult = string(tmp.org_term_mult)
		one.iso_currency_cd = string(tmp.iso_currency_cd)
		one.adjustable_type_cd = string(tmp.adjustable_type_cd)
		one.org_par_bal = string(tmp.org_par_bal)
		one.org_payment_amt = string(tmp.org_payment_amt)
		one.cur_payment = string(tmp.cur_payment)
		one.cur_net_rate = string(tmp.cur_net_rate)
		one.accrual_basis_cd = string(tmp.accrual_basis_cd)
		one.amart_type_cd = string(tmp.amart_type_cd)
		one.pmt_freq = string(tmp.pmt_freq)
		one.pmt_freq_mult = string(tmp.pmt_freq_mult)
		one.last_payment_date = this.byteToDate(tmp.last_payment_date)
		one.next_payment_date = this.byteToDate(tmp.next_payment_date)
		one.remain_no_pmts_c = string(tmp.remain_no_pmts_c)
		one.reprice_freq = string(tmp.reprice_freq)
		one.reprice_freq_mult = string(tmp.reprice_freq_mult)
		one.lrd_balance = string(tmp.lrd_balance)
		one.last_reprice_date = this.byteToDate(tmp.last_reprice_date)
		one.next_reprice_date = this.byteToDate(tmp.next_reprice_date)
		one.domain_id = string(tmp.domain_id)

		select {
		case <-time.After(time.Second * 300):
			cf.log.Error("发送账户信息已超时,退出定价")
			cf.lock.RLock()
			cf.status = 1
			cf.lock.RUnlock()
			close(cf.inputCache)
			return errors.New("连接超时")
		case cf.inputCache <- one:
			cf.lock.RLock()
			if cf.status != 0 {
				close(cf.inputCache)
				cf.log.Debug("stop batch", cf.dispatchId)
				cf.lock.RUnlock()
				return nil
			} else {
				cf.lock.RUnlock()
			}
		}
	}
	close(cf.inputCache)
	return nil
}

// Author: huangzhanwei
// Time:2016-07-01
// format byte value to date value
func (this *acctInfo) byteToDate(dt []byte) string {
	if len(string(dt)) > 10 {
		return string(dt)[0:10]
	} else {
		return string(dt)
	}
}

// Author: huangzhanwei
// Time: 2016-07-01
// format string value to date value
func (this *acctInfo) stringToDate(dt string) string {
	if len(dt) > 10 {
		return dt[0:10]
	} else {
		return dt
	}
}

// Author: huangzhanwei
// Time:2016-07-01
// format byte value to date value
func ByteToDate(dt []byte) string {
	if len(string(dt)) > 10 {
		return string(dt)[0:10]
	} else {
		return string(dt)
	}
}

// Author: huangzhanwei
// Time : 2016-06-30
// format string value to date value
//
func StringToDate(dt string) string {
	if len(dt) > 10 {
		return dt[0:10]
	} else {
		return dt
	}
}

// Author: huangzhanwei
// Time: 2016-07-03
// get all account info
//
func GetAcctInfo(as_of_date string, domain_id string, offset, limit int) ([]acctInfo, error) {
	r := new(acctInfo)
	return r.get(as_of_date, domain_id, offset, limit)
}

// Author:huangzhanwei
// Time : 2016-07-05
// get all account info by stream
//
func GetAcctInfoStream(as_of_date string, domain_id string, offset int, limit int, cf *CalcFTP) error {

	r := new(acctInfo)
	return r.getAcctStream(as_of_date, domain_id, offset, limit, cf)
}
