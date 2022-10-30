package main

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"
)

func TestGoodsList() {
	c := context.Background()
	Rsp, err := BrandClient.GoodsList(c, &proto.GoodsFilterRequest{
		Pages:       1,
		PagePerNums: 100,
		TopCategory: 130361,
		PriceMin:    20,
		PriceMax:    100,
	})
	if err != nil {
		log.Fatal("获取商品失败", err)
	}
	fmt.Println(Rsp.Total)
	for _, goods := range Rsp.Data {
		fmt.Println("name:", goods.Name, "price:", goods.ShopPrice, "category:", goods.CategoryId, goods.Category.Id)
	}

}

func TestGetGoodsDetail() {
	c := context.Background()
	goods, err := BrandClient.GetGoodsDetail(c, &proto.GoodInfoRequest{
		Id: 848,
	})
	if err != nil {
		log.Fatal("获取商品失败", err)
	}
	fmt.Println("name:", goods.Name, "price:", goods.ShopPrice, "category:", goods.CategoryId, goods.Category.Id)

}

func TestBatchGetGoods() {
	c := context.Background()
	var id []int32
	for i := 422; i < 450; i++ {
		id = append(id, int32(i))
	}

	fmt.Println(id)
	Rsp, err := BrandClient.BatchGetGoods(c, &proto.BatchGoodsIdInfo{
		Id: id,
	})
	if err != nil {
		log.Fatal("获取商品失败", err)
	}

	for _, good := range Rsp.Data {
		fmt.Println(good.Name)
	}
}

func TestCreateGoods() {
	c := context.Background()
	images := model.GormList{
		"https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
		"https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
		"https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
		"https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
	}

	descImages := model.GormList{
		"https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
		"https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
		"https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
	}

	Rsp, err := BrandClient.CreateGoods(c, &proto.CreateGoodsInfo{
		Name:            "台湾四季青香水柠檬（5斤装）屏东尤力克新鲜免邮奶茶店一点点用",
		GoodsSn:         "32",
		MarketPrice:     12.8,
		ShopPrice:       7.8,
		GoodsBrief:      "台湾四季青香水柠檬（5斤装）屏东尤力克新鲜免邮奶茶店一点点用",
		Images:          images,
		DescImages:      descImages,
		GoodsFrontImage: "https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
		BrandId:         1113,
		CategoryId:      238014,
	})
	if err != nil {
		log.Fatal("添加商品失败", err)
	}
	fmt.Println(Rsp.Id)
}

func TestUpdateGoods() {
	c := context.Background()
	BrandClient.UpdateGoods(c, &proto.CreateGoodsInfo{
		Id:              848,
		Name:            "台湾四季青香水柠檬",
		GoodsSn:         "32",
		MarketPrice:     22.8,
		ShopPrice:       8.8,
		GoodsBrief:      "台湾四季青香水柠檬（5斤装）屏东尤力克新鲜免邮奶茶店一点点用",
		GoodsFrontImage: "https://py-go.oss-cn-beijing.aliyuncs.com/goods_images/fd23760f7dfe98279edc9d059b3b3431",
		BrandId:         1113,
		CategoryId:      238014,
		OnSale:          true,
		IsHot:           true,
	})
}
