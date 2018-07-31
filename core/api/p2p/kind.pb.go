// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/p2p/kind.proto

package p2p

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import berty_entity1 "github.com/berty/berty/core/entity"
import _ "github.com/gogo/protobuf/gogoproto"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Kind int32

const (
	Kind_Unknown Kind = 0
	// Sent events are used by the UI.
	Kind_Sent Kind = 101
	// Ack events are created and sent after receiving an event from a peer.
	Kind_Ack Kind = 102
	// Ping events can be use to check and measure the availability of a peer.
	Kind_Ping                   Kind = 103
	Kind_ContactRequest         Kind = 201
	Kind_ContactRequestAccepted Kind = 202
	Kind_ContactShareMe         Kind = 203
	Kind_ContactShare           Kind = 204
)

var Kind_name = map[int32]string{
	0:   "Unknown",
	101: "Sent",
	102: "Ack",
	103: "Ping",
	201: "ContactRequest",
	202: "ContactRequestAccepted",
	203: "ContactShareMe",
	204: "ContactShare",
}
var Kind_value = map[string]int32{
	"Unknown":                0,
	"Sent":                   101,
	"Ack":                    102,
	"Ping":                   103,
	"ContactRequest":         201,
	"ContactRequestAccepted": 202,
	"ContactShareMe":         203,
	"ContactShare":           204,
}

func (x Kind) String() string {
	return proto.EnumName(Kind_name, int32(x))
}
func (Kind) EnumDescriptor() ([]byte, []int) { return fileDescriptorKind, []int{0} }

