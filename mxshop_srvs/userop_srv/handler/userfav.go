package handler

import (
	"context"

	"mxshop_srvs/userop_srv/global"
	"mxshop_srvs/userop_srv/model"
	"mxshop_srvs/userop_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (u *UserOpServer) GetFavList(ctx context.Context, req *proto.UserFavRequest) (*proto.UserFavListResponse, error) {
	var rsp proto.UserFavListResponse
	var userFavs []model.UserFav
	var userFavList []*proto.UserFavResponse

	result := global.DB.Where(&model.UserFav{User: req.UserId, Goods: req.GoodsId}).Find(&userFavs)
	rsp.Total = int32(result.RowsAffected)

	for _, userFav := range userFavs {
		userFavList = append(userFavList, &proto.UserFavResponse{
			UserId:  userFav.User,
			GoodsId: userFav.Goods,
		})
	}

	rsp.Data = userFavList

	return &rsp, nil
}

func (u *UserOpServer) AddUserFav(ctx context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	var userFav model.UserFav

	userFav.User = req.UserId
	userFav.Goods = req.GoodsId

	global.DB.Save(&userFav)

	return &emptypb.Empty{}, nil
}

//DeleteUserFav 移除收藏
func (*UserOpServer) DeleteUserFav(ctx context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	//Unscoped 设置了唯一索引使用硬删除
	if result := global.DB.Unscoped().Where("goods=? and user=?", req.GoodsId, req.UserId).Delete(&model.UserFav{}); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}
	return &emptypb.Empty{}, nil
}

func (*UserOpServer) GetUserFavDetail(ctx context.Context, req *proto.UserFavRequest) (*emptypb.Empty, error) {
	var userfav model.UserFav
	if result := global.DB.Where("goods=? and user=?", req.GoodsId, req.UserId).Find(&userfav); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "收藏记录不存在")
	}
	return &emptypb.Empty{}, nil
}
