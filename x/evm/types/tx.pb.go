// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/evm/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	types "github.com/palomachain/paloma/x/valset/types"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = proto.Marshal
	_ = fmt.Errorf
	_ = math.Inf
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgDeployNewSmartContractRequest struct {
	Creator     string            `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"` // Deprecated: Do not use.
	Title       string            `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Description string            `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	AbiJSON     string            `protobuf:"bytes,4,opt,name=abiJSON,proto3" json:"abiJSON,omitempty"`
	BytecodeHex string            `protobuf:"bytes,5,opt,name=bytecodeHex,proto3" json:"bytecodeHex,omitempty"`
	Metadata    types.MsgMetadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata"`
}

func (m *MsgDeployNewSmartContractRequest) Reset()         { *m = MsgDeployNewSmartContractRequest{} }
func (m *MsgDeployNewSmartContractRequest) String() string { return proto.CompactTextString(m) }
func (*MsgDeployNewSmartContractRequest) ProtoMessage()    {}
func (*MsgDeployNewSmartContractRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{0}
}

func (m *MsgDeployNewSmartContractRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgDeployNewSmartContractRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeployNewSmartContractRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgDeployNewSmartContractRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeployNewSmartContractRequest.Merge(m, src)
}

func (m *MsgDeployNewSmartContractRequest) XXX_Size() int {
	return m.Size()
}

func (m *MsgDeployNewSmartContractRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeployNewSmartContractRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeployNewSmartContractRequest proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *MsgDeployNewSmartContractRequest) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetAbiJSON() string {
	if m != nil {
		return m.AbiJSON
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetBytecodeHex() string {
	if m != nil {
		return m.BytecodeHex
	}
	return ""
}

func (m *MsgDeployNewSmartContractRequest) GetMetadata() types.MsgMetadata {
	if m != nil {
		return m.Metadata
	}
	return types.MsgMetadata{}
}

type DeployNewSmartContractResponse struct{}

func (m *DeployNewSmartContractResponse) Reset()         { *m = DeployNewSmartContractResponse{} }
func (m *DeployNewSmartContractResponse) String() string { return proto.CompactTextString(m) }
func (*DeployNewSmartContractResponse) ProtoMessage()    {}
func (*DeployNewSmartContractResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{1}
}

func (m *DeployNewSmartContractResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *DeployNewSmartContractResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DeployNewSmartContractResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *DeployNewSmartContractResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeployNewSmartContractResponse.Merge(m, src)
}

func (m *DeployNewSmartContractResponse) XXX_Size() int {
	return m.Size()
}

func (m *DeployNewSmartContractResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeployNewSmartContractResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeployNewSmartContractResponse proto.InternalMessageInfo

type MsgRemoveSmartContractDeploymentRequest struct {
	Sender           string            `protobuf:"bytes,1,opt,name=Sender,proto3" json:"Sender,omitempty"` // Deprecated: Do not use.
	SmartContractID  uint64            `protobuf:"varint,2,opt,name=smartContractID,proto3" json:"smartContractID,omitempty"`
	ChainReferenceID string            `protobuf:"bytes,3,opt,name=chainReferenceID,proto3" json:"chainReferenceID,omitempty"`
	Metadata         types.MsgMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata"`
}

func (m *MsgRemoveSmartContractDeploymentRequest) Reset() {
	*m = MsgRemoveSmartContractDeploymentRequest{}
}
func (m *MsgRemoveSmartContractDeploymentRequest) String() string { return proto.CompactTextString(m) }
func (*MsgRemoveSmartContractDeploymentRequest) ProtoMessage()    {}
func (*MsgRemoveSmartContractDeploymentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{2}
}

func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest.Merge(m, src)
}

func (m *MsgRemoveSmartContractDeploymentRequest) XXX_Size() int {
	return m.Size()
}

func (m *MsgRemoveSmartContractDeploymentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRemoveSmartContractDeploymentRequest proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *MsgRemoveSmartContractDeploymentRequest) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *MsgRemoveSmartContractDeploymentRequest) GetSmartContractID() uint64 {
	if m != nil {
		return m.SmartContractID
	}
	return 0
}

func (m *MsgRemoveSmartContractDeploymentRequest) GetChainReferenceID() string {
	if m != nil {
		return m.ChainReferenceID
	}
	return ""
}

func (m *MsgRemoveSmartContractDeploymentRequest) GetMetadata() types.MsgMetadata {
	if m != nil {
		return m.Metadata
	}
	return types.MsgMetadata{}
}

type RemoveSmartContractDeploymentResponse struct{}

func (m *RemoveSmartContractDeploymentResponse) Reset()         { *m = RemoveSmartContractDeploymentResponse{} }
func (m *RemoveSmartContractDeploymentResponse) String() string { return proto.CompactTextString(m) }
func (*RemoveSmartContractDeploymentResponse) ProtoMessage()    {}
func (*RemoveSmartContractDeploymentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_631cfc68eb1fd278, []int{3}
}

func (m *RemoveSmartContractDeploymentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *RemoveSmartContractDeploymentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RemoveSmartContractDeploymentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *RemoveSmartContractDeploymentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RemoveSmartContractDeploymentResponse.Merge(m, src)
}

func (m *RemoveSmartContractDeploymentResponse) XXX_Size() int {
	return m.Size()
}

func (m *RemoveSmartContractDeploymentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RemoveSmartContractDeploymentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RemoveSmartContractDeploymentResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgDeployNewSmartContractRequest)(nil), "palomachain.paloma.evm.MsgDeployNewSmartContractRequest")
	proto.RegisterType((*DeployNewSmartContractResponse)(nil), "palomachain.paloma.evm.DeployNewSmartContractResponse")
	proto.RegisterType((*MsgRemoveSmartContractDeploymentRequest)(nil), "palomachain.paloma.evm.MsgRemoveSmartContractDeploymentRequest")
	proto.RegisterType((*RemoveSmartContractDeploymentResponse)(nil), "palomachain.paloma.evm.RemoveSmartContractDeploymentResponse")
}

