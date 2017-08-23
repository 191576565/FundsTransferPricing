package calcLog

import (
	"fmt"
	"logs"
	"os"
	"path"
	"strconv"
	"time"
)

type clog struct {
	file string
	l    *logs.BeeLogger
}

func NewBatchLog(dispatchId string, as_of_date string, flag int) *clog {
	if flag == 0 {

		file := dispatchId + "#" + as_of_date + "#" + time.Now().Format("2006-01-02-15-04-05") + ".log"
		logpath := os.Getenv("GOLOG")
		lp := path.Join(logpath, file)
		logs.Debug(dispatchId, as_of_date, " log file in", lp)

		l := logs.NewLogger(10240)
		l.SetLogger("file", `{"filename":"`+lp+`","maxlines":0,"maxsize":0,"daily":false,"maxdays":2}`)
		l.EnableFuncCallDepth(true)
		l.Async()

		LOGLVL := os.Getenv("LOGLEVEL")
		if LOGLVL != "" {
			lid, err := strconv.Atoi(LOGLVL)
			if err != nil {
				fmt.Println("LOGLEVEL variable invalid. set default logleve is 1")
				l.SetLevel(3)
			}
			l.SetLevel(lid)
		} else {
			l.SetLevel(3)
		}

		l.SetLogFuncCallDepth(3)

		rc := new(clog)
		rc.l = l
		rc.file = file
		rc.l.Debug("start batch info")
		return rc
	} else {
		rc := new(clog)
		rc.l = logs.Lg
		return rc
	}
}

func (this *clog) Debug(v ...interface{}) {

	this.l.Debug(v...)

}
func (this *clog) Info(v ...interface{}) {

	this.l.Info(v...)

}
func (this *clog) Error(v ...interface{}) {

	this.l.Error(v...)

}
func (this *clog) Warn(v ...interface{}) {

	this.l.Warn(v...)

}

func (this *clog) Close() {
	this.l.Close()
}
