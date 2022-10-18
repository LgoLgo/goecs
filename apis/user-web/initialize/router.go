package initialize

import (
	userRouter "E-commerce-system/apis/user-web/router"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func Routers(port int) *server.Hertz {
	Router := server.Default(
		server.WithHostPorts(fmt.Sprintf(":%d", port)),
	)

	ApiGroup := Router.Group("/v1")
	userRouter.InitUserRouter(ApiGroup)

	return Router
}
