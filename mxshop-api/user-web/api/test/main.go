package main

import (
	"fmt"
	"math/rand"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"
	"strings"
	"time"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
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

func main() {
	initialize.InitConfig()

	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", global.ServerConfig.AliSms.Apikey, global.ServerConfig.AliSms.ApiSecret)
	if err != nil {
		panic(err)
	}
	fmt.Println("配置：", global.ServerConfig.AliSms.Apikey, global.ServerConfig.AliSms.ApiSecret)

	width := global.ServerConfig.Verify.Width
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Scheme = "https" // https | http
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = "SendSms"
	request.QueryParams["RegionId"] = "cn-beijing"
	request.QueryParams["PhoneNumbers"] = "17585610985"                                //手机号
	request.QueryParams["SignName"] = global.ServerConfig.Params.SignName              //阿里云验证过的项目名 自己设置
	request.QueryParams["TemplateCode"] = global.ServerConfig.Params.TemplateCode      //阿里云的短信模板号 自己设置
	request.QueryParams["TemplateParam"] = "{\"code\":" + GenerateSmsCode(width) + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
	response, err := client.ProcessCommonRequest(request)
	fmt.Print(client.DoAction(request, response))
	//  fmt.Print(response)
	if err != nil {
		fmt.Print(err.Error())
	}
	fmt.Printf("response is %#v\n", response)
	//json数据解析
}
