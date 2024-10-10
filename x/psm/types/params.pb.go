// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/psm/v1/params.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type Params struct {
	// total $nomUSD can mint
	LimitTotal cosmossdk_io_math.Int `protobuf:"bytes,1,opt,name=limit_total,json=limitTotal,proto3,customtype=cosmossdk.io/math.Int" json:"limit_total"`
	// The price cannot be exactly 1, an acceptable such as 0.9999 (AcceptablePriceRatio = 0.0001)
	AcceptablePriceRatio cosmossdk_io_math.LegacyDec `protobuf:"bytes,2,opt,name=acceptable_price_ratio,json=acceptablePriceRatio,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"acceptable_price_ratio"`
	// feeIn adjustment factor
	AdjustmentFeeIn cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,opt,name=adjustment_feeIn,json=adjustmentFeeIn,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"adjustment_feeIn"`
	// feeIn adjustment factor
	AdjustmentFeeOut cosmossdk_io_math.LegacyDec `protobuf:"bytes,4,opt,name=adjustment_feeOut,json=adjustmentFeeOut,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"adjustment_feeOut"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_e516259d7293aa1e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func init() {
	proto.RegisterType((*Params)(nil), "reserve.psm.v1.Params")
}

func init() { proto.RegisterFile("reserve/psm/v1/params.proto", fileDescriptor_e516259d7293aa1e) }

