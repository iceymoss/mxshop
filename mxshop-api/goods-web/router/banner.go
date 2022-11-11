package router

import (
	"mxshop-api/goods-web/api/banner"
	"mxshop-api/goods-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitBanner(router *gin.RouterGroup) {
	BannerRouter := router.Group("banners").Use(middlewares.Trace())
	{
		BannerRouter.GET("", banner.BannerList)                                                        //获取轮播图列表
		BannerRouter.POST("", middlewares.JWTAuth(), middlewares.IsAdmin(), banner.NewBanner)          //添加轮播图
		BannerRouter.DELETE("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), banner.DeleteBanner) //删除轮播图
		BannerRouter.PUT("/:id", middlewares.JWTAuth(), middlewares.IsAdmin(), banner.UpdateBanner)    //更新轮播图
	}
}
