package ftpCalc

import (
	"fmt"
	"strconv"
	"utils"
)

type adjustHandle func(adj *ftpAdjustment, this *CalcFTP, val acctInfo, m busizMethod, ftpRate string) string

func adjust601(adj *ftpAdjustment, this *CalcFTP, val acctInfo, m busizMethod, ftpRate string) string {

	// only float rate business can be adjusted.
	if val.adjustable_type_cd == "250" {
		curveId, err := adj.getTlpCurve(val)
		if err != nil {
			return ""
		}
		if cm, ok := this.cv[curveId]; ok {
			date := this.getDate(val)
			ci, _ := this.matchCurve(date, cm)
			term := strconv.Itoa(utils.DataInterval(val.origination_date, val.maturity_date))
			return calcTermLiquidity(ci[0].as_of_date, term, "D", ci, this.ct)
		} else {
			fmt.Println("can't not found tlp curve info")
		}
	}
	return ""
}

func adjust603(adj *ftpAdjustment, this *CalcFTP, val acctInfo, m busizMethod, ftpRate string) string {
	curveId, err := adj.getFtpRestoreCurveId(val.busiz_id)
	if err != nil {
		return ""
	}
	if cm, ok := this.cv[curveId]; ok {
		date := val.origination_date
		ci, _ := this.matchCurve(date, cm)
		term := strconv.Itoa(utils.DataInterval(val.origination_date, val.maturity_date))
		return linearInterpolation(ci[0].as_of_date, term, "D", ci, this.ct)
	} else {
		fmt.Println("can't not found tlp curve info")
	}
	return ""
}

func adjust604(adj *ftpAdjustment, this *CalcFTP, val acctInfo, m busizMethod, ftpRate string) string {
	if rc, ok := adj.ri[val.busiz_id]; ok {
		return calcReserveRate(ftpRate, rc)
	} else {
		return ""
	}
}

func init() {
	register("601", adjust601)
	register("603", adjust603)
	register("604", adjust604)
}