func init() { proto.RegisterFile("palomachain/paloma/evm/tx.proto", fileDescriptor_631cfc68eb1fd278) }

var fileDescriptor_631cfc68eb1fd278 = []byte{
	// 467 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0xae, 0xeb, 0xe0, 0xdb, 0x01, 0x64, 0x4d, 0x53, 0x14, 0x41, 0x16, 0x55, 0x62,
	0x2b, 0x1c, 0x12, 0x69, 0x48, 0x88, 0x0b, 0x42, 0xea, 0x7a, 0xd8, 0x90, 0x32, 0xa4, 0xf4, 0xc6,
	0xcd, 0x75, 0x3f, 0xb2, 0x48, 0xb5, 0x1d, 0x62, 0x2f, 0xb4, 0xaf, 0xc0, 0x89, 0x37, 0xe0, 0x25,
	0x78, 0x88, 0x1d, 0x77, 0xe4, 0x84, 0x50, 0x2b, 0x5e, 0x03, 0xa1, 0xc6, 0xdd, 0x94, 0x41, 0xda,
	0x4d, 0xda, 0xed, 0xcb, 0xff, 0xfb, 0xff, 0x2d, 0x7f, 0x3f, 0xdb, 0x81, 0xbd, 0x8c, 0x8d, 0x95,
	0x60, 0xfc, 0x8c, 0xa5, 0x32, 0xb4, 0x75, 0x88, 0x85, 0x08, 0xcd, 0x24, 0xc8, 0x72, 0x65, 0x14,
	0xdd, 0xad, 0x18, 0x02, 0x5b, 0x07, 0x58, 0x08, 0x77, 0x27, 0x51, 0x89, 0x2a, 0x2d, 0xe1, 0xa2,
	0xb2, 0x6e, 0x77, 0xbf, 0x66, 0xb9, 0x82, 0x8d, 0x35, 0x9a, 0x90, 0x2b, 0x21, 0x94, 0xb4, 0xbe,
	0xce, 0x1f, 0x02, 0x7e, 0xa4, 0x93, 0x3e, 0x66, 0x63, 0x35, 0x3d, 0xc5, 0xcf, 0x03, 0xc1, 0x72,
	0x73, 0xa4, 0xa4, 0xc9, 0x19, 0x37, 0x31, 0x7e, 0x3a, 0x47, 0x6d, 0xe8, 0x13, 0xd8, 0xe2, 0x39,
	0x32, 0xa3, 0x72, 0x87, 0xf8, 0xa4, 0xfb, 0xb0, 0xd7, 0x74, 0x48, 0x7c, 0x25, 0xd1, 0x1d, 0xd8,
	0x34, 0xa9, 0x19, 0xa3, 0xd3, 0x5c, 0xf4, 0x62, 0xfb, 0x41, 0x7d, 0xd8, 0x1e, 0xa1, 0xe6, 0x79,
	0x9a, 0x99, 0x54, 0x49, 0x67, 0xa3, 0xec, 0x55, 0x25, 0xea, 0xc0, 0x16, 0x1b, 0xa6, 0xef, 0x06,
	0xef, 0x4f, 0x9d, 0x56, 0xd9, 0xbd, 0xfa, 0x5c, 0x64, 0x87, 0x53, 0x83, 0x5c, 0x8d, 0xf0, 0x18,
	0x27, 0xce, 0xa6, 0xcd, 0x56, 0x24, 0x7a, 0x0c, 0x0f, 0x04, 0x1a, 0x36, 0x62, 0x86, 0x39, 0x6d,
	0x9f, 0x74, 0xb7, 0x0f, 0xf7, 0x83, 0x1a, 0x3e, 0x76, 0xe2, 0x20, 0xd2, 0x49, 0xb4, 0x74, 0xf7,
	0x5a, 0x17, 0x3f, 0xf7, 0x1a, 0xf1, 0x75, 0xba, 0xe3, 0x83, 0xb7, 0x6a, 0x78, 0x9d, 0x29, 0xa9,
	0xb1, 0xf3, 0x9b, 0xc0, 0x41, 0xa4, 0x93, 0x18, 0x85, 0x2a, 0xf0, 0x86, 0xc5, 0x06, 0x05, 0xca,
	0x6b, 0x52, 0x2e, 0xb4, 0x07, 0x28, 0x47, 0x58, 0x05, 0xb5, 0x54, 0x68, 0x17, 0x1e, 0xe9, 0x6a,
	0xfa, 0xa4, 0x5f, 0x12, 0x6b, 0xc5, 0xff, 0xca, 0xf4, 0x05, 0x3c, 0x2e, 0xc7, 0x88, 0xf1, 0x23,
	0xe6, 0x28, 0x39, 0x9e, 0xf4, 0x97, 0x00, 0xff, 0xd3, 0x6f, 0x90, 0x68, 0xdd, 0x8b, 0xc4, 0x01,
	0x3c, 0xbb, 0x65, 0x46, 0x0b, 0xe4, 0xf0, 0x7b, 0x13, 0x36, 0x22, 0x9d, 0xd0, 0x2f, 0x04, 0x76,
	0xeb, 0xd9, 0xd1, 0xd7, 0x41, 0xfd, 0x6d, 0x0d, 0x6e, 0xbb, 0x6b, 0xee, 0xab, 0x55, 0xc9, 0xf5,
	0xa7, 0x44, 0xbf, 0x11, 0x78, 0xba, 0x76, 0xfb, 0xf4, 0xed, 0x9a, 0x3d, 0xdd, 0xe5, 0x70, 0xdd,
	0x37, 0xab, 0x16, 0xb8, 0x13, 0xb6, 0xde, 0xd1, 0xc5, 0xcc, 0x23, 0x97, 0x33, 0x8f, 0xfc, 0x9a,
	0x79, 0xe4, 0xeb, 0xdc, 0x6b, 0x5c, 0xce, 0xbd, 0xc6, 0x8f, 0xb9, 0xd7, 0xf8, 0xf0, 0x3c, 0x49,
	0xcd, 0xd9, 0xf9, 0x30, 0xe0, 0x4a, 0x84, 0x35, 0xef, 0x76, 0x62, 0x7f, 0x04, 0xd3, 0x0c, 0xf5,
	0xb0, 0x5d, 0x3e, 0xdb, 0x97, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x25, 0x0d, 0x30, 0x2f,
	0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ context.Context
	_ grpc.ClientConn
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	DeployNewSmartContract(ctx context.Context, in *MsgDeployNewSmartContractRequest, opts ...grpc.CallOption) (*DeployNewSmartContractResponse, error)
	RemoveSmartContractDeployment(ctx context.Context, in *MsgRemoveSmartContractDeploymentRequest, opts ...grpc.CallOption) (*RemoveSmartContractDeploymentResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) DeployNewSmartContract(ctx context.Context, in *MsgDeployNewSmartContractRequest, opts ...grpc.CallOption) (*DeployNewSmartContractResponse, error) {
	out := new(DeployNewSmartContractResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Msg/DeployNewSmartContract", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RemoveSmartContractDeployment(ctx context.Context, in *MsgRemoveSmartContractDeploymentRequest, opts ...grpc.CallOption) (*RemoveSmartContractDeploymentResponse, error) {
	out := new(RemoveSmartContractDeploymentResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.evm.Msg/RemoveSmartContractDeployment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	DeployNewSmartContract(context.Context, *MsgDeployNewSmartContractRequest) (*DeployNewSmartContractResponse, error)
	RemoveSmartContractDeployment(context.Context, *MsgRemoveSmartContractDeploymentRequest) (*RemoveSmartContractDeploymentResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct{}

func (*UnimplementedMsgServer) DeployNewSmartContract(ctx context.Context, req *MsgDeployNewSmartContractRequest) (*DeployNewSmartContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeployNewSmartContract not implemented")
}

func (*UnimplementedMsgServer) RemoveSmartContractDeployment(ctx context.Context, req *MsgRemoveSmartContractDeploymentRequest) (*RemoveSmartContractDeploymentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveSmartContractDeployment not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_DeployNewSmartContract_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeployNewSmartContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeployNewSmartContract(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Msg/DeployNewSmartContract",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeployNewSmartContract(ctx, req.(*MsgDeployNewSmartContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RemoveSmartContractDeployment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRemoveSmartContractDeploymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RemoveSmartContractDeployment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.evm.Msg/RemoveSmartContractDeployment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RemoveSmartContractDeployment(ctx, req.(*MsgRemoveSmartContractDeploymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "palomachain.paloma.evm.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DeployNewSmartContract",
			Handler:    _Msg_DeployNewSmartContract_Handler,
		},
		{
			MethodName: "RemoveSmartContractDeployment",
			Handler:    _Msg_RemoveSmartContractDeployment_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "palomachain/paloma/evm/tx.proto",
}

func (m *MsgDeployNewSmartContractRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeployNewSmartContractRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeployNewSmartContractRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.BytecodeHex) > 0 {
		i -= len(m.BytecodeHex)
		copy(dAtA[i:], m.BytecodeHex)
		i = encodeVarintTx(dAtA, i, uint64(len(m.BytecodeHex)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AbiJSON) > 0 {
		i -= len(m.AbiJSON)
		copy(dAtA[i:], m.AbiJSON)
		i = encodeVarintTx(dAtA, i, uint64(len(m.AbiJSON)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DeployNewSmartContractResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DeployNewSmartContractResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DeployNewSmartContractResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRemoveSmartContractDeploymentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRemoveSmartContractDeploymentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRemoveSmartContractDeploymentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Metadata.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.ChainReferenceID) > 0 {
		i -= len(m.ChainReferenceID)
		copy(dAtA[i:], m.ChainReferenceID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ChainReferenceID)))
		i--
		dAtA[i] = 0x1a
	}
	if m.SmartContractID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.SmartContractID))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *RemoveSmartContractDeploymentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RemoveSmartContractDeploymentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RemoveSmartContractDeploymentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}

func (m *MsgDeployNewSmartContractRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.AbiJSON)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.BytecodeHex)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *DeployNewSmartContractResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRemoveSmartContractDeploymentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.SmartContractID != 0 {
		n += 1 + sovTx(uint64(m.SmartContractID))
	}
	l = len(m.ChainReferenceID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *RemoveSmartContractDeploymentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *MsgDeployNewSmartContractRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeployNewSmartContractRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeployNewSmartContractRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AbiJSON", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AbiJSON = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BytecodeHex", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BytecodeHex = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *DeployNewSmartContractResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DeployNewSmartContractResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DeployNewSmartContractResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *MsgRemoveSmartContractDeploymentRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRemoveSmartContractDeploymentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRemoveSmartContractDeploymentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field SmartContractID", wireType)
			}
			m.SmartContractID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.SmartContractID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainReferenceID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainReferenceID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Metadata.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func (m *RemoveSmartContractDeploymentResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: RemoveSmartContractDeploymentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RemoveSmartContractDeploymentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}

func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
