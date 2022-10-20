package main

import (
	"crypto/sha512"
	"fmt"
	"log"
	"os"
	"time"

	"mxshop_srvs/user_srv/model"

	"github.com/anaskhan96/go-password-encoder"
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

func main() {
	dsn := "root:Qq/2013XiaoKUang@tcp(127.0.0.1:3306)/mxshop_user_srv?charset=utf8mb4&parseTime=True&loc=Local"

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
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&model.User{})

	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode("admin123", options)
	fmt.Println(len(encodedPwd))
	Newpassword := fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	for i := 0; i < 3; i++ {
		user := &model.User{
			Mobile:   fmt.Sprintf("1758561499%d", i),
			PassWord: Newpassword,
			NickName: fmt.Sprintf("yangkuang%d", i),
			Role:     2,
		}
		db.Save(&user)
	}

	//
	db.AutoMigrate(&model.User{})

	//Using custom options
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
