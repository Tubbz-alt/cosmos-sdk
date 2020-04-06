// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/msg_authorization/types/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
	_ "github.com/regen-network/cosmos-proto"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type Authorization struct {
	// sum defines a set of all acceptable concrete authorization implementations.
	//
	// Types that are valid to be assigned to Sum:
	//	*Authorization_BasicFeeAllowance
	Sum isAuthorization_Sum `protobuf_oneof:"sum"`
}

func (m *Authorization) Reset()         { *m = Authorization{} }
func (m *Authorization) String() string { return proto.CompactTextString(m) }
func (*Authorization) ProtoMessage()    {}
func (*Authorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_f556beb585238ad8, []int{0}
}
func (m *Authorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Authorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Authorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Authorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Authorization.Merge(m, src)
}
func (m *Authorization) XXX_Size() int {
	return m.Size()
}
func (m *Authorization) XXX_DiscardUnknown() {
	xxx_messageInfo_Authorization.DiscardUnknown(m)
}

var xxx_messageInfo_Authorization proto.InternalMessageInfo

type isAuthorization_Sum interface {
	isAuthorization_Sum()
	Equal(interface{}) bool
	MarshalTo([]byte) (int, error)
	Size() int
}

type Authorization_BasicFeeAllowance struct {
	BasicFeeAllowance *SendAuthorization `protobuf:"bytes,1,opt,name=basic_fee_allowance,json=basicFeeAllowance,proto3,oneof" json:"basic_fee_allowance,omitempty"`
}

func (*Authorization_BasicFeeAllowance) isAuthorization_Sum() {}

func (m *Authorization) GetSum() isAuthorization_Sum {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *Authorization) GetBasicFeeAllowance() *SendAuthorization {
	if x, ok := m.GetSum().(*Authorization_BasicFeeAllowance); ok {
		return x.BasicFeeAllowance
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Authorization) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Authorization_BasicFeeAllowance)(nil),
	}
}

type SendAuthorization struct {
	SpendLimit github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=spend_limit,json=spendLimit,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"spend_limit" yaml:"spend_limit"`
}

func (m *SendAuthorization) Reset()         { *m = SendAuthorization{} }
func (m *SendAuthorization) String() string { return proto.CompactTextString(m) }
func (*SendAuthorization) ProtoMessage()    {}
func (*SendAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_f556beb585238ad8, []int{1}
}
func (m *SendAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SendAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SendAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SendAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SendAuthorization.Merge(m, src)
}
func (m *SendAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *SendAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_SendAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_SendAuthorization proto.InternalMessageInfo

// MsgGrantAuthorization grants the provided authorization to the grantee on the granter's
// account with the provided expiration time.
type MsgGrantAuthorization struct {
	Granter       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=granter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"granter,omitempty"`
	Grantee       github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=grantee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"grantee,omitempty"`
	Authorization *Authorization                                `protobuf:"bytes,3,opt,name=authorization,proto3" json:"authorization,omitempty"`
	Time          time.Time                                     `protobuf:"bytes,4,opt,name=time,proto3,stdtime" json:"time"`
}

func (m *MsgGrantAuthorization) Reset()         { *m = MsgGrantAuthorization{} }
func (m *MsgGrantAuthorization) String() string { return proto.CompactTextString(m) }
func (*MsgGrantAuthorization) ProtoMessage()    {}
func (*MsgGrantAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_f556beb585238ad8, []int{2}
}
func (m *MsgGrantAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgGrantAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgGrantAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgGrantAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgGrantAuthorization.Merge(m, src)
}
func (m *MsgGrantAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *MsgGrantAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgGrantAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_MsgGrantAuthorization proto.InternalMessageInfo

// MsgRevokeAuthorization revokes any authorization with the provided sdk.Msg type on the
// granter's account with that has been granted to the grantee.
type MsgRevokeAuthorization struct {
	Granter              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,1,opt,name=granter,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"granter,omitempty"`
	Grantee              github_com_cosmos_cosmos_sdk_types.AccAddress `protobuf:"bytes,2,opt,name=grantee,proto3,casttype=github.com/cosmos/cosmos-sdk/types.AccAddress" json:"grantee,omitempty"`
	AuthorizationMsgType string                                        `protobuf:"bytes,3,opt,name=authorization_msg_type,json=authorizationMsgType,proto3" json:"authorization_msg_type,omitempty" yaml:"authorization_msg_type"`
}

