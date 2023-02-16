// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package schemav2

import (
	context "context"
	schema "github.com/codenotary/immudb/pkg/api/schema"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ImmuServiceClient is the client API for ImmuService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ImmuServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	DocumentInsert(ctx context.Context, in *DocumentInsertRequest, opts ...grpc.CallOption) (*schema.VerifiableTx, error)
	DocumentSearch(ctx context.Context, in *DocumentSearchRequest, opts ...grpc.CallOption) (*DocumentSearchResponse, error)
	CollectionCreate(ctx context.Context, in *CollectionCreateRequest, opts ...grpc.CallOption) (*CollectionInformation, error)
	CollectionList(ctx context.Context, in *CollectionListRequest, opts ...grpc.CallOption) (*CollectionListResponse, error)
	CollectionDelete(ctx context.Context, in *CollectionDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
}

type immuServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewImmuServiceClient(cc grpc.ClientConnInterface) ImmuServiceClient {
	return &immuServiceClient{cc}
}

func (c *immuServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuService/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) DocumentInsert(ctx context.Context, in *DocumentInsertRequest, opts ...grpc.CallOption) (*schema.VerifiableTx, error) {
	out := new(schema.VerifiableTx)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuService/DocumentInsert", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) DocumentSearch(ctx context.Context, in *DocumentSearchRequest, opts ...grpc.CallOption) (*DocumentSearchResponse, error) {
	out := new(DocumentSearchResponse)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuService/DocumentSearch", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) CollectionCreate(ctx context.Context, in *CollectionCreateRequest, opts ...grpc.CallOption) (*CollectionInformation, error) {
	out := new(CollectionInformation)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuService/CollectionCreate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) CollectionList(ctx context.Context, in *CollectionListRequest, opts ...grpc.CallOption) (*CollectionListResponse, error) {
	out := new(CollectionListResponse)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuService/CollectionList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *immuServiceClient) CollectionDelete(ctx context.Context, in *CollectionDeleteRequest, opts ...grpc.CallOption) (*empty.Empty, error) {
	out := new(empty.Empty)
	err := c.cc.Invoke(ctx, "/immudb.schemav2.ImmuService/CollectionDelete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImmuServiceServer is the server API for ImmuService service.
// All implementations should embed UnimplementedImmuServiceServer
// for forward compatibility
type ImmuServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	DocumentInsert(context.Context, *DocumentInsertRequest) (*schema.VerifiableTx, error)
	DocumentSearch(context.Context, *DocumentSearchRequest) (*DocumentSearchResponse, error)
	CollectionCreate(context.Context, *CollectionCreateRequest) (*CollectionInformation, error)
	CollectionList(context.Context, *CollectionListRequest) (*CollectionListResponse, error)
	CollectionDelete(context.Context, *CollectionDeleteRequest) (*empty.Empty, error)
}

// UnimplementedImmuServiceServer should be embedded to have forward compatible implementations.
type UnimplementedImmuServiceServer struct {
}

func (UnimplementedImmuServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedImmuServiceServer) DocumentInsert(context.Context, *DocumentInsertRequest) (*schema.VerifiableTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocumentInsert not implemented")
}
func (UnimplementedImmuServiceServer) DocumentSearch(context.Context, *DocumentSearchRequest) (*DocumentSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DocumentSearch not implemented")
}
func (UnimplementedImmuServiceServer) CollectionCreate(context.Context, *CollectionCreateRequest) (*CollectionInformation, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionCreate not implemented")
}
func (UnimplementedImmuServiceServer) CollectionList(context.Context, *CollectionListRequest) (*CollectionListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionList not implemented")
}
func (UnimplementedImmuServiceServer) CollectionDelete(context.Context, *CollectionDeleteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectionDelete not implemented")
}

// UnsafeImmuServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ImmuServiceServer will
// result in compilation errors.
type UnsafeImmuServiceServer interface {
	mustEmbedUnimplementedImmuServiceServer()
}

func RegisterImmuServiceServer(s grpc.ServiceRegistrar, srv ImmuServiceServer) {
	s.RegisterService(&ImmuService_ServiceDesc, srv)
}

func _ImmuService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuService/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_DocumentInsert_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentInsertRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).DocumentInsert(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuService/DocumentInsert",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).DocumentInsert(ctx, req.(*DocumentInsertRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_DocumentSearch_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DocumentSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).DocumentSearch(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuService/DocumentSearch",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).DocumentSearch(ctx, req.(*DocumentSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_CollectionCreate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).CollectionCreate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuService/CollectionCreate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).CollectionCreate(ctx, req.(*CollectionCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_CollectionList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).CollectionList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuService/CollectionList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).CollectionList(ctx, req.(*CollectionListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ImmuService_CollectionDelete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CollectionDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImmuServiceServer).CollectionDelete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/immudb.schemav2.ImmuService/CollectionDelete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImmuServiceServer).CollectionDelete(ctx, req.(*CollectionDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ImmuService_ServiceDesc is the grpc.ServiceDesc for ImmuService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ImmuService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "immudb.schemav2.ImmuService",
	HandlerType: (*ImmuServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _ImmuService_Login_Handler,
		},
		{
			MethodName: "DocumentInsert",
			Handler:    _ImmuService_DocumentInsert_Handler,
		},
		{
			MethodName: "DocumentSearch",
			Handler:    _ImmuService_DocumentSearch_Handler,
		},
		{
			MethodName: "CollectionCreate",
			Handler:    _ImmuService_CollectionCreate_Handler,
		},
		{
			MethodName: "CollectionList",
			Handler:    _ImmuService_CollectionList_Handler,
		},
		{
			MethodName: "CollectionDelete",
			Handler:    _ImmuService_CollectionDelete_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "schemav2.proto",
}
