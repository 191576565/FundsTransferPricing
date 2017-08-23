package ftp

import (
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"net/http"
)

var adjnum = 5
var adjb = [5]int8{0x01, 0x02, 0x04, 0x08, 0x10}

type FtpBusiz struct {
	Busiz_id        string `json:"id"`      //业务单元编码
	Busiz_desc      string `json:"name"`    //业务单元名称
	Busiz_up_id     string `json:"pId"`     //上级业务单元名称
	Ftp_flag        string `json:"isPrice"` //是否定价
	Al_flag         string
	Busiz_type      string `json:"type"` //是否为叶子
	Ftp_method_name string //`json:"id"`       //定价方法名称
	Ftp_method_id   string //`json:"id"`       //定价方法名称
	Term_cd         string //`json:"id"`       //期限
	Term_cd_mult    string //`json:"id"`       //期限单位
	Point_val       string //`json:"id"`       //点差
	Curve_id        string //曲线编码
	Domain_id       string `json:"region"` //域
	Is_root         string `json:"isroot"`
	Expand          string `json:"open"`
	Adjment_info    string
	Level           string
	Sequnce         string
}
type FtpBusizCtl struct {
	RouteControl
}
type FtpBusiztmp struct {
	Busiz_id        []byte //业务单元编码
	Busiz_desc      []byte //业务单元名称
	Busiz_up_id     []byte //上级业务单元名称
	Ftp_flag        []byte //是否定价
	Al_flag         []byte
	Busiz_type      []byte //是否为叶子
	Ftp_method_name []byte //定价方法名称
	Ftp_method_id   []byte //定价方法ID
	Term_cd         []byte //期限
	Term_cd_mult    []byte //期限单位
	Point_val       []byte //点差
	Curve_id        []byte //曲线编码
	Domain_id       []byte //域
	Is_root         string
	Expand          string
	Adjment_info    []byte
	Level           string
	Sequnce         string
	RouteControl
}
type FtpBusizPage struct {
	RouteControl
}

func (this *FtpBusizPage) Get() {
	this.TplName = "mas/ftp/ftp_busiz.tpl"
}

//......................
type FtpBusizStruct struct {
	Busiz_id    string `json:"id"`   //业务单元编码
	Busiz_up_id string `json:"pId"`  //上级业务单元名称
	Busiz_desc  string `json:"name"` //业务单元名称
}
type FtpBusizStructCtl struct {
	RouteControl
}
type FtpBusizStructt struct {
	Busiz_id    []byte
	Busiz_up_id []byte
	Busiz_desc  []byte
	RouteControl
}

func (this *FtpBusizStructCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()

	var (
		tmp FtpBusizStructt
		one FtpBusizStruct
		rst []FtpBusizStruct
		sql = ""
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	sql = FTP_BUSIZSTRUCT_GET
	rows, err := dbobj.Default.Query(sql, doName)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询期限结构失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&tmp.Busiz_id, &tmp.Busiz_up_id, &tmp.Busiz_desc)
		one.Busiz_id = string(tmp.Busiz_id)
		one.Busiz_up_id = string(tmp.Busiz_up_id)
		one.Busiz_desc = string(tmp.Busiz_desc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("获取层级失败"))
			return
		}
		rst = append(rst, one)
	}
	if rst == nil {
		one.Busiz_id = "-1"
		one.Busiz_desc = "业务单元根节点"
		one.Busiz_up_id = "-1"
		rst = append(rst, one)
	}
	ojs, err := json.Marshal(rst)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
