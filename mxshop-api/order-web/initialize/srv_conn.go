package initialize

import (
	"fmt"

	"mxshop-api/order-web/global"
	"mxshop-api/order-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//InitSrvConn 连接到consul注册中心并对服务做负载均衡
func InitSrvConn() {
	consul := global.ServerConfig.ConsulInfo
	//连接订单服务
	Orderconn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.OrderSerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【订单服务失败】", err.Error())
		return
	}
	OrderClient := proto.NewOrderClient(Orderconn)
	global.OrderSrvClient = OrderClient

	//连接商品服务
	Goodsconn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.GoodsSerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【商品服务失败】", err.Error())
		return
	}
	goodsClient := proto.NewGoodsClient(Goodsconn)
	global.GoodsSrvClient = goodsClient

	Inventoryconn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.InventorySerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【库存服务】失败", err.Error())
		return
	}
	InventoryClient := proto.NewInventoryClient(Inventoryconn)
	global.InventorySrvClient = InventoryClient
}
