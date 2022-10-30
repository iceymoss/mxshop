package handler

import (
	"context"
	"fmt"

	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//CategoryBrandList 品牌分类列表
func (g *GoodsServer) CategoryBrandList(c context.Context, req *proto.CategoryBrandFilterRequest) (*proto.CategoryBrandListResponse, error) {
	var categoryBrands []model.GoodsCategoryBrand
	var categoryBrandListResponse proto.CategoryBrandListResponse

	//方式一：获取总数
	//var total int64
	//global.DB.Model(&model.GoodsCategoryBrand{}).Count(&total)
	//categoryBrandListResponse.Total = int32(total)

	//方式二：获取总数
	result := global.DB.Find(&categoryBrands)
	if result.Error != nil {
		return nil, result.Error
	}
	categoryBrandListResponse.Total = int32(result.RowsAffected)

	//分页
	global.DB.Preload("Category").Preload("Brands").Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&categoryBrands)

	var CategoryBrandResponse []*proto.CategoryBrandResponse
	for _, categoryBrand := range categoryBrands {
		CategoryBrandResponse = append(CategoryBrandResponse, &proto.CategoryBrandResponse{
			Id: categoryBrand.CategoryID,
			Brand: &proto.BrandInfoResponse{
				Id:   categoryBrand.Brands.ID,
				Name: categoryBrand.Brands.Name,
				Logo: categoryBrand.Brands.Logo,
			},
			Category: &proto.CategoryInfoResponse{
				Id:             categoryBrand.Category.ID,
				Name:           categoryBrand.Category.Name,
				Level:          categoryBrand.Category.Level,
				IsTab:          categoryBrand.Category.IsTab,
				ParentCategory: categoryBrand.Category.ParentCategoryID,
			},
		})
	}
	categoryBrandListResponse.Data = CategoryBrandResponse

	return &categoryBrandListResponse, nil
}

//GetCategoryBrandList 根据分类获品牌信息
func (g *GoodsServer) GetCategoryBrandList(c context.Context, req *proto.CategoryInfoRequest) (*proto.BrandListResponse, error) {
	var category model.Category
	var brandListResponse proto.BrandListResponse

	//查询当前分类
	if result := global.DB.Find(&category, req.Id).First(&category); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	fmt.Println("category:", category)

	//查询当前分类下的商品品牌
	var categoryBrands []model.GoodsCategoryBrand
	if result := global.DB.Preload("Brands").Where(&model.GoodsCategoryBrand{CategoryID: category.ID}).Find(&categoryBrands); result.RowsAffected > 0 {
		brandListResponse.Total = int32(result.RowsAffected)
	}

	var brandInfoResponse []*proto.BrandInfoResponse
	for _, categoryBrand := range categoryBrands {
		brandInfoResponse = append(brandInfoResponse, &proto.BrandInfoResponse{
			Id:   categoryBrand.Brands.ID,
			Name: categoryBrand.Brands.Name,
			Logo: categoryBrand.Brands.Logo,
		})
	}
	brandListResponse.Data = brandInfoResponse

	return &brandListResponse, nil
}

//CreateCategoryBrand 创建品牌分类 将商品分类和品牌建立关系
func (g *GoodsServer) CreateCategoryBrand(c context.Context, req *proto.CategoryBrandRequest) (*proto.CategoryBrandResponse, error) {
	var categroy model.Category

	//在创建品牌分类前，必须有对应商品分类,没有商品分类, 是不允许创建品牌分类的
	//查询商品分类
	if result := global.DB.First(&categroy, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	//查询品牌
	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categroyBrand := model.GoodsCategoryBrand{
		CategoryID: categroy.ID,
		BrandsID:   brand.ID,
	}

	//在对应表中插入记录
	global.DB.Save(&categroyBrand)

	return &proto.CategoryBrandResponse{
		Id: int32(categroyBrand.ID),
	}, nil
}

//DeleteCategoryBrand 删除品牌分类
func (g *GoodsServer) DeleteCategoryBrand(c context.Context, req *proto.CategoryBrandRequest) (*empty.Empty, error) {
	if result := global.DB.Delete(&model.GoodsCategoryBrand{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}

	return &empty.Empty{}, nil
}

//UpdateCategoryBrand 更新品牌分类
func (g *GoodsServer) UpdateCategoryBrand(c context.Context, req *proto.CategoryBrandRequest) (*empty.Empty, error) {
	var categoryBrand model.GoodsCategoryBrand
	if result := global.DB.First(&categoryBrand, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌分类不存在")
	}

	//查询商品分类
	var category model.Category
	if result := global.DB.First(&category, req.CategoryId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "商品分类不存在")
	}

	//查询品牌
	var brand model.Brands
	if result := global.DB.First(&brand, req.BrandId); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌不存在")
	}

	categoryBrand.CategoryID = req.CategoryId
	categoryBrand.BrandsID = req.BrandId

	global.DB.Save(&categoryBrand)

	return &empty.Empty{}, nil
}
