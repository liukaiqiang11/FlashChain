// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

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

// ChaincodeSupportClient is the client API for ChaincodeSupport service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaincodeSupportClient interface {
	StartPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	EndPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
	GetBlockInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockchainInfo, error)
	ValidateBlock(ctx context.Context, in *ValidateMessage, opts ...grpc.CallOption) (*ValidateMessage, error)
}

type chaincodeSupportClient struct {
	cc grpc.ClientConnInterface
}

func NewChaincodeSupportClient(cc grpc.ClientConnInterface) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) StartPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.ChaincodeSupport/StartPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeSupportClient) EndPeer(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	out := new(Empty)
	err := c.cc.Invoke(ctx, "/pb.ChaincodeSupport/EndPeer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := c.cc.NewStream(ctx, &ChaincodeSupport_ServiceDesc.Streams[0], "/pb.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *chaincodeSupportClient) GetBlockInfo(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*BlockchainInfo, error) {
	out := new(BlockchainInfo)
	err := c.cc.Invoke(ctx, "/pb.ChaincodeSupport/GetBlockInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *chaincodeSupportClient) ValidateBlock(ctx context.Context, in *ValidateMessage, opts ...grpc.CallOption) (*ValidateMessage, error) {
	out := new(ValidateMessage)
	err := c.cc.Invoke(ctx, "/pb.ChaincodeSupport/ValidateBlock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ChaincodeSupportServer is the server API for ChaincodeSupport service.
// All implementations must embed UnimplementedChaincodeSupportServer
// for forward compatibility
type ChaincodeSupportServer interface {
	StartPeer(context.Context, *Empty) (*Empty, error)
	EndPeer(context.Context, *Empty) (*Empty, error)
	Register(ChaincodeSupport_RegisterServer) error
	GetBlockInfo(context.Context, *Empty) (*BlockchainInfo, error)
	ValidateBlock(context.Context, *ValidateMessage) (*ValidateMessage, error)
	mustEmbedUnimplementedChaincodeSupportServer()
}

// UnimplementedChaincodeSupportServer must be embedded to have forward compatible implementations.
type UnimplementedChaincodeSupportServer struct {
}

func (UnimplementedChaincodeSupportServer) StartPeer(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartPeer not implemented")
}
func (UnimplementedChaincodeSupportServer) EndPeer(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EndPeer not implemented")
}
func (UnimplementedChaincodeSupportServer) Register(ChaincodeSupport_RegisterServer) error {
	return status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedChaincodeSupportServer) GetBlockInfo(context.Context, *Empty) (*BlockchainInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlockInfo not implemented")
}
func (UnimplementedChaincodeSupportServer) ValidateBlock(context.Context, *ValidateMessage) (*ValidateMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidateBlock not implemented")
}
func (UnimplementedChaincodeSupportServer) mustEmbedUnimplementedChaincodeSupportServer() {}

// UnsafeChaincodeSupportServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaincodeSupportServer will
// result in compilation errors.
type UnsafeChaincodeSupportServer interface {
	mustEmbedUnimplementedChaincodeSupportServer()
}

func RegisterChaincodeSupportServer(s grpc.ServiceRegistrar, srv ChaincodeSupportServer) {
	s.RegisterService(&ChaincodeSupport_ServiceDesc, srv)
}

func _ChaincodeSupport_StartPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeSupportServer).StartPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChaincodeSupport/StartPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeSupportServer).StartPeer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaincodeSupport_EndPeer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeSupportServer).EndPeer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChaincodeSupport/EndPeer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeSupportServer).EndPeer(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _ChaincodeSupport_GetBlockInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeSupportServer).GetBlockInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChaincodeSupport/GetBlockInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeSupportServer).GetBlockInfo(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ChaincodeSupport_ValidateBlock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ChaincodeSupportServer).ValidateBlock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.ChaincodeSupport/ValidateBlock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ChaincodeSupportServer).ValidateBlock(ctx, req.(*ValidateMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// ChaincodeSupport_ServiceDesc is the grpc.ServiceDesc for ChaincodeSupport service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ChaincodeSupport_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StartPeer",
			Handler:    _ChaincodeSupport_StartPeer_Handler,
		},
		{
			MethodName: "EndPeer",
			Handler:    _ChaincodeSupport_EndPeer_Handler,
		},
		{
			MethodName: "GetBlockInfo",
			Handler:    _ChaincodeSupport_GetBlockInfo_Handler,
		},
		{
			MethodName: "ValidateBlock",
			Handler:    _ChaincodeSupport_ValidateBlock_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chaincode_shim.proto",
}

// ChaincodeClient is the client API for Chaincode service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ChaincodeClient interface {
	Connect(ctx context.Context, opts ...grpc.CallOption) (Chaincode_ConnectClient, error)
}

type chaincodeClient struct {
	cc grpc.ClientConnInterface
}

func NewChaincodeClient(cc grpc.ClientConnInterface) ChaincodeClient {
	return &chaincodeClient{cc}
}

func (c *chaincodeClient) Connect(ctx context.Context, opts ...grpc.CallOption) (Chaincode_ConnectClient, error) {
	stream, err := c.cc.NewStream(ctx, &Chaincode_ServiceDesc.Streams[0], "/pb.Chaincode/Connect", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeConnectClient{stream}
	return x, nil
}

type Chaincode_ConnectClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeConnectClient struct {
	grpc.ClientStream
}

func (x *chaincodeConnectClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeConnectClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChaincodeServer is the server API for Chaincode service.
// All implementations must embed UnimplementedChaincodeServer
// for forward compatibility
type ChaincodeServer interface {
	Connect(Chaincode_ConnectServer) error
	mustEmbedUnimplementedChaincodeServer()
}

// UnimplementedChaincodeServer must be embedded to have forward compatible implementations.
type UnimplementedChaincodeServer struct {
}

func (UnimplementedChaincodeServer) Connect(Chaincode_ConnectServer) error {
	return status.Errorf(codes.Unimplemented, "method Connect not implemented")
}
func (UnimplementedChaincodeServer) mustEmbedUnimplementedChaincodeServer() {}

// UnsafeChaincodeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ChaincodeServer will
// result in compilation errors.
type UnsafeChaincodeServer interface {
	mustEmbedUnimplementedChaincodeServer()
}

func RegisterChaincodeServer(s grpc.ServiceRegistrar, srv ChaincodeServer) {
	s.RegisterService(&Chaincode_ServiceDesc, srv)
}

func _Chaincode_Connect_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeServer).Connect(&chaincodeConnectServer{stream})
}

type Chaincode_ConnectServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeConnectServer struct {
	grpc.ServerStream
}

func (x *chaincodeConnectServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeConnectServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Chaincode_ServiceDesc is the grpc.ServiceDesc for Chaincode service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Chaincode_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Chaincode",
	HandlerType: (*ChaincodeServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Connect",
			Handler:       _Chaincode_Connect_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chaincode_shim.proto",
}
