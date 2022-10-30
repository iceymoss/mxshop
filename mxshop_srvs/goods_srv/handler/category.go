package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
)

//GetAllCategorysList 获取所有分类
func (g *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []model.Category
	//通过反向查询，查一级内目，查二级内目，查三级内目
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)

	//json序列化
	b, _ := json.Marshal(&categorys)

	return &proto.CategoryListResponse{JsonData: string(b)}, nil
}

//GetSubCategory 获取子分类
func (g *GoodsServer) GetSubCategory(c context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {
	categoryListResponse := proto.SubCategoryListResponse{}

	var category model.Category
	//查询分类是否存在
	//查询当前目录
	if result := global.DB.Find(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "分类不存在")
	}
	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	//当前目录子目录
	var subCategorys []model.Category
	var categoryInfoResponse []*proto.CategoryInfoResponse
	perloads := "SubCategory"
	if category.Level == 1 {
		perloads = "SubCategory.SubCategory"
	}
	global.DB.Where(&model.Category{ParentCategoryID: req.Id}).Preload(perloads).Find(&subCategorys)
	for _, subCategory := range subCategorys {
		categoryInfoResponse = append(categoryInfoResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.ID,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}
	categoryListResponse.SubCategorys = categoryInfoResponse
	return &categoryListResponse, nil
}

//CreateCategory 创建分类
func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {
	category := model.Category{}
	cMap := map[string]interface{}{}
	cMap["name"] = req.Name
	cMap["level"] = req.Level
	//category.Name = req.Name
	//category.Level = req.Level
	if req.Level != 1 {
		//不是一级类目，需要将指向一级类目
		cMap["parent_category_id"] = req.ParentCategory
		//category.ParentCategoryID = req.ParentCategory
	}
	//category.IsTab = req.IsTab
	cMap["is_tab"] = req.IsTab

	//global.DB.Save(&category)
	tr := global.DB.Model(model.Category{}).Create(&cMap)
	fmt.Println(tr)

	return &proto.CategoryInfoResponse{Id: int32(category.ID)}, nil
}

//DeleteCategory 删除分类
func (g *GoodsServer) DeleteCategory(c context.Context, req *proto.DeleteCategoryRequest) (*empty.Empty, error) {
	if result := global.DB.Delete(&model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	return &empty.Empty{}, nil
}

//UpdateCategory 更新分类
func (g *GoodsServer) UpdateCategory(c context.Context, req *proto.CategoryInfoRequest) (*empty.Empty, error) {
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类不存在")
	}
	if req.Name != "" {
		category.Name = req.Name
	}

	if req.ParentCategory != 0 {
		//类目级别
		category.ParentCategoryID = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = req.IsTab
	}

	global.DB.Save(&category)

	return &empty.Empty{}, nil
}
