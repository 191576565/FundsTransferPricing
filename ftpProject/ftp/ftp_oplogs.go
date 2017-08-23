package ftp

import (
	"ftpProject/dbobj"
	"ftpProject/logs"
	"ftpProject/utils"
	"ftpProject/utils/cacheutil"
	"strconv"
	"strings"
)

type HandleLogPage struct {
	RouteControl
}

type HandleLog struct {
	OpUserId  string
	OpOrg     string
	OpApp     string
	OpType    string
	OpContent string
	OpIp      string
	OpDate    string
	cnt       int
}
type HandleLogCtl struct {
	RouteControl
}

func (this *HandleLogCtl) Get() {
	sql := P_HLOG_GET1
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	searchId := r.FormValue("SearchTrue")
	userId := r.FormValue("UserId")
	startDate := r.FormValue("StartDate")
	endDate := r.FormValue("EndDate")
	optype := r.FormValue("OpType")
	opapp := r.FormValue("OpApp")
	roletype := r.FormValue("RoleType")
	offset, _ := strconv.Atoi(r.FormValue("offset"))
	limit, _ := strconv.Atoi(r.FormValue("limit"))
	cond := " 1=1 "
	var sqlparams []string
	targets := " 1=1 "
	if searchId != "" {
		if userId != "" {
			userId = utils.HandleSqlKey(userId)
			cond = " op_user_id = '" + userId + "'"
			sqlparams = append(sqlparams, cond)
		}
		if startDate != "" {
			//cond = " to_timestamp(op_date,'YYYY-MM-DD HH24:MI:SS') >= to_timestamp('" + startDate + "','YYYY-MM-DD HH24:MI:SS')"
			cond = ` TO_DATE(TO_CHAR(to_timestamp(op_date,'YYYY-MM-DD HH24:MI:SS'),'YYYYMMDD') ,'YYYYMMDD')  BETWEEN  DATE` + "'" + startDate + "'"
			sqlparams = append(sqlparams, cond)
		}
		if endDate != "" {
			//cond = " to_timestamp(op_date,'YYYY-MM-DD HH24:MI:SS') <= to_timestamp('" + endDate + "','YYYY-MM-DD HH24:MI:SS')"
			cond = `DATE` + "'" + endDate + "'"
			sqlparams = append(sqlparams, cond)
		}
		if optype != "" {
			optype = utils.HandleSqlKey(optype)
			cond = " op_type like '%" + optype + "%'"
			sqlparams = append(sqlparams, cond)
		}
		if roletype != "" {
			cond = " op_role = '" + roletype + "'"
			sqlparams = append(sqlparams, cond)
		}
		if opapp != "" {
			opapp = utils.HandleSqlKey(opapp)
			cond = " op_app = '" + opapp + "'"
			sqlparams = append(sqlparams, cond)
		}
		//拼接
		targets = strings.Join(sqlparams, " and ")
		if targets == "" {
			targets = " 1=1 "
		}
	}
	//这儿只能查看自己机构和下级机构的信息,取cache里面的data
	var orgdata []OrgInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		orgdata = val.OrgData
	}
	//有两种方法，一种程序遍历，第二种传入sql语句  此处选择第二种orginpara
	var s []string
	for _, val := range orgdata {
		s1 := "'" + val.Org_unit_desc + "'"
		s = append(s, s1)
	}
	tmp := strings.Join(s, ",")
	ORGUINTIN := "(" + tmp + ")"

	sql = strings.Replace(sql, "SQLPARAMS", targets, -1)
	sql = strings.Replace(sql, "ORGINPARA", ORGUINTIN, -1)

	//logs.Debug("handlog sql :", sql, offset, limit+offset)
	rows, err := dbobj.Default.Query(sql, offset, limit+offset)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		return
	}
	var one HandleLog
	var rst []HandleLog
	for rows.Next() {
		err := rows.Scan(
			&one.OpUserId,
			&one.OpOrg,
			&one.OpApp,
			&one.OpType,
			&one.OpContent,
			&one.OpIp,
			&one.OpDate,
			&one.cnt)
		if err != nil {
			logs.Error(err)
			return
		}
		//fmt.Println("one:", one)
		rst = append(rst, one)
	}
	//fmt.Println("rst:", rst)
	//
	this.WritePage(w, one.cnt, rst)
}

func (this *HandleLogPage) Get() {
	this.TplName = "platform/resource/handle_logs_page.tpl"
}
