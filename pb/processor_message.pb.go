// Code generated by protoc-gen-go. DO NOT EDIT.
// source: processor_message.proto

package pb

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

type CPU struct {
	// Brand of the CPU
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	NumberCores          uint32   `protobuf:"varint,3,opt,name=number_cores,json=numberCores,proto3" json:"number_cores,omitempty"`
	NumberThreads        uint32   `protobuf:"varint,4,opt,name=number_threads,json=numberThreads,proto3" json:"number_threads,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,5,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,6,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CPU) Reset()         { *m = CPU{} }
func (m *CPU) String() string { return proto.CompactTextString(m) }
func (*CPU) ProtoMessage()    {}
func (*CPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{0}
}

func (m *CPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CPU.Unmarshal(m, b)
}
func (m *CPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CPU.Marshal(b, m, deterministic)
}
func (m *CPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CPU.Merge(m, src)
}
func (m *CPU) XXX_Size() int {
	return xxx_messageInfo_CPU.Size(m)
}
func (m *CPU) XXX_DiscardUnknown() {
	xxx_messageInfo_CPU.DiscardUnknown(m)
}

var xxx_messageInfo_CPU proto.InternalMessageInfo

func (m *CPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *CPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *CPU) GetNumberCores() uint32 {
	if m != nil {
		return m.NumberCores
	}
	return 0
}

func (m *CPU) GetNumberThreads() uint32 {
	if m != nil {
		return m.NumberThreads
	}
	return 0
}

func (m *CPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *CPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

type GPU struct {
	Brand                string   `protobuf:"bytes,1,opt,name=brand,proto3" json:"brand,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	MinGhz               float64  `protobuf:"fixed64,3,opt,name=min_ghz,json=minGhz,proto3" json:"min_ghz,omitempty"`
	MaxGhz               float64  `protobuf:"fixed64,4,opt,name=max_ghz,json=maxGhz,proto3" json:"max_ghz,omitempty"`
	Memory               *Memory  `protobuf:"bytes,5,opt,name=memory,proto3" json:"memory,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GPU) Reset()         { *m = GPU{} }
func (m *GPU) String() string { return proto.CompactTextString(m) }
func (*GPU) ProtoMessage()    {}
func (*GPU) Descriptor() ([]byte, []int) {
	return fileDescriptor_466578cecc6db379, []int{1}
}

func (m *GPU) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GPU.Unmarshal(m, b)
}
func (m *GPU) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GPU.Marshal(b, m, deterministic)
}
func (m *GPU) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GPU.Merge(m, src)
}
func (m *GPU) XXX_Size() int {
	return xxx_messageInfo_GPU.Size(m)
}
func (m *GPU) XXX_DiscardUnknown() {
	xxx_messageInfo_GPU.DiscardUnknown(m)
}

var xxx_messageInfo_GPU proto.InternalMessageInfo

func (m *GPU) GetBrand() string {
	if m != nil {
		return m.Brand
	}
	return ""
}

func (m *GPU) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *GPU) GetMinGhz() float64 {
	if m != nil {
		return m.MinGhz
	}
	return 0
}

func (m *GPU) GetMaxGhz() float64 {
	if m != nil {
		return m.MaxGhz
	}
	return 0
}

func (m *GPU) GetMemory() *Memory {
	if m != nil {
		return m.Memory
	}
	return nil
}

func init() {
	proto.RegisterType((*CPU)(nil), "techschool.pcbook.CPU")
	proto.RegisterType((*GPU)(nil), "techschool.pcbook.GPU")
}

func init() { proto.RegisterFile("processor_message.proto", fileDescriptor_466578cecc6db379) }

var fileDescriptor_466578cecc6db379 = []byte{
	// 271 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0xc1, 0x4a, 0xf4, 0x30,
	0x14, 0x85, 0xc9, 0xb4, 0xd3, 0x9f, 0xff, 0x8e, 0x23, 0x18, 0x06, 0xa6, 0xba, 0xb1, 0x0e, 0x88,
	0x5d, 0x15, 0xd4, 0x37, 0x70, 0x16, 0x5d, 0x09, 0x43, 0xd1, 0x8d, 0x9b, 0x92, 0x64, 0x42, 0x53,
	0x9c, 0xf4, 0x86, 0xa4, 0xc2, 0xd8, 0xe7, 0xf0, 0x2d, 0x7c, 0x49, 0x31, 0xe9, 0x62, 0x44, 0x5d,
	0xb8, 0x4b, 0xbe, 0xf3, 0xc1, 0xb9, 0x1c, 0x58, 0x1a, 0x8b, 0x42, 0x3a, 0x87, 0xb6, 0xd6, 0xd2,
	0x39, 0xd6, 0xc8, 0xc2, 0x58, 0xec, 0x91, 0x9e, 0xf4, 0x52, 0x28, 0x27, 0x14, 0xe2, 0xae, 0x30,
	0x82, 0x23, 0x3e, 0x9f, 0x2d, 0xb4, 0xd4, 0x68, 0x5f, 0xbf, 0x8a, 0xab, 0x77, 0x02, 0xd1, 0x7a,
	0xf3, 0x48, 0x17, 0x30, 0xe5, 0x96, 0x75, 0xdb, 0x94, 0x64, 0x24, 0xff, 0x5f, 0x85, 0x0f, 0xa5,
	0x10, 0x77, 0x4c, 0xcb, 0x74, 0xe2, 0xa1, 0x7f, 0xd3, 0x0b, 0x38, 0xea, 0x5e, 0x34, 0x97, 0xb6,
	0x16, 0x68, 0xa5, 0x4b, 0xa3, 0x8c, 0xe4, 0xf3, 0x6a, 0x16, 0xd8, 0xfa, 0x13, 0xd1, 0x4b, 0x38,
	0x1e, 0x95, 0x5e, 0x59, 0xc9, 0xb6, 0x2e, 0x8d, 0xbd, 0x34, 0x0f, 0xf4, 0x21, 0x40, 0xba, 0x84,
	0x7f, 0xba, 0xed, 0xea, 0x46, 0x0d, 0xe9, 0x34, 0x23, 0x39, 0xa9, 0x12, 0xdd, 0x76, 0xa5, 0x1a,
	0x7c, 0xc0, 0xf6, 0x3e, 0x48, 0xc6, 0x80, 0xed, 0x4b, 0x35, 0xac, 0xde, 0x08, 0x44, 0xe5, 0x9f,
	0xae, 0x3d, 0xe8, 0x88, 0x7e, 0xeb, 0x88, 0x0f, 0x3b, 0xe8, 0x35, 0x24, 0x61, 0x29, 0x7f, 0xd4,
	0xec, 0xe6, 0xb4, 0xf8, 0xb6, 0x65, 0x71, 0xef, 0x85, 0x6a, 0x14, 0xef, 0xae, 0xe0, 0x5c, 0xa0,
	0x2e, 0x9a, 0xb6, 0xdf, 0x31, 0xfe, 0x83, 0x6e, 0xf8, 0x86, 0x3c, 0x4d, 0x0c, 0xe7, 0x89, 0x1f,
	0xfd, 0xf6, 0x23, 0x00, 0x00, 0xff, 0xff, 0x44, 0x53, 0x3b, 0x15, 0xb8, 0x01, 0x00, 0x00,
}
