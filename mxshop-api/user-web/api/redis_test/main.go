package main

import (
	"context"
	"fmt"
	"log"
	"mxshop-api/user-web/global"
	"mxshop-api/user-web/initialize"

	"github.com/go-redis/redis/v8"
)

func main() {
	//验证码验证
	initialize.InitConfig()
	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", global.ServerConfig.Redis.Host, global.ServerConfig.Redis.Port),
	})
	mobile := "17585610985"
	code := "67122"
	rsp := rdb.Get(context.Background(), mobile)
	value, err := rsp.Result()
	if err != nil {
		log.Fatal("获取验证码失败", err)
		return
	}

	if code != value {
		log.Fatal("验证码错误")
		return
	}
	fmt.Println("验证码验证通过")
	fmt.Println(value)

}
