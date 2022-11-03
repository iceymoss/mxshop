package middlewares

import (
	"mxshop-api/order-web/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

//IsAdmin IsAdmin 识别用户身份类型
func IsAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//将用户信息拿出
		claims, _ := ctx.Get("claims")
		currentUser := claims.(*models.CustomClaims)
		if currentUser.AuthorityId != 2 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"msg": "无权限",
			})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
