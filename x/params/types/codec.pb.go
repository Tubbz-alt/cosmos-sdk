// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: x/params/types/codec.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// ParameterChangeProposal defines a proposal which contains multiple parameter
// changes.
type ParameterChangeProposal struct {
	Title       string         `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Description string         `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	Changes     []*ParamChange `protobuf:"bytes,3,rep,name=changes,proto3" json:"changes,omitempty"`
}

func (m *ParameterChangeProposal) Reset()         { *m = ParameterChangeProposal{} }
func (m *ParameterChangeProposal) String() string { return proto.CompactTextString(m) }
func (*ParameterChangeProposal) ProtoMessage()    {}
func (*ParameterChangeProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ef361ccb69097c, []int{0}
}
func (m *ParameterChangeProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParameterChangeProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParameterChangeProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParameterChangeProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParameterChangeProposal.Merge(m, src)
}
func (m *ParameterChangeProposal) XXX_Size() int {
	return m.Size()
}
func (m *ParameterChangeProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_ParameterChangeProposal.DiscardUnknown(m)
}

var xxx_messageInfo_ParameterChangeProposal proto.InternalMessageInfo

func (m *ParameterChangeProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *ParameterChangeProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *ParameterChangeProposal) GetChanges() []*ParamChange {
	if m != nil {
		return m.Changes
	}
	return nil
}

// ParamChange defines a parameter change.
type ParamChange struct {
	Subspace string `protobuf:"bytes,1,opt,name=subspace,proto3" json:"subspace,omitempty"`
	Key      string `protobuf:"bytes,2,opt,name=key,proto3" json:"key,omitempty"`
	Value    string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *ParamChange) Reset()      { *m = ParamChange{} }
func (*ParamChange) ProtoMessage() {}
func (*ParamChange) Descriptor() ([]byte, []int) {
	return fileDescriptor_30ef361ccb69097c, []int{1}
}
func (m *ParamChange) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ParamChange) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ParamChange.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ParamChange) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParamChange.Merge(m, src)
}
func (m *ParamChange) XXX_Size() int {
	return m.Size()
}
func (m *ParamChange) XXX_DiscardUnknown() {
	xxx_messageInfo_ParamChange.DiscardUnknown(m)
}

var xxx_messageInfo_ParamChange proto.InternalMessageInfo

func (m *ParamChange) GetSubspace() string {
	if m != nil {
		return m.Subspace
	}
	return ""
}

func (m *ParamChange) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *ParamChange) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*ParameterChangeProposal)(nil), "cosmos_sdk.x.params.v1.ParameterChangeProposal")
	proto.RegisterType((*ParamChange)(nil), "cosmos_sdk.x.params.v1.ParamChange")
}

func init() { proto.RegisterFile("x/params/types/codec.proto", fileDescriptor_30ef361ccb69097c) }

