package global

import (
	"mxshop_srvs/inventory_srv/config"

	"github.com/go-redsync/redsync/v4"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig *config.ServerConfig = &config.ServerConfig{}
	NacosConfig  *config.NacosServer  = &config.NacosServer{}
	Rs           *redsync.Redsync
)
