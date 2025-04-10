// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             (unknown)
// source: user/v1/user.proto

package userpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	UserService_GetUser_FullMethodName        = "/user.v1.UserService/GetUser"
	UserService_UpdateProfile_FullMethodName  = "/user.v1.UserService/UpdateProfile"
	UserService_BlockUser_FullMethodName      = "/user.v1.UserService/BlockUser"
	UserService_DeleteAccount_FullMethodName  = "/user.v1.UserService/DeleteAccount"
	UserService_SearchUsers_FullMethodName    = "/user.v1.UserService/SearchUsers"
	UserService_UpdateStatus_FullMethodName   = "/user.v1.UserService/UpdateStatus"
	UserService_GetOnlineUsers_FullMethodName = "/user.v1.UserService/GetOnlineUsers"
)

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*User, error)
	BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeleteAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error)
	UpdateStatus(ctx context.Context, in *StatusUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetOnlineUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OnlineUsersResponse, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_GetUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateProfile(ctx context.Context, in *UpdateProfileRequest, opts ...grpc.CallOption) (*User, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(User)
	err := c.cc.Invoke(ctx, UserService_UpdateProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) BlockUser(ctx context.Context, in *BlockUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_BlockUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteAccount(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_DeleteAccount_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SearchUsers(ctx context.Context, in *SearchUsersRequest, opts ...grpc.CallOption) (*SearchUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SearchUsersResponse)
	err := c.cc.Invoke(ctx, UserService_SearchUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateStatus(ctx context.Context, in *StatusUpdate, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, UserService_UpdateStatus_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetOnlineUsers(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*OnlineUsersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OnlineUsersResponse)
	err := c.cc.Invoke(ctx, UserService_GetOnlineUsers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility.
type UserServiceServer interface {
	GetUser(context.Context, *GetUserRequest) (*User, error)
	UpdateProfile(context.Context, *UpdateProfileRequest) (*User, error)
	BlockUser(context.Context, *BlockUserRequest) (*emptypb.Empty, error)
	DeleteAccount(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error)
	UpdateStatus(context.Context, *StatusUpdate) (*emptypb.Empty, error)
	GetOnlineUsers(context.Context, *emptypb.Empty) (*OnlineUsersResponse, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedUserServiceServer struct{}

func (UnimplementedUserServiceServer) GetUser(context.Context, *GetUserRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateProfile(context.Context, *UpdateProfileRequest) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateProfile not implemented")
}
func (UnimplementedUserServiceServer) BlockUser(context.Context, *BlockUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BlockUser not implemented")
}
func (UnimplementedUserServiceServer) DeleteAccount(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccount not implemented")
}
func (UnimplementedUserServiceServer) SearchUsers(context.Context, *SearchUsersRequest) (*SearchUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchUsers not implemented")
}
func (UnimplementedUserServiceServer) UpdateStatus(context.Context, *StatusUpdate) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateStatus not implemented")
}
func (UnimplementedUserServiceServer) GetOnlineUsers(context.Context, *emptypb.Empty) (*OnlineUsersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOnlineUsers not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}
func (UnimplementedUserServiceServer) testEmbeddedByValue()                     {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	// If the following call pancis, it indicates UnimplementedUserServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateProfile(ctx, req.(*UpdateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_BlockUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BlockUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).BlockUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_BlockUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).BlockUser(ctx, req.(*BlockUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_DeleteAccount_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteAccount(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SearchUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchUsersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SearchUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_SearchUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SearchUsers(ctx, req.(*SearchUsersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusUpdate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_UpdateStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateStatus(ctx, req.(*StatusUpdate))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetOnlineUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetOnlineUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserService_GetOnlineUsers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetOnlineUsers(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.v1.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _UserService_GetUser_Handler,
		},
		{
			MethodName: "UpdateProfile",
			Handler:    _UserService_UpdateProfile_Handler,
		},
		{
			MethodName: "BlockUser",
			Handler:    _UserService_BlockUser_Handler,
		},
		{
			MethodName: "DeleteAccount",
			Handler:    _UserService_DeleteAccount_Handler,
		},
		{
			MethodName: "SearchUsers",
			Handler:    _UserService_SearchUsers_Handler,
		},
		{
			MethodName: "UpdateStatus",
			Handler:    _UserService_UpdateStatus_Handler,
		},
		{
			MethodName: "GetOnlineUsers",
			Handler:    _UserService_GetOnlineUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user/v1/user.proto",
}
