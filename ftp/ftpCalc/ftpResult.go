// Copyright 2016 huangzhanwei. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ftpCalc

import (
	"dbobj"
	"fmt"
	"logs"
)

//type ftpCalcResult struct {
//	domain_id      string
//	as_of_date     string
//	account_number string
//	ftp_rate       string
//	ajust_01       []byte
//	ajust_02       []byte
//}

//type dispatchUpdate struct {
//	dispatch_id      string
//	input_source_cd  string
//	output_source_cd string
//}

type ftpResult struct {
	domain_id      string
	as_of_date     string
	account_number string
	ftp_rate       string
	ajust_01       string
	ajust_02       string
	ajust_03       string
	ajust_04       string
	ajust_05       string
}

//var rstCache = make(chan ftpResult, 10240)
//var ftprst *ftpCalcResult

//func (this *ftpCalcResult) save(v1, v2, v3, v4, v5, v6 string) {
//	sql := "insert /*+APPEND*/ into ftp_calc_result(domain_id,as_of_date,account_number,ftp_rate,ajust_01,ajust_02) values(:1,to_date(:2,'YYYY-MM-DD'),:3,:4,:5,:6)"
//	err := dbobj.Default.Exec(sql, v1, v2, v3, v4, v5, v6)
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//}

//func (this *ftpCalcResult) del() {
//	sql := "truncate table ftp_calc_result"
//	err := dbobj.Default.Exec(sql)
//	if err != nil {
//		logs.Error(err)
//		return
//	}
//}

//func UpdateFtpRate(v1, v2, v3, v4, v5, v6 string) {
//	//ftprst.save(v1, v2, v3, v4, v5, v6)
//	var one ftpResult
//	one.domain_id = v1
//	one.as_of_date = v2
//	one.account_number = v3
//	one.ftp_rate = v4
//	one.ajust_01 = v5
//	one.ajust_02 = v6
//	rstCache <- one
//}

//func ClearFtpRate() {
//	ftprst.del()
//}

//func insertData() {
//	sql := "insert /*+APPEND*/ into ftp_calc_result nologging (domain_id,as_of_date,account_number,ftp_rate,ajust_01,ajust_02) values(:1,to_date(:2,'YYYY-MM-DD'),:3,:4,:5,:6)"
//	var index = 0
//	tx, _ := dbobj.Default.Begin()
//	for {
//		select {
//		case val := <-rstCache:
//			_, err := tx.Exec(sql,
//				val.domain_id,
//				val.as_of_date,
//				val.account_number,
//				val.ftp_rate,
//				val.ajust_01,
//				val.ajust_02)
//			if err != nil {
//				fmt.Println(err.Error())
//			} else {
//				index++
//			}
//			if index > 300 {
//				go tx.Commit()
//				tx, _ = dbobj.Default.Begin()
//				index = 0
//			}
//		case <-time.After(time.Second * 2):
//			if index > 0 {
//				go tx.Commit()
//				index = 0
//				tx, _ = dbobj.Default.Begin()
//			}
//		}
//	}
//}

func updateDispatchStatus(dispatchId, as_of_date string, status string) {

	msg := "未定义。"
	if status == "1" {
		msg = "批次正在运行中。"
	} else if status == "2" {
		msg = "批次异常终止，请查看日志文件。"
	} else if status == "3" {
		msg = "批次运行已完成。"
	} else if status == "4" {
		msg = "批次被手动停止。"
	} else if status == "5" {
		msg = "批次运行已完成"
	}
	err := dbobj.Default.Exec(FTPCALC_RESULT, status, msg, dispatchId, as_of_date)
	if err != nil {
		fmt.Println(err)
		logs.Error(err)
	}
}

//func updateResultRows() {
//	//get output list
//	//
//	fmt.Println("init ftp result rows")
//	for {
//		sql := `select t.dispatch_id,t.input_source_cd,t.output_result_cd from FTP_DISPATCH_LIST t`
//		rows, err := dbobj.Default.Query(sql)
//		defer rows.Close()
//		if err != nil {
//			logs.Error(err)
//			return
//		}
//		var one dispatchUpdate
//		var rst []dispatchUpdate
//		for rows.Next() {
//			err := rows.Scan(&one.dispatch_id,
//				&one.input_source_cd,
//				&one.output_source_cd)
//			if err != nil {
//				logs.Error(err)
//			}
//			rst = append(rst, one)
//		}
//		for _, val := range rst {
//			sql = `select as_of_date,count(*) as cnt from ` + val.output_source_cd + ` group by as_of_date`
//			type line struct {
//				as_of_date string
//				cnt        int
//			}
//			rows, err := dbobj.Default.Query(sql)
//			if err != nil {
//				logs.Error(err)
//			}
//			var lineone line
//			var linerst []line
//			for rows.Next() {
//				err := rows.Scan(&lineone.as_of_date,
//					&lineone.cnt)
//				if err != nil {
//					logs.Error(err)
//				}
//				linerst = append(linerst, lineone)
//			}
//			//fmt.Println("hanghsu:", linerst, val.dispatch_id)
//			sql = `update  FTP_DISPATCH_PRO t set t.cur_rows = :1 where t.dispatch_id = :2 and t.dispatch_date = to_date(:3,'YYYY-MM-DD')`

//			for _, li := range linerst {
//				asofdate := li.as_of_date[0:10]
//				err := dbobj.Default.Exec(sql, li.cnt, val.dispatch_id, asofdate)
//				if err != nil {
//					fmt.Println(err)
//					logs.Error(err)
//				}
//			}
//		}
//		time.Sleep(time.Second * 5)
//	}
//}

//func init() {
//	ftprst = new(ftpCalcResult)
//	//go insertData()
//	//	go updateResultRows()
//}
