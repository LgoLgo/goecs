package order

import (
	"apis/order-web/api"
	"apis/order-web/forms"
	"apis/order-web/global"
	"apis/order-web/models"
	"apis/order-web/proto/gen"
	"context"
	"net/http"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
)

func List(c context.Context, ctx *app.RequestContext) {
	userId, _ := ctx.Get("userId")
	claims, _ := ctx.Get("claims")

	request := proto.OrderFilterRequest{}

	// Returns all orders if an admin user
	model := claims.(*models.CustomClaims)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	pages := ctx.DefaultQuery("p", "0")
	pagesInt, _ := strconv.Atoi(pages)
	request.Pages = int32(pagesInt)

	perNums := ctx.DefaultQuery("pnum", "0")
	perNumsInt, _ := strconv.Atoi(perNums)
	request.PagePerNums = int32(perNumsInt)

	request.Pages = int32(pagesInt)
	request.PagePerNums = int32(perNumsInt)

	rsp, err := global.OrderSrvClient.OrderList(c, &request)
	if err != nil {
		zap.S().Errorw("Failed to get order list")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	reMap := utils.H{
		"total": rsp.Total,
	}
	orderList := make([]interface{}, 0)

	for _, item := range rsp.Data {
		tmpMap := map[string]interface{}{}

		tmpMap["id"] = item.Id
		tmpMap["status"] = item.Status
		tmpMap["pay_type"] = item.PayType
		tmpMap["user"] = item.UserId
		tmpMap["post"] = item.Post
		tmpMap["total"] = item.Total
		tmpMap["address"] = item.Address
		tmpMap["name"] = item.Name
		tmpMap["mobile"] = item.Mobile
		tmpMap["order_sn"] = item.OrderSn
		tmpMap["id"] = item.Id
		tmpMap["add_time"] = item.AddTime

		orderList = append(orderList, tmpMap)
	}
	reMap["data"] = orderList
	ctx.JSON(http.StatusOK, reMap)
}

func New(c context.Context, ctx *app.RequestContext) {
	orderForm := forms.CreateOrderForm{}
	if err := ctx.BindAndValidate(&orderForm); err != nil {
		api.HandleValidatorError(ctx, err)
	}
	userId, _ := ctx.Get("userId")
	rsp, err := global.OrderSrvClient.CreateOrder(context.WithValue(c, "hertzContext", ctx), &proto.OrderRequest{
		UserId:  int32(userId.(uint)),
		Name:    orderForm.Name,
		Mobile:  orderForm.Mobile,
		Address: orderForm.Address,
		Post:    orderForm.Post,
	})
	if err != nil {
		zap.S().Errorw("New order failed")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}

	// Generate Alipay payment url
	client, err := alipay.New(global.ServerConfig.AliPayInfo.AppID, global.ServerConfig.AliPayInfo.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("Failed to instantiate Alipay")
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(global.ServerConfig.AliPayInfo.AliPublicKey)
	if err != nil {
		zap.S().Errorw("Failed to load Alipay public key")
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": err.Error(),
		})
		return
	}

	p := alipay.TradePagePay{}
	p.NotifyURL = global.ServerConfig.AliPayInfo.NotifyURL
	p.ReturnURL = global.ServerConfig.AliPayInfo.ReturnURL
	p.Subject = rsp.OrderSn
	p.OutTradeNo = rsp.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(rsp.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("Failed to generate payment url")
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, utils.H{
		"id":         rsp.Id,
		"alipay_url": url.String(),
	})
}

func Detail(c context.Context, ctx *app.RequestContext) {
	id := ctx.Param("id")
	userId, _ := ctx.Get("userId")
	i, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, utils.H{
			"msg": "url format error",
		})
		return
	}

	// Returns all orders if an admin user
	request := proto.OrderRequest{
		Id: int32(i),
	}
	claims, _ := ctx.Get("claims")
	model := claims.(*models.CustomClaims)
	if model.AuthorityId == 1 {
		request.UserId = int32(userId.(uint))
	}

	rsp, err := global.OrderSrvClient.OrderDetail(c, &request)
	if err != nil {
		zap.S().Errorw("Failed to get order details")
		api.HandleGRPCErrorToHTTP(err, ctx)
		return
	}
	reMap := utils.H{}
	reMap["id"] = rsp.OrderInfo.Id
	reMap["status"] = rsp.OrderInfo.Status
	reMap["user"] = rsp.OrderInfo.UserId
	reMap["post"] = rsp.OrderInfo.Post
	reMap["total"] = rsp.OrderInfo.Total
	reMap["address"] = rsp.OrderInfo.Address
	reMap["name"] = rsp.OrderInfo.Name
	reMap["mobile"] = rsp.OrderInfo.Mobile
	reMap["pay_type"] = rsp.OrderInfo.PayType
	reMap["order_sn"] = rsp.OrderInfo.OrderSn

	goodsList := make([]interface{}, 0)
	for _, item := range rsp.Goods {
		tmpMap := utils.H{
			"id":    item.GoodsId,
			"name":  item.GoodsName,
			"image": item.GoodsImage,
			"price": item.GoodsPrice,
			"nums":  item.Nums,
		}

		goodsList = append(goodsList, tmpMap)
	}
	reMap["goods"] = goodsList

	// Generate Alipay payment url
	client, err := alipay.New(global.ServerConfig.AliPayInfo.AppID, global.ServerConfig.AliPayInfo.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("Failed to instantiate Alipay")
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(global.ServerConfig.AliPayInfo.AliPublicKey)
	if err != nil {
		zap.S().Errorw("Failed to load Alipay public key")
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": err.Error(),
		})
		return
	}

	p := alipay.TradePagePay{}
	p.NotifyURL = global.ServerConfig.AliPayInfo.NotifyURL
	p.ReturnURL = global.ServerConfig.AliPayInfo.ReturnURL
	p.Subject = rsp.OrderInfo.OrderSn
	p.OutTradeNo = rsp.OrderInfo.OrderSn
	p.TotalAmount = strconv.FormatFloat(float64(rsp.OrderInfo.Total), 'f', 2, 64)
	p.ProductCode = "FAST_INSTANT_TRADE_PAY"

	url, err := client.TradePagePay(p)
	if err != nil {
		zap.S().Errorw("Failed to generate payment url")
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": err.Error(),
		})
		return
	}
	reMap["alipay_url"] = url.String()

	ctx.JSON(http.StatusOK, reMap)
}
