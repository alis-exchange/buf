// Copyright 2020-2024 Buf Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        (unknown)
// source: buf/alpha/wasmplugin/v1/wasmplugin.proto

package wasmpluginv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WASM_ABI specifies the abi this module expects buf to provide.
type WasmABI int32

const (
	WasmABI_WASM_ABI_UNSPECIFIED WasmABI = 0
	// Plugins compiled against
	// https://github.com/WebAssembly/WASI/releases/tag/snapshot-01.
	WasmABI_WASM_ABI_WASI_SNAPSHOT_PREVIEW1 WasmABI = 1
	// Plugins compiled with `GOOS=js` and `GOARCH=wasm`.
	WasmABI_WASM_ABI_GOJS WasmABI = 2
)

// Enum value maps for WasmABI.
var (
	WasmABI_name = map[int32]string{
		0: "WASM_ABI_UNSPECIFIED",
		1: "WASM_ABI_WASI_SNAPSHOT_PREVIEW1",
		2: "WASM_ABI_GOJS",
	}
	WasmABI_value = map[string]int32{
		"WASM_ABI_UNSPECIFIED":            0,
		"WASM_ABI_WASI_SNAPSHOT_PREVIEW1": 1,
		"WASM_ABI_GOJS":                   2,
	}
)

func (x WasmABI) Enum() *WasmABI {
	p := new(WasmABI)
	*p = x
	return p
}

func (x WasmABI) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (WasmABI) Descriptor() protoreflect.EnumDescriptor {
	return file_buf_alpha_wasmplugin_v1_wasmplugin_proto_enumTypes[0].Descriptor()
}

func (WasmABI) Type() protoreflect.EnumType {
	return &file_buf_alpha_wasmplugin_v1_wasmplugin_proto_enumTypes[0]
}

func (x WasmABI) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use WasmABI.Descriptor instead.
func (WasmABI) EnumDescriptor() ([]byte, []int) {
	return file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescGZIP(), []int{0}
}

// ExecConfig describes buf specific extensions for a wasm plugin.  A ExecConfig may
// be encoded in a custom WASM section named ".bufplugin", see
// (https://webassembly.github.io/spec/core/binary/modules.html#binary-customsec)
// for more info.
type ExecConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// When no ABI is provided, buf will make a best effort guess based on
	// the functions the wasm module exports.
	WasmAbi WasmABI `protobuf:"varint,1,opt,name=wasm_abi,json=wasmAbi,proto3,enum=buf.alpha.wasmplugin.v1.WasmABI" json:"wasm_abi,omitempty"`
	// The arguments that should be passed when running this plugin. Useful
	// for interpreted languages where the main wasm bundle is only the
	// interpreter.
	Args []string `protobuf:"bytes,2,rep,name=args,proto3" json:"args,omitempty"`
	// Files that should be made available to the WASI fs when running this
	// plugin. Useful for interpreted languages where the main wasm bundle
	// is only the interpreter.
	Files []*File `protobuf:"bytes,3,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *ExecConfig) Reset() {
	*x = ExecConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecConfig) ProtoMessage() {}

func (x *ExecConfig) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecConfig.ProtoReflect.Descriptor instead.
func (*ExecConfig) Descriptor() ([]byte, []int) {
	return file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescGZIP(), []int{0}
}

func (x *ExecConfig) GetWasmAbi() WasmABI {
	if x != nil {
		return x.WasmAbi
	}
	return WasmABI_WASM_ABI_UNSPECIFIED
}

func (x *ExecConfig) GetArgs() []string {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ExecConfig) GetFiles() []*File {
	if x != nil {
		return x.Files
	}
	return nil
}

// File represents a file that must be made available to the wasi plugin.
type File struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Normalized path with `/` as directory separator.
	Path     string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Contents []byte `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
}

func (x *File) Reset() {
	*x = File{}
	if protoimpl.UnsafeEnabled {
		mi := &file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *File) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*File) ProtoMessage() {}

