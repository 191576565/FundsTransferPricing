package ftp

import (
	"encoding/json"
	"ftpProject/conf"
	"ftpProject/dbobj"
	"ftpProject/logs"
	"ftpProject/utils/cacheutil"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
)

type ReturnMsg struct {
	ErrorCode string
	ErrorMsg  string
}

type RouteControl struct {
	Datasid    string
	Userid     string
	Roleid     string
	RoleName   string
	Domainid   string
	DomainName string
	Orgid      string
	OrgName    string
	beego.Controller
}

func (this *RouteControl) Prepare() {
	logs.Debug("uri:", this.Ctx.Request.RequestURI)
	datasid := this.GetSession("datasid")
	if datasid == "" {
		logs.Error("session data 失效")
		this.ShowPageError("登录超时，请重新登录", conf.FtpConf.AuthLoginIp)
	}
	this.Datasid = datasid
	this.Userid = this.GetSession("userid")
	this.Domainid = this.GetSession("domainid")
	this.DomainName = this.GetSession("doName")
	this.Orgid = this.GetSession("orgid")
	this.OrgName = this.GetSession("orgName")
	this.Roleid = this.GetSession("roleId")
	this.RoleName = this.GetSession("roleName")

	err := cacheutil.BeeCache.Refresh(this.Userid)
	if err != nil {
		logs.Error("cache data 失效")
		this.ShowPageError("登录超时，请重新登录", conf.FtpConf.AuthLoginIp)
	}
	logs.Debug("datasid  userid  domainid orgid is :", this.Datasid, this.Userid, this.Domainid, this.Orgid)
}

func (this *RouteControl) WritePage(w http.ResponseWriter, total int, rows interface{}) {
	type page struct {
		Total int         `json:"total"`
		Rows  interface{} `json:"rows"`
	}
	var p page
	p.Total = total
	p.Rows = rows
	//	fmt.Println(p)
	//	if rows == nil {
	//		s := `{"total":0,"rows":[]}`
	//		w.Write([]byte(s))
	//		return
	//	}
	ijs, err := json.Marshal(p)
	if err != nil {
		logs.Error(err)
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte("打包json数据失败"))
		return
	}
	w.Write(ijs)
}

func (this *RouteControl) WriteJson(w http.ResponseWriter, v interface{}) error {
	ojs, err := json.Marshal(v)
	if err != nil {
		logs.Error(err.Error())
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(err.Error()))
		return err
	}
	w.Write(ojs)
	return nil
}
func (this *RouteControl) WriteJsonStr(v interface{}) error {
	ojs, err := json.Marshal(v)
	if err != nil {
		logs.Error(err.Error())
		return err
	}
	this.Ctx.ResponseWriter.Write(ojs)
	return nil
}
func (this *RouteControl) SetSession(m map[string]string) error {
	sess := this.StartSession()
	for key, val := range m {
		err := sess.Set(key, val)
		if err != nil {
			logs.Error("set key failed in beego session", err)
			return err
		}
	}
	return nil
}
func (this *RouteControl) GetSession(key string) string {
	sess := this.StartSession()
	if sess.Get(key) != nil {
		//logs.Debug("Get info success in session :", key)
		return sess.Get(key).(string)
	} else {
		//logs.Debug("Get info failed in session, value is nil :", key)
		return ""
	}
}
func (this *RouteControl) InsertLogToDB(optype string, opcontent string, opapp string, args ...string) {
	var (
		user = ""
		role = ""
		org  = ""
	)
	if 0 == len(args) {
		user = this.Userid
		if user == "" {
			user = "无名"
		}
		role = this.RoleName
		if role == "" {
			role = "无名"
		}
		org = this.OrgName
		if org == "" {
			org = "无名"
		}
	} else if 3 == len(args) {
		user = args[0]
		role = args[1]
		org = args[2]
	} else {
		user = "不明"
		role = "不明"
		org = "不明"
	}

	opdate := time.Now().Format("2006-01-02 15:04:05")
	ip := strings.Split(this.Ctx.Request.RemoteAddr, ":")[0]
	if ip == "[" {
		ip = "localhost"
	}
	//opapp
	opapp = "FTP"
	sql := dbobj.LOG_TO_DB
	err := dbobj.Default.Exec(sql, user, org, optype, opcontent, ip, opdate, role, opapp)
	if err != nil {
		logs.Error("插入日志失败", user, org, optype, opcontent, ip, opdate, role, opapp)
		logs.Error(err)
		return
	}
}
func (this *RouteControl) ShowPageError(content, url string) {
	//	params := map[string]string{
	//		"login":   "1",
	//		"apitype": "1",
	//		"userid":  this.Userid,
	//		"sid":     this.Datasid,
	//	}

	//	urll := LogoutaddParams(url, params)
	urll := url

	this.Data["errcontent"] = content
	this.Data["loginurl"] = urll
	this.TplName = "404.tpl"
	this.Render()
	this.StopRun()
}
