// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: entity/contact.proto

package entity

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"
import _ "github.com/golang/protobuf/ptypes/timestamp"

import time "time"

import types "github.com/gogo/protobuf/types"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

type Contact_Status int32

const (
	Contact_Unknown         Contact_Status = 0
	Contact_IsFriend        Contact_Status = 1
	Contact_IsTrustedFriend Contact_Status = 2
	Contact_IsRequested     Contact_Status = 3
	Contact_RequestedMe     Contact_Status = 4
	Contact_IsBlocked       Contact_Status = 5
	Contact_Myself          Contact_Status = 42
)

var Contact_Status_name = map[int32]string{
	0:  "Unknown",
	1:  "IsFriend",
	2:  "IsTrustedFriend",
	3:  "IsRequested",
	4:  "RequestedMe",
	5:  "IsBlocked",
	42: "Myself",
}
var Contact_Status_value = map[string]int32{
	"Unknown":         0,
	"IsFriend":        1,
	"IsTrustedFriend": 2,
	"IsRequested":     3,
	"RequestedMe":     4,
	"IsBlocked":       5,
	"Myself":          42,
}

func (x Contact_Status) String() string {
	return proto.EnumName(Contact_Status_name, int32(x))
}
func (Contact_Status) EnumDescriptor() ([]byte, []int) { return fileDescriptorContact, []int{0, 0} }

