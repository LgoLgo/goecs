package router

import (
	"apis/goods-web/api/brands"
	"apis/goods-web/middlewares"

	"github.com/cloudwego/hertz/pkg/route"
)

func InitBrandRouter(Router *route.RouterGroup) {
	BrandRouter := Router.Group("brands").Use(middlewares.Trace())
	{
		BrandRouter.GET("", brands.BrandList)
		BrandRouter.DELETE("/:id", brands.DeleteBrand)
		BrandRouter.POST("", brands.NewBrand)
		BrandRouter.PUT("/:id", brands.UpdateBrand)
	}

	CategoryBrandRouter := Router.Group("categorybrands")
	{
		CategoryBrandRouter.GET("", brands.CategoryBrandList)
		CategoryBrandRouter.DELETE("/:id", brands.DeleteCategoryBrand)
		CategoryBrandRouter.POST("", brands.NewCategoryBrand)
		CategoryBrandRouter.PUT("/:id", brands.UpdateCategoryBrand)
		CategoryBrandRouter.GET("/:id", brands.GetCategoryBrandList)
	}
}
