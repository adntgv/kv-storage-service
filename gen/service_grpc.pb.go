// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package gen

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

// KeyValueClient is the client API for KeyValue service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type KeyValueClient interface {
	Create(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Reply, error)
	Update(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Reply, error)
	Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Reply, error)
	Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Reply, error)
	GetHistory(ctx context.Context, in *Key, opts ...grpc.CallOption) (*HistoryReply, error)
	Clear(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Reply, error)
}

type keyValueClient struct {
	cc grpc.ClientConnInterface
}

func NewKeyValueClient(cc grpc.ClientConnInterface) KeyValueClient {
	return &keyValueClient{cc}
}

func (c *keyValueClient) Create(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/service.KeyValue/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueClient) Update(ctx context.Context, in *Pair, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/service.KeyValue/Update", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueClient) Get(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/service.KeyValue/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueClient) Delete(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/service.KeyValue/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueClient) GetHistory(ctx context.Context, in *Key, opts ...grpc.CallOption) (*HistoryReply, error) {
	out := new(HistoryReply)
	err := c.cc.Invoke(ctx, "/service.KeyValue/GetHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *keyValueClient) Clear(ctx context.Context, in *Key, opts ...grpc.CallOption) (*Reply, error) {
	out := new(Reply)
	err := c.cc.Invoke(ctx, "/service.KeyValue/Clear", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// KeyValueServer is the server API for KeyValue service.
// All implementations must embed UnimplementedKeyValueServer
// for forward compatibility
type KeyValueServer interface {
	Create(context.Context, *Pair) (*Reply, error)
	Update(context.Context, *Pair) (*Reply, error)
	Get(context.Context, *Key) (*Reply, error)
	Delete(context.Context, *Key) (*Reply, error)
	GetHistory(context.Context, *Key) (*HistoryReply, error)
	Clear(context.Context, *Key) (*Reply, error)
	mustEmbedUnimplementedKeyValueServer()
}

// UnimplementedKeyValueServer must be embedded to have forward compatible implementations.
type UnimplementedKeyValueServer struct {
}

func (UnimplementedKeyValueServer) Create(context.Context, *Pair) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedKeyValueServer) Update(context.Context, *Pair) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedKeyValueServer) Get(context.Context, *Key) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedKeyValueServer) Delete(context.Context, *Key) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedKeyValueServer) GetHistory(context.Context, *Key) (*HistoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHistory not implemented")
}
func (UnimplementedKeyValueServer) Clear(context.Context, *Key) (*Reply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Clear not implemented")
}
func (UnimplementedKeyValueServer) mustEmbedUnimplementedKeyValueServer() {}

// UnsafeKeyValueServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to KeyValueServer will
// result in compilation errors.
type UnsafeKeyValueServer interface {
	mustEmbedUnimplementedKeyValueServer()
}

func RegisterKeyValueServer(s grpc.ServiceRegistrar, srv KeyValueServer) {
	s.RegisterService(&KeyValue_ServiceDesc, srv)
}

func _KeyValue_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KeyValue/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).Create(ctx, req.(*Pair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValue_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Pair)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KeyValue/Update",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).Update(ctx, req.(*Pair))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValue_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KeyValue/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).Get(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValue_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KeyValue/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).Delete(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValue_GetHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).GetHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KeyValue/GetHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).GetHistory(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

func _KeyValue_Clear_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Key)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(KeyValueServer).Clear(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.KeyValue/Clear",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(KeyValueServer).Clear(ctx, req.(*Key))
	}
	return interceptor(ctx, in, info, handler)
}

// KeyValue_ServiceDesc is the grpc.ServiceDesc for KeyValue service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var KeyValue_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.KeyValue",
	HandlerType: (*KeyValueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _KeyValue_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _KeyValue_Update_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _KeyValue_Get_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _KeyValue_Delete_Handler,
		},
		{
			MethodName: "GetHistory",
			Handler:    _KeyValue_GetHistory_Handler,
		},
		{
			MethodName: "Clear",
			Handler:    _KeyValue_Clear_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service.proto",
}
