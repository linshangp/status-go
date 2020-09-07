// Code generated by protoc-gen-go. DO NOT EDIT.
// source: application_metadata_message.proto

package protobuf

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ApplicationMetadataMessage_Type int32

const (
	ApplicationMetadataMessage_UNKNOWN                                 ApplicationMetadataMessage_Type = 0
	ApplicationMetadataMessage_CHAT_MESSAGE                            ApplicationMetadataMessage_Type = 1
	ApplicationMetadataMessage_CONTACT_UPDATE                          ApplicationMetadataMessage_Type = 2
	ApplicationMetadataMessage_MEMBERSHIP_UPDATE_MESSAGE               ApplicationMetadataMessage_Type = 3
	ApplicationMetadataMessage_PAIR_INSTALLATION                       ApplicationMetadataMessage_Type = 4
	ApplicationMetadataMessage_SYNC_INSTALLATION                       ApplicationMetadataMessage_Type = 5
	ApplicationMetadataMessage_REQUEST_ADDRESS_FOR_TRANSACTION         ApplicationMetadataMessage_Type = 6
	ApplicationMetadataMessage_ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION  ApplicationMetadataMessage_Type = 7
	ApplicationMetadataMessage_DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION ApplicationMetadataMessage_Type = 8
	ApplicationMetadataMessage_REQUEST_TRANSACTION                     ApplicationMetadataMessage_Type = 9
	ApplicationMetadataMessage_SEND_TRANSACTION                        ApplicationMetadataMessage_Type = 10
	ApplicationMetadataMessage_DECLINE_REQUEST_TRANSACTION             ApplicationMetadataMessage_Type = 11
	ApplicationMetadataMessage_SYNC_INSTALLATION_CONTACT               ApplicationMetadataMessage_Type = 12
	ApplicationMetadataMessage_SYNC_INSTALLATION_ACCOUNT               ApplicationMetadataMessage_Type = 13
	ApplicationMetadataMessage_SYNC_INSTALLATION_PUBLIC_CHAT           ApplicationMetadataMessage_Type = 14
	ApplicationMetadataMessage_CONTACT_CODE_ADVERTISEMENT              ApplicationMetadataMessage_Type = 15
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REGISTRATION          ApplicationMetadataMessage_Type = 16
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REGISTRATION_RESPONSE ApplicationMetadataMessage_Type = 17
	ApplicationMetadataMessage_PUSH_NOTIFICATION_QUERY                 ApplicationMetadataMessage_Type = 18
	ApplicationMetadataMessage_PUSH_NOTIFICATION_QUERY_RESPONSE        ApplicationMetadataMessage_Type = 19
	ApplicationMetadataMessage_PUSH_NOTIFICATION_REQUEST               ApplicationMetadataMessage_Type = 20
	ApplicationMetadataMessage_PUSH_NOTIFICATION_RESPONSE              ApplicationMetadataMessage_Type = 21
	ApplicationMetadataMessage_EMOJI_REACTION                          ApplicationMetadataMessage_Type = 22
	ApplicationMetadataMessage_GROUP_CHAT_INVITATION                   ApplicationMetadataMessage_Type = 23
)

var ApplicationMetadataMessage_Type_name = map[int32]string{
	0:  "UNKNOWN",
	1:  "CHAT_MESSAGE",
	2:  "CONTACT_UPDATE",
	3:  "MEMBERSHIP_UPDATE_MESSAGE",
	4:  "PAIR_INSTALLATION",
	5:  "SYNC_INSTALLATION",
	6:  "REQUEST_ADDRESS_FOR_TRANSACTION",
	7:  "ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION",
	8:  "DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION",
	9:  "REQUEST_TRANSACTION",
	10: "SEND_TRANSACTION",
	11: "DECLINE_REQUEST_TRANSACTION",
	12: "SYNC_INSTALLATION_CONTACT",
	13: "SYNC_INSTALLATION_ACCOUNT",
	14: "SYNC_INSTALLATION_PUBLIC_CHAT",
	15: "CONTACT_CODE_ADVERTISEMENT",
	16: "PUSH_NOTIFICATION_REGISTRATION",
	17: "PUSH_NOTIFICATION_REGISTRATION_RESPONSE",
	18: "PUSH_NOTIFICATION_QUERY",
	19: "PUSH_NOTIFICATION_QUERY_RESPONSE",
	20: "PUSH_NOTIFICATION_REQUEST",
	21: "PUSH_NOTIFICATION_RESPONSE",
	22: "EMOJI_REACTION",
	23: "GROUP_CHAT_INVITATION",
}

