package ftp

import (
	"fmt"
	"ftpProject/dbobj"
	"runtime"
)

type TestCtl struct {
	RouteControl
}

func (this *TestCtl) Prepare() {

}
func (this *TestCtl) Get() {
	i := runtime.NumGoroutine()
	this.Data["json"] = &map[string]int{"number": i}
	this.ServeJSON()
}
func (this *TestCtl) Put() {
	tx, _ := dbobj.Default.Begin()
	_, err := tx.Exec("insert into pt_group values(?,?,?,?,?)", 22, 22, "hujian", "hujian", "hujian")
	if err != nil {
		fmt.Println("tx insert error")
		return
	}
	tx.Commit()
	this.Data["json"] = "success"
	this.ServeJSON()
}
func (this *TestCtl) Post() {
	err := dbobj.Default.Exec("insert into pt_group values(?,?,?,?,?)", 11, 11, "hujian", "hujian", "hujian")
	if err != nil {
		fmt.Println("insert error")
		return
	}
	this.Data["json"] = "success"
	this.ServeJSON()
}
