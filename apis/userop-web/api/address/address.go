package address

import (
	"apis/userop-web/api"
	"apis/userop-web/forms"
	"apis/userop-web/global"
	"apis/userop-web/models"
	"apis/userop-web/proto/gen"
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"
)

func List(c context.Context, ctx *app.RequestContext) {
	request := &proto.AddressRequest{}

	claims, _ := ctx.Get("claims")
	currentUser := claims.(*models.CustomClaims)

	if currentUser.AuthorityId != 2 {
		userId, _ := ctx.Get("userId")
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.AddressClient.GetAddressList(c, request)
	if err != nil {
		zap.S().Errorw("Get address failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	reMap := utils.H{
		"total": rsp.Total,
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["user_id"] = value.UserId
		reMap["province"] = value.Province
		reMap["city"] = value.City
		reMap["district"] = value.District
		reMap["address"] = value.Address
		reMap["signer_name"] = value.SignerName
		reMap["signer_mobile"] = value.SignerMobile

		result = append(result, reMap)
	}

	reMap["data"] = result

	ctx.JSON(http.StatusOK, reMap)
}

func New(c context.Context, ctx *app.RequestContext) {
	addressForm := forms.AddressForm{}
	if err := ctx.BindAndValidate(&addressForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	rsp, err := global.AddressClient.CreateAddress(c, &proto.AddressRequest{
		UserId:       int32(userId.(uint)),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("Create address error")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.H{
		"id": rsp.Id,
	})
}

func Delete(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.AddressClient.DeleteAddress(c, &proto.AddressRequest{Id: int32(i)})
	if err != nil {
		zap.S().Errorw("Delete address error")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.H{
		"msg": "Delete successful",
	})
}

func Update(c context.Context, ctx *app.RequestContext) {
	addressForm := forms.AddressForm{}
	if err := ctx.BindAndValidate(&addressForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.AddressClient.UpdateAddress(c, &proto.AddressRequest{
		Id:           int32(i),
		Province:     addressForm.Province,
		City:         addressForm.City,
		District:     addressForm.District,
		Address:      addressForm.Address,
		SignerName:   addressForm.SignerName,
		SignerMobile: addressForm.SignerMobile,
	})
	if err != nil {
		zap.S().Errorw("Update address failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	ctx.JSON(http.StatusOK, utils.H{})
}
