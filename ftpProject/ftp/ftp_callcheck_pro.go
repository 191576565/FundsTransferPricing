package ftp

import (
	"database/sql"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"
	"ftpProject/utils/cacheutil"

	"ftpProject/utils"
	"strconv"
	"strings"
)

type FtpOrgInfo struct {
	ReturnMsg
	RouteControl
}

type FtpOrgInfoPage struct {
	RouteControl
}

func (this *FtpOrgInfoPage) Get() {
	this.TplName = "mas/ftp/ftp_org_info.tpl"
}
func (this *FtpOrgInfo) Get() {
	param1 := "2016-10-10"
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	param2 := doName
	err := dbobj.Default.Exec("call ts01_proc_adj(to_date(:1,'YYYY-MM-DD'), :2)", param1, param2)
	if err != nil {
		logs.Error(err)
		fmt.Println("err:", err.Error())
	}
	//fmt.Println(result)
}

//
//
//
//
//
func (this *FtpOrgInfo) Post() {
}
func (this *FtpOrgInfo) Put() {
}
func (this *FtpOrgInfo) Delete() {
}

type FtpCallPCheckProc struct {
	ReturnMsg
	RouteControl
}

func (this *FtpCallPCheckProc) Post() {
	w := this.Ctx.ResponseWriter
	param1 := "2016-10-10"
	//param2 := "FTP"
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	param2 := doName
	err := dbobj.Default.Exec("call PROC_FTP_ADJ_GDATA(to_date(:1,'YYYY-MM-DD'), :2)", param1, param2)
	if err != nil {
		logs.Error(err)
		fmt.Println("err:", err.Error())
		this.ErrorCode = "0"
		this.ErrorMsg = "调用数据生成存储过程失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	err = dbobj.Default.Exec("call PROC_FTP_ADJ_CHK (to_date(:1,'YYYY-MM-DD'), :2)", param1, param2)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "调用数据校验存储过程失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "调用政策性调节项校验存储过程"
	this.InsertLogToDB(policyadjcallp, opcontent, myapp)

	this.ErrorCode = "1"
	this.ErrorMsg = "调用存储过程成功"
	this.WriteJson(w, this.ReturnMsg)
	//fmt.Println(result)
}

type FtpPCheckResult struct {
	Uuid            string
	Adj_id          string
	Adj_desc        string
	Org_unit_id     string
	Org_unit_desc   string
	Iso_currency_cd string
	Adj_dyn_dim     string
	Dyn_name        string
	Term_str        string
	Term_end        string
	Last_date       string
	Adj_bp          string
	Eff_str_date    string
	Eff_end_date    string
	Buz_str_date    string
	Buz_end_date    string
	Domain_id       string
	Domain_name     string
	Memo            string
	Cnt             int
}
type FtpPCheckResultCtl struct {
	ReturnMsg
	RouteControl
}
type FtpPCheckResultB struct {
	Uuid            []byte
	Adj_id          []byte
	Adj_desc        []byte
	Org_unit_id     []byte
	Org_unit_desc   []byte
	Iso_currency_cd []byte
	Adj_dyn_dim     []byte
	Dyn_name        []byte
	Term_str        []byte
	Term_end        []byte
	Last_date       []byte
	Adj_bp          []byte
	Eff_str_date    []byte
	Eff_end_date    []byte
	Buz_str_date    []byte
	Buz_end_date    []byte
	Domain_id       []byte
	Domain_name     []byte
	Memo            []byte
	Cnt             int
}

func (this *FtpPCheckResultCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		db   = dbobj.DefaultDB()
		msql = ""
		one  FtpPCheckResult
		all  []FtpPCheckResult
		tmp  FtpPCheckResultB
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	orgId := this.Orgid
	if orgId == "" {
		logs.Error("seesion中机构号为空")
		return
	}
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	sAdjid := r.FormValue("keyword1")
	sother := r.FormValue("keyword2")
	//拼搜索sql
	cond := " 1=1"
	var sqlparams []string
	if sAdjid != "" || sother != "" {
		sAdjid = utils.HandleSqlKey(sAdjid)
		sother = utils.HandleSqlKey(sother)
		if sAdjid == "803" || sAdjid == "805" {
			cond = "t.adj_id=" + `'` + sAdjid + `'`
			sqlparams = append(sqlparams, cond)
			if sother != "" {
				cond = "to_number(t.term_str)<" + sother + " and to_number(t.term_end)>=" + sother
				sqlparams = append(sqlparams, cond)
			}
		} else {
			if sAdjid != "" {
				cond = "t.adj_id=" + `'` + sAdjid + `'`
				sqlparams = append(sqlparams, cond)
			}
			if sother != "" {
				//				cond = "(upper(t.adj_dyn_dim) like upper("
				//				cond += "'%" + sother + `%') escape '\' or upper(t.DYN_NAME) like upper(`
				//				cond += "'%" + sother + `%') escape '\')`
				//				sqlparams = append(sqlparams, cond)
				//				cond = "(upper(t.adj_dyn_dim) like upper("
				//				cond += "'%" + sother + `%') escape '\' or upper(t.DYN_NAME) like upper(`
				//				cond += "'%" + sother + `%') escape '\')`
				cond = "(upper(t.adj_dyn_dim) like upper("
				cond += "'%" + sother + `%') escape '\')`
				//or upper(t.DYN_NAME) like upper(`
				//cond += "'%" + sother + `%') escape '\')`
				sqlparams = append(sqlparams, cond)
			}
		}
		cond = strings.Join(sqlparams, " and ")
	}
	//fmt.Println("cond", cond)
	msql = FTP_PCHECK_GET1

	var orgdata []OrgInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		orgdata = val.OrgData
	}
	//有两种方法，一种程序遍历，第二种传入sql语句  此处选择第二种orginpara
	var s []string
	for _, val := range orgdata {
		s1 := "'" + val.Uuid + "'"
		s = append(s, s1)
	}
	tmps := strings.Join(s, ",")
	ORGUINTIN := "(" + tmps + ")"
	//
	msql = strings.Replace(msql, "SSQQLL", cond, -1)
	msql = strings.Replace(msql, "ORGINPARA", ORGUINTIN, -1)

	//fmt.Println("sql", sql)
	var rows *sql.Rows
	var err error
	if db == "oracle" {
		rows, err = dbobj.Default.Query(msql, doName, offset, offset+limit)
	} else if db == "db2" {
		rows, err = dbobj.Default.Query(msql, doName, offset, offset+limit)
	}
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&tmp.Uuid,
			&tmp.Adj_id,
			&tmp.Adj_desc,
			&tmp.Org_unit_id,
			&tmp.Iso_currency_cd,
			&tmp.Adj_dyn_dim,
			&tmp.Dyn_name,
			&tmp.Term_str,
			&tmp.Term_end,
			&tmp.Last_date,
			&tmp.Adj_bp,
			&tmp.Eff_str_date,
			&tmp.Eff_end_date,
			&tmp.Buz_str_date,
			&tmp.Buz_end_date,
			&tmp.Domain_id,
			&tmp.Memo,
			&tmp.Cnt)
		if err != nil {
			logs.Error(err)
			return
		}
		//赋值
		one.Uuid = string(tmp.Uuid)
		one.Adj_id = string(tmp.Adj_id)
		one.Adj_desc = string(tmp.Adj_desc)
		one.Org_unit_id = string(tmp.Org_unit_id)
		for _, val := range orgdata {
			if val.Uuid == one.Org_unit_id {
				one.Org_unit_desc = val.Org_unit_desc
			}
		}
		//one.Org_unit_desc = this.OrgName
		one.Iso_currency_cd = string(tmp.Iso_currency_cd)
		one.Adj_dyn_dim = string(tmp.Adj_dyn_dim)
		one.Dyn_name = string(tmp.Dyn_name)
		one.Term_str = string(tmp.Term_str)
		one.Term_end = string(tmp.Term_end)
		one.Last_date = string(tmp.Last_date)
		one.Adj_bp = string(tmp.Adj_bp)
		one.Eff_str_date = string(tmp.Eff_str_date)
		one.Eff_end_date = string(tmp.Eff_end_date)
		one.Buz_str_date = string(tmp.Buz_str_date)
		one.Buz_end_date = string(tmp.Buz_end_date)
		one.Domain_id = string(tmp.Domain_id)
		one.Domain_name = this.DomainName
		one.Memo = string(tmp.Memo)
		one.Cnt = tmp.Cnt
		//
		all = append(all, one)
	}
	this.WritePage(w, one.Cnt, all)
}
