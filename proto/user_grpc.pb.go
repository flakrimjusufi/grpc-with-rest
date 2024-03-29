// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: proto/user.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUserByName(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	UpdateUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Message, error)
	ListUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*ListUser, error)
	GetUserByName(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	GetUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error)
	SayHello(ctx context.Context, in *User, opts ...grpc.CallOption) (*Message, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) CreateUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/CreateUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserByName(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) UpdateUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/UpdateUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) DeleteUser(ctx context.Context, in *User, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/user.UserService/DeleteUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) ListUsers(ctx context.Context, in *User, opts ...grpc.CallOption) (*ListUser, error) {
	out := new(ListUser)
	err := c.cc.Invoke(ctx, "/user.UserService/ListUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByName(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByID(ctx context.Context, in *User, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/user.UserService/GetUserByID", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) SayHello(ctx context.Context, in *User, opts ...grpc.CallOption) (*Message, error) {
	out := new(Message)
	err := c.cc.Invoke(ctx, "/user.UserService/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	CreateUser(context.Context, *User) (*User, error)
	UpdateUserByName(context.Context, *User) (*User, error)
	UpdateUserByID(context.Context, *User) (*User, error)
	DeleteUser(context.Context, *User) (*Message, error)
	ListUsers(context.Context, *User) (*ListUser, error)
	GetUserByName(context.Context, *User) (*User, error)
	GetUserByID(context.Context, *User) (*User, error)
	SayHello(context.Context, *User) (*Message, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) CreateUser(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateUser not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserByName(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserByName not implemented")
}
func (UnimplementedUserServiceServer) UpdateUserByID(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateUserByID not implemented")
}
func (UnimplementedUserServiceServer) DeleteUser(context.Context, *User) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUser not implemented")
}
func (UnimplementedUserServiceServer) ListUsers(context.Context, *User) (*ListUser, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListUsers not implemented")
}
func (UnimplementedUserServiceServer) GetUserByName(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByName not implemented")
}
func (UnimplementedUserServiceServer) GetUserByID(context.Context, *User) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByID not implemented")
}
func (UnimplementedUserServiceServer) SayHello(context.Context, *User) (*Message, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_CreateUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).CreateUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/CreateUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).CreateUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserByName(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_UpdateUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).UpdateUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/UpdateUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).UpdateUserByID(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_DeleteUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).DeleteUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/DeleteUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).DeleteUser(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_ListUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).ListUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/ListUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).ListUsers(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByName(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByID_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByID(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/GetUserByID",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByID(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(User)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.UserService/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).SayHello(ctx, req.(*User))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateUser",
			Handler:    _UserService_CreateUser_Handler,
		},
		{
			MethodName: "UpdateUserByName",
			Handler:    _UserService_UpdateUserByName_Handler,
		},
		{
			MethodName: "UpdateUserByID",
			Handler:    _UserService_UpdateUserByID_Handler,
		},
		{
			MethodName: "DeleteUser",
			Handler:    _UserService_DeleteUser_Handler,
		},
		{
			MethodName: "ListUsers",
			Handler:    _UserService_ListUsers_Handler,
		},
		{
			MethodName: "GetUserByName",
			Handler:    _UserService_GetUserByName_Handler,
		},
		{
			MethodName: "GetUserByID",
			Handler:    _UserService_GetUserByID_Handler,
		},
		{
			MethodName: "SayHello",
			Handler:    _UserService_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}

// CreditCardServiceClient is the client API for CreditCardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CreditCardServiceClient interface {
	CreditCards(ctx context.Context, in *CreditCard, opts ...grpc.CallOption) (*ListCreditCards, error)
	GetCreditCardByUserName(ctx context.Context, in *CreditCard, opts ...grpc.CallOption) (*CreditCard, error)
	CreateCreditCardApplication(ctx context.Context, in *CreditCardApplication, opts ...grpc.CallOption) (*CreditCardApplication, error)
	GetCreditCardApplicationByName(ctx context.Context, in *CreditCardApplication, opts ...grpc.CallOption) (*CreditCardApplication, error)
}

type creditCardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCreditCardServiceClient(cc grpc.ClientConnInterface) CreditCardServiceClient {
	return &creditCardServiceClient{cc}
}

func (c *creditCardServiceClient) CreditCards(ctx context.Context, in *CreditCard, opts ...grpc.CallOption) (*ListCreditCards, error) {
	out := new(ListCreditCards)
	err := c.cc.Invoke(ctx, "/user.CreditCardService/CreditCards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditCardServiceClient) GetCreditCardByUserName(ctx context.Context, in *CreditCard, opts ...grpc.CallOption) (*CreditCard, error) {
	out := new(CreditCard)
	err := c.cc.Invoke(ctx, "/user.CreditCardService/GetCreditCardByUserName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditCardServiceClient) CreateCreditCardApplication(ctx context.Context, in *CreditCardApplication, opts ...grpc.CallOption) (*CreditCardApplication, error) {
	out := new(CreditCardApplication)
	err := c.cc.Invoke(ctx, "/user.CreditCardService/CreateCreditCardApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *creditCardServiceClient) GetCreditCardApplicationByName(ctx context.Context, in *CreditCardApplication, opts ...grpc.CallOption) (*CreditCardApplication, error) {
	out := new(CreditCardApplication)
	err := c.cc.Invoke(ctx, "/user.CreditCardService/GetCreditCardApplicationByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CreditCardServiceServer is the server API for CreditCardService service.
// All implementations must embed UnimplementedCreditCardServiceServer
// for forward compatibility
type CreditCardServiceServer interface {
	CreditCards(context.Context, *CreditCard) (*ListCreditCards, error)
	GetCreditCardByUserName(context.Context, *CreditCard) (*CreditCard, error)
	CreateCreditCardApplication(context.Context, *CreditCardApplication) (*CreditCardApplication, error)
	GetCreditCardApplicationByName(context.Context, *CreditCardApplication) (*CreditCardApplication, error)
	mustEmbedUnimplementedCreditCardServiceServer()
}

// UnimplementedCreditCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCreditCardServiceServer struct {
}

func (UnimplementedCreditCardServiceServer) CreditCards(context.Context, *CreditCard) (*ListCreditCards, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreditCards not implemented")
}
func (UnimplementedCreditCardServiceServer) GetCreditCardByUserName(context.Context, *CreditCard) (*CreditCard, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreditCardByUserName not implemented")
}
func (UnimplementedCreditCardServiceServer) CreateCreditCardApplication(context.Context, *CreditCardApplication) (*CreditCardApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateCreditCardApplication not implemented")
}
func (UnimplementedCreditCardServiceServer) GetCreditCardApplicationByName(context.Context, *CreditCardApplication) (*CreditCardApplication, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCreditCardApplicationByName not implemented")
}
func (UnimplementedCreditCardServiceServer) mustEmbedUnimplementedCreditCardServiceServer() {}

// UnsafeCreditCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CreditCardServiceServer will
// result in compilation errors.
type UnsafeCreditCardServiceServer interface {
	mustEmbedUnimplementedCreditCardServiceServer()
}

func RegisterCreditCardServiceServer(s grpc.ServiceRegistrar, srv CreditCardServiceServer) {
	s.RegisterService(&CreditCardService_ServiceDesc, srv)
}

func _CreditCardService_CreditCards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardServiceServer).CreditCards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CreditCardService/CreditCards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardServiceServer).CreditCards(ctx, req.(*CreditCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditCardService_GetCreditCardByUserName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCard)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardServiceServer).GetCreditCardByUserName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CreditCardService/GetCreditCardByUserName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardServiceServer).GetCreditCardByUserName(ctx, req.(*CreditCard))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditCardService_CreateCreditCardApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCardApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardServiceServer).CreateCreditCardApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CreditCardService/CreateCreditCardApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardServiceServer).CreateCreditCardApplication(ctx, req.(*CreditCardApplication))
	}
	return interceptor(ctx, in, info, handler)
}

func _CreditCardService_GetCreditCardApplicationByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreditCardApplication)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CreditCardServiceServer).GetCreditCardApplicationByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/user.CreditCardService/GetCreditCardApplicationByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CreditCardServiceServer).GetCreditCardApplicationByName(ctx, req.(*CreditCardApplication))
	}
	return interceptor(ctx, in, info, handler)
}

// CreditCardService_ServiceDesc is the grpc.ServiceDesc for CreditCardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CreditCardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.CreditCardService",
	HandlerType: (*CreditCardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreditCards",
			Handler:    _CreditCardService_CreditCards_Handler,
		},
		{
			MethodName: "GetCreditCardByUserName",
			Handler:    _CreditCardService_GetCreditCardByUserName_Handler,
		},
		{
			MethodName: "CreateCreditCardApplication",
			Handler:    _CreditCardService_CreateCreditCardApplication_Handler,
		},
		{
			MethodName: "GetCreditCardApplicationByName",
			Handler:    _CreditCardService_GetCreditCardApplicationByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/user.proto",
}
