// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.20.1
// source: message.proto

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

// TextServiceClient is the client API for TextService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TextServiceClient interface {
	TextFunc(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextResponse, error)
}

type textServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTextServiceClient(cc grpc.ClientConnInterface) TextServiceClient {
	return &textServiceClient{cc}
}

func (c *textServiceClient) TextFunc(ctx context.Context, in *TextRequest, opts ...grpc.CallOption) (*TextResponse, error) {
	out := new(TextResponse)
	err := c.cc.Invoke(ctx, "/TextService/TextFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TextServiceServer is the server API for TextService service.
// All implementations must embed UnimplementedTextServiceServer
// for forward compatibility
type TextServiceServer interface {
	TextFunc(context.Context, *TextRequest) (*TextResponse, error)
	mustEmbedUnimplementedTextServiceServer()
}

// UnimplementedTextServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTextServiceServer struct {
}

func (UnimplementedTextServiceServer) TextFunc(context.Context, *TextRequest) (*TextResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TextFunc not implemented")
}
func (UnimplementedTextServiceServer) mustEmbedUnimplementedTextServiceServer() {}

// UnsafeTextServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TextServiceServer will
// result in compilation errors.
type UnsafeTextServiceServer interface {
	mustEmbedUnimplementedTextServiceServer()
}

func RegisterTextServiceServer(s grpc.ServiceRegistrar, srv TextServiceServer) {
	s.RegisterService(&TextService_ServiceDesc, srv)
}

func _TextService_TextFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TextServiceServer).TextFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/TextService/TextFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TextServiceServer).TextFunc(ctx, req.(*TextRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TextService_ServiceDesc is the grpc.ServiceDesc for TextService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TextService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "TextService",
	HandlerType: (*TextServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TextFunc",
			Handler:    _TextService_TextFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// BytesServiceClient is the client API for BytesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BytesServiceClient interface {
	BytesFunc(ctx context.Context, in *Bytesmessage, opts ...grpc.CallOption) (*Bytesmessage, error)
}

type bytesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBytesServiceClient(cc grpc.ClientConnInterface) BytesServiceClient {
	return &bytesServiceClient{cc}
}

func (c *bytesServiceClient) BytesFunc(ctx context.Context, in *Bytesmessage, opts ...grpc.CallOption) (*Bytesmessage, error) {
	out := new(Bytesmessage)
	err := c.cc.Invoke(ctx, "/BytesService/BytesFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BytesServiceServer is the server API for BytesService service.
// All implementations must embed UnimplementedBytesServiceServer
// for forward compatibility
type BytesServiceServer interface {
	BytesFunc(context.Context, *Bytesmessage) (*Bytesmessage, error)
	mustEmbedUnimplementedBytesServiceServer()
}

// UnimplementedBytesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBytesServiceServer struct {
}

func (UnimplementedBytesServiceServer) BytesFunc(context.Context, *Bytesmessage) (*Bytesmessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BytesFunc not implemented")
}
func (UnimplementedBytesServiceServer) mustEmbedUnimplementedBytesServiceServer() {}

// UnsafeBytesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BytesServiceServer will
// result in compilation errors.
type UnsafeBytesServiceServer interface {
	mustEmbedUnimplementedBytesServiceServer()
}

func RegisterBytesServiceServer(s grpc.ServiceRegistrar, srv BytesServiceServer) {
	s.RegisterService(&BytesService_ServiceDesc, srv)
}

func _BytesService_BytesFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Bytesmessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BytesServiceServer).BytesFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BytesService/BytesFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BytesServiceServer).BytesFunc(ctx, req.(*Bytesmessage))
	}
	return interceptor(ctx, in, info, handler)
}

// BytesService_ServiceDesc is the grpc.ServiceDesc for BytesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BytesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BytesService",
	HandlerType: (*BytesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BytesFunc",
			Handler:    _BytesService_BytesFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// UploadServiceClient is the client API for UploadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UploadServiceClient interface {
	UploadFunc(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error)
}

type uploadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUploadServiceClient(cc grpc.ClientConnInterface) UploadServiceClient {
	return &uploadServiceClient{cc}
}

