package initialize

import (
	"fmt"

	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

//InitSrvConn 连接到consul注册中心并对服务做负载均衡
func InitSrvConn() {
	consul := global.ServerConfig.ConsulInfo
	conn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.UserSerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败", err.Error())
		return
	}
	goodsClient := proto.NewGoodsClient(conn)
	global.GoodsSrvClient = goodsClient
}
