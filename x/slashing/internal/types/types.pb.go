// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/slashing/internal/types/types.proto

package types

import (
	bytes "bytes"
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "github.com/golang/protobuf/ptypes/timestamp"
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

// MsgUnjail - struct for unjailing jailed validator
type MsgUnjail struct {
	ValidatorAddr github_com_cosmos_cosmos_sdk_types.ValAddress `protobuf:"bytes,1,opt,name=validator_addr,json=validatorAddr,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ValAddress" json:"validator_addr,omitempty" yaml:"address"`
}

func (m *MsgUnjail) Reset()         { *m = MsgUnjail{} }
func (m *MsgUnjail) String() string { return proto.CompactTextString(m) }
func (*MsgUnjail) ProtoMessage()    {}
func (*MsgUnjail) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{0}
}
func (m *MsgUnjail) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUnjail) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUnjail.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUnjail) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUnjail.Merge(m, src)
}
func (m *MsgUnjail) XXX_Size() int {
	return m.Size()
}
func (m *MsgUnjail) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUnjail.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUnjail proto.InternalMessageInfo

func (m *MsgUnjail) GetValidatorAddr() github_com_cosmos_cosmos_sdk_types.ValAddress {
	if m != nil {
		return m.ValidatorAddr
	}
	return nil
}

// ValidatorSigningInfo defines the signing info for a validator
type ValidatorSigningInfo struct {
	Address             github_com_cosmos_cosmos_sdk_types.ConsAddress `protobuf:"bytes,1,opt,name=address,proto3,casttype=github.com/cosmos/cosmos-sdk/types.ConsAddress" json:"address,omitempty"`
	StartHeight         int64                                          `protobuf:"varint,2,opt,name=start_height,json=startHeight,proto3" json:"start_height,omitempty" yaml:"start_height"`
	IndexOffset         int64                                          `protobuf:"varint,3,opt,name=index_offset,json=indexOffset,proto3" json:"index_offset,omitempty" yaml:"index_offset"`
	JailedUntil         time.Time                                      `protobuf:"bytes,4,opt,name=jailed_until,json=jailedUntil,proto3,stdtime" json:"jailed_until" yaml:"jailed_until"`
	Tombstoned          bool                                           `protobuf:"varint,5,opt,name=tombstoned,proto3" json:"tombstoned,omitempty"`
	MissedBlocksCounter int64                                          `protobuf:"varint,6,opt,name=missed_blocks_counter,json=missedBlocksCounter,proto3" json:"missed_blocks_counter,omitempty" yaml:"missed_blocks_counter"`
}

func (m *ValidatorSigningInfo) Reset()      { *m = ValidatorSigningInfo{} }
func (*ValidatorSigningInfo) ProtoMessage() {}
func (*ValidatorSigningInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2b882c3b0cdd6f57, []int{1}
}
func (m *ValidatorSigningInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorSigningInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorSigningInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorSigningInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorSigningInfo.Merge(m, src)
}
func (m *ValidatorSigningInfo) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorSigningInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorSigningInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorSigningInfo proto.InternalMessageInfo

func (m *ValidatorSigningInfo) GetAddress() github_com_cosmos_cosmos_sdk_types.ConsAddress {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ValidatorSigningInfo) GetStartHeight() int64 {
	if m != nil {
		return m.StartHeight
	}
	return 0
}

func (m *ValidatorSigningInfo) GetIndexOffset() int64 {
	if m != nil {
		return m.IndexOffset
	}
	return 0
}

func (m *ValidatorSigningInfo) GetJailedUntil() time.Time {
	if m != nil {
		return m.JailedUntil
	}
	return time.Time{}
}

func (m *ValidatorSigningInfo) GetTombstoned() bool {
	if m != nil {
		return m.Tombstoned
	}
	return false
}

func (m *ValidatorSigningInfo) GetMissedBlocksCounter() int64 {
	if m != nil {
		return m.MissedBlocksCounter
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgUnjail)(nil), "cosmos_sdk.x.slashing.v1.MsgUnjail")
	proto.RegisterType((*ValidatorSigningInfo)(nil), "cosmos_sdk.x.slashing.v1.ValidatorSigningInfo")
}

func init() {
	proto.RegisterFile("x/slashing/internal/types/types.proto", fileDescriptor_2b882c3b0cdd6f57)
}

