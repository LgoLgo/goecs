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
		OrderRouter.GET("", order.List)       // 订单列表
		OrderRouter.POST("", order.New)       // 新建订单
		OrderRouter.GET("/:id", order.Detail) // 订单详情
	}
	PayRouter := Router.Group("pay")
	{
		PayRouter.POST("alipay/notify", pay.Notify)
	}
}
