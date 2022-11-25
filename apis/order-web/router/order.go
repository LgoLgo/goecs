package router

import (
	"github.com/cloudwego/hertz/pkg/route"

	"apis/order-web/api/order"
	"apis/order-web/api/pay"
	"apis/order-web/middlewares"
)

func InitOrderRouter(Router *route.RouterGroup) {
	OrderRouter := Router.Group("orders").Use(middlewares.JWTAuth()).Use(middlewares.Trace())
	{
		OrderRouter.GET("", order.List)
		OrderRouter.POST("", order.New)
		OrderRouter.GET("/:id", order.Detail)
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
