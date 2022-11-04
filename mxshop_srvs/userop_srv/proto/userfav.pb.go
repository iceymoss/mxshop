// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.13.0
// source: userfav.proto

package proto

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type UserFavRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	GoodsId int32 `protobuf:"varint,2,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
}

func (x *UserFavRequest) Reset() {
	*x = UserFavRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userfav_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFavRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFavRequest) ProtoMessage() {}

func (x *UserFavRequest) ProtoReflect() protoreflect.Message {
	mi := &file_userfav_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFavRequest.ProtoReflect.Descriptor instead.
func (*UserFavRequest) Descriptor() ([]byte, []int) {
	return file_userfav_proto_rawDescGZIP(), []int{0}
}

func (x *UserFavRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserFavRequest) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

type UserFavResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId  int32 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	GoodsId int32 `protobuf:"varint,2,opt,name=goodsId,proto3" json:"goodsId,omitempty"`
}

func (x *UserFavResponse) Reset() {
	*x = UserFavResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userfav_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFavResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFavResponse) ProtoMessage() {}

func (x *UserFavResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userfav_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFavResponse.ProtoReflect.Descriptor instead.
func (*UserFavResponse) Descriptor() ([]byte, []int) {
	return file_userfav_proto_rawDescGZIP(), []int{1}
}

func (x *UserFavResponse) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *UserFavResponse) GetGoodsId() int32 {
	if x != nil {
		return x.GoodsId
	}
	return 0
}

type UserFavListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int32              `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
	Data  []*UserFavResponse `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *UserFavListResponse) Reset() {
	*x = UserFavListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_userfav_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserFavListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserFavListResponse) ProtoMessage() {}

func (x *UserFavListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_userfav_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserFavListResponse.ProtoReflect.Descriptor instead.
func (*UserFavListResponse) Descriptor() ([]byte, []int) {
	return file_userfav_proto_rawDescGZIP(), []int{2}
}

func (x *UserFavListResponse) GetTotal() int32 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *UserFavListResponse) GetData() []*UserFavResponse {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_userfav_proto protoreflect.FileDescriptor

var file_userfav_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x75, 0x73, 0x65, 0x72, 0x66, 0x61, 0x76, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x42, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64,
	0x22, 0x43, 0x0a, 0x0f, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x67,
	0x6f, 0x6f, 0x64, 0x73, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x67, 0x6f,
	0x6f, 0x64, 0x73, 0x49, 0x64, 0x22, 0x51, 0x0a, 0x13, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x74, 0x6f, 0x74,
	0x61, 0x6c, 0x12, 0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x32, 0xec, 0x01, 0x0a, 0x07, 0x55, 0x73, 0x65,
	0x72, 0x46, 0x61, 0x76, 0x12, 0x33, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x46, 0x61, 0x76, 0x4c, 0x69,
	0x73, 0x74, 0x12, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x41, 0x64, 0x64,
	0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x12, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61,
	0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x12, 0x38, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61,
	0x76, 0x12, 0x0f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3b, 0x0a, 0x10, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12, 0x0f,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x46, 0x61, 0x76, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x0a, 0x5a, 0x08, 0x2f, 0x2e, 0x3b, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_userfav_proto_rawDescOnce sync.Once
	file_userfav_proto_rawDescData = file_userfav_proto_rawDesc
)

func file_userfav_proto_rawDescGZIP() []byte {
	file_userfav_proto_rawDescOnce.Do(func() {
		file_userfav_proto_rawDescData = protoimpl.X.CompressGZIP(file_userfav_proto_rawDescData)
	})
	return file_userfav_proto_rawDescData
}

var file_userfav_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_userfav_proto_goTypes = []interface{}{
	(*UserFavRequest)(nil),      // 0: UserFavRequest
	(*UserFavResponse)(nil),     // 1: UserFavResponse
	(*UserFavListResponse)(nil), // 2: UserFavListResponse
	(*empty.Empty)(nil),         // 3: google.protobuf.Empty
}
var file_userfav_proto_depIdxs = []int32{
	1, // 0: UserFavListResponse.data:type_name -> UserFavResponse
	0, // 1: UserFav.GetFavList:input_type -> UserFavRequest
	0, // 2: UserFav.AddUserFav:input_type -> UserFavRequest
	0, // 3: UserFav.DeleteUserFav:input_type -> UserFavRequest
	0, // 4: UserFav.GetUserFavDetail:input_type -> UserFavRequest
	2, // 5: UserFav.GetFavList:output_type -> UserFavListResponse
	3, // 6: UserFav.AddUserFav:output_type -> google.protobuf.Empty
	3, // 7: UserFav.DeleteUserFav:output_type -> google.protobuf.Empty
	3, // 8: UserFav.GetUserFavDetail:output_type -> google.protobuf.Empty
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_userfav_proto_init() }
func file_userfav_proto_init() {
	if File_userfav_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_userfav_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFavRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userfav_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFavResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_userfav_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserFavListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_userfav_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_userfav_proto_goTypes,
		DependencyIndexes: file_userfav_proto_depIdxs,
		MessageInfos:      file_userfav_proto_msgTypes,
	}.Build()
	File_userfav_proto = out.File
	file_userfav_proto_rawDesc = nil
	file_userfav_proto_goTypes = nil
	file_userfav_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserFavClient is the client API for UserFav service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserFavClient interface {
	GetFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error)
	AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type userFavClient struct {
	cc grpc.ClientConnInterface
}

func NewUserFavClient(cc grpc.ClientConnInterface) UserFavClient {
	return &userFavClient{cc}
}

func (c *userFavClient) GetFavList(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*UserFavListResponse, error) {
	out := new(UserFavListResponse)
	err := c.cc.Invoke(ctx, "/UserFav/GetFavList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFavClient) AddUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserFav/AddUserFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFavClient) DeleteUserFav(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserFav/DeleteUserFav", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userFavClient) GetUserFavDetail(ctx context.Context, in *UserFavRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/UserFav/GetUserFavDetail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserFavServer is the server API for UserFav service.
type UserFavServer interface {
	GetFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error)
	AddUserFav(context.Context, *UserFavRequest) (*empty.Empty, error)
	DeleteUserFav(context.Context, *UserFavRequest) (*empty.Empty, error)
	GetUserFavDetail(context.Context, *UserFavRequest) (*empty.Empty, error)
}

// UnimplementedUserFavServer can be embedded to have forward compatible implementations.
type UnimplementedUserFavServer struct {
}

func (*UnimplementedUserFavServer) GetFavList(context.Context, *UserFavRequest) (*UserFavListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFavList not implemented")
}
func (*UnimplementedUserFavServer) AddUserFav(context.Context, *UserFavRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddUserFav not implemented")
}
func (*UnimplementedUserFavServer) DeleteUserFav(context.Context, *UserFavRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserFav not implemented")
}
func (*UnimplementedUserFavServer) GetUserFavDetail(context.Context, *UserFavRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFavDetail not implemented")
}

func RegisterUserFavServer(s *grpc.Server, srv UserFavServer) {
	s.RegisterService(&_UserFav_serviceDesc, srv)
}

func _UserFav_GetFavList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).GetFavList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserFav/GetFavList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).GetFavList(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFav_AddUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).AddUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserFav/AddUserFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).AddUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFav_DeleteUserFav_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).DeleteUserFav(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserFav/DeleteUserFav",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).DeleteUserFav(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserFav_GetUserFavDetail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserFavRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserFavServer).GetUserFavDetail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UserFav/GetUserFavDetail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserFavServer).GetUserFavDetail(ctx, req.(*UserFavRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _UserFav_serviceDesc = grpc.ServiceDesc{
	ServiceName: "UserFav",
	HandlerType: (*UserFavServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetFavList",
			Handler:    _UserFav_GetFavList_Handler,
		},
		{
			MethodName: "AddUserFav",
			Handler:    _UserFav_AddUserFav_Handler,
		},
		{
			MethodName: "DeleteUserFav",
			Handler:    _UserFav_DeleteUserFav_Handler,
		},
		{
			MethodName: "GetUserFavDetail",
			Handler:    _UserFav_GetUserFavDetail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "userfav.proto",
}
