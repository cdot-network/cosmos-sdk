// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cosmos/feegrant/v1beta1/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// MsgGrantFeeAllowance adds permission for Grantee to spend up to Allowance
// of fees from the account of Granter.
type MsgGrantFeeAllowance struct {
	Granter   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=granter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"granter,omitempty"`
	Grantee   github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=grantee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"grantee,omitempty"`
	Allowance *types.Any                                    `protobuf:"bytes,3,opt,name=allowance,proto3" json:"allowance,omitempty"`
}

func (m *MsgGrantFeeAllowance) Reset()         { *m = MsgGrantFeeAllowance{} }
func (m *MsgGrantFeeAllowance) String() string { return proto.CompactTextString(m) }
func (*MsgGrantFeeAllowance) ProtoMessage()    {}
func (*MsgGrantFeeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{0}
}
func (m *MsgGrantFeeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantFeeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantFeeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantFeeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantFeeAllowance.Merge(m, src)
}
func (m *MsgGrantFeeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantFeeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantFeeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantFeeAllowance proto.InternalMessageInfo

// MsgGrantFeeAllowanceResponse defines the Msg/GrantFeeAllowanceResponse response type.
type MsgGrantFeeAllowanceResponse struct {
}

func (m *MsgGrantFeeAllowanceResponse) Reset()         { *m = MsgGrantFeeAllowanceResponse{} }
func (m *MsgGrantFeeAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgGrantFeeAllowanceResponse) ProtoMessage()    {}
func (*MsgGrantFeeAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{1}
}
func (m *MsgGrantFeeAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantFeeAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantFeeAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantFeeAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantFeeAllowanceResponse.Merge(m, src)
}
func (m *MsgGrantFeeAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantFeeAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantFeeAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantFeeAllowanceResponse proto.InternalMessageInfo

// MsgRevokeFeeAllowance removes any existing FeeAllowance from Granter to Grantee.
type MsgRevokeFeeAllowance struct {
	Granter github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=granter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"granter,omitempty"`
	Grantee github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=grantee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"grantee,omitempty"`
}

func (m *MsgRevokeFeeAllowance) Reset()         { *m = MsgRevokeFeeAllowance{} }
func (m *MsgRevokeFeeAllowance) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeFeeAllowance) ProtoMessage()    {}
func (*MsgRevokeFeeAllowance) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{2}
}
func (m *MsgRevokeFeeAllowance) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeFeeAllowance) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeFeeAllowance.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeFeeAllowance) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeFeeAllowance.Merge(m, src)
}
func (m *MsgRevokeFeeAllowance) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeFeeAllowance) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeFeeAllowance.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeFeeAllowance proto.InternalMessageInfo

func (m *MsgRevokeFeeAllowance) GetGranter() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Granter
	}
	return nil
}

func (m *MsgRevokeFeeAllowance) GetGrantee() github_com_cosmos_cosmos_sdk_types.AccAddress {
	if m != nil {
		return m.Grantee
	}
	return nil
}

// MsgRevokeFeeAllowanceResponse defines the Msg/RevokeFeeAllowanceResponse response type.
type MsgRevokeFeeAllowanceResponse struct {
}

func (m *MsgRevokeFeeAllowanceResponse) Reset()         { *m = MsgRevokeFeeAllowanceResponse{} }
func (m *MsgRevokeFeeAllowanceResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeFeeAllowanceResponse) ProtoMessage()    {}
func (*MsgRevokeFeeAllowanceResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd44ad7946dad783, []int{3}
}
func (m *MsgRevokeFeeAllowanceResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeFeeAllowanceResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeFeeAllowanceResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeFeeAllowanceResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeFeeAllowanceResponse.Merge(m, src)
}
func (m *MsgRevokeFeeAllowanceResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeFeeAllowanceResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeFeeAllowanceResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeFeeAllowanceResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgGrantFeeAllowance)(nil), "cosmos.feegrant.v1beta1.MsgGrantFeeAllowance")
	proto.RegisterType((*MsgGrantFeeAllowanceResponse)(nil), "cosmos.feegrant.v1beta1.MsgGrantFeeAllowanceResponse")
	proto.RegisterType((*MsgRevokeFeeAllowance)(nil), "cosmos.feegrant.v1beta1.MsgRevokeFeeAllowance")
	proto.RegisterType((*MsgRevokeFeeAllowanceResponse)(nil), "cosmos.feegrant.v1beta1.MsgRevokeFeeAllowanceResponse")
}

