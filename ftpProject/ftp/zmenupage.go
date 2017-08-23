package ftp

import "ftpProject/utils/cacheutil"

type MenuPageM struct {
	Res_id       string
	Res_name     string
	Res_bg_color string
	Res_class    string
	Res_url      string
	Res_img      string
	Group_id     string
	Res_up_id    string
}
type MenuPage struct {
	RouteControl
}

func (this *MenuPage) Get() {
	//	r := this.Ctx.Request
	//取缓存
	var resdata []FtpResData
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		resdata = val.ResData
	}
	//fmt.Println(resdata[0])
	//MenuPageM 30208000000000
	var one MenuPageM
	var rst []MenuPageM
	if this.GetString("Id") != "" {

		for _, val := range resdata {
			if val.Res_type == "1" {
				one.Res_id = val.Res_id
				one.Res_name = val.Res_name
				one.Res_bg_color = val.Res_color
				one.Res_class = val.Res_css
				one.Res_url = val.Res_url
				one.Res_img = val.Res_icon
				one.Group_id = "11"
				one.Res_up_id = val.Res_up_id
				rst = append(rst, one)
			}
		}
	}
	//	var rst1 []MenuPageM
	//	var rst2 []MenuPageM
	//	for _, val := range rst {
	//		if val.Res_up_id == "30208000000000" {
	//			rst1 = append(rst1, val)
	//		} else {
	//			rst2 = append(rst2, val)
	//		}
	//	}
	//	rst1 = append(rst1, rst2...)
	//
	//	this.Data["json"] = &rst
	//	this.ServeJSON()
	this.WriteJsonStr(rst)
}
