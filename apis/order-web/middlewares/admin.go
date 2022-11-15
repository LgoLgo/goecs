package middlewares

import (
	"apis/order-web/models"
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"net/http"
)

func IsAdminAuth() app.HandlerFunc {
	//将一些共用的代码抽出来然后共用 - 版本管理
	//如果不抽出来
	return func(ctx context.Context, c *app.RequestContext) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			c.JSON(http.StatusForbidden, utils.H{
				"msg": "无权限",
			})
			c.Abort()
			return
		}
		c.Next(ctx)
	}

}
