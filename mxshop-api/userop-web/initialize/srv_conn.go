package initialize

import (
	"fmt"

	"mxshop-api/userop-web/global"
	"mxshop-api/userop-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//InitSrvConn 连接到consul注册中心并对服务做负载均衡
func InitSrvConn() {
	consul := global.ServerConfig.ConsulInfo
	//连接用户操作服务
	UserOpConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.UserOpSerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[InitSrvConn] 连接 【订单服务失败】", err.Error())
		return
	}
	MessageClient := proto.NewMessageClient(UserOpConn)
	global.MessageSrvClient = MessageClient

	AddressClient := proto.NewAddressClient(UserOpConn)
	global.AddressSrvClient = AddressClient

	UserFavClient := proto.NewUserFavClient(UserOpConn)
	global.UserFavSrvClient = UserFavClient

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
}
