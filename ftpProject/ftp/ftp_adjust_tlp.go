package ftp

import (
	"ftpProject/dbobj"
	"ftpProject/logs"
	
	"net/http"
)

type AdjustTlp struct {
	Curve_id   string
	Curve_desc string
}
type AdjustTlpCtl struct {
	RouteControl
}

func (this *AdjustTlpCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	ctype := r.FormValue("CurveType")
	sql := FTP_ADJTLP_GET1
	rows, err := dbobj.Default.Query(sql, ctype, doName)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询TLP曲线信息失败"))
		return
	}
	var one AdjustTlp
	var rst []AdjustTlp
	for rows.Next() {
		err := rows.Scan(&one.Curve_id,
			&one.Curve_desc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询TLP曲线信息失败"))
			return
		}
		rst = append(rst, one)
	}
	this.WriteJson(w, rst)
}
