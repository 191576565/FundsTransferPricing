package ftp

import (
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
)

type FtpAdjInfoPage struct {
	RouteControl
}

func (this *FtpAdjInfoPage) Get() {
	this.TplName = "mas/ftp/ftp_adjust_info.tpl"
}

//---------------------------
type FtpAdjInfo struct {
	AdjId         string
	AdjDesc       string
	AdjType       string
	AdjTypeDesc   string
	AdjStatus     string
	AdjStatusDesc string
}
type FtpAdjInfoCtl struct {
	ReturnMsg
	RouteControl
}

func (this *FtpAdjInfoCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql     = ""
		keyword = ""
	)
	keyword = r.FormValue("keyword")
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//fmt.Println("domain", doName)
	sql = FTP_ADJINFO_GET1
	rows, err := dbobj.Default.Query(sql, doName)
	if keyword != "" {
		sql = FTP_ADJINFO_GET2
		ss := utils.HandleSqlKey(r.FormValue("keyword"))
		searchword := "%"
		searchword = searchword + ss
		searchword = searchword + "%"
		rows, err = dbobj.Default.Query(sql, doName, searchword, searchword)
	}
	defer rows.Close()
	if err != nil {
		logs.Error(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("获取调节项信息失败"))
		return
	}
	var one FtpAdjInfo
	var rst []FtpAdjInfo
	for rows.Next() {
		err := rows.Scan(&one.AdjId,
			&one.AdjDesc,
			&one.AdjType,
			&one.AdjTypeDesc,
			&one.AdjStatus,
			&one.AdjStatusDesc,
		)
		if err != nil {
			logs.Error(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte(err.Error()))
			return
		}
		rst = append(rst, one)
	}
	this.WriteJson(w, rst)
}
func (this *FtpAdjInfoCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql = ""
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	adjid := r.FormValue("AdjId")
	adjdesc := r.FormValue("AdjDesc")
	adjtypeid := r.FormValue("AdjTypeId")
	adjstatus := r.FormValue("AdjStatus")
	sql = FTP_ADJINFO_POST1
	err := dbobj.Default.Exec(sql, adjid, adjdesc, adjtypeid, adjstatus, doName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "插入调节项失败，请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "插入调节项,编码为:" + adjid
	this.InsertLogToDB(adjadd, opcontent, myapp)
	//结束插入调节项
	this.ErrorCode = "1"
	this.ErrorMsg = "插入调节项成功"
	this.WriteJson(w, this.ReturnMsg)
}
func (this *FtpAdjInfoCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql = ""
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	adjid := r.FormValue("AdjId")
	adjdesc := r.FormValue("AdjDesc")
	adjtypeid := r.FormValue("AdjTypeId")
	adjstatus := r.FormValue("AdjStatus")
	if adjstatus == "1" {
		var totalrow = "0"
		qsql := FTP_ADJINFO_P
		row := dbobj.Default.QueryRow(qsql, adjid, doName, adjid, doName)
		err := row.Scan(&totalrow)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "编辑调节项信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		if totalrow != "0" {
			this.ErrorCode = "0"
			this.ErrorMsg = "该调节项已被引用，无法设置为不启用"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	sql = FTP_ADJINFO_PUT1
	err := dbobj.Default.Exec(sql, adjdesc, adjtypeid, adjstatus, adjid, doName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "编辑调节项失败，请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "编辑调节项,编码为:" + adjid
	this.InsertLogToDB(adjedit, opcontent, myapp)
	//结束插入调节项
	this.ErrorCode = "1"
	this.ErrorMsg = "编辑调节项成功"
	this.WriteJson(w, this.ReturnMsg)
}
func (this *FtpAdjInfoCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql = ""
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}

	adjid := r.FormValue("AdjId")
	if true {
		var totalrow = "0"
		qsql := FTP_ADJINFO_D
		row := dbobj.Default.QueryRow(qsql, adjid, doName, adjid, doName)
		err := row.Scan(&totalrow)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "删除调节项信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		if totalrow != "0" {
			this.ErrorCode = "0"
			this.ErrorMsg = "该调节项已被引用，无法删除"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	sql = FTP_ADJINFO_DELETE1
	err := dbobj.Default.Exec(sql, adjid, doName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "删除调节项失败，请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "删除调节项,编码为:" + adjid
	this.InsertLogToDB(adjdelete, opcontent, myapp)
	//结束插入调节项
	this.ErrorCode = "1"
	this.ErrorMsg = "删除调节项成功"
	this.WriteJson(w, this.ReturnMsg)
}
