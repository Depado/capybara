// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: pb/capybara.proto

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

// CapybaraClient is the client API for Capybara service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CapybaraClient interface {
	// Acquires a lock
	ClaimLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error)
	// Release a lock
	ReleaseLock(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error)
	// CRUD operations
	Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error)
}

type capybaraClient struct {
	cc grpc.ClientConnInterface
}

func NewCapybaraClient(cc grpc.ClientConnInterface) CapybaraClient {
	return &capybaraClient{cc}
}

func (c *capybaraClient) ClaimLock(ctx context.Context, in *LockRequest, opts ...grpc.CallOption) (*LockResponse, error) {
	out := new(LockResponse)
	err := c.cc.Invoke(ctx, "/pb.Capybara/ClaimLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capybaraClient) ReleaseLock(ctx context.Context, in *ReleaseRequest, opts ...grpc.CallOption) (*ReleaseResponse, error) {
	out := new(ReleaseResponse)
	err := c.cc.Invoke(ctx, "/pb.Capybara/ReleaseLock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capybaraClient) Put(ctx context.Context, in *PutRequest, opts ...grpc.CallOption) (*PutResponse, error) {
	out := new(PutResponse)
	err := c.cc.Invoke(ctx, "/pb.Capybara/Put", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capybaraClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, "/pb.Capybara/Delete", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *capybaraClient) Get(ctx context.Context, in *GetRequest, opts ...grpc.CallOption) (*GetResponse, error) {
	out := new(GetResponse)
	err := c.cc.Invoke(ctx, "/pb.Capybara/Get", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CapybaraServer is the server API for Capybara service.
// All implementations must embed UnimplementedCapybaraServer
// for forward compatibility
type CapybaraServer interface {
	// Acquires a lock
	ClaimLock(context.Context, *LockRequest) (*LockResponse, error)
	// Release a lock
	ReleaseLock(context.Context, *ReleaseRequest) (*ReleaseResponse, error)
	// CRUD operations
	Put(context.Context, *PutRequest) (*PutResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	mustEmbedUnimplementedCapybaraServer()
}

// UnimplementedCapybaraServer must be embedded to have forward compatible implementations.
type UnimplementedCapybaraServer struct {
}

func (UnimplementedCapybaraServer) ClaimLock(context.Context, *LockRequest) (*LockResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ClaimLock not implemented")
}
func (UnimplementedCapybaraServer) ReleaseLock(context.Context, *ReleaseRequest) (*ReleaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReleaseLock not implemented")
}
func (UnimplementedCapybaraServer) Put(context.Context, *PutRequest) (*PutResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Put not implemented")
}
func (UnimplementedCapybaraServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedCapybaraServer) Get(context.Context, *GetRequest) (*GetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Get not implemented")
}
func (UnimplementedCapybaraServer) mustEmbedUnimplementedCapybaraServer() {}

// UnsafeCapybaraServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CapybaraServer will
// result in compilation errors.
type UnsafeCapybaraServer interface {
	mustEmbedUnimplementedCapybaraServer()
}

func RegisterCapybaraServer(s grpc.ServiceRegistrar, srv CapybaraServer) {
	s.RegisterService(&Capybara_ServiceDesc, srv)
}

func _Capybara_ClaimLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LockRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapybaraServer).ClaimLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Capybara/ClaimLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapybaraServer).ClaimLock(ctx, req.(*LockRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capybara_ReleaseLock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapybaraServer).ReleaseLock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Capybara/ReleaseLock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapybaraServer).ReleaseLock(ctx, req.(*ReleaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capybara_Put_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapybaraServer).Put(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Capybara/Put",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapybaraServer).Put(ctx, req.(*PutRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capybara_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapybaraServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Capybara/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapybaraServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Capybara_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CapybaraServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Capybara/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CapybaraServer).Get(ctx, req.(*GetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Capybara_ServiceDesc is the grpc.ServiceDesc for Capybara service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Capybara_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Capybara",
	HandlerType: (*CapybaraServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ClaimLock",
			Handler:    _Capybara_ClaimLock_Handler,
		},
		{
			MethodName: "ReleaseLock",
			Handler:    _Capybara_ReleaseLock_Handler,
		},
		{
			MethodName: "Put",
			Handler:    _Capybara_Put_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _Capybara_Delete_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Capybara_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/capybara.proto",
}
