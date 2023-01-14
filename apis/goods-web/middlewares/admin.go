package middlewares

import (
	"apis/goods-web/models"
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func IsAdminAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		claims, _ := c.Get("claims")
		currentUser := claims.(*models.CustomClaims)

		if currentUser.AuthorityId != 2 {
			c.JSON(http.StatusForbidden, utils.H{
				"msg": "No permission",
			})
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
