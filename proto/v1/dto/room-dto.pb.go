// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/v1/dto/room-dto.proto

package opt

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

type ListRoomReply struct {
	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *ListRoomReply) Reset()         { *m = ListRoomReply{} }
func (m *ListRoomReply) String() string { return proto.CompactTextString(m) }
func (*ListRoomReply) ProtoMessage()    {}
func (*ListRoomReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_6db50da38d0e1e85, []int{0}
}
func (m *ListRoomReply) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ListRoomReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ListRoomReply.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ListRoomReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListRoomReply.Merge(m, src)
}
func (m *ListRoomReply) XXX_Size() int {
	return m.Size()
}
func (m *ListRoomReply) XXX_DiscardUnknown() {
	xxx_messageInfo_ListRoomReply.DiscardUnknown(m)
}

var xxx_messageInfo_ListRoomReply proto.InternalMessageInfo

func (m *ListRoomReply) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ListRoomReply) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*ListRoomReply)(nil), "dto.ListRoomReply")
}

func init() { proto.RegisterFile("proto/v1/dto/room-dto.proto", fileDescriptor_6db50da38d0e1e85) }

var fileDescriptor_6db50da38d0e1e85 = []byte{
	// 168 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x2f, 0x33, 0xd4, 0x4f, 0x29, 0xc9, 0xd7, 0x2f, 0xca, 0xcf, 0xcf, 0xd5, 0x4d, 0x29,
	0xc9, 0xd7, 0x03, 0x8b, 0x0a, 0x31, 0xa7, 0x94, 0xe4, 0x2b, 0x19, 0x73, 0xf1, 0xfa, 0x64, 0x16,
	0x97, 0x04, 0xe5, 0xe7, 0xe7, 0x06, 0xa5, 0x16, 0xe4, 0x54, 0x0a, 0xf1, 0x71, 0x31, 0x65, 0xa6,
	0x48, 0x30, 0x2a, 0x30, 0x6a, 0x30, 0x07, 0x31, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0xe4, 0x25,
	0xe6, 0xa6, 0x4a, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4e, 0x76, 0x27, 0x1e, 0xc9,
	0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e,
	0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7, 0x10, 0xa5, 0x92, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97,
	0x9c, 0x9f, 0xab, 0x9f, 0x98, 0x53, 0x9a, 0x9d, 0xa8, 0x6b, 0xae, 0x9f, 0x9c, 0x5f, 0x94, 0xaa,
	0x0f, 0x77, 0x48, 0x7e, 0x41, 0x49, 0x12, 0x1b, 0x98, 0x67, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff,
	0x91, 0x00, 0x3c, 0xf2, 0x9f, 0x00, 0x00, 0x00,
}

func (m *ListRoomReply) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ListRoomReply) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ListRoomReply) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintRoomDto(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintRoomDto(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintRoomDto(dAtA []byte, offset int, v uint64) int {
	offset -= sovRoomDto(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ListRoomReply) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovRoomDto(uint64(m.Id))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovRoomDto(uint64(l))
	}
	return n
}

func sovRoomDto(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozRoomDto(x uint64) (n int) {
	return sovRoomDto(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ListRoomReply) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowRoomDto
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
			return fmt.Errorf("proto: ListRoomReply: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ListRoomReply: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomDto
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowRoomDto
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
				return ErrInvalidLengthRoomDto
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthRoomDto
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipRoomDto(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthRoomDto
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
func skipRoomDto(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowRoomDto
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
					return 0, ErrIntOverflowRoomDto
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
					return 0, ErrIntOverflowRoomDto
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
				return 0, ErrInvalidLengthRoomDto
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupRoomDto
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthRoomDto
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthRoomDto        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowRoomDto          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupRoomDto = fmt.Errorf("proto: unexpected end of group")
)
