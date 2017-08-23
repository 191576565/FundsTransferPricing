package ftp

import (
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"net/http"
	"strings"
	"time"
)

//期限结构表
type FtpCurveSave struct {
	Struct_code  string
	Term_cd      string
	Term_cd_mult string
	Domain_id    string
	Sort_id      string
}
type FtpCurveSaveCtl struct {
	RouteControl
}

//所有期限结构，供新增时候选择
func (this *FtpCurveSaveCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		one FtpCurveSave
		rst []FtpCurveSave
		sql = FTP_CURVESAVE_GET
	)

	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询期限结构失败"))
		return
	}

	for rows.Next() {
		err := rows.Scan(&one.Struct_code, &one.Term_cd, &one.Term_cd_mult, &one.Domain_id, &one.Sort_id)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询期限结构失败，请检查期限点是否存在异常"))
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

//
func (this *FtpCurveSaveCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//	optype := r.FormValue("type")
	//	if optype == "edit" {
	//		ojs := this.PPut(w.ResponseWriter, r)
	//		w.Write(ojs)
	//	}
	var (
		errmsg ReturnMsg
		err    error
		//		dbtp   = dbobj.DefaultDB()
		sql = ""
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//

	Curve_id := r.FormValue("Curve_id")
	Curve_desc := r.FormValue("Curve_desc")
	Curve_type := r.FormValue("Curve_type") //add ftp第二期
	Iso_currency_cd := r.FormValue("Iso_currency_cd")
	Domain_id := doName                            //r.FormValue("Domain_id")
	Create_date := time.Now().Format("2006-01-02") //2016.8.14  修改编辑时候曲线日期更新
	terms := r.FormValue("terms")
	rep_id := r.FormValue("Rep_Id")
	if Curve_type != "1" {
		rep_id = ""
	}
	//
	if !utils.ValidNumber(Curve_id) {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "曲线编码错误，请输入数字"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//
	if !utils.ValidHanWord(Curve_desc) {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "曲线名称错误"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}

	tx, _ := dbobj.Default.Begin()

	//step 1:插入曲线信息

	sql = FTP_CURVESAVE_POST1

	_, err = tx.Exec(sql, Curve_id, Curve_desc, Iso_currency_cd, Create_date, Domain_id, Curve_type, rep_id)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "插入曲线信息失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//这里插入期限结构
	qixianduan := strings.Split(terms, ",")

	if qixianduan == nil {
		logs.Error(qixianduan)
		//删除期限段信息

		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "提交曲线结构点信息失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		tx.Rollback()
		return
	}

	sql = FTP_CURVESAVE_POST2

	for _, val := range qixianduan {
		//2016.12.12取消虚拟列，以便兼容oracle 11g以下
		s := Domain_id + "_" + Curve_id + "_" + val
		_, err := tx.Exec(sql, Curve_id, val, Domain_id, s)
		if err != nil {
			logs.Error(err)
			//			sql = "delete form mas_curve_info_struct_node where curve_id = :1 and domain_id = :2"
			//			dbobj.Default.Exec(sql, Curve_id, Domain_id)
			//			sql = "delete from MAS_CURVE_DEFINE where curve_id = :1 and domain_id = :2"
			//			dbobj.Default.Exec(sql, Curve_id, Domain_id)
			//返回信息
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "提交" + val + "结构点信息失败"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
	}
	tx.Commit()
	//
	opcontent := "新增曲线编号为:" + Curve_id + " 所属域为:" + Domain_id
	this.InsertLogToDB(curveadd, opcontent, myapp)
	//成功
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "提交成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)

}

//删除曲线
func (this *FtpCurveSaveCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		//		dbtp   = dbobj.DefaultDB()
		sql    = ""
		errmsg ReturnMsg
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//
	Curve_id := r.FormValue("curveCode")
	domainid := doName //r.FormValue("Domain_id")
	//检查曲线是否别引用1
	allrow := "0"
	sql = FTP_CURVE_D1
	row := dbobj.Default.QueryRow(sql, Curve_id, doName)
	row.Scan(&allrow)
	if allrow != "0" {
		logs.Error("曲线被业务单元引用")
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "该曲线已被业务单元引用，删除失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//查看曲线是否被引用2
	allrow = "0"
	sql = FTP_CURVE_D2

	row = dbobj.Default.QueryRow(sql, Curve_id, doName)
	row.Scan(&allrow)
	if allrow != "0" {
		logs.Error("曲线被期限流动性溢价调节项引用")
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "该曲线被期限流动性溢价调节项引用，删除失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//查看曲线是否被引用3
	allrow = "0"
	sql = FTP_CURVE_D3

	row = dbobj.Default.QueryRow(sql, Curve_id, doName)
	row.Scan(&allrow)
	if allrow != "0" {
		logs.Error("曲线被司库利润还原调节项引用")
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "该曲线被司库利润还原调节项引用，删除失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//step 1:删除曲线

	sql = FTP_CURVESAVE_DELETE1

	tx, _ := dbobj.Default.Begin()
	_, err := tx.Exec(sql, Curve_id, doName)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除曲线失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//删除期限结构

	sql = FTP_CURVESAVE_DELETE2

	_, err1 := tx.Exec(sql, Curve_id, doName)
	if err1 != nil {
		logs.Error(err1)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除曲线对应值失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		tx.Rollback()
		return
	}
	//删除曲线对应点值
	sql = FTP_CURVESAVE_DELETE3
	err2 := dbobj.Default.Exec(sql, domainid, Curve_id)
	if err2 != nil {
		logs.Error(err2)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除曲线对应值失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		tx.Rollback()
		return
	}
	//
	tx.Commit()
	//
	opcontent := "删除曲线编号为:" + Curve_id + " 所属域为:" + domainid
	this.InsertLogToDB(curvedelete, opcontent, myapp)
	//删除成功
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "删除曲线,期限段，期限点值成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

func (this *FtpCurveSaveCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		errmsg ReturnMsg
		err    error
		sql    = ""
		//		dbtp   = dbobj.DefaultDB()
	)
	r.ParseForm()
	//增加域 2016.9.10
	doName := this.Domainid

	Curve_id := r.FormValue("Curve_id")
	Curve_desc := r.FormValue("Curve_desc")
	Curve_type := r.FormValue("Curve_type") //add ftp第二期
	Iso_currency_cd := r.FormValue("Iso_currency_cd")
	rep_id := r.FormValue("Rep_Id")

	rep_attr := r.FormValue("Rep_Attr")

	Domain_id := doName //r.FormValue("Domain_id")
	terms := r.FormValue("terms")

	if !utils.ValidNumber(Curve_id) {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "曲线编码错误，请输入数字"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//
	if !utils.ValidHanWord(Curve_desc) {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "曲线名称错误"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	tx, _ := dbobj.Default.Begin()
	//step 1:更新曲线定义表
	sql = FTP_CURVESAVE_PUT
	_, err = tx.Exec(sql, Curve_desc, Curve_type, Iso_currency_cd, rep_id, Curve_id, Domain_id)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "更新曲线信息失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//add 2017.3.3
	//update FTP_ADJUST_TERM_LIQUIDITY set reprice_freq_range="" where curve_id=?
	sql = FTP_CURVE_U1
	_, err = tx.Exec(sql, rep_attr, Curve_id)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "更新曲线重定价频率失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//step 2:这里更新期限结构，首先删除
	qixianduan := strings.Split(terms, ",")
	sql = FTP_CURVESAVE_PUT1
	_, err = tx.Exec(sql, Curve_id, doName)
	if err != nil {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "更新期限结构信息失败"
		//??这里要删除曲线？
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		tx.Rollback()
		w.Write(ojs)
		return

	}
	if qixianduan == nil {
		logs.Error(qixianduan)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "提交曲线结构点信息失败"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		tx.Rollback()
		w.Write(ojs)

		return
	}

	sql = FTP_CURVESAVE_PUT2

	for _, val := range qixianduan {
		s := Domain_id + "_" + Curve_id + "_" + val
		_, err := tx.Exec(sql, Curve_id, val, Domain_id, s)
		if err != nil {
			logs.Error(err)
			//返回信息
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "提交" + val + "结构点信息失败"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			tx.Rollback()
			w.Write(ojs)

			return
		}
	}
	//step 3:删除曲线值
	ends := "'" + Domain_id + "_" + Curve_id + "%'"
	endsql := " and t.curve_uuid like " + ends
	targetsql := `delete from mas_curve_info t where t.curve_uuid not in `
	ss := "( "
	nlen := len(qixianduan)
	for i, val := range qixianduan {
		s := "'" + Domain_id + "_" + Curve_id + "_" + val + "'"
		ss = ss + s
		if i < nlen-1 {
			ss = ss + ","
		}
	}
	ss = ss + ")"
	targetsql = targetsql + ss + endsql
	_, err = tx.Exec(targetsql)
	if err != nil {
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除曲线值失败，请确认"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		tx.Rollback()
		w.Write(ojs)

		return
	}
	//
	tx.Commit()
	//
	opcontent := "编辑曲线编号为:" + Curve_id + " 所属域为:" + Domain_id
	this.InsertLogToDB(curveedit, opcontent, myapp)
	//成功
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "更新成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}

	w.Write(ojs)

}
