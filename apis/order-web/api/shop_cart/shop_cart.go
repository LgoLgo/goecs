package shop_cart

import (
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"go.uber.org/zap"

	"apis/order-web/api"
	"apis/order-web/forms"
	"apis/order-web/global"
	"apis/order-web/proto/gen"
)

func List(c context.Context, ctx *app.RequestContext) {
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CartItemList(c, &proto.UserInfo{
		Id: int32(userId.(uint)),
	})
	if err != nil {
		zap.S().Errorw("[List] Query [Shopping Cart List] failed")
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

	// Request product service to obtain product information
	goodsRsp, err := global.GoodsSrvClient.BatchGetGoods(c, &proto.BatchGoodsIdInfo{
		Id: ids,
	})
	if err != nil {
		zap.S().Errorw("[List] Batch query [Product List] failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	reMap := utils.H{
		"total": rsp.Total,
	}
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
	// Add item to cart
	itemForm := forms.ShopCartItemForm{}
	if err := ctx.BindAndValidate(&itemForm); err != nil {
		api.HandleValidatorError(ctx, err)
		return
	}

	// For the sake of rigor, remember to check if the item exists before adding it to the cart
	_, err := global.GoodsSrvClient.GetGoodsDetail(c, &proto.GoodInfoRequest{
		Id: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[List] Query [Product Information] failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	// If the quantity added to the cart now does not match the quantity in stock
	invRsp, err := global.InventorySrvClient.InvDetail(c, &proto.GoodsInvInfo{
		GoodsId: itemForm.GoodsId,
	})
	if err != nil {
		zap.S().Errorw("[List] Query [Inventory Information] failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	if invRsp.Num < itemForm.Nums {
		ctx.JSON(http.StatusBadRequest, utils.H{
			"nums": "Inventory shortage",
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
		zap.S().Errorw("Add to cart failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	ctx.JSON(http.StatusOK, utils.H{
		"id": rsp.Id,
	})
}

func Update(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.H{
			"msg": "url format error",
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
		zap.S().Errorw("Failed to update cart record")
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
			"msg": "url format error",
		})
		return
	}

	userId, _ := ctx.Get("userId")
	_, err = global.OrderSrvClient.DeleteCartItem(c, &proto.CartItemRequest{
		UserId:  int32(userId.(uint)),
		GoodsId: int32(i),
	})
	if err != nil {
		zap.S().Errorw("Failed to delete cart record")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	ctx.Status(http.StatusOK)
}