func (this *FtpBusizCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	// 首先查出根节点
	var (
		sql         = ""
		rootbusizid string
	)
	//增加域 2016.9.10
	doName := this.Domainid
	sql = FTP_BUSIZ_GET1
	row := dbobj.Default.QueryRow(sql, doName)
	err := row.Scan(&rootbusizid)
	if err != nil {
		logs.Error("没有对应根节点，域名为:", doName)
		err = dbobj.Default.Exec(P_DOMIAN_I1, doName, doName+"业务单元层级", "-1", "1", "0", doName, "0")

		if err != nil {
			logs.Error("插入对应根节点失败，域名为:", doName)
			w.Write([]byte("插入根节点失败，请联系管理员"))
			return
		}
		rootbusizid = doName
	}
	//	rows.Close() //9.13 add
	sql = FTP_BUSIZ_GET2
	rows, err := dbobj.Default.Query(sql, doName, rootbusizid)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询期限结构失败"))
		return
	}
	var (
		one FtpBusiz
		tmp FtpBusiztmp
		rst []FtpBusiz
	)

	for rows.Next() {
		err := rows.Scan(&tmp.Busiz_id,
			&tmp.Busiz_desc,
			&tmp.Busiz_up_id,
			&tmp.Busiz_type,
			&tmp.Ftp_flag,
			&tmp.Al_flag,
			&tmp.Ftp_method_name,
			&tmp.Ftp_method_id,
			&tmp.Term_cd,
			&tmp.Term_cd_mult,
			&tmp.Point_val,
			&tmp.Curve_id,
			&tmp.Domain_id,
			&tmp.Is_root,
			&tmp.Expand,
			&tmp.Adjment_info,
			&tmp.Level,
			&tmp.Sequnce)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询期限结构失败，请检查期限点是否存在异常"))
			return
		}
		one.Busiz_desc = string(tmp.Busiz_desc)
		one.Busiz_id = string(tmp.Busiz_id)
		one.Busiz_type = string(tmp.Busiz_type)
		one.Busiz_up_id = string(tmp.Busiz_up_id)
		one.Domain_id = string(tmp.Domain_id)
		one.Ftp_flag = string(tmp.Ftp_flag)
		one.Al_flag = string(tmp.Al_flag)
		one.Ftp_method_name = string(tmp.Ftp_method_name)
		one.Ftp_method_id = string(tmp.Ftp_method_id)
		one.Point_val = string(tmp.Point_val)
		one.Curve_id = string(tmp.Curve_id)
		one.Term_cd = string(tmp.Term_cd)
		one.Term_cd_mult = string(tmp.Term_cd_mult)
		one.Is_root = tmp.Is_root
		one.Expand = tmp.Expand
		one.Adjment_info = string(tmp.Adjment_info)
		one.Sequnce = tmp.Sequnce
		one.Level = tmp.Level
		rst = append(rst, one)
	}
	ojs, err := json.Marshal(rst)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

type FtpAdjRel struct {
	AdjId     string
	AdjTypeId string
}

