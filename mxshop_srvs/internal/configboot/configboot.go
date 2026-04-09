// Package configboot 统一处理 .env、本地 YAML 与 Nacos bootstrap YAML 的选择逻辑。
package configboot

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// ServiceRoot 返回调用方所在微服务根目录的绝对路径（在 initialize.InitConfig 内调用时请传 skip=1）。
func ServiceRoot(skip int) string {
	_, f, _, ok := runtime.Caller(skip)
	if !ok {
		return ""
	}
	// .../xxx_srv/initialize/config.go -> .../xxx_srv
	return filepath.Clean(filepath.Join(filepath.Dir(f), ".."))
}

// LoadDotEnv 从微服务目录加载 .env（文件不存在时忽略）。
func LoadDotEnv(serviceRoot string) {
	_ = godotenv.Load(filepath.Join(serviceRoot, ".env"))
}

// ConfigSource 返回 nacos 或 local。默认 nacos，与改造前行为一致。
func ConfigSource() string {
	s := strings.ToLower(strings.TrimSpace(os.Getenv("MXSHOP_CONFIG_SOURCE")))
	switch s {
	case "local":
		return "local"
	default:
		return "nacos"
	}
}

// NacosBootstrapProfile 返回用于选择 config-debug.yaml / config-por.yaml 的 profile。
// 优先读 MXSHOP_NACOS_BOOTSTRAP_PROFILE；未设置时沿用 MXSHOP_DEBUG（true -> debug）。
func NacosBootstrapProfile() string {
	if p := strings.ToLower(strings.TrimSpace(os.Getenv("MXSHOP_NACOS_BOOTSTRAP_PROFILE"))); p == "debug" || p == "por" {
		return p
	}
	v := viper.New()
	v.AutomaticEnv()
	if v.GetBool("MXSHOP_DEBUG") {
		return "debug"
	}
	return "por"
}

// ResolveBootstrapYAML 返回 Nacos bootstrap 用的 yaml 绝对路径。
func ResolveBootstrapYAML(serviceRoot string) string {
	if custom := strings.TrimSpace(os.Getenv("MXSHOP_BOOTSTRAP_FILE")); custom != "" {
		if filepath.IsAbs(custom) {
			return custom
		}
		return filepath.Join(serviceRoot, custom)
	}
	prof := NacosBootstrapProfile()
	return filepath.Join(serviceRoot, fmt.Sprintf("config-%s.yaml", prof))
}

// LocalServerConfigPath 返回本地 Server 全量配置文件路径（默认 local-server.yaml；相对路径相对微服务根目录）。
func LocalServerConfigPath(serviceRoot string) string {
	p := strings.TrimSpace(os.Getenv("MXSHOP_LOCAL_SERVER_CONFIG"))
	if p == "" {
		p = "local-server.yaml"
	}
	if filepath.IsAbs(p) {
		return p
	}
	return filepath.Join(serviceRoot, p)
}

// NacosClientDirs 将 Nacos SDK 日志与缓存放到各微服务目录下，避免依赖进程工作目录。
func NacosClientDirs(serviceRoot string) (logDir, cacheDir string) {
	return filepath.Join(serviceRoot, "tmp/nacos/log"),
		filepath.Join(serviceRoot, "tmp/nacos/cache")
}

// ReadLocalServerFile 读取本地全量配置（YAML，与 Nacos 下发 JSON 字段一致，依赖 struct 的 mapstructure tag）。
func ReadLocalServerFile(path string, target interface{}) error {
	v := viper.New()
	v.SetConfigFile(path)
	if err := v.ReadInConfig(); err != nil {
		return err
	}
	return v.Unmarshal(target)
}
