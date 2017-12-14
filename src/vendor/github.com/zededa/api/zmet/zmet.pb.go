// Code generated by protoc-gen-go. DO NOT EDIT.
// source: zmet.proto

/*
Package zmet is a generated protocol buffer package.

It is generated from these files:
	zmet.proto

It has these top-level messages:
	ZInfoManufacturer
	ZInfoNetwork
	ZinfoPeripheral
	ZInfoSW
	ZInfoDevice
	ZInfoHypervisor
	ZInfoApp
	ZInfoMsg
	CpuMetric
	MemoryMetric
	NetworkMetric
	DeviceMetric
	ZMetricMsg
*/
package zmet

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

//
// Broadly there are two types
// Info : information that is discovered/rarely changes
// Metrics: information that gets updated periodically
// Protobuf definitions in this file follow the convention.
type ZInfoTypes int32

const (
	ZInfoTypes_ZiNop        ZInfoTypes = 0
	ZInfoTypes_ZiDevice     ZInfoTypes = 1
	ZInfoTypes_ZiHypervisor ZInfoTypes = 2
	ZInfoTypes_ZiApp        ZInfoTypes = 3
)

var ZInfoTypes_name = map[int32]string{
	0: "ZiNop",
	1: "ZiDevice",
	2: "ZiHypervisor",
	3: "ZiApp",
}
var ZInfoTypes_value = map[string]int32{
	"ZiNop":        0,
	"ZiDevice":     1,
	"ZiHypervisor": 2,
	"ZiApp":        3,
}

func (x ZInfoTypes) String() string {
	return proto.EnumName(ZInfoTypes_name, int32(x))
}
func (ZInfoTypes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ZPeripheralTypes int32

const (
	ZPeripheralTypes_ZpNone    ZPeripheralTypes = 0
	ZPeripheralTypes_ZpStorage ZPeripheralTypes = 1
	ZPeripheralTypes_ZpNetwork ZPeripheralTypes = 2
)

var ZPeripheralTypes_name = map[int32]string{
	0: "ZpNone",
	1: "ZpStorage",
	2: "ZpNetwork",
}
var ZPeripheralTypes_value = map[string]int32{
	"ZpNone":    0,
	"ZpStorage": 1,
	"ZpNetwork": 2,
}

func (x ZPeripheralTypes) String() string {
	return proto.EnumName(ZPeripheralTypes_name, int32(x))
}
func (ZPeripheralTypes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ZSwState int32

const (
	ZSwState_INVALID          ZSwState = 0
	ZSwState_INITIAL          ZSwState = 1
	ZSwState_DOWNLOAD_STARTED ZSwState = 2
	ZSwState_DOWNLOADED       ZSwState = 3
	ZSwState_DELIVERED        ZSwState = 4
	ZSwState_INSTALLED        ZSwState = 5
)

var ZSwState_name = map[int32]string{
	0: "INVALID",
	1: "INITIAL",
	2: "DOWNLOAD_STARTED",
	3: "DOWNLOADED",
	4: "DELIVERED",
	5: "INSTALLED",
}
var ZSwState_value = map[string]int32{
	"INVALID":          0,
	"INITIAL":          1,
	"DOWNLOAD_STARTED": 2,
	"DOWNLOADED":       3,
	"DELIVERED":        4,
	"INSTALLED":        5,
}

func (x ZSwState) String() string {
	return proto.EnumName(ZSwState_name, int32(x))
}
func (ZSwState) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type ZmetricTypes int32

const (
	ZmetricTypes_ZmNop    ZmetricTypes = 0
	ZmetricTypes_ZmDevice ZmetricTypes = 1
	ZmetricTypes_ZmApp    ZmetricTypes = 3
)

var ZmetricTypes_name = map[int32]string{
	0: "ZmNop",
	1: "ZmDevice",
	3: "ZmApp",
}
var ZmetricTypes_value = map[string]int32{
	"ZmNop":    0,
	"ZmDevice": 1,
	"ZmApp":    3,
}

func (x ZmetricTypes) String() string {
	return proto.EnumName(ZmetricTypes_name, int32(x))
}
func (ZmetricTypes) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

// Manufacturing info, product name, model, version etc.
type ZInfoManufacturer struct {
	Manufacturer string `protobuf:"bytes,1,opt,name=manufacturer" json:"manufacturer,omitempty"`
	ProductName  string `protobuf:"bytes,2,opt,name=productName" json:"productName,omitempty"`
	Version      string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
	SerialNumber string `protobuf:"bytes,4,opt,name=serialNumber" json:"serialNumber,omitempty"`
	UUID         string `protobuf:"bytes,5,opt,name=UUID" json:"UUID,omitempty"`
}

func (m *ZInfoManufacturer) Reset()                    { *m = ZInfoManufacturer{} }
func (m *ZInfoManufacturer) String() string            { return proto.CompactTextString(m) }
func (*ZInfoManufacturer) ProtoMessage()               {}
func (*ZInfoManufacturer) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *ZInfoManufacturer) GetManufacturer() string {
	if m != nil {
		return m.Manufacturer
	}
	return ""
}

func (m *ZInfoManufacturer) GetProductName() string {
	if m != nil {
		return m.ProductName
	}
	return ""
}

func (m *ZInfoManufacturer) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *ZInfoManufacturer) GetSerialNumber() string {
	if m != nil {
		return m.SerialNumber
	}
	return ""
}

func (m *ZInfoManufacturer) GetUUID() string {
	if m != nil {
		return m.UUID
	}
	return ""
}

type ZInfoNetwork struct {
	IPAddr  string `protobuf:"bytes,1,opt,name=IPAddr" json:"IPAddr,omitempty"`
	GwAddr  string `protobuf:"bytes,2,opt,name=gwAddr" json:"gwAddr,omitempty"`
	MacAddr string `protobuf:"bytes,3,opt,name=macAddr" json:"macAddr,omitempty"`
	DevName string `protobuf:"bytes,4,opt,name=devName" json:"devName,omitempty"`
}

