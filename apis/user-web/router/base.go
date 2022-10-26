package router

import (
	"E-commerce-system/apis/user-web/api"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitBaseRouter(Router *route.RouterGroup) {
	BaseRouter := Router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetCaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