var fileDescriptor_30ef361ccb69097c = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xbb, 0x4e, 0xf3, 0x30,
	0x14, 0xc7, 0xe3, 0x2f, 0x1f, 0x37, 0x77, 0x01, 0x0b, 0x41, 0x94, 0xc1, 0x54, 0x65, 0xa9, 0x84,
	0xb0, 0xc5, 0x65, 0x42, 0x62, 0x01, 0xb1, 0x57, 0x5d, 0x90, 0x58, 0x90, 0xe3, 0x98, 0x34, 0xea,
	0xe5, 0x58, 0xb6, 0x53, 0xb5, 0x6f, 0xd1, 0x91, 0x91, 0xc7, 0x61, 0xec, 0xc8, 0x88, 0xda, 0x17,
	0x41, 0xb5, 0x5b, 0x14, 0x04, 0x93, 0xcf, 0xff, 0x5c, 0x7f, 0x3e, 0x07, 0xa7, 0x13, 0xae, 0x85,
	0x11, 0x43, 0xcb, 0xdd, 0x54, 0x2b, 0xcb, 0x25, 0xe4, 0x4a, 0x32, 0x6d, 0xc0, 0x01, 0x39, 0x92,
	0x60, 0x87, 0x60, 0x9f, 0x6d, 0xde, 0x67, 0x13, 0x16, 0xd2, 0xd8, 0xf8, 0x22, 0x3d, 0xf8, 0x95,
	0x9a, 0x5e, 0x8f, 0xd5, 0x28, 0x07, 0xc3, 0x8b, 0xd2, 0xf5, 0xaa, 0x8c, 0x49, 0x18, 0xf2, 0x02,
	0x0a, 0xe0, 0x3e, 0x9a, 0x55, 0x2f, 0x5e, 0x79, 0xe1, 0xad, 0x50, 0xd5, 0x9a, 0x21, 0x7c, 0xdc,
	0x59, 0xb5, 0x55, 0x4e, 0x99, 0xfb, 0x9e, 0x18, 0x15, 0xaa, 0x63, 0x40, 0x83, 0x15, 0x03, 0x72,
	0x88, 0xb7, 0x5c, 0xe9, 0x06, 0x2a, 0x41, 0x4d, 0xd4, 0xde, 0xeb, 0x06, 0x41, 0x9a, 0xb8, 0x91,
	0x2b, 0x2b, 0x4d, 0xa9, 0x5d, 0x09, 0xa3, 0xe4, 0x9f, 0x8f, 0xd5, 0x5d, 0xe4, 0x16, 0xef, 0x48,
	0xdf, 0xc9, 0x26, 0x71, 0x33, 0x6e, 0x37, 0x2e, 0x4f, 0xd9, 0xdf, 0xdf, 0x60, 0x7e, 0x72, 0x98,
	0xda, 0xdd, 0xd4, 0xb4, 0x1e, 0x71, 0xa3, 0xe6, 0x27, 0x29, 0xde, 0xb5, 0x55, 0x66, 0xb5, 0x90,
	0x1b, 0x90, 0x6f, 0x4d, 0xf6, 0x71, 0xdc, 0x57, 0xd3, 0x35, 0xc3, 0xca, 0x5c, 0x31, 0x8f, 0xc5,
	0xa0, 0x52, 0x49, 0x1c, 0x98, 0xbd, 0xb8, 0xf9, 0xff, 0xfa, 0x76, 0x12, 0xdd, 0x3d, 0xbc, 0x2f,
	0x28, 0x9a, 0x2f, 0x28, 0xfa, 0x5c, 0x50, 0x34, 0x5b, 0xd2, 0x68, 0xbe, 0xa4, 0xd1, 0xc7, 0x92,
	0x46, 0x4f, 0x67, 0xb5, 0xa5, 0x05, 0xd4, 0xf5, 0x73, 0x6e, 0xf3, 0x3e, 0xff, 0x79, 0x9f, 0x6c,
	0xdb, 0x6f, 0xee, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x83, 0x93, 0x48, 0x72, 0xb8, 0x01, 0x00,
	0x00,
}

func (m *ParameterChangeProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParameterChangeProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParameterChangeProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Changes) > 0 {
		for iNdEx := len(m.Changes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Changes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCodec(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ParamChange) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ParamChange) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ParamChange) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Key) > 0 {
		i -= len(m.Key)
		copy(dAtA[i:], m.Key)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Key)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Subspace) > 0 {
		i -= len(m.Subspace)
		copy(dAtA[i:], m.Subspace)
		i = encodeVarintCodec(dAtA, i, uint64(len(m.Subspace)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCodec(dAtA []byte, offset int, v uint64) int {
	offset -= sovCodec(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ParameterChangeProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	if len(m.Changes) > 0 {
		for _, e := range m.Changes {
			l = e.Size()
			n += 1 + l + sovCodec(uint64(l))
		}
	}
	return n
}

func (m *ParamChange) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Subspace)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Key)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovCodec(uint64(l))
	}
	return n
}

func sovCodec(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCodec(x uint64) (n int) {
	return sovCodec(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ParameterChangeProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: ParameterChangeProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParameterChangeProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Changes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Changes = append(m.Changes, &ParamChange{})
			if err := m.Changes[len(m.Changes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func (m *ParamChange) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCodec
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
			return fmt.Errorf("proto: ParamChange: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ParamChange: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subspace", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subspace = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Key", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Key = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCodec
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
				return ErrInvalidLengthCodec
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCodec
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCodec(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthCodec
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthCodec
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
func skipCodec(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
					return 0, ErrIntOverflowCodec
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
				return 0, ErrInvalidLengthCodec
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCodec
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCodec
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCodec        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCodec          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCodec = fmt.Errorf("proto: unexpected end of group")
)
