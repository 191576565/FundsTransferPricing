package ftp

import (
	"encoding/json"
	"ftpProject/conf"
	"ftpProject/logs"

	"ftpProject/utils/cacheutil"
	"ftpProject/utils/dohttpclient"
)

type CacheData struct {
	ResData  []FtpResData
	OrgData  []OrgInfo
	CucyData []CurrencyInfo
}
type FtpResData struct {
	Res_id    string
	Res_name  string
	Res_url   string
	Res_up_id string
	Res_css   string
	Res_color string
	Res_icon  string
	Res_type  string
}
type TransData struct {
	Domain_id     string
	Domain_name   string
	Org_unit_id   string
	Org_unit_desc string
	User_id       string
	User_name     string
	Role_id       string
	Role_name     string
	Rest_items    []FtpResData
}
type OrgInfo struct {
	Uuid          string
	Org_unit_id   string
	Org_unit_desc string
	Org_up_uuid   string
}
type CurrencyInfo struct {
	IsoCurrencyCd   string
	IsoCurrencyDesc string
}

//func makeOrginfo() []OrgInfo {
//	var one OrgInfo
//	var rst []OrgInfo
//	//1	10000000	江北分行	00000000

//	one.OrgUintID = "10000000"
//	one.OrgUpID = "00000000"
//	one.OrgUintDesc = "江北分行"
//	rst = append(rst, one)
//	one.OrgUintID = "10100000"
//	one.OrgUpID = "10000000"
//	one.OrgUintDesc = "观音桥支行"
//	rst = append(rst, one)
//	one.OrgUintID = "10200000"
//	one.OrgUpID = "10000000"
//	one.OrgUintDesc = "鲤鱼池支行"
//	rst = append(rst, one)
//	one.OrgUintID = "10300000"
//	one.OrgUpID = "10000000"
//	one.OrgUintDesc = "黄泥磅支行"
//	rst = append(rst, one)
//	return rst
//}
func makeCurrencyInfo() []CurrencyInfo {
	var one CurrencyInfo
	var rst []CurrencyInfo
	one.IsoCurrencyCd = "THB"
	one.IsoCurrencyDesc = "泰铢"
	rst = append(rst, one)
	one.IsoCurrencyCd = "CNY"
	one.IsoCurrencyDesc = "人民币"
	rst = append(rst, one)
	one.IsoCurrencyCd = "USD"
	one.IsoCurrencyDesc = "美元"
	rst = append(rst, one)
	one.IsoCurrencyCd = "EUR"
	one.IsoCurrencyDesc = "欧元"
	rst = append(rst, one)
	one.IsoCurrencyCd = "HKD"
	one.IsoCurrencyDesc = "港币"
	rst = append(rst, one)
	one.IsoCurrencyCd = "GBP"
	one.IsoCurrencyDesc = "英磅"
	rst = append(rst, one)
	one.IsoCurrencyCd = "JPY"
	one.IsoCurrencyDesc = "日元"
	rst = append(rst, one)
	return rst
}

type JsonData struct {
	Code string
	Msg  string
	Data TransData
	Org  []OrgInfo
}
type LoginIndexPage struct {
	RouteControl
}

//覆盖Prepare
func (this *LoginIndexPage) Prepare() {
}
func (this *LoginIndexPage) Get() {
	sid := this.GetString("sid")
	userid := this.GetString("userid")
	logs.Debug("跳转成功，sid:", sid)
	logs.Debug("跳转成功，userid:", userid)
	//请求菜单数据
	url := conf.FtpConf.AuthIp
	params := map[string]string{
		"login":   "0",
		"apitype": "1",
		"userid":  userid,
		"sid":     sid,
	}
	body, err := dohttpclient.DoHttpGet(url, params)
	if err != nil {
		logs.Error("请求菜单数据失败:", err)
		this.ShowPageError("请求菜单数据失败,请检查配置权限系统地址配置是否正确", conf.FtpConf.AuthLoginIp)
	}
	var authData JsonData

	err = json.Unmarshal(body, &authData)

	if err != nil {
		logs.Error("json解析失败:", err)
		this.ShowPageError("传输数据格式有误，请联系管理员", conf.FtpConf.AuthLoginIp)
	}
	if authData.Code != "200" {
		this.ShowPageError("用户信息有误，请联系管理员", conf.FtpConf.AuthLoginIp)
	}

	if len(authData.Data.Rest_items) == 0 {
		this.ShowPageError("菜单数据为空，请确认用户所属角色已分配菜单", conf.FtpConf.AuthLoginIp)
	}
	if len(authData.Org) == 0 {
		this.ShowPageError("机构数据为空，请确认用户所属机构", conf.FtpConf.AuthLoginIp)
	}

	//session 操作
	sessmap := map[string]string{
		"datasid":  sid,
		"userid":   userid,
		"domainid": authData.Data.Domain_id,
		"doName":   authData.Data.Domain_name,
		"orgid":    authData.Data.Org_unit_id,
		"orgName":  authData.Data.Org_unit_desc,
		"roleId":   authData.Data.Role_id,
		"roleName": authData.Data.Role_name,
	}
	this.SetSession(sessmap)

	//cache 操作
	var ca CacheData
	//存入资源
	ca.ResData = authData.Data.Rest_items
	//存入机构
	ca.OrgData = authData.Org
	//存入币种  自己控制  暂不
	//ca.CucyData = makeCurrencyInfo()
	cacheutil.PutIntoBeeCache(userid, ca)

	//日志操作
	this.Userid = userid
	this.OrgName = authData.Data.Org_unit_desc
	this.RoleName = authData.Data.Role_name
	this.InsertLogToDB("登录", "用户-登录", "FTP")

	//
	this.TplName = "theme/default/index.tpl"
}
