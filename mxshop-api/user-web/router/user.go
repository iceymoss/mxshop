package router

import (
	"mxshop-api/user-web/api"
	"mxshop-api/user-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.RouterGroup) {
	UserRouter := router.Group("user")
	{
		//中间件的参数位置需要按业务需求而定
		UserRouter.GET("list", middlewares.JWTAuth(), middlewares.IsAdmin(), api.GetUserList)
		UserRouter.POST("login", api.PassWordLogin)
		UserRouter.POST("register", api.Register)
	}

	StoreRouter := router.Group("store")
	StoreRouter.Use(middlewares.JWTAuth())
	{
		StoreRouter.GET("storelist", api.GetStoreList)
	}
}
