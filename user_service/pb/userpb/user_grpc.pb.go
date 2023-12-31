// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: user.proto

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
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	GetUserData(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*UserData, error)
	GetAvailableDriver(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DriverData, error)
	CreateDriverData(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetDriverStatusOnline(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetDriverStatusOngoing(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	SetDriverStatusOffline(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	VerifyNewUser(ctx context.Context, in *UserCredential, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, "/user.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetUserData(ctx context.Context, in *EmailRequest, opts ...grpc.CallOption) (*UserData, error) {
	out := new(UserData)
	err := c.cc.Invoke(ctx, "/user.User/GetUserData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetAvailableDriver(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*DriverData, error) {
	out := new(DriverData)
	err := c.cc.Invoke(ctx, "/user.User/GetAvailableDriver", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) CreateDriverData(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.User/CreateDriverData", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetDriverStatusOnline(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.User/SetDriverStatusOnline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetDriverStatusOngoing(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.User/SetDriverStatusOngoing", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) SetDriverStatusOffline(ctx context.Context, in *DriverId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.User/SetDriverStatusOffline", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) VerifyNewUser(ctx context.Context, in *UserCredential, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/user.User/VerifyNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	GetUserData(context.Context, *EmailRequest) (*UserData, error)
	GetAvailableDriver(context.Context, *emptypb.Empty) (*DriverData, error)
	CreateDriverData(context.Context, *DriverId) (*emptypb.Empty, error)
	SetDriverStatusOnline(context.Context, *DriverId) (*emptypb.Empty, error)
	SetDriverStatusOngoing(context.Context, *DriverId) (*emptypb.Empty, error)
	SetDriverStatusOffline(context.Context, *DriverId) (*emptypb.Empty, error)
	VerifyNewUser(context.Context, *UserCredential) (*emptypb.Empty, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) GetUserData(context.Context, *EmailRequest) (*UserData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserData not implemented")
}
func (UnimplementedUserServer) GetAvailableDriver(context.Context, *emptypb.Empty) (*DriverData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAvailableDriver not implemented")
}
func (UnimplementedUserServer) CreateDriverData(context.Context, *DriverId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDriverData not implemented")
}
func (UnimplementedUserServer) SetDriverStatusOnline(context.Context, *DriverId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDriverStatusOnline not implemented")
}
func (UnimplementedUserServer) SetDriverStatusOngoing(context.Context, *DriverId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDriverStatusOngoing not implemented")
}
func (UnimplementedUserServer) SetDriverStatusOffline(context.Context, *DriverId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetDriverStatusOffline not implemented")
}
func (UnimplementedUserServer) VerifyNewUser(context.Context, *UserCredential) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyNewUser not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetUserData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetUserData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetUserData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetUserData(ctx, req.(*EmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetAvailableDriver_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAvailableDriver(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/GetAvailableDriver",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAvailableDriver(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_CreateDriverData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).CreateDriverData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/CreateDriverData",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).CreateDriverData(ctx, req.(*DriverId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetDriverStatusOnline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetDriverStatusOnline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SetDriverStatusOnline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetDriverStatusOnline(ctx, req.(*DriverId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetDriverStatusOngoing_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetDriverStatusOngoing(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SetDriverStatusOngoing",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetDriverStatusOngoing(ctx, req.(*DriverId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_SetDriverStatusOffline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DriverId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).SetDriverStatusOffline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/SetDriverStatusOffline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).SetDriverStatusOffline(ctx, req.(*DriverId))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_VerifyNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).VerifyNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.User/VerifyNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).VerifyNewUser(ctx, req.(*UserCredential))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "GetUserData",
			Handler:    _User_GetUserData_Handler,
		},
		{
			MethodName: "GetAvailableDriver",
			Handler:    _User_GetAvailableDriver_Handler,
		},
		{
			MethodName: "CreateDriverData",
			Handler:    _User_CreateDriverData_Handler,
		},
		{
			MethodName: "SetDriverStatusOnline",
			Handler:    _User_SetDriverStatusOnline_Handler,
		},
		{
			MethodName: "SetDriverStatusOngoing",
			Handler:    _User_SetDriverStatusOngoing_Handler,
		},
		{
			MethodName: "SetDriverStatusOffline",
			Handler:    _User_SetDriverStatusOffline_Handler,
		},
		{
			MethodName: "VerifyNewUser",
			Handler:    _User_VerifyNewUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