var fileDescriptor_2b882c3b0cdd6f57 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x94, 0x96, 0x72, 0x09, 0x1d, 0x5c, 0x10, 0x56, 0x84, 0x7c, 0x96, 0x25, 0x50,
	0x96, 0xda, 0xa2, 0x6c, 0xd9, 0x70, 0x16, 0x90, 0x40, 0x48, 0xa6, 0xed, 0xc0, 0x80, 0x75, 0xce,
	0x5d, 0xce, 0xd7, 0xd8, 0x77, 0x91, 0xef, 0x52, 0x92, 0x6f, 0xd1, 0x91, 0xb1, 0x1f, 0x84, 0x0f,
	0xd0, 0xb1, 0x23, 0x93, 0x41, 0xc9, 0xc2, 0x9c, 0xb1, 0x13, 0xb2, 0x2f, 0xa6, 0x11, 0x62, 0x60,
	0xb1, 0xfd, 0x7e, 0xf7, 0xfe, 0xef, 0xef, 0xe7, 0xbf, 0xe1, 0xf3, 0x79, 0xa8, 0x72, 0xac, 0x32,
	0x2e, 0x58, 0xc8, 0x85, 0xa6, 0xa5, 0xc0, 0x79, 0xa8, 0x17, 0x53, 0xaa, 0xcc, 0x35, 0x98, 0x96,
	0x52, 0x4b, 0xdb, 0x19, 0x49, 0x55, 0x48, 0x95, 0x28, 0x32, 0x09, 0xe6, 0x41, 0xab, 0x08, 0x2e,
	0x5e, 0xf6, 0x5e, 0xe8, 0x8c, 0x97, 0x24, 0x99, 0xe2, 0x52, 0x2f, 0xc2, 0xa6, 0x39, 0x64, 0x92,
	0xc9, 0xbb, 0x27, 0x33, 0xa1, 0x87, 0x98, 0x94, 0x2c, 0xa7, 0xa6, 0x25, 0x9d, 0x8d, 0x43, 0xcd,
	0x0b, 0xaa, 0x34, 0x2e, 0xa6, 0xa6, 0xc1, 0xff, 0x02, 0x1f, 0xbe, 0x57, 0xec, 0x54, 0x9c, 0x63,
	0x9e, 0xdb, 0xe7, 0xf0, 0xe0, 0x02, 0xe7, 0x9c, 0x60, 0x2d, 0xcb, 0x04, 0x13, 0x52, 0x3a, 0xc0,
	0x03, 0xfd, 0x6e, 0x34, 0x5c, 0x57, 0xe8, 0x60, 0x81, 0x8b, 0x7c, 0xe0, 0xd7, 0x94, 0x2a, 0xe5,
	0xdf, 0x56, 0xe8, 0x88, 0x71, 0x9d, 0xcd, 0xd2, 0x60, 0x24, 0x8b, 0xd0, 0xbc, 0xe8, 0xe6, 0x76,
	0xa4, 0xc8, 0x64, 0xb3, 0xc7, 0x19, 0xce, 0x5f, 0x1b, 0x45, 0xfc, 0xe8, 0xcf, 0xe8, 0x9a, 0xf8,
	0xdf, 0x76, 0xe0, 0xe3, 0xb3, 0x96, 0x7c, 0xe4, 0x4c, 0x70, 0xc1, 0xde, 0x8a, 0xb1, 0xb4, 0xdf,
	0xc1, 0x07, 0x1b, 0x93, 0x8d, 0xfb, 0xf1, 0x6d, 0x85, 0x82, 0xff, 0xf0, 0x1a, 0x4a, 0xa1, 0x5a,
	0xb3, 0x76, 0x84, 0x3d, 0x80, 0x5d, 0xa5, 0x71, 0xa9, 0x93, 0x8c, 0x72, 0x96, 0x69, 0xe7, 0x9e,
	0x07, 0xfa, 0x3b, 0xd1, 0xd3, 0x75, 0x85, 0x0e, 0xcd, 0x42, 0xdb, 0xa7, 0x7e, 0xdc, 0x69, 0xca,
	0x37, 0x4d, 0x55, 0x6b, 0xb9, 0x20, 0x74, 0x9e, 0xc8, 0xf1, 0x58, 0x51, 0xed, 0xec, 0xfc, 0xad,
	0xdd, 0x3e, 0xf5, 0xe3, 0x4e, 0x53, 0x7e, 0x68, 0x2a, 0xfb, 0x33, 0xec, 0xd6, 0x9f, 0x94, 0x92,
	0x64, 0x26, 0x34, 0xcf, 0x9d, 0xfb, 0x1e, 0xe8, 0x77, 0x8e, 0x7b, 0x81, 0xc9, 0x23, 0x68, 0xf3,
	0x08, 0x4e, 0xda, 0x3c, 0x22, 0x74, 0x5d, 0x21, 0xeb, 0x6e, 0xf6, 0xb6, 0xda, 0xbf, 0xfc, 0x81,
	0x40, 0xdc, 0x31, 0xe8, 0xb4, 0x26, 0xb6, 0x0b, 0xa1, 0x96, 0x45, 0xaa, 0xb4, 0x14, 0x94, 0x38,
	0xbb, 0x1e, 0xe8, 0xef, 0xc7, 0x5b, 0xc4, 0x3e, 0x81, 0x4f, 0x0a, 0xae, 0x14, 0x25, 0x49, 0x9a,
	0xcb, 0xd1, 0x44, 0x25, 0x23, 0x39, 0xab, 0x7f, 0x34, 0x67, 0xaf, 0x59, 0xc2, 0x5b, 0x57, 0xe8,
	0x99, 0x31, 0xfa, 0x67, 0x9b, 0x1f, 0x1f, 0x1a, 0x1e, 0x35, 0x78, 0x68, 0xe8, 0x60, 0xff, 0xeb,
	0x15, 0xb2, 0x7e, 0x5d, 0x21, 0x10, 0xa1, 0xeb, 0xa5, 0x0b, 0x6e, 0x96, 0x2e, 0xf8, 0xb9, 0x74,
	0xc1, 0xe5, 0xca, 0xb5, 0x6e, 0x56, 0xae, 0xf5, 0x7d, 0xe5, 0x5a, 0x9f, 0x76, 0x9b, 0x34, 0xd2,
	0xbd, 0x66, 0xc5, 0x57, 0xbf, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbc, 0x6f, 0xa4, 0x75, 0xeb, 0x02,
	0x00, 0x00,
}

