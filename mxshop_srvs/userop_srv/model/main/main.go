package main

import (
	"log"
	"os"
	"time"

	"mxshop_srvs/userop_srv/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//Paginate 将数据进行分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func main() {
	dsn := "root:Qq/2013XiaoKUang@tcp(127.0.0.1:3306)/mxshop_userop_srv?charset=utf8mb4&parseTime=True&loc=Local"

	//用于输出使用的sql语句
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)

	//打开mysql服务中对应的数据库
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//以自定义名称表名写入数据库
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.LeavingMessages{}, &model.Address{}, &model.UserFav{})

}
