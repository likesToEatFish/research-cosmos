// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: reserve/oracle/events.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
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

type EventBandIBCAckSuccess struct {
	AckResult string `protobuf:"bytes,1,opt,name=ack_result,json=ackResult,proto3" json:"ack_result,omitempty"`
	ClientId  int64  `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *EventBandIBCAckSuccess) Reset()         { *m = EventBandIBCAckSuccess{} }
func (m *EventBandIBCAckSuccess) String() string { return proto.CompactTextString(m) }
func (*EventBandIBCAckSuccess) ProtoMessage()    {}
func (*EventBandIBCAckSuccess) Descriptor() ([]byte, []int) {
	return fileDescriptor_5441448c19065114, []int{0}
}
func (m *EventBandIBCAckSuccess) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBandIBCAckSuccess) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBandIBCAckSuccess.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBandIBCAckSuccess) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBandIBCAckSuccess.Merge(m, src)
}
func (m *EventBandIBCAckSuccess) XXX_Size() int {
	return m.Size()
}
func (m *EventBandIBCAckSuccess) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBandIBCAckSuccess.DiscardUnknown(m)
}

var xxx_messageInfo_EventBandIBCAckSuccess proto.InternalMessageInfo

func (m *EventBandIBCAckSuccess) GetAckResult() string {
	if m != nil {
		return m.AckResult
	}
	return ""
}

func (m *EventBandIBCAckSuccess) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type EventBandIBCAckError struct {
	AckError string `protobuf:"bytes,1,opt,name=ack_error,json=ackError,proto3" json:"ack_error,omitempty"`
	ClientId int64  `protobuf:"varint,2,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *EventBandIBCAckError) Reset()         { *m = EventBandIBCAckError{} }
func (m *EventBandIBCAckError) String() string { return proto.CompactTextString(m) }
func (*EventBandIBCAckError) ProtoMessage()    {}
func (*EventBandIBCAckError) Descriptor() ([]byte, []int) {
	return fileDescriptor_5441448c19065114, []int{1}
}
func (m *EventBandIBCAckError) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBandIBCAckError) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBandIBCAckError.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBandIBCAckError) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBandIBCAckError.Merge(m, src)
}
func (m *EventBandIBCAckError) XXX_Size() int {
	return m.Size()
}
func (m *EventBandIBCAckError) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBandIBCAckError.DiscardUnknown(m)
}

var xxx_messageInfo_EventBandIBCAckError proto.InternalMessageInfo

func (m *EventBandIBCAckError) GetAckError() string {
	if m != nil {
		return m.AckError
	}
	return ""
}

