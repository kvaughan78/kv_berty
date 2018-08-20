// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/p2p/envelope.proto

/*
	Package p2p is a generated protocol buffer package.

	It is generated from these files:
		api/p2p/envelope.proto
		api/p2p/event.proto
		api/p2p/kind.proto
		api/p2p/service.proto

	It has these top-level messages:
		Envelope
		Event
		SentAttrs
		AckAttrs
		PingAttrs
		ContactRequestAttrs
		ContactRequestAcceptedAttrs
		ContactShareMeAttrs
		ContactShareAttrs
		ConversationInviteAttrs
		ConversationNewMessageAttrs
		Void
*/
package p2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Envelope struct {
	ChannelID       string `protobuf:"bytes,1,opt,name=channel_id,json=channelId,proto3" json:"channel_id,omitempty"`
	EncryptedEvent  []byte `protobuf:"bytes,2,opt,name=encrypted_event,json=encryptedEvent,proto3" json:"encrypted_event,omitempty"`
	SignerPublicKey []byte `protobuf:"bytes,3,opt,name=signer_public_key,json=signerPublicKey,proto3" json:"signer_public_key,omitempty"`
}

func (m *Envelope) Reset()                    { *m = Envelope{} }
func (m *Envelope) String() string            { return proto.CompactTextString(m) }
func (*Envelope) ProtoMessage()               {}
func (*Envelope) Descriptor() ([]byte, []int) { return fileDescriptorEnvelope, []int{0} }

func (m *Envelope) GetChannelID() string {
	if m != nil {
		return m.ChannelID
	}
	return ""
}

func (m *Envelope) GetEncryptedEvent() []byte {
	if m != nil {
		return m.EncryptedEvent
	}
	return nil
}

func (m *Envelope) GetSignerPublicKey() []byte {
	if m != nil {
		return m.SignerPublicKey
	}
	return nil
}

func init() {
	proto.RegisterType((*Envelope)(nil), "berty.p2p.Envelope")
}
func (m *Envelope) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Envelope) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ChannelID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintEnvelope(dAtA, i, uint64(len(m.ChannelID)))
		i += copy(dAtA[i:], m.ChannelID)
	}
	if len(m.EncryptedEvent) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintEnvelope(dAtA, i, uint64(len(m.EncryptedEvent)))
		i += copy(dAtA[i:], m.EncryptedEvent)
	}
	if len(m.SignerPublicKey) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintEnvelope(dAtA, i, uint64(len(m.SignerPublicKey)))
		i += copy(dAtA[i:], m.SignerPublicKey)
	}
	return i, nil
}

func encodeVarintEnvelope(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Envelope) Size() (n int) {
	var l int
	_ = l
	l = len(m.ChannelID)
	if l > 0 {
		n += 1 + l + sovEnvelope(uint64(l))
	}
	l = len(m.EncryptedEvent)
	if l > 0 {
		n += 1 + l + sovEnvelope(uint64(l))
	}
	l = len(m.SignerPublicKey)
	if l > 0 {
		n += 1 + l + sovEnvelope(uint64(l))
	}
	return n
}

func sovEnvelope(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozEnvelope(x uint64) (n int) {
	return sovEnvelope(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Envelope) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEnvelope
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
			return fmt.Errorf("proto: Envelope: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Envelope: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChannelID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ChannelID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptedEvent", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptedEvent = append(m.EncryptedEvent[:0], dAtA[iNdEx:postIndex]...)
			if m.EncryptedEvent == nil {
				m.EncryptedEvent = []byte{}
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignerPublicKey", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEnvelope
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthEnvelope
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignerPublicKey = append(m.SignerPublicKey[:0], dAtA[iNdEx:postIndex]...)
			if m.SignerPublicKey == nil {
				m.SignerPublicKey = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEnvelope(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthEnvelope
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
func skipEnvelope(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEnvelope
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
					return 0, ErrIntOverflowEnvelope
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
					return 0, ErrIntOverflowEnvelope
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
				return 0, ErrInvalidLengthEnvelope
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowEnvelope
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
				next, err := skipEnvelope(dAtA[start:])
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
	ErrInvalidLengthEnvelope = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEnvelope   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/p2p/envelope.proto", fileDescriptorEnvelope) }

var fileDescriptorEnvelope = []byte{
	// 231 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4b, 0x2c, 0xc8, 0xd4,
	0x2f, 0x30, 0x2a, 0xd0, 0x4f, 0xcd, 0x2b, 0x4b, 0xcd, 0xc9, 0x2f, 0x48, 0xd5, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0xe2, 0x4c, 0x4a, 0x2d, 0x2a, 0xa9, 0xd4, 0x2b, 0x30, 0x2a, 0x90, 0x12, 0x49,
	0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xea, 0x83, 0x58, 0x10, 0x05, 0x4a, 0xbd, 0x8c, 0x5c, 0x1c, 0xae,
	0x50, 0x3d, 0x42, 0x3a, 0x5c, 0x5c, 0xc9, 0x19, 0x89, 0x79, 0x79, 0xa9, 0x39, 0xf1, 0x99, 0x29,
	0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x9c, 0x4e, 0xbc, 0x8f, 0xee, 0xc9, 0x73, 0x3a, 0x43, 0x44, 0x3d,
	0x5d, 0x82, 0x38, 0xa1, 0x0a, 0x3c, 0x53, 0x84, 0xd4, 0xb9, 0xf8, 0x53, 0xf3, 0x92, 0x8b, 0x2a,
	0x0b, 0x4a, 0x52, 0x53, 0xe2, 0x53, 0xcb, 0x52, 0xf3, 0x4a, 0x24, 0x98, 0x14, 0x18, 0x35, 0x78,
	0x82, 0xf8, 0xe0, 0xc2, 0xae, 0x20, 0x51, 0x21, 0x2d, 0x2e, 0xc1, 0xe2, 0xcc, 0xf4, 0xbc, 0xd4,
	0xa2, 0xf8, 0x82, 0xd2, 0xa4, 0x9c, 0xcc, 0xe4, 0xf8, 0xec, 0xd4, 0x4a, 0x09, 0x66, 0xb0, 0x52,
	0x7e, 0x88, 0x44, 0x00, 0x58, 0xdc, 0x3b, 0xb5, 0xd2, 0xc9, 0xf4, 0xc4, 0x23, 0x39, 0xc6, 0x0b,
	0x8f, 0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x39, 0x3d, 0xb3,
	0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0xec, 0x1b, 0x28, 0x99, 0x9c, 0x5f, 0x94,
	0xaa, 0x0f, 0xf5, 0x76, 0x12, 0x1b, 0xd8, 0x37, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x71,
	0xb6, 0x7e, 0xc9, 0x08, 0x01, 0x00, 0x00,
}
