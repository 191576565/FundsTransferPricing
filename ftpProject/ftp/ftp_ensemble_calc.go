package ftp

import (
	"encoding/json"
	"errors"
	"fmt"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"ftpProject/utils"
	"io"
	"net/rpc"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/config"

	"github.com/tealeg/xlsx"
)

type FtpEnsembleCalaPage struct {
	RouteControl
}

func (this *FtpEnsembleCalaPage) Get() {
	this.TplName = "mas/ftp/ftp_ensemble_calc.tpl"
}

type FtpEnsembleCalc struct {
	ReturnMsg
	RouteControl
}
type FtpEnsembleCalcAuto struct {
	ReturnMsg
	RouteControl
}

const (
	CELLA = iota
	CELLB
	CELLC
	CELLD
	CELLE
	CELLF
	CELLG
	CELLH
	CELLI
	CELLJ
	CELLK
)
const ReadNum = CELLK
const SHEETNUM = 6

var SHEETR = [SHEETNUM]int{6, 9, 12, 12, 8, 12}
var ResultName = [10]string{
	"贷款小计",
	"直贴合计",
	"转贴现合计",
	"交易性金融资产",
	"可供出售金融资产",
	"其他资金资产合计",
	"存款合计",
	"资金负债合计",
	"所有资产合计",
	"所有负债合计",
}
var ResultName1 = [5]string{
	"存贷利差切割分析",
	"存贷利差",
	"贷款利差",
	"存款利差",
	"存贷错配利差",
}
var ResultName2 = [5]string{
	"生息资产和付息负债利差切割分析",
	"全行利差",
	"资产利差",
	"负债利差",
	"全行错配利差",
}
var sheetcalc = [6]string{
	"Sheet1",
	"Sheet2",
	"Sheet3",
	"Sheet4",
	"Sheet5",
	"Sheet6",
}

type CommonPara struct {
	OrgTerm       string  //原始期限
	OrgTermMult   string  //原始期限单位
	StockBalance  float64 //存量余额
	Ratio         float64 //占比
	WeightRate    float64 //加权利率
	Income        float64 //收入
	FtpValue      float64 //FTP价格
	ProfitRe      float64 //司库还原
	AftpValue     float64 //调整后FTP价格
	FtpDiffer     float64 //FTP利差
	BeforeFtpCost float64 //调整前FTP成本
	AfterFtpCost  float64 //调整后FTP成本
	FtpProfit     float64 //FTP利润

}
type Acalc struct {
	Id     string
	Bid    string
	Desc   string
	sum    float64
	Params []CommonPara
}
type CommonParass struct {
	OrgTerm       string //原始期限
	OrgTermMult   string //原始期限单位
	StockBalance  string //存量余额
	Ratio         string //占比
	WeightRate    string //加权利率
	Income        string //收入
	FtpValue      string //FTP价格
	ProfitRe      string //司库还原
	AftpValue     string //调整后FTP价格
	FtpDiffer     string //FTP利差
	BeforeFtpCost string //调整前FTP成本
	AfterFtpCost  string //调整后FTP成本
	FtpProfit     string //FTP利润
}
type Acalcss struct {
	Id     string
	Bid    string
	Desc   string
	sum    string
	Params []CommonParass
}
type TotalResult struct {
	Zcid          string  //资产项
	StockBalance  float64 //存量余额
	ExternTa      float64 //外部利率
	ExternTax     float64 //外部利息
	FtpValue      float64 //FTP价格
	AftpValue     float64 //调整后FTP价格
	FtpDiffer     float64 //FTP利差
	BeforeFtpCost float64 //调整前FTP成本
	AfterFtpCost  float64 //调整后FTP成本
	FtpProfit     float64 //FTP利润
}
type TotalResultss struct {
	Zcid          string //资产项
	StockBalance  string //存量余额
	ExternTa      string //外部利率
	ExternTax     string //外部利息
	FtpValue      string //FTP价格
	AftpValue     string //调整后FTP价格
	FtpDiffer     string //FTP利差
	BeforeFtpCost string //调整前FTP成本
	AfterFtpCost  string //调整后FTP成本
	FtpProfit     string //FTP利润
}
type XlsxJson struct {
	ReturnMsg
	DfileName string
	Single    *[]Acalcss
	Result    *[]TotalResultss
}

