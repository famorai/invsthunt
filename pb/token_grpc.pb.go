// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// TokenServiceClient is the client API for TokenService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TokenServiceClient interface {
	AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error)
	AddTokenDetail(ctx context.Context, in *Token, opts ...grpc.CallOption) (TokenService_AddTokenDetailClient, error)
	AddTokens(ctx context.Context, opts ...grpc.CallOption) (TokenService_AddTokensClient, error)
	AddTokenStreamBi(ctx context.Context, opts ...grpc.CallOption) (TokenService_AddTokenStreamBiClient, error)
}

type tokenServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTokenServiceClient(cc grpc.ClientConnInterface) TokenServiceClient {
	return &tokenServiceClient{cc}
}

func (c *tokenServiceClient) AddToken(ctx context.Context, in *Token, opts ...grpc.CallOption) (*Token, error) {
	out := new(Token)
	err := c.cc.Invoke(ctx, "/pb.TokenService/AddToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tokenServiceClient) AddTokenDetail(ctx context.Context, in *Token, opts ...grpc.CallOption) (TokenService_AddTokenDetailClient, error) {
	stream, err := c.cc.NewStream(ctx, &TokenService_ServiceDesc.Streams[0], "/pb.TokenService/AddTokenDetail", opts...)
	if err != nil {
		return nil, err
	}
	x := &tokenServiceAddTokenDetailClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type TokenService_AddTokenDetailClient interface {
	Recv() (*TokenResultStream, error)
	grpc.ClientStream
}

type tokenServiceAddTokenDetailClient struct {
	grpc.ClientStream
}

func (x *tokenServiceAddTokenDetailClient) Recv() (*TokenResultStream, error) {
	m := new(TokenResultStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tokenServiceClient) AddTokens(ctx context.Context, opts ...grpc.CallOption) (TokenService_AddTokensClient, error) {
	stream, err := c.cc.NewStream(ctx, &TokenService_ServiceDesc.Streams[1], "/pb.TokenService/AddTokens", opts...)
	if err != nil {
		return nil, err
	}
	x := &tokenServiceAddTokensClient{stream}
	return x, nil
}

type TokenService_AddTokensClient interface {
	Send(*Token) error
	CloseAndRecv() (*Tokens, error)
	grpc.ClientStream
}

type tokenServiceAddTokensClient struct {
	grpc.ClientStream
}

func (x *tokenServiceAddTokensClient) Send(m *Token) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tokenServiceAddTokensClient) CloseAndRecv() (*Tokens, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Tokens)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *tokenServiceClient) AddTokenStreamBi(ctx context.Context, opts ...grpc.CallOption) (TokenService_AddTokenStreamBiClient, error) {
	stream, err := c.cc.NewStream(ctx, &TokenService_ServiceDesc.Streams[2], "/pb.TokenService/AddTokenStreamBi", opts...)
	if err != nil {
		return nil, err
	}
	x := &tokenServiceAddTokenStreamBiClient{stream}
	return x, nil
}

type TokenService_AddTokenStreamBiClient interface {
	Send(*Token) error
	Recv() (*TokenResultStream, error)
	grpc.ClientStream
}

type tokenServiceAddTokenStreamBiClient struct {
	grpc.ClientStream
}

func (x *tokenServiceAddTokenStreamBiClient) Send(m *Token) error {
	return x.ClientStream.SendMsg(m)
}

func (x *tokenServiceAddTokenStreamBiClient) Recv() (*TokenResultStream, error) {
	m := new(TokenResultStream)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TokenServiceServer is the server API for TokenService service.
// All implementations must embed UnimplementedTokenServiceServer
// for forward compatibility
type TokenServiceServer interface {
	AddToken(context.Context, *Token) (*Token, error)
	AddTokenDetail(*Token, TokenService_AddTokenDetailServer) error
	AddTokens(TokenService_AddTokensServer) error
	AddTokenStreamBi(TokenService_AddTokenStreamBiServer) error
	mustEmbedUnimplementedTokenServiceServer()
}

// UnimplementedTokenServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTokenServiceServer struct {
}

func (UnimplementedTokenServiceServer) AddToken(context.Context, *Token) (*Token, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddToken not implemented")
}
func (UnimplementedTokenServiceServer) AddTokenDetail(*Token, TokenService_AddTokenDetailServer) error {
	return status.Errorf(codes.Unimplemented, "method AddTokenDetail not implemented")
}
func (UnimplementedTokenServiceServer) AddTokens(TokenService_AddTokensServer) error {
	return status.Errorf(codes.Unimplemented, "method AddTokens not implemented")
}
func (UnimplementedTokenServiceServer) AddTokenStreamBi(TokenService_AddTokenStreamBiServer) error {
	return status.Errorf(codes.Unimplemented, "method AddTokenStreamBi not implemented")
}
func (UnimplementedTokenServiceServer) mustEmbedUnimplementedTokenServiceServer() {}

// UnsafeTokenServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TokenServiceServer will
// result in compilation errors.
type UnsafeTokenServiceServer interface {
	mustEmbedUnimplementedTokenServiceServer()
}

func RegisterTokenServiceServer(s grpc.ServiceRegistrar, srv TokenServiceServer) {
	s.RegisterService(&TokenService_ServiceDesc, srv)
}

func _TokenService_AddToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Token)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TokenServiceServer).AddToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.TokenService/AddToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TokenServiceServer).AddToken(ctx, req.(*Token))
	}
	return interceptor(ctx, in, info, handler)
}

