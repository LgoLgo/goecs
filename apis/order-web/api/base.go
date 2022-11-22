package api

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

func HandleGRPCErrorToHTTP(err error, c *app.RequestContext) {
	// Convert gRPC code to HTTP status code
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.NotFound:
				c.JSON(http.StatusNotFound, utils.H{
					"msg": e.Message(),
				})
			case codes.Internal:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg:": "Internal error",
				})
			case codes.InvalidArgument:
				c.JSON(http.StatusBadRequest, utils.H{
					"msg": "Parameter error",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg": "商品服务不可用",
				})
			default:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg": e.Code(),
				})
			}
			return
		}
	}
}

func HandleValidatorError(c *app.RequestContext, err error) {
	c.JSON(http.StatusOK, utils.H{
		"msg": err.Error(),
	})
	return
}
