package pay

import (
	"apis/order-web/proto/gen"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/adaptor"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/smartwalle/alipay/v3"
	"go.uber.org/zap"
	"net/http"

	"apis/order-web/global"
)

func Notify(c context.Context, ctx *app.RequestContext) {
	//支付宝回调通知
	client, err := alipay.New(global.ServerConfig.AliPayInfo.AppID, global.ServerConfig.AliPayInfo.PrivateKey, false)
	if err != nil {
		zap.S().Errorw("实例化支付宝失败")
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": err.Error(),
		})
		return
	}
	err = client.LoadAliPayPublicKey(global.ServerConfig.AliPayInfo.AliPublicKey)
	if err != nil {
		zap.S().Errorw("加载支付宝的公钥失败")
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

	_, err = global.OrderSrvClient.UpdateOrderStatus(context.Background(), &proto.OrderStatus{
		OrderSn: noti.OutTradeNo,
		Status:  string(noti.TradeStatus),
	})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, utils.H{})
		return
	}
	ctx.String(http.StatusOK, "success")
}
