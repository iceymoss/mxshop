// Package migrateenv 与 mxshop_srvs 共用同一环境变量约定，便于统一部署脚本。
package migrateenv

import (
	"os"
	"strings"
)

// AutoMigrateEnabled 当环境变量 MXSHOP_AUTO_MIGRATE 为 1、true、yes、on（不区分大小写）时返回 true。
func AutoMigrateEnabled() bool {
	v := strings.TrimSpace(strings.ToLower(os.Getenv("MXSHOP_AUTO_MIGRATE")))
	switch v {
	case "1", "true", "yes", "on":
		return true
	default:
		return false
	}
}