func (m *ZInfoNetwork) Reset()                    { *m = ZInfoNetwork{} }
func (m *ZInfoNetwork) String() string            { return proto.CompactTextString(m) }
func (*ZInfoNetwork) ProtoMessage()               {}
func (*ZInfoNetwork) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ZInfoNetwork) GetIPAddr() string {
	if m != nil {
		return m.IPAddr
	}
	return ""
}

func (m *ZInfoNetwork) GetGwAddr() string {
	if m != nil {
		return m.GwAddr
	}
	return ""
}

func (m *ZInfoNetwork) GetMacAddr() string {
	if m != nil {
		return m.MacAddr
	}
	return ""
}

func (m *ZInfoNetwork) GetDevName() string {
	if m != nil {
		return m.DevName
	}
	return ""
}

type ZinfoPeripheral struct {
	Ztype     ZPeripheralTypes   `protobuf:"varint,1,opt,name=ztype,enum=ZPeripheralTypes" json:"ztype,omitempty"`
	Pluggable bool               `protobuf:"varint,2,opt,name=pluggable" json:"pluggable,omitempty"`
	Minfo     *ZInfoManufacturer `protobuf:"bytes,3,opt,name=minfo" json:"minfo,omitempty"`
}

func (m *ZinfoPeripheral) Reset()                    { *m = ZinfoPeripheral{} }
func (m *ZinfoPeripheral) String() string            { return proto.CompactTextString(m) }
func (*ZinfoPeripheral) ProtoMessage()               {}
func (*ZinfoPeripheral) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ZinfoPeripheral) GetZtype() ZPeripheralTypes {
	if m != nil {
		return m.Ztype
	}
	return ZPeripheralTypes_ZpNone
}

func (m *ZinfoPeripheral) GetPluggable() bool {
	if m != nil {
		return m.Pluggable
	}
	return false
}

func (m *ZinfoPeripheral) GetMinfo() *ZInfoManufacturer {
	if m != nil {
		return m.Minfo
	}
	return nil
}

type ZInfoSW struct {
	SwVersion string   `protobuf:"bytes,2,opt,name=swVersion" json:"swVersion,omitempty"`
	SwHash    string   `protobuf:"bytes,3,opt,name=swHash" json:"swHash,omitempty"`
	State     ZSwState `protobuf:"varint,4,opt,name=state,enum=ZSwState" json:"state,omitempty"`
	Activated bool     `protobuf:"varint,5,opt,name=activated" json:"activated,omitempty"`
}

func (m *ZInfoSW) Reset()                    { *m = ZInfoSW{} }
func (m *ZInfoSW) String() string            { return proto.CompactTextString(m) }
func (*ZInfoSW) ProtoMessage()               {}
func (*ZInfoSW) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ZInfoSW) GetSwVersion() string {
	if m != nil {
		return m.SwVersion
	}
	return ""
}

func (m *ZInfoSW) GetSwHash() string {
	if m != nil {
		return m.SwHash
	}
	return ""
}

func (m *ZInfoSW) GetState() ZSwState {
	if m != nil {
		return m.State
	}
	return ZSwState_INVALID
}

func (m *ZInfoSW) GetActivated() bool {
	if m != nil {
		return m.Activated
	}
	return false
}

// Base device info, as discovered by Xen (or OS on bare metal)
type ZInfoDevice struct {
	MachineArch  string             `protobuf:"bytes,4,opt,name=machineArch" json:"machineArch,omitempty"`
	CpuArch      string             `protobuf:"bytes,5,opt,name=cpuArch" json:"cpuArch,omitempty"`
	Platform     string             `protobuf:"bytes,6,opt,name=platform" json:"platform,omitempty"`
	Ncpu         uint32             `protobuf:"varint,7,opt,name=ncpu" json:"ncpu,omitempty"`
	Memory       uint64             `protobuf:"varint,8,opt,name=memory" json:"memory,omitempty"`
	Storage      uint64             `protobuf:"varint,9,opt,name=storage" json:"storage,omitempty"`
	Devices      []*ZinfoPeripheral `protobuf:"bytes,10,rep,name=devices" json:"devices,omitempty"`
	Minfo        *ZInfoManufacturer `protobuf:"bytes,11,opt,name=minfo" json:"minfo,omitempty"`
	Software     *ZInfoSW           `protobuf:"bytes,12,opt,name=software" json:"software,omitempty"`
	Network      []*ZInfoNetwork    `protobuf:"bytes,13,rep,name=network" json:"network,omitempty"`
	SoftwareList []*ZInfoSW         `protobuf:"bytes,14,rep,name=softwareList" json:"softwareList,omitempty"`
}

func (m *ZInfoDevice) Reset()                    { *m = ZInfoDevice{} }
func (m *ZInfoDevice) String() string            { return proto.CompactTextString(m) }
func (*ZInfoDevice) ProtoMessage()               {}
func (*ZInfoDevice) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *ZInfoDevice) GetMachineArch() string {
	if m != nil {
		return m.MachineArch
	}
	return ""
}

func (m *ZInfoDevice) GetCpuArch() string {
	if m != nil {
		return m.CpuArch
	}
	return ""
}

func (m *ZInfoDevice) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *ZInfoDevice) GetNcpu() uint32 {
	if m != nil {
		return m.Ncpu
	}
	return 0
}

func (m *ZInfoDevice) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ZInfoDevice) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *ZInfoDevice) GetDevices() []*ZinfoPeripheral {
	if m != nil {
		return m.Devices
	}
	return nil
}

func (m *ZInfoDevice) GetMinfo() *ZInfoManufacturer {
	if m != nil {
		return m.Minfo
	}
	return nil
}

func (m *ZInfoDevice) GetSoftware() *ZInfoSW {
	if m != nil {
		return m.Software
	}
	return nil
}

func (m *ZInfoDevice) GetNetwork() []*ZInfoNetwork {
	if m != nil {
		return m.Network
	}
	return nil
}

