package ftp

import (
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
)

//曲线点值结构
type FtpCurveInfoStruct struct {
	CurveId    string
	StructCode string
	DomainId   string
	Uuid       string
}
type FtpCurveInfoStructCtl struct {
	RouteControl
}
type StructYelid struct {
	StructCode  string
	StructValue string
}

func (this *FtpCurveInfoStructCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	curveid := r.FormValue("curve_id")
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	sql := FTP_CURVEDATA_GET
	rows, err := dbobj.Default.Query(sql, curveid, doName)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询期限结构失败,请联系管理员"))
		return
	}
	var one FtpCurveInfoStruct
	var rst []FtpCurveInfoStruct
	for rows.Next() {
		err := rows.Scan(&one.CurveId, &one.StructCode, &one.DomainId, &one.Uuid)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询期限结构失败，请检查期限点是否存在异常"))
			return
		}
		rst = append(rst, one)
	}
	ojs, err := json.Marshal(rst)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
func (this *FtpCurveInfoStructCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//	optype := r.FormValue("type")
	//	if optype == "edit" {
	//		this.PPut(w.ResponseWriter, r)
	//		return
	//	}
	var (
		err    error
		errmsg ReturnMsg
		all    []StructYelid
		sql    = ""
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}

	curveid := r.FormValue("Curve_Id")
	domainid := doName //r.FormValue("Domain_Id")
	asofdate := r.FormValue("date")
	mjson := []byte(r.FormValue("JSON"))
	err = json.Unmarshal(mjson, &all)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "新增失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	if !utils.ValidDate(asofdate) {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "曲线日期错误,应为YYYY-MM-DD"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	valueflag := false
	for _, val := range all {
		if val.StructValue == "" {
			continue
		}
		valueflag = true
		if !utils.ValidBalance(val.StructValue) {
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "期限点值错误，应该为小数或者整数"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
	}
	if valueflag == false {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "请至少填写一个期限点值"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	tx, _ := dbobj.Default.Begin()
	//先查重
	var allrows string = "0"
	sql = FTP_CURVE_C1
	row := tx.QueryRow(sql, doName, curveid, asofdate)
	err = row.Scan(&allrows)
	if err != nil {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "新增失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	if allrows != "0" {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "已存在相同日期的数据，请重新选择日期"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//

	sql = FTP_CURVEDATA_POST2
	for _, val := range all {
		curveuuid := domainid + "_" + curveid + "_" + val.StructCode
		logs.Info(curveuuid, asofdate, val.StructValue)
		_, err = tx.Exec(sql, curveuuid, asofdate, val.StructValue)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "新增曲线值失败,请联系管理员"
			tx.Rollback()
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
	}
	//
	tx.Commit()
	//
	opcontent := "新增曲线点值,对应曲线编号,日期,所属域为:" + curveid + " " + asofdate + " " + domainid
	this.InsertLogToDB(curvedadd, opcontent, myapp)
	//
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "新增曲线值成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

type CurveInfoDelete struct {
	Domain_Id string
	Curve_Id  string
	Date      []string
}

func (this *FtpCurveInfoStructCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var errmsg ReturnMsg
	//	dbtp := dbobj.DefaultDB()
	sql := ""

	sql = FTP_CURVEDATA_DELETE

	var all CurveInfoDelete
	mjson := []byte(r.FormValue("JSON"))
	err := json.Unmarshal(mjson, &all)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}

	domainid := all.Domain_Id
	curveid := all.Curve_Id
	tx, _ := dbobj.Default.Begin()
	for _, d := range all.Date {
		_, err := tx.Exec(sql, domainid, curveid, d)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "请先到业务单元配置中解除曲线引用后，再删除曲线"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
		opcontent := "删除曲线点值,对应曲线编号,日期,所属域为:" + curveid + " " + d + " " + domainid
		this.InsertLogToDB(curveddelete, opcontent, myapp)
	}
	//t
	tx.Commit()
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "删除曲线成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *FtpCurveInfoStructCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//	if sys.Privilege.Access(w, r) == false {
	//		w.WriteHeader(http.StatusForbidden)
	//		return
	//	}
	curveid := r.FormValue("Curve_Id")
	domainid := r.FormValue("Domain_Id")
	asofdate := r.FormValue("date")
	mjson := []byte(r.FormValue("JSON"))

	var (
		err    error
		errmsg ReturnMsg
		all    []StructYelid
	)
	err = json.Unmarshal(mjson, &all)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "JSON解析失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}

	for _, val := range all {
		if val.StructValue == "" {
			continue
		}
		if !utils.ValidBalance(val.StructValue) {
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "期限点值错误，应该为小数或者整数"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
	}
	tx, _ := dbobj.Default.Begin()
	sql := FTP_CURVEDATA_PUT1
	_, err = tx.Exec(sql, domainid, curveid, asofdate)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "更新曲线值失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}

	sql = FTP_CURVEDATA_PUT2
	for _, val := range all {
		curveuuid := domainid + "_" + curveid + "_" + val.StructCode
		_, err = tx.Exec(sql, curveuuid, asofdate, val.StructValue)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "更新曲线值失败"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	//
	opcontent := "编辑曲线点值,对应曲线编号,日期,所属域为:" + curveid + " " + asofdate + " " + domainid
	this.InsertLogToDB(curvededit, opcontent, myapp)
	//
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "更新曲线值成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
