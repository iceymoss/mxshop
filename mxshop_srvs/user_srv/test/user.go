package main

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/user_srv/initialize"
	"mxshop_srvs/user_srv/proto"
	"time"

	"google.golang.org/grpc"
)

var userClient proto.UserClient
var Conn *grpc.ClientConn

func Init() {
	var err error
	//使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
	Conn, err = grpc.Dial("localhost:58496", grpc.WithInsecure())
	if err != nil {
		log.Panicln("连接失败", err)
	}
	userClient = proto.NewUserClient(Conn)
}

func TestGetUserList() {
	c := context.Background()
	ResUserInfo, err := userClient.GetUserInfoList(c, &proto.PageInfo{
		Pn:    1,
		Psize: 5,
	})
	if err != nil {
		log.Fatal("获取用户列表失败", err)
	}
	for _, user := range ResUserInfo.Data {
		fmt.Println("用户列表：", user)
		res, err := userClient.CheckPassWord(c, &proto.PasswordCheckInfo{
			Password:          "admin123",
			EncryptedPassword: user.Password,
		})
		if err != nil {
			log.Fatal("密码校验失败", err)

		}
		fmt.Println("密码核对：", res.Success)
		fmt.Println("通过Id查询")
		IdFindUser, err := userClient.GetUserById(c, &proto.IdRequest{
			Id: user.Id,
		})
		if err != nil {
			log.Fatal("id查询失败", err)
		}
		fmt.Println(IdFindUser)
	}
}

func TestGetUserByMobile() {
	c := context.Background()
	for i := 0; i < 11; i++ {
		user, err := userClient.GetUserByMobile(c, &proto.MobileRequest{
			Mobile: fmt.Sprintf("1758561098%d", i),
		})
		if err != nil {
			log.Fatal("数据查询失败", err)
		}
		fmt.Println(user)
	}
}

func TestUpdateUserInfo() {
	c := context.Background()
	_, err := userClient.UpdateUser(c, &proto.UpdateUserInfo{
		Id:       2,
		NickName: "babay",
		Gender:   "male",
		Birthday: uint64(time.Now().Unix()),
	})
	if err != nil {
		log.Fatal("更新失败", err)
	}
}

func TestCreateUserInfo() {
	c := context.Background()
	resr, err := userClient.CreateUser(c, &proto.CreateUserInfo{
		NickName: "ice_moss",
		Password: "admin123",
		Mobile:   "17500000001",
	})
	if err != nil {
		log.Fatal("创建失败", err)
	}
	fmt.Println(resr)
}

func main() {
	Init()
	TestGetUserList()
	initialize.InitConfig()
	TestGetUserByMobile()
	TestUpdateUserInfo()
	TestCreateUserInfo()
	Conn.Close()
}
