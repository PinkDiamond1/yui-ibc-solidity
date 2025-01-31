// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: client/ibft2/IBFT2.proto

package ibft2

import (
	fmt "fmt"
	_ "github.com/datachainlab/solidity-protobuf/protobuf-solidity/src/protoc/go"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	client "github.com/hyperledger-labs/yui-ibc-solidity/pkg/ibc/client"
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

type ClientState struct {
	ChainId         string        `protobuf:"bytes,1,opt,name=chain_id,json=chainId,proto3" json:"chain_id,omitempty"`
	IbcStoreAddress []byte        `protobuf:"bytes,2,opt,name=ibc_store_address,json=ibcStoreAddress,proto3" json:"ibc_store_address,omitempty"`
	LatestHeight    client.Height `protobuf:"bytes,3,opt,name=latest_height,json=latestHeight,proto3" json:"latest_height"`
}

func (m *ClientState) Reset()         { *m = ClientState{} }
func (m *ClientState) String() string { return proto.CompactTextString(m) }
func (*ClientState) ProtoMessage()    {}
func (*ClientState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc22388094e2414, []int{0}
}
func (m *ClientState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClientState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClientState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClientState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClientState.Merge(m, src)
}
func (m *ClientState) XXX_Size() int {
	return m.Size()
}
func (m *ClientState) XXX_DiscardUnknown() {
	xxx_messageInfo_ClientState.DiscardUnknown(m)
}

var xxx_messageInfo_ClientState proto.InternalMessageInfo

type ConsensusState struct {
	Timestamp  uint64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Root       []byte   `protobuf:"bytes,2,opt,name=root,proto3" json:"root,omitempty"`
	Validators [][]byte `protobuf:"bytes,3,rep,name=validators,proto3" json:"validators,omitempty"`
}

func (m *ConsensusState) Reset()         { *m = ConsensusState{} }
func (m *ConsensusState) String() string { return proto.CompactTextString(m) }
func (*ConsensusState) ProtoMessage()    {}
func (*ConsensusState) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc22388094e2414, []int{1}
}
func (m *ConsensusState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusState.Merge(m, src)
}
func (m *ConsensusState) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusState) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusState.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusState proto.InternalMessageInfo

