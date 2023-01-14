package router

import (
	"apis/order-web/api/shop_cart"
	"apis/order-web/middlewares"

	"github.com/cloudwego/hertz/pkg/route"
)

func InitShopCartRouter(Router *route.RouterGroup) {
	GoodsRouter := Router.Group("shopcarts").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("", shop_cart.List)
		GoodsRouter.DELETE("/:id", shop_cart.Delete)
		GoodsRouter.POST("", shop_cart.New)
		GoodsRouter.PATCH("/:id", shop_cart.Update)
	}
}
