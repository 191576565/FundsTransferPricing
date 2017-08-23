package ftp

import (
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"
	"ftpProject/utils/cacheutil"

	"net/http"
	"strings"
)

const (
	myapp          = "内部资金转移定价"               // Level 1  Added by lt 20160918   #!#:1
	curve_def_mgr  = myapp + "-曲线定义-"         // Level 2  Added by lt 20160918   #!#:1
	curve_val_mgr  = curve_def_mgr + "-曲线点值-" // Level 3  Added by lt 20160918   #!#:1
	bbusiz_mgr     = myapp + "-定价规则-"         // Level 2  Added by lt 20160918   #!#:1
	price_calc_mgr = myapp + "-基础价格计算-"
	valuecalcrun   = myapp + "-单笔试算"
	adjinfo_mgr    = myapp + "-调节项"
	policyadj_mgr  = myapp + "-政策性调节项"
	product_mgr    = myapp + "-产品信息"
	zbj_mgr        = myapp + "-准备金"
	termliq_mgr    = myapp + "-期限流动性溢价"
	treasure_mgr   = myapp + "-司库利润还原"
	all_input      = myapp + "-全量导入"
	//
	xlsxcalc = myapp + "-整体试算-手工计算"
	xlsxauto = myapp + "-整体试算-系统计算"
	bexport  = myapp + "-业务方案-导出"
	//
	busizadd    = bbusiz_mgr + "新增"
	busizedit   = bbusiz_mgr + "编辑"
	busizdelete = bbusiz_mgr + "删除"
	//
	curveadd    = curve_def_mgr + "新增"
	curveedit   = curve_def_mgr + "编辑"
	curvedelete = curve_def_mgr + "删除"
	//
	curvedadd      = curve_val_mgr + "新增"
	curvededit     = curve_val_mgr + "编辑"
	curveddelete   = curve_val_mgr + "删除"
	curveinfoinput = curve_val_mgr + "增量导入"
	//
	batchrun    = price_calc_mgr + "批次运行"
	batchadd    = price_calc_mgr + "批次新增"
	batchedit   = price_calc_mgr + "批次编辑"
	batchdelete = price_calc_mgr + "批次删除"
	batchstop   = price_calc_mgr + "批次停止"
	batchclear  = price_calc_mgr + "批次清除"
	//
	redemptionedit = "期限点值编辑"
	//调节项
	adjadd    = adjinfo_mgr + "新增"
	adjedit   = adjinfo_mgr + "编辑"
	adjdelete = adjinfo_mgr + "删除"
	//政策性调节项
	policyadjadd    = policyadj_mgr + "新增"
	policyadjedit   = policyadj_mgr + "编辑"
	policyadjdelete = policyadj_mgr + "删除"
	policyadjexport = policyadj_mgr + "导出"
	policyadjimport = policyadj_mgr + "导入"
	//--调用存储过程
	policyadjcallp = policyadj_mgr + "调用校验存储过程"
	//产品信息
	productadd    = product_mgr + "新增"
	productedit   = product_mgr + "编辑"
	productdelete = product_mgr + "删除"
	//准备金
	zbjadd    = zbj_mgr + "新增"
	zbjedit   = zbj_mgr + "编辑"
	zbjdelete = zbj_mgr + "删除"
	//期限流动性溢价
	termliqadd    = termliq_mgr + "新增"
	termliqedit   = termliq_mgr + "编辑"
	termliqdelete = termliq_mgr + "删除"
	//司库利润
	treasureadd    = treasure_mgr + "新增"
	treasureedit   = treasure_mgr + "编辑"
	treasuredelete = treasure_mgr + "删除"
)

//方法
type FtpSysMethod struct {
	FtpMethodId   string
	FtpMethodName string
}
type FtpSysMethodCtl struct {
	RouteControl
}