func init() { proto.RegisterFile("cosmos/feegrant/v1beta1/tx.proto", fileDescriptor_dd44ad7946dad783) }

var fileDescriptor_dd44ad7946dad783 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x48, 0xce, 0x2f, 0xce,
	0xcd, 0x2f, 0xd6, 0x4f, 0x4b, 0x4d, 0x4d, 0x2f, 0x4a, 0xcc, 0x2b, 0xd1, 0x2f, 0x33, 0x4c, 0x4a,
	0x2d, 0x49, 0x34, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x87, 0xa8,
	0xd0, 0x83, 0xa9, 0xd0, 0x83, 0xaa, 0x90, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xab, 0xd1, 0x07,
	0xb1, 0x20, 0xca, 0xa5, 0x24, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0xbc, 0xa4, 0xd2,
	0x34, 0xfd, 0xc4, 0xbc, 0x4a, 0x98, 0x14, 0xc4, 0xa4, 0x78, 0x88, 0x1e, 0xa8, 0xb1, 0x60, 0x8e,
	0xd2, 0x5f, 0x46, 0x2e, 0x11, 0xdf, 0xe2, 0x74, 0x77, 0x90, 0x05, 0x6e, 0xa9, 0xa9, 0x8e, 0x39,
	0x39, 0xf9, 0xe5, 0x89, 0x79, 0xc9, 0xa9, 0x42, 0xde, 0x5c, 0xec, 0x60, 0x5b, 0x53, 0x8b, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x78, 0x9c, 0x0c, 0x7f, 0xdd, 0x93, 0xd7, 0x4d, 0xcf, 0x2c, 0xc9, 0x28,
	0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0x85, 0x1a, 0x03, 0xa5, 0x74, 0x8b, 0x53, 0xb2, 0xf5, 0x4b, 0x2a,
	0x0b, 0x52, 0x8b, 0xf5, 0x1c, 0x93, 0x93, 0x1d, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x83, 0x60,
	0x26, 0x20, 0x0c, 0x4b, 0x95, 0x60, 0xa2, 0xd0, 0xb0, 0x54, 0x21, 0x57, 0x2e, 0xce, 0x44, 0x98,
	0x33, 0x25, 0x98, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0x44, 0xf4, 0x20, 0x9e, 0xd7, 0x83, 0x79, 0x5e,
	0xcf, 0x31, 0xaf, 0xd2, 0x49, 0xf0, 0xd4, 0x16, 0x5d, 0x5e, 0x64, 0x4f, 0x79, 0x06, 0x21, 0x74,
	0x5a, 0xb1, 0x74, 0x2c, 0x90, 0x67, 0x50, 0x92, 0xe3, 0x92, 0xc1, 0xe6, 0xfd, 0xa0, 0xd4, 0xe2,
	0x82, 0xfc, 0xbc, 0xe2, 0x54, 0xa5, 0x8d, 0x8c, 0x5c, 0xa2, 0xbe, 0xc5, 0xe9, 0x41, 0xa9, 0x65,
	0xf9, 0xd9, 0xa9, 0x43, 0x23, 0x80, 0x94, 0xe4, 0xb9, 0x64, 0xb1, 0x3a, 0x19, 0xe6, 0x29, 0xa3,
	0x7f, 0x8c, 0x5c, 0xcc, 0xbe, 0xc5, 0xe9, 0x42, 0x95, 0x5c, 0x82, 0x98, 0x11, 0xaf, 0xab, 0x87,
	0x23, 0xdd, 0xe9, 0x61, 0x0b, 0x28, 0x29, 0x53, 0x92, 0x94, 0xc3, 0x9c, 0x20, 0x54, 0xc3, 0x25,
	0x84, 0x25, 0x4c, 0xf5, 0xf0, 0x19, 0x86, 0xa9, 0x5e, 0xca, 0x8c, 0x34, 0xf5, 0x30, 0xdb, 0x9d,
	0xdc, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39, 0xc6, 0x09, 0x8f,
	0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x0a, 0x7f, 0x98, 0x57, 0x20,
	0xb2, 0x2b, 0x38, 0xf8, 0x93, 0xd8, 0xc0, 0x09, 0xce, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x04,
	0x77, 0xff, 0x0b, 0xce, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// GrantFeeAllowance grants fee allowance to the grantee on the granter's
	// account with the provided expiration time.
	GrantFeeAllowance(ctx context.Context, in *MsgGrantFeeAllowance, opts ...grpc.CallOption) (*MsgGrantFeeAllowanceResponse, error)
	// RevokeFeeAllowance revokes any fee allowance of granter's account that
	// has been granted to the grantee.
	RevokeFeeAllowance(ctx context.Context, in *MsgRevokeFeeAllowance, opts ...grpc.CallOption) (*MsgRevokeFeeAllowanceResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) GrantFeeAllowance(ctx context.Context, in *MsgGrantFeeAllowance, opts ...grpc.CallOption) (*MsgGrantFeeAllowanceResponse, error) {
	out := new(MsgGrantFeeAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Msg/GrantFeeAllowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) RevokeFeeAllowance(ctx context.Context, in *MsgRevokeFeeAllowance, opts ...grpc.CallOption) (*MsgRevokeFeeAllowanceResponse, error) {
	out := new(MsgRevokeFeeAllowanceResponse)
	err := c.cc.Invoke(ctx, "/cosmos.feegrant.v1beta1.Msg/RevokeFeeAllowance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// GrantFeeAllowance grants fee allowance to the grantee on the granter's
	// account with the provided expiration time.
	GrantFeeAllowance(context.Context, *MsgGrantFeeAllowance) (*MsgGrantFeeAllowanceResponse, error)
	// RevokeFeeAllowance revokes any fee allowance of granter's account that
	// has been granted to the grantee.
	RevokeFeeAllowance(context.Context, *MsgRevokeFeeAllowance) (*MsgRevokeFeeAllowanceResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) GrantFeeAllowance(ctx context.Context, req *MsgGrantFeeAllowance) (*MsgGrantFeeAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GrantFeeAllowance not implemented")
}
func (*UnimplementedMsgServer) RevokeFeeAllowance(ctx context.Context, req *MsgRevokeFeeAllowance) (*MsgRevokeFeeAllowanceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RevokeFeeAllowance not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_GrantFeeAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgGrantFeeAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).GrantFeeAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Msg/GrantFeeAllowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).GrantFeeAllowance(ctx, req.(*MsgGrantFeeAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_RevokeFeeAllowance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRevokeFeeAllowance)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RevokeFeeAllowance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cosmos.feegrant.v1beta1.Msg/RevokeFeeAllowance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RevokeFeeAllowance(ctx, req.(*MsgRevokeFeeAllowance))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cosmos.feegrant.v1beta1.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GrantFeeAllowance",
			Handler:    _Msg_GrantFeeAllowance_Handler,
		},
		{
			MethodName: "RevokeFeeAllowance",
			Handler:    _Msg_RevokeFeeAllowance_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cosmos/feegrant/v1beta1/tx.proto",
}

func (m *MsgGrantFeeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Allowance != nil {
		{
			size, err := m.Allowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgGrantFeeAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantFeeAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantFeeAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgRevokeFeeAllowance) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeFeeAllowanceResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeFeeAllowanceResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeFeeAllowanceResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgGrantFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Allowance != nil {
		l = m.Allowance.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgGrantFeeAllowanceResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgRevokeFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRevokeFeeAllowanceResponse) Size() (n int) {
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
func (m *MsgGrantFeeAllowance) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGrantFeeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantFeeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
			m.Granter = append(m.Granter[:0], dAtA[iNdEx:postIndex]...)
			if m.Granter == nil {
				m.Granter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = append(m.Grantee[:0], dAtA[iNdEx:postIndex]...)
			if m.Grantee == nil {
				m.Grantee = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Allowance", wireType)
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
			if m.Allowance == nil {
				m.Allowance = &types.Any{}
			}
			if err := m.Allowance.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgGrantFeeAllowanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgGrantFeeAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantFeeAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgRevokeFeeAllowance) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevokeFeeAllowance: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeFeeAllowance: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
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
			m.Granter = append(m.Granter[:0], dAtA[iNdEx:postIndex]...)
			if m.Granter == nil {
				m.Granter = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Grantee", wireType)
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
			m.Grantee = append(m.Grantee[:0], dAtA[iNdEx:postIndex]...)
			if m.Grantee == nil {
				m.Grantee = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgRevokeFeeAllowanceResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgRevokeFeeAllowanceResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeFeeAllowanceResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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