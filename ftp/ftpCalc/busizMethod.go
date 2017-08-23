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

type busizMethod struct {
	Busiz_id      string
	Ftp_method_id string
	Curve_id      string
	Term_cd       string
	Term_cd_mult  string
	Point_val     string
	Domain_id     string
}

type busizMethodTmp struct {
	Busiz_id      string
	Ftp_method_id []byte
	Curve_id      []byte
	Term_cd       []byte
	Term_cd_mult  []byte
	Point_val     []byte
	Domain_id     string
}

func (this *busizMethod) get(domain_id string) (map[string]busizMethod, error) {

	var one busizMethodTmp
	rst := make(map[string]busizMethod)
	rows, err := dbobj.Default.Query(FTPCALC_BUSIZ_RULES, domain_id)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	for rows.Next() {
		err := rows.Scan(&one.Busiz_id,
			&one.Ftp_method_id,
			&one.Curve_id,
			&one.Term_cd,
			&one.Term_cd_mult,
			&one.Point_val,
			&one.Domain_id)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		var tmp busizMethod
		tmp.Busiz_id = one.Busiz_id
		tmp.Ftp_method_id = string(one.Ftp_method_id)
		tmp.Curve_id = string(one.Curve_id)
		tmp.Term_cd = string(one.Term_cd)
		tmp.Term_cd_mult = string(one.Term_cd_mult)
		tmp.Point_val = string(one.Point_val)
		tmp.Domain_id = one.Domain_id
		rst[tmp.Busiz_id] = tmp
	}
	return rst, nil
}

func GetBusizMethod(domain_id string, p *CalcFTP) map[string]busizMethod {
	defer func() {
		p.wg.Done()
	}()
	r := new(busizMethod)
	bm, err := r.get(domain_id)
	if err != nil {
		p.log.Error(p.dispatchId, "->batch init business rules failed. error message is :", err)
		p.lock.RLock()
		p.status = 2
		p.lock.RUnlock()
		return nil
	}
	p.log.Info(p.dispatchId, "-> init business rules info complied. return status is 0")
	return bm
}
