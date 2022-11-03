package initialize

import (
	"encoding/json"
	"fmt"

	"mxshop_srvs/order_srv/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

//GetEnvInfo 通过环境变量 将线上线下环境隔离
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	//通过环境变量配置线下线上
	debug := GetEnvInfo("MXSHOP_DEBUG")
	configFilePrefix := "config"
	configFileName := fmt.Sprintf("order_srv/%s-por.yaml", configFilePrefix)
	if debug {
		//Users/feng/go/src/mxshop_srvs/order_srv/config-debug.yaml
		configFileName = fmt.Sprintf("order_srv/%s-debug.yaml", configFilePrefix)
	}

	//实例化对象
	v := viper.New()
	//读取配置文件
	v.SetConfigFile(configFileName)
	//读入文件
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	//将数据放入global.ServerConfig 这个对象如何在其他文件中使用--全局变量
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}

	zap.S().Info("配置信息", global.NacosConfig)

	//配置nacos服务去获取配置文件
	//服务端配置
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.NacosInfo.Host,
			Port:   global.NacosConfig.NacosInfo.Port,
		},
	}

	//客服端配置
	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.NacosInfo.NamespaceId, // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		//RotateTime:          "1h",
		//MaxAge:              3,
		LogLevel: "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	//获取配置
	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.NacosInfo.DataId,
		Group:  global.NacosConfig.NacosInfo.Group,
	})

	if err != nil {
		panic(err)
	}

	//将配置文件映射至global.ServerConfig中
	err = json.Unmarshal([]byte(content), &global.ServerConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("serverconfig:", global.ServerConfig)

	//v.WatchConfig()
	//v.OnConfigChange(func(e fsnotify.Event) {
	//	zap.S().Infof("配置文件产生变化：%v", global.ServerConfig)
	//	_ = v.ReadInConfig()
	//	v.Unmarshal(global.ServerConfig)
	//	zap.S().Infof("配置信息：%v", global.ServerConfig)
	//})
}