func (m *ZInfoDevice) GetSoftwareList() []*ZInfoSW {
	if m != nil {
		return m.SoftwareList
	}
	return nil
}

// Device info from DOM0 perspective, if it exists.
type ZInfoHypervisor struct {
	Ncpu     uint32   `protobuf:"varint,3,opt,name=ncpu" json:"ncpu,omitempty"`
	Memory   uint64   `protobuf:"varint,4,opt,name=memory" json:"memory,omitempty"`
	Storage  uint64   `protobuf:"varint,5,opt,name=storage" json:"storage,omitempty"`
	Software *ZInfoSW `protobuf:"bytes,6,opt,name=software" json:"software,omitempty"`
}

func (m *ZInfoHypervisor) Reset()                    { *m = ZInfoHypervisor{} }
func (m *ZInfoHypervisor) String() string            { return proto.CompactTextString(m) }
func (*ZInfoHypervisor) ProtoMessage()               {}
func (*ZInfoHypervisor) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *ZInfoHypervisor) GetNcpu() uint32 {
	if m != nil {
		return m.Ncpu
	}
	return 0
}

func (m *ZInfoHypervisor) GetMemory() uint64 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ZInfoHypervisor) GetStorage() uint64 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *ZInfoHypervisor) GetSoftware() *ZInfoSW {
	if m != nil {
		return m.Software
	}
	return nil
}

type ZInfoApp struct {
	AppID    string   `protobuf:"bytes,1,opt,name=AppID" json:"AppID,omitempty"`
	Ncpu     uint32   `protobuf:"varint,2,opt,name=ncpu" json:"ncpu,omitempty"`
	Memory   uint32   `protobuf:"varint,3,opt,name=memory" json:"memory,omitempty"`
	Storage  uint32   `protobuf:"varint,4,opt,name=storage" json:"storage,omitempty"`
	Software *ZInfoSW `protobuf:"bytes,5,opt,name=software" json:"software,omitempty"`
}

func (m *ZInfoApp) Reset()                    { *m = ZInfoApp{} }
func (m *ZInfoApp) String() string            { return proto.CompactTextString(m) }
func (*ZInfoApp) ProtoMessage()               {}
func (*ZInfoApp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ZInfoApp) GetAppID() string {
	if m != nil {
		return m.AppID
	}
	return ""
}

func (m *ZInfoApp) GetNcpu() uint32 {
	if m != nil {
		return m.Ncpu
	}
	return 0
}

func (m *ZInfoApp) GetMemory() uint32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *ZInfoApp) GetStorage() uint32 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *ZInfoApp) GetSoftware() *ZInfoSW {
	if m != nil {
		return m.Software
	}
	return nil
}

type ZInfoMsg struct {
	Ztype ZInfoTypes `protobuf:"varint,1,opt,name=ztype,enum=ZInfoTypes" json:"ztype,omitempty"`
	DevId string     `protobuf:"bytes,2,opt,name=devId" json:"devId,omitempty"`
	// Types that are valid to be assigned to InfoContent:
	//	*ZInfoMsg_Dinfo
	//	*ZInfoMsg_Hinfo
	//	*ZInfoMsg_Ainfo
	InfoContent isZInfoMsg_InfoContent `protobuf_oneof:"InfoContent"`
}

func (m *ZInfoMsg) Reset()                    { *m = ZInfoMsg{} }
func (m *ZInfoMsg) String() string            { return proto.CompactTextString(m) }
func (*ZInfoMsg) ProtoMessage()               {}
func (*ZInfoMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type isZInfoMsg_InfoContent interface {
	isZInfoMsg_InfoContent()
}

type ZInfoMsg_Dinfo struct {
	Dinfo *ZInfoDevice `protobuf:"bytes,3,opt,name=dinfo,oneof"`
}
type ZInfoMsg_Hinfo struct {
	Hinfo *ZInfoHypervisor `protobuf:"bytes,4,opt,name=hinfo,oneof"`
}
type ZInfoMsg_Ainfo struct {
	Ainfo *ZInfoApp `protobuf:"bytes,5,opt,name=ainfo,oneof"`
}

func (*ZInfoMsg_Dinfo) isZInfoMsg_InfoContent() {}
func (*ZInfoMsg_Hinfo) isZInfoMsg_InfoContent() {}
func (*ZInfoMsg_Ainfo) isZInfoMsg_InfoContent() {}

func (m *ZInfoMsg) GetInfoContent() isZInfoMsg_InfoContent {
	if m != nil {
		return m.InfoContent
	}
	return nil
}

func (m *ZInfoMsg) GetZtype() ZInfoTypes {
	if m != nil {
		return m.Ztype
	}
	return ZInfoTypes_ZiNop
}

func (m *ZInfoMsg) GetDevId() string {
	if m != nil {
		return m.DevId
	}
	return ""
}

func (m *ZInfoMsg) GetDinfo() *ZInfoDevice {
	if x, ok := m.GetInfoContent().(*ZInfoMsg_Dinfo); ok {
		return x.Dinfo
	}
	return nil
}

func (m *ZInfoMsg) GetHinfo() *ZInfoHypervisor {
	if x, ok := m.GetInfoContent().(*ZInfoMsg_Hinfo); ok {
		return x.Hinfo
	}
	return nil
}

func (m *ZInfoMsg) GetAinfo() *ZInfoApp {
	if x, ok := m.GetInfoContent().(*ZInfoMsg_Ainfo); ok {
		return x.Ainfo
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ZInfoMsg) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ZInfoMsg_OneofMarshaler, _ZInfoMsg_OneofUnmarshaler, _ZInfoMsg_OneofSizer, []interface{}{
		(*ZInfoMsg_Dinfo)(nil),
		(*ZInfoMsg_Hinfo)(nil),
		(*ZInfoMsg_Ainfo)(nil),
	}
}

func _ZInfoMsg_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ZInfoMsg)
	// InfoContent
	switch x := m.InfoContent.(type) {
	case *ZInfoMsg_Dinfo:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dinfo); err != nil {
			return err
		}
	case *ZInfoMsg_Hinfo:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Hinfo); err != nil {
			return err
		}
	case *ZInfoMsg_Ainfo:
		b.EncodeVarint(5<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Ainfo); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ZInfoMsg.InfoContent has unexpected type %T", x)
	}
	return nil
}

