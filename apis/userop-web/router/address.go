package router

import (
	"apis/userop-web/api/address"
	"apis/userop-web/middlewares"

	"github.com/cloudwego/hertz/pkg/route"
)

func InitAddressRouter(Router *route.RouterGroup) {
	AddressRouter := Router.Group("address")
	{
		AddressRouter.GET("", middlewares.JWTAuth(), address.List)
		AddressRouter.DELETE("/:id", middlewares.JWTAuth(), address.Delete)
		AddressRouter.POST("", middlewares.JWTAuth(), address.New)
		AddressRouter.PUT("/:id", middlewares.JWTAuth(), address.Update)
	}
}
