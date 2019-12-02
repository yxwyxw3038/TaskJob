package daemon

import (
	"TaskJob/model"
	"TaskJob/setting"
	"TaskJob/util"
	"encoding/json"
	"errors"
	"time"

	"github.com/kirinlabs/HttpRequest"
)

func TaskRun(info model.RequestInfoModel) {
	var err error
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
	baseUrl := setting.ServerSetting.Url
	url := baseUrl + info.Url
	req := HttpRequest.NewRequest()
	hasmap := make(map[string]interface{})
	if info.Parameter != nil && len(info.Parameter) > 0 {
		for i := 0; i < len(info.Parameter); i++ {
			hasmap[info.Parameter[i].Column] = info.Parameter[i].Value
		}
	}
	var res *HttpRequest.Response
	switch info.Action {
	case "Post":
		res, err = req.Post(url, hasmap)
		break
	case "Get":
		res, err = req.Get(url, hasmap)
		break
	default:
		logger.Error("请求类型异常")
		return
	}
	if err != nil {
		logger.Error(err.Error())
		return
	}
	body, err := res.Body()
	if err != nil {
		logger.Error(err.Error())
		return
	}
	var tempJson = string(body)
	var publicResult model.PublicResult
	err = json.Unmarshal([]byte(tempJson), &publicResult)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	if publicResult.Code == "1" {
		if publicResult.Data != nil && publicResult.Data != "" {
			nextInfo, err := toRunInfo(publicResult.Data)
			if err != nil {
				logger.Error(err.Error())
				return
			}
			if nextInfo != nil && len(*nextInfo) > 0 {
				for i := 0; i < len(*nextInfo); i++ {
					var model model.RequestInfoModel
					model = (*nextInfo)[i]
					if model.IntervalTime > 0 {
						time.Sleep(time.Second * time.Duration(model.IntervalTime))
					}
					TaskRun(model)
				}
			}
		}
	} else {
		logger.Error(publicResult.Reason)
		return
	}

}
func toRunInfo(data interface{}) (*[]model.RequestInfoModel, error) {
	var err error
	logger := util.InitZapLog()
	var info []model.RequestInfoModel
	switch value := data.(type) {
	case string:
		err = json.Unmarshal([]byte(value), &info)
		if err != nil {
			logger.Error(err.Error())
			return nil, err
		}
		break
	case []model.RequestInfoModel:
		info = value
	default:
		return nil, errors.New("数据异常")
	}
	return &info, err
}