type SentAttrs struct {
	IDs []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (m *SentAttrs) Reset()                    { *m = SentAttrs{} }
func (m *SentAttrs) String() string            { return proto.CompactTextString(m) }
func (*SentAttrs) ProtoMessage()               {}
func (*SentAttrs) Descriptor() ([]byte, []int) { return fileDescriptorKind, []int{0} }

func (m *SentAttrs) GetIDs() []string {
	if m != nil {
		return m.IDs
	}
	return nil
}

type AckAttrs struct {
	IDs    []string `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	ErrMsg string   `protobuf:"bytes,2,opt,name=ErrMsg,proto3" json:"ErrMsg,omitempty"`
}

func (m *AckAttrs) Reset()                    { *m = AckAttrs{} }
func (m *AckAttrs) String() string            { return proto.CompactTextString(m) }
func (*AckAttrs) ProtoMessage()               {}
func (*AckAttrs) Descriptor() ([]byte, []int) { return fileDescriptorKind, []int{1} }

func (m *AckAttrs) GetIDs() []string {
	if m != nil {
		return m.IDs
	}
	return nil
}

func (m *AckAttrs) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

type PingAttrs struct {
}

func (m *PingAttrs) Reset()                    { *m = PingAttrs{} }
func (m *PingAttrs) String() string            { return proto.CompactTextString(m) }
func (*PingAttrs) ProtoMessage()               {}
func (*PingAttrs) Descriptor() ([]byte, []int) { return fileDescriptorKind, []int{2} }

type ContactRequestAttrs struct {
	Me           *berty_entity1.Contact `protobuf:"bytes,1,opt,name=me" json:"me,omitempty"`
	IntroMessage string                 `protobuf:"bytes,2,opt,name=intro_message,json=introMessage,proto3" json:"intro_message,omitempty"`
}

func (m *ContactRequestAttrs) Reset()                    { *m = ContactRequestAttrs{} }
func (m *ContactRequestAttrs) String() string            { return proto.CompactTextString(m) }
func (*ContactRequestAttrs) ProtoMessage()               {}
func (*ContactRequestAttrs) Descriptor() ([]byte, []int) { return fileDescriptorKind, []int{3} }

func (m *ContactRequestAttrs) GetMe() *berty_entity1.Contact {
	if m != nil {
		return m.Me
	}
	return nil
}

func (m *ContactRequestAttrs) GetIntroMessage() string {
	if m != nil {
		return m.IntroMessage
	}
	return ""
}

type ContactRequestAcceptedAttrs struct {
}

func (m *ContactRequestAcceptedAttrs) Reset()                    { *m = ContactRequestAcceptedAttrs{} }
func (m *ContactRequestAcceptedAttrs) String() string            { return proto.CompactTextString(m) }
func (*ContactRequestAcceptedAttrs) ProtoMessage()               {}
func (*ContactRequestAcceptedAttrs) Descriptor() ([]byte, []int) { return fileDescriptorKind, []int{4} }

type ContactShareMeAttrs struct {
	Me *berty_entity1.Contact `protobuf:"bytes,1,opt,name=me" json:"me,omitempty"`
}

func (m *ContactShareMeAttrs) Reset()                    { *m = ContactShareMeAttrs{} }
func (m *ContactShareMeAttrs) String() string            { return proto.CompactTextString(m) }
func (*ContactShareMeAttrs) ProtoMessage()               {}
func (*ContactShareMeAttrs) Descriptor() ([]byte, []int) { return fileDescriptorKind, []int{5} }

func (m *ContactShareMeAttrs) GetMe() *berty_entity1.Contact {
	if m != nil {
		return m.Me
	}
	return nil
}

type ContactShareAttrs struct {
	Contact *berty_entity1.Contact `protobuf:"bytes,1,opt,name=contact" json:"contact,omitempty"`
}

func (m *ContactShareAttrs) Reset()                    { *m = ContactShareAttrs{} }
func (m *ContactShareAttrs) String() string            { return proto.CompactTextString(m) }
func (*ContactShareAttrs) ProtoMessage()               {}
func (*ContactShareAttrs) Descriptor() ([]byte, []int) { return fileDescriptorKind, []int{6} }

func (m *ContactShareAttrs) GetContact() *berty_entity1.Contact {
	if m != nil {
		return m.Contact
	}
	return nil
}

func init() {
	proto.RegisterType((*SentAttrs)(nil), "berty.p2p.SentAttrs")
	proto.RegisterType((*AckAttrs)(nil), "berty.p2p.AckAttrs")
	proto.RegisterType((*PingAttrs)(nil), "berty.p2p.PingAttrs")
	proto.RegisterType((*ContactRequestAttrs)(nil), "berty.p2p.ContactRequestAttrs")
	proto.RegisterType((*ContactRequestAcceptedAttrs)(nil), "berty.p2p.ContactRequestAcceptedAttrs")
	proto.RegisterType((*ContactShareMeAttrs)(nil), "berty.p2p.ContactShareMeAttrs")
	proto.RegisterType((*ContactShareAttrs)(nil), "berty.p2p.ContactShareAttrs")
	proto.RegisterEnum("berty.p2p.Kind", Kind_name, Kind_value)
}
func (m *SentAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SentAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, s := range m.IDs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	return i, nil
}

func (m *AckAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AckAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, s := range m.IDs {
			dAtA[i] = 0xa
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if len(m.ErrMsg) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKind(dAtA, i, uint64(len(m.ErrMsg)))
		i += copy(dAtA[i:], m.ErrMsg)
	}
	return i, nil
}

func (m *PingAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PingAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ContactRequestAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactRequestAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Me != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(m.Me.Size()))
		n1, err := m.Me.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	if len(m.IntroMessage) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintKind(dAtA, i, uint64(len(m.IntroMessage)))
		i += copy(dAtA[i:], m.IntroMessage)
	}
	return i, nil
}

func (m *ContactRequestAcceptedAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactRequestAcceptedAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *ContactShareMeAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactShareMeAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Me != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(m.Me.Size()))
		n2, err := m.Me.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *ContactShareAttrs) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContactShareAttrs) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Contact != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintKind(dAtA, i, uint64(m.Contact.Size()))
		n3, err := m.Contact.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n3
	}
	return i, nil
}

func encodeVarintKind(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *SentAttrs) Size() (n int) {
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, s := range m.IDs {
			l = len(s)
			n += 1 + l + sovKind(uint64(l))
		}
	}
	return n
}

func (m *AckAttrs) Size() (n int) {
	var l int
	_ = l
	if len(m.IDs) > 0 {
		for _, s := range m.IDs {
			l = len(s)
			n += 1 + l + sovKind(uint64(l))
		}
	}
	l = len(m.ErrMsg)
	if l > 0 {
		n += 1 + l + sovKind(uint64(l))
	}
	return n
}

func (m *PingAttrs) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ContactRequestAttrs) Size() (n int) {
	var l int
	_ = l
	if m.Me != nil {
		l = m.Me.Size()
		n += 1 + l + sovKind(uint64(l))
	}
	l = len(m.IntroMessage)
	if l > 0 {
		n += 1 + l + sovKind(uint64(l))
	}
	return n
}

func (m *ContactRequestAcceptedAttrs) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *ContactShareMeAttrs) Size() (n int) {
	var l int
	_ = l
	if m.Me != nil {
		l = m.Me.Size()
		n += 1 + l + sovKind(uint64(l))
	}
	return n
}

func (m *ContactShareAttrs) Size() (n int) {
	var l int
	_ = l
	if m.Contact != nil {
		l = m.Contact.Size()
		n += 1 + l + sovKind(uint64(l))
	}
	return n
}

func sovKind(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozKind(x uint64) (n int) {
	return sovKind(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SentAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
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
			return fmt.Errorf("proto: SentAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SentAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
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
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDs = append(m.IDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
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
func (m *AckAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
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
			return fmt.Errorf("proto: AckAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AckAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDs", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
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
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDs = append(m.IDs, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ErrMsg", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
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
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ErrMsg = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
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
func (m *PingAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
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
			return fmt.Errorf("proto: PingAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PingAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
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
func (m *ContactRequestAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
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
			return fmt.Errorf("proto: ContactRequestAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactRequestAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Me", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
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
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Me == nil {
				m.Me = &berty_entity1.Contact{}
			}
			if err := m.Me.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IntroMessage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
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
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IntroMessage = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
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
func (m *ContactRequestAcceptedAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
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
			return fmt.Errorf("proto: ContactRequestAcceptedAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactRequestAcceptedAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
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
func (m *ContactShareMeAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
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
			return fmt.Errorf("proto: ContactShareMeAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactShareMeAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Me", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
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
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Me == nil {
				m.Me = &berty_entity1.Contact{}
			}
			if err := m.Me.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
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
func (m *ContactShareAttrs) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowKind
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
			return fmt.Errorf("proto: ContactShareAttrs: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContactShareAttrs: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contact", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowKind
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
				return ErrInvalidLengthKind
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Contact == nil {
				m.Contact = &berty_entity1.Contact{}
			}
			if err := m.Contact.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipKind(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthKind
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
func skipKind(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowKind
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
					return 0, ErrIntOverflowKind
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
					return 0, ErrIntOverflowKind
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
				return 0, ErrInvalidLengthKind
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowKind
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
				next, err := skipKind(dAtA[start:])
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
	ErrInvalidLengthKind = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowKind   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("api/p2p/kind.proto", fileDescriptorKind) }

var fileDescriptorKind = []byte{
	// 395 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0xcd, 0xca, 0xd3, 0x40,
	0x14, 0xfd, 0xa6, 0xf9, 0x68, 0x9a, 0xdb, 0x2a, 0xd3, 0x69, 0x2d, 0xb5, 0xc5, 0x58, 0x52, 0x94,
	0xe2, 0x22, 0x81, 0x8a, 0x3b, 0x5d, 0xb4, 0xd6, 0x85, 0x48, 0x41, 0x5a, 0xdc, 0xb8, 0x91, 0x74,
	0x32, 0xa6, 0x43, 0xc8, 0x4c, 0x9c, 0x4c, 0x91, 0xbe, 0x82, 0x4f, 0xe0, 0x23, 0xf9, 0xb7, 0xf0,
	0x09, 0x44, 0xe2, 0x8b, 0x48, 0x32, 0xa9, 0x54, 0x11, 0xe5, 0xdb, 0x0c, 0x77, 0x0e, 0xe7, 0xcc,
	0x3d, 0xe7, 0xce, 0x05, 0x12, 0x66, 0x3c, 0xc8, 0xe6, 0x59, 0x90, 0x70, 0x11, 0xf9, 0x99, 0x92,
	0x5a, 0x12, 0x67, 0xc7, 0x94, 0x3e, 0xfa, 0xd9, 0x3c, 0x1b, 0xf5, 0x99, 0xd0, 0x5c, 0x1f, 0x03,
	0x2a, 0x85, 0x0e, 0xa9, 0x36, 0x84, 0x51, 0x3f, 0x96, 0xb1, 0xac, 0xca, 0xa0, 0xac, 0x0c, 0xea,
	0xdd, 0x05, 0x67, 0xcb, 0x84, 0x5e, 0x68, 0xad, 0x72, 0x72, 0x13, 0x2c, 0x1e, 0xe5, 0x43, 0x34,
	0xb1, 0x66, 0xce, 0xd2, 0x2e, 0xbe, 0xdd, 0xb6, 0x9e, 0xae, 0xf2, 0x4d, 0x89, 0x79, 0x8f, 0xa0,
	0xb5, 0xa0, 0xc9, 0xff, 0x68, 0x64, 0x00, 0xcd, 0x27, 0x4a, 0xad, 0xf3, 0x78, 0xd8, 0x98, 0xa0,
	0x99, 0xb3, 0xa9, 0x6f, 0x5e, 0x1b, 0x9c, 0xe7, 0x5c, 0xc4, 0x95, 0xde, 0x0b, 0xa1, 0xf7, 0xd8,
	0x58, 0xdb, 0xb0, 0x37, 0x07, 0x96, 0xd7, 0xdd, 0xef, 0x40, 0x23, 0x65, 0x43, 0x34, 0x41, 0xb3,
	0xf6, 0xfc, 0x86, 0x6f, 0xe2, 0x98, 0x24, 0xfe, 0x89, 0xde, 0x48, 0x19, 0x99, 0xc2, 0x35, 0x2e,
	0xb4, 0x92, 0xaf, 0x52, 0x96, 0xe7, 0x61, 0xcc, 0xea, 0x4e, 0x9d, 0x0a, 0x5c, 0x1b, 0xcc, 0xbb,
	0x05, 0xe3, 0x3f, 0x5a, 0x50, 0xca, 0x32, 0xcd, 0x22, 0xe3, 0xe0, 0xe1, 0x2f, 0x07, 0xdb, 0x7d,
	0xa8, 0xd8, 0x9a, 0x5d, 0xc5, 0x81, 0xb7, 0x82, 0xee, 0xb9, 0xda, 0x68, 0x03, 0xb0, 0xeb, 0x79,
	0xff, 0xfb, 0x81, 0x13, 0xeb, 0xde, 0x3b, 0x04, 0x97, 0xcf, 0xb8, 0x88, 0x48, 0x1b, 0xec, 0x17,
	0x22, 0x11, 0xf2, 0xad, 0xc0, 0x17, 0xa4, 0x05, 0x97, 0xe5, 0x7f, 0x60, 0x46, 0x6c, 0xb0, 0x16,
	0x34, 0xc1, 0xaf, 0x4b, 0xa8, 0x9c, 0x1d, 0x8e, 0x49, 0x0f, 0xae, 0xff, 0x9e, 0x0a, 0x7f, 0x44,
	0x64, 0x0c, 0x83, 0xbf, 0x47, 0xc5, 0x9f, 0xd0, 0x99, 0xa2, 0x0e, 0x8a, 0x3f, 0x23, 0xd2, 0x85,
	0xce, 0x39, 0x88, 0xbf, 0xa0, 0xe5, 0x83, 0x0f, 0x85, 0x8b, 0xbe, 0x16, 0x2e, 0xfa, 0x5e, 0xb8,
	0xe8, 0xfd, 0x0f, 0xf7, 0xe2, 0xe5, 0x34, 0xe6, 0x7a, 0x7f, 0xd8, 0xf9, 0x54, 0xa6, 0x41, 0x15,
	0xa4, 0x3e, 0xa9, 0x54, 0x2c, 0xa8, 0xf7, 0x6f, 0xd7, 0xac, 0x96, 0xe8, 0xfe, 0xcf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xf0, 0x5c, 0x30, 0x86, 0x91, 0x02, 0x00, 0x00,
}