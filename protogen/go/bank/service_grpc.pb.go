// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.12.4
// source: proto/bank/service.proto

package bank

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

const (
	BankService_GetCurrentBalance_FullMethodName    = "/bank.BankService/GetCurrentBalance"
	BankService_FetchExchangeRate_FullMethodName    = "/bank.BankService/FetchExchangeRate"
	BankService_SummarizeTransaction_FullMethodName = "/bank.BankService/SummarizeTransaction"
	BankService_TransferMultiple_FullMethodName     = "/bank.BankService/TransferMultiple"
)

// BankServiceClient is the client API for BankService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BankServiceClient interface {
	GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error)
	FetchExchangeRate(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (BankService_FetchExchangeRateClient, error)
	SummarizeTransaction(ctx context.Context, opts ...grpc.CallOption) (BankService_SummarizeTransactionClient, error)
	TransferMultiple(ctx context.Context, opts ...grpc.CallOption) (BankService_TransferMultipleClient, error)
}

type bankServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBankServiceClient(cc grpc.ClientConnInterface) BankServiceClient {
	return &bankServiceClient{cc}
}

func (c *bankServiceClient) GetCurrentBalance(ctx context.Context, in *CurrentBalanceRequest, opts ...grpc.CallOption) (*CurrentBalanceResponse, error) {
	out := new(CurrentBalanceResponse)
	err := c.cc.Invoke(ctx, BankService_GetCurrentBalance_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bankServiceClient) FetchExchangeRate(ctx context.Context, in *ExchangeRateRequest, opts ...grpc.CallOption) (BankService_FetchExchangeRateClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[0], BankService_FetchExchangeRate_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceFetchExchangeRateClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type BankService_FetchExchangeRateClient interface {
	Recv() (*ExchangeRateResponse, error)
	grpc.ClientStream
}

type bankServiceFetchExchangeRateClient struct {
	grpc.ClientStream
}

func (x *bankServiceFetchExchangeRateClient) Recv() (*ExchangeRateResponse, error) {
	m := new(ExchangeRateResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankServiceClient) SummarizeTransaction(ctx context.Context, opts ...grpc.CallOption) (BankService_SummarizeTransactionClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[1], BankService_SummarizeTransaction_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceSummarizeTransactionClient{stream}
	return x, nil
}

type BankService_SummarizeTransactionClient interface {
	Send(*Transaction) error
	CloseAndRecv() (*TransactionSummary, error)
	grpc.ClientStream
}

type bankServiceSummarizeTransactionClient struct {
	grpc.ClientStream
}

func (x *bankServiceSummarizeTransactionClient) Send(m *Transaction) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bankServiceSummarizeTransactionClient) CloseAndRecv() (*TransactionSummary, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(TransactionSummary)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *bankServiceClient) TransferMultiple(ctx context.Context, opts ...grpc.CallOption) (BankService_TransferMultipleClient, error) {
	stream, err := c.cc.NewStream(ctx, &BankService_ServiceDesc.Streams[2], BankService_TransferMultiple_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &bankServiceTransferMultipleClient{stream}
	return x, nil
}

type BankService_TransferMultipleClient interface {
	Send(*TransferRequest) error
	Recv() (*TransferResponse, error)
	grpc.ClientStream
}

type bankServiceTransferMultipleClient struct {
	grpc.ClientStream
}

func (x *bankServiceTransferMultipleClient) Send(m *TransferRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *bankServiceTransferMultipleClient) Recv() (*TransferResponse, error) {
	m := new(TransferResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BankServiceServer is the server API for BankService service.
// All implementations must embed UnimplementedBankServiceServer
// for forward compatibility
type BankServiceServer interface {
	GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error)
	FetchExchangeRate(*ExchangeRateRequest, BankService_FetchExchangeRateServer) error
	SummarizeTransaction(BankService_SummarizeTransactionServer) error
	TransferMultiple(BankService_TransferMultipleServer) error
	mustEmbedUnimplementedBankServiceServer()
}

// UnimplementedBankServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBankServiceServer struct {
}

func (UnimplementedBankServiceServer) GetCurrentBalance(context.Context, *CurrentBalanceRequest) (*CurrentBalanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCurrentBalance not implemented")
}
func (UnimplementedBankServiceServer) FetchExchangeRate(*ExchangeRateRequest, BankService_FetchExchangeRateServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchExchangeRate not implemented")
}
func (UnimplementedBankServiceServer) SummarizeTransaction(BankService_SummarizeTransactionServer) error {
	return status.Errorf(codes.Unimplemented, "method SummarizeTransaction not implemented")
}
func (UnimplementedBankServiceServer) TransferMultiple(BankService_TransferMultipleServer) error {
	return status.Errorf(codes.Unimplemented, "method TransferMultiple not implemented")
}
func (UnimplementedBankServiceServer) mustEmbedUnimplementedBankServiceServer() {}

// UnsafeBankServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BankServiceServer will
// result in compilation errors.
type UnsafeBankServiceServer interface {
	mustEmbedUnimplementedBankServiceServer()
}

func RegisterBankServiceServer(s grpc.ServiceRegistrar, srv BankServiceServer) {
	s.RegisterService(&BankService_ServiceDesc, srv)
}

func _BankService_GetCurrentBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CurrentBalanceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BankServiceServer).GetCurrentBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BankService_GetCurrentBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BankServiceServer).GetCurrentBalance(ctx, req.(*CurrentBalanceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BankService_FetchExchangeRate_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ExchangeRateRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BankServiceServer).FetchExchangeRate(m, &bankServiceFetchExchangeRateServer{stream})
}

type BankService_FetchExchangeRateServer interface {
	Send(*ExchangeRateResponse) error
	grpc.ServerStream
}

type bankServiceFetchExchangeRateServer struct {
	grpc.ServerStream
}

func (x *bankServiceFetchExchangeRateServer) Send(m *ExchangeRateResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _BankService_SummarizeTransaction_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BankServiceServer).SummarizeTransaction(&bankServiceSummarizeTransactionServer{stream})
}

type BankService_SummarizeTransactionServer interface {
	SendAndClose(*TransactionSummary) error
	Recv() (*Transaction, error)
	grpc.ServerStream
}

type bankServiceSummarizeTransactionServer struct {
	grpc.ServerStream
}

func (x *bankServiceSummarizeTransactionServer) SendAndClose(m *TransactionSummary) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bankServiceSummarizeTransactionServer) Recv() (*Transaction, error) {
	m := new(Transaction)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _BankService_TransferMultiple_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BankServiceServer).TransferMultiple(&bankServiceTransferMultipleServer{stream})
}

type BankService_TransferMultipleServer interface {
	Send(*TransferResponse) error
	Recv() (*TransferRequest, error)
	grpc.ServerStream
}

type bankServiceTransferMultipleServer struct {
	grpc.ServerStream
}

func (x *bankServiceTransferMultipleServer) Send(m *TransferResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *bankServiceTransferMultipleServer) Recv() (*TransferRequest, error) {
	m := new(TransferRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BankService_ServiceDesc is the grpc.ServiceDesc for BankService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BankService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "bank.BankService",
	HandlerType: (*BankServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCurrentBalance",
			Handler:    _BankService_GetCurrentBalance_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchExchangeRate",
			Handler:       _BankService_FetchExchangeRate_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "SummarizeTransaction",
			Handler:       _BankService_SummarizeTransaction_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "TransferMultiple",
			Handler:       _BankService_TransferMultiple_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "proto/bank/service.proto",
}
