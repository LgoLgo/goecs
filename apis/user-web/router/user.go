package router

import (
	"apis/user-web/api"
	middlewares "apis/user-web/middleware"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitUserRouter(Router *route.RouterGroup) {
	UserRouter := Router.Group("user")
	{
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdminAuth(), api.GetUserList)
		UserRouter.POST("pwd_login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}
}