type Contact struct {
	ID                    string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" gorm:"primary_key"`
	CreatedAt             time.Time      `protobuf:"bytes,2,opt,name=created_at,json=createdAt,stdtime" json:"created_at"`
	UpdatedAt             time.Time      `protobuf:"bytes,3,opt,name=updated_at,json=updatedAt,stdtime" json:"updated_at"`
	DeletedAt             *time.Time     `protobuf:"bytes,4,opt,name=deleted_at,json=deletedAt,stdtime" json:"deleted_at,omitempty"`
	Sigchain              []byte         `protobuf:"bytes,10,opt,name=sigchain,proto3" json:"sigchain,omitempty"`
	Status                Contact_Status `protobuf:"varint,11,opt,name=status,proto3,enum=berty.entity.Contact_Status" json:"status,omitempty"`
	Devices               []*Device      `protobuf:"bytes,12,rep,name=devices" json:"devices,omitempty"`
	DisplayName           string         `protobuf:"bytes,13,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	DisplayStatus         string         `protobuf:"bytes,14,opt,name=display_status,json=displayStatus,proto3" json:"display_status,omitempty"`
	OverrideDisplayName   string         `protobuf:"bytes,15,opt,name=override_display_name,json=overrideDisplayName,proto3" json:"override_display_name,omitempty"`
	OverrideDisplayStatus string         `protobuf:"bytes,16,opt,name=override_display_status,json=overrideDisplayStatus,proto3" json:"override_display_status,omitempty"`
}

func (m *Contact) Reset()                    { *m = Contact{} }
func (m *Contact) String() string            { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()               {}
func (*Contact) Descriptor() ([]byte, []int) { return fileDescriptorContact, []int{0} }

func (m *Contact) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Contact) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func (m *Contact) GetUpdatedAt() time.Time {
	if m != nil {
		return m.UpdatedAt
	}
	return time.Time{}
}

func (m *Contact) GetDeletedAt() *time.Time {
	if m != nil {
		return m.DeletedAt
	}
	return nil
}

func (m *Contact) GetSigchain() []byte {
	if m != nil {
		return m.Sigchain
	}
	return nil
}

func (m *Contact) GetStatus() Contact_Status {
	if m != nil {
		return m.Status
	}
	return Contact_Unknown
}

func (m *Contact) GetDevices() []*Device {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *Contact) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Contact) GetDisplayStatus() string {
	if m != nil {
		return m.DisplayStatus
	}
	return ""
}

func (m *Contact) GetOverrideDisplayName() string {
	if m != nil {
		return m.OverrideDisplayName
	}
	return ""
}

func (m *Contact) GetOverrideDisplayStatus() string {
	if m != nil {
		return m.OverrideDisplayStatus
	}
	return ""
}

func init() {
	proto.RegisterType((*Contact)(nil), "berty.entity.Contact")
	proto.RegisterEnum("berty.entity.Contact_Status", Contact_Status_name, Contact_Status_value)
}
func (m *Contact) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contact) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.ID) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintContact(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintContact(dAtA, i, uint64(types.SizeOfStdTime(m.CreatedAt)))
	n1, err := types.StdTimeMarshalTo(m.CreatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	dAtA[i] = 0x1a
	i++
	i = encodeVarintContact(dAtA, i, uint64(types.SizeOfStdTime(m.UpdatedAt)))
	n2, err := types.StdTimeMarshalTo(m.UpdatedAt, dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	if m.DeletedAt != nil {
		dAtA[i] = 0x22
		i++
		i = encodeVarintContact(dAtA, i, uint64(types.SizeOfStdTime(*m.DeletedAt)))
		n3, err := types.StdTimeMarshalTo(*m.DeletedAt, dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	if len(m.Sigchain) > 0 {
		dAtA[i] = 0x52
		i++
		i = encodeVarintContact(dAtA, i, uint64(len(m.Sigchain)))
		i += copy(dAtA[i:], m.Sigchain)
	}
	if m.Status != 0 {
		dAtA[i] = 0x58
		i++
		i = encodeVarintContact(dAtA, i, uint64(m.Status))
	}
	if len(m.Devices) > 0 {
		for _, msg := range m.Devices {
			dAtA[i] = 0x62
			i++
			i = encodeVarintContact(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.DisplayName) > 0 {
		dAtA[i] = 0x6a
		i++
		i = encodeVarintContact(dAtA, i, uint64(len(m.DisplayName)))
		i += copy(dAtA[i:], m.DisplayName)
	}
	if len(m.DisplayStatus) > 0 {
		dAtA[i] = 0x72
		i++
		i = encodeVarintContact(dAtA, i, uint64(len(m.DisplayStatus)))
		i += copy(dAtA[i:], m.DisplayStatus)
	}
	if len(m.OverrideDisplayName) > 0 {
		dAtA[i] = 0x7a
		i++
		i = encodeVarintContact(dAtA, i, uint64(len(m.OverrideDisplayName)))
		i += copy(dAtA[i:], m.OverrideDisplayName)
	}
	if len(m.OverrideDisplayStatus) > 0 {
		dAtA[i] = 0x82
		i++
		dAtA[i] = 0x1
		i++
		i = encodeVarintContact(dAtA, i, uint64(len(m.OverrideDisplayStatus)))
		i += copy(dAtA[i:], m.OverrideDisplayStatus)
	}
	return i, nil
}

func encodeVarintContact(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *Contact) Size() (n int) {
	var l int
	_ = l
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovContact(uint64(l))
	}
	l = types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovContact(uint64(l))
	l = types.SizeOfStdTime(m.UpdatedAt)
	n += 1 + l + sovContact(uint64(l))
	if m.DeletedAt != nil {
		l = types.SizeOfStdTime(*m.DeletedAt)
		n += 1 + l + sovContact(uint64(l))
	}
	l = len(m.Sigchain)
	if l > 0 {
		n += 1 + l + sovContact(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovContact(uint64(m.Status))
	}
	if len(m.Devices) > 0 {
		for _, e := range m.Devices {
			l = e.Size()
			n += 1 + l + sovContact(uint64(l))
		}
	}
	l = len(m.DisplayName)
	if l > 0 {
		n += 1 + l + sovContact(uint64(l))
	}
	l = len(m.DisplayStatus)
	if l > 0 {
		n += 1 + l + sovContact(uint64(l))
	}
	l = len(m.OverrideDisplayName)
	if l > 0 {
		n += 1 + l + sovContact(uint64(l))
	}
	l = len(m.OverrideDisplayStatus)
	if l > 0 {
		n += 2 + l + sovContact(uint64(l))
	}
	return n
}

func sovContact(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozContact(x uint64) (n int) {
	return sovContact(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Contact) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowContact
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
			return fmt.Errorf("proto: Contact: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contact: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
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
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UpdatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := types.StdTimeUnmarshal(&m.UpdatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeletedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeletedAt == nil {
				m.DeletedAt = new(time.Time)
			}
			if err := types.StdTimeUnmarshal(m.DeletedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sigchain", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
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
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + byteLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sigchain = append(m.Sigchain[:0], dAtA[iNdEx:postIndex]...)
			if m.Sigchain == nil {
				m.Sigchain = []byte{}
			}
			iNdEx = postIndex
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= (Contact_Status(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Devices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Devices = append(m.Devices, &Device{})
			if err := m.Devices[len(m.Devices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
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
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DisplayStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
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
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DisplayStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverrideDisplayName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
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
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OverrideDisplayName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OverrideDisplayStatus", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowContact
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
				return ErrInvalidLengthContact
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OverrideDisplayStatus = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipContact(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthContact
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
func skipContact(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowContact
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
					return 0, ErrIntOverflowContact
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
					return 0, ErrIntOverflowContact
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
				return 0, ErrInvalidLengthContact
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowContact
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
				next, err := skipContact(dAtA[start:])
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
	ErrInvalidLengthContact = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowContact   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("entity/contact.proto", fileDescriptorContact) }

var fileDescriptorContact = []byte{
	// 519 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0x6a, 0xdb, 0x4c,
	0x14, 0x8d, 0x1c, 0xc7, 0x3f, 0x57, 0xfe, 0x11, 0x63, 0x87, 0x4f, 0x98, 0x60, 0xfb, 0x33, 0x14,
	0x4c, 0x29, 0x12, 0xb8, 0xa1, 0x8b, 0x6e, 0x8a, 0x1d, 0x53, 0xf0, 0x22, 0x5d, 0xa8, 0xe9, 0xa6,
	0x1b, 0x23, 0x6b, 0x6e, 0x94, 0xc1, 0x96, 0xc6, 0x9d, 0x19, 0xa5, 0xe8, 0x2d, 0xba, 0xec, 0x6b,
	0xf4, 0x2d, 0xb2, 0xec, 0x13, 0xb8, 0xc5, 0x7d, 0x83, 0x3e, 0x41, 0xb1, 0x34, 0x0a, 0x49, 0xbb,
	0x28, 0xdd, 0x08, 0xdd, 0x73, 0xcf, 0x39, 0xf7, 0x70, 0xef, 0x40, 0x17, 0x63, 0xc5, 0x54, 0xea,
	0x06, 0x3c, 0x56, 0x7e, 0xa0, 0x9c, 0xad, 0xe0, 0x8a, 0x93, 0xc6, 0x0a, 0x85, 0x4a, 0x9d, 0xbc,
	0xd7, 0xeb, 0x86, 0x3c, 0xe4, 0x59, 0xc3, 0x3d, 0xfc, 0xe5, 0x9c, 0xde, 0x20, 0xe4, 0x3c, 0xdc,
	0xa0, 0x9b, 0x55, 0xab, 0xe4, 0xda, 0x55, 0x2c, 0x42, 0xa9, 0xfc, 0x68, 0xab, 0x09, 0x1d, 0x6d,
	0x4d, 0xf1, 0x96, 0x05, 0x98, 0x83, 0xa3, 0x2f, 0x27, 0x50, 0xbd, 0xc8, 0x67, 0x91, 0x67, 0x50,
	0x62, 0xd4, 0x36, 0x86, 0xc6, 0xb8, 0x3e, 0x3b, 0xdb, 0xef, 0x06, 0xa5, 0xc5, 0xfc, 0xe7, 0x6e,
	0x40, 0x42, 0x2e, 0xa2, 0x97, 0xa3, 0xad, 0x60, 0x91, 0x2f, 0xd2, 0xe5, 0x1a, 0xd3, 0x91, 0x57,
	0x62, 0x94, 0x5c, 0x00, 0x04, 0x02, 0x7d, 0x85, 0x74, 0xe9, 0x2b, 0xbb, 0x34, 0x34, 0xc6, 0xe6,
	0xa4, 0xe7, 0xe4, 0x21, 0x9c, 0x22, 0x84, 0x73, 0x55, 0x84, 0x98, 0xd5, 0xee, 0x76, 0x83, 0xa3,
	0x4f, 0xdf, 0x06, 0x86, 0x57, 0xd7, 0xba, 0xa9, 0x3a, 0x98, 0x24, 0x5b, 0x5a, 0x98, 0x1c, 0xff,
	0x8b, 0x89, 0xd6, 0x4d, 0x15, 0x79, 0x05, 0x40, 0x71, 0x83, 0xda, 0xa4, 0xfc, 0x57, 0x93, 0x72,
	0x6e, 0xa0, 0x35, 0x53, 0x45, 0x7a, 0x50, 0x93, 0x2c, 0x0c, 0x6e, 0x7c, 0x16, 0xdb, 0x30, 0x34,
	0xc6, 0x0d, 0xef, 0xbe, 0x26, 0xe7, 0x50, 0x91, 0xca, 0x57, 0x89, 0xb4, 0xcd, 0xa1, 0x31, 0x6e,
	0x4d, 0xce, 0x9c, 0x87, 0xb7, 0x70, 0xf4, 0xee, 0x9c, 0xb7, 0x19, 0xc7, 0xd3, 0x5c, 0xe2, 0x40,
	0x35, 0x5f, 0xb3, 0xb4, 0x1b, 0xc3, 0xe3, 0xb1, 0x39, 0xe9, 0x3e, 0x96, 0xcd, 0xb3, 0xa6, 0x57,
	0x90, 0xc8, 0xff, 0xd0, 0xa0, 0x4c, 0x6e, 0x37, 0x7e, 0xba, 0x8c, 0xfd, 0x08, 0xed, 0xe6, 0xe1,
	0x08, 0x9e, 0xa9, 0xb1, 0x37, 0x7e, 0x84, 0xe4, 0x09, 0xb4, 0x0a, 0x8a, 0x0e, 0xd4, 0xca, 0x48,
	0x4d, 0x8d, 0xe6, 0x09, 0xc8, 0x04, 0x4e, 0xf9, 0x2d, 0x0a, 0xc1, 0x28, 0x2e, 0x1f, 0x59, 0xb6,
	0x33, 0x76, 0xa7, 0x68, 0xce, 0x1f, 0x58, 0xbf, 0x80, 0xff, 0xfe, 0xd0, 0xe8, 0x19, 0x56, 0xa6,
	0x3a, 0xfd, 0x4d, 0x95, 0xcf, 0x1a, 0x25, 0x50, 0xd1, 0x53, 0x4d, 0xa8, 0xbe, 0x8b, 0xd7, 0x31,
	0xff, 0x18, 0x5b, 0x47, 0xa4, 0x01, 0xb5, 0x85, 0x7c, 0x2d, 0x18, 0xc6, 0xd4, 0x32, 0x48, 0x07,
	0xda, 0x0b, 0x79, 0x25, 0x12, 0xa9, 0x90, 0x6a, 0xb0, 0x44, 0xda, 0x60, 0x2e, 0xa4, 0x87, 0x1f,
	0x12, 0x3c, 0xc0, 0xd6, 0xf1, 0x01, 0xb8, 0x2f, 0x2f, 0xd1, 0x2a, 0x93, 0x26, 0xd4, 0x17, 0x72,
	0xb6, 0xe1, 0xc1, 0x1a, 0xa9, 0x75, 0x42, 0x00, 0x2a, 0x97, 0xa9, 0xc4, 0xcd, 0xb5, 0xf5, 0x74,
	0x76, 0x7e, 0xb7, 0xef, 0x1b, 0x5f, 0xf7, 0x7d, 0xe3, 0xfb, 0xbe, 0x6f, 0x7c, 0xfe, 0xd1, 0x3f,
	0x7a, 0x3f, 0x0a, 0x99, 0xba, 0x49, 0x56, 0x4e, 0xc0, 0x23, 0x37, 0xdb, 0xb3, 0xfe, 0x06, 0x5c,
	0xa0, 0x9b, 0xaf, 0x7c, 0x55, 0xc9, 0x5e, 0xc2, 0xf3, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5b,
	0x64, 0xc9, 0xc6, 0x62, 0x03, 0x00, 0x00,
}
