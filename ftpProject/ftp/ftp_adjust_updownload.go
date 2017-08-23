package ftp

import (
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"
	"ftpProject/utils/cacheutil"

	"ftpProject/utils"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

const startRow = 10
const readCellnum = 14

var checkform = [14]string{"adj_id", "org_unit_id", "iso_currency_cd", "adj_dyn_dim", "term_str", "term_end", "last_date", "adj_bp", "eff_str_date", "eff_end_date", "buz_str_date", "buz_end_date", "domain_id", "memo"}

type FtpAdjUpload struct {
	ReturnMsg
	RouteControl
}

func (this *FtpAdjUpload) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败,请严格按照模板上传"
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
		this.ErrorMsg = "上传失败,请严格按照模板上传"
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
	//fmt.Fprintln(w, "upload ok!")
	//解析excel

	//excelFileName := "test.xlsx"

	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败,请严格按照模板上传"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	sql := FTP_ADJPOLICYDOWN_POST1
	tx, _ := dbobj.Default.Begin()
	//
	var orgdata []OrgInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		orgdata = val.OrgData
	}
	//
	for _, sheet := range xlFile.Sheets {
		for j, row := range sheet.Rows {
			//前面为示例数据，不做输入
			if j == 1 {
				var srows = make([]string, 14, 14)
				for k, cell := range row.Cells {
					s, _ := cell.String()
					s = strings.TrimSpace(s)
					if k == 0 && s == "" {
						break
					}
					srows[k] = s
					//最后一列数据
					if k == readCellnum-1 {
						break
					}
				}
				for i, val := range srows {
					if val != checkform[i] {
						this.ErrorCode = "0"
						this.ErrorMsg = "导入失败,表头有错误,请严格按照模板格式导入"
						this.WriteJson(w, this.ReturnMsg)
						return
					}
				}
			}
			if j < startRow {
				continue
			}

			var srows = make([]string, 14, 14)
			for k, cell := range row.Cells {
				s, _ := cell.String()
				s = strings.TrimSpace(s)
				if k == 0 && s == "" {
					break
				}
				srows[k] = s
				//最后一列数据
				if k == readCellnum-1 {
					break
				}
			}

			if srows[0] != "" {
				//fmt.Println("j,srow", j, srows)
				//
				b := false
				for _, val := range orgdata {
					if val.Org_unit_id == srows[1] {
						srows[1] = val.Uuid
						b = true
						break
					}
				}
				if !b {
					logs.Error("导入数据出错，未找到机构信息，请确认该用户有此机构权限，错误表格行数为:", j+1)
					tx.Rollback()
					jj := strconv.Itoa(j + 1)
					this.ErrorCode = "0"
					this.ErrorMsg = "导入失败第" + jj + "行" + ",未找到匹配机构信息"
					this.WriteJson(w, this.ReturnMsg)
					return
				}
				//
				_, err := tx.Exec(sql,
					srows[0],
					srows[1],
					srows[2],
					srows[3],
					srows[4],
					srows[5],
					srows[6],
					srows[7],
					srows[8],
					srows[9],
					srows[10],
					srows[11],
					srows[12],
					srows[13])
				if err != nil {
					tx.Rollback()
					logs.Error(err)
					logs.Error("导入数据出错，错误表格行数为:", j+1)
					jj := strconv.Itoa(j + 1)
					this.ErrorCode = "0"
					this.ErrorMsg = "导入失败第" + jj + "行"
					this.WriteJson(w, this.ReturnMsg)
					return
				}
			}
		}
	}
	tx.Commit()
	//	//文件关闭
	//	f.Close()
	//	//删除临时文件
	//	err = os.Remove(excelFileName)
	//	if err != nil {
	//		logs.Error(err)
	//		this.ErrorCode = "1"
	//		this.ErrorMsg = "上传成功，删除临时文件失败，请手动删除：" + excelFileName
	//		this.WriteJson(w, this.ReturnMsg)
	//	}
	//成功的日志
	opcontent := "导入政策性调节项"
	this.InsertLogToDB(policyadjimport, opcontent, myapp)
	this.ErrorCode = "1"
	this.ErrorMsg = "上传成功"
	this.WriteJson(w, this.ReturnMsg)

}

//type FtpAdjDownload struct {
//	ReturnMsg
//	RouteControl
//}

