package initialize

import (
	"encoding/json"
	"fmt"

	"mxshop_srvs/internal/configboot"
	"mxshop_srvs/order_srv/global"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// GetEnvInfo 通过环境变量将线上线下环境隔离（兼容旧代码；bootstrap  profile 优先读 MXSHOP_NACOS_BOOTSTRAP_PROFILE）。
func GetEnvInfo(env string) bool {
	viper.AutomaticEnv()
	return viper.GetBool(env)
}

func InitConfig() {
	serviceRoot := configboot.ServiceRoot(1)
	configboot.LoadDotEnv(serviceRoot)

	if configboot.ConfigSource() == "local" {
		localPath := configboot.LocalServerConfigPath(serviceRoot)
		if err := configboot.ReadLocalServerFile(localPath, global.ServerConfig); err != nil {
			panic(err)
		}
		zap.S().Infow("已使用本地 Server 配置（未连接 Nacos）", "path", localPath)
		fmt.Println("serverconfig:", global.ServerConfig)
		return
	}

	bootstrapPath := configboot.ResolveBootstrapYAML(serviceRoot)
	v := viper.New()
	v.SetConfigFile(bootstrapPath)
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := v.Unmarshal(global.NacosConfig); err != nil {
		panic(err)
	}

	zap.S().Info("配置信息", global.NacosConfig)

	logDir, cacheDir := configboot.NacosClientDirs(serviceRoot)
	sc := []constant.ServerConfig{
		{
			IpAddr: global.NacosConfig.NacosInfo.Host,
			Port:   global.NacosConfig.NacosInfo.Port,
		},
	}

	cc := constant.ClientConfig{
		NamespaceId:         global.NacosConfig.NacosInfo.NamespaceId,
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              logDir,
		CacheDir:            cacheDir,
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": sc,
		"clientConfig":  cc,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: global.NacosConfig.NacosInfo.DataId,
		Group:  global.NacosConfig.NacosInfo.Group,
	})

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal([]byte(content), global.ServerConfig)
	if err != nil {
		panic(err)
	}
	fmt.Println("serverconfig:", global.ServerConfig)
}
