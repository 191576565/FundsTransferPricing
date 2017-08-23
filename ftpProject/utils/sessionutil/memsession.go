/*
*memery 内存session
*用sid启动session  sid获取session值
*
*hujian
**/
package sessionutil

import (
	"ftpProject/conf"
	"ftpProject/logs"
	"strconv"
	"time"

	//"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/session"
)

//此函数加在beego session中的mem_session.go中
//func NewMempder() *MemProvider {
//	return mempder
//}
type memConfig struct {
	Gclifetime  int64
	Maxlifetime int64
}
type mprovider struct {
	memp   *session.MemProvider
	config memConfig
}

func newMprovider(maxlifetime, gclifetime int64) *mprovider {
	m := &mprovider{
		memp: session.NewMempder(),
	}
	m.memp.SessionInit(maxlifetime, "hujian")
	m.config.Gclifetime = gclifetime
	m.config.Maxlifetime = maxlifetime
	return m
}

func (this *mprovider) start(sid string) (session session.Store, err error) {
	session, err = this.memp.SessionRead(sid)
	if err != nil {
		return nil, err
	}
	return
}
func (this *mprovider) Set(sid string, m map[string]string) error {
	sess, _ := this.start(sid)
	for key, val := range m {
		err := sess.Set(key, val)
		if err != nil {
			logs.Error("set user info failed in this session ->", err)
			return err
		}
	}
	return nil
}
func (this *mprovider) Get(sid, key string) string {
	sess, _ := this.start(sid)
	if sess.Get(key) != nil {
		logs.Debug("Get info success in session :", key)
		return sess.Get(key).(string)
	} else {
		logs.Debug("Get info failed in session, value is nil :", key)
		return ""
	}
}
func (this *mprovider) GC() {
	this.memp.SessionGC()
	time.AfterFunc(time.Duration(this.config.Gclifetime)*time.Second, func() { this.GC() })
}
func (this *mprovider) DestroyMemSess(sid string) {
	this.memp.SessionDestroy(sid)
	logs.Debug("session destroy success ! sid=", sid)
}
func (this *mprovider) IsExitAndUpdateTime(sid string) bool {
	return this.memp.SessExitUpdateTime(sid)
}

var GobalMemP *mprovider

func hhhinit() {
	//	workPath, err := os.Getwd()
	//	if err != nil {
	//		panic(err)
	//	}
	//	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	//	config, err := config.NewConfig("ini", appConfigPath)
	//	if err != nil {
	//		logs.Error(err)
	//	}
	//	sessiontime, _ := config.Int64("Session.time")
	sessiontime, err := strconv.ParseInt(conf.FtpConf.SessTime, 10, 64)
	if err != nil {
		panic(err)
	}
	logs.Debug("session time is : ", sessiontime)
	GobalMemP = newMprovider(sessiontime, sessiontime)
	go GobalMemP.GC()
}
