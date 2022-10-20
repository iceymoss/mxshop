package handler

import (
	"context"
	"crypto/sha512"
	"fmt"
	"strings"
	"time"

	"mxshop_srvs/user_srv/global"
	"mxshop_srvs/user_srv/model"
	"mxshop_srvs/user_srv/proto"

	"github.com/anaskhan96/go-password-encoder"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

//UserServer 用户服务
type UserServer struct{}

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

//ModelToResponse 赋值
func ModelToResponse(user model.User) proto.UserInfoResponse {
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		Password: user.PassWord,
		Mobile:   user.Mobile,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
	}
	if user.Birthday != nil {
		userInfoRsp.Birthday = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

//GetUserInfoList 获取用户信息列表
func (s *UserServer) GetUserInfoList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	fmt.Println("用户服务")
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)
	global.DB.Scopes(Paginate(int(req.Pn), int(req.Psize))).Find(&users)

	for _, user := range users {
		userInfoRsp := ModelToResponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

//GetUserByMobile 根据Mobile获取用户
func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

//GetUserById 通过id查询用户
func (s *UserServer) GetUserById(ctx context.Context, req *proto.IdRequest) (*proto.UserInfoResponse, error) {
	var user model.User
	//主键查询
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}
	if result.Error != nil {
		return nil, result.Error
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

//CreateUser 创建用户
func (s *UserServer) CreateUser(ctx context.Context, req *proto.CreateUserInfo) (*proto.UserInfoResponse, error) {
	//新建用户
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 1 {
		return nil, status.Errorf(codes.AlreadyExists, "用户已存在")
	}

	user.Mobile = req.Mobile
	user.NickName = req.NickName

	//密码加密
	options := &password.Options{16, 100, 32, sha512.New}
	salt, encodedPwd := password.Encode(req.Password, options)
	user.PassWord = fmt.Sprintf("$pbkdf2-sha512$%s$%s", salt, encodedPwd)

	result = global.DB.Create(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}

	userInfoRsp := ModelToResponse(user)
	return &userInfoRsp, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *proto.UpdateUserInfo) (*empty.Empty, error) {
	//个人中心更新用户
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "用户不存在")
	}

	birthDay := time.Unix(int64(req.Birthday), 0)
	user.NickName = req.NickName
	user.Birthday = &birthDay
	user.Gender = req.Gender

	result = global.DB.Save(&user)
	if result.Error != nil {
		return nil, status.Errorf(codes.Internal, result.Error.Error())
	}
	return &empty.Empty{}, nil
}

//CheckPassWord 校验密码
func (s *UserServer) CheckPassWord(ctx context.Context, req *proto.PasswordCheckInfo) (*proto.CheckResponse, error) {
	Passwordinfo := strings.Split(req.EncryptedPassword, "$")
	fmt.Println(Passwordinfo)
	options := &password.Options{16, 100, 32, sha512.New}
	check := password.Verify(req.Password, Passwordinfo[2], Passwordinfo[3], options)
	return &proto.CheckResponse{Success: check}, nil
}
