package router

import (
	"E-commerce-system/apis/user-web/api"
	"github.com/cloudwego/hertz/pkg/route"
)

func InitUserRouter(Router *route.RouterGroup) {
	UserRouter := Router.Group("user")

	UserRouter.GET("list", api.GetUserList)
}
