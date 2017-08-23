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
	"database/sql"
	"dbobj"
	"errors"
	"mas/ftp/calcLog"
	"strconv"
)

type curve struct {
	oldest_curve []curveInfo
	oldestdate   string
	curve_id     string
	value        map[string][]curveInfo
}

type curveInfo struct {
	as_of_date  string
	struct_code string
	yield       string
}

type curveInfoTmp struct {
	as_of_date  sql.NullString
	struct_code sql.NullString
	yield       sql.NullString
}

func (this *curveInfo) get(domain_id string, curve_id string, year string, log calcLog.Calclog) ([]curveInfo, error) {
	rows, err := dbobj.Default.Query(FTPCALC_CURVE_YIELD, domain_id, curve_id, year)
	defer rows.Close()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	var one curveInfoTmp
	var tmp curveInfo
	var rst []curveInfo
	for rows.Next() {
		err := rows.Scan(&one.as_of_date, &one.struct_code, &one.yield)
		if err != nil {
			log.Error("There is no curve found in this domain")
			return nil, err
		}
		tmp.as_of_date = StringToDate(one.as_of_date.String)
		tmp.struct_code = one.struct_code.String
		tmp.yield = one.yield.String
		rst = append(rst, tmp)
	}
	return rst, nil
}

func (this *curveInfo) getCurveList(domain string, log calcLog.Calclog) ([]string, error) {
	var rst []string
	var one string

	rows, err := dbobj.Default.Query(FTPCALC_CURVE_YIELD1, domain)
	defer rows.Close()
	if err != nil {
		log.Error("get curve id failed.", err)
		return rst, err
	}
	for rows.Next() {
		err := rows.Scan(&one)
		if err != nil {
			log.Error("get curve id failed.", err)
			return rst, err
		}
		rst = append(rst, one)
	}
	if len(rst) > 0 {
		return rst, nil
	} else {
		return rst, errors.New("没有已经定义的曲线。请配置曲线信息")
	}
}

func (this *curveInfo) getOldestCurveDate(domain string, curve_id string, oldestdate string, log calcLog.Calclog) ([]curveInfo, string, error) {
	var one curveInfoTmp
	var rst []curveInfo

	rows, err := dbobj.Default.Query(FTPCALC_CURVE_YIELD2, domain, curve_id, oldestdate)
	defer rows.Close()
	if err != nil {
		log.Error("It is not found oldest curve info values.", err)
		return rst, "", err
	}
	var tmp curveInfo
	for rows.Next() {
		err := rows.Scan(&one.struct_code,
			&one.yield,
			&one.as_of_date)
		if err != nil {
			log.Error("There is no oldest curve value.", err)
			return rst, "", err
		}
		tmp.as_of_date = StringToDate(one.as_of_date.String)
		tmp.struct_code = one.struct_code.String
		tmp.yield = one.yield.String
		rst = append(rst, tmp)
	}
	if tmp.as_of_date == "" {
		log.Error(curve_id, domain, oldestdate, "无法取到最早的一条曲线信息。请确认mas_curve_define，mas_curve_info_struct_node，mas_curve_info，系统默认情况下只计算2005年1月1号之后的曲线信息")
		return rst, "", errors.New("无法取到最早的一条曲线信息。请确认mas_curve_define，mas_curve_info_struct_node，mas_curve_info，系统默认情况下只计算2005年1月1号之后的曲线信息")
	} else {
		return rst, tmp.as_of_date, nil
	}
}

func (this *curveInfo) getCurve(domain_id string, curve_id string, as_of_date string, log calcLog.Calclog) (*curve, error) {

	rst := new(curve)
	rst.value = make(map[string][]curveInfo)

	ci, d, err := this.getOldestCurveDate(domain_id, curve_id, "2005-01-01", log)
	if err != nil {
		return nil, err
	} else if len(ci) == 0 {
		return rst, errors.New("can't find curve info values.")
	}

	//获取最早的一条曲线信息
	rst.oldest_curve = ci
	rst.oldestdate = StringToDate(d)
	//曲线编码信息
	rst.curve_id = curve_id

	//最早曲线信息年份
	sdate, _ := strconv.Atoi(d[0:4])

	//最新日期年份
	edate, _ := strconv.Atoi(as_of_date)

	for sdate <= edate {
		cv, err := this.get(domain_id, curve_id, strconv.Itoa(sdate), log)
		if err != nil {
			log.Info("this domain ,", curve_id, " has no curve info values.", domain_id)
			return nil, err
		}

		rst.value[strconv.Itoa(sdate)] = cv
		sdate = sdate + 1
	}
	return rst, nil
}

func GetCurveVal(domain string, as_of_date string, p *CalcFTP) map[string]*curve {

	defer func() {
		p.wg.Done()
	}()

	cd := new(curveInfo)
	rst := make(map[string]*curve)

	list, err := cd.getCurveList(domain, p.log)
	if err != nil {
		p.log.Error(p.dispatchId, "-> init curve info failed. error message is:", err)
		p.lock.RLock()
		p.status = 2
		p.lock.RUnlock()
		return nil
	}

	for _, val := range list {
		r, err := cd.getCurve(domain, val, as_of_date, p.log)
		if err != nil {
			continue
		}

		rst[val] = r

	}

	if len(rst) == 0 {
		p.log.Error(p.dispatchId, "-> can't found curve info in this domain.  please check the batch info.")
		p.lock.RLock()
		p.status = 2
		p.lock.RUnlock()
		return nil
	} else {
		p.log.Info(p.dispatchId, "-> init curve info complied.")
		return rst
	}
}
