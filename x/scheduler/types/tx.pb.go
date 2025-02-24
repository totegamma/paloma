// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: palomachain/paloma/scheduler/tx.proto

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

type MsgCreateJob struct {
	Creator  string            `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"` // Deprecated: Do not use.
	Job      *Job              `protobuf:"bytes,2,opt,name=job,proto3" json:"job,omitempty"`
	Metadata types.MsgMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata"`
}

func (m *MsgCreateJob) Reset()         { *m = MsgCreateJob{} }
func (m *MsgCreateJob) String() string { return proto.CompactTextString(m) }
func (*MsgCreateJob) ProtoMessage()    {}
func (*MsgCreateJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f63022306b4a0a9, []int{0}
}

func (m *MsgCreateJob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgCreateJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateJob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgCreateJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateJob.Merge(m, src)
}

func (m *MsgCreateJob) XXX_Size() int {
	return m.Size()
}

func (m *MsgCreateJob) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateJob.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateJob proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *MsgCreateJob) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateJob) GetJob() *Job {
	if m != nil {
		return m.Job
	}
	return nil
}

func (m *MsgCreateJob) GetMetadata() types.MsgMetadata {
	if m != nil {
		return m.Metadata
	}
	return types.MsgMetadata{}
}

type MsgCreateJobResponse struct{}

func (m *MsgCreateJobResponse) Reset()         { *m = MsgCreateJobResponse{} }
func (m *MsgCreateJobResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateJobResponse) ProtoMessage()    {}
func (*MsgCreateJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f63022306b4a0a9, []int{1}
}

func (m *MsgCreateJobResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgCreateJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateJobResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgCreateJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateJobResponse.Merge(m, src)
}

