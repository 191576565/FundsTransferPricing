package ftp

import (
	"ftpProject/dbobj"
	"ftpProject/logs"

	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var StNode = [24]string{
	"曲线编号",
	"日期",
	"1D",
	"7D",
	"14D",
	"1M",
	"2M",
	"3M",
	"4M",
	"5M",
	"6M",
	"7M",
	"8M",
	"9M",
	"10M",
	"11M",
	"1Y",
	"2Y",
	"3Y",
	"5Y",
	"10Y",
	"15Y",
	"20Y",
	"30Y",
}

type CurveInfoInput struct {
	ReturnMsg
	RouteControl
}

func (this *CurveInfoInput) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败,请按照模板格式上传"
		this.WriteJson(w, this.ReturnMsg)

		return
	}
	defer file.Close()
	filetype := path.Ext(handler.Filename)
	if filetype != ".xlsx" {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "文件格式不对，请用模板上传"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	ts := time.Now().Unix()
	tss := strconv.FormatInt(ts, 10)

	excelFileName := "./updownload/" + tss + handler.Filename
	f, err := os.OpenFile(excelFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败,请按照模板格式上传"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	io.Copy(f, file)
	defer func() {
		f.Close()
		err := os.Remove(excelFileName)
		if err != nil {
			logs.Error("上传成功，删除临时文件失败，请手动删除：" + excelFileName)
		}
	}()
	//excelFileName := "1234.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败,请按照模板格式上传"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	tx, _ := dbobj.Default.Begin()
	sql := FTP_CURVEINFO_INPUT
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	for _, sheet := range xlFile.Sheets {
		for j, row := range sheet.Rows {
			//前面为示例数据，不做输入
			if j == 0 { //表头校验
				srows := make([]string, 24, 30)
				for k, cell := range row.Cells {
					s, _ := cell.String()
					s = strings.TrimSpace(s)
					srows[k] = s
					if k > 23 {
						break
					}
				}
				for i, val := range srows {
					if val != StNode[i] {
						this.ErrorCode = "0"
						this.ErrorMsg = "导入失败,表头有错误,请严格按照模板格式导入"
						this.WriteJson(w, this.ReturnMsg)
						return
					}
				}
			}
			//
			if j < 4 {
				continue
			}
			//			var srows = make([]string, 24, 30)
			var busiz_id = ""
			var date = ""
			for k, cell := range row.Cells {
				s, _ := cell.String()
				s = strings.TrimSpace(s)
				if k == 0 {
					busiz_id = s
				}
				if k == 1 {
					date = s
				}
				if k > 23 {
					break
				}
				if k > 1 {
					if s != "" {
						uuid := doName + "_" + busiz_id + "_" + StNode[k]
						//fmt.Println("uuid asofdate  yeild:", uuid, date, s)
						_, err := tx.Exec(sql, uuid, date, s)
						if err != nil {
							logs.Error(err)
							jj := strconv.Itoa(j + 1)
							kk := strconv.Itoa(k + 1)
							tx.Rollback()
							this.ErrorCode = "0"
							this.ErrorMsg = "导入失败第" + jj + "行," + kk + "列"
							this.WriteJson(w, this.ReturnMsg)
							return
						}
					}
				}
			}
		}
	}

	//成功的日志
	opcontent := "曲线值增量导入"
	this.InsertLogToDB(curveinfoinput, opcontent, myapp)
	//
	tx.Commit()
	this.ErrorCode = "1"
	this.ErrorMsg = "导入成功"
	this.WriteJson(w, this.ReturnMsg)
}
