package api

import (
	"E-commerce-system/apis/user-web/global"
	"E-commerce-system/apis/user-web/global/response"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

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
	userConn, err := grpc.Dial(fmt.Sprintf("%s:%d", global.ServerConfig.UserSrvInfo.Host, global.ServerConfig.UserSrvInfo.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		zap.S().Errorw("[GetUserList] connected error",
			"msg", err.Error(),
		)
	}
	// 调用接口
	userSrcClient := proto.NewUserClient(userConn)

	pn := c.DefaultQuery("pn", "0")
	pnInt, _ := strconv.Atoi(pn)
	pSize := c.DefaultQuery("psize", "10")
	pSizeInt, _ := strconv.Atoi(pSize)
	rsp, err := userSrcClient.GetUserList(ctx, &proto.PageInfo{
		Pn:    uint32(pnInt),
		PSize: uint32(pSizeInt),
	})
	if err != nil {
		zap.S().Errorw("[GetUserList] query user list error.")
		HandleGRPCErrorToHTTP(err, c)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		user := response.UserResponse{
			Id:       value.Id,
			NickName: value.NickName,
			Birthday: response.JsonTime(time.Unix(int64(value.BirthDay), 0)),
			Gender:   value.Gender,
			Mobile:   value.Mobile,
		}
		result = append(result, user)
	}
	c.JSON(http.StatusOK, result)
}
