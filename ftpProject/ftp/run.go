package ftp

import "github.com/astaxie/beego"

func Start() {
	//beego额外配置
	beego.SetLogFuncCall(true)
	beego.SetStaticPath("/updownload", "updownload")
	InitRouter()
	beego.Run()
}

//路由
func InitRouter() {
	beego.Router("/logout", &LogOutCtl{})
	beego.Router("/test", &TestCtl{})
	//
	beego.Router("/platform/select", &GoEntry{})
	beego.Router("/platform/FtpSysMethod", &FtpSysMethodCtl{})
	//登录后页面
	beego.Router("/platform/LoginIndexPage", &LoginIndexPage{})
	//首页面加载资源
	beego.Router("/platform/DefaultMenu", &DefaultMenu{})
	//页面加载资源
	beego.Router("/platform/MenuPage", &MenuPage{})

	beego.Router("/platform/OrgTree", &OrgTreeCtl{})
	//-------------------------------------------------------------
	beego.Router("/platform/HandleLogsPage", &HandleLogPage{})
	beego.Router("/platform/HandleLogs", &HandleLogCtl{})
	//币种
	beego.Router("/platform/CurrencyStatus", &CurrencyStatusCtl{})
	beego.Router("/platform/ComCurrencyPage", &ComCurrencyPage{})
	beego.Router("/platform/ComCurrency", &ComCurrencyCtl{})
	//-------------------------------------------------------------

	//业务单元
	{
		beego.Router("/platform/FtpBusizPage", &FtpBusizPage{})
		beego.Router("/platform/FtpBusiz", &FtpBusizCtl{})
		//业务单元层级树
		beego.Router("/platform/FtpBusizStruct", &FtpBusizStructCtl{})
		//--注册的公共，动态下拉框的路由
		//定价方法
		beego.Router("/platform/FtpftpMethod", &FtpSysMethodCtl{})
		//域名
		//beego.Router("/platform/ftpDomainInfo", &ftpDomainInfo{})
		//曲线
		beego.Router("/platform/MasCurveDefine", &MasCurveDefineCtl{})
		//调节项
		beego.Router("/platform/FtpAdjustment", &FtpAdjustmentCtl{})
		//redemption
		beego.Router("/platform/FtpRedemption", &FtpRedemptionCtl{})

	}

	//曲线信息
	{
		//曲线定义
		beego.Router("/platform/FtpCurveDefPage", &FtpCurveDefPage{})
		beego.Router("/platform/FtpCurveDef", &FtpCurveDefCtl{})
		//曲线类型
		beego.Router("/platform/FtpCurveType", &FtpCurveTypeCtl{})
		//曲线信息，点击查询页面出来
		beego.Router("/platform/FtpCurveInfoPage", &FtpCurveInfoPage{})
		beego.Router("/platform/FtpCurveInfo", &FtpCurveInfoCtl{})
		//曲线值保存
		beego.Router("/platform/FtpCurveSave", &FtpCurveSaveCtl{})
		//
		beego.Router("/platform/FtpCurveInfoStruct", &FtpCurveInfoStructCtl{})
		//FtpRepType
		beego.Router("/platform/FtpRepType", &FtpRepTypeCtl{})

	}

	//价格试算
	{
		beego.Router("/platform/FtpValueCalcPage", &FtpValueCalcPage{})
		beego.Router("/platform/FtpValueCalc", &FtpValueCalcCtl{})
		//2016.8.11 单笔试算用
		beego.Router("/platform/BusizInfoCalc", &BusizInfoCalcCtl{})
		beego.Router("/platform/RateAdjustType", &RateAdjustTypeCtl{})
		beego.Router("/platform/AccrualCdAttr", &AccrualCdAttrCtl{})
		beego.Router("/platform/PaymentTypeAttr", &PaymentTypeAttrCtl{})
		beego.Router("/platform/MasDimCurrency", &MasDimCurrencyCtl{})
	}

	//基础价格计算
	{
		beego.Router("/platform/FtpBaseValueCalPage", &FtpBaseValueCalPage{})
		//基础价格计算
		beego.Router("/platform/FtpPatchCalc", &FtpPatchCalc{})
		beego.Router("/platform/FtpBaseValueCal", &FtpBaseValueCalCtl{})
		//+动态输入输出
		beego.Router("/platform/FtpDispatchInput", &FtpDispatchInputCtl{})
		beego.Router("/platform/FtpDispatchOutput", &FtpDispatchOutputCtl{})
		//基础价格试算FtpDispatchRealt
		beego.Router("/platform/FtpDispatchRealt", &FtpDispatchRealtCtl{})
	}

	//调节项
	{
		beego.Router("/mas/ftp/adjust/page", &FtpAdjustDefPage{})
		beego.Router("/mas/ftp/adjust/policy/page", &FtpPolicyAdjustDefPage{})
		beego.Router("/mas/ftp/adjust/info", &FtpAdjustDefCtl{})
		beego.Router("/mas/ftp/adjust/termLiq/page", &FtpAdjustTermLiqPage{})
		beego.Router("/mas/ftp/adjust/termLiq/config", &FtpAdjustTermLiqInfoCtl{})

		beego.Router("/mas/ftp/adjust/treasurep/page", &FtpAdjustTreasurePage{})
		beego.Router("/mas/ftp/adjust/treasurep/config", &FtpAdjustTreasureInfoCtl{})

		beego.Router("/mas/ftp/adjust/termLiq/tlp", &AdjustTlpCtl{})
		beego.Router("/mas/ftp/adjust/reserve/page", &FtpAdjustReservePage{})
		beego.Router("/mas/ftp/adjust/reserve/config", &FtpAdjustReserveInfoCtl{})
		//FtpAdjInfo  FtpAdjInfoPage
		beego.Router("/mas/ftp/FtpAdjInfoPage", &FtpAdjInfoPage{})
		beego.Router("/mas/ftp/FtpAdjInfo", &FtpAdjInfoCtl{})
		//调节项状态
		beego.Router("/mas/ftp/FtpAdjStatus", &FtpAdjStatusCtl{})
		//调节项类型
		beego.Router("/mas/ftp/FtpAdjType", &FtpAdjTypeCtl{})
		//z准备金业务单元下拉框
		beego.Router("/mas/ftp/FtpReBusiz", &FtpReBusizCtl{})
	}

	//产品
	{
		//FtpProductInfo
		beego.Router("/mas/ftp/FtpProductInfoPage", &FtpProductInfoPage{})
		beego.Router("/mas/ftp/FtpProductInfo", &FtpProductInfoCtl{})
		beego.Router("/mas/ftp/ProductTree", &ProductTreeCtl{})
	}
	//政策调节FtpAdjustPolicyPage
	{
		beego.Router("/mas/ftp/FtpAdjustPolicyPage", &FtpAdjustPolicyPage{})
		beego.Router("/mas/ftp/FtpAdjustPolicy", &FtpAdjustPolicyCtl{})
		//FtpAlType资产或者负债
		beego.Router("/mas/ftp/FtpAlType", &FtpAlTypeCtl{})
		//业务单元含有政策调节的
		beego.Router("/mas/ftp/FtpBusizPolicy", &FtpBusizPolicyCtl{})
		beego.Router("/mas/ftp/IndustryTree", &IndustryTreeCtl{})
		//FtpAdjUpload
		beego.Router("/mas/ftp/FtpAdjUpload", &FtpAdjUpload{})
		//FtpAdjDownload
		beego.Router("/mas/ftp/FtpAdjDownload", &FtpAdjDownloadCtl{})
		//TopOrg顶级机构
		beego.Router("/mas/ftp/TopOrg", &TopOrgCtl{})
		//调用存储过程
		beego.Router("/mas/ftp/FtpCallPCheckProc", &FtpCallPCheckProc{})
		beego.Router("/mas/ftp/FtpPCheckResult", &FtpPCheckResultCtl{})

	}

	//整体试算FtpEnsembleCalaPage
	{
		beego.Router("/platform/FtpEnsembleCalaPage", &FtpEnsembleCalaPage{})
		beego.Router("/platform/FtpEnsembleCala", &FtpEnsembleCalc{})
		//FtpEnsembleCalcAuto
		beego.Router("/platform/FtpEnsembleCalcAuto", &FtpEnsembleCalcAuto{})

	}

	beego.Router("/mas/ftp/FtpFormBackup", &FtpFormBackup{})
	beego.Router("/mas/ftp/CurveInfoInput", &CurveInfoInput{})
	/*
		//测试FtpOrgInfo  CurveDown FtpFormBackup CurveInfoInput
		beego.Router("/mas/ftp/FtpOrgInfo", &FtpOrgInfo{})


	*/
}
