package main

import (
	"fmt"
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/base"
	"github.com/alibaba/sentinel-golang/core/flow"
	"log"
)

func main() {
	//先初始化sentinel
	err := sentinel.InitDefault()
	if err != nil {
		log.Fatalf("初始化sentinel 异常: %v", err)
	}

	//配置限流规则
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "some-test",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //匀速通过
			Threshold:              10,
			StatIntervalInMs:       1000,
		},

		{
			Resource:               "some-test2",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, //直接拒绝
			Threshold:              10,
			StatIntervalInMs:       1000,
		},
	})

	if err != nil {
		log.Fatalf("加载规则失败: %v", err)
	}

	for i := 0; i < 12; i++ {
		e, b := sentinel.Entry("some-test", sentinel.WithTrafficType(base.Inbound))
		if b != nil {
			fmt.Println("限流了")
		} else {
			fmt.Println("检查通过")
			e.Exit()
		}
	}
}
