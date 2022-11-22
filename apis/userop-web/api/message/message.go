package message

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"

	"apis/userop-web/api"
	"apis/userop-web/forms"
	"apis/userop-web/global"
	"apis/userop-web/models"
	"apis/userop-web/proto/gen"
)

func List(c context.Context, ctx *app.RequestContext) {
	request := &proto.MessageRequest{}

	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")
	model := claims.(*models.CustomClaims)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.MessageClient.MessageList(c, request)
	if err != nil {
		zap.S().Errorw("Get message error")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": rsp.Total,
	}
	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["user_id"] = value.UserId
		reMap["type"] = value.MessageType
		reMap["subject"] = value.Subject
		reMap["message"] = value.Message
		reMap["file"] = value.File

		result = append(result, reMap)
	}
	reMap["data"] = result

	ctx.JSON(http.StatusOK, reMap)
}

func New(c context.Context, ctx *app.RequestContext) {
	userId, _ := ctx.Get("userId")

	messageForm := forms.MessageForm{}
	if err := ctx.BindAndValidate(&messageForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	rsp, err := global.MessageClient.CreateMessage(c, &proto.MessageRequest{
		UserId:      int32(userId.(uint)),
		MessageType: messageForm.MessageType,
		Subject:     messageForm.Subject,
		Message:     messageForm.Message,
		File:        messageForm.File,
	})

	if err != nil {
		zap.S().Errorw("Add message failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.H{
		"id": rsp.Id,
	})
}
