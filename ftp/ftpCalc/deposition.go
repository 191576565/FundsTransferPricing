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
)

type precipitation struct {
	busiz_id     string
	term_cd      string
	term_cd_mult string
	weight       string
}

func (this *precipitation) getConfig(domain_id string) (map[string][]precipitation, error) {

	rows, err := dbobj.Default.Query(FTPCALC_DEPOSITION, domain_id)

	defer rows.Close()
	if err != nil {
		fmt.Println(FTPCALC_DEPOSITION, err)
		return nil, err
	}
	var one precipitation
	var rst []precipitation
	ret := make(map[string][]precipitation)
	lastBusizId := "#*####^^^###"
	oneflat := true
	for rows.Next() {
		err := rows.Scan(&one.busiz_id, &one.term_cd, &one.term_cd_mult, &one.weight)
		if err != nil {
			fmt.Println("查询数据库失败")
			return nil, err
		}
		if lastBusizId == "#*####^^^###" {
			rst = append(rst, one)
			lastBusizId = one.busiz_id
		} else if one.busiz_id == lastBusizId {
			rst = append(rst, one)
		} else {
			oneflat = false
			ret[lastBusizId] = rst
			var tmp []precipitation
			lastBusizId = one.busiz_id
			rst = append(tmp, one)
		}
	}

	if oneflat == true {
		ret[lastBusizId] = rst
	} else {
		ret[one.busiz_id] = rst
	}

	return ret, nil
}

func getPrecipitation(domain_id string, p *CalcFTP) map[string][]precipitation {
	defer func() {
		p.wg.Done()
	}()
	r := new(precipitation)
	d, err := r.getConfig(domain_id)
	if err != nil {
		p.log.Error(p.dispatchId, "-> init precipitation config values failed. error info is :", err)
		p.lock.RLock()
		p.status = 2
		p.lock.RUnlock()
		return nil
	}
	p.log.Info(p.dispatchId, "-> init precipitation info successfully. return status is 0")
	return d
}
