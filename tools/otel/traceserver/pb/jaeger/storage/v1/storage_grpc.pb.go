/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

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

// SpanWriterPluginClient is the client API for SpanWriterPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpanWriterPluginClient interface {
	// spanstore/Writer
	WriteSpan(ctx context.Context, in *WriteSpanRequest, opts ...grpc.CallOption) (*WriteSpanResponse, error)
	Close(ctx context.Context, in *CloseWriterRequest, opts ...grpc.CallOption) (*CloseWriterResponse, error)
}

type spanWriterPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSpanWriterPluginClient(cc grpc.ClientConnInterface) SpanWriterPluginClient {
	return &spanWriterPluginClient{cc}
}

func (c *spanWriterPluginClient) WriteSpan(ctx context.Context, in *WriteSpanRequest, opts ...grpc.CallOption) (*WriteSpanResponse, error) {
	out := new(WriteSpanResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.SpanWriterPlugin/WriteSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spanWriterPluginClient) Close(ctx context.Context, in *CloseWriterRequest, opts ...grpc.CallOption) (*CloseWriterResponse, error) {
	out := new(CloseWriterResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.SpanWriterPlugin/Close", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpanWriterPluginServer is the server API for SpanWriterPlugin service.
// All implementations must embed UnimplementedSpanWriterPluginServer
// for forward compatibility
type SpanWriterPluginServer interface {
	// spanstore/Writer
	WriteSpan(context.Context, *WriteSpanRequest) (*WriteSpanResponse, error)
	Close(context.Context, *CloseWriterRequest) (*CloseWriterResponse, error)
	mustEmbedUnimplementedSpanWriterPluginServer()
}

// UnimplementedSpanWriterPluginServer must be embedded to have forward compatible implementations.
type UnimplementedSpanWriterPluginServer struct {
}

func (UnimplementedSpanWriterPluginServer) WriteSpan(context.Context, *WriteSpanRequest) (*WriteSpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteSpan not implemented")
}
func (UnimplementedSpanWriterPluginServer) Close(context.Context, *CloseWriterRequest) (*CloseWriterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Close not implemented")
}
func (UnimplementedSpanWriterPluginServer) mustEmbedUnimplementedSpanWriterPluginServer() {}

// UnsafeSpanWriterPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpanWriterPluginServer will
// result in compilation errors.
type UnsafeSpanWriterPluginServer interface {
	mustEmbedUnimplementedSpanWriterPluginServer()
}

func RegisterSpanWriterPluginServer(s grpc.ServiceRegistrar, srv SpanWriterPluginServer) {
	s.RegisterService(&SpanWriterPlugin_ServiceDesc, srv)
}

func _SpanWriterPlugin_WriteSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteSpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanWriterPluginServer).WriteSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.SpanWriterPlugin/WriteSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanWriterPluginServer).WriteSpan(ctx, req.(*WriteSpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpanWriterPlugin_Close_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CloseWriterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanWriterPluginServer).Close(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.SpanWriterPlugin/Close",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanWriterPluginServer).Close(ctx, req.(*CloseWriterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpanWriterPlugin_ServiceDesc is the grpc.ServiceDesc for SpanWriterPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpanWriterPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v1.SpanWriterPlugin",
	HandlerType: (*SpanWriterPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteSpan",
			Handler:    _SpanWriterPlugin_WriteSpan_Handler,
		},
		{
			MethodName: "Close",
			Handler:    _SpanWriterPlugin_Close_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jaeger/storage/v1/storage.proto",
}

// StreamingSpanWriterPluginClient is the client API for StreamingSpanWriterPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingSpanWriterPluginClient interface {
	WriteSpanStream(ctx context.Context, opts ...grpc.CallOption) (StreamingSpanWriterPlugin_WriteSpanStreamClient, error)
}

type streamingSpanWriterPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingSpanWriterPluginClient(cc grpc.ClientConnInterface) StreamingSpanWriterPluginClient {
	return &streamingSpanWriterPluginClient{cc}
}

func (c *streamingSpanWriterPluginClient) WriteSpanStream(ctx context.Context, opts ...grpc.CallOption) (StreamingSpanWriterPlugin_WriteSpanStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingSpanWriterPlugin_ServiceDesc.Streams[0], "/jaeger.storage.v1.StreamingSpanWriterPlugin/WriteSpanStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingSpanWriterPluginWriteSpanStreamClient{stream}
	return x, nil
}

type StreamingSpanWriterPlugin_WriteSpanStreamClient interface {
	Send(*WriteSpanRequest) error
	CloseAndRecv() (*WriteSpanResponse, error)
	grpc.ClientStream
}

type streamingSpanWriterPluginWriteSpanStreamClient struct {
	grpc.ClientStream
}

func (x *streamingSpanWriterPluginWriteSpanStreamClient) Send(m *WriteSpanRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingSpanWriterPluginWriteSpanStreamClient) CloseAndRecv() (*WriteSpanResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(WriteSpanResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingSpanWriterPluginServer is the server API for StreamingSpanWriterPlugin service.
// All implementations must embed UnimplementedStreamingSpanWriterPluginServer
// for forward compatibility
type StreamingSpanWriterPluginServer interface {
	WriteSpanStream(StreamingSpanWriterPlugin_WriteSpanStreamServer) error
	mustEmbedUnimplementedStreamingSpanWriterPluginServer()
}

// UnimplementedStreamingSpanWriterPluginServer must be embedded to have forward compatible implementations.
type UnimplementedStreamingSpanWriterPluginServer struct {
}

func (UnimplementedStreamingSpanWriterPluginServer) WriteSpanStream(StreamingSpanWriterPlugin_WriteSpanStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method WriteSpanStream not implemented")
}
func (UnimplementedStreamingSpanWriterPluginServer) mustEmbedUnimplementedStreamingSpanWriterPluginServer() {
}

// UnsafeStreamingSpanWriterPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingSpanWriterPluginServer will
// result in compilation errors.
type UnsafeStreamingSpanWriterPluginServer interface {
	mustEmbedUnimplementedStreamingSpanWriterPluginServer()
}

func RegisterStreamingSpanWriterPluginServer(s grpc.ServiceRegistrar, srv StreamingSpanWriterPluginServer) {
	s.RegisterService(&StreamingSpanWriterPlugin_ServiceDesc, srv)
}

func _StreamingSpanWriterPlugin_WriteSpanStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingSpanWriterPluginServer).WriteSpanStream(&streamingSpanWriterPluginWriteSpanStreamServer{stream})
}

type StreamingSpanWriterPlugin_WriteSpanStreamServer interface {
	SendAndClose(*WriteSpanResponse) error
	Recv() (*WriteSpanRequest, error)
	grpc.ServerStream
}

type streamingSpanWriterPluginWriteSpanStreamServer struct {
	grpc.ServerStream
}

func (x *streamingSpanWriterPluginWriteSpanStreamServer) SendAndClose(m *WriteSpanResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingSpanWriterPluginWriteSpanStreamServer) Recv() (*WriteSpanRequest, error) {
	m := new(WriteSpanRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingSpanWriterPlugin_ServiceDesc is the grpc.ServiceDesc for StreamingSpanWriterPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingSpanWriterPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v1.StreamingSpanWriterPlugin",
	HandlerType: (*StreamingSpanWriterPluginServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WriteSpanStream",
			Handler:       _StreamingSpanWriterPlugin_WriteSpanStream_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "jaeger/storage/v1/storage.proto",
}

// SpanReaderPluginClient is the client API for SpanReaderPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SpanReaderPluginClient interface {
	// spanstore/Reader
	GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (SpanReaderPlugin_GetTraceClient, error)
	GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error)
	GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error)
	FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (SpanReaderPlugin_FindTracesClient, error)
	FindTraceIDs(ctx context.Context, in *FindTraceIDsRequest, opts ...grpc.CallOption) (*FindTraceIDsResponse, error)
}

type spanReaderPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewSpanReaderPluginClient(cc grpc.ClientConnInterface) SpanReaderPluginClient {
	return &spanReaderPluginClient{cc}
}

func (c *spanReaderPluginClient) GetTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (SpanReaderPlugin_GetTraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpanReaderPlugin_ServiceDesc.Streams[0], "/jaeger.storage.v1.SpanReaderPlugin/GetTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &spanReaderPluginGetTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpanReaderPlugin_GetTraceClient interface {
	Recv() (*SpansResponseChunk, error)
	grpc.ClientStream
}

type spanReaderPluginGetTraceClient struct {
	grpc.ClientStream
}

func (x *spanReaderPluginGetTraceClient) Recv() (*SpansResponseChunk, error) {
	m := new(SpansResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *spanReaderPluginClient) GetServices(ctx context.Context, in *GetServicesRequest, opts ...grpc.CallOption) (*GetServicesResponse, error) {
	out := new(GetServicesResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.SpanReaderPlugin/GetServices", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spanReaderPluginClient) GetOperations(ctx context.Context, in *GetOperationsRequest, opts ...grpc.CallOption) (*GetOperationsResponse, error) {
	out := new(GetOperationsResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.SpanReaderPlugin/GetOperations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *spanReaderPluginClient) FindTraces(ctx context.Context, in *FindTracesRequest, opts ...grpc.CallOption) (SpanReaderPlugin_FindTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &SpanReaderPlugin_ServiceDesc.Streams[1], "/jaeger.storage.v1.SpanReaderPlugin/FindTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &spanReaderPluginFindTracesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type SpanReaderPlugin_FindTracesClient interface {
	Recv() (*SpansResponseChunk, error)
	grpc.ClientStream
}

type spanReaderPluginFindTracesClient struct {
	grpc.ClientStream
}

func (x *spanReaderPluginFindTracesClient) Recv() (*SpansResponseChunk, error) {
	m := new(SpansResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *spanReaderPluginClient) FindTraceIDs(ctx context.Context, in *FindTraceIDsRequest, opts ...grpc.CallOption) (*FindTraceIDsResponse, error) {
	out := new(FindTraceIDsResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.SpanReaderPlugin/FindTraceIDs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SpanReaderPluginServer is the server API for SpanReaderPlugin service.
// All implementations must embed UnimplementedSpanReaderPluginServer
// for forward compatibility
type SpanReaderPluginServer interface {
	// spanstore/Reader
	GetTrace(*GetTraceRequest, SpanReaderPlugin_GetTraceServer) error
	GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error)
	GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error)
	FindTraces(*FindTracesRequest, SpanReaderPlugin_FindTracesServer) error
	FindTraceIDs(context.Context, *FindTraceIDsRequest) (*FindTraceIDsResponse, error)
	mustEmbedUnimplementedSpanReaderPluginServer()
}

// UnimplementedSpanReaderPluginServer must be embedded to have forward compatible implementations.
type UnimplementedSpanReaderPluginServer struct {
}

func (UnimplementedSpanReaderPluginServer) GetTrace(*GetTraceRequest, SpanReaderPlugin_GetTraceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetTrace not implemented")
}
func (UnimplementedSpanReaderPluginServer) GetServices(context.Context, *GetServicesRequest) (*GetServicesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServices not implemented")
}
func (UnimplementedSpanReaderPluginServer) GetOperations(context.Context, *GetOperationsRequest) (*GetOperationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOperations not implemented")
}
func (UnimplementedSpanReaderPluginServer) FindTraces(*FindTracesRequest, SpanReaderPlugin_FindTracesServer) error {
	return status.Errorf(codes.Unimplemented, "method FindTraces not implemented")
}
func (UnimplementedSpanReaderPluginServer) FindTraceIDs(context.Context, *FindTraceIDsRequest) (*FindTraceIDsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindTraceIDs not implemented")
}
func (UnimplementedSpanReaderPluginServer) mustEmbedUnimplementedSpanReaderPluginServer() {}

// UnsafeSpanReaderPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SpanReaderPluginServer will
// result in compilation errors.
type UnsafeSpanReaderPluginServer interface {
	mustEmbedUnimplementedSpanReaderPluginServer()
}

func RegisterSpanReaderPluginServer(s grpc.ServiceRegistrar, srv SpanReaderPluginServer) {
	s.RegisterService(&SpanReaderPlugin_ServiceDesc, srv)
}

func _SpanReaderPlugin_GetTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpanReaderPluginServer).GetTrace(m, &spanReaderPluginGetTraceServer{stream})
}

type SpanReaderPlugin_GetTraceServer interface {
	Send(*SpansResponseChunk) error
	grpc.ServerStream
}

type spanReaderPluginGetTraceServer struct {
	grpc.ServerStream
}

func (x *spanReaderPluginGetTraceServer) Send(m *SpansResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _SpanReaderPlugin_GetServices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetServicesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanReaderPluginServer).GetServices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.SpanReaderPlugin/GetServices",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanReaderPluginServer).GetServices(ctx, req.(*GetServicesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpanReaderPlugin_GetOperations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOperationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanReaderPluginServer).GetOperations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.SpanReaderPlugin/GetOperations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanReaderPluginServer).GetOperations(ctx, req.(*GetOperationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SpanReaderPlugin_FindTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(FindTracesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(SpanReaderPluginServer).FindTraces(m, &spanReaderPluginFindTracesServer{stream})
}

type SpanReaderPlugin_FindTracesServer interface {
	Send(*SpansResponseChunk) error
	grpc.ServerStream
}

type spanReaderPluginFindTracesServer struct {
	grpc.ServerStream
}

func (x *spanReaderPluginFindTracesServer) Send(m *SpansResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

func _SpanReaderPlugin_FindTraceIDs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FindTraceIDsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SpanReaderPluginServer).FindTraceIDs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.SpanReaderPlugin/FindTraceIDs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SpanReaderPluginServer).FindTraceIDs(ctx, req.(*FindTraceIDsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SpanReaderPlugin_ServiceDesc is the grpc.ServiceDesc for SpanReaderPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SpanReaderPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v1.SpanReaderPlugin",
	HandlerType: (*SpanReaderPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetServices",
			Handler:    _SpanReaderPlugin_GetServices_Handler,
		},
		{
			MethodName: "GetOperations",
			Handler:    _SpanReaderPlugin_GetOperations_Handler,
		},
		{
			MethodName: "FindTraceIDs",
			Handler:    _SpanReaderPlugin_FindTraceIDs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetTrace",
			Handler:       _SpanReaderPlugin_GetTrace_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "FindTraces",
			Handler:       _SpanReaderPlugin_FindTraces_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "jaeger/storage/v1/storage.proto",
}

// ArchiveSpanWriterPluginClient is the client API for ArchiveSpanWriterPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchiveSpanWriterPluginClient interface {
	// spanstore/Writer
	WriteArchiveSpan(ctx context.Context, in *WriteSpanRequest, opts ...grpc.CallOption) (*WriteSpanResponse, error)
}

type archiveSpanWriterPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveSpanWriterPluginClient(cc grpc.ClientConnInterface) ArchiveSpanWriterPluginClient {
	return &archiveSpanWriterPluginClient{cc}
}

func (c *archiveSpanWriterPluginClient) WriteArchiveSpan(ctx context.Context, in *WriteSpanRequest, opts ...grpc.CallOption) (*WriteSpanResponse, error) {
	out := new(WriteSpanResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.ArchiveSpanWriterPlugin/WriteArchiveSpan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ArchiveSpanWriterPluginServer is the server API for ArchiveSpanWriterPlugin service.
// All implementations must embed UnimplementedArchiveSpanWriterPluginServer
// for forward compatibility
type ArchiveSpanWriterPluginServer interface {
	// spanstore/Writer
	WriteArchiveSpan(context.Context, *WriteSpanRequest) (*WriteSpanResponse, error)
	mustEmbedUnimplementedArchiveSpanWriterPluginServer()
}

// UnimplementedArchiveSpanWriterPluginServer must be embedded to have forward compatible implementations.
type UnimplementedArchiveSpanWriterPluginServer struct {
}

func (UnimplementedArchiveSpanWriterPluginServer) WriteArchiveSpan(context.Context, *WriteSpanRequest) (*WriteSpanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WriteArchiveSpan not implemented")
}
func (UnimplementedArchiveSpanWriterPluginServer) mustEmbedUnimplementedArchiveSpanWriterPluginServer() {
}

// UnsafeArchiveSpanWriterPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchiveSpanWriterPluginServer will
// result in compilation errors.
type UnsafeArchiveSpanWriterPluginServer interface {
	mustEmbedUnimplementedArchiveSpanWriterPluginServer()
}

func RegisterArchiveSpanWriterPluginServer(s grpc.ServiceRegistrar, srv ArchiveSpanWriterPluginServer) {
	s.RegisterService(&ArchiveSpanWriterPlugin_ServiceDesc, srv)
}

func _ArchiveSpanWriterPlugin_WriteArchiveSpan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WriteSpanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ArchiveSpanWriterPluginServer).WriteArchiveSpan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.ArchiveSpanWriterPlugin/WriteArchiveSpan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ArchiveSpanWriterPluginServer).WriteArchiveSpan(ctx, req.(*WriteSpanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ArchiveSpanWriterPlugin_ServiceDesc is the grpc.ServiceDesc for ArchiveSpanWriterPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArchiveSpanWriterPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v1.ArchiveSpanWriterPlugin",
	HandlerType: (*ArchiveSpanWriterPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "WriteArchiveSpan",
			Handler:    _ArchiveSpanWriterPlugin_WriteArchiveSpan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jaeger/storage/v1/storage.proto",
}

// ArchiveSpanReaderPluginClient is the client API for ArchiveSpanReaderPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ArchiveSpanReaderPluginClient interface {
	// spanstore/Reader
	GetArchiveTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (ArchiveSpanReaderPlugin_GetArchiveTraceClient, error)
}

type archiveSpanReaderPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewArchiveSpanReaderPluginClient(cc grpc.ClientConnInterface) ArchiveSpanReaderPluginClient {
	return &archiveSpanReaderPluginClient{cc}
}

func (c *archiveSpanReaderPluginClient) GetArchiveTrace(ctx context.Context, in *GetTraceRequest, opts ...grpc.CallOption) (ArchiveSpanReaderPlugin_GetArchiveTraceClient, error) {
	stream, err := c.cc.NewStream(ctx, &ArchiveSpanReaderPlugin_ServiceDesc.Streams[0], "/jaeger.storage.v1.ArchiveSpanReaderPlugin/GetArchiveTrace", opts...)
	if err != nil {
		return nil, err
	}
	x := &archiveSpanReaderPluginGetArchiveTraceClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ArchiveSpanReaderPlugin_GetArchiveTraceClient interface {
	Recv() (*SpansResponseChunk, error)
	grpc.ClientStream
}

type archiveSpanReaderPluginGetArchiveTraceClient struct {
	grpc.ClientStream
}

func (x *archiveSpanReaderPluginGetArchiveTraceClient) Recv() (*SpansResponseChunk, error) {
	m := new(SpansResponseChunk)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ArchiveSpanReaderPluginServer is the server API for ArchiveSpanReaderPlugin service.
// All implementations must embed UnimplementedArchiveSpanReaderPluginServer
// for forward compatibility
type ArchiveSpanReaderPluginServer interface {
	// spanstore/Reader
	GetArchiveTrace(*GetTraceRequest, ArchiveSpanReaderPlugin_GetArchiveTraceServer) error
	mustEmbedUnimplementedArchiveSpanReaderPluginServer()
}

// UnimplementedArchiveSpanReaderPluginServer must be embedded to have forward compatible implementations.
type UnimplementedArchiveSpanReaderPluginServer struct {
}

func (UnimplementedArchiveSpanReaderPluginServer) GetArchiveTrace(*GetTraceRequest, ArchiveSpanReaderPlugin_GetArchiveTraceServer) error {
	return status.Errorf(codes.Unimplemented, "method GetArchiveTrace not implemented")
}
func (UnimplementedArchiveSpanReaderPluginServer) mustEmbedUnimplementedArchiveSpanReaderPluginServer() {
}

// UnsafeArchiveSpanReaderPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ArchiveSpanReaderPluginServer will
// result in compilation errors.
type UnsafeArchiveSpanReaderPluginServer interface {
	mustEmbedUnimplementedArchiveSpanReaderPluginServer()
}

func RegisterArchiveSpanReaderPluginServer(s grpc.ServiceRegistrar, srv ArchiveSpanReaderPluginServer) {
	s.RegisterService(&ArchiveSpanReaderPlugin_ServiceDesc, srv)
}

func _ArchiveSpanReaderPlugin_GetArchiveTrace_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTraceRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ArchiveSpanReaderPluginServer).GetArchiveTrace(m, &archiveSpanReaderPluginGetArchiveTraceServer{stream})
}

type ArchiveSpanReaderPlugin_GetArchiveTraceServer interface {
	Send(*SpansResponseChunk) error
	grpc.ServerStream
}

type archiveSpanReaderPluginGetArchiveTraceServer struct {
	grpc.ServerStream
}

func (x *archiveSpanReaderPluginGetArchiveTraceServer) Send(m *SpansResponseChunk) error {
	return x.ServerStream.SendMsg(m)
}

// ArchiveSpanReaderPlugin_ServiceDesc is the grpc.ServiceDesc for ArchiveSpanReaderPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ArchiveSpanReaderPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v1.ArchiveSpanReaderPlugin",
	HandlerType: (*ArchiveSpanReaderPluginServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetArchiveTrace",
			Handler:       _ArchiveSpanReaderPlugin_GetArchiveTrace_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "jaeger/storage/v1/storage.proto",
}

// DependenciesReaderPluginClient is the client API for DependenciesReaderPlugin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DependenciesReaderPluginClient interface {
	// dependencystore/Reader
	GetDependencies(ctx context.Context, in *GetDependenciesRequest, opts ...grpc.CallOption) (*GetDependenciesResponse, error)
}

type dependenciesReaderPluginClient struct {
	cc grpc.ClientConnInterface
}

func NewDependenciesReaderPluginClient(cc grpc.ClientConnInterface) DependenciesReaderPluginClient {
	return &dependenciesReaderPluginClient{cc}
}

func (c *dependenciesReaderPluginClient) GetDependencies(ctx context.Context, in *GetDependenciesRequest, opts ...grpc.CallOption) (*GetDependenciesResponse, error) {
	out := new(GetDependenciesResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.DependenciesReaderPlugin/GetDependencies", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DependenciesReaderPluginServer is the server API for DependenciesReaderPlugin service.
// All implementations must embed UnimplementedDependenciesReaderPluginServer
// for forward compatibility
type DependenciesReaderPluginServer interface {
	// dependencystore/Reader
	GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error)
	mustEmbedUnimplementedDependenciesReaderPluginServer()
}

// UnimplementedDependenciesReaderPluginServer must be embedded to have forward compatible implementations.
type UnimplementedDependenciesReaderPluginServer struct {
}

func (UnimplementedDependenciesReaderPluginServer) GetDependencies(context.Context, *GetDependenciesRequest) (*GetDependenciesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDependencies not implemented")
}
func (UnimplementedDependenciesReaderPluginServer) mustEmbedUnimplementedDependenciesReaderPluginServer() {
}

// UnsafeDependenciesReaderPluginServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DependenciesReaderPluginServer will
// result in compilation errors.
type UnsafeDependenciesReaderPluginServer interface {
	mustEmbedUnimplementedDependenciesReaderPluginServer()
}

func RegisterDependenciesReaderPluginServer(s grpc.ServiceRegistrar, srv DependenciesReaderPluginServer) {
	s.RegisterService(&DependenciesReaderPlugin_ServiceDesc, srv)
}

func _DependenciesReaderPlugin_GetDependencies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDependenciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DependenciesReaderPluginServer).GetDependencies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.DependenciesReaderPlugin/GetDependencies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DependenciesReaderPluginServer).GetDependencies(ctx, req.(*GetDependenciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DependenciesReaderPlugin_ServiceDesc is the grpc.ServiceDesc for DependenciesReaderPlugin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DependenciesReaderPlugin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v1.DependenciesReaderPlugin",
	HandlerType: (*DependenciesReaderPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDependencies",
			Handler:    _DependenciesReaderPlugin_GetDependencies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jaeger/storage/v1/storage.proto",
}

// PluginCapabilitiesClient is the client API for PluginCapabilities service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PluginCapabilitiesClient interface {
	Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error)
}

type pluginCapabilitiesClient struct {
	cc grpc.ClientConnInterface
}

func NewPluginCapabilitiesClient(cc grpc.ClientConnInterface) PluginCapabilitiesClient {
	return &pluginCapabilitiesClient{cc}
}

func (c *pluginCapabilitiesClient) Capabilities(ctx context.Context, in *CapabilitiesRequest, opts ...grpc.CallOption) (*CapabilitiesResponse, error) {
	out := new(CapabilitiesResponse)
	err := c.cc.Invoke(ctx, "/jaeger.storage.v1.PluginCapabilities/Capabilities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PluginCapabilitiesServer is the server API for PluginCapabilities service.
// All implementations must embed UnimplementedPluginCapabilitiesServer
// for forward compatibility
type PluginCapabilitiesServer interface {
	Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error)
	mustEmbedUnimplementedPluginCapabilitiesServer()
}

// UnimplementedPluginCapabilitiesServer must be embedded to have forward compatible implementations.
type UnimplementedPluginCapabilitiesServer struct {
}

func (UnimplementedPluginCapabilitiesServer) Capabilities(context.Context, *CapabilitiesRequest) (*CapabilitiesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Capabilities not implemented")
}
func (UnimplementedPluginCapabilitiesServer) mustEmbedUnimplementedPluginCapabilitiesServer() {}

// UnsafePluginCapabilitiesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PluginCapabilitiesServer will
// result in compilation errors.
type UnsafePluginCapabilitiesServer interface {
	mustEmbedUnimplementedPluginCapabilitiesServer()
}

func RegisterPluginCapabilitiesServer(s grpc.ServiceRegistrar, srv PluginCapabilitiesServer) {
	s.RegisterService(&PluginCapabilities_ServiceDesc, srv)
}

func _PluginCapabilities_Capabilities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CapabilitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PluginCapabilitiesServer).Capabilities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jaeger.storage.v1.PluginCapabilities/Capabilities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PluginCapabilitiesServer).Capabilities(ctx, req.(*CapabilitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PluginCapabilities_ServiceDesc is the grpc.ServiceDesc for PluginCapabilities service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PluginCapabilities_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jaeger.storage.v1.PluginCapabilities",
	HandlerType: (*PluginCapabilitiesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Capabilities",
			Handler:    _PluginCapabilities_Capabilities_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jaeger/storage/v1/storage.proto",
}
