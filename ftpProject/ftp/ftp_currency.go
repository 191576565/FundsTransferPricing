package ftp

import (
	"database/sql"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"time"
)

//
type ComCurrencyPage struct {
	RouteControl
}

func (this *ComCurrencyPage) Get() {
	this.TplName = "platform/resource/Com_Currency_Page.tpl"
}

//
type ComCurrency struct {
	IsoCurrencyCd         string
	IsoCurrencyDesc       string
	CurrencyOwner         string
	EffectiveDate         string
	IsoCurrencyStatus     string
	IsoCurrencyStatusDesc string
	SortOrder             string
	MemoBk                string
}
type ComCurrencyCtl struct {
	ReturnMsg
	RouteControl
}

func (this *ComCurrencyCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		qsql = ""
		rows *sql.Rows
		err  error
	)
	r.ParseForm()
	keyword := r.FormValue("keyword")

	if keyword != "" {
		qsql = P_CURRENCY_GET1
		ss := utils.HandleSqlKey(keyword)
		searchkey := "%" + ss + "%"
		rows, err = dbobj.Default.Query(qsql, searchkey, searchkey)
	} else {
		qsql = P_CURRENCY_GET2
		rows, err = dbobj.Default.Query(qsql)
	}

	defer rows.Close()
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "查询币种信息出错"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	var (
		one ComCurrency
		rst []ComCurrency
		bk  []byte
	)
	for rows.Next() {
		err := rows.Scan(
			&one.IsoCurrencyCd,
			&one.IsoCurrencyDesc,
			&one.CurrencyOwner,
			&one.EffectiveDate,
			&one.IsoCurrencyStatus,
			&one.IsoCurrencyStatusDesc,
			&one.SortOrder,
			&bk,
		)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "取值币种信息出错"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		one.MemoBk = string(bk)
		rst = append(rst, one)
	}
	this.WriteJson(w, rst)
}
func (this *ComCurrencyCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
	)
	r.ParseForm()

	IsoCurrencyCd := r.FormValue("IsoCurrencyCd")
	IsoCurrencyDesc := r.FormValue("IsoCurrencyDesc")
	CurrencyOwner := this.Userid
	if CurrencyOwner == "" {
		logs.Error("取值session中用户失败")
		this.ErrorCode = "0"
		this.ErrorMsg = "取值session中用户失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	EffectiveDate := time.Now().Format("2006-01-02") //r.FormValue("EffectiveDate")
	IsoCurrencyStatus := r.FormValue("IsoCurrencyStatus")
	SortOrder := r.FormValue("SortOrder")
	MemoBk := r.FormValue("MemoBk")

	sql = P_CURRENCY_POST1
	err := dbobj.Default.Exec(sql,
		IsoCurrencyCd,
		IsoCurrencyDesc,
		CurrencyOwner,
		EffectiveDate,
		IsoCurrencyStatus,
		SortOrder,
		MemoBk,
	)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "新增币种信息失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	//成功的日志
	opcontent := "新增币种信息,币种编码为:" + IsoCurrencyCd
	this.InsertLogToDB("币种-新增", opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "新增币种信息成功"
	this.WriteJson(w, this.ReturnMsg)
}
func (this *ComCurrencyCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
	)
	r.ParseForm()
	IsoCurrencyCd := r.FormValue("IsoCurrencyCd")
	IsoCurrencyDesc := r.FormValue("IsoCurrencyDesc")
	IsoCurrencyStatus := r.FormValue("IsoCurrencyStatus")
	SortOrder := r.FormValue("SortOrder")
	MemoBk := r.FormValue("MemoBk")

	//增加判断11.3 IsoCurrencyStatus
	if IsoCurrencyStatus == "1" {
		var totalrow string = "0"
		qsql := P_CURRENCY_DC
		row := dbobj.Default.QueryRow(qsql, IsoCurrencyCd)
		err := row.Scan(&totalrow)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "编辑币种信息失败，请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		if totalrow != "0" {
			this.ErrorCode = "0"
			this.ErrorMsg = "该币种已被曲线引用，无法设置为失效"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	sql = P_CURRENCY_PUT1
	err := dbobj.Default.Exec(sql,
		IsoCurrencyDesc,
		IsoCurrencyStatus,
		SortOrder,
		MemoBk,
		IsoCurrencyCd,
	)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "编辑币种信息失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	//成功的日志
	opcontent := "编辑币种信息,币种编码为:" + IsoCurrencyCd
	this.InsertLogToDB("币种-编辑", opcontent, myapp)
	//
	fmt.Println("heree")
	this.ErrorCode = "1"
	this.ErrorMsg = "编辑币种信息成功"
	this.WriteJson(w, this.ReturnMsg)
}
func (this *ComCurrencyCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
	)
	r.ParseForm()
	IsoCurrencyCd := r.FormValue("IsoCurrencyCd")
	sql = P_CURRENCY_DELETE1
	err := dbobj.Default.Exec(sql, IsoCurrencyCd)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "删除失败，请查看币种是否被曲线引用"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "删除币种信息,币种编码为:" + IsoCurrencyCd
	this.InsertLogToDB("币种-删除", opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "删除币种信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

//----------------------
type CurrencyStatusCtl struct {
	RouteControl
}
type CurrencyStatus struct {
	Cstatus     string
	CstatusDesc string
}

func (this *CurrencyStatusCtl) Get() {
	w := this.Ctx.ResponseWriter

	var (
		sql = ""
		one CurrencyStatus
		all []CurrencyStatus
	)
	sql = P_CURRENCYSTATUS_GET1
	rows, err := dbobj.Default.Query(sql)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.Cstatus, &one.CstatusDesc)
		if err != nil {
			logs.Error(err)
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}
