package ftp

import (
	"database/sql"
	"ftpProject/dbobj"
	"ftpProject/logs"
	
	
	"net/http"
	"time"
)

type FtpProductInfo struct {
	ProductId         string
	ProductName       string
	ProductParentId   string
	ProductParentName string
	CreationTime      string
	Creater           string
	DomainId          string
	Memo              string
	Level             string
}
type FtpProductInfoCtl struct {
	ReturnMsg
	RouteControl
}
type FtpProductInfoPage struct {
	RouteControl
}

func (this *FtpProductInfoPage) Get() {
	this.TplName = "mas/ftp/ftp_product_info.tpl"
}
func (this *FtpProductInfoCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one FtpProductInfo
		all []FtpProductInfo
	)

	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}

	sql = FTP_PRODUCTINFO_GET1
	rows, err := dbobj.Default.Query(sql, doName, doName)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	var (
		bk    []byte
		pname []byte
	)
	for rows.Next() {
		err := rows.Scan(
			&one.ProductId,
			&one.ProductName,
			&one.ProductParentId,
			&pname,
			&one.CreationTime,
			&one.Creater,
			&one.DomainId,
			&bk,
			&one.Level)
		if err != nil {
			logs.Error(err)
			return
		}
		one.Memo = string(bk)
		one.ProductParentName = string(pname)
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

func (this *FtpProductInfoCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
	)
	r.ParseForm()
	sql = FTP_PRODUCTINFO_POST1
	ProductId := r.FormValue("ProductId")
	ProductName := r.FormValue("ProductName")
	ProductParentId := r.FormValue("ProductParentId")
	CreationTime := time.Now().Format("2006-01-02")
	Creater := this.Userid
	if Creater == "" {
		logs.Error("session中userid为空，退出")
		return
	}
	//	Creater := "hujian"
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	//DomainId := r.FormValue("DomainId")
	Memo := r.FormValue("Memo")
	err := dbobj.Default.Exec(sql, ProductId, ProductName, ProductParentId, CreationTime, Creater, doName, Memo)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "新增产品信息失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "新增产品信息,产品ID为：" + ProductId
	this.InsertLogToDB( productadd, opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "新增产品信息成功"
	this.WriteJson(w, this.ReturnMsg)
}
func (this *FtpProductInfoCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
	)
	r.ParseForm()
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}

	sql = FTP_PRODUCTINFO_PUT1
	ProductId := r.FormValue("ProductId")
	ProductName := r.FormValue("ProductName")
	ProductParentId := r.FormValue("ProductParentId")
	Memo := r.FormValue("Memo")
	err := dbobj.Default.Exec(sql, ProductName, ProductParentId, Memo, ProductId, doName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "编辑产品信息失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "编辑产品信息,产品ID为：" + ProductId
	this.InsertLogToDB( productedit, opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "编辑产品信息成功"
	this.WriteJson(w, this.ReturnMsg)
}
func (this *FtpProductInfoCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		msql = ""
		db   = dbobj.DefaultDB()
	)
	r.ParseForm()
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	ProductId := r.FormValue("ProductId")
	//add 11.4
	var totalrow = "0"
	qsql := FTP_PRDT_D1
	//*sql.Row
	var row *sql.Row
	if db == "oracle" {
		row = dbobj.Default.QueryRow(qsql, doName, doName, ProductId)
	} else if db == "db2" {
		row = dbobj.Default.QueryRow(qsql, ProductId, doName, doName, doName)
	}
	err := row.Scan(&totalrow)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "删除产品信息失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	if totalrow != "0" {
		this.ErrorCode = "0"
		this.ErrorMsg = "该产品或下级产品已被引用，删除失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	msql = FTP_PRODUCTINFO_DELETE1
	err = dbobj.Default.Exec(msql, ProductId, doName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "删除产品信息失败，请检查是否有用户已分配该产品"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "删除产品信息,产品ID为：" + ProductId
	this.InsertLogToDB( productdelete, opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "删除产品信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

//tree
//机构树
type ProductTree struct {
	ProductId       string `json:"id"`
	ProductParentId string `json:"pId"`
	ProductName     string `json:"name"`
}
type ProductTreeCtl struct {
	RouteControl
}

func (this *ProductTreeCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		sql = ""

		one ProductTree
		all []ProductTree
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	sql = FTP_PRDTREE_G1

	rows, err := dbobj.Default.Query(sql, doName)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询产品树失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.ProductId, &one.ProductParentId, &one.ProductName)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("获取产品树层级失败"))
			return
		}
		all = append(all, one)
	}
	if all == nil {
		one.ProductId = "-1"
		one.ProductName = "产品信息根节点"
		one.ProductParentId = "-1"
		all = append(all, one)
	}
	this.WriteJson(w, all)
}
