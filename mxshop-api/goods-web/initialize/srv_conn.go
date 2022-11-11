package initialize

import (
	"fmt"
	"mxshop-api/goods-web/global"
	"mxshop-api/goods-web/proto"
	"mxshop-api/goods-web/utils/otgrpc"

	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/opentracing/opentracing-go"
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
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败", err.Error())
		return
	}
	goodsClient := proto.NewGoodsClient(conn)
	global.GoodsSrvClient = goodsClient

	InvtConn, err := grpc.Dial(
		fmt.Sprintf("consul://%s:%d/%s?wait=14s&tag=srv", consul.Host, consul.Port, global.ServerConfig.InventSerInfo.Name),
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
		grpc.WithUnaryInterceptor(otgrpc.OpenTracingClientInterceptor(opentracing.GlobalTracer())),
	)

	fmt.Println(global.ServerConfig)

	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败", err.Error())
		return
	}
	global.InventoryClient = proto.NewInventoryClient(InvtConn)
	global.GoodsSrvClient = goodsClient

}
