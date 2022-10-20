package global

import (
	"mxshop_srvs/user_srv/config"

	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig  *config.NacosServer  = &config.NacosServer{}
)
