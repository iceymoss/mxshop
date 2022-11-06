package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"mxshop-api/userop-web/global"
	"mxshop-api/userop-web/initialize"
	"mxshop-api/userop-web/utils"
	"mxshop-api/userop-web/utils/register/consul"
	myvalidator "mxshop-api/userop-web/validator"

	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//1. 初始化logger
	initialize.InitLogger()

	//2. 初始化配置
	initialize.InitConfig()

	//3. 初始化Router
	Router := initialize.Routers()

	//4. 初始化翻译
	if err := initialize.InitTrans("zh"); err != nil {
		panic(err)
	}

	//5. 初始化grpc Client
	initialize.InitSrvConn()

	//将线上线下address隔离，固定本地端口，线上动态端口
	viper.AutomaticEnv()
	debug := viper.GetBool("MXSHOP_DEBUG")
	if !debug {
		port, err := utils.GetFreePort()
		global.ServerConfig.Port = port
		if err != nil {
			zap.S().Errorw("获取端口失败", err)
		}
	}

	//注册验证器,将自定义validator放入翻译器中
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("mobile", myvalidator.ValidatorMobile)
		_ = v.RegisterTranslation("mobile", global.Trans, func(ut ut.Translator) error {
			return ut.Add("mobile", "{0} 非法的手机号码!", true) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("mobile", fe.Field())
			return t
		})
	}

	//将服务注册到注册中心
	registry_client := consul.NewRegistryClient(global.ServerConfig.ConsulInfo.Host, global.ServerConfig.ConsulInfo.Port)
	id := uuid.NewV4()
	serverIdstr := fmt.Sprintf("%s", id)
	err := registry_client.Register(global.ServerConfig.Host, global.ServerConfig.Port, global.ServerConfig.Name, global.ServerConfig.Tag, serverIdstr)
	if err != nil {
		zap.S().Panic("服务注册失败", err.Error())
	}

	//启动服务
	go func() {
		zap.S().Debugf("服务启动：%d", global.ServerConfig.Port)
		if err := Router.Run(fmt.Sprintf(":%d", global.ServerConfig.Port)); err != nil {
			zap.S().Panic("启动失败", err.Error())
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