func (this *FtpBusizCtl) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		//dbtp   = dbobj.DefaultDB()
		sql    = ""
		errmsg ReturnMsg
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}

	sql = FTP_BUSIZ_POST1

	busiz_id := r.FormValue("Busiz_id")
	busiz_desc := r.FormValue("Busiz_desc")
	busiz_up_id := r.FormValue("Busiz_up_id")
	busiz_type := r.FormValue("Busiz_type")
	ftp_flag := r.FormValue("Ftp_flag")
	al_flag := r.FormValue("AL_flag")
	domain_id := doName //r.FormValue("Domain_id")

	//数据库事物
	tx, _ := dbobj.Default.Begin()
	_, err := tx.Exec(sql, busiz_id, busiz_desc, busiz_up_id, ftp_flag, busiz_type, domain_id, al_flag)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "插入业务单元信息失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}

	//插入关联
	ftp_method_id := r.FormValue("ftp_method_id")
	if ftp_flag != "1" {
		curve_id := r.FormValue("curve_id")
		//curve_id := "100"
		term_cd := r.FormValue("term_cd")
		term_cd_mult := r.FormValue("term_cd_mult")
		point_val := r.FormValue("point_val")
		h_domain_id := doName //r.FormValue("Domain_id")
		//由于db2 插入""不是null  要处理空值
		if term_cd == "" {
			term_cd = "0"
		}
		if point_val == "" {
			point_val = "0"
		}
		sql = FTP_BUSIZ_POST2

		logs.Debug(busiz_id, ftp_method_id, curve_id, term_cd, term_cd_mult, point_val, h_domain_id)
		_, err = tx.Exec(sql, busiz_id, ftp_method_id, curve_id, term_cd, term_cd_mult, point_val, h_domain_id)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "插入业务单元与方法关联失败,请联系管理员"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
		//插入调节项

		sql = FTP_BUSIZ_POST3

		var all []FtpAdjRel
		//fmt.Println("json", r.FormValue("JSON"))
		mjson := []byte(r.FormValue("JSON"))
		err = json.Unmarshal(mjson, &all)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "插入调节项失败,请联系管理员"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
		for _, val := range all {
			_, err = tx.Exec(sql, busiz_id, val.AdjId, val.AdjTypeId, doName)
			if err != nil {
				logs.Error(err)
				errmsg.ErrorCode = "0"
				errmsg.ErrorMsg = "插入调节项失败,请联系管理员"
				ojs, err := json.Marshal(errmsg)
				if err != nil {
					logs.Error(err)
				}
				w.Write(ojs)
				tx.Rollback()
				return
			}
		}
	}

	//提交事物
	tx.Commit()
	//成功的日志
	opcontent := "新增业务单元编号为:" + busiz_id + " 所属域为:" + domain_id
	this.InsertLogToDB(busizadd, opcontent, myapp)
	//结束插入调节项
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "插入成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)

}
func (this *FtpBusizCtl) Put() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		//dbtp   = dbobj.DefaultDB()
		sql    = ""
		errmsg ReturnMsg
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	busiz_id := r.FormValue("Busiz_id")
	busiz_desc := r.FormValue("Busiz_desc")
	busiz_up_id := r.FormValue("Busiz_up_id")
	busiz_type := r.FormValue("Busiz_type")
	ftp_flag := r.FormValue("Ftp_flag")
	al_flag := r.FormValue("AL_flag")
	if al_flag == "" {
		al_flag = "0"
	}
	domain_id := doName //r.FormValue("Domain_id")

	sql = FTP_BUSIZ_PUT1

	tx, _ := dbobj.Default.Begin()
	_, err := tx.Exec(sql, busiz_desc, busiz_up_id, ftp_flag, busiz_type, al_flag, busiz_id, domain_id)
	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "编辑业务单元信息失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}

	//''''''''
	ftp_method_id := r.FormValue("ftp_method_id")

	h_domain_id := doName //r.FormValue("Domain_id")
	if ftp_flag == "0" {

		sql = FTP_BUSIZ_PUT2

		_, err = tx.Exec(sql, busiz_id, domain_id)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "删除业务单元与方法关联失败,请联系管理员"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
		//16.11.10
		if ftp_method_id != "104" {
			sql = FTP_BUSIZ_D5
			_, err = tx.Exec(sql, domain_id, busiz_id)
			if err != nil {
				logs.Error(err)
				errmsg.ErrorCode = "0"
				errmsg.ErrorMsg = "更新沉淀期限结构值失败,请联系管理员"
				ojs, err := json.Marshal(errmsg)
				if err != nil {
					logs.Error(err)
				}
				w.Write(ojs)
				tx.Rollback()
				return
			}
		}
		curve_id := r.FormValue("curve_id")
		term_cd := r.FormValue("term_cd")
		term_cd_mult := r.FormValue("term_cd_mult")
		point_val := r.FormValue("point_val")
		if point_val == "" {
			point_val = "0"
		}
		sql = FTP_BUSIZ_PUT3

		logs.Debug(busiz_id, ftp_method_id, curve_id, term_cd, term_cd_mult, point_val, h_domain_id)
		_, err = tx.Exec(sql, busiz_id, ftp_method_id, curve_id, term_cd, term_cd_mult, point_val, h_domain_id)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "插入业务单元与方法关联失败,请联系管理员"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}

		sql = FTP_BUSIZ_PUT4

		_, err = tx.Exec(sql, busiz_id, doName)
		if err != nil {
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "编辑调节项失败,请联系管理员"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			return
		}
		var all []FtpAdjRel
		//fmt.Println("json", r.FormValue("JSON"))
		mjson := []byte(r.FormValue("JSON"))
		err = json.Unmarshal(mjson, &all)
		if err != nil {
			logs.Error(err)
			errmsg.ErrorCode = "0"
			errmsg.ErrorMsg = "关联调节项失败,请联系管理员"
			ojs, err := json.Marshal(errmsg)
			if err != nil {
				logs.Error(err)
			}
			w.Write(ojs)
			tx.Rollback()
			return
		}
		sql = FTP_BUSIZ_PUT5
		var dadj int8 = 0
		for _, val := range all {
			switch val.AdjId {
			case "601":
				dadj |= 0x01
			case "602":
				dadj |= 0x02
			case "603":
				dadj |= 0x04
			case "604":
				dadj |= 0x08
			case "605":
				dadj |= 0x10
			default:
			}
			_, err = tx.Exec(sql, busiz_id, val.AdjId, val.AdjTypeId, doName)
			if err != nil {
				logs.Error(err)
				errmsg.ErrorCode = "0"
				errmsg.ErrorMsg = "插入调节项失败,请联系管理员"
				ojs, err := json.Marshal(errmsg)
				if err != nil {
					logs.Error(err)
				}
				w.Write(ojs)
				tx.Rollback()
				return
			}
		}
		//fmt.Println("dadj:", dadj)
		var dsql = ""
		for i := 0; i < adjnum; i++ {
			if (adjb[i] & dadj) == 0 {
				switch i {
				case 0:
					dsql = FTP_BUSIZ_D2 //`select count(*) from FTP_ADJUST_TERM_LIQUIDITY where busiz_id=:1 and domain_id=:2`
				case 1:
					sql = ``
				case 2:
					dsql = FTP_BUSIZ_D3 //`select count(*) from FTP_ADJUST_FTP_RESTORE where busiz_id=:1 and domain_id=:2`
				case 3:
					dsql = FTP_BUSIZ_D4 //`select count(*) from FTP_ADJUST_CAPITAL_RESERVES where busiz_id=:1 and domain_id=:2`
				case 4:
				default:
				}
			}
			if dsql != "" {
				var allrow = "0"
				row := tx.QueryRow(dsql, busiz_id, domain_id)
				_ = row.Scan(&allrow)

				if allrow != "0" {
					errmsg.ErrorCode = "0"
					errmsg.ErrorMsg = "更新的调节项已经配置值，请解除配置后再更新"
					ojs, err := json.Marshal(errmsg)
					if err != nil {
						logs.Error(err)
					}
					w.Write(ojs)
					tx.Rollback()
					return
				}
			}
			dsql = ""
		}
	} else { //不定价
		/*
		   		sql = FTP_BUSIZ_PUT2

		   		_, err = tx.Exec(sql, busiz_id, domain_id)
		   		if err != nil {
		   			logs.Error(err)
		   			errmsg.ErrorCode = "0"
		   			errmsg.ErrorMsg = "删除业务单元与方法关联失败"
		   			ojs, err := json.Marshal(errmsg)
		   			if err != nil {
		   				logs.Error(err)
		   			}
		   			w.Write(ojs)
		   			tx.Rollback()
		   			return
		   		}

		   		sql = FTP_BUSIZ_PUT4
		   		fmt.Println("delete 0>1")
		   		_, err = tx.Exec(sql, busiz_id, doName)
		   		if err != nil {
		   			errmsg.ErrorCode = "0"
		   			errmsg.ErrorMsg = "编辑调节项失败"
		   			ojs, err := json.Marshal(errmsg)
		   			if err != nil {
		   				logs.Error(err)
		   			}
		   			w.Write(ojs)
		   			return
		   		}

		   		//16.11.10
		   		if ftp_method_id != "104" {
		   			sql = `delete from FTP_BUSIZ_REDEMPTION_CURVE t where t.domain_id=:1 and t.busiz_id=:2`
		   			_, err = tx.Exec(sql, domain_id, busiz_id)
		   			if err != nil {
		   				logs.Error(err)
		   				errmsg.ErrorCode = "0"
		   				errmsg.ErrorMsg = "沉淀率方法值失败"
		   				ojs, err := json.Marshal(errmsg)
		   				if err != nil {
		   					logs.Error(err)
		   				}
		   				w.Write(ojs)
		   				tx.Rollback()
		   				return
		   			}
		   		}
		   		var dsql = ""
		   		for i := 0; i < adjnum; i++ {

		   			switch i {
		   			case 0:
		   				dsql = `delete from FTP_ADJUST_TERM_LIQUIDITY t
		         where (t.busiz_id,t.domain_id) in
		         (select t.busiz_id,t.domain_id from FTP_BUSIZ_INFO t where t.domain_id=:1 start with busiz_id=:2
		         connect by prior busiz_id= busiz_up_id)`
		   			case 1:
		   				dsql = ``
		   			case 2:
		   				dsql = `delete from FTP_ADJUST_FTP_RESTORE t
		         where (t.busiz_id,t.domain_id) in
		         (select t.busiz_id,t.domain_id from FTP_BUSIZ_INFO t where t.domain_id=:1 start with busiz_id=:2
		         connect by prior busiz_id= busiz_up_id)`
		   			case 3:
		   				dsql = `delete from FTP_ADJUST_CAPITAL_RESERVES t
		         where (t.busiz_id,t.domain_id) in
		         (select t.busiz_id,t.domain_id from FTP_BUSIZ_INFO t where t.domain_id=:1 start with busiz_id=:2
		         connect by prior busiz_id= busiz_up_id)`
		   			case 4:
		   			default:
		   			}

		   			if dsql != "" {
		   				fmt.Println("i,sql:", i, dsql, busiz_id, domain_id)
		   				_, err = tx.Exec(dsql, domain_id, busiz_id)
		   				if err != nil {
		   					logs.Error(err)
		   					errmsg.ErrorCode = "0"
		   					errmsg.ErrorMsg = "删除调节项值失败"
		   					ojs, err := json.Marshal(errmsg)
		   					if err != nil {
		   						logs.Error(err)
		   					}
		   					w.Write(ojs)
		   					tx.Rollback()
		   					return
		   				}
		   			}
		   			dsql = ""
		   		}
		*/
	}

	//
	tx.Commit()
	//
	opcontent := "编辑业务单元编号为:" + busiz_id + " 所属域为:" + h_domain_id
	this.InsertLogToDB(busizedit, opcontent, myapp)
	//完成编辑调节项
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "编辑成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
func (this *FtpBusizCtl) Delete() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()

	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//dbtp := dbobj.DefaultDB()
	sql := ""
	Busiz_id := r.FormValue("Busiz_id")
	Domain_id := doName //r.FormValue("Domain_id")
	var errmsg ReturnMsg
	//删除关联

	sql = FTP_BUSIZ_DELETE1

	tx, _ := dbobj.Default.Begin()
	_, err := tx.Exec(sql, Domain_id, Busiz_id)

	if err != nil {
		logs.Error(err)
		//fmt.Println(sql)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除业务单元与方法关联失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	//
	//删除关联

	sql = FTP_BUSIZ_DELETE2

	_, err = tx.Exec(sql, Domain_id, Busiz_id)

	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除调节项关联失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		tx.Rollback()
		return
	}
	//沉淀率法单独处理16.11.10
	sql = FTP_BUSIZ_D1
	_, err = tx.Exec(sql, Domain_id, Busiz_id)

	if err != nil {
		logs.Error(err)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "删除沉淀期限结构值失败,请联系管理员"
		ojs, err := json.Marshal(errmsg)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		tx.Rollback()
		return
	}
	//
	var dsql = ""
	for i := 0; i < adjnum; i++ {

		switch i {
		case 0:
			dsql = FTP_BUSIZ_D2
		case 1:
			sql = ``
		case 2:
			dsql = FTP_BUSIZ_D3
		case 3:
			dsql = FTP_BUSIZ_D4
		case 4:
		default:
		}

		if dsql != "" {
			var allrow = "0"
			row := tx.QueryRow(dsql, Busiz_id, Domain_id)
			_ = row.Scan(&allrow)
			if allrow != "0" {
				errmsg.ErrorCode = "0"
				errmsg.ErrorMsg = "删除的调节项已经配置值，请解除配置后再删除"
				ojs, err := json.Marshal(errmsg)
				if err != nil {
					logs.Error(err)
				}
				w.Write(ojs)
				tx.Rollback()
				return
			}
		}
		dsql = ""
	}
	//
	sql = FTP_BUSIZ_DELETE3
	_, err = tx.Exec(sql, Domain_id, Busiz_id)

	if err != nil {
		logs.Error(err)
		//fmt.Println(sql)
		errmsg.ErrorCode = "0"
		errmsg.ErrorMsg = "请检查该业务单元是否已配置调整项"
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
	opcontent := "删除业务单元编号为:" + Busiz_id + " 所属域为:" + Domain_id
	this.InsertLogToDB(busizdelete, opcontent, myapp)
	//
	errmsg.ErrorCode = "1"
	errmsg.ErrorMsg = "删除成功"
	ojs, err := json.Marshal(errmsg)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