func (m *MsgRevokeAuthorization) Reset()         { *m = MsgRevokeAuthorization{} }
func (m *MsgRevokeAuthorization) String() string { return proto.CompactTextString(m) }
func (*MsgRevokeAuthorization) ProtoMessage()    {}
func (*MsgRevokeAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_f556beb585238ad8, []int{3}
}
func (m *MsgRevokeAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRevokeAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRevokeAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRevokeAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRevokeAuthorization.Merge(m, src)
}
func (m *MsgRevokeAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *MsgRevokeAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRevokeAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRevokeAuthorization proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Authorization)(nil), "cosmos_sdk.x.msgauth.v1.Authorization")
	proto.RegisterType((*SendAuthorization)(nil), "cosmos_sdk.x.msgauth.v1.SendAuthorization")
	proto.RegisterType((*MsgGrantAuthorization)(nil), "cosmos_sdk.x.msgauth.v1.MsgGrantAuthorization")
	proto.RegisterType((*MsgRevokeAuthorization)(nil), "cosmos_sdk.x.msgauth.v1.MsgRevokeAuthorization")
}

func init() {
	proto.RegisterFile("x/msg_authorization/types/types.proto", fileDescriptor_f556beb585238ad8)
}

var fileDescriptor_f556beb585238ad8 = []byte{
	// 548 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xd4, 0x53, 0xb1, 0x6f, 0xd3, 0x4c,
	0x1c, 0xf5, 0x25, 0xf9, 0x3e, 0xca, 0x85, 0x22, 0xc5, 0x85, 0x10, 0x45, 0xc2, 0x2e, 0x96, 0xa8,
	0xa2, 0xa2, 0x9c, 0x95, 0xb2, 0xa0, 0x6c, 0x31, 0x52, 0x01, 0xb5, 0x59, 0x4c, 0x25, 0x24, 0x84,
	0x64, 0x39, 0xf6, 0xf5, 0x62, 0xc5, 0xf6, 0x59, 0xbe, 0x4b, 0x48, 0x10, 0x7f, 0x00, 0x03, 0x43,
	0x99, 0x59, 0x3a, 0x22, 0x66, 0x76, 0xd6, 0x8a, 0xa9, 0x23, 0x53, 0x8a, 0x92, 0x85, 0xb9, 0x23,
	0x13, 0x3a, 0x3b, 0x51, 0x6d, 0x85, 0x22, 0x24, 0x26, 0x16, 0xcb, 0x3e, 0xbf, 0xf7, 0xfc, 0x7e,
	0xef, 0xf7, 0x0c, 0xef, 0x8e, 0xf5, 0x80, 0x11, 0xcb, 0x1e, 0xf2, 0x3e, 0x8d, 0xbd, 0x57, 0x36,
	0xf7, 0x68, 0xa8, 0xf3, 0x49, 0x84, 0x59, 0x7a, 0x45, 0x51, 0x4c, 0x39, 0x95, 0x6f, 0x39, 0x94,
	0x05, 0x94, 0x59, 0xcc, 0x1d, 0xa0, 0x31, 0x0a, 0x18, 0x11, 0x04, 0x34, 0x6a, 0xd5, 0xef, 0xf1,
	0xbe, 0x17, 0xbb, 0x56, 0x64, 0xc7, 0x7c, 0xa2, 0x27, 0x58, 0x3d, 0x85, 0x36, 0xb3, 0x0f, 0xa9,
	0x4a, 0x7d, 0x6b, 0x15, 0x4c, 0x28, 0xa1, 0x17, 0x77, 0x0b, 0x5c, 0x65, 0xc5, 0x40, 0x5d, 0x25,
	0x94, 0x12, 0x1f, 0xa7, 0xac, 0xde, 0xf0, 0x50, 0xe7, 0x5e, 0x80, 0x19, 0xb7, 0x83, 0x28, 0x05,
	0x68, 0x6f, 0x01, 0x5c, 0xef, 0x64, 0xa7, 0x90, 0x5f, 0xc0, 0x8d, 0x9e, 0xcd, 0x3c, 0xc7, 0x3a,
	0xc4, 0xd8, 0xb2, 0x7d, 0x9f, 0xbe, 0xb4, 0x43, 0x07, 0xd7, 0xc0, 0x26, 0x68, 0x94, 0x77, 0xb6,
	0xd1, 0x25, 0x13, 0xa1, 0xa7, 0x38, 0x74, 0x73, 0x42, 0x8f, 0x25, 0xb3, 0x92, 0x08, 0xed, 0x62,
	0xdc, 0x59, 0xca, 0xb4, 0xab, 0xdf, 0x8f, 0x55, 0xf0, 0xe5, 0x53, 0xf3, 0xfa, 0x76, 0xf6, 0xf8,
	0x89, 0xf1, 0x1f, 0x2c, 0xb2, 0x61, 0xa0, 0xbd, 0x07, 0xb0, 0xb2, 0xa2, 0x24, 0xbf, 0x86, 0x65,
	0x16, 0xe1, 0xd0, 0xb5, 0x7c, 0x2f, 0xf0, 0x78, 0x0d, 0x6c, 0x16, 0x1b, 0xe5, 0x9d, 0x8d, 0xac,
	0x95, 0x51, 0x0b, 0x3d, 0xa4, 0x5e, 0x68, 0xec, 0x9e, 0x4c, 0x55, 0xe9, 0x7c, 0xaa, 0xca, 0x13,
	0x3b, 0xf0, 0xdb, 0x5a, 0x86, 0xa5, 0x7d, 0x3c, 0x53, 0x1b, 0xc4, 0xe3, 0xfd, 0x61, 0x0f, 0x39,
	0x34, 0x58, 0x24, 0xbc, 0x4c, 0x9d, 0xb9, 0x83, 0x45, 0x6e, 0x42, 0x86, 0x99, 0x30, 0x61, 0xee,
	0x0b, 0x62, 0x7b, 0xed, 0xcd, 0xb1, 0x2a, 0x09, 0xdb, 0xda, 0xe7, 0x02, 0xbc, 0xd9, 0x65, 0xe4,
	0x51, 0x6c, 0x87, 0x3c, 0xef, 0x70, 0x0f, 0x5e, 0x21, 0xe2, 0x14, 0xc7, 0x49, 0x50, 0xd7, 0x8c,
	0xd6, 0x8f, 0xa9, 0xda, 0xfc, 0x83, 0xcf, 0x75, 0x1c, 0xa7, 0xe3, 0xba, 0x31, 0x66, 0xcc, 0x5c,
	0x2a, 0x5c, 0x88, 0xe1, 0x5a, 0xe1, 0x2f, 0xc5, 0xb0, 0xbc, 0x0f, 0xd7, 0x73, 0x2d, 0xad, 0x15,
	0x93, 0x45, 0x6e, 0x5d, 0xba, 0xc8, 0xdc, 0x60, 0x66, 0x9e, 0x2c, 0x3f, 0x80, 0x25, 0xd1, 0xa0,
	0x5a, 0x29, 0x11, 0xa9, 0xa3, 0xb4, 0x5e, 0x68, 0x59, 0x2f, 0x74, 0xb0, 0xac, 0x97, 0xb1, 0x26,
	0x36, 0x71, 0x74, 0xa6, 0x02, 0x33, 0x61, 0xb4, 0x4b, 0x22, 0x45, 0xed, 0x5d, 0x01, 0x56, 0xbb,
	0x8c, 0x98, 0x78, 0x44, 0x07, 0xf8, 0x5f, 0x89, 0xf0, 0x19, 0xac, 0xe6, 0x52, 0xb0, 0xc4, 0xaf,
	0x2f, 0xe0, 0x49, 0x96, 0x57, 0x8d, 0x3b, 0xe7, 0x53, 0xf5, 0x76, 0x5a, 0xb8, 0x5f, 0xe3, 0x34,
	0xf3, 0x46, 0xee, 0x45, 0x97, 0x91, 0x83, 0x49, 0xb4, 0xc8, 0xc4, 0xd8, 0xfb, 0x30, 0x53, 0xc0,
	0xc9, 0x4c, 0x01, 0xa7, 0x33, 0x05, 0x7c, 0x9b, 0x29, 0xe0, 0x68, 0xae, 0x48, 0xa7, 0x73, 0x45,
	0xfa, 0x3a, 0x57, 0xa4, 0xe7, 0xbf, 0x37, 0x3d, 0xd6, 0x99, 0x6f, 0xb3, 0xbe, 0x17, 0x92, 0xd4,
	0x7f, 0xef, 0xff, 0x64, 0x15, 0xf7, 0x7f, 0x06, 0x00, 0x00, 0xff, 0xff, 0x63, 0x70, 0xf4, 0x6f,
	0xa1, 0x04, 0x00, 0x00,
}