func _ZInfoMsg_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ZInfoMsg)
	switch tag {
	case 3: // InfoContent.dinfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ZInfoDevice)
		err := b.DecodeMessage(msg)
		m.InfoContent = &ZInfoMsg_Dinfo{msg}
		return true, err
	case 4: // InfoContent.hinfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ZInfoHypervisor)
		err := b.DecodeMessage(msg)
		m.InfoContent = &ZInfoMsg_Hinfo{msg}
		return true, err
	case 5: // InfoContent.ainfo
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(ZInfoApp)
		err := b.DecodeMessage(msg)
		m.InfoContent = &ZInfoMsg_Ainfo{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ZInfoMsg_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ZInfoMsg)
	// InfoContent
	switch x := m.InfoContent.(type) {
	case *ZInfoMsg_Dinfo:
		s := proto.Size(x.Dinfo)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ZInfoMsg_Hinfo:
		s := proto.Size(x.Hinfo)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ZInfoMsg_Ainfo:
		s := proto.Size(x.Ainfo)
		n += proto.SizeVarint(5<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type CpuMetric struct {
	UpTime         uint32  `protobuf:"varint,2,opt,name=upTime" json:"upTime,omitempty"`
	CpuUtilization float64 `protobuf:"fixed64,3,opt,name=cpuUtilization" json:"cpuUtilization,omitempty"`
	Usr            float64 `protobuf:"fixed64,4,opt,name=usr" json:"usr,omitempty"`
	Nice           float64 `protobuf:"fixed64,5,opt,name=nice" json:"nice,omitempty"`
	System         float64 `protobuf:"fixed64,6,opt,name=system" json:"system,omitempty"`
	Io             float64 `protobuf:"fixed64,7,opt,name=io" json:"io,omitempty"`
	Irq            float64 `protobuf:"fixed64,8,opt,name=irq" json:"irq,omitempty"`
	Soft           float64 `protobuf:"fixed64,9,opt,name=soft" json:"soft,omitempty"`
	Steal          float64 `protobuf:"fixed64,10,opt,name=steal" json:"steal,omitempty"`
	Guest          float64 `protobuf:"fixed64,11,opt,name=guest" json:"guest,omitempty"`
	Idle           float64 `protobuf:"fixed64,12,opt,name=idle" json:"idle,omitempty"`
}

func (m *CpuMetric) Reset()                    { *m = CpuMetric{} }
func (m *CpuMetric) String() string            { return proto.CompactTextString(m) }
func (*CpuMetric) ProtoMessage()               {}
func (*CpuMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CpuMetric) GetUpTime() uint32 {
	if m != nil {
		return m.UpTime
	}
	return 0
}

func (m *CpuMetric) GetCpuUtilization() float64 {
	if m != nil {
		return m.CpuUtilization
	}
	return 0
}

func (m *CpuMetric) GetUsr() float64 {
	if m != nil {
		return m.Usr
	}
	return 0
}

func (m *CpuMetric) GetNice() float64 {
	if m != nil {
		return m.Nice
	}
	return 0
}

func (m *CpuMetric) GetSystem() float64 {
	if m != nil {
		return m.System
	}
	return 0
}

func (m *CpuMetric) GetIo() float64 {
	if m != nil {
		return m.Io
	}
	return 0
}

func (m *CpuMetric) GetIrq() float64 {
	if m != nil {
		return m.Irq
	}
	return 0
}

func (m *CpuMetric) GetSoft() float64 {
	if m != nil {
		return m.Soft
	}
	return 0
}

func (m *CpuMetric) GetSteal() float64 {
	if m != nil {
		return m.Steal
	}
	return 0
}

func (m *CpuMetric) GetGuest() float64 {
	if m != nil {
		return m.Guest
	}
	return 0
}

func (m *CpuMetric) GetIdle() float64 {
	if m != nil {
		return m.Idle
	}
	return 0
}

type MemoryMetric struct {
	UsedMem         uint32  `protobuf:"varint,2,opt,name=usedMem" json:"usedMem,omitempty"`
	AvailMem        uint32  `protobuf:"varint,3,opt,name=availMem" json:"availMem,omitempty"`
	UsedPercentage  float64 `protobuf:"fixed64,4,opt,name=usedPercentage" json:"usedPercentage,omitempty"`
	AvailPercentage float64 `protobuf:"fixed64,5,opt,name=availPercentage" json:"availPercentage,omitempty"`
}

func (m *MemoryMetric) Reset()                    { *m = MemoryMetric{} }
func (m *MemoryMetric) String() string            { return proto.CompactTextString(m) }
func (*MemoryMetric) ProtoMessage()               {}
func (*MemoryMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *MemoryMetric) GetUsedMem() uint32 {
	if m != nil {
		return m.UsedMem
	}
	return 0
}

func (m *MemoryMetric) GetAvailMem() uint32 {
	if m != nil {
		return m.AvailMem
	}
	return 0
}

func (m *MemoryMetric) GetUsedPercentage() float64 {
	if m != nil {
		return m.UsedPercentage
	}
	return 0
}

func (m *MemoryMetric) GetAvailPercentage() float64 {
	if m != nil {
		return m.AvailPercentage
	}
	return 0
}

type NetworkMetric struct {
	IName   string `protobuf:"bytes,1,opt,name=iName" json:"iName,omitempty"`
	TxBytes uint64 `protobuf:"varint,2,opt,name=txBytes" json:"txBytes,omitempty"`
	RxBytes uint64 `protobuf:"varint,3,opt,name=rxBytes" json:"rxBytes,omitempty"`
	TxDrops uint64 `protobuf:"varint,4,opt,name=txDrops" json:"txDrops,omitempty"`
	RxDrops uint64 `protobuf:"varint,5,opt,name=rxDrops" json:"rxDrops,omitempty"`
	TxRate  uint64 `protobuf:"varint,6,opt,name=txRate" json:"txRate,omitempty"`
	RxRate  uint64 `protobuf:"varint,7,opt,name=rxRate" json:"rxRate,omitempty"`
}

func (m *NetworkMetric) Reset()                    { *m = NetworkMetric{} }
func (m *NetworkMetric) String() string            { return proto.CompactTextString(m) }
func (*NetworkMetric) ProtoMessage()               {}
func (*NetworkMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *NetworkMetric) GetIName() string {
	if m != nil {
		return m.IName
	}
	return ""
}

func (m *NetworkMetric) GetTxBytes() uint64 {
	if m != nil {
		return m.TxBytes
	}
	return 0
}

func (m *NetworkMetric) GetRxBytes() uint64 {
	if m != nil {
		return m.RxBytes
	}
	return 0
}

func (m *NetworkMetric) GetTxDrops() uint64 {
	if m != nil {
		return m.TxDrops
	}
	return 0
}

func (m *NetworkMetric) GetRxDrops() uint64 {
	if m != nil {
		return m.RxDrops
	}
	return 0
}

func (m *NetworkMetric) GetTxRate() uint64 {
	if m != nil {
		return m.TxRate
	}
	return 0
}

func (m *NetworkMetric) GetRxRate() uint64 {
	if m != nil {
		return m.RxRate
	}
	return 0
}

type DeviceMetric struct {
	Cpu     *CpuMetric       `protobuf:"bytes,1,opt,name=cpu" json:"cpu,omitempty"`
	Memory  *MemoryMetric    `protobuf:"bytes,2,opt,name=memory" json:"memory,omitempty"`
	Network []*NetworkMetric `protobuf:"bytes,3,rep,name=network" json:"network,omitempty"`
}

func (m *DeviceMetric) Reset()                    { *m = DeviceMetric{} }
func (m *DeviceMetric) String() string            { return proto.CompactTextString(m) }
func (*DeviceMetric) ProtoMessage()               {}
func (*DeviceMetric) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *DeviceMetric) GetCpu() *CpuMetric {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *DeviceMetric) GetMemory() *MemoryMetric {
	if m != nil {
		return m.Memory
	}
	return nil
}

func (m *DeviceMetric) GetNetwork() []*NetworkMetric {
	if m != nil {
		return m.Network
	}
	return nil
}

type ZMetricMsg struct {
	DevID string       `protobuf:"bytes,1,opt,name=devID" json:"devID,omitempty"`
	Ztype ZmetricTypes `protobuf:"varint,2,opt,name=ztype,enum=ZmetricTypes" json:"ztype,omitempty"`
	// Types that are valid to be assigned to MetricContent:
	//	*ZMetricMsg_Dm
	MetricContent isZMetricMsg_MetricContent `protobuf_oneof:"MetricContent"`
}

func (m *ZMetricMsg) Reset()                    { *m = ZMetricMsg{} }
func (m *ZMetricMsg) String() string            { return proto.CompactTextString(m) }
func (*ZMetricMsg) ProtoMessage()               {}
func (*ZMetricMsg) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type isZMetricMsg_MetricContent interface {
	isZMetricMsg_MetricContent()
}

type ZMetricMsg_Dm struct {
	Dm *DeviceMetric `protobuf:"bytes,4,opt,name=dm,oneof"`
}

func (*ZMetricMsg_Dm) isZMetricMsg_MetricContent() {}

func (m *ZMetricMsg) GetMetricContent() isZMetricMsg_MetricContent {
	if m != nil {
		return m.MetricContent
	}
	return nil
}

func (m *ZMetricMsg) GetDevID() string {
	if m != nil {
		return m.DevID
	}
	return ""
}

func (m *ZMetricMsg) GetZtype() ZmetricTypes {
	if m != nil {
		return m.Ztype
	}
	return ZmetricTypes_ZmNop
}

func (m *ZMetricMsg) GetDm() *DeviceMetric {
	if x, ok := m.GetMetricContent().(*ZMetricMsg_Dm); ok {
		return x.Dm
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ZMetricMsg) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ZMetricMsg_OneofMarshaler, _ZMetricMsg_OneofUnmarshaler, _ZMetricMsg_OneofSizer, []interface{}{
		(*ZMetricMsg_Dm)(nil),
	}
}

func _ZMetricMsg_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ZMetricMsg)
	// MetricContent
	switch x := m.MetricContent.(type) {
	case *ZMetricMsg_Dm:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Dm); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("ZMetricMsg.MetricContent has unexpected type %T", x)
	}
	return nil
}

