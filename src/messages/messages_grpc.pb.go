// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package messages

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

// ImageServiceClient is the client API for ImageService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImageServiceClient interface {
	ApplyCustomFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	ApplyGaussianBlurFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	ApplyBoxBlurFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	ApplyEdgeDetectionFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	ApplySharpenFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
	ApplyNegative(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error)
}

type imageServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImageServiceClient(cc grpc.ClientConnInterface) ImageServiceClient {
	return &imageServiceClient{cc}
}

func (c *imageServiceClient) ApplyCustomFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/ImageService/ApplyCustomFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ApplyGaussianBlurFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/ImageService/ApplyGaussianBlurFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ApplyBoxBlurFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/ImageService/ApplyBoxBlurFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ApplyEdgeDetectionFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/ImageService/ApplyEdgeDetectionFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ApplySharpenFilter(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/ImageService/ApplySharpenFilter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *imageServiceClient) ApplyNegative(ctx context.Context, in *ImageRequest, opts ...grpc.CallOption) (*ImageResponse, error) {
	out := new(ImageResponse)
	err := c.cc.Invoke(ctx, "/ImageService/ApplyNegative", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImageServiceServer is the server API for ImageService service.
// All implementations must embed UnimplementedImageServiceServer
// for forward compatibility
type ImageServiceServer interface {
	ApplyCustomFilter(context.Context, *ImageRequest) (*ImageResponse, error)
	ApplyGaussianBlurFilter(context.Context, *ImageRequest) (*ImageResponse, error)
	ApplyBoxBlurFilter(context.Context, *ImageRequest) (*ImageResponse, error)
	ApplyEdgeDetectionFilter(context.Context, *ImageRequest) (*ImageResponse, error)
	ApplySharpenFilter(context.Context, *ImageRequest) (*ImageResponse, error)
	ApplyNegative(context.Context, *ImageRequest) (*ImageResponse, error)
	mustEmbedUnimplementedImageServiceServer()
}

// UnimplementedImageServiceServer must be embedded to have forward compatible implementations.
type UnimplementedImageServiceServer struct {
}

func (UnimplementedImageServiceServer) ApplyCustomFilter(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyCustomFilter not implemented")
}
func (UnimplementedImageServiceServer) ApplyGaussianBlurFilter(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyGaussianBlurFilter not implemented")
}
func (UnimplementedImageServiceServer) ApplyBoxBlurFilter(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyBoxBlurFilter not implemented")
}
func (UnimplementedImageServiceServer) ApplyEdgeDetectionFilter(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyEdgeDetectionFilter not implemented")
}
func (UnimplementedImageServiceServer) ApplySharpenFilter(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplySharpenFilter not implemented")
}
func (UnimplementedImageServiceServer) ApplyNegative(context.Context, *ImageRequest) (*ImageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApplyNegative not implemented")
}
func (UnimplementedImageServiceServer) mustEmbedUnimplementedImageServiceServer() {}

// UnsafeImageServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImageServiceServer will
// result in compilation errors.
type UnsafeImageServiceServer interface {
	mustEmbedUnimplementedImageServiceServer()
}

func RegisterImageServiceServer(s grpc.ServiceRegistrar, srv ImageServiceServer) {
	s.RegisterService(&ImageService_ServiceDesc, srv)
}

func _ImageService_ApplyCustomFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ApplyCustomFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageService/ApplyCustomFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ApplyCustomFilter(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ApplyGaussianBlurFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ApplyGaussianBlurFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageService/ApplyGaussianBlurFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ApplyGaussianBlurFilter(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ApplyBoxBlurFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ApplyBoxBlurFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageService/ApplyBoxBlurFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ApplyBoxBlurFilter(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ApplyEdgeDetectionFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ApplyEdgeDetectionFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageService/ApplyEdgeDetectionFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ApplyEdgeDetectionFilter(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ApplySharpenFilter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ApplySharpenFilter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageService/ApplySharpenFilter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ApplySharpenFilter(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImageService_ApplyNegative_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ImageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImageServiceServer).ApplyNegative(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ImageService/ApplyNegative",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImageServiceServer).ApplyNegative(ctx, req.(*ImageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImageService_ServiceDesc is the grpc.ServiceDesc for ImageService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImageService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ImageService",
	HandlerType: (*ImageServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ApplyCustomFilter",
			Handler:    _ImageService_ApplyCustomFilter_Handler,
		},
		{
			MethodName: "ApplyGaussianBlurFilter",
			Handler:    _ImageService_ApplyGaussianBlurFilter_Handler,
		},
		{
			MethodName: "ApplyBoxBlurFilter",
			Handler:    _ImageService_ApplyBoxBlurFilter_Handler,
		},
		{
			MethodName: "ApplyEdgeDetectionFilter",
			Handler:    _ImageService_ApplyEdgeDetectionFilter_Handler,
		},
		{
			MethodName: "ApplySharpenFilter",
			Handler:    _ImageService_ApplySharpenFilter_Handler,
		},
		{
			MethodName: "ApplyNegative",
			Handler:    _ImageService_ApplyNegative_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/messages.proto",
}
