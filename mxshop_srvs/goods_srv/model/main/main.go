package main

import (
	"context"
	"log"
	"mxshop_srvs/goods_srv/model"
	"os"
	"strconv"
	"time"

	"github.com/olivere/elastic/v7"

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

func InitMysql() {
	dsn := "root:Qq/2013XiaoKUang@tcp(127.0.0.1:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

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
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		//以自定义名称表名写入数据库
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
}

//MysqlToEs 将数据写入es中
func MysqlToEs() {
	dsn := "root:Qq/2013XiaoKUang@tcp(127.0.0.1:3306)/mxshop_goods_srv?charset=utf8mb4&parseTime=True&loc=Local"

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

	url := "http://localhost:9200/"
	logger := log.New(os.Stdout, "mxshop", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false), elastic.SetTraceLog(logger))
	if err != nil {
		panic(err)
	}

	var goodslist []model.Goods
	db.Find(&goodslist)
	for _, g := range goodslist {
		goods := model.EsGoods{
			ID:          g.ID,
			CategoryID:  g.CategoryID,
			BrandsID:    g.BrandsID,
			OnSale:      g.OnSale,
			ShipFree:    g.ShipFree,
			IsNew:       g.IsNew,
			IsHot:       g.IsHot,
			Name:        g.Name,
			ClickNum:    g.ClickNum,
			SoldNum:     g.SoldNum,
			FavNum:      g.FavNum,
			MarketPrice: g.MarketPrice,
			GoodsBrief:  g.GoodsBrief,
			ShopPrice:   g.ShopPrice,
		}

		_, err = client.Index().Index(model.EsGoods{}.GetIndexName()).BodyJson(goods).Id(strconv.Itoa(int(goods.ID))).Do(context.Background())
		if err != nil {
			panic(err)
		}
	}
}

func main() {
	MysqlToEs()

	//db.AutoMigrate(&model.Category{})

	//var categorys []model.Category
	//
	////查询当前分类
	//if result := db.Find(&categorys, 135482).First(&categorys); result.RowsAffected == 0 {
	//	status.Errorf(codes.InvalidArgument, "商品分类不存在")
	//}

	//for _, value := range categorys {
	//	fmt.Println(value.Name)
	//	fmt.Println(value.ID)
	//	fmt.Println(value.SubCategory)
	//}
	//
	//var categorybrand []model.GoodsCategoryBrand
	//var total int64
	//db.Model(&model.GoodsCategoryBrand{}).Count(&total)
	//fmt.Println(total)
	//
	//db.Preload("Category").Preload("Brands").Scopes(Paginate(1, 5)).Find(&categorybrand)
	//
	//for _, value := range categorybrand {
	//	fmt.Println(value.Brands.Name)
	//	fmt.Println(value.Category.Name)
	//}

	//var goods []model.Goods
	//_ = db.Find(&goods)
	//i := 0
	//for _, value := range goods {
	//	fmt.Println(value.Name)
	//	i++
	//}
	//fmt.Println("共计商品:", i)
	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("admin123", options)
	//fmt.Println(len(encodedPwd))
	//Newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//
	//for i := 0; i < 3; i++ {
	//	user := &model.User{
	//		Mobile:   fmt.Sprintf("1758561499%d", i),
	//		PassWord: Newpassword,
	//		NickName: fmt.Sprintf("yangkuang%d", i),
	//		Role:     2,
	//	}
	//	db.Save(&user)
	//}

	//
	//db.AutoMigrate(&model.User{})

	// Using custom options
	//options := &password.Options{16, 100, 32, sha512.New}
	//salt, encodedPwd := password.Encode("generic password", options)
	//fmt.Println(len(encodedPwd))
	//Newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)
	//fmt.Println(len(Newpassword))
	//
	//Passwordinfo := strings.Split(Newpassword, "$")
	//fmt.Println(Passwordinfo)
	//check := password.Verify("generic password", Passwordinfo[2], encodedPwd, options)
	//fmt.Println(check) // true
}