func _ZMetricMsg_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ZMetricMsg)
	switch tag {
	case 4: // MetricContent.dm
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(DeviceMetric)
		err := b.DecodeMessage(msg)
		m.MetricContent = &ZMetricMsg_Dm{msg}
		return true, err
	default:
		return false, nil
	}
}

func _ZMetricMsg_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ZMetricMsg)
	// MetricContent
	switch x := m.MetricContent.(type) {
	case *ZMetricMsg_Dm:
		s := proto.Size(x.Dm)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*ZInfoManufacturer)(nil), "ZInfoManufacturer")
	proto.RegisterType((*ZInfoNetwork)(nil), "ZInfoNetwork")
	proto.RegisterType((*ZinfoPeripheral)(nil), "ZinfoPeripheral")
	proto.RegisterType((*ZInfoSW)(nil), "ZInfoSW")
	proto.RegisterType((*ZInfoDevice)(nil), "ZInfoDevice")
	proto.RegisterType((*ZInfoHypervisor)(nil), "ZInfoHypervisor")
	proto.RegisterType((*ZInfoApp)(nil), "ZInfoApp")
	proto.RegisterType((*ZInfoMsg)(nil), "ZInfoMsg")
	proto.RegisterType((*CpuMetric)(nil), "cpuMetric")
	proto.RegisterType((*MemoryMetric)(nil), "memoryMetric")
	proto.RegisterType((*NetworkMetric)(nil), "networkMetric")
	proto.RegisterType((*DeviceMetric)(nil), "deviceMetric")
	proto.RegisterType((*ZMetricMsg)(nil), "ZMetricMsg")
	proto.RegisterEnum("ZInfoTypes", ZInfoTypes_name, ZInfoTypes_value)
	proto.RegisterEnum("ZPeripheralTypes", ZPeripheralTypes_name, ZPeripheralTypes_value)
	proto.RegisterEnum("ZSwState", ZSwState_name, ZSwState_value)
	proto.RegisterEnum("ZmetricTypes", ZmetricTypes_name, ZmetricTypes_value)
}