type Header struct {
	BesuHeaderRlp     []byte        `protobuf:"bytes,1,opt,name=besu_header_rlp,json=besuHeaderRlp,proto3" json:"besu_header_rlp,omitempty"`
	Seals             [][]byte      `protobuf:"bytes,2,rep,name=seals,proto3" json:"seals,omitempty"`
	TrustedHeight     client.Height `protobuf:"bytes,3,opt,name=trusted_height,json=trustedHeight,proto3" json:"trusted_height"`
	AccountStateProof []byte        `protobuf:"bytes,4,opt,name=account_state_proof,json=accountStateProof,proto3" json:"account_state_proof,omitempty"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_0cc22388094e2414, []int{2}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Header.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(m, src)
}
func (m *Header) XXX_Size() int {
	return m.Size()
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ClientState)(nil), "ibc.lightclients.ibft2.v1.ClientState")
	proto.RegisterType((*ConsensusState)(nil), "ibc.lightclients.ibft2.v1.ConsensusState")
	proto.RegisterType((*Header)(nil), "ibc.lightclients.ibft2.v1.Header")
}

func init() { proto.RegisterFile("client/ibft2/IBFT2.proto", fileDescriptor_0cc22388094e2414) }

var fileDescriptor_0cc22388094e2414 = []byte{
	// 482 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x92, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x63, 0x12, 0x5a, 0x7a, 0x4d, 0x5a, 0xf5, 0xda, 0xc1, 0xad, 0x90, 0x89, 0x32, 0xa0,
	0x08, 0x29, 0xb6, 0x08, 0x48, 0x88, 0x8d, 0xa6, 0x12, 0x6a, 0x37, 0xe4, 0x32, 0xb1, 0x58, 0x77,
	0xe7, 0x8b, 0x7d, 0xe2, 0xe2, 0xb3, 0xee, 0x3d, 0x57, 0x64, 0x65, 0x62, 0xe4, 0x03, 0xf0, 0x11,
	0x58, 0xf9, 0x0e, 0x19, 0x3b, 0x32, 0x21, 0x48, 0xbe, 0x08, 0xf2, 0x5d, 0x10, 0x9d, 0x60, 0xf2,
	0xfd, 0x7f, 0xf7, 0xf7, 0x7b, 0xef, 0xaf, 0x7b, 0x24, 0x14, 0x5a, 0xc9, 0x0a, 0x13, 0xc5, 0xe7,
	0x38, 0x4d, 0xae, 0x66, 0xaf, 0xdf, 0x4e, 0xe3, 0xda, 0x1a, 0x34, 0xf4, 0x54, 0x71, 0x11, 0x6b,
	0x55, 0x94, 0xe8, 0x2d, 0x10, 0x3b, 0x4f, 0x7c, 0xf3, 0xf4, 0xec, 0xa4, 0x30, 0x85, 0x71, 0xae,
	0xa4, 0x3d, 0xf9, 0x1f, 0xce, 0x46, 0x60, 0xb4, 0xca, 0x15, 0x2e, 0x27, 0x4e, 0xf3, 0x66, 0x3e,
	0x91, 0x1f, 0x50, 0x56, 0xa0, 0x4c, 0x05, 0x5b, 0xcf, 0xf1, 0xb6, 0xdd, 0x85, 0xfb, 0x78, 0x38,
	0xfa, 0x14, 0x90, 0x7d, 0x0f, 0xae, 0x91, 0xa1, 0xa4, 0xa7, 0xe4, 0x81, 0x28, 0x99, 0xaa, 0x32,
	0x95, 0x87, 0xc1, 0x30, 0x18, 0xef, 0xa5, 0xbb, 0x4e, 0x5f, 0xe5, 0xf4, 0x09, 0x39, 0x52, 0x5c,
	0x64, 0x80, 0xc6, 0xca, 0x8c, 0xe5, 0xb9, 0x95, 0x00, 0xe1, 0xbd, 0x61, 0x30, 0xee, 0xa7, 0x87,
	0x8a, 0x8b, 0xeb, 0x96, 0x9f, 0x7b, 0x4c, 0xa7, 0x64, 0xa0, 0x19, 0x4a, 0xc0, 0xac, 0x94, 0x6d,
	0x8c, 0xb0, 0x3b, 0x0c, 0xc6, 0xfb, 0xd3, 0xdd, 0xf8, 0xd2, 0xc9, 0x59, 0x6f, 0xf5, 0xe3, 0x51,
	0x27, 0xed, 0x7b, 0x8f, 0x67, 0x23, 0x4e, 0x0e, 0x2e, 0x4c, 0x05, 0xb2, 0x82, 0x06, 0xfc, 0x30,
	0x0f, 0xc9, 0x1e, 0xaa, 0x85, 0x04, 0x64, 0x8b, 0xda, 0x4d, 0xd3, 0x4b, 0xff, 0x02, 0x4a, 0x49,
	0xcf, 0x1a, 0x83, 0xdb, 0x11, 0xdc, 0x99, 0x46, 0x84, 0xdc, 0x30, 0xad, 0x72, 0x86, 0xc6, 0x42,
	0xd8, 0x1d, 0x76, 0xc7, 0xfd, 0xf4, 0x0e, 0x19, 0x7d, 0x0d, 0xc8, 0xce, 0xa5, 0x64, 0xb9, 0xb4,
	0xf4, 0x31, 0x39, 0xe4, 0x12, 0x9a, 0xac, 0x74, 0x32, 0xb3, 0xda, 0xb7, 0xe8, 0xa7, 0x83, 0x16,
	0x7b, 0x53, 0xaa, 0x6b, 0x7a, 0x42, 0xee, 0x83, 0x64, 0xba, 0x8d, 0xda, 0x56, 0xf3, 0x82, 0x3e,
	0x27, 0x07, 0x68, 0x1b, 0x40, 0x99, 0xff, 0x33, 0xe1, 0x60, 0x6b, 0xf2, 0x90, 0xc6, 0xe4, 0x98,
	0x09, 0x61, 0x9a, 0x0a, 0x33, 0x68, 0x13, 0x66, 0xb5, 0x35, 0x66, 0x1e, 0xf6, 0x5c, 0xdf, 0xa3,
	0xed, 0x95, 0xcb, 0xfe, 0xa6, 0xbd, 0x98, 0x7d, 0x09, 0x3e, 0x7e, 0x0b, 0x5f, 0x92, 0x17, 0xaf,
	0xca, 0x65, 0x2d, 0xad, 0x96, 0x79, 0x21, 0xed, 0x44, 0x33, 0x0e, 0xc9, 0xb2, 0x51, 0x13, 0xc5,
	0xc5, 0xe4, 0xcf, 0xbb, 0x27, 0xc2, 0x54, 0x68, 0x99, 0x40, 0x48, 0x84, 0xb1, 0x32, 0xc1, 0x65,
	0x2d, 0x61, 0xf5, 0x2b, 0xea, 0xac, 0xd6, 0x51, 0x70, 0xbb, 0x8e, 0x82, 0x9f, 0xeb, 0x28, 0xf8,
	0xbc, 0x89, 0x3a, 0xb7, 0x9b, 0xa8, 0xf3, 0x7d, 0x13, 0x75, 0xde, 0x9d, 0x17, 0x0a, 0xcb, 0x86,
	0xc7, 0xc2, 0x2c, 0x92, 0xff, 0x97, 0xaf, 0xdf, 0x17, 0x89, 0xe2, 0x22, 0xb9, 0xbb, 0xb2, 0x7c,
	0xc7, 0xed, 0xd0, 0xb3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xaf, 0x02, 0xa3, 0x32, 0xc9, 0x02,
	0x00, 0x00,
}

func (m *ClientState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClientState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClientState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.LatestHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIBFT2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.IbcStoreAddress) > 0 {
		i -= len(m.IbcStoreAddress)
		copy(dAtA[i:], m.IbcStoreAddress)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.IbcStoreAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ChainId) > 0 {
		i -= len(m.ChainId)
		copy(dAtA[i:], m.ChainId)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.ChainId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Validators) > 0 {
		for iNdEx := len(m.Validators) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Validators[iNdEx])
			copy(dAtA[i:], m.Validators[iNdEx])
			i = encodeVarintIBFT2(dAtA, i, uint64(len(m.Validators[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Root) > 0 {
		i -= len(m.Root)
		copy(dAtA[i:], m.Root)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.Root)))
		i--
		dAtA[i] = 0x12
	}
	if m.Timestamp != 0 {
		i = encodeVarintIBFT2(dAtA, i, uint64(m.Timestamp))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Header) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Header) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Header) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountStateProof) > 0 {
		i -= len(m.AccountStateProof)
		copy(dAtA[i:], m.AccountStateProof)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.AccountStateProof)))
		i--
		dAtA[i] = 0x22
	}
	{
		size, err := m.TrustedHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintIBFT2(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Seals) > 0 {
		for iNdEx := len(m.Seals) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Seals[iNdEx])
			copy(dAtA[i:], m.Seals[iNdEx])
			i = encodeVarintIBFT2(dAtA, i, uint64(len(m.Seals[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.BesuHeaderRlp) > 0 {
		i -= len(m.BesuHeaderRlp)
		copy(dAtA[i:], m.BesuHeaderRlp)
		i = encodeVarintIBFT2(dAtA, i, uint64(len(m.BesuHeaderRlp)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIBFT2(dAtA []byte, offset int, v uint64) int {
	offset -= sovIBFT2(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ClientState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ChainId)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	l = len(m.IbcStoreAddress)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	l = m.LatestHeight.Size()
	n += 1 + l + sovIBFT2(uint64(l))
	return n
}

func (m *ConsensusState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Timestamp != 0 {
		n += 1 + sovIBFT2(uint64(m.Timestamp))
	}
	l = len(m.Root)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	if len(m.Validators) > 0 {
		for _, b := range m.Validators {
			l = len(b)
			n += 1 + l + sovIBFT2(uint64(l))
		}
	}
	return n
}

func (m *Header) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.BesuHeaderRlp)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	if len(m.Seals) > 0 {
		for _, b := range m.Seals {
			l = len(b)
			n += 1 + l + sovIBFT2(uint64(l))
		}
	}
	l = m.TrustedHeight.Size()
	n += 1 + l + sovIBFT2(uint64(l))
	l = len(m.AccountStateProof)
	if l > 0 {
		n += 1 + l + sovIBFT2(uint64(l))
	}
	return n
}

func sovIBFT2(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIBFT2(x uint64) (n int) {
	return sovIBFT2(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ClientState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIBFT2
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
			return fmt.Errorf("proto: ClientState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClientState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChainId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IbcStoreAddress", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IbcStoreAddress = append(m.IbcStoreAddress[:0], dAtA[iNdEx:postIndex]...)
			if m.IbcStoreAddress == nil {
				m.IbcStoreAddress = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LatestHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.LatestHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIBFT2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIBFT2
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
func (m *ConsensusState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIBFT2
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
			return fmt.Errorf("proto: ConsensusState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Timestamp", wireType)
			}
			m.Timestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Timestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Root", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Root = append(m.Root[:0], dAtA[iNdEx:postIndex]...)
			if m.Root == nil {
				m.Root = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Validators", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Validators = append(m.Validators, make([]byte, postIndex-iNdEx))
			copy(m.Validators[len(m.Validators)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIBFT2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIBFT2
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
func (m *Header) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIBFT2
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
			return fmt.Errorf("proto: Header: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Header: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BesuHeaderRlp", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BesuHeaderRlp = append(m.BesuHeaderRlp[:0], dAtA[iNdEx:postIndex]...)
			if m.BesuHeaderRlp == nil {
				m.BesuHeaderRlp = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seals", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seals = append(m.Seals, make([]byte, postIndex-iNdEx))
			copy(m.Seals[len(m.Seals)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrustedHeight", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.TrustedHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountStateProof", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIBFT2
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
				return ErrInvalidLengthIBFT2
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthIBFT2
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AccountStateProof = append(m.AccountStateProof[:0], dAtA[iNdEx:postIndex]...)
			if m.AccountStateProof == nil {
				m.AccountStateProof = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIBFT2(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthIBFT2
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
func skipIBFT2(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIBFT2
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
					return 0, ErrIntOverflowIBFT2
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
					return 0, ErrIntOverflowIBFT2
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
				return 0, ErrInvalidLengthIBFT2
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIBFT2
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIBFT2
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIBFT2        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIBFT2          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIBFT2 = fmt.Errorf("proto: unexpected end of group")
)
