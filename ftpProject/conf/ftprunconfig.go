package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/config"
)

const (
	FtpLogKey      = "FTP.log"
	SessTimeKey    = "SessionGCMaxLifetime" //"Session.time"
	DbKey          = "Default_Db"
	DbTypeKey      = "DB.type"
	DbtTnsKey      = "DB.tns"
	DbUserKey      = "DB.user"
	DbPassKey      = "DB.passwd"
	CalcIpKey      = "Calc.IP"
	SessTypeKey    = "Session.Type"
	LogLvlKey      = "Log.level"
	AuthIpKey      = "Auth.ip"
	AuthLoginIpKey = "Auth.loginip"
)

type FtpRunConf struct {
	FtpLog      string
	SessTime    string
	Db          string
	DbType      string
	DbtTns      string
	DbUser      string
	DbPass      string
	CalcIp      string
	SessType    string
	LogLvl      string
	AuthIp      string
	AuthLoginIp string
}

var FtpConf FtpRunConf

func init() {
	fmt.Println("[init] 初始化配置文件中参数")
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	config, err := config.NewConfig("ini", appConfigPath)
	if err != nil {
		panic(err)
	}

	//
	FtpConf.FtpLog = config.String(FtpLogKey)
	if FtpConf.FtpLog == "" {
		panic(err)
	}
	//
	FtpConf.SessTime = config.String(SessTimeKey)
	if FtpConf.SessTime == "" {
		panic(err)
	}
	//
	//	FtpConf.Db = config.String(DbKey)
	//	if FtpConf.Db == "" {
	//		panic(err)
	//	}
	//
	FtpConf.DbType = config.String(DbTypeKey)
	if FtpConf.DbType == "" {
		panic(err)
	}
	//
	FtpConf.DbtTns = config.String(DbtTnsKey)
	if FtpConf.DbtTns == "" {
		panic(err)
	}
	//
	FtpConf.DbUser = config.String(DbUserKey)
	if FtpConf.DbUser == "" {
		panic(err)
	}
	//
	FtpConf.DbPass = config.String(DbPassKey)
	if FtpConf.DbPass == "" {
		panic(err)
	}
	//
	FtpConf.CalcIp = config.String(CalcIpKey)
	if FtpConf.CalcIp == "" {
		panic(err)
	}
	//
	FtpConf.LogLvl = config.String(LogLvlKey)
	if FtpConf.LogLvl == "" {
		panic(err)
	}
	//
	FtpConf.AuthIp = config.String(AuthIpKey)
	if FtpConf.AuthIp == "" {
		panic(err)
	}
	//
	FtpConf.AuthLoginIp = config.String(AuthLoginIpKey)
	if FtpConf.AuthLoginIp == "" {
		panic(err)
	}
}
