package router

import (
	"github.com/cloudwego/hertz/pkg/route"

	"apis/userop-web/api/user_fav"
	"apis/userop-web/middlewares"
)

func InitUserFavRouter(Router *route.RouterGroup) {
	UserFavRouter := Router.Group("userfavs")
	{
		UserFavRouter.DELETE("/:id", middlewares.JWTAuth(), user_fav.Delete)
		UserFavRouter.GET("/:id", middlewares.JWTAuth(), user_fav.Detail)
		UserFavRouter.POST("", middlewares.JWTAuth(), user_fav.New)
		UserFavRouter.GET("", middlewares.JWTAuth(), user_fav.List)
	}
}
