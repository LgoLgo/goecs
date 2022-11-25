package router

import (
	"apis/goods-web/api/banners"
	"apis/goods-web/middlewares"

	"github.com/cloudwego/hertz/pkg/route"
)

func InitBannerRouter(Router *route.RouterGroup) {
	BannerRouter := Router.Group("banners").Use(middlewares.Trace())
	{
		BannerRouter.GET("", banners.List)
		BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Delete)
		BannerRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.New)
		BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdminAuth(), banners.Update)
	}
}
