package user_fav

import (
	"apis/userop-web/api"
	"apis/userop-web/forms"
	"apis/userop-web/global"
	"apis/userop-web/proto/gen"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

func List(c context.Context, ctx *app.RequestContext) {
	userId, _ := ctx.Get("userId")
	userFavRsp, err := global.UserFavClient.GetFavList(context.Background(), &proto.UserFavRequest{
		UserId: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("获取收藏列表失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ids := make([]int32, 0)
	for _, item := range userFavRsp.Data {
		ids = append(ids, item.GoodsId)
	}

	if len(ids) == 0 {
		ctx.JSON(http.StatusOK, utils.H{
			"total": 0,
		})
		return
	}

	//请求商品服务
	goods, err := global.GoodsSrvClient.BatchGetGoods(context.Background(), &proto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		zap.S().Errorw("[List] 批量查询【商品列表】失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	reMap := map[string]interface{}{
		"total": userFavRsp.Total,
	}

	goodsList := make([]interface{}, 0)
	for _, item := range userFavRsp.Data {
		data := utils.H{
			"id": item.GoodsId,
		}

		for _, good := range goods.Data {
			if item.GoodsId == good.Id {
				data["name"] = good.Name
				data["shop_price"] = good.ShopPrice
			}
		}

		goodsList = append(goodsList, data)
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func New(c context.Context, ctx *app.RequestContext) {
	userFavForm := forms.UserFavForm{}
	if err := ctx.BindAndValidate(&userFavForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	//缺少一步， 这个时候应该去商品服务查询一下这个是否存在
	userId, _ := ctx.Get("userId")
	_, err := global.UserFavClient.AddUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: userFavForm.GoodsId,
	})

	if err != nil {
		zap.S().Errorw("添加收藏记录失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.H{})
}

func Delete(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	i, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.UserFavClient.DeleteUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		zap.S().Errorw("删除收藏记录失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.H{
		"msg": "删除成功",
	})
}

func Detail(c context.Context, ctx *app.RequestContext) {
	goodsId := ctx.Param("id")
	goodsIdInt, err := strconv.ParseInt(goodsId, 10, 32)
	if err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	userId, _ := ctx.Get("userId")
	_, err = global.UserFavClient.GetUserFavDetail(context.Background(), &proto.UserFavRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(goodsIdInt),
	})
	if err != nil {
		zap.S().Errorw("查询收藏状态失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.Status(http.StatusOK)
}
