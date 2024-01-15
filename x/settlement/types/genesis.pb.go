// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: settlus/settlement/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the settlement module's genesis state.
type GenesisState struct {
	Params  Params                `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Tenants []Tenant              `protobuf:"bytes,2,rep,name=tenants,proto3" json:"tenants"`
	Utxrs   []UTXRWithTenantAndId `protobuf:"bytes,3,rep,name=utxrs,proto3" json:"utxrs"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_50b4d7b20b5c1163, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetTenants() []Tenant {
	if m != nil {
		return m.Tenants
	}
	return nil
}

func (m *GenesisState) GetUtxrs() []UTXRWithTenantAndId {
	if m != nil {
		return m.Utxrs
	}
	return nil
}

type UTXRWithTenantAndId struct {
	TenantId uint64 `protobuf:"varint,1,opt,name=tenant_id,json=tenantId,proto3" json:"tenant_id,omitempty"`
	Id       uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Utxr     UTXR   `protobuf:"bytes,3,opt,name=utxr,proto3" json:"utxr"`
}

func (m *UTXRWithTenantAndId) Reset()         { *m = UTXRWithTenantAndId{} }
func (m *UTXRWithTenantAndId) String() string { return proto.CompactTextString(m) }
func (*UTXRWithTenantAndId) ProtoMessage()    {}
func (*UTXRWithTenantAndId) Descriptor() ([]byte, []int) {
	return fileDescriptor_50b4d7b20b5c1163, []int{1}
}
func (m *UTXRWithTenantAndId) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UTXRWithTenantAndId) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UTXRWithTenantAndId.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UTXRWithTenantAndId) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UTXRWithTenantAndId.Merge(m, src)
}
func (m *UTXRWithTenantAndId) XXX_Size() int {
	return m.Size()
}
func (m *UTXRWithTenantAndId) XXX_DiscardUnknown() {
	xxx_messageInfo_UTXRWithTenantAndId.DiscardUnknown(m)
}

var xxx_messageInfo_UTXRWithTenantAndId proto.InternalMessageInfo

func (m *UTXRWithTenantAndId) GetTenantId() uint64 {
	if m != nil {
		return m.TenantId
	}
	return 0
}

func (m *UTXRWithTenantAndId) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *UTXRWithTenantAndId) GetUtxr() UTXR {
	if m != nil {
		return m.Utxr
	}
	return UTXR{}
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "settlus.settlement.GenesisState")
	proto.RegisterType((*UTXRWithTenantAndId)(nil), "settlus.settlement.UTXRWithTenantAndId")
}

func init() { proto.RegisterFile("settlus/settlement/genesis.proto", fileDescriptor_50b4d7b20b5c1163) }

var fileDescriptor_50b4d7b20b5c1163 = []byte{
	// 308 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x28, 0x4e, 0x2d, 0x29,
	0xc9, 0x29, 0x2d, 0xd6, 0x07, 0xd3, 0xa9, 0xb9, 0xa9, 0x79, 0x25, 0xfa, 0xe9, 0xa9, 0x79, 0xa9,
	0xc5, 0x99, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x42, 0x50, 0x15, 0x7a, 0x08, 0x15,
	0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x69, 0x7d, 0x10, 0x0b, 0xa2, 0x52, 0x4a, 0x19, 0x8b,
	0x59, 0x08, 0x26, 0x44, 0x91, 0xd2, 0x49, 0x46, 0x2e, 0x1e, 0x77, 0x88, 0x05, 0xc1, 0x25, 0x89,
	0x25, 0xa9, 0x42, 0x16, 0x5c, 0x6c, 0x05, 0x89, 0x45, 0x89, 0xb9, 0xc5, 0x12, 0x8c, 0x0a, 0x8c,
	0x1a, 0xdc, 0x46, 0x52, 0x7a, 0x98, 0x16, 0xea, 0x05, 0x80, 0x55, 0x38, 0xb1, 0x9c, 0xb8, 0x27,
	0xcf, 0x10, 0x04, 0x55, 0x2f, 0x64, 0xc5, 0xc5, 0x5e, 0x92, 0x9a, 0x97, 0x98, 0x57, 0x52, 0x2c,
	0xc1, 0xa4, 0xc0, 0x8c, 0x4b, 0x6b, 0x08, 0x58, 0x09, 0x54, 0x2b, 0x4c, 0x83, 0x90, 0x33, 0x17,
	0x6b, 0x69, 0x49, 0x45, 0x51, 0xb1, 0x04, 0x33, 0x58, 0xa7, 0x3a, 0x36, 0x9d, 0xa1, 0x21, 0x11,
	0x41, 0xe1, 0x99, 0x25, 0x19, 0x10, 0x13, 0x1c, 0xf3, 0x52, 0x3c, 0x53, 0xa0, 0xc6, 0x40, 0xf4,
	0x2a, 0x95, 0x71, 0x09, 0x63, 0x51, 0x23, 0x24, 0xcd, 0xc5, 0x09, 0xb1, 0x26, 0x3e, 0x33, 0x05,
	0xec, 0x29, 0x96, 0x20, 0x0e, 0x88, 0x80, 0x67, 0x8a, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04,
	0x13, 0x58, 0x94, 0x29, 0x33, 0x45, 0xc8, 0x88, 0x8b, 0x05, 0x64, 0x98, 0x04, 0x33, 0xd8, 0xf3,
	0x12, 0xb8, 0xdc, 0x01, 0xb5, 0x18, 0xac, 0xd6, 0xc9, 0xf5, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0xb4, 0xd3, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x61, 0xb1, 0x91, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x5f, 0x81, 0x1c, 0x2b, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0xe0, 0x18, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe0, 0xca, 0x5f, 0xd7,
	0x04, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Utxrs) > 0 {
		for iNdEx := len(m.Utxrs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Utxrs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Tenants) > 0 {
		for iNdEx := len(m.Tenants) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Tenants[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *UTXRWithTenantAndId) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UTXRWithTenantAndId) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UTXRWithTenantAndId) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Utxr.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.Id != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.TenantId != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.TenantId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Tenants) > 0 {
		for _, e := range m.Tenants {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Utxrs) > 0 {
		for _, e := range m.Utxrs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func (m *UTXRWithTenantAndId) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.TenantId != 0 {
		n += 1 + sovGenesis(uint64(m.TenantId))
	}
	if m.Id != 0 {
		n += 1 + sovGenesis(uint64(m.Id))
	}
	l = m.Utxr.Size()
	n += 1 + l + sovGenesis(uint64(l))
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tenants", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tenants = append(m.Tenants, Tenant{})
			if err := m.Tenants[len(m.Tenants)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Utxrs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Utxrs = append(m.Utxrs, UTXRWithTenantAndId{})
			if err := m.Utxrs[len(m.Utxrs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func (m *UTXRWithTenantAndId) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: UTXRWithTenantAndId: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UTXRWithTenantAndId: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TenantId", wireType)
			}
			m.TenantId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TenantId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Utxr", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Utxr.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
