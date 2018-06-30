// Code generated by protoc-gen-go. DO NOT EDIT.
// source: config.proto

package config

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import protobuf "google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// *
// Device Descriptor used in Configuration
type DeviceDescriptor struct {
	VendorId             *uint32  `protobuf:"varint,1,opt,name=vendor_id,json=vendorId" json:"vendor_id,omitempty"`
	ProductId            *uint32  `protobuf:"varint,2,opt,name=product_id,json=productId" json:"product_id,omitempty"`
	SerialNumber         *string  `protobuf:"bytes,3,opt,name=serial_number,json=serialNumber" json:"serial_number,omitempty"`
	Path                 *string  `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DeviceDescriptor) Reset()         { *m = DeviceDescriptor{} }
func (m *DeviceDescriptor) String() string { return proto.CompactTextString(m) }
func (*DeviceDescriptor) ProtoMessage()    {}
func (*DeviceDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_6f33bf65e2224337, []int{0}
}
func (m *DeviceDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeviceDescriptor.Unmarshal(m, b)
}
func (m *DeviceDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeviceDescriptor.Marshal(b, m, deterministic)
}
func (dst *DeviceDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeviceDescriptor.Merge(dst, src)
}
func (m *DeviceDescriptor) XXX_Size() int {
	return xxx_messageInfo_DeviceDescriptor.Size(m)
}
func (m *DeviceDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_DeviceDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_DeviceDescriptor proto.InternalMessageInfo

func (m *DeviceDescriptor) GetVendorId() uint32 {
	if m != nil && m.VendorId != nil {
		return *m.VendorId
	}
	return 0
}

func (m *DeviceDescriptor) GetProductId() uint32 {
	if m != nil && m.ProductId != nil {
		return *m.ProductId
	}
	return 0
}

func (m *DeviceDescriptor) GetSerialNumber() string {
	if m != nil && m.SerialNumber != nil {
		return *m.SerialNumber
	}
	return ""
}

func (m *DeviceDescriptor) GetPath() string {
	if m != nil && m.Path != nil {
		return *m.Path
	}
	return ""
}

// *
// Plugin Configuration
type Configuration struct {
	WhitelistUrls        []string                    `protobuf:"bytes,1,rep,name=whitelist_urls,json=whitelistUrls" json:"whitelist_urls,omitempty"`
	BlacklistUrls        []string                    `protobuf:"bytes,2,rep,name=blacklist_urls,json=blacklistUrls" json:"blacklist_urls,omitempty"`
	WireProtocol         *protobuf.FileDescriptorSet `protobuf:"bytes,3,req,name=wire_protocol,json=wireProtocol" json:"wire_protocol,omitempty"`
	KnownDevices         []*DeviceDescriptor         `protobuf:"bytes,4,rep,name=known_devices,json=knownDevices" json:"known_devices,omitempty"`
	ValidUntil           *uint32                     `protobuf:"varint,5,opt,name=valid_until,json=validUntil" json:"valid_until,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *Configuration) Reset()         { *m = Configuration{} }
func (m *Configuration) String() string { return proto.CompactTextString(m) }
func (*Configuration) ProtoMessage()    {}
func (*Configuration) Descriptor() ([]byte, []int) {
	return fileDescriptor_config_6f33bf65e2224337, []int{1}
}
func (m *Configuration) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Configuration.Unmarshal(m, b)
}
func (m *Configuration) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Configuration.Marshal(b, m, deterministic)
}
func (dst *Configuration) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Configuration.Merge(dst, src)
}
func (m *Configuration) XXX_Size() int {
	return xxx_messageInfo_Configuration.Size(m)
}
func (m *Configuration) XXX_DiscardUnknown() {
	xxx_messageInfo_Configuration.DiscardUnknown(m)
}

var xxx_messageInfo_Configuration proto.InternalMessageInfo

func (m *Configuration) GetWhitelistUrls() []string {
	if m != nil {
		return m.WhitelistUrls
	}
	return nil
}

func (m *Configuration) GetBlacklistUrls() []string {
	if m != nil {
		return m.BlacklistUrls
	}
	return nil
}

func (m *Configuration) GetWireProtocol() *protobuf.FileDescriptorSet {
	if m != nil {
		return m.WireProtocol
	}
	return nil
}

func (m *Configuration) GetKnownDevices() []*DeviceDescriptor {
	if m != nil {
		return m.KnownDevices
	}
	return nil
}

func (m *Configuration) GetValidUntil() uint32 {
	if m != nil && m.ValidUntil != nil {
		return *m.ValidUntil
	}
	return 0
}

func init() {
	proto.RegisterType((*DeviceDescriptor)(nil), "DeviceDescriptor")
	proto.RegisterType((*Configuration)(nil), "Configuration")
}

