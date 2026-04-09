package migratehint

import (
	"mxshop-api/internal/migrateenv"

	"go.uber.org/zap"
)

// LogIfEnabled 在 Web 层不执行 AutoMigrate；若环境变量已开启则打印说明，避免误以为本进程会迁库。
func LogIfEnabled(serviceName string) {
	if !migrateenv.AutoMigrateEnabled() {
		return
	}
	zap.S().Infow(
		"MXSHOP_AUTO_MIGRATE 已启用：业务表迁移由 mxshop_srvs 中对应 gRPC 服务在启动时执行，本 BFF 不连接业务 MySQL",
		"service", serviceName,
	)
}
