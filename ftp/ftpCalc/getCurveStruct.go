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

type curveStruct struct {
	term_cd      string
	term_cd_mult string
}

func (this *curveStruct) get(domain_id string) (map[string]curveStruct, error) {
	rst := make(map[string]curveStruct, 56)
	var one curveStruct

	rows, err := dbobj.Default.Query(FTPCALC_CURVE_STRUCT, domain_id)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return rst, err
	}
	for rows.Next() {
		var sc string
		err := rows.Scan(&sc, &one.term_cd, &one.term_cd_mult)
		if err != nil {
			logs.Error(err)
			return rst, err
		}
		rst[sc] = one
	}
	return rst, nil
}

func GetCurveStruct(domain_id string, p *CalcFTP) map[string]curveStruct {
	defer func() {
		p.wg.Done()
	}()
	r := new(curveStruct)
	d, e := r.get(domain_id)
	if e != nil {
		p.log.Error(p.dispatchId, "-> init curve struct info failed.error info is :", e)
		p.lock.RLock()
		p.status = 2
		p.lock.RUnlock()
		return nil
	}
	p.log.Info(p.dispatchId, "-> init curve struct info successfully.return status is 0")
	return d
}
