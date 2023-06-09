// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: shortener.proto

package __

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
	UrlShortener_SaveUrl_FullMethodName = "/api.UrlShortener/SaveUrl"
	UrlShortener_GetUrl_FullMethodName  = "/api.UrlShortener/GetUrl"
)

// UrlShortenerClient is the client API for UrlShortener service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UrlShortenerClient interface {
	SaveUrl(ctx context.Context, in *SaveUrlRequest, opts ...grpc.CallOption) (*SaveUrlResponse, error)
	GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error)
}

type urlShortenerClient struct {
	cc grpc.ClientConnInterface
}

func NewUrlShortenerClient(cc grpc.ClientConnInterface) UrlShortenerClient {
	return &urlShortenerClient{cc}
}

func (c *urlShortenerClient) SaveUrl(ctx context.Context, in *SaveUrlRequest, opts ...grpc.CallOption) (*SaveUrlResponse, error) {
	out := new(SaveUrlResponse)
	err := c.cc.Invoke(ctx, UrlShortener_SaveUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerClient) GetUrl(ctx context.Context, in *GetUrlRequest, opts ...grpc.CallOption) (*GetUrlResponse, error) {
	out := new(GetUrlResponse)
	err := c.cc.Invoke(ctx, UrlShortener_GetUrl_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UrlShortenerServer is the server API for UrlShortener service.
// All implementations must embed UnimplementedUrlShortenerServer
// for forward compatibility
type UrlShortenerServer interface {
	SaveUrl(context.Context, *SaveUrlRequest) (*SaveUrlResponse, error)
	GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error)
	mustEmbedUnimplementedUrlShortenerServer()
}

// UnimplementedUrlShortenerServer must be embedded to have forward compatible implementations.
type UnimplementedUrlShortenerServer struct {
}

func (UnimplementedUrlShortenerServer) SaveUrl(context.Context, *SaveUrlRequest) (*SaveUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveUrl not implemented")
}
func (UnimplementedUrlShortenerServer) GetUrl(context.Context, *GetUrlRequest) (*GetUrlResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUrl not implemented")
}
func (UnimplementedUrlShortenerServer) mustEmbedUnimplementedUrlShortenerServer() {}

// UnsafeUrlShortenerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UrlShortenerServer will
// result in compilation errors.
type UnsafeUrlShortenerServer interface {
	mustEmbedUnimplementedUrlShortenerServer()
}

func RegisterUrlShortenerServer(s grpc.ServiceRegistrar, srv UrlShortenerServer) {
	s.RegisterService(&UrlShortener_ServiceDesc, srv)
}

func _UrlShortener_SaveUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SaveUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).SaveUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortener_SaveUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).SaveUrl(ctx, req.(*SaveUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UrlShortener_GetUrl_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUrlRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UrlShortenerServer).GetUrl(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UrlShortener_GetUrl_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UrlShortenerServer).GetUrl(ctx, req.(*GetUrlRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UrlShortener_ServiceDesc is the grpc.ServiceDesc for UrlShortener service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UrlShortener_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.UrlShortener",
	HandlerType: (*UrlShortenerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveUrl",
			Handler:    _UrlShortener_SaveUrl_Handler,
		},
		{
			MethodName: "GetUrl",
			Handler:    _UrlShortener_GetUrl_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "shortener.proto",
}
