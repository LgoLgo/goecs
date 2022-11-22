package api

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

func GetCaptcha(_ context.Context, ctx *app.RequestContext) {
	driver := base64Captcha.NewDriverDigit(80, 240, 5, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	id, b64s, err := cp.Generate()
	if err != nil {
		zap.S().Errorf("Generate captcha error: %s", err.Error())
		ctx.JSON(http.StatusInternalServerError, utils.H{
			"msg": "Generate captcha error",
		})
		return
	}
	ctx.JSON(http.StatusOK, utils.H{
		"captchaId": id,
		"picPath":   b64s,
	})
}
