package main

import (
	"log"
	"mxshop_srvs/inventory_srv/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//func Md5(code string) string {
//	//实例化一个md5的对象，将code写入其中
//	MD5 := md5.New()
//	_, _ = io.WriteString(MD5, code)
//	return hex.EncodeToString(MD5.Sum(nil))
//}

// Paginate 将数据进行分页
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
	dsn := "root:jhfdijfidjfg@tcp(127.0.0.1:3306)/mxshop_inventory_srv?charset=utf8mb4&parseTime=True&loc=Local"

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

	db.AutoMigrate(&model.Inventory{}, &model.StockSellDetail{})
	//
	//var goodsdetail model.StockSellDetail
	//goodsdetail.OrderSn = "ice_moss"
	//goodsdetail.Detail = []model.GoodsDetail{{1, 2}, {3, 4}}
	//goodsdetail.Status = 1
	//db.Create(&goodsdetail)

	//var goodsdetail model.StockSellDetail
	//db.Where(model.StockSellDetail{OrderSn: "ice_moss"}).First(&goodsdetail)
	//fmt.Println(goodsdetail)

}
