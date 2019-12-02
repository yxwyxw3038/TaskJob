package setting

import (
	"gopkg.in/ini.v1"
	"log"
)

type App struct {
	Port         string
	JumpTime     int32
	IntervalTime int32
	TaskList     string
}
type Server struct {
	Url string
}
type CacheInfo struct {
	CacheDataBaseName string
	AllIntervalTime   int32
}

var AppSetting = &App{}
var ServerSetting = &Server{}
var CacheInfoSetting = &CacheInfo{}
var config *ini.File

func Setup() {
	var err error
	config, err = ini.Load("app.ini")
	if err != nil {
		log.Fatal("Fail to parse 'app.ini': %v", err)
	}
	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("cacheinfo", CacheInfoSetting)
}

func mapTo(section string, v interface{}) {
	err := config.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
