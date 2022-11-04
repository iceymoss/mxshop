package handler

import (
	"context"

	"mxshop_srvs/userop_srv/global"
	"mxshop_srvs/userop_srv/model"
	"mxshop_srvs/userop_srv/proto"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//GetAddressList 获取收货地址列表
func (u *UserOpServer) GetAddressList(ctx context.Context, req *proto.GetAddressRequest) (*proto.AddressListResponse, error) {
	var addressList []model.Address
	var AddressListResponse proto.AddressListResponse
	if result := global.DB.Where(&model.Address{User: req.UserId}).Find(&addressList); result.RowsAffected != 0 {
		AddressListResponse.Total = int32(result.RowsAffected)
	}

	for _, address := range addressList {
		AddressListResponse.Data = append(AddressListResponse.Data, &proto.AddressResponse{
			Id:           address.ID,
			UserId:       address.User,
			Province:     address.Province,
			City:         address.City,
			District:     address.District,
			Address:      address.Address,
			SignerMobile: address.SignerMobile,
			SignerName:   address.SignerName,
		})
	}
	return &AddressListResponse, nil
}

//CreateAddress 新建收货地址
func (u *UserOpServer) CreateAddress(ctx context.Context, req *proto.GetAddressRequest) (*proto.AddressResponse, error) {
	var address model.Address
	address.User = req.UserId
	address.Province = req.Province
	address.City = req.City
	address.District = req.District
	address.Address = req.Address
	address.SignerName = req.SignerName
	address.SignerMobile = req.SignerMobile

	if result := global.DB.Save(&address); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.Internal, "新建地址失败")
	}
	return &proto.AddressResponse{Id: address.ID}, nil
}

//UpdateAddress 更新收货地址
func (u *UserOpServer) UpdateAddress(ctx context.Context, req *proto.GetAddressRequest) (*empty.Empty, error) {
	var address model.Address
	if result := global.DB.Where("id=? and user=?", req.Id, req.UserId).First(&address); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收货地址记录不存在")
	}

	if req.Province != "" {
		address.Province = req.Province
	}
	if req.City != "" {
		address.City = req.City
	}
	if req.District != "" {
		address.District = req.District
	}
	if req.Address != "" {
		address.Address = req.Address
	}
	if req.SignerMobile != "" {
		address.SignerMobile = req.SignerMobile
	}
	if req.SignerName != "" {
		address.SignerName = req.SignerName
	}

	global.DB.Save(&address)

	return &empty.Empty{}, nil
}

//DeleteAddress 删除收货地址
func (u *UserOpServer) DeleteAddress(ctx context.Context, req *proto.GetAddressRequest) (*empty.Empty, error) {
	if result := global.DB.Where("id=? and user=?", req.Id, req.UserId).Delete(&model.Address{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收货地址记录不存在")
	}
	return &empty.Empty{}, nil
}
