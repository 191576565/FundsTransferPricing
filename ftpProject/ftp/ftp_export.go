package ftp

import (
	"database/sql"
	"errors"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

var dstartrow = 1

//曲线定义表下载
func Dcurvedefine(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_CURVE_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}

	for rows.Next() {
		row := sheet.AddRow()
		myslice := make([]string, 7, 10)
		var tmp []byte
		err := rows.Scan(&myslice[0], &myslice[1], &myslice[2], &myslice[3], &myslice[4], &myslice[5], &tmp)
		if err != nil {
			logs.Error(err)
			return err
		}
		myslice[6] = string(tmp)
		//fmt.Println(myslice)
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Icurvedefine(sheet *xlsx.Sheet, tx *sql.Tx) error {

	sql := FTP_CURVE_IMPORT
	myslice := make([]string, 7, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 7 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s

		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3], myslice[4], myslice[5], myslice[6])
		if err != nil {
			logs.Error(err)
			errmsg := "插入曲线定义表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//曲线期限点下载
func Dcurvestruct(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_CURVESTRUCT_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}

	for rows.Next() {
		row := sheet.AddRow()
		myslice := make([]string, 4, 7)
		err := rows.Scan(&myslice[0], &myslice[1], &myslice[2], &myslice[3])
		if err != nil {
			logs.Error(err)
			return err
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Icurvestruct(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_CURVESTRUCT_IMPORT
	myslice := make([]string, 4, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 4 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s

		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3])
		if err != nil {
			logs.Error(err)
			errmsg := "插入曲线期限表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//
//曲线值下载
func Dcurveinfo(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_CURVEINFO_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}

	for rows.Next() {
		var yieldtmp []byte
		row := sheet.AddRow()
		myslice := make([]string, 4, 7)
		err := rows.Scan(&myslice[0], &myslice[1], &myslice[2], &yieldtmp)
		if err != nil {
			logs.Error(err)
			return err
		}
		myslice[3] = string(yieldtmp)
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Icurveinfo(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_CURVEINFO_IMPORT
	myslice := make([]string, 4, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 4 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s

		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3])
		if err != nil {
			logs.Error(err)
			errmsg := "插入曲线值表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//业务单元信息下载
func Dbusizinfo(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_BUSIZ_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}

	for rows.Next() {

		row := sheet.AddRow()
		myslice := make([]string, 7, 10)
		err := rows.Scan(&myslice[0], &myslice[1], &myslice[2], &myslice[3], &myslice[4], &myslice[5], &myslice[6])
		if err != nil {
			logs.Error(err)
			return err
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Ibusizinfo(sheet *xlsx.Sheet, tx *sql.Tx) error {

	sql := FTP_BUSIZ_IMPORT
	myslice := make([]string, 7, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 7 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3], myslice[4], myslice[5], myslice[6])
		if err != nil {
			logs.Error(err)
			errmsg := "插入业务单元表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//业务单元方法
func Dbusizmethod(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_BMETHOD_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}

	for rows.Next() {

		row := sheet.AddRow()
		mybyte := make([][]byte, 9, 10)
		myslice := make([]string, 9, 10)
		err := rows.Scan(&mybyte[0], &mybyte[1], &mybyte[2], &mybyte[3], &mybyte[4], &mybyte[5], &mybyte[6], &mybyte[7], &mybyte[8])
		if err != nil {
			logs.Error(err)
			return err
		}
		for i, val := range mybyte {
			myslice[i] = string(val)
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Ibusizmethod(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_BMETHOD_IMPORT
	myslice := make([]string, 9, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 9 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3], myslice[4], myslice[5], myslice[6], myslice[7], myslice[8])
		if err != nil {
			logs.Error(err)
			errmsg := "插入业务单元与定价方法表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//业务与调整项
func Dbusizadj(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_BADJ_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}
	for rows.Next() {
		row := sheet.AddRow()
		myslice := make([]string, 4, 10)
		err := rows.Scan(&myslice[0], &myslice[1], &myslice[2], &myslice[3])
		if err != nil {
			logs.Error(err)
			return err
		}

		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Ibusizadj(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_BADJ_IMPORT
	myslice := make([]string, 4, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 4 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3])
		if err != nil {
			logs.Error(err)
			errmsg := "插入业务单元与调节项表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//业务单元沉淀率法期限结构
func Dbusizcdlstr(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_BCD_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}
	for rows.Next() {
		row := sheet.AddRow()
		mybyte := make([][]byte, 6, 10)
		myslice := make([]string, 6, 10)
		err := rows.Scan(&mybyte[0], &mybyte[1], &mybyte[2], &mybyte[3], &mybyte[4], &mybyte[5])
		if err != nil {
			logs.Error(err)
			return err
		}
		for i, val := range mybyte {
			myslice[i] = string(val)
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Ibusizcdlstr(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_BCD_IMPORT
	myslice := make([]string, 6, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 6 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3], myslice[4], myslice[5])
		if err != nil {
			logs.Error(err)
			errmsg := "插入沉淀期限结构表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//准备金
func Dadjreverse(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_ADJREVERSE_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}
	for rows.Next() {
		row := sheet.AddRow()
		mybyte := make([][]byte, 7, 10)
		myslice := make([]string, 7, 10)
		err := rows.Scan(&mybyte[0], &mybyte[1], &mybyte[2], &mybyte[3], &mybyte[4], &mybyte[5], &mybyte[6])
		if err != nil {
			logs.Error(err)
			return err
		}
		for i, val := range mybyte {
			myslice[i] = string(val)
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Iadjreverse(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_ADJREVERSE_IMPORT
	myslice := make([]string, 7, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 7 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3], myslice[4], myslice[5], myslice[6])
		if err != nil {
			logs.Error(err)
			errmsg := "插入准本金表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//司库利润还原
func Dadjrestore(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_ADJRESTORE_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}
	for rows.Next() {
		row := sheet.AddRow()
		mybyte := make([][]byte, 4, 10)
		myslice := make([]string, 4, 10)
		err := rows.Scan(&mybyte[0], &mybyte[1], &mybyte[2], &mybyte[3])
		if err != nil {
			logs.Error(err)
			return err
		}
		for i, val := range mybyte {
			myslice[i] = string(val)
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Iadjrestore(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_ADJRESTORE_IMPORT
	myslice := make([]string, 4, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 4 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3])
		if err != nil {
			logs.Error(err)
			errmsg := "插入司库利润还原表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//
//期限流动性溢价
func Dadjliqu(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_TERMLIQU_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}
	for rows.Next() {
		row := sheet.AddRow()
		mybyte := make([][]byte, 5, 10)
		myslice := make([]string, 5, 10)
		err := rows.Scan(&mybyte[0], &mybyte[1], &mybyte[2], &mybyte[3], &mybyte[4])
		if err != nil {
			logs.Error(err)
			return err
		}
		for i, val := range mybyte {
			myslice[i] = string(val)
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Iadjliqu(sheet *xlsx.Sheet, tx *sql.Tx) error {
	sql := FTP_TERMLIQU_IMPORT
	myslice := make([]string, 5, 10)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 5 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3], myslice[4])
		if err != nil {
			logs.Error(err)
			errmsg := "插入期限流动性溢价表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

//政策性调节项
func Dadjpolicy(sheet *xlsx.Sheet, domain string) error {

	sql := FTP_ADJPLOCY_EXPORT
	rows, err := dbobj.Default.Query(sql, domain)
	if err != nil {
		logs.Error(err)
		return err
	}
	for rows.Next() {
		row := sheet.AddRow()
		mybyte := make([][]byte, 15, 20)
		myslice := make([]string, 15, 20)
		err := rows.Scan(&mybyte[0], &mybyte[1], &mybyte[2], &mybyte[3], &mybyte[4], &mybyte[5], &mybyte[6], &mybyte[7], &mybyte[8], &mybyte[9], &mybyte[10], &mybyte[11], &mybyte[12], &mybyte[13], &mybyte[14])
		if err != nil {
			logs.Error(err)
			return err
		}
		for i, val := range mybyte {
			myslice[i] = string(val)
		}
		row.WriteSlice(&myslice, -1)
	}
	return nil
}
func Iadjpolicy(sheet *xlsx.Sheet, tx *sql.Tx) error {

	sql := FTP_ADJPLOCY_IMPORT
	myslice := make([]string, 15, 20)
	for i, row := range sheet.Rows {
		if i < dstartrow {
			continue
		}
		for j, cell := range row.Cells {
			if j >= 15 {
				break
			}
			s, _ := cell.String()
			s = strings.TrimSpace(s)
			myslice[j] = s
		}
		//fmt.Println(myslice)
		_, err := tx.Exec(sql, myslice[0], myslice[1], myslice[2], myslice[3], myslice[4], myslice[5], myslice[6], myslice[7], myslice[8], myslice[9], myslice[10], myslice[11], myslice[12], myslice[13], myslice[14])
		if err != nil {
			logs.Error(err)
			errmsg := "插入政策性调节项表失败,第" + strconv.Itoa(i+1) + "行数据错误"
			return errors.New(errmsg)
		}
	}
	return nil
}

type FtpFormBackup struct {
	ReturnMsg
	RouteControl
}

func (this *FtpFormBackup) Get() {

	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	dcurveflag := r.FormValue("DownCurve")
	dbusizflag := r.FormValue("DownBusiz")
	dadjflag := r.FormValue("DownAdj")

	var (
		file  *xlsx.File
		sheet *xlsx.Sheet
		err   error
	)
	//file = xlsx.NewFile()
	file, _ = xlsx.OpenFile("./updownload/MyBackupFile.xlsx")
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//下载曲线
	if dcurveflag == "1" {
		//曲线
		sheet, _ = file.Sheet["曲线信息"]
		err = Dcurvedefine(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		//
		//期限结构
		sheet, _ = file.Sheet["期限结构"]
		err = Dcurvestruct(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		//
		//点值
		sheet, _ = file.Sheet["曲线点值"]
		err = Dcurveinfo(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	if dbusizflag == "1" {
		//Dbusizinfo
		sheet, _ = file.Sheet["业务单元"]
		err = Dbusizinfo(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		//Dbusizmethod
		sheet, _ = file.Sheet["业务单元与定价方法"]
		err = Dbusizmethod(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}

		sheet, _ = file.Sheet["沉淀期限结构"]
		err = Dbusizcdlstr(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		//Dbusizadj
		sheet, _ = file.Sheet["业务单元与调节项"]
		err = Dbusizadj(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	if dadjflag == "1" {
		sheet, _ = file.Sheet["准备金"]
		err = Dadjreverse(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		//Dadjrestore
		sheet, _ = file.Sheet["司库利润还原"]
		err = Dadjrestore(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		//Dadjliqu
		sheet, _ = file.Sheet["期限流动性溢价"]
		err = Dadjliqu(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		//Dadjpolicy
		sheet, _ = file.Sheet["政策性调整项"]
		err = Dadjpolicy(sheet, doName)
		if err != nil {
			this.ErrorCode = "0"
			this.ErrorMsg = err.Error()
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	username := this.Userid
	if username == "" {
		logs.Error("session中用户名为空")
		return
	}
	msavefile := "./updownload/ftpformbackup/" + username + "MyBackupFile.xlsx"
	mname := username + "业务配置方案.xlsx"
	err = file.Save(msavefile)
	if err != nil {
		fmt.Printf(err.Error())
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+mname)
	http.ServeFile(w, r, msavefile)
	opcontent := "导出业务方案"
	this.InsertLogToDB(bexport, opcontent, myapp)
}

//导入tx *sql.Tx
func (this *FtpFormBackup) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("uploadFile")
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败"
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
		this.ErrorMsg = "上传失败"
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
		this.ErrorMsg = "上传失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}

	for _, val := range sheetnamearry {
		if _, ok := xlFile.Sheet[val]; !ok {
			logs.Error(err)
			logs.Error("缺少" + val + "sheet页面")
			this.ErrorCode = "0"
			this.ErrorMsg = "缺少" + val + "sheet页面"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	tx, _ := dbobj.Default.Begin()
	//校验在跑批没
	var allrow string = "0"
	csql := FTP_CHECTPATCH_RUN
	row := tx.QueryRow(csql, doName)
	err = row.Scan(&allrow)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = "有批次正在运行，无法导入"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	if allrow != "0" {
		this.ErrorCode = "0"
		this.ErrorMsg = "有批次正在运行，无法导入"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	//先删除
	for _, val := range delearrytmp {
		sql := `delete from ` + val + ` where domain_id=` + "'" + doName + "'"
		if val == "MAS_CURVE_INFO" {
			sql = `delete from ` + val + ` where curve_uuid like ` + "'" + doName + "%'"
		}
		_, err := tx.Exec(sql)
		if err != nil {
			tx.Rollback()
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "删除数据出错,请联系管理员"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	//曲线

	sheet := xlFile.Sheet["曲线信息"]
	err = Icurvedefine(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}

	//期限结构
	//sheet, _ = file.AddSheet("期限结构")
	sheet = xlFile.Sheet["期限结构"]
	err = Icurvestruct(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}

	//
	//点值
	//sheet, _ = file.AddSheet("曲线点值")
	sheet = xlFile.Sheet["曲线点值"]
	err = Icurveinfo(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	sheet = xlFile.Sheet["业务单元"]
	err = Ibusizinfo(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	//Dbusizmethod
	//sheet, _ = file.AddSheet("业务单元与定价方法")
	sheet = xlFile.Sheet["业务单元与定价方法"]
	err = Ibusizmethod(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	//Dbusizadj
	//sheet, _ = file.AddSheet("业务单元与调节项")
	sheet = xlFile.Sheet["业务单元与调节项"]
	err = Ibusizadj(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	//sheet, _ = file.AddSheet("期限流动性溢价期限结构")
	sheet = xlFile.Sheet["沉淀期限结构"]
	err = Ibusizcdlstr(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	sheet = xlFile.Sheet["准备金"]
	//sheet, _ = file.AddSheet("准备金")
	err = Iadjreverse(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	//Dadjrestore
	//sheet, _ = file.AddSheet("司库利润还原")
	sheet = xlFile.Sheet["司库利润还原"]

	err = Iadjrestore(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	//Dadjliqu
	//sheet, _ = file.AddSheet("期限流动性溢价")
	sheet = xlFile.Sheet["期限流动性溢价"]

	err = Iadjliqu(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	//Dadjpolicy
	sheet = xlFile.Sheet["政策性调整项"]

	//sheet, _ = file.AddSheet("政策性调整项")
	err = Iadjpolicy(sheet, tx)
	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		tx.Rollback()
		return
	}
	//
	//	f.Close()
	//	err = os.Remove(excelFileName)
	//	if err != nil {
	//		logs.Error(err)
	//		this.ErrorCode = "1"
	//		this.ErrorMsg = "上传成功，删除临时文件失败，请手动删除：" + excelFileName
	//		this.WriteJson(w, this.ReturnMsg)
	//	}
	//
	tx.Commit()
	//
	//成功的日志
	opcontent := "全量导入"
	this.InsertLogToDB(all_input, opcontent, myapp)
	//
	this.ErrorCode = "1"
	this.ErrorMsg = "导入成功"
	this.WriteJson(w, this.ReturnMsg)

}

var delearrytmp = [11]string{
	"FTP_ADJUST_CAPITAL_RESERVES",
	"FTP_ADJUST_FTP_RESTORE",
	"FTP_ADJUST_TERM_LIQUIDITY",
	"FTP_ADJUST_POLICY",
	"FTP_ADJUST_REL",
	"FTP_BUSIZ_REDEMPTION_CURVE",
	"FTP_BUSIZ_METHOD_RELATION",
	"FTP_BUSIZ_INFO",
	"MAS_CURVE_INFO",
	"MAS_CURVE_INFO_STRUCT_NODE",
	"MAS_CURVE_DEFINE",
}
var sheetnamearry = [11]string{
	"曲线信息",
	"期限结构",
	"曲线点值",
	"业务单元",
	"业务单元与定价方法",
	"业务单元与调节项",
	"沉淀期限结构",
	"准备金",
	"司库利润还原",
	"期限流动性溢价",
	"政策性调整项",
}
