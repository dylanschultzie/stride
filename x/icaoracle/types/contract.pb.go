// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stride/icaoracle/contract.proto

package types

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

// InstanitateOracleContract is the contract-specific instantiate message
type MsgInstantiateOracleContract struct {
	AdminAddress string `protobuf:"bytes,1,opt,name=admin_address,json=adminAddress,proto3" json:"admin_address,omitempty"`
}

func (m *MsgInstantiateOracleContract) Reset()         { *m = MsgInstantiateOracleContract{} }
func (m *MsgInstantiateOracleContract) String() string { return proto.CompactTextString(m) }
func (*MsgInstantiateOracleContract) ProtoMessage()    {}
func (*MsgInstantiateOracleContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf036e49b48ee03, []int{0}
}
func (m *MsgInstantiateOracleContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgInstantiateOracleContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgInstantiateOracleContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgInstantiateOracleContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgInstantiateOracleContract.Merge(m, src)
}
func (m *MsgInstantiateOracleContract) XXX_Size() int {
	return m.Size()
}
func (m *MsgInstantiateOracleContract) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgInstantiateOracleContract.DiscardUnknown(m)
}

var xxx_messageInfo_MsgInstantiateOracleContract proto.InternalMessageInfo

func (m *MsgInstantiateOracleContract) GetAdminAddress() string {
	if m != nil {
		return m.AdminAddress
	}
	return ""
}

// ExecuteContractPostMetric is the contract-specific metric update message
type MsgExecuteContractPostMetric struct {
	PostMetric *Metric `protobuf:"bytes,1,opt,name=post_metric,json=postMetric,proto3" json:"post_metric,omitempty"`
}

func (m *MsgExecuteContractPostMetric) Reset()         { *m = MsgExecuteContractPostMetric{} }
func (m *MsgExecuteContractPostMetric) String() string { return proto.CompactTextString(m) }
func (*MsgExecuteContractPostMetric) ProtoMessage()    {}
func (*MsgExecuteContractPostMetric) Descriptor() ([]byte, []int) {
	return fileDescriptor_8bf036e49b48ee03, []int{1}
}
func (m *MsgExecuteContractPostMetric) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgExecuteContractPostMetric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgExecuteContractPostMetric.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgExecuteContractPostMetric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgExecuteContractPostMetric.Merge(m, src)
}
func (m *MsgExecuteContractPostMetric) XXX_Size() int {
	return m.Size()
}
func (m *MsgExecuteContractPostMetric) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgExecuteContractPostMetric.DiscardUnknown(m)
}

var xxx_messageInfo_MsgExecuteContractPostMetric proto.InternalMessageInfo

func (m *MsgExecuteContractPostMetric) GetPostMetric() *Metric {
	if m != nil {
		return m.PostMetric
	}
	return nil
}

func init() {
	proto.RegisterType((*MsgInstantiateOracleContract)(nil), "stride.icaoracle.MsgInstantiateOracleContract")
	proto.RegisterType((*MsgExecuteContractPostMetric)(nil), "stride.icaoracle.MsgExecuteContractPostMetric")
}

func init() { proto.RegisterFile("stride/icaoracle/contract.proto", fileDescriptor_8bf036e49b48ee03) }

var fileDescriptor_8bf036e49b48ee03 = []byte{
	// 245 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x2e, 0x29, 0xca,
	0x4c, 0x49, 0xd5, 0xcf, 0x4c, 0x4e, 0xcc, 0x2f, 0x4a, 0x4c, 0xce, 0x49, 0xd5, 0x4f, 0xce, 0xcf,
	0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0x80, 0x28, 0xd0,
	0x83, 0x2b, 0x90, 0x52, 0xc0, 0xd0, 0x02, 0x67, 0x41, 0xf4, 0x28, 0x39, 0x73, 0xc9, 0xf8, 0x16,
	0xa7, 0x7b, 0xe6, 0x15, 0x97, 0x24, 0xe6, 0x95, 0x64, 0x26, 0x96, 0xa4, 0xfa, 0x83, 0x65, 0x9d,
	0xa1, 0x26, 0x0b, 0x29, 0x73, 0xf1, 0x26, 0xa6, 0xe4, 0x66, 0xe6, 0xc5, 0x27, 0xa6, 0xa4, 0x14,
	0xa5, 0x16, 0x17, 0x4b, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0xf1, 0x80, 0x05, 0x1d, 0x21, 0x62,
	0x4a, 0x91, 0x60, 0x43, 0x5c, 0x2b, 0x52, 0x93, 0x4b, 0x4b, 0xe0, 0x5a, 0x03, 0xf2, 0x8b, 0x4b,
	0x7c, 0x53, 0x4b, 0x8a, 0x32, 0x93, 0x85, 0x2c, 0xb9, 0xb8, 0x0b, 0xf2, 0x8b, 0x4b, 0xe2, 0x73,
	0xc1, 0x5c, 0xb0, 0x11, 0xdc, 0x46, 0x12, 0x7a, 0xe8, 0xce, 0xd5, 0x83, 0x28, 0x0f, 0xe2, 0x2a,
	0x80, 0x6b, 0x75, 0xf2, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4,
	0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xa3,
	0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x60, 0xb0, 0x49, 0xba, 0x3e,
	0x89, 0x49, 0xc5, 0xfa, 0x50, 0x2f, 0x97, 0x99, 0xea, 0x57, 0x20, 0xf9, 0xbb, 0xa4, 0xb2, 0x20,
	0xb5, 0x38, 0x89, 0x0d, 0xec, 0x69, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xae, 0x0c, 0x38,
	0x73, 0x4b, 0x01, 0x00, 0x00,
}

func (m *MsgInstantiateOracleContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgInstantiateOracleContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgInstantiateOracleContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AdminAddress) > 0 {
		i -= len(m.AdminAddress)
		copy(dAtA[i:], m.AdminAddress)
		i = encodeVarintContract(dAtA, i, uint64(len(m.AdminAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgExecuteContractPostMetric) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgExecuteContractPostMetric) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgExecuteContractPostMetric) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PostMetric != nil {
		{
			size, err := m.PostMetric.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintContract(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgInstantiateOracleContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AdminAddress)
	if l > 0 {
		n += 1 + l + sovContract(uint64(l))
	}
	return n
}

func (m *MsgExecuteContractPostMetric) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.PostMetric != nil {
		l = m.PostMetric.Size()
		n += 1 + l + sovContract(uint64(l))
	}
	return n
}

func sovContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozContract(x uint64) (n int) {
	return sovContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgInstantiateOracleContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: MsgInstantiateOracleContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgInstantiateOracleContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdminAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AdminAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func (m *MsgExecuteContractPostMetric) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContract
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
			return fmt.Errorf("proto: MsgExecuteContractPostMetric: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgExecuteContractPostMetric: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PostMetric", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContract
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
				return ErrInvalidLengthContract
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PostMetric == nil {
				m.PostMetric = &Metric{}
			}
			if err := m.PostMetric.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthContract
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
func skipContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
					return 0, ErrIntOverflowContract
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
				return 0, ErrInvalidLengthContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupContract = fmt.Errorf("proto: unexpected end of group")
)