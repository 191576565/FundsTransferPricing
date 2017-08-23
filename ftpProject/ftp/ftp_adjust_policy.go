package ftp

import (
	"database/sql"
	"ftpProject/dbobj"
	"ftpProject/logs"
	"ftpProject/utils/cacheutil"

	"ftpProject/utils"
	"strconv"
	"strings"
	"time"
)

type FtpAdjustPolicyPage struct {
	RouteControl
}

func (this *FtpAdjustPolicyPage) Get() {
	this.TplName = "mas/ftp/ftp_adjust_policy.tpl"
}

type FtpAdjustPolicy struct {
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
	Cnt             int
}
type FtpAdjustPolicyCtl struct {
	ReturnMsg
	RouteControl
}
type FtpAdjustPolicyB struct {
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
	Cnt             int
}

func (this *FtpAdjustPolicyCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		db   = dbobj.DefaultDB()
		msql = ""
		one  FtpAdjustPolicy
		all  []FtpAdjustPolicy
		tmp  FtpAdjustPolicyB
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
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
	msql = FTP_ADJPOLICY_GET1
	//这儿只能查看自己机构和下级机构的信息,取cache里面的data
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
		//
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
		one.Cnt = tmp.Cnt
		//
		all = append(all, one)
	}
	this.WritePage(w, one.Cnt, all)
}

//
//
//
//
//
func (this *FtpAdjustPolicyCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql = ""
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	adj_id := r.FormValue("Adj_id")
	org_unit_id := r.FormValue("Org_unit_id")
	iso_currency_cd := r.FormValue("Iso_currency_cd")
	adj_dyn_dim := r.FormValue("Adj_dyn_dim")
	term_str := r.FormValue("Term_str")
	term_end := r.FormValue("Term_end")
	last_date := time.Now().Format("2006-01-02") //r.FormValue("Last_date")
	adj_bp := r.FormValue("Adj_bp")
	eff_str_date := r.FormValue("Eff_str_date")
	eff_end_date := r.FormValue("Eff_end_date")
	//	buz_str_date := r.FormValue("Buz_str_date")
	//	buz_end_date := r.FormValue("Buz_end_date")
	buz_str_date := r.FormValue("Buz_str_date")
	if buz_str_date == "" {
		buz_str_date = "1900-01-01"
	}
	buz_end_date := r.FormValue("Buz_end_date")
	if buz_end_date == "" {
		buz_end_date = "9999-12-31"
	}
	//domain_id := r.FormValue("Domain_id")
	domain_id := doName
	sql = FTP_ADJPOLICY_POST1
	err := dbobj.Default.Exec(sql,
		adj_id,
		org_unit_id,
		iso_currency_cd,
		adj_dyn_dim,
		term_str,
		term_end,
		last_date,
		adj_bp,
		eff_str_date,
		eff_end_date,
		buz_str_date,
		buz_end_date,
		domain_id)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "插入失败,请联系管理员"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "插入政策性调节项,编码为:" + adj_id
	this.InsertLogToDB(policyadjadd, opcontent, myapp)
	//结束插入调节项
	this.ErrorCode = "1"
	this.ErrorMsg = "插入政策性调节项成功"
	this.WriteJson(w, this.ReturnMsg)

}
func (this *FtpAdjustPolicyCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql = ""
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	uuid := r.FormValue("Uuid")
	//adj_id := r.FormValue("Adj_id")
	org_unit_id := r.FormValue("Org_unit_id")
	iso_currency_cd := r.FormValue("Iso_currency_cd")
	adj_dyn_dim := r.FormValue("Adj_dyn_dim")
	term_str := r.FormValue("Term_str")
	term_end := r.FormValue("Term_end")
	last_date := time.Now().Format("2006-01-02") //r.FormValue("Last_date")
	adj_bp := r.FormValue("Adj_bp")
	eff_str_date := r.FormValue("Eff_str_date")
	eff_end_date := r.FormValue("Eff_end_date")
	buz_str_date := r.FormValue("Buz_str_date")
	if buz_str_date == "" {
		buz_str_date = "1900-01-01"
	}
	buz_end_date := r.FormValue("Buz_end_date")
	if buz_end_date == "" {
		buz_end_date = "9999-12-31"
	}
	//domain_id := r.FormValue("Domain_id")
	//domain_id := doName //r.FormValue("Domain_id")
	sql = FTP_ADJPOLICY_PUT1
	err := dbobj.Default.Exec(sql,
		org_unit_id,
		iso_currency_cd,
		adj_dyn_dim,
		term_str,
		term_end,
		last_date,
		adj_bp,
		eff_str_date,
		eff_end_date,
		buz_str_date,
		buz_end_date,
		uuid)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "编辑失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "编辑政策性调节项,唯一编码为:" + uuid
	this.InsertLogToDB(policyadjedit, opcontent, myapp)
	//结束插入调节项
	this.ErrorCode = "1"
	this.ErrorMsg = "编辑政策性调节项成功"
	this.WriteJson(w, this.ReturnMsg)
}
func (this *FtpAdjustPolicyCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql = ""
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	uuid := r.FormValue("Uuid")
	sql = FTP_ADJPOLICY_DELETE1
	err := dbobj.Default.Exec(sql, uuid, doName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "删除失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	//成功的日志
	opcontent := "删除政策性调节项,唯一编码为:" + uuid
	this.InsertLogToDB(policyadjdelete, opcontent, myapp)
	//结束插入调节项
	this.ErrorCode = "1"
	this.ErrorMsg = "删除政策性调节项成功"
	this.WriteJson(w, this.ReturnMsg)
}
