package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/handler"
	"mxshop_srvs/user_srv/initialize"
	"mxshop_srvs/user_srv/proto"
	"mxshop_srvs/user_srv/utils"

	"github.com/hashicorp/consul/api"
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

	//将线上线下address隔离，固定本地端口，线上动态端口

	IP := flag.String("ip", global.ServerConfig.Host, "ip地址")
	Port := flag.Int("port", 8081, "端口号")
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
	proto.RegisterUserServer(s, &handler.UserServer{})

	//注册健康检查
	//将user_srv服务注册到consul中，让web层可获取其配置信息
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	//DefaultConfig 返回客户端的默认配置
	cfg := api.DefaultConfig()
	cfg.Address = fmt.Sprintf("%s:%d", global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	//生成对应的检查对象
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", global.ServerConfig.Host, *Port),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}
	//生成注册对象
	registration := new(api.AgentServiceRegistration)
	registration.Name = global.ServerConfig.Name
	serverID := fmt.Sprintf("%s", uuid.NewV4())
	registration.ID = serverID
	registration.Port = *Port
	registration.Tags = []string{"user-srv", "ice_moss", "learning"}
	registration.Address = global.ServerConfig.Host
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}

	//开启服务
	go func() {
		err = s.Serve(conn)
		if err != nil {
			zap.S().Errorw("fail server start for GRPC", err)
		}
	}()

	//接收终止信号
	qiut := make(chan os.Signal)
	//接收control+c
	signal.Notify(qiut, syscall.SIGINT, syscall.SIGTERM)
	<-qiut
	err = client.Agent().ServiceDeregister(serverID)
	if err != nil {
		zap.S().Info("注销失败", err)
	}
	zap.S().Info("注销成功")

}
