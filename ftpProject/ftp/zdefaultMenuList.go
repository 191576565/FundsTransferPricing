package ftp

import (
	"ftpProject/conf"
	"ftpProject/logs"

	"ftpProject/utils/cacheutil"
)

type DefaultMenuM struct {
	Res_id string
}
type DefaultMenu struct {
	RouteControl
}

func (this *DefaultMenu) Get() {
	//取缓存
	var resdata []FtpResData
	if val, ok := cacheutil.BeeCache.Get(this.Userid).(CacheData); ok {
		resdata = val.ResData
	} else {
		logs.Error("缓存中菜单信息获取失败")
		this.ShowPageError("登录超时,请重新登录", conf.FtpConf.AuthLoginIp)
	}

	Id := this.GetString("Id")
	var one DefaultMenuM
	var rst []DefaultMenuM
	for _, val := range resdata {
		if val.Res_up_id == Id {
			one.Res_id = val.Res_id
			rst = append(rst, one)
		}
	}
	this.WriteJsonStr(rst)
}
