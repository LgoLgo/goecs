package initialize

import (
	middlewares "apis/user-web/middleware"
	userRouter "apis/user-web/router"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Routers(port int) *server.Hertz {
	Router := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%d", port)),
	)

	//配置跨域
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/v1")
	userRouter.InitUserRouter(ApiGroup)
	userRouter.InitBaseRouter(ApiGroup)

	return Router
}
