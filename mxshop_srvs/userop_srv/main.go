package main

import (
	"flag"
	"fmt"
	"log"
	"mxshop_srvs/userop_srv/handler"
	"net"
	"os"
	"os/signal"
	"syscall"

	"mxshop_srvs/userop_srv/global"
	"mxshop_srvs/userop_srv/initialize"
	"mxshop_srvs/userop_srv/proto"
	"mxshop_srvs/userop_srv/utils"
	"mxshop_srvs/userop_srv/utils/register/consul"

	uuid "github.com/satori/go.uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
)

func main() {

	//1.初始化日志
	initialize.InitLogger()

	//2.初始化配置
	initialize.InitConfig()

	//3.初始化数据库
	initialize.InitDB()
	initialize.InitRs()

	//将线上线下address隔离，固定本地端口，线上动态端口
	IP := flag.String("ip", global.ServerConfig.Host, "ip地址")
	Port := flag.Int("port", 8085, "端口号")
	flag.Parse()

	zap.S().Info(global.ServerConfig)
	//监听端口
	zap.S().Info("ip:", *IP)
	if *Port == 0 {
		*Port, _ = utils.GetFreePort()
	}
	zap.S().Info("port:", *Port)

	conn, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		log.Fatal("监听端口失败", err)
	}

	//实例server
	s := grpc.NewServer()
	//注册处理逻辑
	proto.RegisterMessageServer(s, &handler.UserOpServer{})
	proto.RegisterAddressServer(s, &handler.UserOpServer{})
	proto.RegisterUserFavServer(s, &handler.UserOpServer{})

	//做负载均衡
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//注册健康检查
	//将userop_srv服务注册到consul中，让web层可获取其配置信息
	//将服务注册到注册中心
	registry_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	id := uuid.NewV4()
	serverIdstr := fmt.Sprintf("%s", id)
	err = registry_client.Register(global.ServerConfig.Host, *Port, global.ServerConfig.Name, global.ServerConfig.Tags, serverIdstr)
	if err != nil {
		zap.S().Panic("服务注册失败", err.Error())
	}
	zap.S().Info("服务启动", *Port)

	go func() {
		//开启服务
		err = s.Serve(conn)
		if err != nil {
			zap.S().Errorw("fail server start for GRPC", err)
		}
	}()

	//接收终止信号 Signal表示操作系统信号
	qiut := make(chan os.Signal)
	//接收control+c
	signal.Notify(qiut, syscall.SIGINT, syscall.SIGTERM)
	<-qiut
	err = registry_client.DeRegister(serverIdstr)
	if err != nil {
		zap.S().Info("注销失败", err)
	} else {
		zap.S().Info("注销成功")
	}
}
