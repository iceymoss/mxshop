package handler

import (
	"context"
	"mxshop_srvs/goods_srv/global"
	"mxshop_srvs/goods_srv/model"
	"mxshop_srvs/goods_srv/proto"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"gorm.io/gorm"
)

//Paginate 将数据进行分页
func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func ModelToRsponse(brand model.Brands) proto.BrandInfoResponse {
	return proto.BrandInfoResponse{
		Id:   brand.ID,
		Name: brand.Name,
		Logo: brand.Logo,
	}
}

//BrandList 获取品牌列表
func (g *GoodsServer) BrandList(ctx context.Context, req *proto.BrandFilterRequest) (*proto.BrandListResponse, error) {
	var BrandList proto.BrandListResponse
	var brands []model.Brands
	var BrandInfo []*proto.BrandInfoResponse

	result := global.DB.Find(&brands)
	if result.Error != nil {
		return nil, result.Error
	}

	var total int64
	global.DB.Model(&brands).Count(&total)

	global.DB.Scopes(Paginate(int(req.Pages), int(req.PagePerNums))).Find(&brands)

	for _, brand := range brands {
		BrandInfo = append(BrandInfo, &proto.BrandInfoResponse{
			Id:   brand.ID,
			Name: brand.Name,
			Logo: brand.Logo,
		})
	}
	BrandList.Total = int32(total)
	BrandList.Data = BrandInfo
	return &BrandList, nil
}

//CreateBrand 新增品牌
func (g *GoodsServer) CreateBrand(ctx context.Context, req *proto.BrandRequest) (*proto.BrandInfoResponse, error) {
	if result := global.DB.Where("name=?", req.Name).First(&model.Brands{}); result.RowsAffected == 1 {
		return nil, status.Errorf(codes.InvalidArgument, "品牌已存在")
	}

	//if res := global.DB.First(&model.Brands{}); res.RowsAffected == 1 {
	//	return nil, status.Errorf(codes.InvalidArgument, "品牌也存在")
	//}

	var brand model.Brands
	brand.Name = req.Name
	brand.Logo = req.Logo

	result := global.DB.Create(&brand)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	BrandResponse := ModelToRsponse(brand)
	return &BrandResponse, nil
}

//DeleteBrand 删除品牌
func (g *GoodsServer) DeleteBrand(ctx context.Context, req *proto.BrandRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(&model.Brands{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}
	return &emptypb.Empty{}, nil
}

//UpdateBrand 更新品牌
func (g *GoodsServer) UpdateBrand(ctx context.Context, req *proto.BrandRequest) (*empty.Empty, error) {
	var brand model.Brands
	if result := global.DB.First(&brand, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "品牌不存在")
	}

	if brand.Name != "" {
		brand.Name = req.Name
	}

	if brand.Logo != "" {
		brand.Logo = req.Logo
	}

	global.DB.Save(&brand)
	return &emptypb.Empty{}, nil
}
