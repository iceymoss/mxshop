package initialize

import (
	"net/http"

	"mxshop-api/userop-web/middlewares"
	"mxshop-api/userop-web/router"

	"github.com/gin-gonic/gin"
)

//Routers 初始化及路由分发
func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨越
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("up")
	//分发路由
	ApiGroup = ApiGroup.Group("v1")

	router.InitMessageRouter(ApiGroup) //留言
	router.InitUserFavRouter(ApiGroup) //商品收藏
	router.InitAddressRouter(ApiGroup) //收货地址

	//健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})

	return Router
}
