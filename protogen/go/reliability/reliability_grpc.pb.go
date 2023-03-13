// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/reliability/reliability.proto

package reliability

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

// ReliabilityServiceClient is the client API for ReliabilityService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ReliabilityServiceClient interface {
	UnaryReliability(ctx context.Context, in *ReliabilityRequest, opts ...grpc.CallOption) (*ReliabilityResponse, error)
	ServerStreamingReliability(ctx context.Context, in *ReliabilityRequest, opts ...grpc.CallOption) (ReliabilityService_ServerStreamingReliabilityClient, error)
	ClientStreamingReliability(ctx context.Context, opts ...grpc.CallOption) (ReliabilityService_ClientStreamingReliabilityClient, error)
	BiDirectionalReliability(ctx context.Context, opts ...grpc.CallOption) (ReliabilityService_BiDirectionalReliabilityClient, error)
}

type reliabilityServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewReliabilityServiceClient(cc grpc.ClientConnInterface) ReliabilityServiceClient {
	return &reliabilityServiceClient{cc}
}

func (c *reliabilityServiceClient) UnaryReliability(ctx context.Context, in *ReliabilityRequest, opts ...grpc.CallOption) (*ReliabilityResponse, error) {
	out := new(ReliabilityResponse)
	err := c.cc.Invoke(ctx, "/reliability.ReliabilityService/UnaryReliability", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *reliabilityServiceClient) ServerStreamingReliability(ctx context.Context, in *ReliabilityRequest, opts ...grpc.CallOption) (ReliabilityService_ServerStreamingReliabilityClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReliabilityService_ServiceDesc.Streams[0], "/reliability.ReliabilityService/ServerStreamingReliability", opts...)
	if err != nil {
		return nil, err
	}
	x := &reliabilityServiceServerStreamingReliabilityClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ReliabilityService_ServerStreamingReliabilityClient interface {
	Recv() (*ReliabilityResponse, error)
	grpc.ClientStream
}

type reliabilityServiceServerStreamingReliabilityClient struct {
	grpc.ClientStream
}

func (x *reliabilityServiceServerStreamingReliabilityClient) Recv() (*ReliabilityResponse, error) {
	m := new(ReliabilityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reliabilityServiceClient) ClientStreamingReliability(ctx context.Context, opts ...grpc.CallOption) (ReliabilityService_ClientStreamingReliabilityClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReliabilityService_ServiceDesc.Streams[1], "/reliability.ReliabilityService/ClientStreamingReliability", opts...)
	if err != nil {
		return nil, err
	}
	x := &reliabilityServiceClientStreamingReliabilityClient{stream}
	return x, nil
}

type ReliabilityService_ClientStreamingReliabilityClient interface {
	Send(*ReliabilityRequest) error
	CloseAndRecv() (*ReliabilityResponse, error)
	grpc.ClientStream
}

type reliabilityServiceClientStreamingReliabilityClient struct {
	grpc.ClientStream
}

func (x *reliabilityServiceClientStreamingReliabilityClient) Send(m *ReliabilityRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reliabilityServiceClientStreamingReliabilityClient) CloseAndRecv() (*ReliabilityResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(ReliabilityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *reliabilityServiceClient) BiDirectionalReliability(ctx context.Context, opts ...grpc.CallOption) (ReliabilityService_BiDirectionalReliabilityClient, error) {
	stream, err := c.cc.NewStream(ctx, &ReliabilityService_ServiceDesc.Streams[2], "/reliability.ReliabilityService/BiDirectionalReliability", opts...)
	if err != nil {
		return nil, err
	}
	x := &reliabilityServiceBiDirectionalReliabilityClient{stream}
	return x, nil
}

type ReliabilityService_BiDirectionalReliabilityClient interface {
	Send(*ReliabilityRequest) error
	Recv() (*ReliabilityResponse, error)
	grpc.ClientStream
}

type reliabilityServiceBiDirectionalReliabilityClient struct {
	grpc.ClientStream
}

func (x *reliabilityServiceBiDirectionalReliabilityClient) Send(m *ReliabilityRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *reliabilityServiceBiDirectionalReliabilityClient) Recv() (*ReliabilityResponse, error) {
	m := new(ReliabilityResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReliabilityServiceServer is the server API for ReliabilityService service.
// All implementations must embed UnimplementedReliabilityServiceServer
// for forward compatibility
type ReliabilityServiceServer interface {
	UnaryReliability(context.Context, *ReliabilityRequest) (*ReliabilityResponse, error)
	ServerStreamingReliability(*ReliabilityRequest, ReliabilityService_ServerStreamingReliabilityServer) error
	ClientStreamingReliability(ReliabilityService_ClientStreamingReliabilityServer) error
	BiDirectionalReliability(ReliabilityService_BiDirectionalReliabilityServer) error
	mustEmbedUnimplementedReliabilityServiceServer()
}

// UnimplementedReliabilityServiceServer must be embedded to have forward compatible implementations.
type UnimplementedReliabilityServiceServer struct {
}

func (UnimplementedReliabilityServiceServer) UnaryReliability(context.Context, *ReliabilityRequest) (*ReliabilityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnaryReliability not implemented")
}
func (UnimplementedReliabilityServiceServer) ServerStreamingReliability(*ReliabilityRequest, ReliabilityService_ServerStreamingReliabilityServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerStreamingReliability not implemented")
}
func (UnimplementedReliabilityServiceServer) ClientStreamingReliability(ReliabilityService_ClientStreamingReliabilityServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientStreamingReliability not implemented")
}
func (UnimplementedReliabilityServiceServer) BiDirectionalReliability(ReliabilityService_BiDirectionalReliabilityServer) error {
	return status.Errorf(codes.Unimplemented, "method BiDirectionalReliability not implemented")
}
func (UnimplementedReliabilityServiceServer) mustEmbedUnimplementedReliabilityServiceServer() {}

// UnsafeReliabilityServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ReliabilityServiceServer will
// result in compilation errors.
type UnsafeReliabilityServiceServer interface {
	mustEmbedUnimplementedReliabilityServiceServer()
}

func RegisterReliabilityServiceServer(s grpc.ServiceRegistrar, srv ReliabilityServiceServer) {
	s.RegisterService(&ReliabilityService_ServiceDesc, srv)
}

func _ReliabilityService_UnaryReliability_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReliabilityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ReliabilityServiceServer).UnaryReliability(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/reliability.ReliabilityService/UnaryReliability",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ReliabilityServiceServer).UnaryReliability(ctx, req.(*ReliabilityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ReliabilityService_ServerStreamingReliability_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ReliabilityRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ReliabilityServiceServer).ServerStreamingReliability(m, &reliabilityServiceServerStreamingReliabilityServer{stream})
}

type ReliabilityService_ServerStreamingReliabilityServer interface {
	Send(*ReliabilityResponse) error
	grpc.ServerStream
}

type reliabilityServiceServerStreamingReliabilityServer struct {
	grpc.ServerStream
}

func (x *reliabilityServiceServerStreamingReliabilityServer) Send(m *ReliabilityResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _ReliabilityService_ClientStreamingReliability_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReliabilityServiceServer).ClientStreamingReliability(&reliabilityServiceClientStreamingReliabilityServer{stream})
}

type ReliabilityService_ClientStreamingReliabilityServer interface {
	SendAndClose(*ReliabilityResponse) error
	Recv() (*ReliabilityRequest, error)
	grpc.ServerStream
}

type reliabilityServiceClientStreamingReliabilityServer struct {
	grpc.ServerStream
}

func (x *reliabilityServiceClientStreamingReliabilityServer) SendAndClose(m *ReliabilityResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reliabilityServiceClientStreamingReliabilityServer) Recv() (*ReliabilityRequest, error) {
	m := new(ReliabilityRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ReliabilityService_BiDirectionalReliability_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ReliabilityServiceServer).BiDirectionalReliability(&reliabilityServiceBiDirectionalReliabilityServer{stream})
}

type ReliabilityService_BiDirectionalReliabilityServer interface {
	Send(*ReliabilityResponse) error
	Recv() (*ReliabilityRequest, error)
	grpc.ServerStream
}

type reliabilityServiceBiDirectionalReliabilityServer struct {
	grpc.ServerStream
}

func (x *reliabilityServiceBiDirectionalReliabilityServer) Send(m *ReliabilityResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *reliabilityServiceBiDirectionalReliabilityServer) Recv() (*ReliabilityRequest, error) {
	m := new(ReliabilityRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ReliabilityService_ServiceDesc is the grpc.ServiceDesc for ReliabilityService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ReliabilityService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reliability.ReliabilityService",
	HandlerType: (*ReliabilityServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UnaryReliability",
			Handler:    _ReliabilityService_UnaryReliability_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerStreamingReliability",
			Handler:       _ReliabilityService_ServerStreamingReliability_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "ClientStreamingReliability",
			Handler:       _ReliabilityService_ClientStreamingReliability_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "BiDirectionalReliability",
			Handler:       _ReliabilityService_BiDirectionalReliability_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/reliability/reliability.proto",
}
