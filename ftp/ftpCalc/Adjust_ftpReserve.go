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
	"logs"
	"strconv"
)

type reserveAdjustInfo struct {
	reserve_percent string
	reserve_rate    string
}

func (this *reserveAdjustInfo) get(domain_id string, start_date string, end_date string) map[string]reserveAdjustInfo {

	rows, err := dbobj.Default.Query(FTPCALC_ADJUST_RESERVE, domain_id, start_date, end_date)
	defer rows.Close()
	if err != nil {
		logs.Error(err.Error())
		return nil
	}
	var one reserveAdjustInfo
	rst := make(map[string]reserveAdjustInfo)
	for rows.Next() {
		var bid string
		err := rows.Scan(&bid, &one.reserve_percent, &one.reserve_rate)
		if err != nil {
			logs.Error(err.Error())
			return nil
		}
		rst[bid] = one
	}
	return rst
}

func getReserveInfo(domain_id string, as_of_date string) map[string]reserveAdjustInfo {
	r := new(reserveAdjustInfo)
	return r.get(domain_id, as_of_date, as_of_date)
}

func calcReserveRate(ftpRate string, radjinfo reserveAdjustInfo) string {
	if ftpRate == "" || radjinfo.reserve_percent == "" || radjinfo.reserve_rate == "" {
		return ""
	}
	fr, _ := strconv.ParseFloat(ftpRate, 8)
	rp, _ := strconv.ParseFloat(radjinfo.reserve_percent, 8)
	rr, _ := strconv.ParseFloat(radjinfo.reserve_rate, 8)
	return strconv.FormatFloat(((rr - fr) * (rp / 100)), 'f', 8, 64)
}
