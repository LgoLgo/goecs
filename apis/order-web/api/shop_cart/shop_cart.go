package shop_cart

import (
	"apis/order-web/api"
	"apis/order-web/forms"
	"apis/order-web/proto/gen"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"
	"net/http"
	"strconv"

	"apis/order-web/global"
)

func List(c context.Context, ctx *app.RequestContext) {
	//获取购物车商品
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CartItemList(c, &proto.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("[List] 查询 【购物车列表】失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ids := make([]int32, 0)
	for _, item := range rsp.Data {
		ids = append(ids, item.GoodsId)
	}
	if len(ids) == 0 {
		ctx.JSON(http.StatusOK, utils.H{
			"total": 0,
		})
		return
	}

	//请求商品服务获取商品信息
	goodsRsp, err := global.GoodsSrvClient.BatchGetGoods(c, &proto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		zap.S().Errorw("[List] 批量查询【商品列表】失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	reMap := utils.H{
		"total": rsp.Total,
	}

	/*
		{
			"total":12,
			"data":[
				{
					"id":1,
					"goods_id":421,
					"goods_name":421,
					"goods_price":421,
					"goods_image":421,
					"nums":421,
					"checked":421,
				}
			]
		}
	*/
	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Data {
		for _, good := range goodsRsp.Data {
			if good.Id == item.GoodsId {
				tmpMap := map[string]interface{}{}
				tmpMap["id"] = item.Id
				tmpMap["goods_id"] = item.GoodsId
				tmpMap["good_name"] = good.Name
				tmpMap["good_image"] = good.GoodsFrontImage
				tmpMap["good_price"] = good.ShopPrice
				tmpMap["nums"] = item.Nums
				tmpMap["checked"] = item.Checked

				goodsList = append(goodsList, tmpMap)
			}
		}
	}
	reMap["data"] = goodsList
	ctx.JSON(http.StatusOK, reMap)
}

func New(c context.Context, ctx *app.RequestContext) {
	//添加商品到购物车
	itemForm := forms.ShopCartItemForm{}
	if err := ctx.BindAndValidate(&itemForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	//为了严谨性，添加商品到购物车之前，记得检查一下商品是否存在
	_, err := global.GoodsSrvClient.GetGoodsDetail(c, &proto.GoodInfoRequest{
		Id: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[List] 查询【商品信息】失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	//如果现在添加到购物车的数量和库存的数量不一致
	invRsp, err := global.InventorySrvClient.InvDetail(c, &proto.GoodsInvInfo{
		GoodsId: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[List] 查询【库存信息】失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	if invRsp.Num < itemForm.Nums {
		ctx.JSON(http.StatusBadRequest, utils.H{
			"nums": "库存不足",
		})
		return
	}

	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CreateCartItem(c, &proto.CartItemRequest{
		GoodsId: itemForm.GoodsId,
		UserId:  int32(userId.(uint)),
		Nums:    itemForm.Nums,
	})

	if err != nil {
		zap.S().Errorw("添加到购物车失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.H{
		"id": rsp.Id,
	})
}

func Update(c context.Context, ctx *app.RequestContext) {
	// o/v1/421
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.H{
			"msg": "url格式出错",
		})
		return
	}

	itemForm := forms.ShopCartItemUpdateForm{}
	if err := ctx.BindAndValidate(&itemForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	userId, _ := ctx.Get("userId")
	request := proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
		Nums:    itemForm.Nums,
		Checked: false,
	}
	if itemForm.Checked != nil {
		request.Checked = *itemForm.Checked
	}

	_, err = global.OrderSrvClient.UpdateCartItem(c, &request)
	if err != nil {
		zap.S().Errorw("更新购物车记录失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}

func Delete(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.H{
			"msg": "url格式出错",
		})
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.OrderSrvClient.DeleteCartItem(c, &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		zap.S().Errorw("删除购物车记录失败")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
