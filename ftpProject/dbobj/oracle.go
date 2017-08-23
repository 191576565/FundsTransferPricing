package dbobj

import (
	"database/sql"
	"ftpProject/conf"
	"ftpProject/logs"
	"ftpProject/utils"
	"os"
	"strings"

	_ "github.com/mattn/go-oci8"
)

type dBsonOracle struct {
}

func (this *dBsonOracle) NewDB() *sql.DB {

	nlsLang := os.Getenv("NLS_LANG")
	if !strings.HasSuffix(nlsLang, "UTF8") {
		os.Setenv("NLS_LANG", "AMERICAN_AMERICA.AL32UTF8")
	}
	logs.Info("设置语言环境 NLS_LANG = AMERICAN_AMERICA.AL32UTF8")

	tns := conf.FtpConf.DbtTns
	usr := conf.FtpConf.DbUser
	pad := conf.FtpConf.DbPass
	var err error
	if len(pad) == 24 {
		pad, err = utils.Decrypt(pad)
		if err != nil {
			logs.Error("Decrypt mysql passwd failed.")
			return nil
		}
	}

	tnsname := usr + "/" + pad + "@" + tns

	db, err := sql.Open("oci8", tnsname)
	logs.Info("数据库为链接为 :", tns)
	if err != nil {
		logs.Fatal("连接数据库失败:", err)
		return db
	}
	if len(pad) != 24 {
		psd, err := utils.Encrypt(pad)
		if err != nil {
			logs.Error("decrypt passwd failed.", psd)
		}
		psd = "\"" + psd + "\""
		//		config.Set("DB.passwd", psd)
	}
	db.SetMaxOpenConns(0)
	db.SetConnMaxLifetime(0)
	logs.Debug("创建数据库链接成功")

	if err := db.Ping(); err != nil {
		logs.Error("oracle数据库连接失败：", err)
		return nil
	}
	return db
}