type FtpAdjDownload struct {
	Uuid            string
	Adj_id          string
	Adj_desc        string
	Org_unit_id     string
	Org_unit_desc   string
	Iso_currency_cd string
	Adj_dyn_dim     string
	Dyn_name        string
	Term_str        string
	Term_end        string
	Last_date       string
	Adj_bp          string
	Eff_str_date    string
	Eff_end_date    string
	Buz_str_date    string
	Buz_end_date    string
	Domain_id       string
	Domain_name     string
	Memo            string
	Cnt             int
}
type FtpAdjDownloadCtl struct {
	ReturnMsg
	RouteControl
}
type FtpAdjDownloadB struct {
	Uuid            []byte
	Adj_id          []byte
	Adj_desc        []byte
	Org_unit_id     []byte
	Org_unit_desc   []byte
	Iso_currency_cd []byte
	Adj_dyn_dim     []byte
	Dyn_name        []byte
	Term_str        []byte
	Term_end        []byte
	Last_date       []byte
	Adj_bp          []byte
	Eff_str_date    []byte
	Eff_end_date    []byte
	Buz_str_date    []byte
	Buz_end_date    []byte
	Domain_id       []byte
	Domain_name     []byte
	Memo            []byte
	Cnt             int
}

func (this *FtpAdjDownloadCtl) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		sql = ""
		one FtpAdjDownload
		all []FtpAdjDownload
		tmp FtpAdjDownloadB
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}

	//offset, _ := strconv.Atoi(r.FormValue("offset"))
	//limit, _ := strconv.Atoi(r.FormValue("limit"))
	sAdjid := r.FormValue("keyword1")
	sother := r.FormValue("keyword2")
	//拼搜索sql
	cond := " 1=1"
	var sqlparams []string
	if sAdjid != "" || sother != "" {
		sAdjid = utils.HandleSqlKey(sAdjid)
		sother = utils.HandleSqlKey(sother)
		if sAdjid == "803" || sAdjid == "805" {
			cond = "t.adj_id=" + `'` + sAdjid + `'`
			sqlparams = append(sqlparams, cond)
			if sother != "" {
				cond = "to_number(t.term_str)<" + sother + " and to_number(t.term_end)>=" + sother
				sqlparams = append(sqlparams, cond)
			}
		} else {
			if sAdjid != "" {
				cond = "t.adj_id=" + `'` + sAdjid + `'`
				sqlparams = append(sqlparams, cond)
			}
			if sother != "" {
				cond = "(upper(t.adj_dyn_dim) like upper("
				cond += "'%" + sother + `%') escape '\' or upper(t.DYN_NAME) like upper(`
				cond += "'%" + sother + `%') escape '\')`
				sqlparams = append(sqlparams, cond)
			}
		}
		cond = strings.Join(sqlparams, " and ")
	}
	//fmt.Println("cond", cond)
	//这儿只能查看自己机构和下级机构的信息,取cache里面的data
	var orgdata []OrgInfo
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		orgdata = val.OrgData
	}
	//有两种方法，一种程序遍历，第二种传入sql语句  此处选择第二种orginpara
	var s []string
	for _, val := range orgdata {
		s1 := "'" + val.Uuid + "'"
		s = append(s, s1)
	}
	tmps := strings.Join(s, ",")
	ORGUINTIN := "(" + tmps + ")"
	//
	sql = FTP_ADJPOLICYDOWN_GET1
	sql = strings.Replace(sql, "SSQQLL", cond, -1)
	sql = strings.Replace(sql, "ORGINPARA", ORGUINTIN, -1)

	//fmt.Println("sql", sql)
	rows, err := dbobj.Default.Query(sql, doName)
	if err != nil {
		logs.Error(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(
			&tmp.Uuid,
			&tmp.Adj_id,
			&tmp.Adj_desc,
			&tmp.Org_unit_id,
			&tmp.Iso_currency_cd,
			&tmp.Adj_dyn_dim,
			&tmp.Dyn_name,
			&tmp.Term_str,
			&tmp.Term_end,
			&tmp.Last_date,
			&tmp.Adj_bp,
			&tmp.Eff_str_date,
			&tmp.Eff_end_date,
			&tmp.Buz_str_date,
			&tmp.Buz_end_date,
			&tmp.Domain_id,
			&tmp.Memo,
			&tmp.Cnt)
		if err != nil {
			logs.Error(err)
			return
		}
		//赋值
		one.Uuid = string(tmp.Uuid)
		one.Adj_id = string(tmp.Adj_id)
		one.Adj_desc = string(tmp.Adj_desc)
		//one.Org_unit_id = string(tmp.Org_unit_id)
		//one.Org_unit_desc = this.OrgName
		for _, val := range orgdata {
			if val.Uuid == string(tmp.Org_unit_id) {
				one.Org_unit_id = val.Org_unit_id
				one.Org_unit_desc = val.Org_unit_desc
			}
		}

		one.Iso_currency_cd = string(tmp.Iso_currency_cd)
		one.Adj_dyn_dim = string(tmp.Adj_dyn_dim)
		one.Dyn_name = string(tmp.Dyn_name)
		one.Term_str = string(tmp.Term_str)
		one.Term_end = string(tmp.Term_end)
		one.Last_date = string(tmp.Last_date)
		one.Adj_bp = string(tmp.Adj_bp)
		one.Eff_str_date = string(tmp.Eff_str_date)
		one.Eff_end_date = string(tmp.Eff_end_date)
		one.Buz_str_date = string(tmp.Buz_str_date)
		one.Buz_end_date = string(tmp.Buz_end_date)
		one.Domain_id = string(tmp.Domain_id)
		one.Domain_name = this.DomainName
		one.Memo = string(tmp.Memo)
		one.Cnt = tmp.Cnt
		//
		all = append(all, one)
	}
	var (
		file    *xlsx.File
		sheet   *xlsx.Sheet
		row     *xlsx.Row
		cell    *xlsx.Cell
		xlsxerr error
	)
	file, _ = xlsx.OpenFile("./updownload/TmpPolicyData.xlsx")
	sheet = file.Sheet["Sheet1"]
	if all != nil {
		for _, val := range all {
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.Value = val.Adj_id
			//			cell = row.AddCell()
			//			cell.Value = val.Adj_desc
			cell = row.AddCell()
			cell.Value = val.Org_unit_id
			//			cell = row.AddCell()
			//			cell.Value = val.Org_unit_desc
			cell = row.AddCell()
			cell.Value = val.Iso_currency_cd
			cell = row.AddCell()
			cell.Value = val.Adj_dyn_dim
			//			cell = row.AddCell()
			//			cell.Value = val.Dyn_name
			cell = row.AddCell()
			cell.Value = val.Term_str
			cell = row.AddCell()
			cell.Value = val.Term_end
			cell = row.AddCell()
			cell.Value = val.Last_date
			cell = row.AddCell()
			cell.Value = val.Adj_bp
			cell = row.AddCell()
			cell.Value = val.Eff_str_date
			cell = row.AddCell()
			cell.Value = val.Eff_end_date
			cell = row.AddCell()
			cell.Value = val.Buz_str_date
			cell = row.AddCell()
			cell.Value = val.Buz_end_date
			cell = row.AddCell()
			cell.Value = val.Domain_id
			cell = row.AddCell()
			cell.Value = val.Memo
			//			cell = row.AddCell()
			//			cell.Value = val.Domain_name
		}
	}
	ts := time.Now().Unix()
	tss := strconv.FormatInt(ts, 10)
	fileSavename := "./updownload/policyexport/" + tss + "ExpPolicyData.xlsx"
	xlsxerr = file.Save(fileSavename)
	if xlsxerr != nil {
		fmt.Printf(xlsxerr.Error())
		logs.Error(xlsxerr.Error())
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+tss+"ExpPolicyData.xlsx")
	http.ServeFile(w, r, fileSavename)

	//成功的日志
	opcontent := "导出政策性调节项"
	this.InsertLogToDB(policyadjexport, opcontent, myapp)

	this.ErrorCode = "1"
	this.ErrorMsg = "文件生成成功，准备导出"
	this.WriteJson(w, this.ReturnMsg)
}

//func (this *FtpAdjDownload) Post() {
//	//	excelFileName := "test.xlsx"
//	//	fmt.Println("filename:", excelFileName)
//	//	xlFile, err := xlsx.OpenFile(excelFileName)
//	//	if err != nil {
//	//		fmt.Println("open err", err)
//	//	}
//	//	io.WriteString(w, xlFile)
//}