func (this *FtpSysMethodCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//取业务单元
	sql := `select ftp_method_id,ftp_method_desc from ftp_sys_method`
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询方法表失败"))
		return
	}
	var one FtpSysMethod
	var all []FtpSysMethod
	for rows.Next() {
		err := rows.Scan(&one.FtpMethodId, &one.FtpMethodName)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值方法表失败"))
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

//域名
type SysDomainInfo struct {
	DomainId   string
	DomainName string
	RouteControl
}

func (this *SysDomainInfo) Get() {
	w := this.Ctx.ResponseWriter
	useroleid := "" //sys.Privilege.Get(w, r, "roleId")
	rolestage := ""
	//
	rolestagetmp := "" //sys.Privilege.Get(w, r, "roleStage")
	rolestagearry := strings.Split(rolestagetmp, ",")
	for _, val := range rolestagearry {
		rolestage = val
		break
	}
	if rolestage == "" {
		logs.Error("取值用户所属角色层失败,请重新登录")
		return
	}
	//

	if rolestage <= "1000" {
		//取域
		sql := `select domain_id,domain_name from sys_domain_info`
		rows, err := dbobj.Default.Query(sql)
		defer rows.Close()
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("查询域名表失败"))
			return
		}
		var one SysDomainInfo
		var all []SysDomainInfo
		for rows.Next() {
			err := rows.Scan(&one.DomainId, &one.DomainName)
			if err != nil {
				logs.Error(err)
				w.WriteHeader(http.StatusExpectationFailed)
				w.Write([]byte("取值域名表失败"))
				return
			}
			all = append(all, one)
		}
		ojs, err := json.Marshal(all)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)

	} else {
		doid := this.Domainid
		if useroleid == "" {
			logs.Error("取值用户所域失败,请重新登录")
			return
		}
		doname := "" //sys.Privilege.Get(w, r, "domainDesc")
		if useroleid == "" {
			logs.Error("取值用户所域失败,请重新登录")
			return
		}
		var one SysDomainInfo
		var all []SysDomainInfo
		one.DomainId = doid
		one.DomainName = doname
		all = append(all, one)
		ojs, err := json.Marshal(all)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)

	}

}

//币种
type MasDimCurrency struct {
	IsoCurrencyCd   string
	IsoCurrencyName string
}
type MasDimCurrencyCtl struct {
	RouteControl
}

func (this *MasDimCurrencyCtl) Get() {
	w := this.Ctx.ResponseWriter
	sql := `select iso_currency_cd,iso_currency_desc from mas_dim_currency where status='0' order by sort_id`
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询币种表失败"))
		return
	}
	var one MasDimCurrency
	var all []MasDimCurrency
	for rows.Next() {
		err := rows.Scan(&one.IsoCurrencyCd, &one.IsoCurrencyName)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值币种表失败"))
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
func (this *MasDimCurrencyCtl) Get_old() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//取业务单元
	var curr []CurrencyInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		curr = val.CucyData
	}
	var one MasDimCurrency
	var all []MasDimCurrency
	for _, val := range curr {
		one.IsoCurrencyCd = val.IsoCurrencyCd
		one.IsoCurrencyName = val.IsoCurrencyDesc
		all = append(all, one)
	}
	ojs, err := json.Marshal(all)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}

//曲线表
type MasCurveDefine struct {
	CurveId         string
	CurveDesc       string
	IsoCurrencyDesc string
	CreateDate      string
	DomainId        string
}
type MasCurveDefineCtl struct {
	RouteControl
}

