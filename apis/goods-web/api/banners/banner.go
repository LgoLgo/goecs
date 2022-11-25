package banners

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang/protobuf/ptypes/empty"

	"apis/goods-web/api"
	"apis/goods-web/forms"
	"apis/goods-web/global"
	"apis/goods-web/proto/gen"
)

func List(ctx context.Context, c *app.RequestContext) {
	rsp, err := global.GoodsSrvClient.BannerList(ctx, &empty.Empty{})
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, c)
		return
	}

	result := make([]interface{}, 0)
	for _, value := range rsp.Data {
		reMap := make(map[string]interface{})
		reMap["id"] = value.Id
		reMap["index"] = value.Index
		reMap["image"] = value.Image
		reMap["url"] = value.Url

		result = append(result, reMap)
	}

	c.JSON(http.StatusOK, result)
}

func New(c context.Context, ctx *app.RequestContext) {
	bannerForm := forms.BannerForm{}
	if err := ctx.BindAndValidate(&bannerForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	rsp, err := global.GoodsSrvClient.CreateBanner(c, &proto.BannerRequest{
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
		Image: bannerForm.Image,
	})
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	response := make(map[string]interface{})
	response["id"] = rsp.Id
	response["index"] = rsp.Index
	response["url"] = rsp.Url
	response["image"] = rsp.Image

	ctx.JSON(http.StatusOK, response)
}

func Update(c context.Context, ctx *app.RequestContext) {
	bannerForm := forms.BannerForm{}
	if err := ctx.BindAndValidate(&bannerForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	_, err = global.GoodsSrvClient.UpdateBanner(c, &proto.BannerRequest{
		Id:    int32(i),
		Index: int32(bannerForm.Index),
		Url:   bannerForm.Url,
	})
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}

func Delete(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	_, err = global.GoodsSrvClient.DeleteBanner(c, &proto.BannerRequest{Id: int32(i)})
	if err != nil {
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, "")
}
