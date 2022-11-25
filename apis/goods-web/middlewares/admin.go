package middlewares

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	"apis/goods-web/models"
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
