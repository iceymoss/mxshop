package global

import (
	"mxshop-api/oss-web/config"

	ut "github.com/go-playground/universal-translator"
)

var (
	Trans ut.Translator

	ServerConfig *config.ServerConfig = &config.ServerConfig{}

	NacosConfig *config.NacosConfig = &config.NacosConfig{}
)
