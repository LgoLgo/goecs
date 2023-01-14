package api

import (
	"apis/user-web/forms"
	"apis/user-web/global"
	"apis/user-web/validator"
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func GenerateSmsCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		_, _ = fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(_ context.Context, ctx *app.RequestContext) {
	validator.ValidateMobile()
	sendSmsForm := forms.SendSmsForm{}
	if err := ctx.BindAndValidate(&sendSmsForm); err != nil {
		HandleValidatorError(ctx, err)
		return
	}
	smsCode := GenerateSmsCode(6)
	zap.S().Infof("smsCode = %s", smsCode)

	// Save your smsCode - redis
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.RedisInfo.Host, global.ServerConfig.RedisInfo.Port),
	})
	rdb.Set(context.Background(), sendSmsForm.Mobile, smsCode, time.Duration(global.ServerConfig.RedisInfo.Expire)*time.Second)

	ctx.JSON(http.StatusOK, utils.H{
		"msg": "Send successful",
	})
}
