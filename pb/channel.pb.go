// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: channel.proto

package pb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type SetChannelRequest struct {
}

func (m *SetChannelRequest) Reset()         { *m = SetChannelRequest{} }
func (m *SetChannelRequest) String() string { return proto.CompactTextString(m) }
func (*SetChannelRequest) ProtoMessage()    {}
func (*SetChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_channel_ca5238ae9de34f8c, []int{0}
}
func (m *SetChannelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetChannelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetChannelRequest.Merge(dst, src)
}
func (m *SetChannelRequest) XXX_Size() int {
	return m.Size()
}
func (m *SetChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetChannelRequest proto.InternalMessageInfo

type SetChannelResponse struct {
}

func (m *SetChannelResponse) Reset()         { *m = SetChannelResponse{} }
func (m *SetChannelResponse) String() string { return proto.CompactTextString(m) }
func (*SetChannelResponse) ProtoMessage()    {}
func (*SetChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_channel_ca5238ae9de34f8c, []int{1}
}
func (m *SetChannelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SetChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SetChannelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *SetChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetChannelResponse.Merge(dst, src)
}
func (m *SetChannelResponse) XXX_Size() int {
	return m.Size()
}
func (m *SetChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetChannelResponse proto.InternalMessageInfo

type GetChannelRequest struct {
}

func (m *GetChannelRequest) Reset()         { *m = GetChannelRequest{} }
func (m *GetChannelRequest) String() string { return proto.CompactTextString(m) }
func (*GetChannelRequest) ProtoMessage()    {}
func (*GetChannelRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_channel_ca5238ae9de34f8c, []int{2}
}
func (m *GetChannelRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChannelRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChannelRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetChannelRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChannelRequest.Merge(dst, src)
}
func (m *GetChannelRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetChannelRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChannelRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetChannelRequest proto.InternalMessageInfo

type GetChannelResponse struct {
}

func (m *GetChannelResponse) Reset()         { *m = GetChannelResponse{} }
func (m *GetChannelResponse) String() string { return proto.CompactTextString(m) }
func (*GetChannelResponse) ProtoMessage()    {}
func (*GetChannelResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_channel_ca5238ae9de34f8c, []int{3}
}
func (m *GetChannelResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetChannelResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetChannelResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *GetChannelResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetChannelResponse.Merge(dst, src)
}
func (m *GetChannelResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetChannelResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetChannelResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetChannelResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*SetChannelRequest)(nil), "state.SetChannelRequest")
	proto.RegisterType((*SetChannelResponse)(nil), "state.SetChannelResponse")
	proto.RegisterType((*GetChannelRequest)(nil), "state.GetChannelRequest")
	proto.RegisterType((*GetChannelResponse)(nil), "state.GetChannelResponse")
}
func (m *SetChannelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetChannelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *SetChannelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SetChannelResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetChannelRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChannelRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *GetChannelResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetChannelResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintChannel(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SetChannelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *SetChannelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetChannelRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *GetChannelResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovChannel(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozChannel(x uint64) (n int) {
	return sovChannel(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SetChannelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChannel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetChannelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetChannelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipChannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChannel
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
func (m *SetChannelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChannel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: SetChannelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SetChannelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipChannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChannel
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
func (m *GetChannelRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChannel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetChannelRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChannelRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipChannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChannel
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
func (m *GetChannelResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowChannel
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: GetChannelResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetChannelResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipChannel(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthChannel
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
func skipChannel(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowChannel
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
					return 0, ErrIntOverflowChannel
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowChannel
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthChannel
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowChannel
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipChannel(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthChannel = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowChannel   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("channel.proto", fileDescriptor_channel_ca5238ae9de34f8c) }

var fileDescriptor_channel_ca5238ae9de34f8c = []byte{
	// 119 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4d, 0xce, 0x48, 0xcc,
	0xcb, 0x4b, 0xcd, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2d, 0x2e, 0x49, 0x2c, 0x49,
	0x55, 0x12, 0xe6, 0x12, 0x0c, 0x4e, 0x2d, 0x71, 0x86, 0x48, 0x05, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x28, 0x89, 0x70, 0x09, 0x21, 0x0b, 0x16, 0x17, 0xe4, 0xe7, 0x15, 0x83, 0x95, 0xba, 0x63,
	0x53, 0xea, 0x8e, 0xa1, 0xd4, 0x49, 0xe6, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f,
	0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18,
	0xa2, 0x98, 0x0a, 0x92, 0x92, 0xd8, 0xc0, 0x2e, 0x30, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0x06,
	0xf9, 0xf7, 0x4e, 0x92, 0x00, 0x00, 0x00,
}
