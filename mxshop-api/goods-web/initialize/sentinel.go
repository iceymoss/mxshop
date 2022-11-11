package initialize

import (
	"log"

	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
)

//InitSentinel 初始化sentinel
func InitSentinel() {
	//基于sentinel的qps限流
	//必须初始化
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
	}

	//配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "goods-list",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //超过直接拒绝
			Threshold:              3,           //请求次数
			StatIntervalInMs:       2000,        //允许时间内
		},
		{
			Resource:               "goods-stock",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //超过直接拒绝
			Threshold:              3,           //请求次数
			StatIntervalInMs:       2000,        //允许时间内
		},
	})

	if err != nil {
		log.Fatalf("Unexpected error: %+v", err)
		return
	}

}
