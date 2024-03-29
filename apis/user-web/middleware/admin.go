package middlewares

import (
	"apis/user-web/models"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func IsAdminAuth() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, utils.H{
				"msg": "You are not admin.",
			})
			ctx.Abort()
			return
		}
		ctx.Next(c)
	}
}
