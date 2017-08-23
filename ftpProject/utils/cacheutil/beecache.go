package cacheutil

import (
	"ftpProject/conf"
	"ftpProject/logs"
	"strconv"
	"time"

	"github.com/astaxie/beego/cache"
)

//加在beego memcache代码中
//func NNewMemoryCache() *MemoryCache {
//	cache := MemoryCache{items: make(map[string]*MemoryItem)}
//	return &cache
//}
//func (bc *MemoryCache) Refresh(name string) error {
//	bc.Lock()
//	defer bc.Unlock()
//	if it, ok := bc.items[name]; ok {
//		it.createdTime = time.Now()
//	}
//	return nil
//}
func newMcache(config string) *cache.MemoryCache {
	memcache := cache.NNewMemoryCache()
	err := memcache.StartAndGC(config)
	if err != nil {
		return nil
	}
	return memcache
}

//var NewMcache(`{"interval":60}`)
var BeeCache *cache.MemoryCache

func PutIntoBeeCache(name string, value interface{}) {
	btime, _ := strconv.ParseInt(conf.FtpConf.SessTime, 10, 64)
	//为了比session时间延迟一分钟
	btime += 60
	t := time.Duration(btime)
	BeeCache.Put(name, value, t*time.Second)
}
func init() {
	stime := conf.FtpConf.SessTime
	config := `{"interval":`
	if stime != "" {
		config += stime + "}"
	} else {
		config = `{"interval":3600}`
	}
	BeeCache = newMcache(config)
	logs.Debug("cache初始化成功")
}
