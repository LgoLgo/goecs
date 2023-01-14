package router

import (
	"apis/userop-web/api/message"
	"apis/userop-web/middlewares"

	"github.com/cloudwego/hertz/pkg/route"
)

func InitMessageRouter(Router *route.RouterGroup) {
	MessageRouter := Router.Group("message").Use(middlewares.JWTAuth())
	{
		MessageRouter.GET("", message.List)
		MessageRouter.POST("", message.New)
	}
}
