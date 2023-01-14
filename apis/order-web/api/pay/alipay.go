package pay

import (
	"apis/order-web/global"
	"apis/order-web/proto/gen"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
)

func Notify(c context.Context, ctx *app.RequestContext) {
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
	req, err := adaptor.GetCompatRequest(&ctx.Request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.H{})
		return
	}
	noti, err := client.GetTradeNotification(req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.H{})
		return
	}

	_, err = global.OrderSrvClient.UpdateOrderStatus(c, &proto.OrderStatus{
		OrderSn: noti.OutTradeNo,
		Status:  string(noti.TradeStatus),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.H{})
		return
	}
	ctx.String(http.StatusOK, "success")
}