func (this *ValidatorSigningInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*ValidatorSigningInfo)
	if !ok {
		that2, ok := that.(ValidatorSigningInfo)
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
	if !bytes.Equal(this.Address, that1.Address) {
		return false
	}
	if this.StartHeight != that1.StartHeight {
		return false
	}
	if this.IndexOffset != that1.IndexOffset {
		return false
	}
	if !this.JailedUntil.Equal(that1.JailedUntil) {
		return false
	}
	if this.Tombstoned != that1.Tombstoned {
		return false
	}
	if this.MissedBlocksCounter != that1.MissedBlocksCounter {
		return false
	}
	return true
}
func (m *MsgUnjail) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUnjail) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUnjail) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ValidatorAddr) > 0 {
		i -= len(m.ValidatorAddr)
		copy(dAtA[i:], m.ValidatorAddr)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.ValidatorAddr)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorSigningInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorSigningInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorSigningInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.MissedBlocksCounter != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.MissedBlocksCounter))
		i--
		dAtA[i] = 0x30
	}
	if m.Tombstoned {
		i--
		if m.Tombstoned {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x28
	}
	n1, err1 := github_com_gogo_protobuf_types.StdTimeMarshalTo(m.JailedUntil, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintTypes(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	if m.IndexOffset != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.IndexOffset))
		i--
		dAtA[i] = 0x18
	}
	if m.StartHeight != 0 {
		i = encodeVarintTypes(dAtA, i, uint64(m.StartHeight))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Address)))
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
func (m *MsgUnjail) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ValidatorAddr)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ValidatorSigningInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.StartHeight != 0 {
		n += 1 + sovTypes(uint64(m.StartHeight))
	}
	if m.IndexOffset != 0 {
		n += 1 + sovTypes(uint64(m.IndexOffset))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdTime(m.JailedUntil)
	n += 1 + l + sovTypes(uint64(l))
	if m.Tombstoned {
		n += 2
	}
	if m.MissedBlocksCounter != 0 {
		n += 1 + sovTypes(uint64(m.MissedBlocksCounter))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUnjail) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgUnjail: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUnjail: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ValidatorAddr", wireType)
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
			m.ValidatorAddr = append(m.ValidatorAddr[:0], dAtA[iNdEx:postIndex]...)
			if m.ValidatorAddr == nil {
				m.ValidatorAddr = []byte{}
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
func (m *ValidatorSigningInfo) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: ValidatorSigningInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorSigningInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = append(m.Address[:0], dAtA[iNdEx:postIndex]...)
			if m.Address == nil {
				m.Address = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartHeight", wireType)
			}
			m.StartHeight = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.StartHeight |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IndexOffset", wireType)
			}
			m.IndexOffset = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IndexOffset |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JailedUntil", wireType)
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
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(&m.JailedUntil, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tombstoned", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Tombstoned = bool(v != 0)
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field MissedBlocksCounter", wireType)
			}
			m.MissedBlocksCounter = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.MissedBlocksCounter |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
