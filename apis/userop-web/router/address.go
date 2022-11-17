package router

import (
	"apis/userop-web/api/address"
	"apis/userop-web/middlewares"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitAddressRouter(Router *route.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)          // 轮播图列表页
		AddressRouter.DELETE("/:id", middlewares.JWTAuth(), address.Delete) // 删除轮播图
		AddressRouter.POST("", middlewares.JWTAuth(), address.New)          //新建轮播图
		AddressRouter.PUT("/:id", middlewares.JWTAuth(), address.Update)    //修改轮播图信息
	}
}
