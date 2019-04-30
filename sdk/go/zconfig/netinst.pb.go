// Code generated by protoc-gen-go. DO NOT EDIT.
// source: netinst.proto

package zconfig

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

type ZNetworkInstType int32

const (
	ZNetworkInstType_ZNetInstFirst       ZNetworkInstType = 0
	ZNetworkInstType_ZnetInstSwitch      ZNetworkInstType = 1
	ZNetworkInstType_ZnetInstLocal       ZNetworkInstType = 2
	ZNetworkInstType_ZnetInstCloud       ZNetworkInstType = 3
	ZNetworkInstType_ZnetInstMesh        ZNetworkInstType = 4
	ZNetworkInstType_ZnetInstHoneyPot    ZNetworkInstType = 5
	ZNetworkInstType_ZnetInstTransparent ZNetworkInstType = 6
	ZNetworkInstType_ZNetInstLast        ZNetworkInstType = 255
)

var ZNetworkInstType_name = map[int32]string{
	0:   "ZNetInstFirst",
	1:   "ZnetInstSwitch",
	2:   "ZnetInstLocal",
	3:   "ZnetInstCloud",
	4:   "ZnetInstMesh",
	5:   "ZnetInstHoneyPot",
	6:   "ZnetInstTransparent",
	255: "ZNetInstLast",
}

var ZNetworkInstType_value = map[string]int32{
	"ZNetInstFirst":       0,
	"ZnetInstSwitch":      1,
	"ZnetInstLocal":       2,
	"ZnetInstCloud":       3,
	"ZnetInstMesh":        4,
	"ZnetInstHoneyPot":    5,
	"ZnetInstTransparent": 6,
	"ZNetInstLast":        255,
}

func (x ZNetworkInstType) String() string {
	return proto.EnumName(ZNetworkInstType_name, int32(x))
}

func (ZNetworkInstType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5d61ed8cf2f4078e, []int{0}
}

type AddressType int32

const (
	AddressType_First      AddressType = 0
	AddressType_IPV4       AddressType = 1
	AddressType_IPV6       AddressType = 2
	AddressType_CryptoIPV4 AddressType = 3
	AddressType_CryptoIPV6 AddressType = 4
	AddressType_Last       AddressType = 255
)

var AddressType_name = map[int32]string{
	0:   "First",
	1:   "IPV4",
	2:   "IPV6",
	3:   "CryptoIPV4",
	4:   "CryptoIPV6",
	255: "Last",
}

var AddressType_value = map[string]int32{
	"First":      0,
	"IPV4":       1,
	"IPV6":       2,
	"CryptoIPV4": 3,
	"CryptoIPV6": 4,
	"Last":       255,
}

func (x AddressType) String() string {
	return proto.EnumName(AddressType_name, int32(x))
}

func (AddressType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5d61ed8cf2f4078e, []int{1}
}

type ZNetworkOpaqueConfigType int32

const (
	ZNetworkOpaqueConfigType_ZNetOConfigVPN  ZNetworkOpaqueConfigType = 0
	ZNetworkOpaqueConfigType_ZNetOConfigLisp ZNetworkOpaqueConfigType = 1
)

var ZNetworkOpaqueConfigType_name = map[int32]string{
	0: "ZNetOConfigVPN",
	1: "ZNetOConfigLisp",
}

var ZNetworkOpaqueConfigType_value = map[string]int32{
	"ZNetOConfigVPN":  0,
	"ZNetOConfigLisp": 1,
}

func (x ZNetworkOpaqueConfigType) String() string {
	return proto.EnumName(ZNetworkOpaqueConfigType_name, int32(x))
}

func (ZNetworkOpaqueConfigType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_5d61ed8cf2f4078e, []int{2}
}

