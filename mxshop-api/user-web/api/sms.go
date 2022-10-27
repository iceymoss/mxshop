package api

import (
	"context"
	"fmt"
	"math/rand"
	"mxshop-api/user-web/forms"
	"mxshop-api/user-web/global"
	"net/http"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

//GenerateSmsCode 生成width长度的验证码
func GenerateSmsCode(width int) string {
	number := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(number)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", number[rand.Intn(r)])
	}
	return sb.String()
}

func SendSms(c *gin.Context) {
	//表单验证
	SendSmsForm := forms.SendSmsForm{}
	if err := c.ShouldBind(&SendSmsForm); err != nil {
		HandleValidatorErr(c, err)
		return
	}

	//配置阿里云accessKey
	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.ServerConfig.AliSms.Apikey, global.ServerConfig.AliSms.ApiSecret)
	if err != nil {
		panic(err)
	}

	mobile := SendSmsForm.Mobile
	width := global.ServerConfig.Verify.Width
	smsCode := GenerateSmsCode(width)
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = mobile                                  //手机号
	request.QueryParams["SignName"] = global.ServerConfig.Params.SignName         //阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = global.ServerConfig.Params.TemplateCode //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + smsCode + "}"           //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	response, err := client.ProcessCommonRequest(request)
	fmt.Print(client.DoAction(request, response))
	//  fmt.Print(response)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)

	//连接redis服务器
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.Redis.Host, global.ServerConfig.Redis.Port),
	})

	//写入数据库
	rdb.Set(context.Background(), mobile, smsCode, time.Second*time.Duration(global.ServerConfig.Redis.Expir))

	c.JSON(http.StatusOK, gin.H{
		"msg": "发送成功",
	})

}
