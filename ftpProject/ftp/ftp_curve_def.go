package ftp

import (
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"
	"ftpProject/utils/cacheutil"

	"ftpProject/utils"
	"net/http"
	"strings"
)

type FtpCurveDef struct {
	Curve_id          string
	Curve_desc        string
	Curve_type        string
	Curve_type_desc   string
	Iso_currency_cd   string
	Iso_currency_desc string
	Create_date       string
	Latest_date       string
	Domain_id         string
	Domain_desc       string
	Rep_id            string
	Rep_desc          string
	All_Struct_Code   string
	cnt               int
}
type FtpCurveDefCtl struct {
	RouteControl
}
type FtpCurveDefPage struct {
	RouteControl
}

type FtpCurveQuery struct {
	Curve_id   string
	Curve_desc string
	RouteControl
}

type ftpLatestCurveInfo struct {
	domain_id  string
	curve_id   string
	as_of_date string
}

func (this *FtpCurveDefPage) Get() {
	this.TplName = "mas/ftp/ftp_curve_def.tpl"
}
func (this *FtpCurveDefCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	ncurve, err := getLatestCurve()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("暂时没有配置任务曲线信息"))
		return
	}
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//

	sql := FTP_CURVEDEF_GET1
	row, err := dbobj.Default.Query(sql, doName)
	//

	//fmt.Println(searchword)
	if r.FormValue("keyword") != "" {
		sql = FTP_CURVEDEF_GET2
		ss := utils.HandleSqlKey(r.FormValue("keyword"))
		searchword := "%"
		searchword = searchword + ss
		searchword = searchword + "%"
		row, err = dbobj.Default.Query(sql, searchword, searchword, doName)
	}
	//

	defer row.Close()
	if err != nil {
		logs.Error(err)
		return
	}
	var (
		one FtpCurveDef
		rst []FtpCurveDef
	)
	for row.Next() {
		var (
			tmprepid   []byte
			tmprepdesc []byte
		)
		err := row.Scan(&one.Curve_id,
			&one.Curve_desc,
			&one.Curve_type,
			&one.Curve_type_desc,
			&one.Iso_currency_cd,
			&one.Iso_currency_desc,
			&one.Create_date,
			&one.Domain_id,
			&tmprepid,
			&tmprepdesc,
			&one.cnt)
		if err != nil {
			logs.Error(err)
			return
		}
		one.Rep_id = string(tmprepid)
		one.Rep_desc = string(tmprepdesc)
		one.Domain_desc = this.DomainName
		sql := FTP_CURVEDEF_GET3
		roww, err := dbobj.Default.Query(sql, one.Curve_id, doName)
		if err != nil {
			logs.Error(err)
			return
		}
		var struct_sum []string
		var struct_cc string
		var struct_cct []byte
		for roww.Next() {
			err := roww.Scan(&struct_cct)
			if err != nil {
				logs.Error(err)
				return
			}
			struct_cc = string(struct_cct)
			struct_sum = append(struct_sum, struct_cc)
		}

		roww.Close() //2016.9.13 add
		one.All_Struct_Code = strings.Join(struct_sum, ",")
		one.Latest_date = getLatestDate(ncurve, one.Domain_id, one.Curve_id)
		rst = append(rst, one)
	}
	ojs, err := json.Marshal(rst)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *FtpCurveDefCtl) Get_old() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	ncurve, err := getLatestCurve()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("暂时没有配置任务曲线信息"))
		return
	}
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//

	sql := FTP_CURVEDEF_GET1
	row, err := dbobj.Default.Query(sql, doName)
	//

	//fmt.Println(searchword)
	if r.FormValue("keyword") != "" {
		sql = FTP_CURVEDEF_GET2
		ss := utils.HandleSqlKey(r.FormValue("keyword"))
		searchword := "%"
		searchword = searchword + ss
		searchword = searchword + "%"
		row, err = dbobj.Default.Query(sql, searchword, searchword, doName)
	}
	//

	defer row.Close()
	if err != nil {
		logs.Error(err)
		return
	}
	var (
		one FtpCurveDef
		rst []FtpCurveDef
	)
	for row.Next() {
		var (
			tmprepid   []byte
			tmprepdesc []byte
		)
		err := row.Scan(&one.Curve_id,
			&one.Curve_desc,
			&one.Curve_type,
			&one.Curve_type_desc,
			&one.Iso_currency_cd,
			&one.Create_date,
			&one.Domain_id,
			&tmprepid,
			&tmprepdesc,
			&one.cnt)
		if err != nil {
			logs.Error(err)
			return
		}
		one.Rep_id = string(tmprepid)
		one.Rep_desc = string(tmprepdesc)
		one.Domain_desc = this.DomainName
		sql := FTP_CURVEDEF_GET3
		roww, err := dbobj.Default.Query(sql, one.Curve_id, doName)
		if err != nil {
			logs.Error(err)
			return
		}
		var struct_sum []string
		var struct_cc string
		var struct_cct []byte
		for roww.Next() {
			err := roww.Scan(&struct_cct)
			if err != nil {
				logs.Error(err)
				return
			}
			struct_cc = string(struct_cct)
			struct_sum = append(struct_sum, struct_cc)
		}

		roww.Close() //2016.9.13 add
		one.All_Struct_Code = strings.Join(struct_sum, ",")
		one.Latest_date = getLatestDate(ncurve, one.Domain_id, one.Curve_id)
		rst = append(rst, one)
	}
	//匹配币种
	var curveCucyData []CurrencyInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		curveCucyData = val.CucyData
	}
	for i, curve := range rst {
		for _, curr := range curveCucyData {
			if curve.Iso_currency_cd == curr.IsoCurrencyCd {
				rst[i].Iso_currency_desc = curr.IsoCurrencyDesc
			}
		}
	}
	//
	ojs, err := json.Marshal(rst)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func getLatestCurve() ([]ftpLatestCurveInfo, error) {
	dbtp := dbobj.DefaultDB()
	sql := ""
	switch dbtp {
	case "oracle":
		sql = FTP_GETLATEST_CURVE
	case "db2":
		sql = `select
				   d.domain_id
				   ,d.curve_id
				   ,TO_CHAR(max(i.as_of_date),'YYYY-MM-DD')
				from mas_curve_define d
				inner join mas_curve_info_struct_node n
				on d.curve_id = n.curve_id
				and d.domain_id = n.domain_id
				inner join mas_curve_info i
				on n.uuid = i.curve_uuid
				group by d.domain_id,d.curve_id
				`
	}

	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	var rst []ftpLatestCurveInfo
	var one ftpLatestCurveInfo
	for rows.Next() {
		err := rows.Scan(&one.domain_id, &one.curve_id, &one.as_of_date)
		if err != nil {
			logs.Error(err)
			return nil, err
		}
		rst = append(rst, one)
	}
	return rst, nil
}

func getLatestDate(rst []ftpLatestCurveInfo, domain_id string, curve_id string) string {
	for _, val := range rst {
		if val.domain_id == domain_id && val.curve_id == curve_id {
			return val.as_of_date
		}
	}
	return ""
}
