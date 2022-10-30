package main

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

//var BrandClient proto.GoodsClient
//var Conn *grpc.ClientConn
//
//func Init() {
//	var err error
//	//使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
//	Conn, err = grpc.Dial("192.168.3.103:8080", grpc.WithInsecure())
//	if err != nil {
//		log.Panicln("连接失败", err)
//	}
//	BrandClient = proto.NewGoodsClient(Conn)
//}

func TestGetBannerList() {
	c := context.Background()
	Rsp, err := BrandClient.BannerList(c, &emptypb.Empty{})
	if err != nil {
		log.Fatal("获取轮播图列表失败", err)
	}
	for _, value := range Rsp.Data {
		fmt.Println("url：", value.Url)
	}
	fmt.Println("总数:", Rsp.Total)
}

func TestCreateBanner() {
	c := context.Background()
	res, err := BrandClient.CreateBanner(c, &proto.BannerRequest{
		Id:    6,
		Index: 6,
		Image: "http://shop.projectsedu.com/media/banner/banner2_GmcsBvj.jpg",
		Url:   "http://shop.projectsedu.com/media/banner",
	})
	if err != nil {
		log.Fatal("新增轮播图失败", err)
	}
	fmt.Println("轮播图信息：", res)
}

func TestDeleteBanner() {
	c := context.Background()
	_, err := BrandClient.DeleteBanner(c, &proto.BannerRequest{
		Id: 5,
	})
	if err != nil {
		log.Fatal("删除失败", err.Error())
	}
}

func TestUpdateBanner() {
	c := context.Background()
	_, err := BrandClient.UpdateBanner(c, &proto.BannerRequest{
		Id:    7,
		Index: 7,
		Image: "http://shop.projectsedu.com/banner2_GmcsBvj.jpg",
		Url:   "http:ice_moos.top/dhfidf/dhfidhf",
	})
	if err != nil {
		log.Fatal("更新失败", err)
	}
}

//func main() {
//	Init()
//	TestGetBannerList()
//	TestCreateBanner()
//	//TestDeleteBanner()
//	TestUpdateBanner()
//}
