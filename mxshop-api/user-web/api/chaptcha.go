package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

//图片验证对象
var store = base64Captcha.DefaultMemStore

//GetChaptcha 生成图片验证码
func GetChaptcha(ctx *gin.Context) {
	digit := base64Captcha.NewDriverDigit(80, 240, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(digit, store)
	id, base64, err := cp.Generate()
	if err != nil {
		zap.S().Info("生成验证码失败", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"msg": "生成图片验证码失败",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"captcha":    base64,
		"captcha_id": id,
	})
}
