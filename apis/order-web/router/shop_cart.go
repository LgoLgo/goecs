package router

import (
	"apis/order-web/api/shop_cart"
	"apis/order-web/middlewares"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitShopCartRouter(Router *route.RouterGroup) {
	GoodsRouter := Router.Group("shopcarts").Use(middlewares.JWTAuth())
	{
		GoodsRouter.GET("", shop_cart.List)          //购物车列表
		GoodsRouter.DELETE("/:id", shop_cart.Delete) //删除条目
		GoodsRouter.POST("", shop_cart.New)          //添加商品到购物车
		GoodsRouter.PATCH("/:id", shop_cart.Update)  //修改条目
	}
}
