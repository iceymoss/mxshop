package initialize

import (
	migrateenv "mxshop_srvs/internal/migrateenv"
	"mxshop_srvs/inventory_srv/global"
	"mxshop_srvs/inventory_srv/model"

	"go.uber.org/zap"
)

// RunAutoMigrate 在 MXSHOP_AUTO_MIGRATE 启用时对库存模型执行 GORM AutoMigrate。
func RunAutoMigrate() {
	if !migrateenv.AutoMigrateEnabled() {
		return
	}
	if global.DB == nil {
		return
	}
	err := global.DB.AutoMigrate(
		&model.Inventory{},
		&model.StockSellDetail{},
	)
	if err != nil {
		zap.S().Fatalw("数据库自动迁移失败", "error", err)
	}
	zap.S().Infow("数据库表迁移已完成", "MXSHOP_AUTO_MIGRATE", true)
}
