package initialize

import (
	"mxshop-api/user-web/middlewares"
	"mxshop-api/user-web/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Routers 初始化及路由分发
func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨越
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("u")
	//分发路由
	ApiGroup = ApiGroup.Group("v1")

	router.InitUserRouter(ApiGroup)

	router.InitBaseRouter(ApiGroup)

	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})

	return Router
}
