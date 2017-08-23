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
	"strconv"
	"utils"
)

type lineinterpolate struct {
}

func linearInterpolation(date string, org_term string, org_term_mult string, c []curveInfo, t map[string]curveStruct) string {
	term := 0
	asOfDate := date
	if len(c) > 0 {
		asOfDate = c[0].as_of_date
	}

	switch org_term_mult {
	case "D":
		term, _ = strconv.Atoi(org_term)
	case "M":
		num, _ := strconv.Atoi(org_term)
		mDate, _ := utils.AddMonths(asOfDate, num)
		term = utils.DataInterval(asOfDate, mDate)
	case "Y":
		num, _ := strconv.Atoi(org_term)
		mDate, _ := utils.AddMonths(asOfDate, 12*num)
		term = utils.DataInterval(asOfDate, mDate)
	}

	lastDate := 0
	lastYield := 0.0000000001
	nextDate := 99999999
	nextYield := 0.0000000001

	for _, val := range c {
		curDate := 0
		switch t[val.struct_code].term_cd_mult {
		case "D":
			curDate, _ = strconv.Atoi(t[val.struct_code].term_cd)
		case "M":
			num, _ := strconv.Atoi(t[val.struct_code].term_cd)
			mDate, _ := utils.AddMonths(val.as_of_date, num)
			curDate = utils.DataInterval(val.as_of_date, mDate)
		case "Y":
			num, _ := strconv.Atoi(t[val.struct_code].term_cd)
			mDate, _ := utils.AddMonths(val.as_of_date, 12*num)
			curDate = utils.DataInterval(val.as_of_date, mDate)
		default:
		}
		if curDate >= lastDate && curDate <= term {
			lastDate = curDate
			lastYield, _ = strconv.ParseFloat(val.yield, 64)
		}
		if curDate <= nextDate && curDate >= term {
			nextDate = curDate
			nextYield, _ = strconv.ParseFloat(val.yield, 64)
		}
	}
	if nextDate == lastDate {
		return strconv.FormatFloat(nextYield, 'f', 6, 64)
	}
	if lastDate == 0 {
		return strconv.FormatFloat(float64(term)*float64(nextYield)/float64(nextDate), 'f', 6, 64)
	}
	return strconv.FormatFloat(float64(nextYield)-float64(nextDate-term)*float64(nextYield-lastYield)/float64(nextDate-lastDate), 'f', 6, 64)
}
