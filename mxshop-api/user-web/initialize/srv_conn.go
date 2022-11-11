package initialize

import (
	"fmt"

	"mxshop-api/user-web/global"
	"mxshop-api/user-web/proto"
	"mxshop-api/user-web/utils/otgrpc"

	"github.com/hashicorp/consul/api"
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

	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败", err.Error())
		return
	}
	userClient := proto.NewUserClient(conn)
	global.UserClient = userClient
}

// InitSrvConn2 InitSrvConn 从consul配置中获取信息并连接grpc服务
//InitSrvConn 连接注册中心的第一个版本
func InitSrvConn2() {
	cfg := api.DefaultConfig()
	consulInfo := global.ServerConfig.ConsulInfo
	cfg.Address = fmt.Sprintf("%s:%d", consulInfo.Host, consulInfo.Port)

	userSrvHost := ""
	userSrvPort := 0
	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == \"%s\"", global.ServerConfig.UserSerInfo.Name))
	//data, err := client.Agent().ServicesWithFilter(fmt.Sprintf("Service == user_srv_test"))
	if err != nil {
		panic(err)
	}
	for _, v := range data {
		userSrvHost = v.Address
		userSrvPort = v.Port
		break
	}

	if userSrvHost == "" {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】", err.Error())
		return
	}

	//待解决问题：1.服务下线了，2.改ip了，3.改端口了
	//拨号连接grpc服务
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", userSrvHost, userSrvPort), grpc.WithInsecure())
	if err != nil {
		zap.S().Errorw("[GetUserList] 连接 【用户服务失败】", err.Error())
		return
	}

	userClient := proto.NewUserClient(conn)
	global.UserClient = userClient
}
