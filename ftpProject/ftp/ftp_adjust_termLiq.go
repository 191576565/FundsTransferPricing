package ftp

import (
	"database/sql"
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
	"strconv"
)

type FtpAdjustTermLiqPage struct {
	RouteControl
}

type FtpAdjustTermLiqInfo struct {
	//Uuid        string
	Busiz_id    string
	Busiz_desc  string
	Curve_id    string
	Curve_desc  string
	Domain_id   string
	Domain_desc string
	cnt         int
}
type FtpAdjustTermLiqInfoCtl struct {
	ReturnMsg
	RouteControl
}
type FtpAdjustTermLiqInfoT struct {
	//Uuid        []byte
	Busiz_id    []byte
	Busiz_desc  []byte
	Curve_id    []byte
	Curve_desc  []byte
	Domain_id   []byte
	Domain_desc []byte
}

func (this *FtpAdjustTermLiqPage) Get() {
	this.TplName = "mas/ftp/ftp_adjust_term_config.tpl"
}

func (this *FtpAdjustTermLiqInfoCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		qsql = ""
		rows *sql.Rows
		err  error
	)
	r.ParseForm()
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	keyword := r.FormValue("keyword")
	adjid := r.FormValue("AdjId")
	if keyword == "" {
		qsql = FTP_ADJTERML_GET1
		rows, err = dbobj.Default.Query(qsql, adjid, doName, offset, limit+offset)

	} else {
		qsql = FTP_ADJTERML_GET2
		ss := utils.HandleSqlKey(keyword)
		searchword := "%"
		searchword = searchword + ss
		searchword = searchword + "%"
		rows, err = dbobj.Default.Query(qsql, adjid, doName, searchword, searchword, offset, limit+offset)
	}

	defer rows.Close()
	if err != nil {
		logs.Info(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("获取期限流动性溢价信息失败"))
		return
	}
	var (
		one FtpAdjustTermLiqInfoT
		tmp FtpAdjustTermLiqInfo
		rst []FtpAdjustTermLiqInfo
	)
	for rows.Next() {
		err := rows.Scan(
			&one.Busiz_id,
			&one.Busiz_desc,
			&one.Curve_id,
			&one.Curve_desc,
			&one.Domain_id,
			&tmp.cnt)
		if err != nil {
			logs.Error(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询期限流动性溢价信息失败"))
			return
		}
		//tmp.Uuid = string(one.Uuid)
		tmp.Busiz_id = string(one.Busiz_id)
		tmp.Busiz_desc = string(one.Busiz_desc)
		tmp.Curve_id = string(one.Curve_id)
		tmp.Curve_desc = string(one.Curve_desc)
		tmp.Domain_id = string(one.Domain_id)
		tmp.Domain_desc = this.DomainName
		rst = append(rst, tmp)
	}
	this.WritePage(w, tmp.cnt, rst)
}

func (this *FtpAdjustTermLiqInfoCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	sql := FTP_ADJTERML_POST1
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	var rst []FtpAdjustTermLiqInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err)
	}
	//add 10.26
	var (
		one string
		all []string
	)
	sqll := FTP_ADJTERML_POST2
	for _, val := range rst {
		row := dbobj.Default.QueryRow(sqll, val.Curve_id)
		err := row.Scan(&one)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "查询曲线对应重定价频率失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		all = append(all, one)
	}
	//
	ts, _ := dbobj.Default.Begin()
	for i, val := range rst {
		_, err := ts.Exec(sql, val.Busiz_id, val.Curve_id, doName, all[i])
		if err != nil {
			ts.Rollback()
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "新增期限流动性溢价信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	ts.Commit()
	//成功的日志
	for _, val := range rst {
		opcontent := "新增期限流动性溢价信息,业务单元为：" + val.Busiz_id
		this.InsertLogToDB(termliqadd, opcontent, myapp)
		break
	}

	//
	this.ErrorCode = "1"
	this.ErrorMsg = "新增期限流动性溢价信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

func (this *FtpAdjustTermLiqInfoCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	sql := FTP_ADJTERML_DELETE1
	var rst []FtpAdjustTermLiqInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err.Error())

		this.ErrorCode = "0"
		this.ErrorMsg = "删除期限流动性溢价信息失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	for _, val := range rst {
		err := dbobj.Default.Exec(sql, val.Busiz_id)
		if err != nil {
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "删除期限流动性溢价信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	//成功的日志
	for _, val := range rst {
		opcontent := "删除期限流动性溢价信息,业务单元为：" + val.Busiz_id
		this.InsertLogToDB(termliqdelete, opcontent, myapp)
	}

	//
	this.ErrorCode = "1"
	this.ErrorMsg = "删除期限流动性溢价信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

func (this *FtpAdjustTermLiqInfoCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var sql = ""
	var rst []FtpAdjustTermLiqInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	busizid := r.FormValue("BusizId")
	if err != nil {
		logs.Error(err)
	}
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	ts, _ := dbobj.Default.Begin()

	//先删除

	sql = FTP_ADJTERML_PUT1
	if len(rst) == 0 {

		_, err := ts.Exec(sql, busizid, doName)
		if err != nil {
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "编辑期限流动性溢价信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	for _, val := range rst {
		sql = FTP_ADJTERML_PUT1
		_, err := ts.Exec(sql, val.Busiz_id, doName)
		if err != nil {
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "编辑期限流动性溢价信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
		}
		break
	}
	var (
		one string
		all []string
	)
	sqll := FTP_ADJTERML_POST2
	for _, val := range rst {
		row := dbobj.Default.QueryRow(sqll, val.Curve_id)
		err := row.Scan(&one)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "查询曲线对应重定价频率失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		all = append(all, one)
	}
	//
	sql = FTP_ADJTERML_POST1

	for i, val := range rst {
		_, err := ts.Exec(sql, val.Busiz_id, val.Curve_id, doName, all[i])
		if err != nil {
			ts.Rollback()
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "新增期限流动性溢价信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}

	ts.Commit()

	//	for _, val := range rst {

	//		opcontent := "编辑期限流动性溢价信息,业务单元为：" + val.Busiz_id
	//		this.InsertLogToDB(termliqedit, opcontent, myapp)
	//		break
	//	}
	opcontent := "编辑期限流动性溢价信息,业务单元为：" + busizid
	this.InsertLogToDB(termliqedit, opcontent, myapp)
	//
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "编辑期限流动性溢价信息成功"
	this.WriteJson(w, this.ReturnMsg)
}
