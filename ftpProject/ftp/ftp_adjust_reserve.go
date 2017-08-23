package ftp

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
	"strconv"
	"time"
)

type FtpAdjustReservePage struct {
	RouteControl
}

type FtpAdjustReserveInfo struct {
	Uuid            string
	Busiz_id        string
	Busiz_desc      string
	Reserve_percent string
	Reserve_rate    string
	Eff_start       string
	Eff_end         string
	Domain_id       string
	Domain_desc     string
	cnt             int
}
type FtpAdjustReserveInfoCtl struct {
	ReturnMsg
	RouteControl
}
type FtpAdjustReserveInfoT struct {
	Uuid            []byte
	Busiz_id        []byte
	Busiz_desc      []byte
	Reserve_percent []byte
	Reserve_rate    []byte
	Eff_start       []byte
	Eff_end         []byte
	Domain_id       []byte
	Domain_desc     []byte
}

func (this *FtpAdjustReservePage) Get() {
	this.TplName = "mas/ftp/ftp_adjust_reserve.tpl"
}

func (this *FtpAdjustReserveInfoCtl) Get() {
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

	if keyword == "" {
		qsql = FTP_ADJRESERVE_GET1
		rows, err = dbobj.Default.Query(qsql, doName, offset, limit+offset)

	} else {
		qsql = FTP_ADJRESERVE_GET2
		ss := utils.HandleSqlKey(keyword)
		searchword := "%"
		searchword = searchword + ss
		searchword = searchword + "%"
		rows, err = dbobj.Default.Query(qsql, doName, searchword, searchword, offset, limit+offset)
	}

	defer rows.Close()
	if err != nil {
		logs.Info(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("获取准备金调节项信息失败"))
		return
	}
	var (
		tmp FtpAdjustReserveInfo
		one FtpAdjustReserveInfoT
		rst []FtpAdjustReserveInfo
	)
	for rows.Next() {
		err := rows.Scan(&one.Uuid,
			&one.Busiz_id,
			&one.Busiz_desc,
			&one.Reserve_percent,
			&one.Reserve_rate,
			&one.Eff_start,
			&one.Eff_end,
			&one.Domain_id,
			&tmp.cnt)
		if err != nil {
			logs.Error(err.Error())
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询准备金调节项信息失败"))
			return
		}
		tmp.Uuid = string(one.Uuid)
		tmp.Busiz_id = string(one.Busiz_id)
		tmp.Busiz_desc = string(one.Busiz_desc)
		tmp.Reserve_percent = string(one.Reserve_percent)
		tmp.Reserve_rate = string(one.Reserve_rate)
		tmp.Eff_start = string(one.Eff_start)
		tmp.Eff_end = string(one.Eff_end)
		tmp.Domain_id = string(one.Domain_id)
		tmp.Domain_desc = this.DomainName
		rst = append(rst, tmp)
	}
	this.WritePage(w, tmp.cnt, rst)
}

//新增时校验
type checkdate struct {
	startdate string
	enddate   string
}

func (this *FtpAdjustReserveInfoCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	sql := FTP_ADJRESERVE_POST1
	bi := r.FormValue("busiz_id")
	rp := r.FormValue("reserve_percent")
	rr := r.FormValue("reserve_rate")
	strdate := r.FormValue("str_date")
	enddate := r.FormValue("end_date")
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	did := doName
	//新增时校验
	sqlcheck := FTP_ADJREVERSE_ADDC
	var one checkdate
	var all []checkdate
	rows, cerr := dbobj.Default.Query(sqlcheck, did, bi)
	if cerr != nil {
		logs.Error(cerr)
		this.ErrorCode = "0"
		this.ErrorMsg = "新增准备金配置信息失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	for rows.Next() {
		err := rows.Scan(&one.startdate, &one.enddate)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "新增准备金配置信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		all = append(all, one)
	}
	cstart, _ := time.Parse("2006-01-02", strdate)
	cend, _ := time.Parse("2006-01-02", enddate)
	for _, val := range all {
		repflag := false
		t1, _ := time.Parse("2006-01-02", val.startdate)
		t2, _ := time.Parse("2006-01-02", val.enddate)
		fmt.Println(cstart, cend, t1, t2)
		if (cstart.Before(t1) || cstart.Equal(t1)) && (cend.After(t1) || cend.Equal(t1)) {
			repflag = true
		} else if cstart.Before(t2) && (cstart.After(t1) || cstart.Equal(t1)) {
			repflag = true
		} else if cend.After(t1) && (cend.Before(t2) || cend.Equal(t2)) {
			repflag = true
		}
		if repflag == true {
			logs.Error("开始日期结束日期有重复")
			this.ErrorCode = "0"
			this.ErrorMsg = "新增失败,开始日期或结束日期与已有的配置存在重复，请检查"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	//
	err := dbobj.Default.Exec(sql, bi, rp, rr, strdate, enddate, did)
	if err != nil {
		logs.Error(err.Error())
		this.ErrorCode = "0"
		this.ErrorMsg = "新增准备金配置信息失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "新增准备金配置信息,业务单元为：" + bi
	this.InsertLogToDB(zbjadd, opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "新增准备金配置信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

func (this *FtpAdjustReserveInfoCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	sql := FTP_ADJRESERVE_DELETE1
	var bi = ""
	var rst []FtpAdjustReserveInfo
	err := json.Unmarshal([]byte(r.FormValue("JSON")), &rst)
	if err != nil {
		logs.Error(err.Error())
		this.ErrorCode = "0"
		this.ErrorMsg = "删除准备金配置信息失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	for _, val := range rst {
		bi = val.Busiz_id
		err := dbobj.Default.Exec(sql, val.Uuid)
		if err != nil {
			logs.Error(err.Error())
			this.ErrorCode = "0"
			this.ErrorMsg = "删除准备金配置信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	//成功的日志
	opcontent := "删除准备金配置信息,业务单元为：" + bi
	this.InsertLogToDB(zbjdelete, opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "删除准备金配置信息成功"
	this.WriteJson(w, this.ReturnMsg)
}

func (this *FtpAdjustReserveInfoCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	sql := FTP_ADJRESERVE_PUT1
	uuid := r.FormValue("Uuid")
	bi := r.FormValue("busiz_id")
	rp := r.FormValue("reserve_percent")
	rr := r.FormValue("reserve_rate")
	strdate := r.FormValue("str_date")
	enddate := r.FormValue("end_date")
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	did := doName

	//新增时校验
	sqlcheck := FTP_ADJREVERSE_DEC
	var one checkdate
	var all []checkdate
	rows, cerr := dbobj.Default.Query(sqlcheck, did, bi, uuid)
	if cerr != nil {
		logs.Error(cerr)
		this.ErrorCode = "0"
		this.ErrorMsg = "编辑准备金配置信息失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	for rows.Next() {
		err := rows.Scan(&one.startdate, &one.enddate)
		if err != nil {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "新增准备金配置信息失败,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		all = append(all, one)
	}
	cstart, _ := time.Parse("2006-01-02", strdate)
	cend, _ := time.Parse("2006-01-02", enddate)
	for _, val := range all {
		repflag := false
		t1, _ := time.Parse("2006-01-02", val.startdate)
		t2, _ := time.Parse("2006-01-02", val.enddate)

		if (cstart.Before(t1) || cstart.Equal(t1)) && (cend.After(t1) || cend.Equal(t1)) {
			repflag = true
		} else if cstart.Before(t2) && (cstart.After(t1) || cstart.Equal(t1)) {
			repflag = true
		} else if cend.After(t1) && (cend.Before(t2) || cend.Equal(t2)) {
			repflag = true
		}
		if repflag == true {
			logs.Error("开始日期结束日期有重复")
			this.ErrorCode = "0"
			this.ErrorMsg = "新增失败,开始日期或结束日期与已有的配置存在重复，请检查"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}

	err := dbobj.Default.Exec(sql, rp, rr, strdate, enddate, uuid)
	if err != nil {
		logs.Error(err.Error())
		this.ErrorCode = "0"
		this.ErrorMsg = "编辑准备金配置信息失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	//成功的日志
	opcontent := "编辑准备金配置信息,业务单元为：" + bi
	this.InsertLogToDB(zbjedit, opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "编辑准备金配置信息成功"

	this.WriteJson(w, this.ReturnMsg)
}
