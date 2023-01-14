package initialize

import (
	"apis/userop-web/middlewares"
	"apis/userop-web/router"
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Routers(port int) *server.Hertz {
	Router := server.Default(server.WithHostPorts(fmt.Sprintf(":%d", port)))
	Router.GET("/health", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, utils.H{
			"code":    http.StatusOK,
			"success": true,
		})
	})

	// Configure cross-domain.
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("/up/v1")
	router.InitUserFavRouter(ApiGroup)
	router.InitMessageRouter(ApiGroup)
	router.InitAddressRouter(ApiGroup)

	return Router
}
