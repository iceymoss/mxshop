package global

import (
	"mxshop_srvs/goods_srv/config"

	"github.com/olivere/elastic/v7"
	"gorm.io/gorm"
)

var (
	DB           *gorm.DB
	ServerConfig config.ServerConfig
	NacosConfig  *config.NacosServer = &config.NacosServer{}
	EsClient     *elastic.Client
)
