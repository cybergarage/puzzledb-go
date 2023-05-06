// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: metric.proto

package grpc

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
	Metric_ListMetric_FullMethodName = "/Metric/ListMetric"
	Metric_GetMetric_FullMethodName  = "/Metric/GetMetric"
)

// MetricClient is the client API for Metric service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricClient interface {
	ListMetric(ctx context.Context, in *ListMetricRequest, opts ...grpc.CallOption) (*ListMetricResponse, error)
	GetMetric(ctx context.Context, in *GetMetricRequest, opts ...grpc.CallOption) (*GetMetricResponse, error)
}

type metricClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricClient(cc grpc.ClientConnInterface) MetricClient {
	return &metricClient{cc}
}

func (c *metricClient) ListMetric(ctx context.Context, in *ListMetricRequest, opts ...grpc.CallOption) (*ListMetricResponse, error) {
	out := new(ListMetricResponse)
	err := c.cc.Invoke(ctx, Metric_ListMetric_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *metricClient) GetMetric(ctx context.Context, in *GetMetricRequest, opts ...grpc.CallOption) (*GetMetricResponse, error) {
	out := new(GetMetricResponse)
	err := c.cc.Invoke(ctx, Metric_GetMetric_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricServer is the server API for Metric service.
// All implementations must embed UnimplementedMetricServer
// for forward compatibility
type MetricServer interface {
	ListMetric(context.Context, *ListMetricRequest) (*ListMetricResponse, error)
	GetMetric(context.Context, *GetMetricRequest) (*GetMetricResponse, error)
	mustEmbedUnimplementedMetricServer()
}

// UnimplementedMetricServer must be embedded to have forward compatible implementations.
type UnimplementedMetricServer struct {
}

func (UnimplementedMetricServer) ListMetric(context.Context, *ListMetricRequest) (*ListMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMetric not implemented")
}
func (UnimplementedMetricServer) GetMetric(context.Context, *GetMetricRequest) (*GetMetricResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMetric not implemented")
}
func (UnimplementedMetricServer) mustEmbedUnimplementedMetricServer() {}

// UnsafeMetricServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricServer will
// result in compilation errors.
type UnsafeMetricServer interface {
	mustEmbedUnimplementedMetricServer()
}

func RegisterMetricServer(s grpc.ServiceRegistrar, srv MetricServer) {
	s.RegisterService(&Metric_ServiceDesc, srv)
}

func _Metric_ListMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).ListMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_ListMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).ListMetric(ctx, req.(*ListMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Metric_GetMetric_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMetricRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricServer).GetMetric(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Metric_GetMetric_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricServer).GetMetric(ctx, req.(*GetMetricRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Metric_ServiceDesc is the grpc.ServiceDesc for Metric service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Metric_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "Metric",
	HandlerType: (*MetricServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMetric",
			Handler:    _Metric_ListMetric_Handler,
		},
		{
			MethodName: "GetMetric",
			Handler:    _Metric_GetMetric_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "metric.proto",
}
