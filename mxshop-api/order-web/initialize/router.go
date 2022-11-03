package initialize

import (
	"net/http"

	"mxshop-api/order-web/middlewares"
	"mxshop-api/order-web/router"

	"github.com/gin-gonic/gin"
)

//Routers 初始化及路由分发
func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨越
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("o")
	//分发路由
	ApiGroup = ApiGroup.Group("v1")

	router.InitShopCartRouter(ApiGroup) //购物车
	router.InitOrderRouter(ApiGroup)    //订单

	//健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})

	return Router
}