func (this *Authorization) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Authorization)
	if !ok {
		that2, ok := that.(Authorization)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if that1.Sum == nil {
		if this.Sum != nil {
			return false
		}
	} else if this.Sum == nil {
		return false
	} else if !this.Sum.Equal(that1.Sum) {
		return false
	}
	return true
}
func (this *Authorization_BasicFeeAllowance) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Authorization_BasicFeeAllowance)
	if !ok {
		that2, ok := that.(Authorization_BasicFeeAllowance)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.BasicFeeAllowance.Equal(that1.BasicFeeAllowance) {
		return false
	}
	return true
}
func (this *SendAuthorization) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SendAuthorization)
	if !ok {
		that2, ok := that.(SendAuthorization)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.SpendLimit) != len(that1.SpendLimit) {
		return false
	}
	for i := range this.SpendLimit {
		if !this.SpendLimit[i].Equal(&that1.SpendLimit[i]) {
			return false
		}
	}
	return true
}
func (this *MsgGrantAuthorization) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgGrantAuthorization)
	if !ok {
		that2, ok := that.(MsgGrantAuthorization)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Granter, that1.Granter) {
		return false
	}
	if !bytes.Equal(this.Grantee, that1.Grantee) {
		return false
	}
	if !this.Authorization.Equal(that1.Authorization) {
		return false
	}
	if !this.Time.Equal(that1.Time) {
		return false
	}
	return true
}
func (this *MsgRevokeAuthorization) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*MsgRevokeAuthorization)
	if !ok {
		that2, ok := that.(MsgRevokeAuthorization)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !bytes.Equal(this.Granter, that1.Granter) {
		return false
	}
	if !bytes.Equal(this.Grantee, that1.Grantee) {
		return false
	}
	if this.AuthorizationMsgType != that1.AuthorizationMsgType {
		return false
	}
	return true
}
func (this *Authorization) GetFeeAllowanceI() FeeAllowanceI {
	if x := this.GetBasicFeeAllowance(); x != nil {
		return x
	}
	return nil
}

