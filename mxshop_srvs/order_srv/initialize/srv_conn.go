package initialize

import (
	"fmt"
	"mxshop_srvs/order_srv/global"
	"mxshop_srvs/order_srv/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//InitSrvConn 连接到consul注册中心并对服务做负载均衡
func InitSrvConn() {
	//连接商品服务
	consul := global.ServerConfig.ConsulInfo
	GoodsConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.GoodsSerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[ InitSrvConn] 连接 【商品服务】失败", err.Error())
		return
	}
	goodsClient := proto.NewGoodsClient(GoodsConn)
	global.GoodsSrvClient = goodsClient

	//连接库存服务
	InvConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.InventorySerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[ InitSrvConn] 连接 【库存服务】失败", err.Error())
		return
	}
	inventoryClient := proto.NewInventoryClient(InvConn)
	global.InventorySrvClient = inventoryClient
}