var ApplicationMetadataMessage_Type_value = map[string]int32{
	"UNKNOWN":                                 0,
	"CHAT_MESSAGE":                            1,
	"CONTACT_UPDATE":                          2,
	"MEMBERSHIP_UPDATE_MESSAGE":               3,
	"PAIR_INSTALLATION":                       4,
	"SYNC_INSTALLATION":                       5,
	"REQUEST_ADDRESS_FOR_TRANSACTION":         6,
	"ACCEPT_REQUEST_ADDRESS_FOR_TRANSACTION":  7,
	"DECLINE_REQUEST_ADDRESS_FOR_TRANSACTION": 8,
	"REQUEST_TRANSACTION":                     9,
	"SEND_TRANSACTION":                        10,
	"DECLINE_REQUEST_TRANSACTION":             11,
	"SYNC_INSTALLATION_CONTACT":               12,
	"SYNC_INSTALLATION_ACCOUNT":               13,
	"SYNC_INSTALLATION_PUBLIC_CHAT":           14,
	"CONTACT_CODE_ADVERTISEMENT":              15,
	"PUSH_NOTIFICATION_REGISTRATION":          16,
	"PUSH_NOTIFICATION_REGISTRATION_RESPONSE": 17,
	"PUSH_NOTIFICATION_QUERY":                 18,
	"PUSH_NOTIFICATION_QUERY_RESPONSE":        19,
	"PUSH_NOTIFICATION_REQUEST":               20,
	"PUSH_NOTIFICATION_RESPONSE":              21,
	"EMOJI_REACTION":                          22,
	"GROUP_CHAT_INVITATION":                   23,
}

func (x ApplicationMetadataMessage_Type) String() string {
	return proto.EnumName(ApplicationMetadataMessage_Type_name, int32(x))
}

func (ApplicationMetadataMessage_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ad09a6406fcf24c7, []int{0, 0}
}

type ApplicationMetadataMessage struct {
	// Signature of the payload field
	Signature []byte `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	// This is the encoded protobuf of the application level message, i.e ChatMessage
	Payload []byte `protobuf:"bytes,2,opt,name=payload,proto3" json:"payload,omitempty"`
	// The type of protobuf message sent
	Type                 ApplicationMetadataMessage_Type `protobuf:"varint,3,opt,name=type,proto3,enum=protobuf.ApplicationMetadataMessage_Type" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *ApplicationMetadataMessage) Reset()         { *m = ApplicationMetadataMessage{} }
func (m *ApplicationMetadataMessage) String() string { return proto.CompactTextString(m) }
func (*ApplicationMetadataMessage) ProtoMessage()    {}
func (*ApplicationMetadataMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_ad09a6406fcf24c7, []int{0}
}

