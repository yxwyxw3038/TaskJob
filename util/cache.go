package util

import (
	"github.com/muesli/cache2go"
	"TaskJob/setting"
	"time"
	"errors"
)

func NewCache() *cache2go.CacheTable {
	// setting.Setup()
	CacheDataBaseName := setting.CacheInfoSetting.CacheDataBaseName
	return cache2go.Cache(CacheDataBaseName)
}

func GetCacheStr(code string) (string,error) {
	// defer func(){
	// 	if p := recover(); p != nil {
	// 	 return "",error.new("获取缓存异常")
	// 	}
	// }()
	cache := NewCache()
	res, err := cache.Value(code)
	if err != nil {
		return "" ,err
	}
    return res.Data().(string),nil
}

func GetCacheCount(code string) (int64,error) {
	// defer func(){
	// 	if p := recover(); p != nil {
	// 	 return "",error.new("获取缓存异常")
	// 	}
	// }()
	cache := NewCache()
	res, err := cache.Value(code)
	if err != nil {
		return 0 ,err
	}
    return res.AccessCount(),nil
}

func AutoCacheCount(code string) (bool,error) {
	// defer func(){
	// 	if p := recover(); p != nil {
	// 	 return false,error.new("自动设置缓存异常")
	// 	}
	// }()
	res, err :=GetCacheCount(code)
	if err != nil {
		err=RegisterCodeInt(code)
		if err != nil {
			return false,err
		}
	    return true,nil

	}
	if res==0 {
		return false,errors.New("缓存内容异常") 
	}
    if res>=3 {
		return false,nil 
	}
	return true,nil
}

func RegisterCodeInt(code string ) error {
	// defer func(){
	// 	if p := recover(); p != nil {
	// 	 return error.new("注册缓存异常")
	// 	}
	// }()
	cache := NewCache()
	AllIntervalTime := setting.CacheInfoSetting.AllIntervalTime
	cache.Add(code, time.Second * time.Duration(AllIntervalTime), 1)
	return nil
}