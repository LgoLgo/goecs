package router

import (
	"apis/userop-web/api/message"
	"apis/userop-web/middlewares"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitMessageRouter(Router *route.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	{
		MessageRouter.GET("", message.List) // 轮播图列表页
		MessageRouter.POST("", message.New) //新建轮播图
	}
}
