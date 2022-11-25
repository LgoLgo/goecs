package router

import (
	"apis/goods-web/api/category"
	"apis/goods-web/middlewares"

	"github.com/cloudwego/hertz/pkg/route"
)

func InitCategoryRouter(Router *route.RouterGroup) {
	CategoryRouter := Router.Group("categorys").Use(middlewares.Trace())
	{
		CategoryRouter.GET("", category.List)
		CategoryRouter.DELETE("/:id", category.Delete)
		CategoryRouter.GET("/:id", category.Detail)
		CategoryRouter.POST("", category.New)
		CategoryRouter.PUT("/:id", category.Update)
	}
}
