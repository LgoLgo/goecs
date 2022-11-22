package router

import (
	"github.com/cloudwego/hertz/pkg/route"

	"apis/user-web/api"
)

func InitBaseRouter(Router *route.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