func (m *ApplicationMetadataMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ApplicationMetadataMessage.Unmarshal(m, b)
}
func (m *ApplicationMetadataMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ApplicationMetadataMessage.Marshal(b, m, deterministic)
}
func (m *ApplicationMetadataMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ApplicationMetadataMessage.Merge(m, src)
}
func (m *ApplicationMetadataMessage) XXX_Size() int {
	return xxx_messageInfo_ApplicationMetadataMessage.Size(m)
}
func (m *ApplicationMetadataMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ApplicationMetadataMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ApplicationMetadataMessage proto.InternalMessageInfo

func (m *ApplicationMetadataMessage) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

func (m *ApplicationMetadataMessage) GetPayload() []byte {
	if m != nil {
		return m.Payload
	}
	return nil
}

func (m *ApplicationMetadataMessage) GetType() ApplicationMetadataMessage_Type {
	if m != nil {
		return m.Type
	}
	return ApplicationMetadataMessage_UNKNOWN
}

func init() {
	proto.RegisterEnum("protobuf.ApplicationMetadataMessage_Type", ApplicationMetadataMessage_Type_name, ApplicationMetadataMessage_Type_value)
	proto.RegisterType((*ApplicationMetadataMessage)(nil), "protobuf.ApplicationMetadataMessage")
}

func init() { proto.RegisterFile("application_metadata_message.proto", fileDescriptor_ad09a6406fcf24c7) }

var fileDescriptor_ad09a6406fcf24c7 = []byte{
	// 493 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x53, 0xd3, 0x4c,
	0x14, 0xc6, 0xdf, 0xf2, 0x9f, 0x03, 0x6f, 0x5d, 0x0e, 0xd4, 0x16, 0x10, 0xa8, 0xd5, 0x51, 0xd4,
	0x99, 0x5e, 0xe8, 0xb5, 0x17, 0xcb, 0xe6, 0xd0, 0xae, 0x36, 0x9b, 0xb0, 0xbb, 0xc1, 0xe1, 0x6a,
	0x27, 0x48, 0x64, 0x3a, 0x03, 0x34, 0x43, 0xc3, 0x45, 0x3f, 0xa9, 0x9f, 0xc2, 0xef, 0xe0, 0x24,
	0x6d, 0x2d, 0xd8, 0x22, 0x57, 0x99, 0x7d, 0x9e, 0xdf, 0x39, 0x67, 0xce, 0xb3, 0x59, 0x68, 0xc4,
	0x69, 0x7a, 0xd5, 0xfd, 0x1e, 0x67, 0xdd, 0xde, 0x8d, 0xbb, 0x4e, 0xb2, 0xf8, 0x22, 0xce, 0x62,
	0x77, 0x9d, 0xf4, 0xfb, 0xf1, 0x65, 0xd2, 0x4c, 0x6f, 0x7b, 0x59, 0x0f, 0x57, 0x8a, 0xcf, 0xf9,
	0xdd, 0x8f, 0xc6, 0xaf, 0x25, 0xd8, 0xe1, 0x93, 0x02, 0x7f, 0xc4, 0xfb, 0x43, 0x1c, 0x5f, 0xc0,
	0x6a, 0xbf, 0x7b, 0x79, 0x13, 0x67, 0x77, 0xb7, 0x49, 0xad, 0x54, 0x2f, 0x1d, 0xae, 0xeb, 0x89,
	0x80, 0x35, 0x58, 0x4e, 0xe3, 0xc1, 0x55, 0x2f, 0xbe, 0xa8, 0xcd, 0x15, 0xde, 0xf8, 0x88, 0x9f,
	0x61, 0x21, 0x1b, 0xa4, 0x49, 0x6d, 0xbe, 0x5e, 0x3a, 0x2c, 0x7f, 0x7c, 0xd7, 0x1c, 0xcf, 0x6b,
	0x3e, 0x3e, 0xab, 0x69, 0x07, 0x69, 0xa2, 0x8b, 0xb2, 0xc6, 0xcf, 0x45, 0x58, 0xc8, 0x8f, 0xb8,
	0x06, 0xcb, 0x91, 0xfa, 0xaa, 0x82, 0x6f, 0x8a, 0xfd, 0x87, 0x0c, 0xd6, 0x45, 0x9b, 0x5b, 0xe7,
	0x93, 0x31, 0xbc, 0x45, 0xac, 0x84, 0x08, 0x65, 0x11, 0x28, 0xcb, 0x85, 0x75, 0x51, 0xe8, 0x71,
	0x4b, 0x6c, 0x0e, 0xf7, 0x60, 0xdb, 0x27, 0xff, 0x88, 0xb4, 0x69, 0xcb, 0x70, 0x24, 0xff, 0x29,
	0x99, 0xc7, 0x0a, 0x6c, 0x84, 0x5c, 0x6a, 0x27, 0x95, 0xb1, 0xbc, 0xd3, 0xe1, 0x56, 0x06, 0x8a,
	0x2d, 0xe4, 0xb2, 0x39, 0x53, 0xe2, 0xa1, 0xbc, 0x88, 0xaf, 0xe0, 0x40, 0xd3, 0x49, 0x44, 0xc6,
	0x3a, 0xee, 0x79, 0x9a, 0x8c, 0x71, 0xc7, 0x81, 0x76, 0x56, 0x73, 0x65, 0xb8, 0x28, 0xa0, 0x25,
	0x7c, 0x0f, 0x6f, 0xb8, 0x10, 0x14, 0x5a, 0xf7, 0x14, 0xbb, 0x8c, 0x1f, 0xe0, 0xad, 0x47, 0xa2,
	0x23, 0x15, 0x3d, 0x09, 0xaf, 0x60, 0x15, 0x36, 0xc7, 0xd0, 0x7d, 0x63, 0x15, 0xb7, 0x80, 0x19,
	0x52, 0xde, 0x03, 0x15, 0xf0, 0x00, 0x76, 0xff, 0xee, 0x7d, 0x1f, 0x58, 0xcb, 0xa3, 0x99, 0x5a,
	0xd2, 0x8d, 0x02, 0x64, 0xeb, 0xb3, 0x6d, 0x2e, 0x44, 0x10, 0x29, 0xcb, 0xfe, 0xc7, 0x97, 0xb0,
	0x37, 0x6d, 0x87, 0xd1, 0x51, 0x47, 0x0a, 0x97, 0xdf, 0x0b, 0x2b, 0xe3, 0x3e, 0xec, 0x8c, 0xef,
	0x43, 0x04, 0x1e, 0x39, 0xee, 0x9d, 0x92, 0xb6, 0xd2, 0x90, 0x4f, 0xca, 0xb2, 0x67, 0xd8, 0x80,
	0xfd, 0x30, 0x32, 0x6d, 0xa7, 0x02, 0x2b, 0x8f, 0xa5, 0x18, 0xb6, 0xd0, 0xd4, 0x92, 0xc6, 0xea,
	0x61, 0xe4, 0x2c, 0x4f, 0xe8, 0xdf, 0x8c, 0xd3, 0x64, 0xc2, 0x40, 0x19, 0x62, 0x1b, 0xb8, 0x0b,
	0xd5, 0x69, 0xf8, 0x24, 0x22, 0x7d, 0xc6, 0x10, 0x5f, 0x43, 0xfd, 0x11, 0x73, 0xd2, 0x62, 0x33,
	0xdf, 0x7a, 0xd6, 0xbc, 0x22, 0x3f, 0xb6, 0x95, 0xaf, 0x34, 0xcb, 0x1e, 0x95, 0x57, 0xf2, 0x5f,
	0x90, 0xfc, 0xe0, 0x8b, 0x74, 0x9a, 0x46, 0x39, 0x3f, 0xc7, 0x6d, 0xa8, 0xb4, 0x74, 0x10, 0x85,
	0x45, 0x2c, 0x4e, 0xaa, 0x53, 0x69, 0x87, 0xdb, 0x55, 0xcf, 0x97, 0x8a, 0x97, 0xf0, 0xe9, 0x77,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xa0, 0x4d, 0x74, 0xca, 0xa6, 0x03, 0x00, 0x00,
}
