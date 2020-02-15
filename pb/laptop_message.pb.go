// Code generated by protoc-gen-go. DO NOT EDIT.
// source: laptop_message.proto

package pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Laptop struct {
	Id       string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Brand    string     `protobuf:"bytes,2,opt,name=brand,proto3" json:"brand,omitempty"`
	Name     string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Cpu      *CPU       `protobuf:"bytes,4,opt,name=cpu,proto3" json:"cpu,omitempty"`
	Memery   *Memory    `protobuf:"bytes,5,opt,name=memery,proto3" json:"memery,omitempty"`
	Gpus     []*GPU     `protobuf:"bytes,6,rep,name=gpus,proto3" json:"gpus,omitempty"`
	Storages []*Storage `protobuf:"bytes,7,rep,name=storages,proto3" json:"storages,omitempty"`
	Screen   *Screen    `protobuf:"bytes,8,opt,name=screen,proto3" json:"screen,omitempty"`
	Keyboard *Keyboard  `protobuf:"bytes,9,opt,name=keyboard,proto3" json:"keyboard,omitempty"`
	// Types that are valid to be assigned to Weight:
	//	*Laptop_WeightKg
	//	*Laptop_WeightLb
	Weight               isLaptop_Weight      `protobuf_oneof:"weight"`
	PriceUsd             float64              `protobuf:"fixed64,12,opt,name=price_usd,json=priceUsd,proto3" json:"price_usd,omitempty"`
	ReleaseYear          uint32               `protobuf:"varint,13,opt,name=release_year,json=releaseYear,proto3" json:"release_year,omitempty"`
	UpdatedAt            *timestamp.Timestamp `protobuf:"bytes,14,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Laptop) Reset()         { *m = Laptop{} }
func (m *Laptop) String() string { return proto.CompactTextString(m) }
func (*Laptop) ProtoMessage()    {}
func (*Laptop) Descriptor() ([]byte, []int) {
	return fileDescriptor_07a3824d46f4b731, []int{0}
}

func (m *Laptop) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Laptop.Unmarshal(m, b)
}
func (m *Laptop) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Laptop.Marshal(b, m, deterministic)
}
func (m *Laptop) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Laptop.Merge(m, src)
}
func (m *Laptop) XXX_Size() int {
	return xxx_messageInfo_Laptop.Size(m)
}
func (m *Laptop) XXX_DiscardUnknown() {
	xxx_messageInfo_Laptop.DiscardUnknown(m)
}

var xxx_messageInfo_Laptop proto.InternalMessageInfo

func (m *Laptop) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Laptop) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *Laptop) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Laptop) GetCpu() *CPU {
	if m != nil {
		return m.Cpu
	}
	return nil
}

func (m *Laptop) GetMemery() *Memory {
	if m != nil {
		return m.Memery
	}
	return nil
}

func (m *Laptop) GetGpus() []*GPU {
	if m != nil {
		return m.Gpus
	}
	return nil
}

func (m *Laptop) GetStorages() []*Storage {
	if m != nil {
		return m.Storages
	}
	return nil
}

func (m *Laptop) GetScreen() *Screen {
	if m != nil {
		return m.Screen
	}
	return nil
}

func (m *Laptop) GetKeyboard() *Keyboard {
	if m != nil {
		return m.Keyboard
	}
	return nil
}

type isLaptop_Weight interface {
	isLaptop_Weight()
}

type Laptop_WeightKg struct {
	WeightKg float64 `protobuf:"fixed64,10,opt,name=weight_kg,json=weightKg,proto3,oneof"`
}

type Laptop_WeightLb struct {
	WeightLb float64 `protobuf:"fixed64,11,opt,name=weight_lb,json=weightLb,proto3,oneof"`
}

func (*Laptop_WeightKg) isLaptop_Weight() {}

func (*Laptop_WeightLb) isLaptop_Weight() {}

func (m *Laptop) GetWeight() isLaptop_Weight {
	if m != nil {
		return m.Weight
	}
	return nil
}

func (m *Laptop) GetWeightKg() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightKg); ok {
		return x.WeightKg
	}
	return 0
}

func (m *Laptop) GetWeightLb() float64 {
	if x, ok := m.GetWeight().(*Laptop_WeightLb); ok {
		return x.WeightLb
	}
	return 0
}

func (m *Laptop) GetPriceUsd() float64 {
	if m != nil {
		return m.PriceUsd
	}
	return 0
}

func (m *Laptop) GetReleaseYear() uint32 {
	if m != nil {
		return m.ReleaseYear
	}
	return 0
}

func (m *Laptop) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Laptop) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Laptop_WeightKg)(nil),
		(*Laptop_WeightLb)(nil),
	}
}

func init() {
	proto.RegisterType((*Laptop)(nil), "techschool.pcbook.Laptop")
}

func init() { proto.RegisterFile("laptop_message.proto", fileDescriptor_07a3824d46f4b731) }

var fileDescriptor_07a3824d46f4b731 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0xd2, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0x07, 0x70, 0xd2, 0x66, 0x21, 0x79, 0xdd, 0x26, 0x61, 0x95, 0x61, 0x3a, 0x21, 0x02, 0xa7,
	0x88, 0x43, 0x26, 0x40, 0x02, 0x71, 0x64, 0x1c, 0x40, 0xda, 0x90, 0x26, 0x8f, 0x1e, 0xe0, 0x12,
	0x39, 0xc9, 0xc3, 0x8d, 0x1a, 0xd7, 0x96, 0xed, 0x08, 0xf5, 0x6b, 0xf1, 0x09, 0x51, 0x9d, 0x74,
	0x62, 0x5d, 0x7b, 0x8b, 0xff, 0xef, 0xe7, 0xe8, 0xf9, 0xe9, 0xc1, 0xb4, 0xe5, 0xda, 0x29, 0x5d,
	0x48, 0xb4, 0x96, 0x0b, 0xcc, 0xb5, 0x51, 0x4e, 0x91, 0x27, 0x0e, 0xab, 0x85, 0xad, 0x16, 0x4a,
	0xb5, 0xb9, 0xae, 0x4a, 0xa5, 0x96, 0xb3, 0x67, 0xda, 0xa8, 0x0a, 0xad, 0x55, 0xe6, 0xbe, 0x9d,
	0x4d, 0x25, 0x4a, 0x65, 0xd6, 0x3b, 0xe9, 0x53, 0xeb, 0x94, 0xe1, 0x02, 0x77, 0xb1, 0xad, 0x0c,
	0xe2, 0x6a, 0x27, 0x3d, 0x5b, 0xe2, 0xba, 0x54, 0xdc, 0xd4, 0x3b, 0xf9, 0x4b, 0xa1, 0x94, 0x68,
	0xf1, 0xc2, 0x9f, 0xca, 0xee, 0xf7, 0x85, 0x6b, 0x24, 0x5a, 0xc7, 0xa5, 0xee, 0xc1, 0xeb, 0xbf,
	0x21, 0x44, 0xd7, 0xfe, 0x01, 0xe4, 0x14, 0x46, 0x4d, 0x4d, 0x83, 0x34, 0xc8, 0x12, 0x36, 0x6a,
	0x6a, 0x32, 0x85, 0xa3, 0xd2, 0xf0, 0x55, 0x4d, 0x47, 0x3e, 0xea, 0x0f, 0x84, 0x40, 0xb8, 0xe2,
	0x12, 0xe9, 0xd8, 0x87, 0xfe, 0x9b, 0x64, 0x30, 0xae, 0x74, 0x47, 0xc3, 0x34, 0xc8, 0x26, 0xef,
	0xce, 0xf2, 0x07, 0x4f, 0xcf, 0xbf, 0xdc, 0xcc, 0xd9, 0x86, 0x90, 0xb7, 0x10, 0x49, 0x94, 0x68,
	0xd6, 0xf4, 0xc8, 0xe3, 0xe7, 0x7b, 0xf0, 0x77, 0x3f, 0x0d, 0x36, 0x40, 0xf2, 0x06, 0x42, 0xa1,
	0x3b, 0x4b, 0xa3, 0x74, 0x7c, 0xe0, 0xef, 0x5f, 0x6f, 0xe6, 0xcc, 0x1b, 0xf2, 0x01, 0xe2, 0x61,
	0x6a, 0x96, 0x3e, 0xf6, 0x7e, 0xb6, 0xc7, 0xdf, 0xf6, 0x84, 0xdd, 0xd9, 0x4d, 0x5b, 0xfd, 0x58,
	0x69, 0x7c, 0xb0, 0xad, 0x5b, 0x0f, 0xd8, 0x00, 0xc9, 0x47, 0x88, 0xb7, 0x33, 0xa7, 0x89, 0xbf,
	0x74, 0xbe, 0xe7, 0xd2, 0xd5, 0x40, 0xd8, 0x1d, 0x26, 0x2f, 0x20, 0xf9, 0x83, 0x8d, 0x58, 0xb8,
	0x62, 0x29, 0x28, 0xa4, 0x41, 0x16, 0x7c, 0x7b, 0xc4, 0xe2, 0x3e, 0xba, 0x12, 0xff, 0x95, 0xdb,
	0x92, 0x4e, 0xee, 0x97, 0xaf, 0x4b, 0x72, 0x0e, 0x89, 0x36, 0x4d, 0x85, 0x45, 0x67, 0x6b, 0x7a,
	0xbc, 0x29, 0xb3, 0xd8, 0x07, 0x73, 0x5b, 0x93, 0x57, 0x70, 0x6c, 0xb0, 0x45, 0x6e, 0xb1, 0x58,
	0x23, 0x37, 0xf4, 0x24, 0x0d, 0xb2, 0x13, 0x36, 0x19, 0xb2, 0x9f, 0xc8, 0x0d, 0xf9, 0x04, 0xd0,
	0xe9, 0x9a, 0x3b, 0xac, 0x0b, 0xee, 0xe8, 0xa9, 0x6f, 0x7c, 0x96, 0xf7, 0x5b, 0x92, 0x6f, 0xb7,
	0x24, 0xff, 0xb1, 0xdd, 0x12, 0x96, 0x0c, 0xfa, 0xb3, 0xbb, 0x8c, 0x21, 0xea, 0xdb, 0xb8, 0x0c,
	0x7f, 0x8d, 0x74, 0x59, 0x46, 0x9e, 0xbf, 0xff, 0x17, 0x00, 0x00, 0xff, 0xff, 0x61, 0x9f, 0x88,
	0xa0, 0x01, 0x03, 0x00, 0x00,
}
