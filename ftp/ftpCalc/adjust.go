// Author : huangzhanwei
// Time: 2016-09-09
// All rights reserved.
// This package used to calculate ftp adjust rate.
//
package ftpCalc

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// define 20 adjust item.
// key is the adjust id number.
// value is the describe of the adjust id number.
var (
	ADJUST_TYPE = map[string]string{
		"601": "Term Liquility Adjust",
		"602": "Reserve",
		"603": "Reserve",
		"604": "Reserve Capital Adjust",
		"605": "Reserve",
		"606": "Reserve",
		"607": "Reserve",
		"608": "Reserve",
		"609": "Reserve",
		"610": "Reserve",
		"801": "Reserve",
		"802": "Reserve",
		"803": "Reserve",
		"804": "Reserve",
		"805": "Reserve",
		"806": "Reserve",
		"807": "Reserve",
		"808": "Reserve",
		"809": "Reserve",
		"810": "Reserve",
	}
)

var adjustFunc = make(map[string]adjustHandle)
var ftpAdjust = &ftpAdjustment{}

// register adjust handle func
//
func register(key string, handle adjustHandle) {
	if _, dup := adjustFunc[key]; dup {
		fmt.Println("reregister.")
		return
	}
	adjustFunc[key] = handle
}

type ftpAdjustment struct {
	// used for term liquidity
	//
	tlp map[string][]termLiquidity
	ri  map[string]reserveAdjustInfo
	fp  map[string]string
}

// this function return adjust desc.
//
func (this *ftpAdjustment) getAdjustDesc(key string) string {

	if typ, ok := ADJUST_TYPE[key]; ok {
		return typ
	} else {
		return ""
	}
}

// this function return wheather adjust id is valid.
//
func (this *ftpAdjustment) isValid(key string) bool {

	if _, ok := ADJUST_TYPE[key]; ok {
		return true
	} else {
		return false
	}
}

func (this *ftpAdjustment) getFtpRestoreCurveId(bid string) (string, error) {
	if bid, ok := this.fp[bid]; ok {
		return bid, nil
	} else {
		return "", errors.New("no curve found in map")
	}
}

func (this *ftpAdjustment) getTlpCurve(acct acctInfo) (string, error) {

	if tl, ok := this.tlp[acct.busiz_id]; ok {

		for _, val := range tl {
			// parse filter condidion
			cond := strings.TrimSpace(val.reprice_freq_range)

			condlen := len(cond)
			if condlen < 3 {
				return "", nil
			}
			firstChar := cond[0]
			lastChar := cond[condlen-1]

			cond = cond[1:(condlen - 1)]
			cond = strings.ToUpper(cond)
			cd := strings.Split(cond, ",")
			if len(cd) != 2 {
				fmt.Println("filter condidion invalid.")
				return "", errors.New("filter condidion invalid.")
			}
			minValStr := cd[0]
			maxValStr := cd[1]
			minVal := minValStr[0:(len(minValStr) - 1)]
			maxVal := maxValStr[0:(len(maxValStr) - 1)]
			minValMult := minValStr[len(minValStr)-1]
			maxValMult := maxValStr[len(maxValStr)-1]
			if len(minValStr) == 1 {
				minVal = minValStr
				minValMult = 'D'
			}
			if len(maxValStr) == 1 {
				maxVal = maxValStr
				maxValMult = 'D'
			}

			min := this.changeOrgTermToMonth(minVal, minValMult)

			max := this.changeOrgTermToMonth(maxVal, maxValMult)

			freq := this.changeOrgTermToMonth(acct.reprice_freq, acct.reprice_freq_mult[0])

			if firstChar == '[' && lastChar == ']' && freq >= min && freq <= max {
				return val.curve_id, nil
			} else if firstChar == '[' && lastChar == ')' && freq >= min && freq < max {
				return val.curve_id, nil
			} else if firstChar == '(' && lastChar == ')' && freq > min && freq < max {
				return val.curve_id, nil
			} else if firstChar == '(' && lastChar == ']' && freq >= min && freq <= max {
				return val.curve_id, nil
			}
		}
	}
	return "", errors.New("no curve found in rules talbes.")
}

func (this *ftpAdjustment) changeOrgTermToMonth(orgTerm string, orgMult byte) float64 {
	switch orgMult {
	case 'D':
		ot, err := strconv.Atoi(orgTerm)
		if err != nil {
			return 0
		}
		return float64(ot) / 30
	case 'M':
		ot, err := strconv.Atoi(orgTerm)
		if err != nil {
			return 0
		}
		return float64(ot)
	case 'Y':
		ot, err := strconv.Atoi(orgTerm)
		if err != nil {
			return 0
		}
		return float64(ot * 12)
	default:
		fmt.Println("no match")
		return 0
	}
}

// this function return handle function of adjust id
//
func (this *ftpAdjustment) getHandle(key string) (adjustHandle, bool) {

	if val, ok := adjustFunc[key]; ok {
		return val, true
	} else {
		return nil, false
	}
}

func ftpAdjustInit(domain_id string, p *CalcFTP, as_of_date string) *ftpAdjustment {
	defer func() {
		p.wg.Done()
	}()
	r := new(ftpAdjustment)
	r.tlp = getTermLiquidityInfo(domain_id)
	r.ri = getReserveInfo(domain_id, as_of_date)
	r.fp = getFtpRestoreCurve(domain_id)
	p.log.Info(p.dispatchId, "-> init ftpAdjustment info complied. return status is 0")
	return r
}

// this function calc adjust rate
// if adjust id register in adjustFunc. it will be run
//
func (this *ftpAdjustment) calcAdjustRate(tf *CalcFTP, val acctInfo, m busizMethod, ftpRate string) map[string]string {

	rst := make(map[string]string)

	for key, f := range adjustFunc {

		rst[key] = f(this, tf, val, m, ftpRate)

	}
	return rst
}
