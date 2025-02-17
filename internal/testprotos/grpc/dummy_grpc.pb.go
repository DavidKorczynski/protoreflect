// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.7
// source: grpc/dummy.proto

package grpc

import (
	context "context"
	testprotos "github.com/jhump/protoreflect/internal/testprotos"
	pkg "github.com/jhump/protoreflect/internal/testprotos/pkg"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DummyServiceClient is the client API for DummyService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DummyServiceClient interface {
	DoSomething(ctx context.Context, in *DummyRequest, opts ...grpc.CallOption) (*pkg.Bar, error)
	DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (DummyService_DoSomethingElseClient, error)
	DoSomethingAgain(ctx context.Context, in *pkg.Bar, opts ...grpc.CallOption) (DummyService_DoSomethingAgainClient, error)
	DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (DummyService_DoSomethingForeverClient, error)
}

type dummyServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDummyServiceClient(cc grpc.ClientConnInterface) DummyServiceClient {
	return &dummyServiceClient{cc}
}

func (c *dummyServiceClient) DoSomething(ctx context.Context, in *DummyRequest, opts ...grpc.CallOption) (*pkg.Bar, error) {
	out := new(pkg.Bar)
	err := c.cc.Invoke(ctx, "/testprotos.DummyService/DoSomething", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dummyServiceClient) DoSomethingElse(ctx context.Context, opts ...grpc.CallOption) (DummyService_DoSomethingElseClient, error) {
	stream, err := c.cc.NewStream(ctx, &DummyService_ServiceDesc.Streams[0], "/testprotos.DummyService/DoSomethingElse", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyServiceDoSomethingElseClient{stream}
	return x, nil
}

type DummyService_DoSomethingElseClient interface {
	Send(*testprotos.TestMessage) error
	CloseAndRecv() (*DummyResponse, error)
	grpc.ClientStream
}

type dummyServiceDoSomethingElseClient struct {
	grpc.ClientStream
}

func (x *dummyServiceDoSomethingElseClient) Send(m *testprotos.TestMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyServiceDoSomethingElseClient) CloseAndRecv() (*DummyResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(DummyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyServiceClient) DoSomethingAgain(ctx context.Context, in *pkg.Bar, opts ...grpc.CallOption) (DummyService_DoSomethingAgainClient, error) {
	stream, err := c.cc.NewStream(ctx, &DummyService_ServiceDesc.Streams[1], "/testprotos.DummyService/DoSomethingAgain", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyServiceDoSomethingAgainClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type DummyService_DoSomethingAgainClient interface {
	Recv() (*testprotos.AnotherTestMessage, error)
	grpc.ClientStream
}

type dummyServiceDoSomethingAgainClient struct {
	grpc.ClientStream
}

func (x *dummyServiceDoSomethingAgainClient) Recv() (*testprotos.AnotherTestMessage, error) {
	m := new(testprotos.AnotherTestMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *dummyServiceClient) DoSomethingForever(ctx context.Context, opts ...grpc.CallOption) (DummyService_DoSomethingForeverClient, error) {
	stream, err := c.cc.NewStream(ctx, &DummyService_ServiceDesc.Streams[2], "/testprotos.DummyService/DoSomethingForever", opts...)
	if err != nil {
		return nil, err
	}
	x := &dummyServiceDoSomethingForeverClient{stream}
	return x, nil
}

type DummyService_DoSomethingForeverClient interface {
	Send(*DummyRequest) error
	Recv() (*DummyResponse, error)
	grpc.ClientStream
}

type dummyServiceDoSomethingForeverClient struct {
	grpc.ClientStream
}

func (x *dummyServiceDoSomethingForeverClient) Send(m *DummyRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *dummyServiceDoSomethingForeverClient) Recv() (*DummyResponse, error) {
	m := new(DummyResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DummyServiceServer is the server API for DummyService service.
// All implementations must embed UnimplementedDummyServiceServer
// for forward compatibility
type DummyServiceServer interface {
	DoSomething(context.Context, *DummyRequest) (*pkg.Bar, error)
	DoSomethingElse(DummyService_DoSomethingElseServer) error
	DoSomethingAgain(*pkg.Bar, DummyService_DoSomethingAgainServer) error
	DoSomethingForever(DummyService_DoSomethingForeverServer) error
	mustEmbedUnimplementedDummyServiceServer()
}

// UnimplementedDummyServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDummyServiceServer struct {
}

func (UnimplementedDummyServiceServer) DoSomething(context.Context, *DummyRequest) (*pkg.Bar, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DoSomething not implemented")
}
func (UnimplementedDummyServiceServer) DoSomethingElse(DummyService_DoSomethingElseServer) error {
	return status.Errorf(codes.Unimplemented, "method DoSomethingElse not implemented")
}
func (UnimplementedDummyServiceServer) DoSomethingAgain(*pkg.Bar, DummyService_DoSomethingAgainServer) error {
	return status.Errorf(codes.Unimplemented, "method DoSomethingAgain not implemented")
}
func (UnimplementedDummyServiceServer) DoSomethingForever(DummyService_DoSomethingForeverServer) error {
	return status.Errorf(codes.Unimplemented, "method DoSomethingForever not implemented")
}
func (UnimplementedDummyServiceServer) mustEmbedUnimplementedDummyServiceServer() {}

// UnsafeDummyServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DummyServiceServer will
// result in compilation errors.
type UnsafeDummyServiceServer interface {
	mustEmbedUnimplementedDummyServiceServer()
}

func RegisterDummyServiceServer(s grpc.ServiceRegistrar, srv DummyServiceServer) {
	s.RegisterService(&DummyService_ServiceDesc, srv)
}

func _DummyService_DoSomething_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DummyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DummyServiceServer).DoSomething(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testprotos.DummyService/DoSomething",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DummyServiceServer).DoSomething(ctx, req.(*DummyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DummyService_DoSomethingElse_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyServiceServer).DoSomethingElse(&dummyServiceDoSomethingElseServer{stream})
}

type DummyService_DoSomethingElseServer interface {
	SendAndClose(*DummyResponse) error
	Recv() (*testprotos.TestMessage, error)
	grpc.ServerStream
}

type dummyServiceDoSomethingElseServer struct {
	grpc.ServerStream
}

func (x *dummyServiceDoSomethingElseServer) SendAndClose(m *DummyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyServiceDoSomethingElseServer) Recv() (*testprotos.TestMessage, error) {
	m := new(testprotos.TestMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _DummyService_DoSomethingAgain_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(pkg.Bar)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(DummyServiceServer).DoSomethingAgain(m, &dummyServiceDoSomethingAgainServer{stream})
}

type DummyService_DoSomethingAgainServer interface {
	Send(*testprotos.AnotherTestMessage) error
	grpc.ServerStream
}

type dummyServiceDoSomethingAgainServer struct {
	grpc.ServerStream
}

func (x *dummyServiceDoSomethingAgainServer) Send(m *testprotos.AnotherTestMessage) error {
	return x.ServerStream.SendMsg(m)
}

func _DummyService_DoSomethingForever_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(DummyServiceServer).DoSomethingForever(&dummyServiceDoSomethingForeverServer{stream})
}

type DummyService_DoSomethingForeverServer interface {
	Send(*DummyResponse) error
	Recv() (*DummyRequest, error)
	grpc.ServerStream
}

type dummyServiceDoSomethingForeverServer struct {
	grpc.ServerStream
}

func (x *dummyServiceDoSomethingForeverServer) Send(m *DummyResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *dummyServiceDoSomethingForeverServer) Recv() (*DummyRequest, error) {
	m := new(DummyRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// DummyService_ServiceDesc is the grpc.ServiceDesc for DummyService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DummyService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "testprotos.DummyService",
	HandlerType: (*DummyServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DoSomething",
			Handler:    _DummyService_DoSomething_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "DoSomethingElse",
			Handler:       _DummyService_DoSomethingElse_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "DoSomethingAgain",
			Handler:       _DummyService_DoSomethingAgain_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DoSomethingForever",
			Handler:       _DummyService_DoSomethingForever_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "grpc/dummy.proto",
}
