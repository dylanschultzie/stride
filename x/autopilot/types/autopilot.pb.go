// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/autopilot/autopilot.proto

package types

import (
	fmt "fmt"
	proto "github.com/cosmos/gogoproto/proto"
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

type TransferCallback struct {
	FallbackAddress string `protobuf:"bytes,1,opt,name=fallback_address,json=fallbackAddress,proto3" json:"fallback_address,omitempty"`
}

func (m *TransferCallback) Reset()         { *m = TransferCallback{} }
func (m *TransferCallback) String() string { return proto.CompactTextString(m) }
func (*TransferCallback) ProtoMessage()    {}
func (*TransferCallback) Descriptor() ([]byte, []int) {
	return fileDescriptor_cf12981bf14863a6, []int{0}
}
func (m *TransferCallback) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TransferCallback) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TransferCallback.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TransferCallback) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TransferCallback.Merge(m, src)
}
func (m *TransferCallback) XXX_Size() int {
	return m.Size()
}
func (m *TransferCallback) XXX_DiscardUnknown() {
	xxx_messageInfo_TransferCallback.DiscardUnknown(m)
}

var xxx_messageInfo_TransferCallback proto.InternalMessageInfo

func (m *TransferCallback) GetFallbackAddress() string {
	if m != nil {
		return m.FallbackAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*TransferCallback)(nil), "stride.autopilot.TransferCallback")
}

func init() { proto.RegisterFile("stride/autopilot/autopilot.proto", fileDescriptor_cf12981bf14863a6) }

var fileDescriptor_cf12981bf14863a6 = []byte{
	// 174 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x2e, 0x29, 0xca,
	0x4c, 0x49, 0xd5, 0x4f, 0x2c, 0x2d, 0xc9, 0x2f, 0xc8, 0xcc, 0xc9, 0x2f, 0x41, 0xb0, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x04, 0x20, 0x2a, 0xf4, 0xe0, 0xe2, 0x4a, 0xb6, 0x5c, 0x02, 0x21,
	0x45, 0x89, 0x79, 0xc5, 0x69, 0xa9, 0x45, 0xce, 0x89, 0x39, 0x39, 0x49, 0x89, 0xc9, 0xd9, 0x42,
	0x9a, 0x5c, 0x02, 0x69, 0x50, 0x76, 0x7c, 0x62, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3,
	0x02, 0xa3, 0x06, 0x67, 0x10, 0x3f, 0x4c, 0xdc, 0x11, 0x22, 0xec, 0xe4, 0x7b, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xc6, 0xe9, 0x99, 0x25, 0x19, 0xa5, 0x49, 0x7a, 0xc9,
	0xf9, 0xb9, 0xfa, 0xc1, 0x60, 0x5b, 0x75, 0x7d, 0x12, 0x93, 0x8a, 0xf5, 0xa1, 0x6e, 0x2c, 0x33,
	0x34, 0xd3, 0xaf, 0x40, 0x72, 0x69, 0x49, 0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0xd8, 0x99, 0xc6,
	0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7b, 0x78, 0xa0, 0xef, 0xca, 0x00, 0x00, 0x00,
}

func (m *TransferCallback) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TransferCallback) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TransferCallback) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.FallbackAddress) > 0 {
		i -= len(m.FallbackAddress)
		copy(dAtA[i:], m.FallbackAddress)
		i = encodeVarintAutopilot(dAtA, i, uint64(len(m.FallbackAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintAutopilot(dAtA []byte, offset int, v uint64) int {
	offset -= sovAutopilot(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TransferCallback) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.FallbackAddress)
	if l > 0 {
		n += 1 + l + sovAutopilot(uint64(l))
	}
	return n
}

func sovAutopilot(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAutopilot(x uint64) (n int) {
	return sovAutopilot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TransferCallback) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAutopilot
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
			return fmt.Errorf("proto: TransferCallback: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TransferCallback: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FallbackAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAutopilot
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
				return ErrInvalidLengthAutopilot
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAutopilot
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FallbackAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAutopilot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAutopilot
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
func skipAutopilot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAutopilot
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
					return 0, ErrIntOverflowAutopilot
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
					return 0, ErrIntOverflowAutopilot
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
				return 0, ErrInvalidLengthAutopilot
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAutopilot
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAutopilot
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAutopilot        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAutopilot          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAutopilot = fmt.Errorf("proto: unexpected end of group")
)