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
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xc3, 0x30,
	0x18, 0x86, 0xc9, 0xda, 0x55, 0xfc, 0xe6, 0x04, 0xc3, 0x60, 0xd5, 0x53, 0x1d, 0x08, 0x3d, 0x15,
	0xd4, 0x7f, 0xe0, 0x0e, 0x3d, 0x09, 0x52, 0xf4, 0xe2, 0xa5, 0x24, 0xd9, 0xc7, 0x32, 0x34, 0xf9,
	0x4a, 0x52, 0x61, 0xf6, 0x77, 0xf8, 0x2f, 0xfc, 0x93, 0x62, 0xd2, 0xc3, 0x44, 0x3c, 0x78, 0x4b,
	0x9e, 0xf7, 0x81, 0xf7, 0xe3, 0x85, 0x65, 0xe7, 0x48, 0xa1, 0xf7, 0xe4, 0x5a, 0x83, 0xde, 0x8b,
	0x2d, 0x56, 0x9d, 0xa3, 0x9e, 0xf8, 0x59, 0x8f, 0x4a, 0x7b, 0xa5, 0x89, 0x5e, 0xab, 0x4e, 0x49,
	0xa2, 0x97, 0x8b, 0x85, 0x41, 0x43, 0xee, 0xfd, 0xa7, 0xb8, 0xfa, 0x64, 0x90, 0xac, 0x1f, 0x9e,
	0xf8, 0x02, 0xa6, 0xd2, 0x09, 0xbb, 0xc9, 0x59, 0xc1, 0xca, 0xe3, 0x26, 0x7e, 0x38, 0x87, 0xd4,
	0x0a, 0x83, 0xf9, 0x24, 0xc0, 0xf0, 0xe6, 0x97, 0x70, 0x62, 0xdf, 0x8c, 0x44, 0xd7, 0x2a, 0x72,
	0xe8, 0xf3, 0xa4, 0x60, 0xe5, 0xbc, 0x99, 0x45, 0xb6, 0xfe, 0x46, 0xfc, 0x0a, 0x4e, 0x47, 0xa5,
	0xd7, 0x0e, 0xc5, 0xc6, 0xe7, 0x69, 0x90, 0xe6, 0x91, 0x3e, 0x46, 0xc8, 0x97, 0x70, 0x64, 0x76,
	0xb6, 0xdd, 0xea, 0x21, 0x9f, 0x16, 0xac, 0x64, 0x4d, 0x66, 0x76, 0xb6, 0xd6, 0x43, 0x08, 0xc4,
	0x3e, 0x04, 0xd9, 0x18, 0x88, 0x7d, 0xad, 0x87, 0xd5, 0x07, 0x83, 0xa4, 0xfe, 0xd7, 0xb5, 0x07,
	0x1d, 0xc9, 0x5f, 0x1d, 0xe9, 0x61, 0x07, 0xbf, 0x86, 0x2c, 0x2e, 0x15, 0x8e, 0x9a, 0xdd, 0x9c,
	0x57, 0xbf, 0xb6, 0xac, 0xee, 0x83, 0xd0, 0x8c, 0xe2, 0x5d, 0xfa, 0x3c, 0xe9, 0xa4, 0xcc, 0xc2,
	0xa2, 0xb7, 0x5f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x5e, 0x46, 0xcf, 0xd6, 0x95, 0x01, 0x00, 0x00,
}