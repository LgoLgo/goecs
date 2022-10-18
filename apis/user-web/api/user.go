package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	proto "E-commerce-system/apis/user-web/proto/gen"
)

func HandleGRPCErrorToHTTP(err error, c *app.RequestContext) {
	// 将 gRPC 的 code 转换成 HTTP 的状态码
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
					"msg": "Argument error",
				})
			case codes.Unavailable:
				c.JSON(http.StatusInternalServerError, utils.H{
					"msg": "Server error",
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

func GetUserList(ctx context.Context, c *app.RequestContext) {
	ip := "127.0.0.1"
	port := 50051

	// 拨号连接用户 gRPC 服务
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", ip, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] connected error",
			"msg", err.Error(),
		)
	}
	// 调用接口
	userSrcClient := proto.NewUserClient(userConn)

	rsp, err := userSrcClient.GetUserList(ctx, &proto.PageInfo{
		Pn:    0,
		PSize: 0,
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] query user list error.")
		HandleGRPCErrorToHTTP(err, c)
		return
	}

	zap.S().Debug("get user list")
}