func _TokenService_AddTokenDetail_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Token)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TokenServiceServer).AddTokenDetail(m, &tokenServiceAddTokenDetailServer{stream})
}

type TokenService_AddTokenDetailServer interface {
	Send(*TokenResultStream) error
	grpc.ServerStream
}

type tokenServiceAddTokenDetailServer struct {
	grpc.ServerStream
}

func (x *tokenServiceAddTokenDetailServer) Send(m *TokenResultStream) error {
	return x.ServerStream.SendMsg(m)
}

func _TokenService_AddTokens_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TokenServiceServer).AddTokens(&tokenServiceAddTokensServer{stream})
}

type TokenService_AddTokensServer interface {
	SendAndClose(*Tokens) error
	Recv() (*Token, error)
	grpc.ServerStream
}

type tokenServiceAddTokensServer struct {
	grpc.ServerStream
}

func (x *tokenServiceAddTokensServer) SendAndClose(m *Tokens) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tokenServiceAddTokensServer) Recv() (*Token, error) {
	m := new(Token)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _TokenService_AddTokenStreamBi_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TokenServiceServer).AddTokenStreamBi(&tokenServiceAddTokenStreamBiServer{stream})
}

type TokenService_AddTokenStreamBiServer interface {
	Send(*TokenResultStream) error
	Recv() (*Token, error)
	grpc.ServerStream
}

type tokenServiceAddTokenStreamBiServer struct {
	grpc.ServerStream
}

func (x *tokenServiceAddTokenStreamBiServer) Send(m *TokenResultStream) error {
	return x.ServerStream.SendMsg(m)
}

func (x *tokenServiceAddTokenStreamBiServer) Recv() (*Token, error) {
	m := new(Token)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TokenService_ServiceDesc is the grpc.ServiceDesc for TokenService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TokenService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.TokenService",
	HandlerType: (*TokenServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddToken",
			Handler:    _TokenService_AddToken_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddTokenDetail",
			Handler:       _TokenService_AddTokenDetail_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "AddTokens",
			Handler:       _TokenService_AddTokens_Handler,
			ClientStreams: true,
		},
		{
			StreamName:    "AddTokenStreamBi",
			Handler:       _TokenService_AddTokenStreamBi_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "token.proto",
}
