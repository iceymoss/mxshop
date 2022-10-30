package initialize

import (
	"mxshop-api/goods-web/middlewares"
	"mxshop-api/goods-web/router"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Routers 初始化及路由分发
func Routers() *gin.Engine {
	Router := gin.Default()

	//配置跨越
	Router.Use(middlewares.Cors())

	ApiGroup := Router.Group("g")
	//分发路由
	ApiGroup = ApiGroup.Group("v1")

	router.InitGoodsRouter(ApiGroup) //商品
	router.InitCategory(ApiGroup)    //商品分类
	router.InitBanner(ApiGroup)      //商品轮播图
	router.InitBrands(ApiGroup)      //品牌，分类-品牌

	//健康检查
	Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "health",
		})
	})

	return Router
}
