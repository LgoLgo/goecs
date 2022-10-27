package api

import (
	"apis/user-web/forms"
	"apis/user-web/validator"
	"context"
	"fmt"
	"go.uber.org/zap"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"apis/user-web/global"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-redis/redis/v8"
)

func GenerateSmsCode(width int) string {
	// 生成width长度的短信验证码
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(c context.Context, ctx *app.RequestContext) {
	validator.ValidateMobile() // 手机号自定义表单验证设置
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.BindAndValidate(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	smsCode := GenerateSmsCode(6)
	zap.S().Infof("smsCode = %s", smsCode)

	// 将验证码保存起来 - redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)

	ctx.JSON(http.StatusOK, utils.H{
		"msg": "发送成功",
	})
}