func (m *EventBandIBCAckError) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type EventBandIBCResponseTimeout struct {
	ClientId int64 `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *EventBandIBCResponseTimeout) Reset()         { *m = EventBandIBCResponseTimeout{} }
func (m *EventBandIBCResponseTimeout) String() string { return proto.CompactTextString(m) }
func (*EventBandIBCResponseTimeout) ProtoMessage()    {}
func (*EventBandIBCResponseTimeout) Descriptor() ([]byte, []int) {
	return fileDescriptor_5441448c19065114, []int{2}
}
func (m *EventBandIBCResponseTimeout) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EventBandIBCResponseTimeout) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EventBandIBCResponseTimeout.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EventBandIBCResponseTimeout) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventBandIBCResponseTimeout.Merge(m, src)
}
func (m *EventBandIBCResponseTimeout) XXX_Size() int {
	return m.Size()
}
func (m *EventBandIBCResponseTimeout) XXX_DiscardUnknown() {
	xxx_messageInfo_EventBandIBCResponseTimeout.DiscardUnknown(m)
}

var xxx_messageInfo_EventBandIBCResponseTimeout proto.InternalMessageInfo

func (m *EventBandIBCResponseTimeout) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

type SetBandIBCPriceEvent struct {
	Relayer     string                        `protobuf:"bytes,1,opt,name=relayer,proto3" json:"relayer,omitempty"`
	Symbols     []string                      `protobuf:"bytes,2,rep,name=symbols,proto3" json:"symbols,omitempty"`
	Prices      []cosmossdk_io_math.LegacyDec `protobuf:"bytes,3,rep,name=prices,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"prices"`
	ResolveTime uint64                        `protobuf:"varint,4,opt,name=resolve_time,json=resolveTime,proto3" json:"resolve_time,omitempty"`
	RequestId   uint64                        `protobuf:"varint,5,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	ClientId    int64                         `protobuf:"varint,6,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty"`
}

func (m *SetBandIBCPriceEvent) Reset()         { *m = SetBandIBCPriceEvent{} }
func (m *SetBandIBCPriceEvent) String() string { return proto.CompactTextString(m) }
func (*SetBandIBCPriceEvent) ProtoMessage()    {}
func (*SetBandIBCPriceEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_5441448c19065114, []int{3}
}
func (m *SetBandIBCPriceEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetBandIBCPriceEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetBandIBCPriceEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SetBandIBCPriceEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetBandIBCPriceEvent.Merge(m, src)
}
func (m *SetBandIBCPriceEvent) XXX_Size() int {
	return m.Size()
}
func (m *SetBandIBCPriceEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_SetBandIBCPriceEvent.DiscardUnknown(m)
}

var xxx_messageInfo_SetBandIBCPriceEvent proto.InternalMessageInfo

func (m *SetBandIBCPriceEvent) GetRelayer() string {
	if m != nil {
		return m.Relayer
	}
	return ""
}

func (m *SetBandIBCPriceEvent) GetSymbols() []string {
	if m != nil {
		return m.Symbols
	}
	return nil
}

func (m *SetBandIBCPriceEvent) GetResolveTime() uint64 {
	if m != nil {
		return m.ResolveTime
	}
	return 0
}

func (m *SetBandIBCPriceEvent) GetRequestId() uint64 {
	if m != nil {
		return m.RequestId
	}
	return 0
}

func (m *SetBandIBCPriceEvent) GetClientId() int64 {
	if m != nil {
		return m.ClientId
	}
	return 0
}

func init() {
	proto.RegisterType((*EventBandIBCAckSuccess)(nil), "reserve.oracle.EventBandIBCAckSuccess")
	proto.RegisterType((*EventBandIBCAckError)(nil), "reserve.oracle.EventBandIBCAckError")
	proto.RegisterType((*EventBandIBCResponseTimeout)(nil), "reserve.oracle.EventBandIBCResponseTimeout")
	proto.RegisterType((*SetBandIBCPriceEvent)(nil), "reserve.oracle.SetBandIBCPriceEvent")
}

func init() { proto.RegisterFile("reserve/oracle/events.proto", fileDescriptor_5441448c19065114) }

var fileDescriptor_5441448c19065114 = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0x3f, 0x6f, 0xd4, 0x40,
	0x10, 0xc5, 0xbd, 0xb9, 0x70, 0xc4, 0x0b, 0xa2, 0xb0, 0x4e, 0xc8, 0xc2, 0xc2, 0x31, 0xa1, 0xb9,
	0xca, 0x2e, 0xe8, 0xa0, 0xc2, 0x90, 0xc2, 0x12, 0x45, 0xe4, 0xa4, 0xa2, 0x39, 0xed, 0xad, 0x47,
	0x8e, 0xe5, 0x3f, 0x63, 0x76, 0xd6, 0x27, 0xfc, 0x2d, 0xf8, 0x58, 0x29, 0x53, 0x22, 0x8a, 0x08,
	0xdd, 0x49, 0x7c, 0x0e, 0xb4, 0x3e, 0x5b, 0xca, 0x9d, 0x94, 0xce, 0xf3, 0xde, 0xbc, 0x9f, 0xf6,
	0x79, 0x97, 0x7b, 0x0a, 0x08, 0xd4, 0x06, 0x22, 0x54, 0x42, 0x56, 0x10, 0xc1, 0x06, 0x1a, 0x4d,
	0x61, 0xab, 0x50, 0xa3, 0xf3, 0x6a, 0x34, 0xc3, 0xbd, 0xf9, 0x66, 0x91, 0x63, 0x8e, 0x83, 0x15,
	0x99, 0xaf, 0xfd, 0xd6, 0xc5, 0x0d, 0x7f, 0x7d, 0x69, 0x52, 0xb1, 0x68, 0xb2, 0x24, 0xfe, 0xf2,
	0x59, 0x96, 0xd7, 0x9d, 0x94, 0x40, 0xe4, 0xbc, 0xe5, 0x5c, 0xc8, 0x72, 0xa5, 0x80, 0xba, 0x4a,
	0xbb, 0x2c, 0x60, 0x4b, 0x3b, 0xb5, 0x85, 0x2c, 0xd3, 0x41, 0x70, 0x3c, 0x6e, 0xcb, 0xaa, 0x80,
	0x46, 0xaf, 0x8a, 0xcc, 0x3d, 0x09, 0xd8, 0x72, 0x96, 0x9e, 0xed, 0x85, 0x24, 0xbb, 0xb8, 0xe2,
	0x8b, 0x23, 0xea, 0xa5, 0x52, 0xa8, 0x4c, 0xc8, 0x30, 0xc1, 0x0c, 0x23, 0xf2, 0x4c, 0x3c, 0x32,
	0x9f, 0x26, 0x7e, 0xe4, 0xde, 0x63, 0x62, 0x0a, 0xd4, 0x62, 0x43, 0x70, 0x53, 0xd4, 0x80, 0xdd,
	0xd1, 0x69, 0xd8, 0x51, 0xf6, 0x1f, 0xe3, 0x8b, 0x6b, 0x98, 0xa2, 0x57, 0xaa, 0x90, 0x30, 0xb0,
	0x1c, 0x97, 0x3f, 0x57, 0x50, 0x89, 0x1e, 0xa6, 0xc3, 0x4c, 0xa3, 0x71, 0xa8, 0xaf, 0xd7, 0x58,
	0x91, 0x7b, 0x12, 0xcc, 0x8c, 0x33, 0x8e, 0xce, 0x27, 0x3e, 0x6f, 0x0d, 0x81, 0xdc, 0x99, 0x31,
	0xe2, 0xf7, 0x77, 0x0f, 0xe7, 0xd6, 0x9f, 0x87, 0x73, 0x4f, 0x22, 0xd5, 0x48, 0x94, 0x95, 0x61,
	0x81, 0x51, 0x2d, 0xf4, 0x6d, 0xf8, 0x0d, 0x72, 0x21, 0xfb, 0xaf, 0x20, 0xd3, 0x31, 0xe2, 0xbc,
	0xe3, 0x2f, 0x15, 0x10, 0x56, 0x1b, 0x58, 0xe9, 0xa2, 0x06, 0xf7, 0x34, 0x60, 0xcb, 0xd3, 0xf4,
	0xc5, 0xa8, 0x99, 0x32, 0xe6, 0xb7, 0x2b, 0xf8, 0xd1, 0x01, 0x0d, 0x55, 0x9e, 0x0d, 0x0b, 0xf6,
	0xa8, 0x24, 0xd9, 0x61, 0xd1, 0xf9, 0x61, 0xd1, 0x38, 0xb9, 0xdb, 0xfa, 0xec, 0x7e, 0xeb, 0xb3,
	0xbf, 0x5b, 0x9f, 0xfd, 0xda, 0xf9, 0xd6, 0xfd, 0xce, 0xb7, 0x7e, 0xef, 0x7c, 0xeb, 0x7b, 0x94,
	0x17, 0xfa, 0xb6, 0x5b, 0x87, 0x12, 0xeb, 0x08, 0x1b, 0xac, 0xfb, 0xe1, 0xf6, 0x25, 0x56, 0xd1,
	0xf4, 0x84, 0x7e, 0x4e, 0x8f, 0x48, 0xf7, 0x2d, 0xd0, 0x7a, 0x3e, 0x2c, 0x7c, 0xf8, 0x1f, 0x00,
	0x00, 0xff, 0xff, 0x05, 0xdc, 0xb3, 0x3d, 0x63, 0x02, 0x00, 0x00,
}

func (m *EventBandIBCAckSuccess) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBandIBCAckSuccess) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBandIBCAckSuccess) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ClientId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AckResult) > 0 {
		i -= len(m.AckResult)
		copy(dAtA[i:], m.AckResult)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AckResult)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBandIBCAckError) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBandIBCAckError) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBandIBCAckError) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ClientId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.AckError) > 0 {
		i -= len(m.AckError)
		copy(dAtA[i:], m.AckError)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.AckError)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EventBandIBCResponseTimeout) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EventBandIBCResponseTimeout) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EventBandIBCResponseTimeout) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ClientId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *SetBandIBCPriceEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetBandIBCPriceEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SetBandIBCPriceEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ClientId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ClientId))
		i--
		dAtA[i] = 0x30
	}
	if m.RequestId != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.RequestId))
		i--
		dAtA[i] = 0x28
	}
	if m.ResolveTime != 0 {
		i = encodeVarintEvents(dAtA, i, uint64(m.ResolveTime))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Prices) > 0 {
		for iNdEx := len(m.Prices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size := m.Prices[iNdEx].Size()
				i -= size
				if _, err := m.Prices[iNdEx].MarshalTo(dAtA[i:]); err != nil {
					return 0, err
				}
				i = encodeVarintEvents(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Symbols) > 0 {
		for iNdEx := len(m.Symbols) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Symbols[iNdEx])
			copy(dAtA[i:], m.Symbols[iNdEx])
			i = encodeVarintEvents(dAtA, i, uint64(len(m.Symbols[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Relayer) > 0 {
		i -= len(m.Relayer)
		copy(dAtA[i:], m.Relayer)
		i = encodeVarintEvents(dAtA, i, uint64(len(m.Relayer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvents(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvents(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *EventBandIBCAckSuccess) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AckResult)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ClientId != 0 {
		n += 1 + sovEvents(uint64(m.ClientId))
	}
	return n
}

func (m *EventBandIBCAckError) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AckError)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if m.ClientId != 0 {
		n += 1 + sovEvents(uint64(m.ClientId))
	}
	return n
}

func (m *EventBandIBCResponseTimeout) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClientId != 0 {
		n += 1 + sovEvents(uint64(m.ClientId))
	}
	return n
}

func (m *SetBandIBCPriceEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Relayer)
	if l > 0 {
		n += 1 + l + sovEvents(uint64(l))
	}
	if len(m.Symbols) > 0 {
		for _, s := range m.Symbols {
			l = len(s)
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if len(m.Prices) > 0 {
		for _, e := range m.Prices {
			l = e.Size()
			n += 1 + l + sovEvents(uint64(l))
		}
	}
	if m.ResolveTime != 0 {
		n += 1 + sovEvents(uint64(m.ResolveTime))
	}
	if m.RequestId != 0 {
		n += 1 + sovEvents(uint64(m.RequestId))
	}
	if m.ClientId != 0 {
		n += 1 + sovEvents(uint64(m.ClientId))
	}
	return n
}

func sovEvents(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvents(x uint64) (n int) {
	return sovEvents(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EventBandIBCAckSuccess) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBandIBCAckSuccess: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBandIBCAckSuccess: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckResult", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckResult = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			m.ClientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBandIBCAckError) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBandIBCAckError: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBandIBCAckError: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AckError", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AckError = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			m.ClientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *EventBandIBCResponseTimeout) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: EventBandIBCResponseTimeout: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EventBandIBCResponseTimeout: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			m.ClientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func (m *SetBandIBCPriceEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvents
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
			return fmt.Errorf("proto: SetBandIBCPriceEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetBandIBCPriceEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Relayer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Relayer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbols", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbols = append(m.Symbols, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prices", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
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
				return ErrInvalidLengthEvents
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvents
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.LegacyDec
			m.Prices = append(m.Prices, v)
			if err := m.Prices[len(m.Prices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ResolveTime", wireType)
			}
			m.ResolveTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ResolveTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RequestId", wireType)
			}
			m.RequestId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RequestId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClientId", wireType)
			}
			m.ClientId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvents
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClientId |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvents(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvents
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
func skipEvents(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
					return 0, ErrIntOverflowEvents
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
				return 0, ErrInvalidLengthEvents
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvents
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvents
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvents        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvents          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvents = fmt.Errorf("proto: unexpected end of group")
)