func init() { proto.RegisterFile("config.proto", fileDescriptor_config_6f33bf65e2224337) }

var fileDescriptor_config_6f33bf65e2224337 = []byte{
	// 337 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x51, 0x4b, 0xf3, 0x30,
	0x14, 0x86, 0xe9, 0xba, 0x0f, 0xbe, 0x66, 0xad, 0x68, 0xae, 0x8a, 0x22, 0x96, 0x0d, 0xa1, 0x57,
	0x19, 0x4c, 0xf0, 0x07, 0xcc, 0xa1, 0xec, 0x46, 0xa4, 0xba, 0xeb, 0xd2, 0x36, 0xd9, 0x76, 0x58,
	0xd6, 0x53, 0x92, 0x74, 0x03, 0xff, 0x80, 0x3f, 0x5b, 0x69, 0x32, 0xab, 0x78, 0x97, 0x3c, 0xef,
	0x7b, 0xe0, 0x9c, 0x87, 0x84, 0x15, 0xd6, 0x6b, 0xd8, 0xb0, 0x46, 0xa1, 0xc1, 0xcb, 0x64, 0x83,
	0xb8, 0x91, 0x62, 0x6a, 0x7f, 0x65, 0xbb, 0x9e, 0x72, 0xa1, 0x2b, 0x05, 0x8d, 0x41, 0xe5, 0x1a,
	0xe3, 0x0f, 0x8f, 0x9c, 0x2f, 0xc4, 0x01, 0x2a, 0xb1, 0xe8, 0x23, 0x7a, 0x45, 0x82, 0x83, 0xa8,
	0x39, 0xaa, 0x1c, 0x78, 0xec, 0x25, 0x5e, 0x1a, 0x65, 0xff, 0x1d, 0x58, 0x72, 0x7a, 0x4d, 0x48,
	0xa3, 0x90, 0xb7, 0x95, 0xe9, 0xd2, 0x81, 0x4d, 0x83, 0x13, 0x59, 0x72, 0x3a, 0x21, 0x91, 0x16,
	0x0a, 0x0a, 0x99, 0xd7, 0xed, 0xbe, 0x14, 0x2a, 0xf6, 0x13, 0x2f, 0x0d, 0xb2, 0xd0, 0xc1, 0x67,
	0xcb, 0x28, 0x25, 0xc3, 0xa6, 0x30, 0xdb, 0x78, 0x68, 0x33, 0xfb, 0x1e, 0x7f, 0x7a, 0x24, 0x7a,
	0xb0, 0xcb, 0xb7, 0xaa, 0x30, 0x80, 0x35, 0xbd, 0x25, 0x67, 0xc7, 0x2d, 0x18, 0x21, 0x41, 0x9b,
	0xbc, 0x55, 0x52, 0xc7, 0x5e, 0xe2, 0xa7, 0x41, 0x16, 0xf5, 0x74, 0xa5, 0xa4, 0xee, 0x6a, 0xa5,
	0x2c, 0xaa, 0xdd, 0x4f, 0x6d, 0xe0, 0x6a, 0x3d, 0xb5, 0xb5, 0x27, 0x12, 0x1d, 0x41, 0x89, 0xdc,
	0xde, 0x5d, 0xa1, 0x8c, 0xfd, 0x64, 0x90, 0x8e, 0x66, 0x63, 0xe6, 0x1c, 0xb1, 0x6f, 0x47, 0xec,
	0x11, 0xe4, 0x2f, 0x19, 0xaf, 0xc2, 0x64, 0x61, 0x37, 0xf8, 0x72, 0x9a, 0xa3, 0xf7, 0x24, 0xda,
	0xd5, 0x78, 0xac, 0x73, 0x6e, 0xbd, 0xe9, 0x78, 0x98, 0xf8, 0xe9, 0x68, 0x76, 0xc1, 0xfe, 0x7a,
	0xcc, 0x42, 0xdb, 0x73, 0x58, 0xd3, 0x1b, 0x32, 0x3a, 0x14, 0x12, 0x78, 0xde, 0xd6, 0x06, 0x64,
	0xfc, 0xcf, 0x9a, 0x23, 0x16, 0xad, 0x3a, 0x32, 0xbf, 0x23, 0x93, 0x0a, 0xf7, 0x4c, 0x17, 0x06,
	0xf5, 0x16, 0x64, 0x51, 0x6a, 0x66, 0x94, 0x78, 0x47, 0xc5, 0x24, 0x94, 0xfd, 0x7e, 0xf3, 0xf0,
	0xcd, 0x42, 0xe7, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0xf2, 0x40, 0xce, 0x0d, 0xf1, 0x01, 0x00,
	0x00,
}
