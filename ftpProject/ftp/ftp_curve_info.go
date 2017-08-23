package ftp

import (
	"bytes"
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"
	
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type FtpCurveInfo struct {
	As_of_date  string
	Curve_id    string
	Curve_desc  string
	Struct_code string
	Yieldt      []byte
	Yield       string
	Domain_id   string
	cnt         int
}
type FtpCurveInfoCtl struct {
	RouteControl
}
type FtpCurveInfot struct {
	As_of_date  string
	Curve_id    string
	Curve_desc  string
	Struct_code string
	Yieldt      []byte
	Domain_id   string
	cnt         int
}
type curinfo struct {
	Curve_id    string
	Curve_desc  string
	Domain_id   string
	As_of_date  string
	Curve_yield []curveYield
}

type curveYield struct {
	Struct_code string
	Yield       string
}

type FtpCurveInfoPage struct {
	RouteControl
}

func (this *FtpCurveInfoPage) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	did := doName //r.FormValue("DomainId")
	cid := r.FormValue("CurveId")
	sql := FTP_CURVEINFOPAGE_GET
	rows, err := dbobj.Default.Query(sql, did, cid)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询期限结构失败"))
		return
	}

	var one curveYield
	var rst []curveYield
	for rows.Next() {
		err := rows.Scan(&one.Struct_code)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询期限结构失败，请检查期限点是否存在异常"))
			return
		}
		rst = append(rst, one)
	}
	logs.Debug(rst)
	output := make(map[string]interface{}, 3)
	output["CurveId"] = cid
	output["DomainId"] = did
	output["Cstruct"] = rst
	ojs, err := json.Marshal(output)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *FtpCurveInfoCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()

	startDate := r.FormValue("StartDate")
	endDate := r.FormValue("EndDate")
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	if endDate == "" {
		endDate = time.Now().Format("2006-01-02")
	}
	if startDate == "" {
		y, _, _ := time.Now().Date()
		startDate = strconv.Itoa(y-1) + "-" + "01" + "-" + "01"
	}

	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))

	curveId := r.FormValue("CurveId")
	domainId := doName //r.FormValue("DomainId")

	if curveId == "" {
		curveId = "10001"
	}
	if domainId == "" {
		domainId = "FTP"
	}
	//fmt.Println("doname", doName)
	//fmt.Println("domainId", startDate, endDate, curveId, domainId, offset, offset+limit)
	sql := FTP_CURVEINFO_GET

	row, err := dbobj.Default.Query(sql, startDate, endDate, curveId, domainId, offset, offset+limit)
	logs.Info(startDate, endDate, curveId, domainId, offset, offset+limit)
	defer row.Close()
	if err != nil {
		logs.Error(err)
		return
	}
	var one FtpCurveInfo

	logs.Debug("Query business info success.")
	doid := ""
	curid := ""
	asdate := ""
	flag := false
	var onecurinfo curinfo
	var rstcurinfo []curinfo
	var onecuryield curveYield
	var endflag = false
	//fmt.Println("row")
	for row.Next() {
		err := row.Scan(&one.Domain_id,
			&one.Curve_id,
			&one.Curve_desc,
			&one.As_of_date,
			&one.Struct_code,
			&one.Yieldt,
			&one.cnt)
		if err != nil {
			logs.Error(err)
			return
		}
		if flag == false {
			flag = true
			doid = one.Domain_id
			curid = one.Curve_id
			asdate = one.As_of_date
			onecurinfo.Domain_id = one.Domain_id
			onecurinfo.Curve_id = one.Curve_id
			onecurinfo.Curve_desc = one.Curve_desc
			onecurinfo.As_of_date = one.As_of_date
			onecuryield.Struct_code = one.Struct_code
			onecuryield.Yield = string(one.Yieldt) //one.Yield
			onecurinfo.Curve_yield = append(onecurinfo.Curve_yield, onecuryield)
		} else if flag == true {
			if doid == one.Domain_id && curid == one.Curve_id && asdate == one.As_of_date {
				endflag = false
				onecuryield.Struct_code = one.Struct_code
				onecuryield.Yield = string(one.Yieldt) //one.Yield
				onecurinfo.Curve_yield = append(onecurinfo.Curve_yield, onecuryield)
			} else {
				endflag = false
				rstcurinfo = append(rstcurinfo, onecurinfo)
				//刷新数组
				var tmp []curveYield
				onecurinfo.Curve_yield = tmp
				doid = one.Domain_id
				curid = one.Curve_id
				asdate = one.As_of_date
				onecurinfo.Domain_id = one.Domain_id
				onecurinfo.Curve_id = one.Curve_id
				onecurinfo.Curve_desc = one.Curve_desc
				onecurinfo.As_of_date = one.As_of_date
				onecuryield.Struct_code = one.Struct_code
				onecuryield.Struct_code = one.Struct_code
				onecuryield.Yield = string(one.Yieldt) //one.Yield
				onecurinfo.Curve_yield = append(onecurinfo.Curve_yield, onecuryield)
			}
		}
	}
	if endflag == false {
		rstcurinfo = append(rstcurinfo, onecurinfo)
	}
	hz, _ := template.ParseFiles("./views/mas/ftp/ftp_curve_tranf.tpl")
	b := bytes.NewBuffer(make([]byte, 0))
	hz.Execute(b, rstcurinfo)
	str := b.String()

	val := `{"total":"` + strconv.Itoa(one.cnt) + `","rows":` + str + `}`
	w.Write([]byte(val))
}
