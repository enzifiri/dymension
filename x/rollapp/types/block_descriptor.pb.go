// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: dymension/rollapp/block_descriptor.proto

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

// BlockDescriptor defines a single rollapp chain block description.
type BlockDescriptor struct {
	// height is the height of the block
	Height uint64 `protobuf:"varint,1,opt,name=height,proto3" json:"height,omitempty"`
	// stateRoot is a 32 byte array of the hash of the block (state root of the block)
	StateRoot []byte `protobuf:"bytes,2,opt,name=stateRoot,proto3" json:"stateRoot,omitempty"`
	// intermediateStatesRoot is a 32 byte array representing
	// the root of a Merkle tree built from the ISRs of the block (Intermediate State Roots)
	IntermediateStatesRoots [][]byte `protobuf:"bytes,3,rep,name=intermediateStatesRoots,proto3" json:"intermediateStatesRoots,omitempty"`
}

func (m *BlockDescriptor) Reset()         { *m = BlockDescriptor{} }
func (m *BlockDescriptor) String() string { return proto.CompactTextString(m) }
func (*BlockDescriptor) ProtoMessage()    {}
func (*BlockDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e537259fa860c20, []int{0}
}
func (m *BlockDescriptor) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockDescriptor.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDescriptor.Merge(m, src)
}
func (m *BlockDescriptor) XXX_Size() int {
	return m.Size()
}
func (m *BlockDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDescriptor proto.InternalMessageInfo

func (m *BlockDescriptor) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockDescriptor) GetStateRoot() []byte {
	if m != nil {
		return m.StateRoot
	}
	return nil
}

func (m *BlockDescriptor) GetIntermediateStatesRoots() [][]byte {
	if m != nil {
		return m.IntermediateStatesRoots
	}
	return nil
}

// BlockDescriptors defines list of BlockDescriptor.
type BlockDescriptors struct {
	BD []BlockDescriptor `protobuf:"bytes,1,rep,name=BD,proto3" json:"BD"`
}

func (m *BlockDescriptors) Reset()         { *m = BlockDescriptors{} }
func (m *BlockDescriptors) String() string { return proto.CompactTextString(m) }
func (*BlockDescriptors) ProtoMessage()    {}
func (*BlockDescriptors) Descriptor() ([]byte, []int) {
	return fileDescriptor_9e537259fa860c20, []int{1}
}
func (m *BlockDescriptors) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockDescriptors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockDescriptors.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockDescriptors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockDescriptors.Merge(m, src)
}
func (m *BlockDescriptors) XXX_Size() int {
	return m.Size()
}
func (m *BlockDescriptors) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockDescriptors.DiscardUnknown(m)
}

var xxx_messageInfo_BlockDescriptors proto.InternalMessageInfo

func (m *BlockDescriptors) GetBD() []BlockDescriptor {
	if m != nil {
		return m.BD
	}
	return nil
}

func init() {
	proto.RegisterType((*BlockDescriptor)(nil), "dymensionxyz.dymension.rollapp.BlockDescriptor")
	proto.RegisterType((*BlockDescriptors)(nil), "dymensionxyz.dymension.rollapp.BlockDescriptors")
}

func init() {
	proto.RegisterFile("dymension/rollapp/block_descriptor.proto", fileDescriptor_9e537259fa860c20)
}

var fileDescriptor_9e537259fa860c20 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x48, 0xa9, 0xcc, 0x4d,
	0xcd, 0x2b, 0xce, 0xcc, 0xcf, 0xd3, 0x2f, 0xca, 0xcf, 0xc9, 0x49, 0x2c, 0x28, 0xd0, 0x4f, 0xca,
	0xc9, 0x4f, 0xce, 0x8e, 0x4f, 0x49, 0x2d, 0x4e, 0x2e, 0xca, 0x2c, 0x28, 0xc9, 0x2f, 0xd2, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x83, 0xab, 0xac, 0xa8, 0xac, 0xd2, 0x83, 0x73, 0xf4, 0xa0,
	0xda, 0xa4, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4a, 0xf5, 0x41, 0x2c, 0x88, 0x2e, 0xa5, 0x46,
	0x46, 0x2e, 0x7e, 0x27, 0x90, 0x81, 0x2e, 0x70, 0xf3, 0x84, 0xc4, 0xb8, 0xd8, 0x32, 0x52, 0x33,
	0xd3, 0x33, 0x4a, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0xa0, 0x3c, 0x21, 0x19, 0x2e, 0xce,
	0xe2, 0x92, 0xc4, 0x92, 0xd4, 0xa0, 0xfc, 0xfc, 0x12, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x9e, 0x20,
	0x84, 0x80, 0x90, 0x05, 0x97, 0x78, 0x66, 0x5e, 0x49, 0x6a, 0x51, 0x6e, 0x6a, 0x4a, 0x66, 0x62,
	0x49, 0x6a, 0x30, 0x48, 0xa2, 0x18, 0x24, 0x53, 0x2c, 0xc1, 0xac, 0xc0, 0xac, 0xc1, 0x13, 0x84,
	0x4b, 0x5a, 0x29, 0x92, 0x4b, 0x00, 0xcd, 0x09, 0xc5, 0x42, 0xae, 0x5c, 0x4c, 0x4e, 0x2e, 0x12,
	0x8c, 0x0a, 0xcc, 0x1a, 0xdc, 0x46, 0xfa, 0x7a, 0xf8, 0xbd, 0xa6, 0x87, 0xa6, 0xdb, 0x89, 0xe5,
	0xc4, 0x3d, 0x79, 0x86, 0x20, 0x26, 0x27, 0x17, 0x27, 0xbf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c,
	0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e,
	0x3c, 0x96, 0x63, 0x88, 0x32, 0x49, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5,
	0x47, 0x36, 0x1e, 0xc1, 0xd1, 0x2f, 0x33, 0xd6, 0xaf, 0x80, 0x87, 0x7a, 0x49, 0x65, 0x41, 0x6a,
	0x71, 0x12, 0x1b, 0x38, 0xd4, 0x8c, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x34, 0x33, 0x6a, 0x74,
	0x97, 0x01, 0x00, 0x00,
}

