package initialize

import (
	sentinel "github.com/alibaba/sentinel-golang/api"
	"github.com/alibaba/sentinel-golang/core/flow"
	"go.uber.org/zap"
)

func InitSentinel() {
	err := sentinel.InitDefault()
	if err != nil {
		zap.S().Fatalf("Initsentinel error: %v", err)
	}

	// Configure current limiting rules
	_, err = flow.LoadRules([]*flow.Rule{
		{
			Resource:               "goods-list",
			TokenCalculateStrategy: flow.Direct,
			ControlBehavior:        flow.Reject, // pass at a constant speed
			Threshold:              20,          // 100ms can only have 1W concurrency, 1s is 10W concurrency
			StatIntervalInMs:       6000,
		},
	})

	if err != nil {
		zap.S().Fatalf("Failed to load rules:%v", err)
	}
}
