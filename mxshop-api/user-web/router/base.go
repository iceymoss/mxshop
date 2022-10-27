package router

import (
	"mxshop-api/user-web/api"

	"github.com/gin-gonic/gin"
)

//InitBaseRouter 注册图片验证码路由
func InitBaseRouter(router *gin.RouterGroup) {
	//设置路由
	BaseRouter := router.Group("base")
	{
		BaseRouter.GET("captcha", api.GetChaptcha)
		BaseRouter.POST("send_sms", api.SendSms)
	}
}
