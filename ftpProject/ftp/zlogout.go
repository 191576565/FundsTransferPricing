package ftp

import (
	"ftpProject/conf"
	"ftpProject/logs"
	"net/url"
	"strings"
)

const (
	userlogintext = "用户-登录"
	userlogoutext = "用户-登出"
)

type LogOutCtl struct {
	RouteControl
}

func (this *LogOutCtl) Get() {

	opcontent := "退出登录,用户名为:" + this.Userid
	this.InsertLogToDB(userlogoutext, opcontent, "平台")
	//清除session
	this.DestroySession()
	//cache不清除，如果清除了，影响其他在线用户，cache会自动失效
	logs.Debug("redirect to :", conf.FtpConf.AuthLoginIp)
	//this.Redirect(conf.FtpConf.AuthLoginIp, 302)
	//this.StopRun()
	//this.StopRun()
	//	params := map[string]string{
	//		"login":   "1",
	//		"apitype": "1",
	//		"userid":  this.Userid,
	//		"sid":     this.Datasid,
	//	}

	//url := LogoutaddParams(conf.FtpConf.AuthLoginIp)
	//fmt.Println("userlogout:", url)
	url := conf.FtpConf.AuthLoginIp
	this.Ctx.WriteString(url)
}
func paramsToString(params map[string]string) string {
	values := url.Values{}
	for k, v := range params {
		values.Set(k, v)
	}

	return values.Encode()
}
func LogoutaddParams(url_ string, params map[string]string) string {
	if len(params) == 0 {
		return url_
	}

	if !strings.Contains(url_, "?") {
		url_ += "?"
	}

	if strings.HasSuffix(url_, "?") || strings.HasSuffix(url_, "&") {
		url_ += paramsToString(params)
	} else {
		url_ += "&" + paramsToString(params)
	}

	return url_
}

//func (this *LogOutCtl) Get() {

//	opcontent := "退出登录,用户名为:" + this.Userid
//	this.InsertLogToDB(userlogoutext, opcontent, "平台")
//	//清除session
//	this.DestroySession()
//	//cache不清除，如果清除了，影响其他在线用户，cache会自动失效
//	logs.Debug("redirect to :", conf.FtpConf.AuthLoginIp)
//	//this.Redirect(conf.FtpConf.AuthLoginIp, 302)
//	//this.StopRun()
//	//this.StopRun()
//	this.Ctx.WriteString(conf.FtpConf.AuthLoginIp)
//}
