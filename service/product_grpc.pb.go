// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.5
// source: product.proto

package service

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

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProductStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	UpdataProductStockClientStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_UpdataProductStockClientStreamClient, error)
	GetProductStockServerStream(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (ProdService_GetProductStockServerStreamClient, error)
	SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_SayHelloStreamClient, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProductStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/service.ProdService/GetProductStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *prodServiceClient) UpdataProductStockClientStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_UpdataProductStockClientStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProdService_ServiceDesc.Streams[0], "/service.ProdService/UpdataProductStockClientStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &prodServiceUpdataProductStockClientStreamClient{stream}
	return x, nil
}

type ProdService_UpdataProductStockClientStreamClient interface {
	Send(*ProductRequest) error
	CloseAndRecv() (*ProductResponse, error)
	grpc.ClientStream
}

type prodServiceUpdataProductStockClientStreamClient struct {
	grpc.ClientStream
}

func (x *prodServiceUpdataProductStockClientStreamClient) Send(m *ProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *prodServiceUpdataProductStockClientStreamClient) CloseAndRecv() (*ProductResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *prodServiceClient) GetProductStockServerStream(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (ProdService_GetProductStockServerStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProdService_ServiceDesc.Streams[1], "/service.ProdService/GetProductStockServerStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &prodServiceGetProductStockServerStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProdService_GetProductStockServerStreamClient interface {
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type prodServiceGetProductStockServerStreamClient struct {
	grpc.ClientStream
}

func (x *prodServiceGetProductStockServerStreamClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *prodServiceClient) SayHelloStream(ctx context.Context, opts ...grpc.CallOption) (ProdService_SayHelloStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &ProdService_ServiceDesc.Streams[2], "/service.ProdService/SayHelloStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &prodServiceSayHelloStreamClient{stream}
	return x, nil
}

type ProdService_SayHelloStreamClient interface {
	Send(*ProductRequest) error
	Recv() (*ProductResponse, error)
	grpc.ClientStream
}

type prodServiceSayHelloStreamClient struct {
	grpc.ClientStream
}

func (x *prodServiceSayHelloStreamClient) Send(m *ProductRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *prodServiceSayHelloStreamClient) Recv() (*ProductResponse, error) {
	m := new(ProductResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProdServiceServer is the server API for ProdService service.
// All implementations must embed UnimplementedProdServiceServer
// for forward compatibility
type ProdServiceServer interface {
	GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error)
	UpdataProductStockClientStream(ProdService_UpdataProductStockClientStreamServer) error
	GetProductStockServerStream(*ProductRequest, ProdService_GetProductStockServerStreamServer) error
	SayHelloStream(ProdService_SayHelloStreamServer) error
	mustEmbedUnimplementedProdServiceServer()
}

// UnimplementedProdServiceServer must be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (UnimplementedProdServiceServer) GetProductStock(context.Context, *ProductRequest) (*ProductResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProductStock not implemented")
}
func (UnimplementedProdServiceServer) UpdataProductStockClientStream(ProdService_UpdataProductStockClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdataProductStockClientStream not implemented")
}
func (UnimplementedProdServiceServer) GetProductStockServerStream(*ProductRequest, ProdService_GetProductStockServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method GetProductStockServerStream not implemented")
}
func (UnimplementedProdServiceServer) SayHelloStream(ProdService_SayHelloStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloStream not implemented")
}
func (UnimplementedProdServiceServer) mustEmbedUnimplementedProdServiceServer() {}

// UnsafeProdServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ProdServiceServer will
// result in compilation errors.
type UnsafeProdServiceServer interface {
	mustEmbedUnimplementedProdServiceServer()
}

func RegisterProdServiceServer(s grpc.ServiceRegistrar, srv ProdServiceServer) {
	s.RegisterService(&ProdService_ServiceDesc, srv)
}

func _ProdService_GetProductStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProductStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.ProdService/GetProductStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProductStock(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ProdService_UpdataProductStockClientStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProdServiceServer).UpdataProductStockClientStream(&prodServiceUpdataProductStockClientStreamServer{stream})
}

type ProdService_UpdataProductStockClientStreamServer interface {
	SendAndClose(*ProductResponse) error
	Recv() (*ProductRequest, error)
	grpc.ServerStream
}

type prodServiceUpdataProductStockClientStreamServer struct {
	grpc.ServerStream
}

func (x *prodServiceUpdataProductStockClientStreamServer) SendAndClose(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *prodServiceUpdataProductStockClientStreamServer) Recv() (*ProductRequest, error) {
	m := new(ProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ProdService_GetProductStockServerStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProductRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProdServiceServer).GetProductStockServerStream(m, &prodServiceGetProductStockServerStreamServer{stream})
}

type ProdService_GetProductStockServerStreamServer interface {
	Send(*ProductResponse) error
	grpc.ServerStream
}

type prodServiceGetProductStockServerStreamServer struct {
	grpc.ServerStream
}

func (x *prodServiceGetProductStockServerStreamServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ProdService_SayHelloStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ProdServiceServer).SayHelloStream(&prodServiceSayHelloStreamServer{stream})
}

type ProdService_SayHelloStreamServer interface {
	Send(*ProductResponse) error
	Recv() (*ProductRequest, error)
	grpc.ServerStream
}

type prodServiceSayHelloStreamServer struct {
	grpc.ServerStream
}

func (x *prodServiceSayHelloStreamServer) Send(m *ProductResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *prodServiceSayHelloStreamServer) Recv() (*ProductRequest, error) {
	m := new(ProductRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProdService_ServiceDesc is the grpc.ServiceDesc for ProdService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ProdService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProductStock",
			Handler:    _ProdService_GetProductStock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UpdataProductStockClientStream",
			Handler:       _ProdService_UpdataProductStockClientStream_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "GetProductStockServerStream",
			Handler:       _ProdService_GetProductStockServerStream_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SayHelloStream",
			Handler:       _ProdService_SayHelloStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "product.proto",
}