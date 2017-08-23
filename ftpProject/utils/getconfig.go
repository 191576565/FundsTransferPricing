package utils

import (
	"ftpProject/logs"
	"os"
	"path/filepath"

	"github.com/astaxie/beego/config"
)

func GetAppConf(key string) string {
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	appConfigPath := filepath.Join(workPath, "conf", "app.conf")
	config, err := config.NewConfig("ini", appConfigPath)

	if err != nil {
		logs.Error(err)
		return ""
	}
	return config.String(key)
}
