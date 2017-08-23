package ftp

import (
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
)

type FtpPolicyAdjustDefPage struct {
	RouteControl
}

func (this *FtpPolicyAdjustDefPage) Get() {
	this.TplName = "mas/ftp/ftp_adjust_list_policy.tpl"
}

//------------------------
type FtpAdjustDefPage struct {
	RouteControl
}

func (this *FtpAdjustDefPage) Get() {
	this.TplName = "mas/ftp/ftp_adjust_list_inner.tpl"
}

//------------------------
type FtpAdjustDef struct {
	Adjustment_id        string
	Adjustment_name      string
	Adjustment_type      string
	Adjustment_type_desc string
}
type FtpAdjustDefCtl struct {
	RouteControl
}

func (this *FtpAdjustDefCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	typeId := r.FormValue("TypeId")
	keyword := r.FormValue("keyword")
	sql := FTP_ADJLIST_GET1

	rows, err := dbobj.Default.Query(sql, typeId, doName)
	if keyword != "" {
		word := utils.HandleSqlKey(keyword)
		searchkey := "%" + word + "%"
		//fmt.Println("s:", searchkey)
		sql = FTP_ADJLIST_G1
		rows, err = dbobj.Default.Query(sql, typeId, doName, searchkey, searchkey)
	}
	defer rows.Close()
	if err != nil {
		logs.Error(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("获取调节项信息失败"))
		return
	}
	var one FtpAdjustDef
	var rst []FtpAdjustDef
	for rows.Next() {
		err := rows.Scan(&one.Adjustment_id,
			&one.Adjustment_name,
			&one.Adjustment_type,
			&one.Adjustment_type_desc)
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