func (c *uploadServiceClient) UploadFunc(ctx context.Context, in *UploadRequest, opts ...grpc.CallOption) (*UploadResponse, error) {
	out := new(UploadResponse)
	err := c.cc.Invoke(ctx, "/UploadService/UploadFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UploadServiceServer is the server API for UploadService service.
// All implementations must embed UnimplementedUploadServiceServer
// for forward compatibility
type UploadServiceServer interface {
	UploadFunc(context.Context, *UploadRequest) (*UploadResponse, error)
	mustEmbedUnimplementedUploadServiceServer()
}

// UnimplementedUploadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUploadServiceServer struct {
}

func (UnimplementedUploadServiceServer) UploadFunc(context.Context, *UploadRequest) (*UploadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UploadFunc not implemented")
}
func (UnimplementedUploadServiceServer) mustEmbedUnimplementedUploadServiceServer() {}

// UnsafeUploadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UploadServiceServer will
// result in compilation errors.
type UnsafeUploadServiceServer interface {
	mustEmbedUnimplementedUploadServiceServer()
}

func RegisterUploadServiceServer(s grpc.ServiceRegistrar, srv UploadServiceServer) {
	s.RegisterService(&UploadService_ServiceDesc, srv)
}

func _UploadService_UploadFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UploadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UploadServiceServer).UploadFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/UploadService/UploadFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UploadServiceServer).UploadFunc(ctx, req.(*UploadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UploadService_ServiceDesc is the grpc.ServiceDesc for UploadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UploadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "UploadService",
	HandlerType: (*UploadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UploadFunc",
			Handler:    _UploadService_UploadFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// DownloadServiceClient is the client API for DownloadService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DownloadServiceClient interface {
	DownloadFunc(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error)
}

type downloadServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDownloadServiceClient(cc grpc.ClientConnInterface) DownloadServiceClient {
	return &downloadServiceClient{cc}
}

func (c *downloadServiceClient) DownloadFunc(ctx context.Context, in *DownloadRequest, opts ...grpc.CallOption) (*DownloadResponse, error) {
	out := new(DownloadResponse)
	err := c.cc.Invoke(ctx, "/DownloadService/DownloadFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DownloadServiceServer is the server API for DownloadService service.
// All implementations must embed UnimplementedDownloadServiceServer
// for forward compatibility
type DownloadServiceServer interface {
	DownloadFunc(context.Context, *DownloadRequest) (*DownloadResponse, error)
	mustEmbedUnimplementedDownloadServiceServer()
}

// UnimplementedDownloadServiceServer must be embedded to have forward compatible implementations.
type UnimplementedDownloadServiceServer struct {
}

func (UnimplementedDownloadServiceServer) DownloadFunc(context.Context, *DownloadRequest) (*DownloadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DownloadFunc not implemented")
}
func (UnimplementedDownloadServiceServer) mustEmbedUnimplementedDownloadServiceServer() {}

// UnsafeDownloadServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DownloadServiceServer will
// result in compilation errors.
type UnsafeDownloadServiceServer interface {
	mustEmbedUnimplementedDownloadServiceServer()
}

func RegisterDownloadServiceServer(s grpc.ServiceRegistrar, srv DownloadServiceServer) {
	s.RegisterService(&DownloadService_ServiceDesc, srv)
}

func _DownloadService_DownloadFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownloadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownloadServiceServer).DownloadFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DownloadService/DownloadFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownloadServiceServer).DownloadFunc(ctx, req.(*DownloadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DownloadService_ServiceDesc is the grpc.ServiceDesc for DownloadService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DownloadService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DownloadService",
	HandlerType: (*DownloadServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DownloadFunc",
			Handler:    _DownloadService_DownloadFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// BigDataServiceClient is the client API for BigDataService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BigDataServiceClient interface {
	BigDataFunc(ctx context.Context, in *BigDataRequest, opts ...grpc.CallOption) (*BigDataResponse, error)
}

type bigDataServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBigDataServiceClient(cc grpc.ClientConnInterface) BigDataServiceClient {
	return &bigDataServiceClient{cc}
}

func (c *bigDataServiceClient) BigDataFunc(ctx context.Context, in *BigDataRequest, opts ...grpc.CallOption) (*BigDataResponse, error) {
	out := new(BigDataResponse)
	err := c.cc.Invoke(ctx, "/BigDataService/BigDataFunc", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BigDataServiceServer is the server API for BigDataService service.
// All implementations must embed UnimplementedBigDataServiceServer
// for forward compatibility
type BigDataServiceServer interface {
	BigDataFunc(context.Context, *BigDataRequest) (*BigDataResponse, error)
	mustEmbedUnimplementedBigDataServiceServer()
}

// UnimplementedBigDataServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBigDataServiceServer struct {
}

func (UnimplementedBigDataServiceServer) BigDataFunc(context.Context, *BigDataRequest) (*BigDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BigDataFunc not implemented")
}
func (UnimplementedBigDataServiceServer) mustEmbedUnimplementedBigDataServiceServer() {}

// UnsafeBigDataServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BigDataServiceServer will
// result in compilation errors.
type UnsafeBigDataServiceServer interface {
	mustEmbedUnimplementedBigDataServiceServer()
}

func RegisterBigDataServiceServer(s grpc.ServiceRegistrar, srv BigDataServiceServer) {
	s.RegisterService(&BigDataService_ServiceDesc, srv)
}

func _BigDataService_BigDataFunc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BigDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BigDataServiceServer).BigDataFunc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/BigDataService/BigDataFunc",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BigDataServiceServer).BigDataFunc(ctx, req.(*BigDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BigDataService_ServiceDesc is the grpc.ServiceDesc for BigDataService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BigDataService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BigDataService",
	HandlerType: (*BigDataServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BigDataFunc",
			Handler:    _BigDataService_BigDataFunc_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

// ServerSideStreamingServiceClient is the client API for ServerSideStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServerSideStreamingServiceClient interface {
	ServerSideStreamingFunc(ctx context.Context, in *Bytesmessage, opts ...grpc.CallOption) (ServerSideStreamingService_ServerSideStreamingFuncClient, error)
}

type serverSideStreamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewServerSideStreamingServiceClient(cc grpc.ClientConnInterface) ServerSideStreamingServiceClient {
	return &serverSideStreamingServiceClient{cc}
}

func (c *serverSideStreamingServiceClient) ServerSideStreamingFunc(ctx context.Context, in *Bytesmessage, opts ...grpc.CallOption) (ServerSideStreamingService_ServerSideStreamingFuncClient, error) {
	stream, err := c.cc.NewStream(ctx, &ServerSideStreamingService_ServiceDesc.Streams[0], "/ServerSideStreamingService/ServerSideStreamingFunc", opts...)
	if err != nil {
		return nil, err
	}
	x := &serverSideStreamingServiceServerSideStreamingFuncClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ServerSideStreamingService_ServerSideStreamingFuncClient interface {
	Recv() (*Bytesmessage, error)
	grpc.ClientStream
}

type serverSideStreamingServiceServerSideStreamingFuncClient struct {
	grpc.ClientStream
}

func (x *serverSideStreamingServiceServerSideStreamingFuncClient) Recv() (*Bytesmessage, error) {
	m := new(Bytesmessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServerSideStreamingServiceServer is the server API for ServerSideStreamingService service.
// All implementations must embed UnimplementedServerSideStreamingServiceServer
// for forward compatibility
type ServerSideStreamingServiceServer interface {
	ServerSideStreamingFunc(*Bytesmessage, ServerSideStreamingService_ServerSideStreamingFuncServer) error
	mustEmbedUnimplementedServerSideStreamingServiceServer()
}

// UnimplementedServerSideStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServerSideStreamingServiceServer struct {
}

func (UnimplementedServerSideStreamingServiceServer) ServerSideStreamingFunc(*Bytesmessage, ServerSideStreamingService_ServerSideStreamingFuncServer) error {
	return status.Errorf(codes.Unimplemented, "method ServerSideStreamingFunc not implemented")
}
func (UnimplementedServerSideStreamingServiceServer) mustEmbedUnimplementedServerSideStreamingServiceServer() {
}

// UnsafeServerSideStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServerSideStreamingServiceServer will
// result in compilation errors.
type UnsafeServerSideStreamingServiceServer interface {
	mustEmbedUnimplementedServerSideStreamingServiceServer()
}

func RegisterServerSideStreamingServiceServer(s grpc.ServiceRegistrar, srv ServerSideStreamingServiceServer) {
	s.RegisterService(&ServerSideStreamingService_ServiceDesc, srv)
}

func _ServerSideStreamingService_ServerSideStreamingFunc_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Bytesmessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServerSideStreamingServiceServer).ServerSideStreamingFunc(m, &serverSideStreamingServiceServerSideStreamingFuncServer{stream})
}

type ServerSideStreamingService_ServerSideStreamingFuncServer interface {
	Send(*Bytesmessage) error
	grpc.ServerStream
}

type serverSideStreamingServiceServerSideStreamingFuncServer struct {
	grpc.ServerStream
}

func (x *serverSideStreamingServiceServerSideStreamingFuncServer) Send(m *Bytesmessage) error {
	return x.ServerStream.SendMsg(m)
}

// ServerSideStreamingService_ServiceDesc is the grpc.ServiceDesc for ServerSideStreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ServerSideStreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ServerSideStreamingService",
	HandlerType: (*ServerSideStreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ServerSideStreamingFunc",
			Handler:       _ServerSideStreamingService_ServerSideStreamingFunc_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "message.proto",
}

// ClientSideStreamingServiceClient is the client API for ClientSideStreamingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClientSideStreamingServiceClient interface {
	ClientSideStreamingFunc(ctx context.Context, opts ...grpc.CallOption) (ClientSideStreamingService_ClientSideStreamingFuncClient, error)
}

type clientSideStreamingServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClientSideStreamingServiceClient(cc grpc.ClientConnInterface) ClientSideStreamingServiceClient {
	return &clientSideStreamingServiceClient{cc}
}

func (c *clientSideStreamingServiceClient) ClientSideStreamingFunc(ctx context.Context, opts ...grpc.CallOption) (ClientSideStreamingService_ClientSideStreamingFuncClient, error) {
	stream, err := c.cc.NewStream(ctx, &ClientSideStreamingService_ServiceDesc.Streams[0], "/ClientSideStreamingService/ClientSideStreamingFunc", opts...)
	if err != nil {
		return nil, err
	}
	x := &clientSideStreamingServiceClientSideStreamingFuncClient{stream}
	return x, nil
}

type ClientSideStreamingService_ClientSideStreamingFuncClient interface {
	Send(*BigDataResponse) error
	CloseAndRecv() (*BigDataResponse, error)
	grpc.ClientStream
}

type clientSideStreamingServiceClientSideStreamingFuncClient struct {
	grpc.ClientStream
}

func (x *clientSideStreamingServiceClientSideStreamingFuncClient) Send(m *BigDataResponse) error {
	return x.ClientStream.SendMsg(m)
}

func (x *clientSideStreamingServiceClientSideStreamingFuncClient) CloseAndRecv() (*BigDataResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(BigDataResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientSideStreamingServiceServer is the server API for ClientSideStreamingService service.
// All implementations must embed UnimplementedClientSideStreamingServiceServer
// for forward compatibility
type ClientSideStreamingServiceServer interface {
	ClientSideStreamingFunc(ClientSideStreamingService_ClientSideStreamingFuncServer) error
	mustEmbedUnimplementedClientSideStreamingServiceServer()
}

// UnimplementedClientSideStreamingServiceServer must be embedded to have forward compatible implementations.
type UnimplementedClientSideStreamingServiceServer struct {
}

func (UnimplementedClientSideStreamingServiceServer) ClientSideStreamingFunc(ClientSideStreamingService_ClientSideStreamingFuncServer) error {
	return status.Errorf(codes.Unimplemented, "method ClientSideStreamingFunc not implemented")
}
func (UnimplementedClientSideStreamingServiceServer) mustEmbedUnimplementedClientSideStreamingServiceServer() {
}

// UnsafeClientSideStreamingServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClientSideStreamingServiceServer will
// result in compilation errors.
type UnsafeClientSideStreamingServiceServer interface {
	mustEmbedUnimplementedClientSideStreamingServiceServer()
}

func RegisterClientSideStreamingServiceServer(s grpc.ServiceRegistrar, srv ClientSideStreamingServiceServer) {
	s.RegisterService(&ClientSideStreamingService_ServiceDesc, srv)
}

func _ClientSideStreamingService_ClientSideStreamingFunc_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ClientSideStreamingServiceServer).ClientSideStreamingFunc(&clientSideStreamingServiceClientSideStreamingFuncServer{stream})
}

type ClientSideStreamingService_ClientSideStreamingFuncServer interface {
	SendAndClose(*BigDataResponse) error
	Recv() (*BigDataResponse, error)
	grpc.ServerStream
}

type clientSideStreamingServiceClientSideStreamingFuncServer struct {
	grpc.ServerStream
}

func (x *clientSideStreamingServiceClientSideStreamingFuncServer) SendAndClose(m *BigDataResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *clientSideStreamingServiceClientSideStreamingFuncServer) Recv() (*BigDataResponse, error) {
	m := new(BigDataResponse)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ClientSideStreamingService_ServiceDesc is the grpc.ServiceDesc for ClientSideStreamingService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClientSideStreamingService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ClientSideStreamingService",
	HandlerType: (*ClientSideStreamingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ClientSideStreamingFunc",
			Handler:       _ClientSideStreamingService_ClientSideStreamingFunc_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "message.proto",
}
