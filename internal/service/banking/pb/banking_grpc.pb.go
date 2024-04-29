// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/banking.proto

package pb

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

// BankingClient is the client API for Banking service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankingClient interface {
	Transfer(ctx context.Context, in *TransferIn, opts ...grpc.CallOption) (*TransferOut, error)
	CreateAccount(ctx context.Context, in *CreateAccountIn, opts ...grpc.CallOption) (*CreateAccountOut, error)
	InfoAccount(ctx context.Context, in *InfoAccountIn, opts ...grpc.CallOption) (*InfoAccountOut, error)
}

type bankingClient struct {
	cc grpc.ClientConnInterface
}

func NewBankingClient(cc grpc.ClientConnInterface) BankingClient {
	return &bankingClient{cc}
}

func (c *bankingClient) Transfer(ctx context.Context, in *TransferIn, opts ...grpc.CallOption) (*TransferOut, error) {
	out := new(TransferOut)
	err := c.cc.Invoke(ctx, "/banking.Banking/Transfer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingClient) CreateAccount(ctx context.Context, in *CreateAccountIn, opts ...grpc.CallOption) (*CreateAccountOut, error) {
	out := new(CreateAccountOut)
	err := c.cc.Invoke(ctx, "/banking.Banking/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankingClient) InfoAccount(ctx context.Context, in *InfoAccountIn, opts ...grpc.CallOption) (*InfoAccountOut, error) {
	out := new(InfoAccountOut)
	err := c.cc.Invoke(ctx, "/banking.Banking/InfoAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BankingServer is the server API for Banking service.
// All implementations must embed UnimplementedBankingServer
// for forward compatibility
type BankingServer interface {
	Transfer(context.Context, *TransferIn) (*TransferOut, error)
	CreateAccount(context.Context, *CreateAccountIn) (*CreateAccountOut, error)
	InfoAccount(context.Context, *InfoAccountIn) (*InfoAccountOut, error)
	mustEmbedUnimplementedBankingServer()
}

// UnimplementedBankingServer must be embedded to have forward compatible implementations.
type UnimplementedBankingServer struct {
}

func (UnimplementedBankingServer) Transfer(context.Context, *TransferIn) (*TransferOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Transfer not implemented")
}
func (UnimplementedBankingServer) CreateAccount(context.Context, *CreateAccountIn) (*CreateAccountOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedBankingServer) InfoAccount(context.Context, *InfoAccountIn) (*InfoAccountOut, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InfoAccount not implemented")
}
func (UnimplementedBankingServer) mustEmbedUnimplementedBankingServer() {}

// UnsafeBankingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankingServer will
// result in compilation errors.
type UnsafeBankingServer interface {
	mustEmbedUnimplementedBankingServer()
}

func RegisterBankingServer(s grpc.ServiceRegistrar, srv BankingServer) {
	s.RegisterService(&Banking_ServiceDesc, srv)
}

func _Banking_Transfer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServer).Transfer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.Banking/Transfer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServer).Transfer(ctx, req.(*TransferIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banking_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.Banking/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServer).CreateAccount(ctx, req.(*CreateAccountIn))
	}
	return interceptor(ctx, in, info, handler)
}

func _Banking_InfoAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoAccountIn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankingServer).InfoAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/banking.Banking/InfoAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankingServer).InfoAccount(ctx, req.(*InfoAccountIn))
	}
	return interceptor(ctx, in, info, handler)
}

// Banking_ServiceDesc is the grpc.ServiceDesc for Banking service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Banking_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "banking.Banking",
	HandlerType: (*BankingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Transfer",
			Handler:    _Banking_Transfer_Handler,
		},
		{
			MethodName: "CreateAccount",
			Handler:    _Banking_CreateAccount_Handler,
		},
		{
			MethodName: "InfoAccount",
			Handler:    _Banking_InfoAccount_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/banking.proto",
}
