package main

import (
	"context"
	"fmt"
	"log"

	"mxshop_srvs/goods_srv/proto"
)

//var BrandClient proto.GoodsClient
//var Conn *grpc.ClientConn
//
//func Init() {
//	var err error
//	//使用grpc.Dial()进行拨号， grpc.WithInsecure()使用不安全的方式连接
//	Conn, err = grpc.Dial("192.168.3.111:8080", grpc.WithInsecure())
//	if err != nil {
//		log.Panicln("连接失败", err)
//	}
//	BrandClient = proto.NewGoodsClient(Conn)
//}

func TestGetBrandList() {
	c := context.Background()
	Rsp, err := BrandClient.BrandList(c, &proto.BrandFilterRequest{
		Pages:       1,
		PagePerNums: 3,
	})
	if err != nil {
		log.Fatal("获取用户列表失败", err)
	}
	for _, value := range Rsp.Data {
		fmt.Println("品牌：", value.Name)
	}
	fmt.Println("总数:", Rsp.Total)
}

func TestCreateBrand() {
	c := context.Background()
	res, err := BrandClient.CreateBrand(c, &proto.BrandRequest{
		Name: "银蕨农场",
		Logo: "https://img30.360buyimg.com/popshop/jfs/t6739/299/1630614351/29051/a47f9456/5954983fN2f665b13.jpg",
	})
	if err != nil {
		log.Fatal("新增商品失败", err)
	}
	fmt.Println("商品：", res)

}

func TestDeleteBrand() {
	c := context.Background()
	_, err := BrandClient.DeleteBrand(c, &proto.BrandRequest{
		Id: 1087,
	})
	if err != nil {
		log.Fatal("删除失败", err.Error())
	}
}

func TestUpdateBrand() {
	c := context.Background()
	_, err := BrandClient.UpdateBrand(c, &proto.BrandRequest{
		Id:   1086,
		Name: "烟台苹果",
		Logo: "https://img30.360buyimg.com",
	})
	if err != nil {
		log.Fatal("更新失败", err)
	}
}

//func TestGetBannerList() {
//	c := context.Background()
//	Rsp, err := BrandClient.BannerList(c, &empty.Empty{})
//	if err != nil {
//		log.Fatal("获取轮播图列表失败", err)
//	}
//	for _, value := range Rsp.Data {
//		fmt.Println("品牌：", value.Image)
//	}
//	fmt.Println("总数:", Rsp.Total)
//}
//
//func TestCreateBanner() {
//	c := context.Background()
//	res, err := BrandClient.CreateBanner(c, &proto.BannerRequest{
//		Id:    5,
//		Index: 5,
//		Image: "http://shop.projectsedu.com/media/banner/banner2_GmcsBvj.jpg",
//		Url:   "http://shop.projectsedu.com/media/banner",
//	})
//	if err != nil {
//		log.Fatal("新增轮播图失败", err)
//	}
//	fmt.Println("轮播图信息：", res)
//}
//
//func TestDeleteBanner() {
//	c := context.Background()
//	_, err := BrandClient.DeleteBanner(c, &proto.BannerRequest{
//		Id: 5,
//	})
//	if err != nil {
//		log.Fatal("删除失败", err.Error())
//	}
//}
//
//func TestUpdateBanner() {
//	c := context.Background()
//	_, err := BrandClient.UpdateBanner(c, &proto.BannerRequest{
//		Id:    5,
//		Index: 5,
//		Image: "http://shop.projectsedu.com/banner2_GmcsBvj.jpg",
//		Url:   "http://shop.projectsedu.com",
//	})
//	if err != nil {
//		log.Fatal("更新失败", err)
//	}
//}

//func main() {
//	Init()
//	TestGetBrandList()
//	TestGetBannerList()
//	//TestCreateBrand()
//	//TestDeleteBrand()
//	//TestUpdateBanner()
//	//TestDeleteBanner()
//}