func (this *FtpEnsembleCalc) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	var (
		xlFile *xlsx.File
		err    error
		xrev   []Acalc
		tmpA   Acalc
		tmpP   CommonPara
		iarry  int  = -1 //读取excel成数组的下标
		flag   bool      //true:新建一个接收excel的结构体
	)
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	//上传start
	r.ParseMultipartForm(32 << 20)
	Iso_currency_cd := r.FormValue("IsoCurrencyCd")

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
	xf, err := os.OpenFile(excelFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	io.Copy(xf, file)
	defer func() {
		xf.Close()
		err := os.Remove(excelFileName)
		if err != nil {
			logs.Error("上传成功，删除临时文件失败，请手动删除：" + excelFileName)
		}
	}()
	//上传end
	xlFile, err = xlsx.OpenFile(excelFileName)
	if err != nil {
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = "上传失败"
		this.WriteJson(w, this.ReturnMsg)
		fmt.Println("open err", err)
		return
	}
	for _, val := range sheetcalc {
		if _, ok := xlFile.Sheet[val]; !ok {
			logs.Error(err)
			this.ErrorCode = "0"
			this.ErrorMsg = "缺少" + val + "的sheet页面,请严格按照模板格式上传"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
	}
	for i, sheet := range xlFile.Sheets {
		//总共有多少个Sheet
		if i > SHEETNUM {
			break
		}
		flag = true
		//--------遍历每行Excel-----------------------
		for j, row := range sheet.Rows {
			//第一行为表头，不读
			if j == 0 {
				continue
			}
			//Sheet3单独处理
			if i == 2 && j == 10 {
				continue
			}
			//每个Sheet读多少行做限定
			if j > SHEETR[i] {
				break
			}
			var rowdata []string
			//--------遍历每行每个单元格-----------------------
			for k, cell := range row.Cells {
				//每行读多少列做限定
				if k > ReadNum {
					break
				}
				s, _ := cell.String()
				s = strings.TrimSpace(s)

				rowdata = append(rowdata, s)
			}
			//fmt.Println("rowdate:", rowdata)
			if rowdata == nil {
				break
			}
			//Sheet3特殊处理
			if i == 2 && j == 11 {
				flag = true
			}
			//Sheet3特殊处理
			if i == 2 && j == 12 {
				flag = true
			}
			//Sheet5特殊处理第3到9
			if i == 4 && j == 2 {
				flag = true
			}
			if flag {
				tmpA.Id = rowdata[CELLA]
				tmpA.Desc = rowdata[CELLB]
				tmpA.Bid = rowdata[CELLC]
				xrev = append(xrev, tmpA)
				iarry++
				flag = false
			}
			//如果FTP价格不为空
			tmpP.FtpValue = 0
			if rowdata[CELLK] != "" {
				f, _ := strconv.ParseFloat(rowdata[CELLK], 64)
				tmpP.FtpValue = f
			}
			tmpP.OrgTerm = rowdata[CELLE]
			tmpP.OrgTermMult = rowdata[CELLF]
			f, _ := strconv.ParseFloat(rowdata[CELLG], 64)
			tmpP.StockBalance = f
			f, _ = strconv.ParseFloat(rowdata[CELLI], 64)
			tmpP.WeightRate = f
			//合计存量余额
			xrev[iarry].sum += tmpP.StockBalance
			xrev[iarry].Params = append(xrev[iarry].Params, tmpP)
		}
	}
	//note:合计特殊处理
	if iarry >= 7 {
		sumts := xrev[6].sum + xrev[7].sum
		xrev[6].sum = sumts
		xrev[7].sum = sumts
	}
	//	fmt.Println("1111---------------------")
	//	for _, val := range xrev {
	//		fmt.Println("id", val.Id)
	//		fmt.Println("bid", val.Bid)
	//		fmt.Println("desc", val.Desc)
	//		fmt.Println("sum", val.sum)
	//		for _, val := range val.Params {
	//			fmt.Println("params", val)
	//		}
	//	}
	//	fmt.Println("1111---------------------")
	err = HandleCallRpc(&xrev, Iso_currency_cd, doName)
	//	fmt.Println("2222---------------------")
	//	for _, val := range xrev {
	//		fmt.Println("id", val.Id)
	//		fmt.Println("bid", val.Bid)
	//		fmt.Println("desc", val.Desc)
	//		fmt.Println("sum", val.sum)
	//		for _, val := range val.Params {
	//			fmt.Println("params", val)
	//		}
	//	}
	//	fmt.Println("2222---------------------")
	if err != nil {
		fmt.Println("err:", err)
		logs.Error(err)
		this.ErrorCode = "0"
		this.ErrorMsg = err.Error()
		this.WriteJson(w, this.ReturnMsg)
		fmt.Println("open err", err)
		return
	}
	//计算单条数据
	for i, val := range xrev {
		for j, _ := range val.Params {
			xrev[i].Params[j].Income = xrev[i].Params[j].StockBalance * xrev[i].Params[j].WeightRate / 100
			xrev[i].Params[j].Ratio = xrev[i].Params[j].StockBalance / xrev[i].sum
			xrev[i].Params[j].AftpValue = xrev[i].Params[j].FtpValue + xrev[i].Params[j].ProfitRe
			//note:资产和负债的ftp利差不一样
			if i <= 5 {
				xrev[i].Params[j].FtpDiffer = xrev[i].Params[j].WeightRate - xrev[i].Params[j].AftpValue

			} else {
				xrev[i].Params[j].FtpDiffer = xrev[i].Params[j].AftpValue - xrev[i].Params[j].WeightRate
			}
			xrev[i].Params[j].BeforeFtpCost = xrev[i].Params[j].StockBalance * xrev[i].Params[j].FtpValue / 100
			xrev[i].Params[j].AfterFtpCost = xrev[i].Params[j].StockBalance * xrev[i].Params[j].AftpValue / 100
			xrev[i].Params[j].FtpProfit = xrev[i].Params[j].StockBalance * xrev[i].Params[j].FtpDiffer / 100
		}
	}
	//打印接受值

	//	fmt.Println("加工好的--------------------")
	//	for _, val := range xrev {
	//		fmt.Println("id", val.Id)
	//		fmt.Println("bid", val.Bid)
	//		fmt.Println("desc", val.Desc)
	//		fmt.Println("sum", val.sum)
	//		for _, val := range val.Params {
	//			fmt.Println("params", val)
	//		}
	//	}
	//	fmt.Println("加工好的over--------------------")

	//封装合计
	var one TotalResult
	var all []TotalResult
	var (
		sumIncomeold    float64 = 0
		sumAftpvalold   float64 = 0
		sumAftpcostold  float64 = 0
		sumAftpdiffold  float64 = 0
		sumBefcostold   float64 = 0
		sumAftcostold   float64 = 0
		sumFtpProfitold float64 = 0
	)
	for i, val := range xrev {
		var sumIncome float64 = 0
		var sumAftpval float64 = 0
		var sumAftpcost float64 = 0
		var sumAftpdiff float64 = 0
		var sumBefcost float64 = 0
		var sumAftcost float64 = 0
		var sumFtpProfit float64 = 0
		one.Zcid = val.Id
		one.StockBalance = val.sum
		for _, val := range val.Params {
			sumIncome += val.Income
			sumAftpval += val.AftpValue
			sumAftpcost += val.AfterFtpCost
			sumAftpdiff += val.FtpProfit
			sumBefcost += val.BeforeFtpCost
			sumAftcost += val.AfterFtpCost
			sumFtpProfit += val.FtpProfit
		}
		//note:特殊处理
		if i == 6 {
			sumIncomeold = sumIncome
			sumAftpvalold = sumAftpval
			sumAftpcostold = sumAftpcost
			sumAftpdiffold = sumAftpdiff
			sumBefcostold = sumBefcost
			sumAftcostold = sumAftcost
			sumFtpProfitold = sumFtpProfit
			continue
		}
		if i == 7 {
			sumIncome += sumIncomeold
			sumAftpval += sumAftpvalold
			sumAftpcost += sumAftpcostold
			sumAftpdiff += sumAftpdiffold
			sumBefcost += sumBefcostold
			sumAftcost += sumAftcostold
			sumFtpProfit += sumFtpProfitold
		}
		one.ExternTa = sumIncome / one.StockBalance * 100
		one.ExternTax = one.StockBalance * one.ExternTa / 100
		one.FtpValue = sumBefcost / one.StockBalance * 100
		one.AftpValue = sumAftpcost / one.StockBalance * 100
		one.FtpDiffer = sumAftpdiff / one.StockBalance * 100
		one.BeforeFtpCost = sumBefcost
		one.AfterFtpCost = sumAftcost
		//one.FtpProfit = one.ExternTax - one.AfterFtpCost
		one.FtpProfit = sumFtpProfit
		all = append(all, one)
	}
	//封装总计
	allzc := all[:6]
	allfz := all[6:8]
	//先清0数据
	one.StockBalance = 0
	one.ExternTax = 0
	one.BeforeFtpCost = 0
	one.AfterFtpCost = 0
	one.FtpProfit = 0
	for _, val := range allzc {
		one.StockBalance += val.StockBalance
		one.ExternTax += val.ExternTax
		one.BeforeFtpCost += val.BeforeFtpCost
		one.AfterFtpCost += val.AfterFtpCost
		one.FtpProfit += val.FtpProfit
	}
	one.Zcid = "10"
	one.ExternTa = one.ExternTax / one.StockBalance * 100
	one.FtpValue = one.BeforeFtpCost / one.StockBalance * 100
	one.AftpValue = one.AfterFtpCost / one.StockBalance * 100
	one.FtpDiffer = one.FtpProfit / one.StockBalance * 100
	all = append(all, one)
	//先清0数据
	one.StockBalance = 0
	one.ExternTax = 0
	one.BeforeFtpCost = 0
	one.AfterFtpCost = 0
	one.FtpProfit = 0
	for _, val := range allfz {
		one.StockBalance += val.StockBalance
		one.ExternTax += val.ExternTax
		one.BeforeFtpCost += val.BeforeFtpCost
		one.AfterFtpCost += val.AfterFtpCost
		one.FtpProfit += val.FtpProfit
	}
	one.Zcid = "11"
	one.ExternTa = one.ExternTax / one.StockBalance * 100
	one.FtpValue = one.BeforeFtpCost / one.StockBalance * 100
	one.AftpValue = one.AfterFtpCost / one.StockBalance * 100
	one.FtpDiffer = one.FtpProfit / one.StockBalance * 100
	all = append(all, one)
	//打印合计值
	//	for _, val := range all {
	//		fmt.Println("alresult:", val)
	//	}

	//将Float64 转成字符串
	var aass []Acalcss
	var css CommonParass
	for _, val := range xrev {
		var ass Acalcss
		ass.Id = val.Id
		ass.Desc = val.Desc
		ass.Bid = val.Bid
		for _, val := range val.Params {
			css.OrgTerm = val.OrgTerm
			css.OrgTermMult = val.OrgTermMult
			css.StockBalance = strconv.FormatFloat(val.StockBalance, 'f', 0, 64)
			css.Ratio = strconv.FormatFloat(val.Ratio*100, 'f', 2, 64) + "%"
			css.WeightRate = strconv.FormatFloat(val.WeightRate, 'f', 4, 64)
			css.Income = strconv.FormatFloat(val.Income, 'f', 0, 64)
			css.FtpValue = strconv.FormatFloat(val.FtpValue, 'f', 4, 64)
			css.ProfitRe = strconv.FormatFloat(val.ProfitRe, 'f', 4, 64)
			css.AftpValue = strconv.FormatFloat(val.AftpValue, 'f', 4, 64)
			css.FtpDiffer = strconv.FormatFloat(val.FtpDiffer, 'f', 4, 64)
			css.BeforeFtpCost = strconv.FormatFloat(val.BeforeFtpCost, 'f', 0, 64)
			css.AfterFtpCost = strconv.FormatFloat(val.AfterFtpCost, 'f', 0, 64)
			css.FtpProfit = strconv.FormatFloat(val.FtpProfit, 'f', 0, 64)
			ass.Params = append(ass.Params, css)
		}
		aass = append(aass, ass)
	}

	//zhuan
	var ones TotalResultss
	var alls []TotalResultss
	for _, val := range all {
		ones.Zcid = val.Zcid
		ones.StockBalance = strconv.FormatFloat(val.StockBalance, 'f', 0, 64)
		ones.ExternTa = strconv.FormatFloat(val.ExternTa, 'f', 4, 64)
		ones.ExternTax = strconv.FormatFloat(val.ExternTax, 'f', 0, 64)
		ones.FtpValue = strconv.FormatFloat(val.FtpValue, 'f', 4, 64)
		ones.AftpValue = strconv.FormatFloat(val.AftpValue, 'f', 4, 64)
		ones.FtpDiffer = strconv.FormatFloat(val.FtpDiffer, 'f', 4, 64)
		ones.BeforeFtpCost = strconv.FormatFloat(val.BeforeFtpCost, 'f', 0, 64)
		ones.AfterFtpCost = strconv.FormatFloat(val.AfterFtpCost, 'f', 0, 64)
		ones.FtpProfit = strconv.FormatFloat(val.FtpProfit, 'f', 0, 64)
		alls = append(alls, ones)
	}
	//
	//为下载excel做数据准备
	var cddiff []string
	var allali []string
	if len(all) >= 10 {
		//存贷利差
		s := strconv.FormatFloat((all[0].ExternTa - all[6].ExternTa), 'f', 4, 64)
		cddiff = append(cddiff, s)
		//贷款利差
		cddiff = append(cddiff, alls[0].FtpDiffer)
		//存款利差
		cddiff = append(cddiff, alls[6].FtpDiffer)
		//存贷错配利差
		ft := all[0].ExternTa - all[6].ExternTa - all[0].FtpDiffer - all[6].FtpDiffer
		s = strconv.FormatFloat(ft, 'f', 4, 64)
		cddiff = append(cddiff, s)

		ft = all[8].ExternTa - all[9].ExternTa
		s = strconv.FormatFloat(ft, 'f', 4, 64)
		allali = append(allali, s)
		//资产利差
		allali = append(allali, alls[8].FtpDiffer)
		//负债利差
		allali = append(allali, alls[9].FtpDiffer)
		//全行错配利差
		ft = all[8].ExternTa - all[9].ExternTa - all[8].FtpDiffer - all[9].FtpDiffer
		s = strconv.FormatFloat(ft, 'f', 4, 64)
		allali = append(allali, s)
	}

	//文件关闭
	//	xf.Close()
	//	//删除临时文件
	//	err = os.Remove(excelFileName)
	//	if err != nil {
	//		logs.Error(err)
	//		this.ErrorCode = "1"
	//		this.ErrorMsg = "上传成功，删除临时文件失败，请手动删除：" + excelFileName
	//		this.WriteJson(w, this.ReturnMsg)
	//	}
	//add 2016.11.1  为下载准备
	myfile, _ := xlsx.OpenFile("./updownload/TmpExportCalcResult.xlsx")
	sheet := myfile.Sheet["Sheet1"]
	for i, val := range alls {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = ResultName[i]
		cell = row.AddCell()
		cell.Value = val.StockBalance
		cell = row.AddCell()
		cell.Value = val.ExternTa
		cell = row.AddCell()
		cell.Value = val.ExternTax
		cell = row.AddCell()
		cell.Value = val.FtpValue
		cell = row.AddCell()
		cell.Value = val.AftpValue
		cell = row.AddCell()
		cell.Value = val.FtpDiffer
		cell = row.AddCell()
		cell.Value = val.BeforeFtpCost
		cell = row.AddCell()
		cell.Value = val.AfterFtpCost
		cell = row.AddCell()
		cell.Value = val.FtpProfit
	}
	row := sheet.AddRow()
	row = sheet.AddRow()
	cell := row.AddCell()
	cell.Value = ResultName1[0]
	for i, val := range cddiff {
		row = sheet.AddRow()
		cell := row.AddCell()
		cell.Value = ResultName1[i+1]
		cell = row.AddCell()
		cell.Value = val
	}
	row = sheet.AddRow()
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = ResultName2[0]
	for i, val := range allali {
		row = sheet.AddRow()
		cell := row.AddCell()
		cell.Value = ResultName2[i+1]
		cell = row.AddCell()
		cell.Value = val
	}
	ts = time.Now().Unix()
	tss = strconv.FormatInt(ts, 10)
	userid := this.Userid
	if userid == "" {
		logs.Error("session中域名为空")
	}
	fileSavename := "./updownload/calcresult/" + userid + tss + "calc.xlsx"

	xlsxerr := myfile.Save(fileSavename)
	if xlsxerr != nil {
		fmt.Printf(xlsxerr.Error())
		logs.Error(xlsxerr.Error())
	}
	//添加日志
	opcontent := "手工计算"
	this.InsertLogToDB(xlsxcalc, opcontent, myapp)
	//打包最终json
	var outjson XlsxJson
	outjson.ErrorCode = "1"
	outjson.ErrorMsg = "成功"
	outjson.DfileName = fileSavename[1:]
	outjson.Result = &alls
	outjson.Single = &aass
	ojs, err := json.Marshal(&outjson)
	if err != nil {
		fmt.Println("json打包失败")
	}
	w.Write(ojs)
}
func HandleCallRpc(xrev *[]Acalc, Iso_currency_cd string, domain_id string) error {
	//fmt.Println("calc start---------")
	//组装数据  准备调接口
	var (
		tmp   ACCTInfo
		args  []ACCTInfo
		reply []ACCTInfo
		sum   int = 0
	)
	for _, val := range *xrev {
		tmp.Busiz_id = val.Bid
		tmp.Iso_currency_cd = Iso_currency_cd
		tmp.Adjustable_type_cd = "0"
		tmp.Domain_id = domain_id
		for _, val := range val.Params {
			tmp.Org_term = val.OrgTerm
			tmp.Org_term_mult = val.OrgTermMult
			//tmp.Accrual_basis_cd = "0"

			//自己加工的数据
			num, _ := strconv.Atoi(tmp.Org_term)
			tmp.Origination_date = time.Now().Format("2006-01-02")
			switch val.OrgTermMult {
			case "D":
				tmp.Maturity_date, _ = utils.AddDays(tmp.Origination_date, num*(1))
			case "M":
				tmp.Maturity_date, _ = utils.AddMonths(tmp.Origination_date, num*(1))
			case "Y":
				tmp.Maturity_date, _ = utils.AddMonths(tmp.Origination_date, num*(12))
			default:
				tmp.Maturity_date = tmp.Origination_date
			}
			sum++
			args = append(args, tmp)
		}
	}
	//	for _, val := range args {
	//		fmt.Println("args:", val)
	//	}
	//add
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")

	red, err := config.NewConfig("ini", appConfigPath)
	if err != nil {
		logs.Error("cant not read ./conf/app.conf.please check this file.")
		return errors.New("读取计算引擎地址出错，请联系管理员检查计算引擎配置文件")
	}
	calcip := red.String("Calc.IP")
	conn, err := rpc.DialHTTP("tcp", calcip)

	if err != nil {
		logs.Error(err)
		return errors.New("与引擎服务建立连接失败,请联系管理员检查引擎服务配置和引擎是否启动")
	}

	err = conn.Call("RpcSrvOfFTP.FtpTrial", args, &reply)

	if err != nil {
		logs.Error(err)
		return errors.New("调用引擎服务失败,请联系管理员检查系统日志")
	}

	if reply == nil {
		logs.Error(err)
		return errors.New("服务器返回值为空,请检查参数是否错误")
	}
	//fmt.Println("sum len(replay):", sum, len(reply))
	if sum != len(reply) {
		fmt.Println("返回条数不对应")
		return errors.New("服务器返回值不对应")
	}
	//已返回ftp价格
	sum = 0
	for i, val := range *xrev {
		for j, _ := range val.Params {
			//赋值ftp价格
			//			fmt.Println("ftpv:", reply[sum].FtpRate)
			//			fmt.Println("zbj:", reply[sum].Adjtwo)
			//			fmt.Println("sk:", reply[sum].Adjthree)
			if (*xrev)[i].Params[j].FtpValue == 0 {
				f, _ := strconv.ParseFloat(reply[sum].FtpRate, 64)
				(*xrev)[i].Params[j].FtpValue = f
			}
			//赋值准备金
			if i > 5 {
				f, _ := strconv.ParseFloat(reply[sum].Adjtwo, 64)
				(*xrev)[i].Params[j].ProfitRe = f
			} else {
				f, _ := strconv.ParseFloat(reply[sum].Adjthree, 64)
				(*xrev)[i].Params[j].ProfitRe = f
			}
			sum++
		}
	}
	return nil
}

//自动计算
type tmpAcalc struct {
	item_id      string
	item_name_l1 string
	item_name_l2 string
	busiz_id     string
	org_par_bal  string
	ratio_bal    string
	cur_net_rate string
	accd_int     string
	adj_int      string
	ftp_rate_b   string
	ftp_rate_a   string
	ftp_margin_a string
	ftp_int_b    string
	ftp_int_a    string
	ftp_profit_a string
	//	domain_id       string
	//	iso_currency_cd string
}

func (this *FtpEnsembleCalcAuto) Get() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	iso_cd := r.FormValue("IsoCurrencyCd")
	var (
		sql = ""
	)

	doName := this.Domainid
	if doName == "" {
		logs.Error("seesion中域名为空")
		return
	}
	//首先调用存储过程
	param2 := time.Now().Format("2006-01-02")

	err := dbobj.Default.Exec("call PROC_FTP_ENSEMBLE_CALC(to_date(:1,'YYYY-MM-DD'), :2)", param2, doName)

	if err != nil {
		this.ErrorCode = "0"
		this.ErrorMsg = "调用存储过程失败"
		this.WriteJson(w, this.ReturnMsg)
		logs.Error(err.Error())
		return
	}
	//调用完毕
	sql = FTP_ECALCAUTO_GET1
	rows, err1 := dbobj.Default.Query(sql, doName, iso_cd)
	if err1 != nil {
		logs.Error(err1)
		this.ErrorCode = "0"
		this.ErrorMsg = "查询结果表失败,请联系管理员检查结果表"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	var one tmpAcalc
	var all []tmpAcalc
	for rows.Next() {
		err := rows.Scan(
			&one.item_id,
			&one.item_name_l1,
			&one.item_name_l2,
			&one.busiz_id,
			&one.org_par_bal,
			&one.ratio_bal,
			&one.cur_net_rate,
			&one.accd_int,
			&one.adj_int,
			&one.ftp_rate_b,
			&one.ftp_rate_a,
			&one.ftp_margin_a,
			&one.ftp_int_b,
			&one.ftp_int_a,
			&one.ftp_profit_a,
			//&one.domain_id,
			//&one.iso_currency_cd,
		)
		if err != nil {
			logs.Error(err)
			return
		}
		all = append(all, one)
	}
	if all == nil {
		this.ErrorCode = "0"
		this.ErrorMsg = "计算失败,请联系管理员检查计算数据"
		this.WriteJson(w, this.ReturnMsg)
		return
	}

	//	for _, val := range all {
	//		fmt.Println("singleVAL:", val)
	//	}
	rows.Close()
	//组装
	var (
		num        = -1
		sold       = "x-x-x"
		tmptarcalc Acalcss
		taracalc   []Acalcss
		tmpC       CommonParass
	)
	for _, val := range all {
		//s := val.item_id[:3]
		s := val.busiz_id
		//fmt.Println("s:", s)
		if s != sold {
			tmptarcalc.Id = strconv.Itoa(num + 2)
			tmptarcalc.Desc = val.item_name_l1
			tmptarcalc.Bid = val.busiz_id
			taracalc = append(taracalc, tmptarcalc)
			sold = s
			num++
		}
		//tmpC.OrgTerm        =val.
		//tmpC.OrgTermMult    =val.
		tmpC.StockBalance = val.org_par_bal
		tmpC.Ratio = val.ratio_bal
		tmpC.WeightRate = val.cur_net_rate
		tmpC.Income = val.accd_int
		tmpC.FtpValue = val.ftp_rate_b
		tmpC.ProfitRe = val.adj_int
		tmpC.AftpValue = val.ftp_rate_a
		tmpC.FtpDiffer = val.ftp_margin_a
		tmpC.BeforeFtpCost = val.ftp_int_b
		tmpC.AfterFtpCost = val.ftp_int_a
		tmpC.FtpProfit = val.ftp_profit_a
		taracalc[num].Params = append(taracalc[num].Params, tmpC)
	}
	//	for _, val := range taracalc {
	//		fmt.Println("id", val.Id)
	//		fmt.Println("bid", val.Bid)
	//		fmt.Println("desc", val.Desc)
	//		fmt.Println("sum", val.sum)
	//		for _, val := range val.Params {
	//			fmt.Println("params", val)
	//		}
	//	}

	sql = FTP_ECALCAUTO_GET2
	rows, err2 := dbobj.Default.Query(sql, doName, iso_cd)
	if err2 != nil {
		logs.Error(err2)
		return
	}
	var onet TotalResultss
	var allt []TotalResultss
	var domainid string
	var isocd string
	num = 0
	for rows.Next() {
		err := rows.Scan(
			&onet.Zcid,
			&onet.StockBalance,
			&onet.ExternTa,
			&onet.ExternTax,
			&onet.FtpValue,
			&onet.AftpValue,
			&onet.FtpDiffer,
			&onet.BeforeFtpCost,
			&onet.AfterFtpCost,
			&onet.FtpProfit,
			&domainid,
			&isocd,
		)
		num++
		onet.Zcid = strconv.Itoa(num)
		if err != nil {
			logs.Error(err)
			return
		}
		allt = append(allt, onet)
	}
	if allt == nil {
		this.ErrorCode = "0"
		this.ErrorMsg = "计算失败,请联系管理员检查数据"
		this.WriteJson(w, this.ReturnMsg)
		return
	}
	rows.Close()
	//为下载slsx做准备
	var cddiff []string
	var allali []string
	if len(allt) >= 10 {
		//存贷利差
		f1, _ := strconv.ParseFloat(allt[0].ExternTa, 64)
		f2, _ := strconv.ParseFloat(allt[6].ExternTa, 64)
		f3, _ := strconv.ParseFloat(allt[0].FtpDiffer, 64)
		f4, _ := strconv.ParseFloat(allt[6].FtpDiffer, 64)
		s := strconv.FormatFloat((f1 - f2), 'f', 4, 64)
		cddiff = append(cddiff, s)
		//贷款利差
		cddiff = append(cddiff, allt[0].FtpDiffer)
		//存款利差
		cddiff = append(cddiff, allt[6].FtpDiffer)
		//存贷错配利差
		ft := f1 - f2 - f3 - f4
		s = strconv.FormatFloat(ft, 'f', 4, 64)
		cddiff = append(cddiff, s)

		//
		f5, _ := strconv.ParseFloat(allt[8].ExternTa, 64)
		f6, _ := strconv.ParseFloat(allt[9].ExternTa, 64)
		f7, _ := strconv.ParseFloat(allt[8].FtpDiffer, 64)
		f8, _ := strconv.ParseFloat(allt[9].FtpDiffer, 64)
		ft = f5 - f6
		s = strconv.FormatFloat(ft, 'f', 4, 64)
		allali = append(allali, s)
		//资产利差
		allali = append(allali, allt[8].FtpDiffer)
		//负债利差
		allali = append(allali, allt[9].FtpDiffer)
		//全行错配利差
		ft = f5 - f6 - f7 - f8
		s = strconv.FormatFloat(ft, 'f', 4, 64)
		allali = append(allali, s)
	}
	myfile, _ := xlsx.OpenFile("./updownload/TmpExportCalcResult.xlsx")
	sheet := myfile.Sheet["Sheet1"]
	for i, val := range allt {
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = ResultName[i]
		cell = row.AddCell()
		cell.Value = val.StockBalance
		cell = row.AddCell()
		cell.Value = val.ExternTa
		cell = row.AddCell()
		cell.Value = val.ExternTax
		cell = row.AddCell()
		cell.Value = val.FtpValue
		cell = row.AddCell()
		cell.Value = val.AftpValue
		cell = row.AddCell()
		cell.Value = val.FtpDiffer
		cell = row.AddCell()
		cell.Value = val.BeforeFtpCost
		cell = row.AddCell()
		cell.Value = val.AfterFtpCost
		cell = row.AddCell()
		cell.Value = val.FtpProfit
	}
	row := sheet.AddRow()
	row = sheet.AddRow()
	cell := row.AddCell()
	cell.Value = ResultName1[0]
	for i, val := range cddiff {
		row = sheet.AddRow()
		cell := row.AddCell()
		cell.Value = ResultName1[i+1]
		cell = row.AddCell()
		cell.Value = val
	}
	row = sheet.AddRow()
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = ResultName2[0]
	for i, val := range allali {
		row = sheet.AddRow()
		cell := row.AddCell()
		cell.Value = ResultName2[i+1]
		cell = row.AddCell()
		cell.Value = val
	}
	ts := time.Now().Unix()
	tss := strconv.FormatInt(ts, 10)
	userid := this.Userid
	if userid == "" {
		logs.Error("session中域名为空")
	}
	fileSavename := "./updownload/calcresult/" + userid + tss + "calcauto.xlsx"
	xlsxerr := myfile.Save(fileSavename)
	if xlsxerr != nil {
		fmt.Printf(xlsxerr.Error())
		logs.Error(xlsxerr.Error())
	}
	//
	opcontent := "系统计算"
	this.InsertLogToDB(xlsxauto, opcontent, myapp)
	//打包最终json
	var outjson XlsxJson
	outjson.ErrorCode = "1"
	outjson.ErrorMsg = "计算成功"
	outjson.DfileName = fileSavename[1:]
	outjson.Result = &allt
	outjson.Single = &taracalc
	ojs, err := json.Marshal(&outjson)
	if err != nil {
		fmt.Println("json打包失败")
	}
	w.Write(ojs)
}
