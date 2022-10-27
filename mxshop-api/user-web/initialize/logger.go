package initialize

import (
	"log"

	"go.uber.org/zap"
)

//	1. zap.S可以配置一个全局的sugar,用来配置全局logger
//	2. 日志级别：debug、info、warn、error、fetal
//	3. S函数和L函数可以配置全局的安全的logger

//InitLogger 初始化日志
func InitLogger() {
	//初始化日志
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal("日志初始化失败", err.Error())
	}
	//使用全局logger
	zap.ReplaceGlobals(logger)
}
