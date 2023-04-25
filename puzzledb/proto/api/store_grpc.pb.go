// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: store.proto

package api

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
	StoreAPI_ListDatabases_FullMethodName   = "/api.StoreAPI/ListDatabases"
	StoreAPI_ListCollections_FullMethodName = "/api.StoreAPI/ListCollections"
)

// StoreAPIClient is the client API for StoreAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StoreAPIClient interface {
	ListDatabases(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error)
	ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error)
}

type storeAPIClient struct {
	cc grpc.ClientConnInterface
}

func NewStoreAPIClient(cc grpc.ClientConnInterface) StoreAPIClient {
	return &storeAPIClient{cc}
}

func (c *storeAPIClient) ListDatabases(ctx context.Context, in *ListDatabasesRequest, opts ...grpc.CallOption) (*ListDatabasesResponse, error) {
	out := new(ListDatabasesResponse)
	err := c.cc.Invoke(ctx, StoreAPI_ListDatabases_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *storeAPIClient) ListCollections(ctx context.Context, in *ListCollectionsRequest, opts ...grpc.CallOption) (*ListCollectionsResponse, error) {
	out := new(ListCollectionsResponse)
	err := c.cc.Invoke(ctx, StoreAPI_ListCollections_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StoreAPIServer is the server API for StoreAPI service.
// All implementations must embed UnimplementedStoreAPIServer
// for forward compatibility
type StoreAPIServer interface {
	ListDatabases(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error)
	ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error)
	mustEmbedUnimplementedStoreAPIServer()
}

// UnimplementedStoreAPIServer must be embedded to have forward compatible implementations.
type UnimplementedStoreAPIServer struct {
}

func (UnimplementedStoreAPIServer) ListDatabases(context.Context, *ListDatabasesRequest) (*ListDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatabases not implemented")
}
func (UnimplementedStoreAPIServer) ListCollections(context.Context, *ListCollectionsRequest) (*ListCollectionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCollections not implemented")
}
func (UnimplementedStoreAPIServer) mustEmbedUnimplementedStoreAPIServer() {}

// UnsafeStoreAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StoreAPIServer will
// result in compilation errors.
type UnsafeStoreAPIServer interface {
	mustEmbedUnimplementedStoreAPIServer()
}

func RegisterStoreAPIServer(s grpc.ServiceRegistrar, srv StoreAPIServer) {
	s.RegisterService(&StoreAPI_ServiceDesc, srv)
}

func _StoreAPI_ListDatabases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreAPIServer).ListDatabases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreAPI_ListDatabases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreAPIServer).ListDatabases(ctx, req.(*ListDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StoreAPI_ListCollections_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCollectionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StoreAPIServer).ListCollections(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StoreAPI_ListCollections_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StoreAPIServer).ListCollections(ctx, req.(*ListCollectionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StoreAPI_ServiceDesc is the grpc.ServiceDesc for StoreAPI service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StoreAPI_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.StoreAPI",
	HandlerType: (*StoreAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListDatabases",
			Handler:    _StoreAPI_ListDatabases_Handler,
		},
		{
			MethodName: "ListCollections",
			Handler:    _StoreAPI_ListCollections_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "store.proto",
}