// Network Instance Opaque config. In future we might add more fields here
// but idea is here. This is service specific configuration.
type NetworkInstanceOpaqueConfig struct {
	Oconfig              string                     `protobuf:"bytes,1,opt,name=oconfig,proto3" json:"oconfig,omitempty"`
	LispConfig           *NetworkInstanceLispConfig `protobuf:"bytes,2,opt,name=lispConfig,proto3" json:"lispConfig,omitempty"`
	Type                 ZNetworkOpaqueConfigType   `protobuf:"varint,3,opt,name=type,proto3,enum=ZNetworkOpaqueConfigType" json:"type,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *NetworkInstanceOpaqueConfig) Reset()         { *m = NetworkInstanceOpaqueConfig{} }
func (m *NetworkInstanceOpaqueConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkInstanceOpaqueConfig) ProtoMessage()    {}
func (*NetworkInstanceOpaqueConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d61ed8cf2f4078e, []int{0}
}

func (m *NetworkInstanceOpaqueConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInstanceOpaqueConfig.Unmarshal(m, b)
}
func (m *NetworkInstanceOpaqueConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInstanceOpaqueConfig.Marshal(b, m, deterministic)
}
func (m *NetworkInstanceOpaqueConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInstanceOpaqueConfig.Merge(m, src)
}
func (m *NetworkInstanceOpaqueConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkInstanceOpaqueConfig.Size(m)
}
func (m *NetworkInstanceOpaqueConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInstanceOpaqueConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInstanceOpaqueConfig proto.InternalMessageInfo

func (m *NetworkInstanceOpaqueConfig) GetOconfig() string {
	if m != nil {
		return m.Oconfig
	}
	return ""
}

func (m *NetworkInstanceOpaqueConfig) GetLispConfig() *NetworkInstanceLispConfig {
	if m != nil {
		return m.LispConfig
	}
	return nil
}

func (m *NetworkInstanceOpaqueConfig) GetType() ZNetworkOpaqueConfigType {
	if m != nil {
		return m.Type
	}
	return ZNetworkOpaqueConfigType_ZNetOConfigVPN
}

// Lisp NetworkInstance config
type NetworkInstanceLispConfig struct {
	LispMSs             []*ZcServicePoint `protobuf:"bytes,1,rep,name=LispMSs,proto3" json:"LispMSs,omitempty"`
	LispInstanceId      uint32            `protobuf:"varint,2,opt,name=LispInstanceId,proto3" json:"LispInstanceId,omitempty"`
	Allocate            bool              `protobuf:"varint,3,opt,name=allocate,proto3" json:"allocate,omitempty"`
	Exportprivate       bool              `protobuf:"varint,4,opt,name=exportprivate,proto3" json:"exportprivate,omitempty"`
	Allocationprefix    []byte            `protobuf:"bytes,5,opt,name=allocationprefix,proto3" json:"allocationprefix,omitempty"`
	Allocationprefixlen uint32            `protobuf:"varint,6,opt,name=allocationprefixlen,proto3" json:"allocationprefixlen,omitempty"`
	// various configuration to dataPlane, lispers.net vs Zededa
	Experimental         bool     `protobuf:"varint,20,opt,name=experimental,proto3" json:"experimental,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkInstanceLispConfig) Reset()         { *m = NetworkInstanceLispConfig{} }
func (m *NetworkInstanceLispConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkInstanceLispConfig) ProtoMessage()    {}
func (*NetworkInstanceLispConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d61ed8cf2f4078e, []int{1}
}

func (m *NetworkInstanceLispConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInstanceLispConfig.Unmarshal(m, b)
}
func (m *NetworkInstanceLispConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInstanceLispConfig.Marshal(b, m, deterministic)
}
func (m *NetworkInstanceLispConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInstanceLispConfig.Merge(m, src)
}
func (m *NetworkInstanceLispConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkInstanceLispConfig.Size(m)
}
func (m *NetworkInstanceLispConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInstanceLispConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInstanceLispConfig proto.InternalMessageInfo

func (m *NetworkInstanceLispConfig) GetLispMSs() []*ZcServicePoint {
	if m != nil {
		return m.LispMSs
	}
	return nil
}

func (m *NetworkInstanceLispConfig) GetLispInstanceId() uint32 {
	if m != nil {
		return m.LispInstanceId
	}
	return 0
}

func (m *NetworkInstanceLispConfig) GetAllocate() bool {
	if m != nil {
		return m.Allocate
	}
	return false
}

func (m *NetworkInstanceLispConfig) GetExportprivate() bool {
	if m != nil {
		return m.Exportprivate
	}
	return false
}

func (m *NetworkInstanceLispConfig) GetAllocationprefix() []byte {
	if m != nil {
		return m.Allocationprefix
	}
	return nil
}

func (m *NetworkInstanceLispConfig) GetAllocationprefixlen() uint32 {
	if m != nil {
		return m.Allocationprefixlen
	}
	return 0
}

func (m *NetworkInstanceLispConfig) GetExperimental() bool {
	if m != nil {
		return m.Experimental
	}
	return false
}

type NetworkInstanceConfig struct {
	Uuidandversion *UUIDandVersion `protobuf:"bytes,1,opt,name=uuidandversion,proto3" json:"uuidandversion,omitempty"`
	Displayname    string          `protobuf:"bytes,2,opt,name=displayname,proto3" json:"displayname,omitempty"`
	// instType - Type of network instance ( local, bridge etc )
	InstType ZNetworkInstType `protobuf:"varint,4,opt,name=instType,proto3,enum=ZNetworkInstType" json:"instType,omitempty"`
	// activate
	//  - True by default. If set to false ( inactivate), the network instance
	//    configuration is downloaded to the device, but the network instance
	//    itself is not created on the device.
	Activate bool `protobuf:"varint,5,opt,name=activate,proto3" json:"activate,omitempty"`
	// port - Only a single port is supported.
	//    This is used as the external connection for the network instance.
	//    This can be a physical (eth0 ) or logical port (vlan 0).
	//    The port name comes from DeviceConfig ( When it is supported in future).
	//    If the user needs multiple physical ports, Device config should be
	//    used to create a label for multiple physical ports.
	Port *Adapter `protobuf:"bytes,20,opt,name=port,proto3" json:"port,omitempty"`
	// cfg - Used to pass some feature-specific configuration to the
	//       network instance. For Ex: Lisp, StriongSwan etc
	Cfg *NetworkInstanceOpaqueConfig `protobuf:"bytes,30,opt,name=cfg,proto3" json:"cfg,omitempty"`
	// type of ipSpec
	IpType AddressType `protobuf:"varint,39,opt,name=ipType,proto3,enum=AddressType" json:"ipType,omitempty"`
	// network ip specification
	Ip *Ipspec `protobuf:"bytes,40,opt,name=ip,proto3" json:"ip,omitempty"`
	// static DNS entry, if we are running DNS/DHCP service
	Dns                  []*ZnetStaticDNSEntry `protobuf:"bytes,41,rep,name=dns,proto3" json:"dns,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *NetworkInstanceConfig) Reset()         { *m = NetworkInstanceConfig{} }
func (m *NetworkInstanceConfig) String() string { return proto.CompactTextString(m) }
func (*NetworkInstanceConfig) ProtoMessage()    {}
func (*NetworkInstanceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d61ed8cf2f4078e, []int{2}
}

func (m *NetworkInstanceConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkInstanceConfig.Unmarshal(m, b)
}
func (m *NetworkInstanceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkInstanceConfig.Marshal(b, m, deterministic)
}
func (m *NetworkInstanceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkInstanceConfig.Merge(m, src)
}
func (m *NetworkInstanceConfig) XXX_Size() int {
	return xxx_messageInfo_NetworkInstanceConfig.Size(m)
}
func (m *NetworkInstanceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkInstanceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkInstanceConfig proto.InternalMessageInfo

func (m *NetworkInstanceConfig) GetUuidandversion() *UUIDandVersion {
	if m != nil {
		return m.Uuidandversion
	}
	return nil
}

func (m *NetworkInstanceConfig) GetDisplayname() string {
	if m != nil {
		return m.Displayname
	}
	return ""
}

func (m *NetworkInstanceConfig) GetInstType() ZNetworkInstType {
	if m != nil {
		return m.InstType
	}
	return ZNetworkInstType_ZNetInstFirst
}

func (m *NetworkInstanceConfig) GetActivate() bool {
	if m != nil {
		return m.Activate
	}
	return false
}

func (m *NetworkInstanceConfig) GetPort() *Adapter {
	if m != nil {
		return m.Port
	}
	return nil
}

func (m *NetworkInstanceConfig) GetCfg() *NetworkInstanceOpaqueConfig {
	if m != nil {
		return m.Cfg
	}
	return nil
}

func (m *NetworkInstanceConfig) GetIpType() AddressType {
	if m != nil {
		return m.IpType
	}
	return AddressType_First
}

func (m *NetworkInstanceConfig) GetIp() *Ipspec {
	if m != nil {
		return m.Ip
	}
	return nil
}

func (m *NetworkInstanceConfig) GetDns() []*ZnetStaticDNSEntry {
	if m != nil {
		return m.Dns
	}
	return nil
}

func init() {
	proto.RegisterEnum("ZNetworkInstType", ZNetworkInstType_name, ZNetworkInstType_value)
	proto.RegisterEnum("AddressType", AddressType_name, AddressType_value)
	proto.RegisterEnum("ZNetworkOpaqueConfigType", ZNetworkOpaqueConfigType_name, ZNetworkOpaqueConfigType_value)
	proto.RegisterType((*NetworkInstanceOpaqueConfig)(nil), "NetworkInstanceOpaqueConfig")
	proto.RegisterType((*NetworkInstanceLispConfig)(nil), "NetworkInstanceLispConfig")
	proto.RegisterType((*NetworkInstanceConfig)(nil), "NetworkInstanceConfig")
}

func init() { proto.RegisterFile("netinst.proto", fileDescriptor_5d61ed8cf2f4078e) }

var fileDescriptor_5d61ed8cf2f4078e = []byte{
	// 711 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x94, 0xed, 0x4e, 0xc3, 0x36,
	0x14, 0x86, 0x49, 0xbf, 0x68, 0x4f, 0x3f, 0x30, 0x2e, 0x13, 0x81, 0xa1, 0xad, 0xaa, 0xd8, 0x56,
	0x90, 0x48, 0xa7, 0x6e, 0x62, 0xd2, 0xfe, 0xb1, 0xb2, 0x8f, 0x4a, 0x50, 0xaa, 0x14, 0x98, 0xd4,
	0x7f, 0xc6, 0x31, 0xc5, 0x22, 0xb5, 0xbd, 0xd8, 0x2d, 0x94, 0xcb, 0xd9, 0x25, 0x6c, 0x57, 0xb6,
	0x2b, 0xd8, 0x14, 0x27, 0xa9, 0x4a, 0x37, 0xf6, 0xcf, 0xe7, 0x39, 0x6f, 0x4e, 0xde, 0xbc, 0xb6,
	0x03, 0x75, 0xc1, 0x0c, 0x17, 0xda, 0x78, 0x2a, 0x92, 0x46, 0x1e, 0xee, 0x04, 0x6c, 0x41, 0xe5,
	0x6c, 0x26, 0x45, 0x0a, 0x6a, 0x82, 0x19, 0x3a, 0x4b, 0xab, 0xf6, 0xef, 0x0e, 0x7c, 0x3a, 0x64,
	0xe6, 0x45, 0x46, 0xcf, 0x03, 0xa1, 0x0d, 0x11, 0x94, 0xdd, 0x28, 0xf2, 0xdb, 0x9c, 0xf5, 0xa5,
	0x78, 0xe4, 0x53, 0xec, 0xc2, 0xb6, 0xa4, 0x76, 0xe9, 0x3a, 0x2d, 0xa7, 0x53, 0xf1, 0xb3, 0x12,
	0x7f, 0x0f, 0x10, 0x72, 0xad, 0x12, 0x9d, 0x9b, 0x6b, 0x39, 0x9d, 0x6a, 0xef, 0xd0, 0xdb, 0x98,
	0x75, 0xb5, 0x52, 0xf8, 0x6b, 0x6a, 0x7c, 0x06, 0x05, 0xb3, 0x54, 0xcc, 0xcd, 0xb7, 0x9c, 0x4e,
	0xa3, 0x77, 0xe0, 0x4d, 0xd2, 0xc7, 0xd6, 0x5f, 0x7d, 0xbb, 0x54, 0xcc, 0xb7, 0xb2, 0xf6, 0x1f,
	0x39, 0x38, 0xf8, 0x70, 0x30, 0x3e, 0x81, 0xed, 0xb8, 0xba, 0x1e, 0x6b, 0xd7, 0x69, 0xe5, 0x3b,
	0xd5, 0xde, 0x8e, 0x37, 0xa1, 0x63, 0x16, 0x2d, 0x38, 0x65, 0x23, 0xc9, 0x85, 0xf1, 0xb3, 0x3e,
	0xfe, 0x12, 0x1a, 0xf1, 0x32, 0x1b, 0x32, 0x08, 0xac, 0xef, 0xba, 0xbf, 0x41, 0xf1, 0x21, 0x94,
	0x49, 0x18, 0x4a, 0x4a, 0x4c, 0xe2, 0xb1, 0xec, 0xaf, 0x6a, 0x7c, 0x0c, 0x75, 0xf6, 0xaa, 0x64,
	0x64, 0x54, 0xc4, 0x17, 0xb1, 0xa0, 0x60, 0x05, 0xef, 0x21, 0x3e, 0x05, 0x94, 0x3e, 0xc1, 0xa5,
	0x50, 0x11, 0x7b, 0xe4, 0xaf, 0x6e, 0xb1, 0xe5, 0x74, 0x6a, 0xfe, 0xbf, 0x38, 0xfe, 0x1a, 0x9a,
	0x9b, 0x2c, 0x64, 0xc2, 0x2d, 0x59, 0x6b, 0xff, 0xd5, 0xc2, 0x6d, 0xa8, 0xb1, 0x57, 0xc5, 0x22,
	0x3e, 0x63, 0xc2, 0x90, 0xd0, 0xdd, 0xb3, 0x16, 0xde, 0xb1, 0xf6, 0x5f, 0x39, 0xf8, 0x64, 0x23,
	0xb4, 0x34, 0xb0, 0xef, 0xa0, 0x31, 0x9f, 0xf3, 0x80, 0x88, 0x60, 0xc1, 0x22, 0xcd, 0xa5, 0xb0,
	0x5b, 0x1b, 0xe7, 0x76, 0x77, 0x37, 0xb8, 0x24, 0x22, 0xb8, 0x4f, 0xb0, 0xbf, 0x21, 0xc3, 0x2d,
	0xa8, 0x06, 0x5c, 0xab, 0x90, 0x2c, 0x05, 0x99, 0x31, 0x9b, 0x5d, 0xc5, 0x5f, 0x47, 0xf8, 0x0c,
	0xca, 0xf1, 0xd9, 0x8b, 0xf7, 0xce, 0xe6, 0xd2, 0xe8, 0xed, 0xae, 0x36, 0x77, 0x90, 0x36, 0xfc,
	0x95, 0xc4, 0xe6, 0x4c, 0x4d, 0x12, 0x63, 0x31, 0xcd, 0x39, 0xad, 0xf1, 0x11, 0x14, 0xe2, 0x40,
	0xed, 0xb7, 0x55, 0x7b, 0x65, 0xef, 0x22, 0x20, 0xca, 0xb0, 0xc8, 0xb7, 0x14, 0x7b, 0x90, 0xa7,
	0x8f, 0x53, 0xf7, 0x33, 0xdb, 0x3c, 0xf2, 0xfe, 0xe7, 0x08, 0xfb, 0xb1, 0x10, 0x1f, 0x43, 0x89,
	0x2b, 0x6b, 0xeb, 0x2b, 0x6b, 0xab, 0xe6, 0x5d, 0x04, 0x41, 0xc4, 0xb4, 0xb6, 0x8e, 0xd2, 0x1e,
	0xde, 0x87, 0x1c, 0x57, 0x6e, 0xc7, 0x0e, 0xdd, 0xf6, 0xb8, 0xd2, 0x8a, 0x51, 0x3f, 0xc7, 0x15,
	0xfe, 0x02, 0xf2, 0x81, 0xd0, 0xee, 0x89, 0x3d, 0x5f, 0x4d, 0x6f, 0x22, 0x98, 0x19, 0x1b, 0x62,
	0x38, 0xbd, 0x1c, 0x8e, 0x7f, 0x14, 0x26, 0x5a, 0xfa, 0x71, 0xff, 0xf4, 0x4f, 0x07, 0xd0, 0xe6,
	0xe7, 0xe2, 0x5d, 0xa8, 0xc7, 0x2c, 0xae, 0x7f, 0xe2, 0x91, 0x36, 0x68, 0x0b, 0x63, 0x68, 0xc4,
	0x23, 0x62, 0x34, 0x7e, 0xe1, 0x86, 0x3e, 0x21, 0xc7, 0xca, 0x52, 0x76, 0x25, 0x29, 0x09, 0x51,
	0x6e, 0x1d, 0xf5, 0x43, 0x39, 0x0f, 0x50, 0x1e, 0x23, 0xa8, 0x65, 0xe8, 0x9a, 0xe9, 0x27, 0x54,
	0xc0, 0x7b, 0x80, 0x32, 0xf2, 0x8b, 0x14, 0x6c, 0x39, 0x92, 0x06, 0x15, 0xf1, 0x3e, 0x34, 0x33,
	0x7a, 0x1b, 0x11, 0xa1, 0x15, 0x89, 0x98, 0x30, 0xa8, 0x84, 0x77, 0xa1, 0x96, 0xb9, 0xb9, 0x22,
	0xda, 0xa0, 0xbf, 0x9d, 0xd3, 0x5f, 0xa1, 0xba, 0x16, 0x06, 0xae, 0x40, 0x31, 0xf3, 0x59, 0x86,
	0xc2, 0x60, 0x74, 0xff, 0x2d, 0x72, 0xd2, 0xd5, 0x39, 0xca, 0xe1, 0x06, 0x40, 0x3f, 0x5a, 0x2a,
	0x23, 0x6d, 0x27, 0xff, 0xae, 0x3e, 0x47, 0x05, 0x5c, 0x81, 0x42, 0x36, 0xb8, 0x0f, 0xee, 0x47,
	0x37, 0xdb, 0x46, 0x30, 0x64, 0xe6, 0x26, 0x41, 0xf7, 0xa3, 0x21, 0xda, 0xc2, 0x4d, 0xd8, 0x59,
	0x63, 0xf1, 0x9d, 0x44, 0xce, 0x0f, 0x3f, 0xc3, 0xe7, 0x54, 0xce, 0xbc, 0x37, 0x16, 0xb0, 0x80,
	0x78, 0x34, 0xce, 0xc1, 0x9b, 0xeb, 0xe4, 0x7a, 0x27, 0x3f, 0xb1, 0xc9, 0xf1, 0x94, 0x9b, 0xa7,
	0xf9, 0x83, 0x47, 0xe5, 0xac, 0x9b, 0xe8, 0xba, 0x6c, 0xc1, 0xba, 0x3a, 0x78, 0xee, 0x4e, 0x65,
	0xf7, 0x2d, 0xf9, 0x61, 0x3d, 0x94, 0xac, 0xf8, 0x9b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x29,
	0xfe, 0xf8, 0xb3, 0x21, 0x05, 0x00, 0x00,
}