package ftp

import (
	"encoding/json"
	"ftpProject/dbobj"
	"ftpProject/logs"

	"net/rpc"
	"os"
	"path/filepath"
	"strconv"

	"github.com/astaxie/beego/config"
	//	"text/template"
)

//批次计算
type FtpPatchCalc struct {
	ReturnMsg
	RouteControl
}
type FtpPatchCalcResult struct {
	ReturnMsg
	PatchCalcValue string
}
type FtpPatchInfo struct {
	PatchId        string `json:"DispatchId"`
	InputSouceCd   string
	OutputResultCd string
	PatchName      string `json:"DispatcName"`
	Asofdate       string `json:"batchDate"`
	DomainId       string
	Offset         string `json:"StartOffset"`
	Limit          string `json:"MaxLimit"`
	patchsumrows   string
}

func (this *FtpPatchCalc) Post() {
	r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
	r.ParseForm()
	var (
		result      FtpPatchCalcResult
		allpatch    []FtpPatchInfo
		allpatchtmp []FtpPatchInfo
		msg         string = "批次名称为:"
		flag        bool   = false
	)
	//增加域 2016.9.10
	doName := this.Domainid
	if doName == "" {
		logs.Error("session中域名为空")
		return
	}
	myjson := []byte(r.FormValue("JSON"))
	err := json.Unmarshal(myjson, &allpatchtmp)
	if err != nil {
		logs.Error("json 解析失败")
		return
	}
	//step add: zdd tichu
	lenth := len(allpatchtmp)
	for i, val := range allpatchtmp {
		sql := `select count(*) from ` + val.InputSouceCd + ` where as_of_date=to_date(:1,'YYYY-MM-DD') and domain_id=:2`
		//		sql := `select count(*) from ` + val.InputSouceCd + ` t inner join ftp_busiz_info t1
		//on t.busiz_id=t1.busiz_id and t.domain_id=t1.domain_id
		//where t1.ftp_flag='0' and t.as_of_date=to_date(:1,'YYYY-MM-DD') and t.domain_id=:2`
		rows, err := dbobj.Default.Query(sql, val.Asofdate, doName)
		if err != nil {
			logs.Error("跑批次时候，查询总行数失败")
			this.ErrorCode = "0"
			this.ErrorMsg = "查询" + val.PatchId + "批次号的输入" + val.InputSouceCd + "表总行数失败,请联系管理员检查此表"
			this.WriteJson(w, this.ReturnMsg)
			return
		}
		var sumrows string
		for rows.Next() {
			err := rows.Scan(&sumrows)
			if err != nil {
				logs.Error("跑批次时候，取值总行数失败")
				this.ErrorCode = "0"
				this.ErrorMsg = "查询" + val.PatchId + "批次号的输入" + val.InputSouceCd + "表总行数失败,请联系管理员检查此表"
				this.WriteJson(w, this.ReturnMsg)
				return
			}
		}

		if sumrows == "0" {
			flag = true
			if i == lenth-1 {
				msg = msg + val.PatchName
			} else {
				msg = msg + val.PatchName + ","
			}
		} else {
			//判断总行数和limit关系
			allrowsInt, _ := strconv.Atoi(sumrows)
			startoffset, _ := strconv.Atoi(val.Offset)
			limitInt, _ := strconv.Atoi(val.Limit)
			if allrowsInt > limitInt+startoffset {
				val.patchsumrows = val.Limit
			} else {
				val.patchsumrows = strconv.Itoa(allrowsInt - startoffset)
			}
			allpatch = append(allpatch, val)
		}
		rows.Close()
	}

	msg = msg + "的源表内没有对应域名的数据"
	//fmt.Println("batch:", allpatch)
	//step add 清除结果表数据
	//	for _, val := range allpatch {
	//		//sql := "delete from " + val.OutputResultCd + " where domain_id=:1"
	//		//err = dbobj.Default.Exec(sql, doName)
	//		err = dbobj.Default.Exec("call PROC_ADD_PARTITIONS(to_date(:1,'YYYY-MM-DD'),:2,:3)", val.Asofdate, val.OutputResultCd, doName)
	//		if err != nil {
	//			logs.Error(err)
	//			logs.Errorf("清除结果表数据失败")
	//			result.ErrorCode = "0"
	//			result.ErrorMsg = "清除结果表数据失败:" + val.PatchName
	//			ojs, err := json.Marshal(result)
	//			if err != nil {
	//				logs.Error(err)
	//			}
	//			w.Write(ojs)
	//			return
	//		}
	//	}

	//向初始化表里面插数据
	//step 1: 删除表中已完成和异常的
	tx, _ := dbobj.Default.Begin()
	sql := FTP_DISPATCHCALC_POST1
	_, err = tx.Exec(sql)
	if err != nil {
		logs.Error("删除停止数据失败")
		tx.Rollback()
		return
	}
	//
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")

	red, err := config.NewConfig("ini", appConfigPath)
	if err != nil {
		tx.Rollback()
		logs.Error("cant not read ./conf/app.conf.please check this file.")
		result.ErrorCode = "0"
		result.ErrorMsg = "读取计算引擎地址出错，请联系管理员检查计算引擎配置文件"
		ojs, err := json.Marshal(result)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}
	calcip := red.String("Calc.IP")
	//
	conn, err := rpc.DialHTTP("tcp", calcip)
	if err != nil {
		logs.Error(err)
		result.ErrorCode = "0"
		result.ErrorMsg = "与引擎服务建立连接失败,请联系管理员检查引擎服务配置和引擎是否启动"
		ojs, err := json.Marshal(result)
		if err != nil {
			logs.Error(err)
		}
		w.Write(ojs)
		return
	}

	for _, val := range allpatch {

		//step 1.1  插入时候先删除
		sql := FTP_DISPATCHCALC_POST2
		_, err = tx.Exec(sql, val.PatchId, val.Asofdate, doName)

		if err != nil {
			tx.Rollback()
			logs.Error(err)
			logs.Errorf("删除失败,批次号: %s", val.PatchId)
			this.ErrorCode = "0"
			this.ErrorMsg = "删除相同批次失败，请联系管理员检查批次表"
			this.WriteJson(w, this.ReturnMsg)
			return
		}

		//step 2: 向表里插入数据
		sql = FTP_DISPATCHCALC_POST3
		_, err = tx.Exec(sql, val.PatchId, val.Asofdate, "1", "0", val.patchsumrows, "运行中", val.PatchName, doName)
		if err != nil {
			logs.Errorf("插入数据失败,批次号: %s", val.PatchId)
			this.ErrorCode = "0"
			this.ErrorMsg = "插入运行批次表失败，请联系管理员检查批次表"
			this.WriteJson(w, this.ReturnMsg)
			tx.Rollback()
			return

		}

	}

	for _, val := range allpatch {
		var para = []string{val.Asofdate, doName, val.Offset, val.Limit, val.PatchId}
		go func(conn *rpc.Client, args []string) {
			var replay string
			err := conn.Call("RpcSrvOfFTP.StartFTPCalc", args, &replay)

			if err != nil {
				logs.Errorf("调用远程服务失败，批次是：%s", para[0])
			}
		}(conn, para)
		//
		opcontent := "批次运行，批次号,日期,所属域分别为:" + val.PatchId + " " + val.Asofdate + " " + val.DomainId
		this.InsertLogToDB(batchrun, opcontent, myapp)
	}
	tx.Commit()
	result.ErrorCode = "1"
	if flag {
		result.ErrorMsg = msg
	} else {
		result.ErrorMsg = "批次运行中"
	}
	ojs, err := json.Marshal(result)
	if err != nil {
		logs.Error(err)
	}
	w.Write(ojs)
}
