// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.11
// source: execute.proto

package executor

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type ExecuteStatus int32

const (
	ExecuteStatus_SUCCESS ExecuteStatus = 0
	ExecuteStatus_FAILED  ExecuteStatus = 1
)

// Enum value maps for ExecuteStatus.
var (
	ExecuteStatus_name = map[int32]string{
		0: "SUCCESS",
		1: "FAILED",
	}
	ExecuteStatus_value = map[string]int32{
		"SUCCESS": 0,
		"FAILED":  1,
	}
)

func (x ExecuteStatus) Enum() *ExecuteStatus {
	p := new(ExecuteStatus)
	*p = x
	return p
}

func (x ExecuteStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ExecuteStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_execute_proto_enumTypes[0].Descriptor()
}

func (ExecuteStatus) Type() protoreflect.EnumType {
	return &file_execute_proto_enumTypes[0]
}

func (x ExecuteStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ExecuteStatus.Descriptor instead.
func (ExecuteStatus) EnumDescriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{0}
}

type ExecuteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Repo string `protobuf:"bytes,1,opt,name=repo,proto3" json:"repo,omitempty"`
	Cmd  string `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *ExecuteRequest) Reset() {
	*x = ExecuteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteRequest) ProtoMessage() {}

func (x *ExecuteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_execute_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteRequest.ProtoReflect.Descriptor instead.
func (*ExecuteRequest) Descriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{0}
}

func (x *ExecuteRequest) GetRepo() string {
	if x != nil {
		return x.Repo
	}
	return ""
}

func (x *ExecuteRequest) GetCmd() string {
	if x != nil {
		return x.Cmd
	}
	return ""
}

type ExecuteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   ExecuteStatus `protobuf:"varint,1,opt,name=status,proto3,enum=executor.ExecuteStatus" json:"status,omitempty"`
	ExitCode int32         `protobuf:"varint,2,opt,name=exit_code,json=exitCode,proto3" json:"exit_code,omitempty"`
	Output   string        `protobuf:"bytes,3,opt,name=output,proto3" json:"output,omitempty"`
	Time     float64       `protobuf:"fixed64,4,opt,name=time,proto3" json:"time,omitempty"`
}

func (x *ExecuteResponse) Reset() {
	*x = ExecuteResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_execute_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExecuteResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExecuteResponse) ProtoMessage() {}

func (x *ExecuteResponse) ProtoReflect() protoreflect.Message {
	mi := &file_execute_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExecuteResponse.ProtoReflect.Descriptor instead.
func (*ExecuteResponse) Descriptor() ([]byte, []int) {
	return file_execute_proto_rawDescGZIP(), []int{1}
}

func (x *ExecuteResponse) GetStatus() ExecuteStatus {
	if x != nil {
		return x.Status
	}
	return ExecuteStatus_SUCCESS
}

func (x *ExecuteResponse) GetExitCode() int32 {
	if x != nil {
		return x.ExitCode
	}
	return 0
}

func (x *ExecuteResponse) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

func (x *ExecuteResponse) GetTime() float64 {
	if x != nil {
		return x.Time
	}
	return 0
}

var File_execute_proto protoreflect.FileDescriptor

var file_execute_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x08, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x1a, 0x3f, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62, 0x75, 0x66, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x48, 0x0a, 0x0e, 0x45, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x04,
	0x72, 0x65, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x04, 0x72, 0x65, 0x70, 0x6f, 0x12, 0x19, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52,
	0x03, 0x63, 0x6d, 0x64, 0x22, 0x8b, 0x01, 0x0a, 0x0f, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x65, 0x78, 0x65, 0x63, 0x75,
	0x74, 0x6f, 0x72, 0x2e, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x65, 0x78, 0x69,
	0x74, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x65, 0x78,
	0x69, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12,
	0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x74, 0x69,
	0x6d, 0x65, 0x2a, 0x28, 0x0a, 0x0d, 0x45, 0x78, 0x65, 0x63, 0x75, 0x74, 0x65, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x00,
	0x12, 0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x42, 0x0e, 0x5a, 0x0c,
	0x70, 0x6b, 0x67, 0x2f, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x6f, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_execute_proto_rawDescOnce sync.Once
	file_execute_proto_rawDescData = file_execute_proto_rawDesc
)

func file_execute_proto_rawDescGZIP() []byte {
	file_execute_proto_rawDescOnce.Do(func() {
		file_execute_proto_rawDescData = protoimpl.X.CompressGZIP(file_execute_proto_rawDescData)
	})
	return file_execute_proto_rawDescData
}

var file_execute_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_execute_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_execute_proto_goTypes = []interface{}{
	(ExecuteStatus)(0),      // 0: executor.ExecuteStatus
	(*ExecuteRequest)(nil),  // 1: executor.ExecuteRequest
	(*ExecuteResponse)(nil), // 2: executor.ExecuteResponse
}
var file_execute_proto_depIdxs = []int32{
	0, // 0: executor.ExecuteResponse.status:type_name -> executor.ExecuteStatus
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_execute_proto_init() }
func file_execute_proto_init() {
	if File_execute_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_execute_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteRequest); i {
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
		file_execute_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExecuteResponse); i {
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
			RawDescriptor: file_execute_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_execute_proto_goTypes,
		DependencyIndexes: file_execute_proto_depIdxs,
		EnumInfos:         file_execute_proto_enumTypes,
		MessageInfos:      file_execute_proto_msgTypes,
	}.Build()
	File_execute_proto = out.File
	file_execute_proto_rawDesc = nil
	file_execute_proto_goTypes = nil
	file_execute_proto_depIdxs = nil
}
