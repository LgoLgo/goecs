package middlewares

import (
	"context"

	"apis/user-web/models"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func IsAdminAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, utils.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
