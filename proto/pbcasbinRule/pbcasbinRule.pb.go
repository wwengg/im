// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pbcasbinRule/pbcasbinRule.proto

package pbcasbinRule

import (
	fmt "fmt"
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

type CheckCasbinRuleArgs struct {
	Role                 int64    `protobuf:"varint,1,opt,name=role,proto3" json:"role,omitempty"`
	Path                 string   `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckCasbinRuleArgs) Reset()         { *m = CheckCasbinRuleArgs{} }
func (m *CheckCasbinRuleArgs) String() string { return proto.CompactTextString(m) }
func (*CheckCasbinRuleArgs) ProtoMessage()    {}
func (*CheckCasbinRuleArgs) Descriptor() ([]byte, []int) {
	return fileDescriptor_97ba891febc1ec3a, []int{0}
}
func (m *CheckCasbinRuleArgs) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckCasbinRuleArgs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckCasbinRuleArgs.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckCasbinRuleArgs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckCasbinRuleArgs.Merge(m, src)
}
func (m *CheckCasbinRuleArgs) XXX_Size() int {
	return m.Size()
}
func (m *CheckCasbinRuleArgs) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckCasbinRuleArgs.DiscardUnknown(m)
}

var xxx_messageInfo_CheckCasbinRuleArgs proto.InternalMessageInfo

func (m *CheckCasbinRuleArgs) GetRole() int64 {
	if m != nil {
		return m.Role
	}
	return 0
}

func (m *CheckCasbinRuleArgs) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *CheckCasbinRuleArgs) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type CheckCasbinRuleReply struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CheckCasbinRuleReply) Reset()         { *m = CheckCasbinRuleReply{} }
func (m *CheckCasbinRuleReply) String() string { return proto.CompactTextString(m) }
func (*CheckCasbinRuleReply) ProtoMessage()    {}
func (*CheckCasbinRuleReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_97ba891febc1ec3a, []int{1}
}
func (m *CheckCasbinRuleReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CheckCasbinRuleReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CheckCasbinRuleReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CheckCasbinRuleReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckCasbinRuleReply.Merge(m, src)
}
func (m *CheckCasbinRuleReply) XXX_Size() int {
	return m.Size()
}
func (m *CheckCasbinRuleReply) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckCasbinRuleReply.DiscardUnknown(m)
}

var xxx_messageInfo_CheckCasbinRuleReply proto.InternalMessageInfo

func (m *CheckCasbinRuleReply) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

func init() {
	proto.RegisterType((*CheckCasbinRuleArgs)(nil), "pbcasbinRule.CheckCasbinRuleArgs")
	proto.RegisterType((*CheckCasbinRuleReply)(nil), "pbcasbinRule.CheckCasbinRuleReply")
}

func init() { proto.RegisterFile("pbcasbinRule/pbcasbinRule.proto", fileDescriptor_97ba891febc1ec3a) }

var fileDescriptor_97ba891febc1ec3a = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x48, 0x4a, 0x4e,
	0x2c, 0x4e, 0xca, 0xcc, 0x0b, 0x2a, 0xcd, 0x49, 0xd5, 0x47, 0xe6, 0xe8, 0x15, 0x14, 0xe5, 0x97,
	0xe4, 0x0b, 0xf1, 0x20, 0x8b, 0x29, 0x85, 0x72, 0x09, 0x3b, 0x67, 0xa4, 0x26, 0x67, 0x3b, 0xc3,
	0x85, 0x1c, 0x8b, 0xd2, 0x8b, 0x85, 0x84, 0xb8, 0x58, 0x8a, 0xf2, 0x73, 0x52, 0x25, 0x18, 0x15,
	0x18, 0x35, 0x98, 0x83, 0xc0, 0x6c, 0x90, 0x58, 0x41, 0x62, 0x49, 0x86, 0x04, 0x93, 0x02, 0xa3,
	0x06, 0x67, 0x10, 0x98, 0x2d, 0x24, 0xc6, 0xc5, 0x96, 0x9b, 0x5a, 0x92, 0x91, 0x9f, 0x22, 0xc1,
	0x0c, 0x16, 0x85, 0xf2, 0x94, 0xd4, 0xb8, 0x44, 0xd0, 0x8c, 0x0d, 0x4a, 0x2d, 0xc8, 0xa9, 0x14,
	0xe2, 0xe3, 0x62, 0xca, 0xcf, 0x06, 0x9b, 0xca, 0x11, 0xc4, 0x94, 0x9f, 0x6d, 0x94, 0xc1, 0xc5,
	0x85, 0x50, 0x22, 0x14, 0xc5, 0xc5, 0x8f, 0xa6, 0x4b, 0x48, 0x51, 0x0f, 0xc5, 0x0b, 0x58, 0xdc,
	0x2a, 0xa5, 0x84, 0x57, 0x09, 0xd8, 0x5e, 0x25, 0x06, 0x27, 0xd5, 0x13, 0x8f, 0xe4, 0x18, 0x2f,
	0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x31, 0x4a, 0xbc, 0xb4, 0x20, 0xb7, 0x58, 0x1f, 0x1c,
	0x20, 0x28, 0x61, 0x94, 0xc4, 0x06, 0x16, 0x33, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x52,
	0xc4, 0x44, 0x47, 0x01, 0x00, 0x00,
}

func (m *CheckCasbinRuleArgs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckCasbinRuleArgs) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckCasbinRuleArgs) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Method) > 0 {
		i -= len(m.Method)
		copy(dAtA[i:], m.Method)
		i = encodeVarintPbcasbinRule(dAtA, i, uint64(len(m.Method)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Path) > 0 {
		i -= len(m.Path)
		copy(dAtA[i:], m.Path)
		i = encodeVarintPbcasbinRule(dAtA, i, uint64(len(m.Path)))
		i--
		dAtA[i] = 0x12
	}
	if m.Role != 0 {
		i = encodeVarintPbcasbinRule(dAtA, i, uint64(m.Role))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *CheckCasbinRuleReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CheckCasbinRuleReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CheckCasbinRuleReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Ok {
		i--
		if m.Ok {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPbcasbinRule(dAtA []byte, offset int, v uint64) int {
	offset -= sovPbcasbinRule(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CheckCasbinRuleArgs) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Role != 0 {
		n += 1 + sovPbcasbinRule(uint64(m.Role))
	}
	l = len(m.Path)
	if l > 0 {
		n += 1 + l + sovPbcasbinRule(uint64(l))
	}
	l = len(m.Method)
	if l > 0 {
		n += 1 + l + sovPbcasbinRule(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *CheckCasbinRuleReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Ok {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovPbcasbinRule(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPbcasbinRule(x uint64) (n int) {
	return sovPbcasbinRule(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CheckCasbinRuleArgs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPbcasbinRule
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
			return fmt.Errorf("proto: CheckCasbinRuleArgs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckCasbinRuleArgs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Role", wireType)
			}
			m.Role = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPbcasbinRule
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Role |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Path", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPbcasbinRule
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
				return ErrInvalidLengthPbcasbinRule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPbcasbinRule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Path = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Method", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPbcasbinRule
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
				return ErrInvalidLengthPbcasbinRule
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPbcasbinRule
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Method = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPbcasbinRule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPbcasbinRule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *CheckCasbinRuleReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPbcasbinRule
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
			return fmt.Errorf("proto: CheckCasbinRuleReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CheckCasbinRuleReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ok", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPbcasbinRule
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
			m.Ok = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipPbcasbinRule(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPbcasbinRule
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipPbcasbinRule(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPbcasbinRule
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
					return 0, ErrIntOverflowPbcasbinRule
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
					return 0, ErrIntOverflowPbcasbinRule
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
				return 0, ErrInvalidLengthPbcasbinRule
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPbcasbinRule
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPbcasbinRule
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPbcasbinRule        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPbcasbinRule          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPbcasbinRule = fmt.Errorf("proto: unexpected end of group")
)
