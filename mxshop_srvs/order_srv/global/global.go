package global

import (
	"mxshop_srvs/order_srv/config"
	"mxshop_srvs/order_srv/proto"

	"github.com/apache/rocketmq-client-go/v2"

	"github.com/apache/rocketmq-client-go/v2/producer"

	"github.com/go-redsync/redsync/v4"
	"gorm.io/gorm"
)

var (
	DB                 *gorm.DB
	ServerConfig       *config.ServerConfig = &config.ServerConfig{}
	NacosConfig        *config.NacosServer  = &config.NacosServer{}
	GoodsSrvClient     proto.GoodsClient
	InventorySrvClient proto.InventoryClient
	Rs                 *redsync.Redsync

	//tset
	GroupInventory producer.Option
	GroupOrder     producer.Option
	MQOrder        rocketmq.Producer
	MQInventory    rocketmq.Producer
)