func (this *MasCurveDefineCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//取业务单元
	sql := FTP_CURVEDEFINE_G1
	rows, err := dbobj.Default.Query(sql, doName)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询曲线表失败"))
		return
	}
	var one MasCurveDefine
	var all []MasCurveDefine
	for rows.Next() {
		err := rows.Scan(&one.CurveId, &one.CurveDesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值曲线信息失败"))
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

//基础价格计算动态输入
type FtpDispatchInput struct {
	InputSourceCd   string
	InputSourceDesc string
}
type FtpDispatchInputCtl struct {
	RouteControl
}

func (this *FtpDispatchInputCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//取业务单元
	sql := `select Input_Source_Cd,Input_Source_Desc from ftp_dispatch_input_conf`
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询批次输入表失败"))
		return
	}
	var one FtpDispatchInput
	var all []FtpDispatchInput
	for rows.Next() {
		err := rows.Scan(&one.InputSourceCd, &one.InputSourceDesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值批次输入表失败"))
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

//基础价格计算动态输出chu
type FtpDispatchOutput struct {
	OutputResultCD   string
	OutputResultDesc string
}
type FtpDispatchOutputCtl struct {
	RouteControl
}

func (this *FtpDispatchOutputCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//取业务单元
	sql := `select output_result_cd,output_result_desc from ftp_dispatch_output_conf`
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询批次输出表失败"))
		return
	}
	var one FtpDispatchOutput
	var all []FtpDispatchOutput
	for rows.Next() {
		err := rows.Scan(&one.OutputResultCD, &one.OutputResultDesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值批次输出表失败"))
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

//调节项动态框
type FtpAdjustment struct {
	AdjustmentId   string
	AdjustmentName string
	AdjtypeId      string
}
type FtpAdjustmentCtl struct {
	RouteControl
}

func (this *FtpAdjustmentCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	sql := FTP_ADJ_G1
	rows, err := dbobj.Default.Query(sql, doName)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询调节项表失败"))
		return
	}
	var one FtpAdjustment
	var all []FtpAdjustment
	for rows.Next() {
		err := rows.Scan(&one.AdjustmentId, &one.AdjustmentName, &one.AdjtypeId)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值调节项表失败"))
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

//曲线类型
type FtpCurveType struct {
	CuType     string
	CuTypeDesc string
}
type FtpCurveTypeCtl struct {
	RouteControl
}

func (this *FtpCurveTypeCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	//取业务单元
	sql := `select t.curve_type,t.curve_type_desc from MAS_CURVE_TYPE t order by t.curve_type`
	rows, err := dbobj.Default.Query(sql)
	defer rows.Close()
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询曲线类型表失败"))
		return
	}
	var one FtpCurveType
	var all []FtpCurveType
	for rows.Next() {
		err := rows.Scan(&one.CuType, &one.CuTypeDesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值曲线类型失败"))
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

//调节项状态框
type FtpAdjStatus struct {
	Astatus     string
	AstatusDesc string
}
type FtpAdjStatusCtl struct {
	RouteControl
}

func (this *FtpAdjStatusCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one FtpAdjStatus
		all []FtpAdjStatus
	)
	sql = `select t.status,t.status_desc from ftp_adj_status t order by t.status desc`
	rows, err := dbobj.Default.Query(sql)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询调节项状态表失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.Astatus, &one.AstatusDesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值调节项状态表失败"))
			return
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

//调节项类型
type FtpAdjType struct {
	Atypeid   string
	Atypename string
}
type FtpAdjTypeCtl struct {
	RouteControl
}

func (this *FtpAdjTypeCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one FtpAdjType
		all []FtpAdjType
	)
	sql = `select t.adj_type_id,t.adj_type_name from FTP_ADJ_TYPE t`
	rows, err := dbobj.Default.Query(sql)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询调节项类型表失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.Atypeid, &one.Atypename)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值调节项类型表失败"))
			return
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

//产品负债标识
type FtpAlType struct {
	Altypeid   string
	Altypedesc string
}
type FtpAlTypeCtl struct {
	RouteControl
}

func (this *FtpAlTypeCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one FtpAlType
		all []FtpAlType
	)
	sql = `select t.AL_FLAG,t.al_FLAG_desc from FTP_AL_TYPE t`
	rows, err := dbobj.Default.Query(sql)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询产品负债标识表失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.Altypeid, &one.Altypedesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值调节项类型表失败"))
			return
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

//行业
//
//业务单元含有政策调节项的
type FtpBusizPolicy struct {
	BusizId   string
	BusizDesc string
}
type FtpBusizPolicyCtl struct {
	RouteControl
}

func (this *FtpBusizPolicyCtl) Get() {

	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one FtpBusizPolicy
		all []FtpBusizPolicy
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	sql = FTP_BP_G1
	rows, err := dbobj.Default.Query(sql, doName)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询业务单元和政策性关联表失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.BusizId, &one.BusizDesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值业务单元和政策性关联表失败"))
			return
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

//产业树
type IndustryTree struct {
	IndustryId       string `json:"id"`
	IndustryParentId string `json:"pId"`
	IndustryName     string `json:"name"`
}
type IndustryTreeCtl struct {
	RouteControl
}

func (this *IndustryTreeCtl) Get() {

	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one IndustryTree
		all []IndustryTree
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	sql = FTP_INDUSTREE_G1

	rows, err := dbobj.Default.Query(sql, doName)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询行业树失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.IndustryId, &one.IndustryParentId, &one.IndustryName)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("获取行业树层级失败"))
			return
		}
		all = append(all, one)
	}
	if all == nil {
		one.IndustryId = "-1"
		one.IndustryName = "行业信息根节点"
		one.IndustryParentId = "-1"
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

//顶级结构  top
type TopOrg struct {
	OrgUnitId   string
	OrgUnitDesc string
}
type TopOrgCtl struct {
	RouteControl
}

func (this *TopOrgCtl) Get() {
	var orgdata []OrgInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		orgdata = val.OrgData
	}
	var one TopOrg
	if orgdata != nil {
		one.OrgUnitId = orgdata[0].Uuid
		one.OrgUnitDesc = orgdata[0].Org_unit_desc
	}
	this.WriteJsonStr(&one)
}

//重定价频率下拉框
//调节项类型
type FtpRepType struct {
	FtpRepId   string
	FtpRepDesc string
	FtpRepAttr string
}
type FtpRepTypeCtl struct {
	RouteControl
}

func (this *FtpRepTypeCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one FtpRepType
		all []FtpRepType
	)
	sql = `select t.rep_id, t.rep_desc,t.rep_calc_attr from FTP_REPRICE_FREQ t order by t.rep_id`
	rows, err := dbobj.Default.Query(sql)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询重定价类型失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.FtpRepId, &one.FtpRepDesc, &one.FtpRepAttr)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值重定价类型失败"))
			return
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

//准备金的业务单元下拉框

type FtpReBusiz struct {
	BusizId   string
	BusizDesc string
}
type FtpReBusizCtl struct {
	RouteControl
}

func (this *FtpReBusizCtl) Get() {
	w := this.Ctx.ResponseWriter
	var (
		sql = ""
		one FtpReBusiz
		all []FtpReBusiz
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	sql = FTP_REB_G1
	rows, err := dbobj.Default.Query(sql, doName)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("查询重定价类型失败"))
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&one.BusizId, &one.BusizDesc)
		if err != nil {
			logs.Error(err)
			w.WriteHeader(http.StatusExpectationFailed)
			w.Write([]byte("取值重定价类型失败"))
			return
		}
		all = append(all, one)
	}
	this.WriteJson(w, all)
}

type OrgTree struct {
	OrgUintID   string `json:"id"`
	OrgUintDesc string `json:"name"`
	OrgUpID     string `json:"pId"`
}
type OrgTreeCtl struct {
	RouteControl
}

func (this *OrgTreeCtl) Get() {
	var orgdata []OrgInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		orgdata = val.OrgData
	}
	var one OrgTree
	var rst []OrgTree
	if orgdata != nil {
		for _, val := range orgdata {
			one.OrgUintDesc = val.Org_unit_desc
			one.OrgUintID = val.Uuid
			one.OrgUpID = val.Org_up_uuid
			rst = append(rst, one)
		}
	}
	this.WriteJsonStr(&rst)
}
