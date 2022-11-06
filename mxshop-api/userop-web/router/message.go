package router

import (
	"mxshop-api/userop-web/api/message"
	"mxshop-api/userop-web/middlewares"

	"github.com/gin-gonic/gin"
)

func InitMessageRouter(router *gin.RouterGroup) {
	MessageRouter := router.Group("message").Use(middlewares.JWTAuth())
	{
		//中间件的参数位置需要按业务需求而定
		MessageRouter.GET("", message.List) //留言列表
		MessageRouter.POST("", message.New) //新建留言
	}
}
