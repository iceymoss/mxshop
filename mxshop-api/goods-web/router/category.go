package router

import (
	"mxshop-api/goods-web/api/category"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitCategory(router *gin.RouterGroup) {
	CategoryRouter := router.Group("categorys").Use(middlewares.Trace())
	{
		CategoryRouter.GET("", category.List)                                                        //商品分类列表
		CategoryRouter.GET("/:id", category.Detail)                                                  //获取商品分类详情
		CategoryRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdmin(), category.New)          //添加分类
		CategoryRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), category.Delete) //删除分类
		CategoryRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), category.Update)    //更新分类
	}
}
