package router

import (
	"mxshop-api/userop-web/api/address"
	"mxshop-api/userop-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitAddressRouter(router *gin.RouterGroup) {
	AddressRouter := router.Group("address").Use(middlewares.JWTAuth())
	{
		AddressRouter.GET("", address.List)          //获取收货地址列表
		AddressRouter.POST("", address.New)          //新建收货地址
		AddressRouter.PUT("/:id", address.Update)    //更新收货地址
		AddressRouter.DELETE("/:id", address.Delete) //删除收货地址
	}
}
