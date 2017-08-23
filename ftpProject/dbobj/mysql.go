package dbobj

import (
	"database/sql"
	"ftpProject/conf"
	"ftpProject/logs"

	_ "github.com/go-sql-driver/mysql"
)

type dBsonMysql struct {
}

func (this *dBsonMysql) NewDB() *sql.DB {

	tns := conf.FtpConf.DbtTns
	usr := conf.FtpConf.DbUser
	pad := conf.FtpConf.DbPass

	db, err := sql.Open("mysql", usr+":"+pad+"@"+tns)
	logs.Info("mysql链接信息为:", usr+":"+pad+"@"+tns)
	if err != nil {
		logs.Fatal("打开mysql数据库失败.", err)
		return db
	}
	if err := db.Ping(); err != nil {
		logs.Error("数据库连接断开了.", err)
		return db
	}
	logs.Debug("创建mysql链接成功.")

	return db
}
