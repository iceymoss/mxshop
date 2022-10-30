package main

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/goods_srv/proto"
)

func TestCategoryBrandList() {
	c := context.Background()
	Rsp, err := BrandClient.CategoryBrandList(c, &proto.CategoryBrandFilterRequest{
		PagePerNums: 100,
		Pages:       1,
	})
	if err != nil {
		log.Fatal("获取品牌分类失败", err)
	}
	fmt.Println("共计", Rsp.Total)
	for _, value := range Rsp.Data {
		fmt.Println("分类", value.Category.Name)
		fmt.Println("品牌", value.Brand.Name)
	}

}

func TestGetCategoryBrandList() {
	fmt.Println("GetCategoryBrandList begin:")
	c := context.Background()
	Rsp, err := BrandClient.GetCategoryBrandList(c, &proto.CategoryInfoRequest{
		Id: int32(130368),
	})
	if err != nil {
		log.Fatal("获取品牌失败", err)
	}
	fmt.Println("总数：", Rsp.Total)
	for _, value := range Rsp.Data {
		fmt.Println("品牌", value.Name)
	}
}

func TestCreateCategoryBrand() {
	c := context.Background()
	BrandClient.CreateCategoryBrand(c, &proto.CategoryBrandRequest{
		CategoryId: 238014,
		BrandId:    1113,
	})
}

func TestDeleteCategoryBrand() {
	c := context.Background()
	BrandClient.DeleteCategoryBrand(c, &proto.CategoryBrandRequest{
		Id: 25799,
	})
}

func TestUpdateCategoryBrand() {
	c := context.Background()
	BrandClient.UpdateCategoryBrand(c, &proto.CategoryBrandRequest{
		Id:         25800,
		CategoryId: 238014,
		BrandId:    1113,
	})
}
