package ftp

import (
	"database/sql"
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
)

type FtpAdjustTreasurePage struct {
	RouteControl
}

type FtpAdjustTreasureInfo struct {
	//Uuid        string
	Busiz_id    string
	Busiz_desc  string
	Curve_id    string
	Curve_desc  string
	Domain_id   string
	Domain_desc string
	cnt         int
}
type FtpAdjustTreasureInfoCtl struct {
	ReturnMsg
	RouteControl
}
type FtpAdjustTreasureInfoT struct {
	//Uuid        []byte
	Busiz_id    []byte
	Busiz_desc  []byte
	Curve_id    []byte
	Curve_desc  []byte
	Domain_id   []byte
	Domain_desc []byte
}

func (this *FtpAdjustTreasurePage) Get() {
	this.TplName = "mas/ftp/ftp_adjust_treasurer_config.tpl"
}

func (this *FtpAdjustTreasureInfoCtl) Get() {
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
	keyword := r.FormValue("keyword")

	adjid := r.FormValue("AdjId")
	if keyword == "" {
		qsql = FTP_ADJTPRO_GET1
		rows, err = dbobj.Default.Query(qsql, adjid, doName)

	} else {
		qsql = FTP_ADJTPRO_GET2
		ss := utils.HandleSqlKey(keyword)
		searchword := "%"
		searchword = searchword + ss
		searchword = searchword + "%"
		rows, err = dbobj.Default.Query(qsql, adjid, doName, searchword, searchword)
	}

	defer rows.Close()
	if err != nil {
		logs.Info(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("获取司库利润信息失败"))
		return
	}
	var (
		one FtpAdjustTreasureInfoT
		tmp FtpAdjustTreasureInfo
		rst []FtpAdjustTreasureInfo
	)
	for rows.Next() {
		err := rows.Scan(
			&one.Busiz_id,
			&one.Busiz_desc,
			&one.Curve_id,
			&one.Curve_desc,
			&one.Domain_id,
			&tmp.cnt)
		//one.Domain_desc = this.DomainName
		if err != nil {
			logs.Error(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询司库利润信息失败"))
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

func (this *FtpAdjustTreasureInfoCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	sql := FTP_ADJTPRO_POST1
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	var rst []FtpAdjustTreasureInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err)
	}
	ts, _ := dbobj.Default.Begin()
	for _, val := range rst {
		_, err := ts.Exec(sql, val.Busiz_id, val.Curve_id, doName)
		if err != nil {
			ts.Rollback()
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "新增司库利润信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	ts.Commit()
	//成功的日志
	for _, val := range rst {
		opcontent := "新增司库利润信息,业务单元为：" + val.Busiz_id
		this.InsertLogToDB(treasureadd, opcontent, myapp)
		break
	}

	//
	this.ErrorCode = "1"
	this.ErrorMsg = "新增司库利润信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

func (this *FtpAdjustTreasureInfoCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	sql := FTP_ADJTPRO_DELETE1
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	var rst []FtpAdjustTreasureInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err.Error())
		this.ErrorCode = "0"
		this.ErrorMsg = "删除司库利润信息失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	for _, val := range rst {
		err := dbobj.Default.Exec(sql, val.Busiz_id, doName)
		if err != nil {
			logs.Error(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("删除司库利润信息失败"))
			return
		}
	}
	//成功的日志
	for _, val := range rst {
		opcontent := "删除删除司库利润信息,业务单元为：" + val.Busiz_id
		this.InsertLogToDB(treasuredelete, opcontent, myapp)
	}

	//
	this.ErrorCode = "1"
	this.ErrorMsg = "删除司库利润信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

func (this *FtpAdjustTreasureInfoCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var sql = ""
	var rst []FtpAdjustTreasureInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err)
	}
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	busizid := r.FormValue("BusizId")
	ts, _ := dbobj.Default.Begin()
	sql = FTP_ADJTPRO_PUT1
	if len(rst) == 0 {
		_, err := ts.Exec(sql, busizid, doName)
		if err != nil {
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "编辑司库利润信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}

	//先删除
	for _, val := range rst {
		sql = FTP_ADJTPRO_PUT1
		_, err := ts.Exec(sql, val.Busiz_id, doName)
		if err != nil {
			logs.Error(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("编辑司库利润信息失败"))
			this.ErrorCode = "0"
			this.ErrorMsg = "编辑司库利润信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
		}
		break
	}
	sql = FTP_ADJTPRO_PUT2

	for _, val := range rst {
		_, err := ts.Exec(sql, val.Busiz_id, val.Curve_id, doName)
		if err != nil {
			ts.Rollback()
			logs.Error(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("编辑司库利润信息失败"))
			this.ErrorCode = "0"
			this.ErrorMsg = "编辑司库利润信息失败"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	ts.Commit()
	//	for _, val := range rst {
	//		opcontent := "编辑司库利润信息,业务单元为：" + val.Busiz_id
	//		this.InsertLogToDB(treasureedit, opcontent, myapp)
	//		break
	//	}
	opcontent := "编辑司库利润信息,业务单元为：" + busizid
	this.InsertLogToDB(treasureedit, opcontent, myapp)

	//
	this.ErrorCode = "1"
	this.ErrorMsg = "编辑司库利润信息成功"
	this.WriteJson(w, this.ReturnMsg)
}
