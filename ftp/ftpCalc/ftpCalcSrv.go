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

//import (
//	"fmt"
//	"logs"
//)

//func StartCalcFTP(as_of_date string, domain_id string, offset int, limit int) {
//	r := new(TrialFtpCalc)
//	r.initData(as_of_date, domain_id, offset, limit)
//	fmt.Println(r.msg)
//}

//func StartCalcFTPByStream(as_of_date string, domain_id string, offset int, limit int) {
//	r := new(CalcFTP)
//	r.calcStream(as_of_date, domain_id, offset, limit)
//	fmt.Println(r.msg)
//}

//func FtpTrial() ([]acctInfo, error) {
//	ONE := acctInfo{account_number: "123451234",
//		as_of_date:         "2016-07-15",
//		busiz_id:           "10000056",
//		origination_date:   "2016-06-01",
//		maturity_date:      "2017-06-01",
//		org_term:           "12",
//		org_term_mult:      "M",
//		iso_currency_cd:    "CNY",
//		adjustable_type_cd: "0",
//		org_par_bal:        "1000",
//		org_payment_amt:    "1000",
//		cur_payment:        "1000",
//		cur_net_rate:       "4.56",
//		accrual_basis_cd:   "2",
//		amart_type_cd:      "500",
//		pmt_freq:           "12",
//		pmt_freq_mult:      "M",
//		last_payment_date:  "2016-06-01",
//		next_payment_date:  "2017-06-01",
//		remain_no_pmts_c:   "1",
//		reprice_freq:       "12",
//		reprice_freq_mult:  "M",
//		lrd_balance:        "1000",
//		last_reprice_date:  "2016-06-01",
//		next_reprice_date:  "2017-06-01",
//		ftpRate:            "0",
//		domain_id:          "FTP"}
//	var domain_id string
//	domain_id = "FTP"
//	var rst []acctInfo
//	rst = append(rst, ONE)
//	r := new(TrialFtpCalc)
//	a, err := r.FtpTrial(rst, domain_id)
//	if err != nil {
//		logs.Error(err)
//		return nil, err
//	}
//	fmt.Println(a)
//	return a, err
//}
