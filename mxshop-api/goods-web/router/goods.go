package router

import (
	"mxshop-api/goods-web/api/goods"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitGoodsRouter(router *gin.RouterGroup) {
	GoodsRouter := router.Group("goods")
	{
		//中间件的参数位置需要按业务需求而定
		GoodsRouter.GET("", goods.List)                                                             //商品列表
		GoodsRouter.GET("/:id", goods.Detail)                                                       //获取商品详情
		GoodsRouter.GET("/:id/stocks", goods.Stocks)                                                //获取库存
		GoodsRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.New)               //添加商品
		GoodsRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.Delete)      //删除商品
		GoodsRouter.PATCH("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.UpdateStatus) //更新商品状态
		GoodsRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), goods.Update)         //更新商品信息
	}
}
