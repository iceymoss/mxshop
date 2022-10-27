package global

import (
	"mxshop-api/user-web/config"
	"mxshop-api/user-web/proto"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans        ut.Translator                                 //声明一个全局翻译器
	ServerConfig *config.ServerConfig = &config.ServerConfig{} //声明配置信息
	UserClient   proto.UserClient                              //grpc Client
	NacosConfig  *config.NacosConfig  = &config.NacosConfig{}
)