func (m *MsgCreateJobResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgCreateJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateJobResponse proto.InternalMessageInfo

type MsgExecuteJob struct {
	Creator  string            `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"` // Deprecated: Do not use.
	JobID    string            `protobuf:"bytes,2,opt,name=jobID,proto3" json:"jobID,omitempty"`
	Payload  []byte            `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Metadata types.MsgMetadata `protobuf:"bytes,4,opt,name=metadata,proto3" json:"metadata"`
}

func (m *MsgExecuteJob) Reset()         { *m = MsgExecuteJob{} }
func (m *MsgExecuteJob) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteJob) ProtoMessage()    {}
func (*MsgExecuteJob) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f63022306b4a0a9, []int{2}
}

func (m *MsgExecuteJob) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgExecuteJob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteJob.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgExecuteJob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteJob.Merge(m, src)
}

func (m *MsgExecuteJob) XXX_Size() int {
	return m.Size()
}

func (m *MsgExecuteJob) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteJob.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteJob proto.InternalMessageInfo

// Deprecated: Do not use.
func (m *MsgExecuteJob) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgExecuteJob) GetJobID() string {
	if m != nil {
		return m.JobID
	}
	return ""
}

func (m *MsgExecuteJob) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *MsgExecuteJob) GetMetadata() types.MsgMetadata {
	if m != nil {
		return m.Metadata
	}
	return types.MsgMetadata{}
}

type MsgExecuteJobResponse struct {
	// points to the ID of the consensus queue message created from the job execution
	MessageID uint64 `protobuf:"varint,1,opt,name=messageID,proto3" json:"messageID,omitempty"`
}

func (m *MsgExecuteJobResponse) Reset()         { *m = MsgExecuteJobResponse{} }
func (m *MsgExecuteJobResponse) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteJobResponse) ProtoMessage()    {}
func (*MsgExecuteJobResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5f63022306b4a0a9, []int{3}
}

func (m *MsgExecuteJobResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}

func (m *MsgExecuteJobResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteJobResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}

func (m *MsgExecuteJobResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteJobResponse.Merge(m, src)
}

func (m *MsgExecuteJobResponse) XXX_Size() int {
	return m.Size()
}

func (m *MsgExecuteJobResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteJobResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteJobResponse proto.InternalMessageInfo

func (m *MsgExecuteJobResponse) GetMessageID() uint64 {
	if m != nil {
		return m.MessageID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgCreateJob)(nil), "palomachain.paloma.scheduler.MsgCreateJob")
	proto.RegisterType((*MsgCreateJobResponse)(nil), "palomachain.paloma.scheduler.MsgCreateJobResponse")
	proto.RegisterType((*MsgExecuteJob)(nil), "palomachain.paloma.scheduler.MsgExecuteJob")
	proto.RegisterType((*MsgExecuteJobResponse)(nil), "palomachain.paloma.scheduler.MsgExecuteJobResponse")
}

func init() {
	proto.RegisterFile("palomachain/paloma/scheduler/tx.proto", fileDescriptor_5f63022306b4a0a9)
}

var fileDescriptor_5f63022306b4a0a9 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x41, 0x6b, 0xe2, 0x50,
	0x10, 0xc7, 0xf3, 0xd4, 0x5d, 0x37, 0xb3, 0xee, 0x25, 0xb8, 0x4b, 0x10, 0xc9, 0xba, 0xc2, 0x8a,
	0xb4, 0x90, 0x80, 0xd2, 0x2f, 0x60, 0x2d, 0x54, 0x21, 0x97, 0x1c, 0x7b, 0x7b, 0x89, 0x8f, 0xa7,
	0x36, 0xc9, 0x84, 0xbc, 0x58, 0xf4, 0x5b, 0xf4, 0x5b, 0xf4, 0xd0, 0x2f, 0xe2, 0xd1, 0x63, 0x4f,
	0xa5, 0xc4, 0x2f, 0x52, 0x4c, 0x8c, 0x46, 0x10, 0x6b, 0xe9, 0x6d, 0x66, 0xf8, 0xcd, 0xe4, 0xff,
	0x9f, 0xcc, 0x83, 0xff, 0x01, 0x75, 0xd1, 0xa3, 0xce, 0x98, 0x4e, 0x7c, 0x23, 0x8d, 0x0d, 0xe1,
	0x8c, 0xd9, 0x68, 0xe6, 0xb2, 0xd0, 0x88, 0xe6, 0x7a, 0x10, 0x62, 0x84, 0x4a, 0x3d, 0x87, 0xe9,
	0x69, 0xac, 0xef, 0xb0, 0x5a, 0xeb, 0xe4, 0x90, 0x29, 0xda, 0xe9, 0x94, 0xa3, 0xdc, 0x03, 0x75,
	0x05, 0x8b, 0x0c, 0x07, 0x3d, 0x0f, 0xfd, 0x2d, 0x57, 0xe5, 0xc8, 0x31, 0x09, 0x8d, 0x4d, 0x94,
	0x56, 0x9b, 0xcf, 0x04, 0x2a, 0xa6, 0xe0, 0xd7, 0x21, 0xa3, 0x11, 0x1b, 0xa2, 0xad, 0xd4, 0xa1,
	0xec, 0x6c, 0x12, 0x0c, 0x55, 0xd2, 0x20, 0x6d, 0xb9, 0x57, 0x50, 0x89, 0x95, 0x95, 0x94, 0x2e,
	0x14, 0xa7, 0x68, 0xab, 0x85, 0x06, 0x69, 0xff, 0xec, 0xfc, 0xd3, 0x4f, 0x19, 0xd0, 0x87, 0x68,
	0x5b, 0x1b, 0x5a, 0xb9, 0x85, 0x1f, 0x1e, 0x8b, 0xe8, 0x88, 0x46, 0x54, 0x2d, 0x26, 0x9d, 0xad,
	0x63, 0x9d, 0xa9, 0x68, 0xdd, 0x14, 0xdc, 0xdc, 0xd2, 0xbd, 0xd2, 0xf2, 0xf5, 0xaf, 0x64, 0xed,
	0xba, 0x9b, 0x7f, 0xa0, 0x9a, 0x17, 0x6b, 0x31, 0x11, 0xa0, 0x2f, 0x58, 0xf3, 0x89, 0xc0, 0x2f,
	0x53, 0xf0, 0x9b, 0x39, 0x73, 0x66, 0xe7, 0xd8, 0xa8, 0xc2, 0xb7, 0x29, 0xda, 0x83, 0x7e, 0x62,
	0x44, 0xb6, 0xd2, 0x44, 0x51, 0xa1, 0x1c, 0xd0, 0x85, 0x8b, 0x74, 0x94, 0xc8, 0xac, 0x58, 0x59,
	0x7a, 0xe0, 0xa0, 0xf4, 0x25, 0x07, 0x57, 0xf0, 0xfb, 0x40, 0x68, 0x66, 0x41, 0xa9, 0x83, 0xec,
	0x31, 0x21, 0x28, 0x67, 0x83, 0x7e, 0x22, 0xb9, 0x64, 0xed, 0x0b, 0x9d, 0x98, 0x40, 0xd1, 0x14,
	0x5c, 0xb9, 0x07, 0x79, 0xff, 0xab, 0x2e, 0x4e, 0xef, 0x3f, 0xbf, 0xa9, 0x5a, 0xe7, 0x7c, 0x76,
	0x27, 0xc9, 0x07, 0xc8, 0x6d, 0xf4, 0xf2, 0xc3, 0x09, 0x7b, 0xb8, 0xd6, 0xfd, 0x04, 0x9c, 0x7d,
	0xaf, 0x37, 0x58, 0xc6, 0x1a, 0x59, 0xc5, 0x1a, 0x79, 0x8b, 0x35, 0xf2, 0xb8, 0xd6, 0xa4, 0xd5,
	0x5a, 0x93, 0x5e, 0xd6, 0x9a, 0x74, 0x67, 0xf0, 0x49, 0x34, 0x9e, 0xd9, 0xba, 0x83, 0x9e, 0x71,
	0xe4, 0xdc, 0xe7, 0xf9, 0xd7, 0xb5, 0x08, 0x98, 0xb0, 0xbf, 0x27, 0xd7, 0xdd, 0x7d, 0x0f, 0x00,
	0x00, 0xff, 0xff, 0x30, 0xf7, 0x97, 0xc4, 0x8a, 0x03, 0x00, 0x00,
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
	CreateJob(ctx context.Context, in *MsgCreateJob, opts ...grpc.CallOption) (*MsgCreateJobResponse, error)
	ExecuteJob(ctx context.Context, in *MsgExecuteJob, opts ...grpc.CallOption) (*MsgExecuteJobResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateJob(ctx context.Context, in *MsgCreateJob, opts ...grpc.CallOption) (*MsgCreateJobResponse, error) {
	out := new(MsgCreateJobResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.scheduler.Msg/CreateJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ExecuteJob(ctx context.Context, in *MsgExecuteJob, opts ...grpc.CallOption) (*MsgExecuteJobResponse, error) {
	out := new(MsgExecuteJobResponse)
	err := c.cc.Invoke(ctx, "/palomachain.paloma.scheduler.Msg/ExecuteJob", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateJob(context.Context, *MsgCreateJob) (*MsgCreateJobResponse, error)
	ExecuteJob(context.Context, *MsgExecuteJob) (*MsgExecuteJobResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct{}

func (*UnimplementedMsgServer) CreateJob(ctx context.Context, req *MsgCreateJob) (*MsgCreateJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateJob not implemented")
}

func (*UnimplementedMsgServer) ExecuteJob(ctx context.Context, req *MsgExecuteJob) (*MsgExecuteJobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteJob not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.scheduler.Msg/CreateJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateJob(ctx, req.(*MsgCreateJob))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ExecuteJob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgExecuteJob)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ExecuteJob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/palomachain.paloma.scheduler.Msg/ExecuteJob",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ExecuteJob(ctx, req.(*MsgExecuteJob))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "palomachain.paloma.scheduler.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateJob",
			Handler:    _Msg_CreateJob_Handler,
		},
		{
			MethodName: "ExecuteJob",
			Handler:    _Msg_ExecuteJob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "palomachain/paloma/scheduler/tx.proto",
}

func (m *MsgCreateJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateJob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateJob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	dAtA[i] = 0x1a
	if m.Job != nil {
		{
			size, err := m.Job.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
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

func (m *MsgCreateJobResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateJobResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateJobResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgExecuteJob) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteJob) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteJob) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Payload) > 0 {
		i -= len(m.Payload)
		copy(dAtA[i:], m.Payload)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Payload)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.JobID) > 0 {
		i -= len(m.JobID)
		copy(dAtA[i:], m.JobID)
		i = encodeVarintTx(dAtA, i, uint64(len(m.JobID)))
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

func (m *MsgExecuteJobResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteJobResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteJobResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MessageID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.MessageID))
		i--
		dAtA[i] = 0x8
	}
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

func (m *MsgCreateJob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Job != nil {
		l = m.Job.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateJobResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgExecuteJob) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.JobID)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Payload)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Metadata.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgExecuteJobResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.MessageID != 0 {
		n += 1 + sovTx(uint64(m.MessageID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}

func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}

func (m *MsgCreateJob) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateJob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateJob: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Job", wireType)
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
			if m.Job == nil {
				m.Job = &Job{}
			}
			if err := m.Job.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
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

func (m *MsgCreateJobResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateJobResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateJobResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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

func (m *MsgExecuteJob) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExecuteJob: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteJob: illegal tag %d (wire type %d)", fieldNum, wire)
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
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
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
			m.JobID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payload", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Payload = append(m.Payload[:0], dAtA[iNdEx:postIndex]...)
			if m.Payload == nil {
				m.Payload = []byte{}
			}
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

func (m *MsgExecuteJobResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgExecuteJobResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteJobResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MessageID", wireType)
			}
			m.MessageID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MessageID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