func init() { proto.RegisterFile("zmet.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 1212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0xcb, 0x6e, 0xdb, 0x46,
	0x14, 0x35, 0x25, 0x51, 0x8f, 0xab, 0x87, 0x99, 0x81, 0x11, 0x10, 0x41, 0x80, 0xb8, 0x6c, 0xda,
	0x08, 0x46, 0xc1, 0x00, 0xee, 0xb6, 0x5d, 0xc8, 0x95, 0x01, 0x0b, 0x90, 0x95, 0x60, 0xec, 0x24,
	0x80, 0x36, 0x05, 0x4d, 0x8e, 0xa5, 0x41, 0x45, 0x71, 0x3a, 0x1c, 0xca, 0x91, 0x81, 0x74, 0xd1,
	0x75, 0xd7, 0xfd, 0x89, 0x7e, 0x43, 0x97, 0xfd, 0xad, 0xb6, 0x98, 0x3b, 0x43, 0x3d, 0x5c, 0x07,
	0xdd, 0xf1, 0x9c, 0x73, 0x39, 0xba, 0x8f, 0x39, 0x97, 0x02, 0xb8, 0x4f, 0x99, 0x0a, 0x85, 0xcc,
	0x54, 0x16, 0xfc, 0xe1, 0xc0, 0x93, 0xe9, 0x68, 0x79, 0x9b, 0x5d, 0x46, 0xcb, 0xe2, 0x36, 0x8a,
	0x55, 0x21, 0x99, 0x24, 0x01, 0x74, 0xd2, 0x1d, 0xec, 0x3b, 0xc7, 0x4e, 0xbf, 0x45, 0xf7, 0x38,
	0x72, 0x0c, 0x6d, 0x21, 0xb3, 0xa4, 0x88, 0xd5, 0x24, 0x4a, 0x99, 0x5f, 0xc1, 0x90, 0x5d, 0x8a,
	0xf8, 0xd0, 0x58, 0x31, 0x99, 0xf3, 0x6c, 0xe9, 0x57, 0x51, 0x2d, 0xa1, 0x3e, 0x3f, 0x67, 0x92,
	0x47, 0x8b, 0x49, 0x91, 0xde, 0x30, 0xe9, 0xd7, 0xcc, 0xf9, 0xbb, 0x1c, 0x21, 0x50, 0x7b, 0xf7,
	0x6e, 0x34, 0xf4, 0x5d, 0xd4, 0xf0, 0x39, 0x90, 0xd0, 0xc1, 0x64, 0x27, 0x4c, 0xdd, 0x65, 0xf2,
	0x27, 0xf2, 0x14, 0xea, 0xa3, 0xb7, 0x83, 0x24, 0x29, 0x33, 0xb4, 0x48, 0xf3, 0xb3, 0x3b, 0xe4,
	0x4d, 0x5a, 0x16, 0xe9, 0x8c, 0xd2, 0x28, 0x46, 0xc1, 0x66, 0x64, 0xa1, 0x56, 0x12, 0xb6, 0xc2,
	0x4a, 0x4c, 0x32, 0x25, 0x0c, 0x7e, 0x75, 0xe0, 0x70, 0xca, 0x97, 0xb7, 0xd9, 0x5b, 0x26, 0xb9,
	0x98, 0x33, 0x19, 0x2d, 0xc8, 0x2b, 0x70, 0xef, 0xd5, 0x5a, 0x30, 0xfc, 0xd9, 0xde, 0xe9, 0x93,
	0x70, 0xba, 0x15, 0xaf, 0xd7, 0x82, 0xe5, 0xd4, 0xe8, 0xe4, 0x39, 0xb4, 0xc4, 0xa2, 0x98, 0xcd,
	0xa2, 0x9b, 0x85, 0x69, 0x51, 0x93, 0x6e, 0x09, 0xd2, 0x07, 0x37, 0xd5, 0x27, 0x63, 0x32, 0xed,
	0x53, 0x12, 0xfe, 0x67, 0x12, 0xd4, 0x04, 0x04, 0xbf, 0x40, 0x03, 0xb5, 0xab, 0x0f, 0xfa, 0xc8,
	0xfc, 0xee, 0xbd, 0xed, 0xab, 0x29, 0x6f, 0x4b, 0xe8, 0xca, 0xf3, 0xbb, 0x8b, 0x28, 0x9f, 0xdb,
	0x02, 0x2d, 0x22, 0x2f, 0xc0, 0xcd, 0x55, 0xa4, 0x4c, 0x75, 0xbd, 0xd3, 0x56, 0x38, 0xbd, 0xba,
	0xbb, 0xd2, 0x04, 0x35, 0xbc, 0x3e, 0x36, 0x8a, 0x15, 0x5f, 0x45, 0x8a, 0x25, 0xd8, 0xf3, 0x26,
	0xdd, 0x12, 0xc1, 0xdf, 0x15, 0x68, 0x63, 0x02, 0x43, 0xb6, 0xe2, 0x31, 0xd3, 0xc3, 0x4f, 0xa3,
	0x78, 0xce, 0x97, 0x6c, 0x20, 0xe3, 0xb9, 0x6d, 0xd9, 0x2e, 0xa5, 0x1b, 0x1a, 0x8b, 0x02, 0x55,
	0x33, 0xc1, 0x12, 0x92, 0x67, 0xd0, 0x14, 0x8b, 0x48, 0xdd, 0x66, 0x32, 0xf5, 0xeb, 0x28, 0x6d,
	0xb0, 0x1e, 0xfa, 0x32, 0x16, 0x85, 0xdf, 0x38, 0x76, 0xfa, 0x5d, 0x8a, 0xcf, 0xba, 0xa4, 0x94,
	0xa5, 0x99, 0x5c, 0xfb, 0xcd, 0x63, 0xa7, 0x5f, 0xa3, 0x16, 0xe9, 0x5f, 0xc8, 0x55, 0x26, 0xa3,
	0x19, 0xf3, 0x5b, 0x28, 0x94, 0x90, 0x9c, 0xe0, 0x30, 0x79, 0xcc, 0x72, 0x1f, 0x8e, 0xab, 0xfd,
	0xf6, 0xa9, 0x17, 0x3e, 0x98, 0x20, 0x2d, 0x03, 0xb6, 0x33, 0x68, 0xff, 0xcf, 0x0c, 0xc8, 0x4b,
	0x68, 0xe6, 0xd9, 0xad, 0xba, 0x8b, 0x24, 0xf3, 0x3b, 0x18, 0xdc, 0x0c, 0xed, 0x50, 0xe8, 0x46,
	0x21, 0xaf, 0xa0, 0xb1, 0x34, 0xb7, 0xd3, 0xef, 0xe2, 0x6f, 0x77, 0xc3, 0xdd, 0x2b, 0x4b, 0x4b,
	0x95, 0x7c, 0x03, 0x9d, 0xf2, 0xa5, 0x31, 0xcf, 0x95, 0xdf, 0xc3, 0xe8, 0xed, 0x91, 0x7b, 0x6a,
	0xf0, 0x09, 0x0e, 0x51, 0xb8, 0x58, 0x0b, 0x26, 0x57, 0x3c, 0xcf, 0xe4, 0xa6, 0x57, 0xd5, 0x47,
	0x7b, 0x55, 0xfb, 0x5c, 0xaf, 0xdc, 0xfd, 0x5e, 0xed, 0x56, 0x55, 0xff, 0x5c, 0x55, 0xc1, 0x6f,
	0x0e, 0x34, 0x91, 0x1d, 0x08, 0x41, 0x8e, 0xc0, 0x1d, 0x08, 0x31, 0x1a, 0x5a, 0xd3, 0x19, 0xb0,
	0x49, 0xa7, 0xf2, 0x68, 0x3a, 0x26, 0xc9, 0x47, 0xd2, 0xa9, 0xa1, 0xf0, 0x68, 0x3a, 0xee, 0x67,
	0xd3, 0xf9, 0xab, 0x4c, 0xe7, 0x32, 0x9f, 0x91, 0x2f, 0xf6, 0xcd, 0xd8, 0x36, 0xf1, 0x7b, 0x36,
	0x3c, 0x02, 0x37, 0x61, 0xab, 0x51, 0x62, 0xfd, 0x62, 0x00, 0x79, 0x09, 0x6e, 0xb2, 0x63, 0xbf,
	0x4e, 0xb8, 0x73, 0xc3, 0x2f, 0x0e, 0xa8, 0x11, 0xf5, 0x05, 0x99, 0x63, 0x54, 0x0d, 0xa3, 0xbc,
	0xf0, 0xc1, 0x1c, 0x74, 0x24, 0x06, 0xe8, 0x44, 0x22, 0x8c, 0x34, 0x89, 0xb7, 0xc2, 0xb2, 0x63,
	0x3a, 0x04, 0x95, 0xb3, 0x2e, 0xb4, 0x35, 0xf7, 0x43, 0xb6, 0x54, 0x6c, 0xa9, 0x82, 0x7f, 0x1c,
	0x68, 0xc5, 0xa2, 0xb8, 0x64, 0x4a, 0xf2, 0x58, 0x77, 0xab, 0x10, 0xd7, 0xdc, 0x2e, 0xd3, 0x2e,
	0xb5, 0x88, 0x7c, 0x0d, 0xbd, 0x58, 0x14, 0xef, 0x14, 0x5f, 0xf0, 0xfb, 0x48, 0x95, 0xeb, 0xd4,
	0xa1, 0x0f, 0x58, 0xe2, 0x41, 0xb5, 0xc8, 0xcd, 0x32, 0x75, 0xa8, 0x7e, 0xc4, 0x99, 0xf0, 0xd8,
	0x74, 0xd2, 0xa1, 0xf8, 0x8c, 0x1b, 0x62, 0x9d, 0x2b, 0x66, 0xcc, 0xe7, 0x50, 0x8b, 0x48, 0x0f,
	0x2a, 0x3c, 0x43, 0xe3, 0x39, 0xb4, 0xc2, 0x33, 0x7d, 0x1a, 0x97, 0x3f, 0xa3, 0xe7, 0x1c, 0xaa,
	0x1f, 0xf5, 0x69, 0x7a, 0x02, 0xe8, 0x36, 0x87, 0xe2, 0xb3, 0xee, 0x6c, 0xae, 0x58, 0xb4, 0xf0,
	0x01, 0x49, 0x03, 0x34, 0x3b, 0x2b, 0x58, 0xae, 0xd0, 0x54, 0x0e, 0x35, 0x40, 0xbf, 0xcf, 0x93,
	0x85, 0x31, 0x8f, 0x43, 0xf1, 0x39, 0xf8, 0xdd, 0x81, 0x8e, 0xb9, 0x14, 0xb6, 0x09, 0x3e, 0x34,
	0x8a, 0x9c, 0x25, 0x97, 0x2c, 0xb5, 0x5d, 0x28, 0xa1, 0xde, 0x1b, 0xd1, 0x2a, 0xe2, 0x0b, 0x2d,
	0x99, 0xeb, 0xb4, 0xc1, 0xba, 0x45, 0x3a, 0xec, 0x2d, 0x93, 0x31, 0x5b, 0xaa, 0xf2, 0x5e, 0x39,
	0xf4, 0x01, 0x4b, 0xfa, 0x70, 0x88, 0xef, 0xec, 0x04, 0x9a, 0xde, 0x3c, 0xa4, 0x83, 0x3f, 0x1d,
	0xe8, 0x5a, 0xab, 0xda, 0xcc, 0x8e, 0xc0, 0xe5, 0xf8, 0x81, 0xb0, 0xd7, 0x9e, 0x97, 0x1f, 0x39,
	0xf5, 0xf1, 0x6c, 0xad, 0x58, 0x8e, 0xf9, 0xd6, 0x68, 0x09, 0xb5, 0x22, 0xad, 0x52, 0x35, 0x8a,
	0xdc, 0x2a, 0xea, 0xe3, 0x50, 0x66, 0x22, 0xb7, 0x36, 0x2d, 0xa1, 0x79, 0xc7, 0x28, 0x6e, 0xf9,
	0x8e, 0x51, 0x9e, 0x42, 0x5d, 0x7d, 0xa4, 0x7a, 0x83, 0xd7, 0x8d, 0xb3, 0x0d, 0xd2, 0xbc, 0x34,
	0x7c, 0xc3, 0xf0, 0x06, 0x05, 0x9f, 0xa0, 0x63, 0x56, 0x9c, 0xcd, 0xfe, 0x39, 0x54, 0xb5, 0x3b,
	0x1d, 0xbc, 0x9a, 0x10, 0x6e, 0x6e, 0x1d, 0xd5, 0x34, 0xf9, 0x6a, 0x63, 0xd4, 0x0a, 0x06, 0x74,
	0xc3, 0xdd, 0xa1, 0x6c, 0x7c, 0xdb, 0xdf, 0x2e, 0xb7, 0x2a, 0xae, 0xab, 0x5e, 0xb8, 0xd7, 0xa3,
	0xcd, 0x76, 0x0b, 0xd6, 0x00, 0x53, 0xc3, 0x69, 0x8b, 0x5a, 0xff, 0x6d, 0x36, 0x06, 0x02, 0xf2,
	0x65, 0x69, 0xdc, 0x0a, 0x1a, 0xb7, 0x1b, 0x4e, 0x53, 0x7c, 0x63, 0xcf, 0xba, 0x2f, 0xa0, 0x92,
	0xa4, 0xd6, 0x7b, 0xdd, 0x70, 0xb7, 0xa4, 0x8b, 0x03, 0x5a, 0x49, 0xd2, 0xb3, 0x43, 0xe8, 0x1a,
	0x6c, 0x4d, 0x75, 0x72, 0x06, 0xb0, 0xdd, 0x00, 0xa4, 0x05, 0xee, 0x94, 0x4f, 0x32, 0xe1, 0x1d,
	0x90, 0x0e, 0x34, 0xa7, 0xdc, 0xd8, 0xdb, 0x73, 0x88, 0x07, 0x9d, 0x29, 0xdf, 0xda, 0xd8, 0xab,
	0x98, 0xd0, 0x81, 0x10, 0x5e, 0xf5, 0xe4, 0x3b, 0xf0, 0x1e, 0x7e, 0xd2, 0x09, 0x40, 0x7d, 0x2a,
	0x26, 0xd9, 0x92, 0x79, 0x07, 0xa4, 0x0b, 0xad, 0xa9, 0xb8, 0x32, 0x3b, 0xcb, 0x73, 0x0c, 0xb4,
	0x1b, 0xde, 0xab, 0x9c, 0xcc, 0xa1, 0x59, 0x7e, 0x5e, 0x49, 0x1b, 0x1a, 0xa3, 0xc9, 0xfb, 0xc1,
	0x78, 0x34, 0xf4, 0x0e, 0x0c, 0x18, 0x5d, 0x8f, 0x06, 0x63, 0xcf, 0x21, 0x47, 0xe0, 0x0d, 0xdf,
	0x7c, 0x98, 0x8c, 0xdf, 0x0c, 0x86, 0x3f, 0x5e, 0x5d, 0x0f, 0xe8, 0xf5, 0xf9, 0xd0, 0xab, 0x90,
	0x1e, 0x40, 0xc9, 0x9e, 0x0f, 0xbd, 0xaa, 0x3e, 0x7a, 0x78, 0x3e, 0x1e, 0xbd, 0x3f, 0xa7, 0xe7,
	0x43, 0xaf, 0xa6, 0xe1, 0x68, 0x72, 0x75, 0x3d, 0x18, 0x8f, 0xcf, 0x87, 0x9e, 0x7b, 0x72, 0x0a,
	0x9d, 0xdd, 0xa6, 0x61, 0x09, 0xe9, 0xb6, 0xda, 0x74, 0x53, 0x2d, 0x0a, 0x58, 0xdb, 0xd9, 0xf7,
	0xf0, 0x22, 0xce, 0xd2, 0xf0, 0x9e, 0x25, 0x2c, 0x89, 0xc2, 0x78, 0x91, 0x15, 0x49, 0x58, 0xe4,
	0xba, 0x0d, 0x31, 0x33, 0xff, 0x0a, 0xa7, 0xcf, 0x66, 0x5c, 0xcd, 0x8b, 0x9b, 0x30, 0xce, 0xd2,
	0xd7, 0x26, 0xee, 0x75, 0x24, 0xf8, 0x6b, 0xfd, 0xbf, 0xf1, 0xa6, 0x8e, 0x21, 0xdf, 0xfe, 0x1b,
	0x00, 0x00, 0xff, 0xff, 0x0c, 0x04, 0xa7, 0xb0, 0x46, 0x0a, 0x00, 0x00,
}