func (m *BlockDescriptor) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockDescriptor) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockDescriptor) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IntermediateStatesRoots) > 0 {
		for iNdEx := len(m.IntermediateStatesRoots) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.IntermediateStatesRoots[iNdEx])
			copy(dAtA[i:], m.IntermediateStatesRoots[iNdEx])
			i = encodeVarintBlockDescriptor(dAtA, i, uint64(len(m.IntermediateStatesRoots[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.StateRoot) > 0 {
		i -= len(m.StateRoot)
		copy(dAtA[i:], m.StateRoot)
		i = encodeVarintBlockDescriptor(dAtA, i, uint64(len(m.StateRoot)))
		i--
		dAtA[i] = 0x12
	}
	if m.Height != 0 {
		i = encodeVarintBlockDescriptor(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BlockDescriptors) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockDescriptors) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockDescriptors) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.BD) > 0 {
		for iNdEx := len(m.BD) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.BD[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintBlockDescriptor(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintBlockDescriptor(dAtA []byte, offset int, v uint64) int {
	offset -= sovBlockDescriptor(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockDescriptor) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Height != 0 {
		n += 1 + sovBlockDescriptor(uint64(m.Height))
	}
	l = len(m.StateRoot)
	if l > 0 {
		n += 1 + l + sovBlockDescriptor(uint64(l))
	}
	if len(m.IntermediateStatesRoots) > 0 {
		for _, b := range m.IntermediateStatesRoots {
			l = len(b)
			n += 1 + l + sovBlockDescriptor(uint64(l))
		}
	}
	return n
}

func (m *BlockDescriptors) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BD) > 0 {
		for _, e := range m.BD {
			l = e.Size()
			n += 1 + l + sovBlockDescriptor(uint64(l))
		}
	}
	return n
}

func sovBlockDescriptor(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBlockDescriptor(x uint64) (n int) {
	return sovBlockDescriptor(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockDescriptor) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockDescriptor
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
			return fmt.Errorf("proto: BlockDescriptor: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockDescriptor: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockDescriptor
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StateRoot", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockDescriptor
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
				return ErrInvalidLengthBlockDescriptor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.StateRoot = append(m.StateRoot[:0], dAtA[iNdEx:postIndex]...)
			if m.StateRoot == nil {
				m.StateRoot = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntermediateStatesRoots", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockDescriptor
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
				return ErrInvalidLengthBlockDescriptor
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBlockDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IntermediateStatesRoots = append(m.IntermediateStatesRoots, make([]byte, postIndex-iNdEx))
			copy(m.IntermediateStatesRoots[len(m.IntermediateStatesRoots)-1], dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockDescriptor
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
func (m *BlockDescriptors) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBlockDescriptor
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
			return fmt.Errorf("proto: BlockDescriptors: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockDescriptors: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BD", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBlockDescriptor
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
				return ErrInvalidLengthBlockDescriptor
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBlockDescriptor
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BD = append(m.BD, BlockDescriptor{})
			if err := m.BD[len(m.BD)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBlockDescriptor(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBlockDescriptor
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
func skipBlockDescriptor(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBlockDescriptor
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
					return 0, ErrIntOverflowBlockDescriptor
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
					return 0, ErrIntOverflowBlockDescriptor
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
				return 0, ErrInvalidLengthBlockDescriptor
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBlockDescriptor
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBlockDescriptor
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBlockDescriptor        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBlockDescriptor          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBlockDescriptor = fmt.Errorf("proto: unexpected end of group")
)
