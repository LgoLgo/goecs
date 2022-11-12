package router

import (
	"apis/oss-web/handler"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitOssRouter(Router *route.RouterGroup) {
	OssRouter := Router.Group("oss")
	{
		OssRouter.GET("/token", handler.Token)
		OssRouter.POST("/callback", handler.Request)
	}
}
