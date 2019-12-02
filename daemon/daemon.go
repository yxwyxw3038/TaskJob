package daemon

import (
	"TaskJob/model"
	"TaskJob/setting"
	"TaskJob/util"
	"encoding/json"
	"strings"
	"sync"
	"time"
)

type DataSyncMutex struct {
	mutex  sync.RWMutex
	isSync bool
}

func Run() {

	logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			switch x := p.(type) {
			case string:
				logger.Error(x)
			case error:
				logger.Error(x.Error())
			default:
				logger.Error("未知异常")
			}

		}
	}()

	logger.Debug("启动后台进程")
	intervalTime := setting.AppSetting.IntervalTime
	taskList := setting.AppSetting.TaskList
	if taskList == "" {
		logger.Error("启动任务列表为空!")
		return
	}
	var date []string
	date = strings.Split(taskList, ",")
	if len(date) <= 0 {
		logger.Error("启动任务列表为空!")
		return
	}
	var wait util.WaitGroupWrapper
	// time.Sleep(time.Second * time.Duration(intervalTime))
	for i := 0; i < len(date); i++ {
		msg := date[i]
		time.Sleep(time.Second * time.Duration(intervalTime))
		wait.Wrap(func() { registerTask(msg, intervalTime) })
	}
	wait.Wait()
}

func registerTask(taskfileName string, intervalTime int32) {
	logger := util.InitZapLog()
	defer func() {
		if p := recover(); p != nil {
			switch x := p.(type) {
			case string:
				logger.Error(x)
			case error:
				logger.Error(x.Error())
			default:
				logger.Error("未知异常")
			}

		}
	}()
	msg, err := util.GetFileInfo(taskfileName)
	if err != nil {

		logger.Error(err.Error())
		return
	}
	var info model.RequestInfoModel
	err = json.Unmarshal([]byte(msg), &info)
	if err != nil {
		logger.Error(err.Error())
		return
	}

	for {
		logger.Sync()
		logger.Debug(info.Url)
		logger.Debug(info.Action)
		TaskRun(info)
		time.Sleep(time.Second * time.Duration(info.IntervalTime))
		// logger.Debug(msg)
		// time.Sleep(time.Second * time.Duration(30))
	}
}