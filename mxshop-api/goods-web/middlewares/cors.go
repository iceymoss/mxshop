package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Cors 解决浏览器跨越问题，后端解决方法
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		Method := c.Request.Method

		//返回给浏览器，告诉浏览器header可以填什么内容
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token, x-token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, DELETE, PATCH, PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		if Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
	}
}