func (x *File) ProtoReflect() protoreflect.Message {
	mi := &file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use File.ProtoReflect.Descriptor instead.
func (*File) Descriptor() ([]byte, []int) {
	return file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescGZIP(), []int{1}
}

func (x *File) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *File) GetContents() []byte {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_buf_alpha_wasmplugin_v1_wasmplugin_proto protoreflect.FileDescriptor

var file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDesc = []byte{
	0x0a, 0x28, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x77, 0x61, 0x73, 0x6d,
	0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x17, 0x62, 0x75, 0x66, 0x2e,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x76, 0x31, 0x22, 0x92, 0x01, 0x0a, 0x0a, 0x45, 0x78, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x3b, 0x0a, 0x08, 0x77, 0x61, 0x73, 0x6d, 0x5f, 0x61, 0x62, 0x69, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x20, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x2e, 0x77, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x57,
	0x61, 0x73, 0x6d, 0x41, 0x42, 0x49, 0x52, 0x07, 0x77, 0x61, 0x73, 0x6d, 0x41, 0x62, 0x69, 0x12,
	0x12, 0x0a, 0x04, 0x61, 0x72, 0x67, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x12, 0x33, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x77,
	0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x46, 0x69, 0x6c,
	0x65, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73,
	0x2a, 0x5b, 0x0a, 0x07, 0x57, 0x61, 0x73, 0x6d, 0x41, 0x42, 0x49, 0x12, 0x18, 0x0a, 0x14, 0x57,
	0x41, 0x53, 0x4d, 0x5f, 0x41, 0x42, 0x49, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46,
	0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x23, 0x0a, 0x1f, 0x57, 0x41, 0x53, 0x4d, 0x5f, 0x41, 0x42,
	0x49, 0x5f, 0x57, 0x41, 0x53, 0x49, 0x5f, 0x53, 0x4e, 0x41, 0x50, 0x53, 0x48, 0x4f, 0x54, 0x5f,
	0x50, 0x52, 0x45, 0x56, 0x49, 0x45, 0x57, 0x31, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x57, 0x41,
	0x53, 0x4d, 0x5f, 0x41, 0x42, 0x49, 0x5f, 0x47, 0x4f, 0x4a, 0x53, 0x10, 0x02, 0x42, 0x80, 0x02,
	0x0a, 0x1b, 0x63, 0x6f, 0x6d, 0x2e, 0x62, 0x75, 0x66, 0x2e, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2e,
	0x77, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x42, 0x0f, 0x57,
	0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x51, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66,
	0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x62, 0x75, 0x66, 0x2f, 0x70, 0x72, 0x69, 0x76, 0x61, 0x74,
	0x65, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x62,
	0x75, 0x66, 0x2f, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x2f, 0x77, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x77, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69,
	0x6e, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x42, 0x41, 0x57, 0xaa, 0x02, 0x17, 0x42, 0x75, 0x66, 0x2e,
	0x41, 0x6c, 0x70, 0x68, 0x61, 0x2e, 0x57, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e,
	0x2e, 0x56, 0x31, 0xca, 0x02, 0x17, 0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c,
	0x57, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x23,
	0x42, 0x75, 0x66, 0x5c, 0x41, 0x6c, 0x70, 0x68, 0x61, 0x5c, 0x57, 0x61, 0x73, 0x6d, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x42, 0x75, 0x66, 0x3a, 0x3a, 0x41, 0x6c, 0x70, 0x68, 0x61,
	0x3a, 0x3a, 0x57, 0x61, 0x73, 0x6d, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x3a, 0x3a, 0x56, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescOnce sync.Once
	file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescData = file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDesc
)

func file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescGZIP() []byte {
	file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescOnce.Do(func() {
		file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescData)
	})
	return file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDescData
}

var file_buf_alpha_wasmplugin_v1_wasmplugin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_buf_alpha_wasmplugin_v1_wasmplugin_proto_goTypes = []interface{}{
	(WasmABI)(0),       // 0: buf.alpha.wasmplugin.v1.WasmABI
	(*ExecConfig)(nil), // 1: buf.alpha.wasmplugin.v1.ExecConfig
	(*File)(nil),       // 2: buf.alpha.wasmplugin.v1.File
}
var file_buf_alpha_wasmplugin_v1_wasmplugin_proto_depIdxs = []int32{
	0, // 0: buf.alpha.wasmplugin.v1.ExecConfig.wasm_abi:type_name -> buf.alpha.wasmplugin.v1.WasmABI
	2, // 1: buf.alpha.wasmplugin.v1.ExecConfig.files:type_name -> buf.alpha.wasmplugin.v1.File
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_buf_alpha_wasmplugin_v1_wasmplugin_proto_init() }
func file_buf_alpha_wasmplugin_v1_wasmplugin_proto_init() {
	if File_buf_alpha_wasmplugin_v1_wasmplugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecConfig); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*File); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_buf_alpha_wasmplugin_v1_wasmplugin_proto_goTypes,
		DependencyIndexes: file_buf_alpha_wasmplugin_v1_wasmplugin_proto_depIdxs,
		EnumInfos:         file_buf_alpha_wasmplugin_v1_wasmplugin_proto_enumTypes,
		MessageInfos:      file_buf_alpha_wasmplugin_v1_wasmplugin_proto_msgTypes,
	}.Build()
	File_buf_alpha_wasmplugin_v1_wasmplugin_proto = out.File
	file_buf_alpha_wasmplugin_v1_wasmplugin_proto_rawDesc = nil
	file_buf_alpha_wasmplugin_v1_wasmplugin_proto_goTypes = nil
	file_buf_alpha_wasmplugin_v1_wasmplugin_proto_depIdxs = nil
}
