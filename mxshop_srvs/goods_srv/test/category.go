package main

import (
	"context"
	"fmt"
	"log"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGetCategoryAllList() {
	c := context.Background()
	Rsp, err := BrandClient.GetAllCategorysList(c, &emptypb.Empty{})
	if err != nil {
		log.Fatal("获取分类列表失败", err)
	}
	fmt.Println(Rsp.JsonData)

}

func TestGetSubCategory() {
	c := context.Background()
	Rsp, err := BrandClient.GetSubCategory(c, &proto.CategoryListRequest{
		Id: 238014,
	})
	if err != nil {
		log.Fatal("获取分类列表失败", err)
	}
	fmt.Println("当前分类", Rsp.Info)
	fmt.Println("子分类", Rsp.SubCategorys)

}

func TestCreateCategory() {
	c := context.Background()
	Rsp, err := BrandClient.CreateCategory(c, &proto.CategoryInfoRequest{
		Name:           "鲜制牛肉",
		ParentCategory: 238014,
		Level:          3,
		IsTab:          true,
	})
	if err != nil {
		log.Fatal("添加分类失败", err)
	}
	fmt.Println(Rsp.Id)
}

func TestDeleteCategory() {
	c := context.Background()
	_, err := BrandClient.DeleteCategory(c, &proto.DeleteCategoryRequest{
		Id: 238015,
	})
	if err != nil {
		log.Fatal("删除分类失败", err)
	}
}

func TestUpdateCategory() {
	c := context.Background()
	_, err := BrandClient.UpdateCategory(c, &proto.CategoryInfoRequest{
		Id:    238014,
		Name:  "优惠牛肉",
		Level: 2,
		IsTab: true,
	})
	if err != nil {
		log.Fatal("更新分类失败", err)
	}

}
