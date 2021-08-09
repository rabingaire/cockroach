// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: storage/enginepb/engine.proto

package enginepb

import proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// EngineType specifies type of storage engine (eg. rocksdb, pebble).
type EngineType int32

const (
	// Denotes the default storage engine. Alias for EngineTypePebble.
	EngineTypeDefault EngineType = 0
	// Denotes Pebble as the underlying storage engine type.
	EngineTypePebble EngineType = 2
)

var EngineType_name = map[int32]string{
	0: "EngineTypeDefault",
	2: "EngineTypePebble",
}
var EngineType_value = map[string]int32{
	"EngineTypeDefault": 0,
	"EngineTypePebble":  2,
}

func (EngineType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_engine_c3997d9635965b2b, []int{0}
}

func init() {
	proto.RegisterEnum("cockroach.storage.enginepb.EngineType", EngineType_name, EngineType_value)
}

func init() {
	proto.RegisterFile("storage/enginepb/engine.proto", fileDescriptor_engine_c3997d9635965b2b)
}

var fileDescriptor_engine_c3997d9635965b2b = []byte{
	// 187 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2d, 0x2e, 0xc9, 0x2f,
	0x4a, 0x4c, 0x4f, 0xd5, 0x4f, 0xcd, 0x4b, 0xcf, 0xcc, 0x4b, 0x2d, 0x48, 0x82, 0x32, 0xf4, 0x0a,
	0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xa4, 0x92, 0xf3, 0x93, 0xb3, 0x8b, 0xf2, 0x13, 0x93, 0x33, 0xf4,
	0xa0, 0x0a, 0xf5, 0x60, 0x0a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0xca, 0xf4, 0x41, 0x2c,
	0x88, 0x0e, 0x2d, 0x7f, 0x2e, 0x2e, 0x57, 0xb0, 0x8a, 0x90, 0xca, 0x82, 0x54, 0x21, 0x51, 0x2e,
	0x41, 0x04, 0xcf, 0x25, 0x35, 0x2d, 0xb1, 0x34, 0xa7, 0x44, 0x80, 0x41, 0x48, 0x84, 0x4b, 0x00,
	0x21, 0x1c, 0x90, 0x9a, 0x94, 0x94, 0x93, 0x2a, 0xc0, 0x24, 0xc5, 0xd1, 0xb1, 0x58, 0x8e, 0x61,
	0xc5, 0x12, 0x39, 0x06, 0x25, 0x16, 0x0e, 0x46, 0x01, 0x46, 0x25, 0x16, 0x0e, 0x66, 0x01, 0x66,
	0x27, 0xad, 0x13, 0x0f, 0xe5, 0x18, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc6,
	0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1,
	0xc6, 0x63, 0x39, 0x86, 0x28, 0x0e, 0x98, 0x93, 0x92, 0xd8, 0xc0, 0x6e, 0x30, 0x06, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xe5, 0x3a, 0x5a, 0xee, 0xd6, 0x00, 0x00, 0x00,
}
