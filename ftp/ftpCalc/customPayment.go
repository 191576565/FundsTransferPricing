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
	"fmt"
	"logs"
	"strconv"
	"utils"
)

type customPay struct {
	pay_amount string
	pay_date   string
	pay_sort   string
	pay_days   string
}

func (this *customPay) initPayment(as_of_date string) map[string][]customPay {
	ret := make(map[string][]customPay)
	rows, err := dbobj.Default.Query(FTPCALC_CUST_CUSTOMPAYMENT2, as_of_date)
	defer rows.Close()
	if err != nil {
		fmt.Println("init payment failed...")
		return ret
	}
	var one customPay
	var rst []customPay

	//	var lastPayDate = ""
	var lastAcct = ""
	var curAcct = "#####"
	var index = 1

	for rows.Next() {
		err := rows.Scan(&curAcct,
			&one.pay_amount,
			&one.pay_date)
		if err != nil {
			logs.Error(err.Error())
			return ret
		}
		one.pay_date = one.pay_date[0:10]
		one.pay_sort = strconv.Itoa(index)
		index++
		if lastAcct == curAcct {
			rst = append(rst, one)
		} else {
			one.pay_sort = "1"
			index = 2
			ret[lastAcct] = rst
			var tmp []customPay
			rst = tmp
			lastAcct = curAcct
			rst = append(rst, one)
		}
	}
	if len(rst) != 0 {
		ret[lastAcct] = rst
	}
	return ret
}

func (this customPay) get(val acctInfo) []customPay {

	rows, err := dbobj.Default.Query(FTPCALC_CUST_CUSTOMPAYMENT, val.account_number, val.as_of_date)
	defer rows.Close()
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	var one customPay
	var rst []customPay
	var lastPayDate = ""
	var index = 1
	for rows.Next() {
		err := rows.Scan(&one.pay_amount,
			&one.pay_date)
		if err != nil {
			logs.Error(err.Error())
			return nil
		}
		one.pay_sort = strconv.Itoa(index)
		index++
		if lastPayDate == "" {
			one.pay_days = strconv.Itoa(utils.DataInterval(val.origination_date, one.pay_date))
		} else {
			one.pay_days = strconv.Itoa(utils.DataInterval(lastPayDate, one.pay_date))
		}
		lastPayDate = one.pay_date
		rst = append(rst, one)
	}
	return rst
}

func getCustomPaymentInfo(val acctInfo) []customPay {
	r := new(customPay)
	return r.get(val)
}
