package dbobj

import (
	"database/sql"
	"fmt"
	"ftpProject/conf"
)

//--------------------

//-------------------

type dBer interface {
	NewDB() *sql.DB
}

//--------------
type dBreal struct {
	dBer
	name string
	db   *sql.DB
}

func (this *dBreal) newDBreal() *sql.DB {
	return this.NewDB()
}

func (this *dBreal) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	rows, err := this.db.Query(sql, args...)
	fmt.Println(sql, args)
	if err != nil {
		if this.db.Ping() != nil {
			this.db = this.newDBreal()
			return this.db.Query(sql, args...)
		}
	}
	return rows, err
}
func (this *dBreal) Exec(sql string, args ...interface{}) error {
	_, err := this.db.Exec(sql, args...)
	fmt.Println(sql, args)
	if err != nil {
		if this.db.Ping() != nil {
			this.db = this.NewDB()
			_, err := this.db.Exec(sql, args...)
			return err
		}
	}
	return err
}

func (this *dBreal) Begin() (*sql.Tx, error) {
	tx, err := this.db.Begin()
	if err != nil {
		if this.db.Ping() != nil {
			this.db = this.NewDB()
			return this.db.Begin()
		}
	}
	return tx, err
}

func (this *dBreal) QueryRow(sql string, args ...interface{}) *sql.Row {
	fmt.Println(sql, args)
	if this.db.Ping() != nil {
		this.db = this.NewDB()
		return this.db.QueryRow(sql, args...)
	}
	return this.db.QueryRow(sql, args...)
}
func DefaultDB() string {
	return Default.name
}

var Default dBreal

func init() {
	var dbt dBreal
	confdb := conf.FtpConf.DbType
	if confdb == "mysql" {
		var msq dBsonMysql
		dbt.dBer = &msq
		dbt.name = "mysql"
	} else if confdb == "oracle" {
		var ora dBsonOracle
		dbt.dBer = &ora
		dbt.name = "oracle"
	}
	dbt.db = dbt.newDBreal()
	Default = dbt
	//Gdb.NewDB()
}
