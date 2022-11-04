package main

import (
	"context"
	"fmt"
	"log"

	"mxshop_srvs/userop_srv/proto"

	"google.golang.org/grpc"
)

var AddressClient proto.AddressClient
var MessageClient proto.MessageClient
var UserFavClient proto.UserFavClient
var Conn *grpc.ClientConn

func Init() {
	var err error
	//使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
	Conn, err = grpc.Dial("10.2.94.231:8084", grpc.WithInsecure())
	if err != nil {
		log.Panicln("连接失败", err)
	}
	AddressClient = proto.NewAddressClient(Conn)
	MessageClient = proto.NewMessageClient(Conn)
	UserFavClient = proto.NewUserFavClient(Conn)
}

func TestCreateAddress() {
	Rsp, err := AddressClient.CreateAddress(context.Background(), &proto.GetAddressRequest{
		UserId:       7,
		Province:     "江苏省",
		City:         "无锡市",
		District:     "梁溪区",
		Address:      "锡山大道333号",
		SignerName:   "ice_moss",
		SignerMobile: "17585610988",
	})
	if err != nil {
		log.Fatal("新建地址失败", err)
		return
	}
	fmt.Println(Rsp.Id)
}

func TestGetAddressList() {
	Rsp, err := AddressClient.GetAddressList(context.Background(), &proto.GetAddressRequest{})
	if err != nil {
		log.Fatal("获取地址列表失败", err)
		return
	}
	fmt.Println(Rsp.Data)
}

func TestUpdateAddress() {
	_, err := AddressClient.UpdateAddress(context.Background(), &proto.GetAddressRequest{
		Id:           4,
		UserId:       7,
		Province:     "江苏省",
		City:         "南京市",
		District:     "鼓楼区",
		Address:      "鼓楼大道333号",
		SignerName:   "ice_moss",
		SignerMobile: "17585610988",
	})
	if err != nil {
		log.Fatal("更新地址失败", err)
		return
	}
}

func TestDelete() {
	_, err := AddressClient.DeleteAddress(context.Background(), &proto.GetAddressRequest{
		Id:     1,
		UserId: 0,
	})
	if err != nil {
		log.Fatal("删除地址失败", err)
		return
	}

}

func TestCreateMessage() {
	Rsp, err := MessageClient.CreateMessage(context.Background(), &proto.MessageRequest{
		UserId:      7,
		MessageType: 1,
		Subject:     "给系统好评",
		Message:     "我已经在本品台购买很多商品了，体验很好",
		File:        "https://aliyun.com/mshop_srv",
	})
	if err != nil {
		log.Fatal("创建失败", err)
		return
	}
	fmt.Println(Rsp.Id)
}

func TestGetMessageList() {
	Rsp, err := MessageClient.GetMessageList(context.Background(), &proto.MessageRequest{})
	if err != nil {
		log.Fatal("获取失败", err)
		return
	}
	fmt.Println(Rsp.Data)

}

func TestAddUserFav() {
	_, err := UserFavClient.AddUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  7,
		GoodsId: 421,
	})
	if err != nil {
		log.Fatal("加入收藏失败", err)
		return
	}
}

func TestGetFavList() {
	Rsp, err := UserFavClient.GetFavList(context.Background(), &proto.UserFavRequest{})
	if err != nil {
		log.Fatal("查询收藏失败", err)
		return
	}
	fmt.Println(Rsp.Data)

}

func TestDeleteUserFav() {
	_, err := UserFavClient.DeleteUserFav(context.Background(), &proto.UserFavRequest{
		UserId:  13,
		GoodsId: 421,
	})
	if err != nil {
		log.Fatal("删除收藏失败", err)
		return
	}
}

func TestGetUserFavDetail() {
	_, err := UserFavClient.GetUserFavDetail(context.Background(), &proto.UserFavRequest{
		UserId:  13,
		GoodsId: 421,
	})
	if err != nil {
		log.Fatal("查询收藏失败", err)
		return
	}

}

func main() {
	Init()
	//TestGetAddressList()
	//TestDelete()
	//TestCreateAddress()
	//TestUpdateAddress()

	//TestCreateMessage()
	//TestGetMessageList()

	//TestAddUserFav()
	//TestGetFavList()

	//TestDeleteUserFav()

	TestGetUserFavDetail()

	Conn.Close()

}
