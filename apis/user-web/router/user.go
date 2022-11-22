package router

import (
	"github.com/cloudwego/hertz/pkg/route"

	"apis/user-web/api"
	middlewares "apis/user-web/middleware"
)

func InitUserRouter(Router *route.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}
}
