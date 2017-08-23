package ftp

import (
	"encoding/json"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"
	
	
)

type FtpRedemption struct {
	Uuid       string
	BusizId    string
	TermCd     string
	TermCdMult string
	Weight     string
	DomainId   string
}
type FtpRedemptionCtl struct {
	RouteControl
}

func (this *FtpRedemptionCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	busizid := r.FormValue("BusizId")
	domainid := doName //r.FormValue("DomainId")

	var (
		sql    = ""
		errmsg ReturnMsg
		one    FtpRedemption
		all    []FtpRedemption
	)
	sql = FTP_REDEMPTION_GET
	rows, err := dbobj.Default.Query(sql, busizid, domainid)
	if err != nil {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "查询失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	defer rows.Close()
	one.BusizId = busizid
	for rows.Next() {
		err := rows.Scan(&one.TermCd, &one.TermCdMult, &one.Weight)
		if err != nil {
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "	取值失败"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
		all = append(all, one)
	}
	ojs, err := json.Marshal(all)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *FtpRedemptionCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql    = ""
		errmsg ReturnMsg
		all    []FtpRedemption
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	busizid := r.FormValue("BusizId")
	domainid := doName //r.FormValue("DomainId")
	mjson := []byte(r.FormValue("JSON"))
	err := json.Unmarshal(mjson, &all)
	fmt.Println(all)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "json解析失败，请重新选择"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//step 1: 首先删除
	sql = FTP_REDEMPTION_PUT1
	err = dbobj.Default.Exec(sql, busizid, domainid)
	if err != nil {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除失败，请重试"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//step 2:再插入
	sql = FTP_REDEMPTION_PUT2
	for _, val := range all {
		err := dbobj.Default.Exec(sql, busizid, val.TermCd, val.TermCdMult, val.Weight, domainid)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "新增偿还值失败"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
	}
	//
	opcontent := "编辑业务单元期限点值,业务单元编号,所属域分别为:" + busizid + " " + domainid
	this.InsertLogToDB( redemptionedit, opcontent, myapp)
	//
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "编辑成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