var fileDescriptor_e516259d7293aa1e = []byte{
	// 357 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0xd2, 0xc1, 0x4a, 0xe3, 0x40,
	0x18, 0x07, 0xf0, 0x64, 0xbb, 0xf4, 0x30, 0xbb, 0xec, 0xb6, 0xa1, 0x4a, 0x6c, 0x21, 0x15, 0x4f,
	0x22, 0x9a, 0xb1, 0xf8, 0x06, 0xa5, 0x14, 0x0a, 0x42, 0x6b, 0xf1, 0x24, 0x62, 0x98, 0x4e, 0xc7,
	0x34, 0x9a, 0xc9, 0x0c, 0x99, 0x2f, 0xc5, 0x5e, 0x7d, 0x02, 0x1f, 0xc3, 0xa3, 0x07, 0x1f, 0xa2,
	0xc7, 0xe2, 0x49, 0x3c, 0x14, 0x69, 0x0f, 0xbe, 0x86, 0x64, 0x12, 0xa9, 0xc5, 0x5b, 0x2f, 0xc3,
	0x37, 0xf3, 0x9f, 0xf9, 0x7d, 0x87, 0xf9, 0x50, 0x2d, 0x66, 0x8a, 0xc5, 0x63, 0x86, 0xa5, 0xe2,
	0x78, 0xdc, 0xc0, 0x92, 0xc4, 0x84, 0x2b, 0x57, 0xc6, 0x02, 0x84, 0xf5, 0x2f, 0x0f, 0x5d, 0xa9,
	0xb8, 0x3b, 0x6e, 0x54, 0xcb, 0x84, 0x07, 0x91, 0xc0, 0x7a, 0xcd, 0xae, 0x54, 0x2b, 0xbe, 0xf0,
	0x85, 0x2e, 0x71, 0x5a, 0xe5, 0xa7, 0x3b, 0x54, 0x28, 0x2e, 0x94, 0x97, 0x05, 0xd9, 0x26, 0x8b,
	0xf6, 0xee, 0x0b, 0xa8, 0xd8, 0xd3, 0x4d, 0xac, 0x33, 0xf4, 0x27, 0x0c, 0x78, 0x00, 0x1e, 0x08,
	0x20, 0xa1, 0x6d, 0xee, 0x9a, 0xfb, 0x7f, 0x9b, 0xc7, 0xd3, 0x79, 0xdd, 0x78, 0x9b, 0xd7, 0xb7,
	0xb2, 0x57, 0x6a, 0x78, 0xeb, 0x06, 0x02, 0x73, 0x02, 0x23, 0xb7, 0x13, 0xc1, 0xcb, 0xf3, 0x11,
	0xca, 0xb9, 0x4e, 0x04, 0x8f, 0x1f, 0x4f, 0x07, 0x66, 0x1f, 0x69, 0xe4, 0x3c, 0x35, 0x2c, 0x1f,
	0x6d, 0x13, 0x4a, 0x99, 0x04, 0x32, 0x08, 0x99, 0x27, 0xe3, 0x80, 0x32, 0x2f, 0x26, 0x10, 0x08,
	0xfb, 0x97, 0xd6, 0x1b, 0xb9, 0x5e, 0xfb, 0xa9, 0x9f, 0x32, 0x9f, 0xd0, 0x49, 0x8b, 0xd1, 0x6f,
	0x3d, 0x5a, 0x8c, 0xf6, 0x2b, 0x2b, 0xb0, 0x97, 0x7a, 0xfd, 0x94, 0xb3, 0x2e, 0x51, 0x89, 0x0c,
	0x6f, 0x12, 0x05, 0x9c, 0x45, 0xe0, 0x5d, 0x33, 0xd6, 0x89, 0xec, 0xc2, 0xa6, 0x2d, 0xfe, 0xaf,
	0xa8, 0x76, 0x2a, 0x59, 0x57, 0xa8, 0xbc, 0xae, 0x77, 0x13, 0xb0, 0x7f, 0x6f, 0xca, 0x97, 0xd6,
	0xf8, 0x6e, 0x02, 0xcd, 0xf6, 0x74, 0xe1, 0x98, 0xb3, 0x85, 0x63, 0xbe, 0x2f, 0x1c, 0xf3, 0x61,
	0xe9, 0x18, 0xb3, 0xa5, 0x63, 0xbc, 0x2e, 0x1d, 0xe3, 0xe2, 0xd0, 0x0f, 0x60, 0x94, 0x0c, 0x5c,
	0x2a, 0x38, 0x16, 0x91, 0xe0, 0x13, 0xfd, 0x6b, 0x54, 0x84, 0xf8, 0x6b, 0x50, 0xee, 0xf4, 0xa8,
	0xc0, 0x44, 0x32, 0x35, 0x28, 0xea, 0xf4, 0xe4, 0x33, 0x00, 0x00, 0xff, 0xff, 0x65, 0xe2, 0x6b,
	0x15, 0x46, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.AdjustmentFeeOut.Size()
		i -= size
		if _, err := m.AdjustmentFeeOut.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.AdjustmentFeeIn.Size()
		i -= size
		if _, err := m.AdjustmentFeeIn.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size := m.AcceptablePriceRatio.Size()
		i -= size
		if _, err := m.AcceptablePriceRatio.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	{
		size := m.LimitTotal.Size()
		i -= size
		if _, err := m.LimitTotal.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintParams(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintParams(dAtA []byte, offset int, v uint64) int {
	offset -= sovParams(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.LimitTotal.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.AcceptablePriceRatio.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.AdjustmentFeeIn.Size()
	n += 1 + l + sovParams(uint64(l))
	l = m.AdjustmentFeeOut.Size()
	n += 1 + l + sovParams(uint64(l))
	return n
}

func sovParams(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozParams(x uint64) (n int) {
	return sovParams(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowParams
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitTotal", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LimitTotal.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AcceptablePriceRatio", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AcceptablePriceRatio.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustmentFeeIn", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdjustmentFeeIn.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AdjustmentFeeOut", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowParams
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
				return ErrInvalidLengthParams
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthParams
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AdjustmentFeeOut.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipParams(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthParams
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
func skipParams(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
					return 0, ErrIntOverflowParams
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
				return 0, ErrInvalidLengthParams
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupParams
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthParams
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthParams        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowParams          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupParams = fmt.Errorf("proto: unexpected end of group")
)
