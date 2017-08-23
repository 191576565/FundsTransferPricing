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
)

//
// term liquidity info
// get all config info
//
type termLiquidity struct {
	curve_id           string
	reprice_freq_range string
}

// Author: huangzhanwei
// Time: 2016-07-01
// this function get all config info from database.
//
func (this *termLiquidity) get(domain_id string) map[string][]termLiquidity {
	rst := make(map[string][]termLiquidity)
	rows, err := dbobj.Default.Query(FTPCALC_ADJUST_TERMLIQ, domain_id)
	defer rows.Close()
	if err != nil {
		logs.Error(err.Error())
		return rst
	}
	var one termLiquidity
	var lstId string
	var tmp []termLiquidity

	for rows.Next() {
		var bid string
		err := rows.Scan(&bid, &one.curve_id, &one.reprice_freq_range)
		if err != nil {
			logs.Error(err.Error())
			return rst
		}
		if lstId == "" || bid == lstId {
			tmp = append(tmp, one)
			lstId = bid
		} else {
			rst[lstId] = tmp
			var tp []termLiquidity
			tmp = tp
			tmp = append(tmp, one)
			lstId = bid
		}
	}
	if len(tmp) > 0 {
		rst[lstId] = tmp
	}
	return rst
}

//
// Author: huangzhanwei
// Time: 2016-07-13
// get all Term Liquidity info
// filter is domain_id
//
func getTermLiquidityInfo(domain_id string) map[string][]termLiquidity {
	r := new(termLiquidity)
	return r.get(domain_id)
}

// Author: huangzhanwei
// Time: 2016-07-15
// use linearInterpolation calculate ftp rate info
//
func calcTermLiquidity(date string, org_term string, org_term_mult string, c []curveInfo, t map[string]curveStruct) string {
	return linearInterpolation(date, org_term, org_term_mult, c, t)
}
