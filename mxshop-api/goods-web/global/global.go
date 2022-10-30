package global

import (
	"mxshop-api/goods-web/config"
	"mxshop-api/goods-web/proto"

	ut "github.com/go-playground/universal-translator"
)

//需要的全局变量
var (
	Trans          ut.Translator                                 //声明一个全局翻译器
	ServerConfig   *config.ServerConfig = &config.ServerConfig{} //声明配置信息
	GoodsSrvClient proto.GoodsClient                             //grpc Client
	NacosConfig    *config.NacosConfig  = &config.NacosConfig{}
)