func (this *Authorization) SetFeeAllowanceI(value FeeAllowanceI) error {
	if value == nil {
		this.Sum = nil
		return nil
	}
	switch vt := value.(type) {
	case *SendAuthorization:
		this.Sum = &Authorization_BasicFeeAllowance{vt}
		return nil
	}
	return fmt.Errorf("can't encode value of type %T as message Authorization", value)
}

func (m *Authorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Authorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Authorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Sum != nil {
		{
			size := m.Sum.Size()
			i -= size
			if _, err := m.Sum.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
		}
	}
	return len(dAtA) - i, nil
}

func (m *Authorization_BasicFeeAllowance) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Authorization_BasicFeeAllowance) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	if m.BasicFeeAllowance != nil {
		{
			size, err := m.BasicFeeAllowance.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}
func (m *SendAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SendAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SendAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for iNdEx := len(m.SpendLimit) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SpendLimit[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTypes(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *MsgGrantAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgGrantAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgGrantAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.Time, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.Time):])
	if err2 != nil {
		return 0, err2
	}
	i -= n2
	i = encodeVarintTypes(dAtA, i, uint64(n2))
	i--
	dAtA[i] = 0x22
	if m.Authorization != nil {
		{
			size, err := m.Authorization.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRevokeAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRevokeAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRevokeAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthorizationMsgType) > 0 {
		i -= len(m.AuthorizationMsgType)
		copy(dAtA[i:], m.AuthorizationMsgType)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.AuthorizationMsgType)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Grantee) > 0 {
		i -= len(m.Grantee)
		copy(dAtA[i:], m.Grantee)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Grantee)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Granter) > 0 {
		i -= len(m.Granter)
		copy(dAtA[i:], m.Granter)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Granter)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Authorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Sum != nil {
		n += m.Sum.Size()
	}
	return n
}

func (m *Authorization_BasicFeeAllowance) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BasicFeeAllowance != nil {
		l = m.BasicFeeAllowance.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}
func (m *SendAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.SpendLimit) > 0 {
		for _, e := range m.SpendLimit {
			l = e.Size()
			n += 1 + l + sovTypes(uint64(l))
		}
	}
	return n
}

func (m *MsgGrantAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.Authorization != nil {
		l = m.Authorization.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.Time)
	n += 1 + l + sovTypes(uint64(l))
	return n
}

func (m *MsgRevokeAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Granter)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.Grantee)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	l = len(m.AuthorizationMsgType)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Authorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: Authorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Authorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BasicFeeAllowance", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			v := &SendAuthorization{}
			if err := v.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			m.Sum = &Authorization_BasicFeeAllowance{v}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *SendAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: SendAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SendAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SpendLimit", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SpendLimit = append(m.SpendLimit, types.Coin{})
			if err := m.SpendLimit[len(m.SpendLimit)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *MsgGrantAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MsgGrantAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgGrantAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
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
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
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
				return fmt.Errorf("proto: wrong wireType = %d for field Authorization", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Authorization == nil {
				m.Authorization = &Authorization{}
			}
			if err := m.Authorization.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.Time, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *MsgRevokeAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: MsgRevokeAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRevokeAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Granter", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
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
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
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
				return fmt.Errorf("proto: wrong wireType = %d for field AuthorizationMsgType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthorizationMsgType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTypes
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
