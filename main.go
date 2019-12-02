package main

import (
	"TaskJob/daemon"
	"TaskJob/setting"
	"TaskJob/util"
	"github.com/urfave/cli"
	"log"
	"os"
	"runtime"
	"strconv"
	"time"
)

// func jump() {
// 	var ch chan int
// 	logger := util.InitZapLog()
// 	//定时任务
// 	setting.Setup()
// 	times := setting.AppSetting.JumpTime
// 	ticker := time.NewTicker(time.Second * time.Duration(times))
// 	go func() {
// 		for range ticker.C {
// 			count:=runtime.NumGoroutine()
// 			logger.Debug("心跳服务，总运行Goroutine数量"+ strconv.Itoa(count))
// 		}
// 		ch <- 1
// 	}()
// 	<-ch
// }
// func jump() {
// 	logger := util.InitZapLog()
// 	count:=runtime.NumGoroutine()
// 	logger.Debug("心跳服务，总运行Goroutine数量"+ strconv.Itoa(count))
// }
func main() {
	//实例化cli
	app := cli.NewApp()
	setting.Setup()
	//Name可以设定应用的名字
	app.Name = "自动任务工具"
	// Version可以设定应用的版本号
	app.Version = "1.0.0"
	portStr := setting.AppSetting.Port
	port, _ := strconv.Atoi(portStr)
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:  "port, p",
			Value: port,
			Usage: "运行端口",
		},
	}

	// 接受os.Args启动程序
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	go daemon.Run()
	//jump()

	// var wait util.WaitGroupWrapper
	// times := setting.AppSetting.JumpTime
	// ticker := time.NewTicker(time.Second * time.Duration(times))
	// for range ticker.C {
	// 	wait.Wrap(jump)
	// 	//jump()
	// }
	logger := util.InitZapLog()
	times := setting.AppSetting.JumpTime
	for {
		logger.Sync()
		count := runtime.NumGoroutine()
		logger.Debug("心跳服务，总运行Goroutine数量" + strconv.Itoa(count))
		time.Sleep(time.Second * time.Duration(times))
	}
}
