package ftp

import (
	"encoding/json"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"

	
	"ftpProject/utils"
	"net/http"
)

type FtpBaseValueCalPage struct {
	RouteControl
}

//批次信息
type FtpBaseValueCal struct {
	DispatchId     string
	DispatcName    string
	InputSouceCd   string
	OutputResultCd string
	DomainId       string
	StartOffset    string
	MaxLimit       string
}
type FtpBaseValueCalCtl struct {
	RouteControl
}

func (this *FtpBaseValueCalPage) Get() {
	this.TplName = "mas/ftp/ftp_BaseValueCalc.tpl"
}

//批次信息
func (this *FtpBaseValueCalCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	sql := FTP_DISPATCHINFO_GET1
	rows, err := dbobj.Default.Query(sql, doName)

	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询批次信息失败"))
		return
	}
	var one FtpBaseValueCal
	var rst []FtpBaseValueCal
	for rows.Next() {
		err := rows.Scan(&one.DispatchId, &one.DispatcName, &one.InputSouceCd, &one.OutputResultCd, &one.DomainId, &one.StartOffset, &one.MaxLimit)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询批次信息失败"))
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

func (this *FtpBaseValueCalCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	var errmsg ReturnMsg
	dbtp := dbobj.DefaultDB()
	sql := ""
	switch dbtp {
	case "mysql":
		sql = ``
	case "oracle":
		sql = FTP_DISPATCHINFO_POST1
	default:
	}
	dispatchid := r.FormValue("DispatchId")
	dispatchname := r.FormValue("DispatcName")
	inputsourcecd := r.FormValue("InputSouceCd")
	outputresultcd := r.FormValue("OutputResultCd")
	domainid := doName //r.FormValue("DomainId")
	startoffset := r.FormValue("StartOffset")
	maxlimit := r.FormValue("MaxLimit")
	//
	if !utils.ValidNumber(startoffset) {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "偏移值应为数字"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//
	if !utils.ValidNumber(maxlimit) {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "长度限制应为数字"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//
	err := dbobj.Default.Exec(sql, dispatchid, dispatchname, inputsourcecd, outputresultcd, domainid, startoffset, maxlimit)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "插入批次失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//
	opcontent := "新增批次,批次编码，所属域分别为:" + dispatchid + " " + domainid
	this.InsertLogToDB(batchadd, opcontent, myapp)
	//

	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "插入成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
func (this *FtpBaseValueCalCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	dbtp := dbobj.DefaultDB()
	sql := ""
	switch dbtp {
	case "mysql":
		sql = ``
	case "oracle":
		sql = FTP_DISPATCHINFO_PUT1
	default:
	}
	dispatchid := r.FormValue("DispatchId")
	dispatchname := r.FormValue("DispatcName")
	inputsourcecd := r.FormValue("InputSouceCd")
	outputresultcd := r.FormValue("OutputResultCd")
	domainid := doName //r.FormValue("DomainId")
	startoffset := r.FormValue("StartOffset")
	maxlimit := r.FormValue("MaxLimit")
	var errmsg ReturnMsg
	err := dbobj.Default.Exec(sql, dispatchname, inputsourcecd, outputresultcd, domainid, startoffset, maxlimit, dispatchid)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "更新批次失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//
	opcontent := "编辑批次,批次编码，所属域分别为:" + dispatchid + " " + domainid
	this.InsertLogToDB(batchedit, opcontent, myapp)
	//
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "更新成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
func (this *FtpBaseValueCalCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	sql := ""
	var all []string
	//fmt.Println("json", r.FormValue("JSON"))
	mjson := []byte(r.FormValue("JSON"))
	err := json.Unmarshal(mjson, &all)
	if err != nil {
		logs.Error("Json解析失败,请联系管理员")
		return
	}
	fmt.Println("all:", all)
	domainid := doName //r.FormValue("DomainId")
	tx, _ := dbobj.Default.Begin()
	var errmsg ReturnMsg
	//step 1: 先查询正在运行的批次

	for _, val := range all {
		sql = FTP_DISPATCHINFO_DELETE1
		var status string = "0"
		fmt.Println(domainid, val)
		rows, err := dbobj.Default.Query(sql, domainid, val)
		if err != nil {
			logs.Error(err)
			logs.Errorf("删除批次运行状态失败,请联系管理员")
		}
		for rows.Next() {
			err := rows.Scan(&status)
			if err != nil {
				logs.Errorf("删除批次运行状态失败")
			}
			if status == "1" {
				errmsg.ErrorCode = "0"
				errmsg.ErrorMsg = val + "正在运行,无法删除"
				ojs, err := json.Marshal(errmsg)
				if err != nil {
					logs.Error(err)
				}
				w.Write(ojs)
				return
			}
		}

	}

	//删除
	for _, val := range all {
		sql = FTP_DISPATCHINFO_DELETE2
		_, err := tx.Exec(sql, val, domainid)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "删除批次运行状态失败"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
		sql = FTP_DISPATCHINFO_DELETE3
		_, err = tx.Exec(sql, val, domainid)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "删除批次失败,请联系管理员"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
		//
		opcontent := "删除批次,批次编码，所属域分别为:" + val + " " + domainid
		this.InsertLogToDB(batchdelete, opcontent, myapp)
		//
	}

	//
	tx.Commit()

	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "删除批次成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
