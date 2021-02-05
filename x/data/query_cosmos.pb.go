// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package data

import (
	context "context"
	types "github.com/regen-network/regen-ledger/types"
	grpc "google.golang.org/grpc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryClient interface {
	// ByHash queries data based on its ContentHash.
	ByHash(ctx context.Context, in *QueryByHashRequest, opts ...grpc.CallOption) (*QueryByHashResponse, error)
	// BySigner queries data based on signers.
	BySigner(ctx context.Context, in *QueryBySignerRequest, opts ...grpc.CallOption) (*QueryBySignerResponse, error)
	ConvertToCompactDataset(ctx context.Context, in *ConvertToCompactDatasetRequest, opts ...grpc.CallOption) (*ConvertToCompactDatasetResponse, error)
}

type queryClient struct {
	cc                       grpc.ClientConnInterface
	_ByHash                  types.Invoker
	_BySigner                types.Invoker
	_ConvertToCompactDataset types.Invoker
}

func NewQueryClient(cc grpc.ClientConnInterface) QueryClient {
	return &queryClient{cc: cc}
}

func (c *queryClient) ByHash(ctx context.Context, in *QueryByHashRequest, opts ...grpc.CallOption) (*QueryByHashResponse, error) {
	if invoker := c._ByHash; invoker != nil {
		var out QueryByHashResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._ByHash, err = invokerConn.Invoker("/regen.data.v1alpha2.Query/ByHash")
		if err != nil {
			var out QueryByHashResponse
			err = c._ByHash(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryByHashResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha2.Query/ByHash", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) BySigner(ctx context.Context, in *QueryBySignerRequest, opts ...grpc.CallOption) (*QueryBySignerResponse, error) {
	if invoker := c._BySigner; invoker != nil {
		var out QueryBySignerResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._BySigner, err = invokerConn.Invoker("/regen.data.v1alpha2.Query/BySigner")
		if err != nil {
			var out QueryBySignerResponse
			err = c._BySigner(ctx, in, &out)
			return &out, err
		}
	}
	out := new(QueryBySignerResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha2.Query/BySigner", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) ConvertToCompactDataset(ctx context.Context, in *ConvertToCompactDatasetRequest, opts ...grpc.CallOption) (*ConvertToCompactDatasetResponse, error) {
	if invoker := c._ConvertToCompactDataset; invoker != nil {
		var out ConvertToCompactDatasetResponse
		err := invoker(ctx, in, &out)
		return &out, err
	}
	if invokerConn, ok := c.cc.(types.InvokerConn); ok {
		var err error
		c._ConvertToCompactDataset, err = invokerConn.Invoker("/regen.data.v1alpha2.Query/ConvertToCompactDataset")
		if err != nil {
			var out ConvertToCompactDatasetResponse
			err = c._ConvertToCompactDataset(ctx, in, &out)
			return &out, err
		}
	}
	out := new(ConvertToCompactDatasetResponse)
	err := c.cc.Invoke(ctx, "/regen.data.v1alpha2.Query/ConvertToCompactDataset", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ByHash queries data based on its ContentHash.
	ByHash(types.Context, *QueryByHashRequest) (*QueryByHashResponse, error)
	// BySigner queries data based on signers.
	BySigner(types.Context, *QueryBySignerRequest) (*QueryBySignerResponse, error)
	ConvertToCompactDataset(types.Context, *ConvertToCompactDatasetRequest) (*ConvertToCompactDatasetResponse, error)
}

func RegisterQueryServer(s grpc.ServiceRegistrar, srv QueryServer) {
	s.RegisterService(&Query_ServiceDesc, srv)
}

func _Query_ByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ByHash(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha2.Query/ByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ByHash(types.UnwrapSDKContext(ctx), req.(*QueryByHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_BySigner_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryBySignerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).BySigner(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha2.Query/BySigner",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).BySigner(types.UnwrapSDKContext(ctx), req.(*QueryBySignerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_ConvertToCompactDataset_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConvertToCompactDatasetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).ConvertToCompactDataset(types.UnwrapSDKContext(ctx), in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/regen.data.v1alpha2.Query/ConvertToCompactDataset",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).ConvertToCompactDataset(types.UnwrapSDKContext(ctx), req.(*ConvertToCompactDatasetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Query_ServiceDesc is the grpc.ServiceDesc for Query service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Query_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "regen.data.v1alpha2.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ByHash",
			Handler:    _Query_ByHash_Handler,
		},
		{
			MethodName: "BySigner",
			Handler:    _Query_BySigner_Handler,
		},
		{
			MethodName: "ConvertToCompactDataset",
			Handler:    _Query_ConvertToCompactDataset_Handler,
		},
	},
	Metadata: "regen/data/v1alpha2/query.proto",
}

const (
	QueryByHashMethod                  = "/regen.data.v1alpha2.Query/ByHash"
	QueryBySignerMethod                = "/regen.data.v1alpha2.Query/BySigner"
	QueryConvertToCompactDatasetMethod = "/regen.data.v1alpha2.Query/ConvertToCompactDataset"
)
