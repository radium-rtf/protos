// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: checker/checker.proto

package checker

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

type TestCase struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input  string `protobuf:"bytes,1,opt,name=input,proto3" json:"input,omitempty"`
	Output string `protobuf:"bytes,2,opt,name=output,proto3" json:"output,omitempty"`
}

func (x *TestCase) Reset() {
	*x = TestCase{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checker_checker_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCase) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCase) ProtoMessage() {}

func (x *TestCase) ProtoReflect() protoreflect.Message {
	mi := &file_checker_checker_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCase.ProtoReflect.Descriptor instead.
func (*TestCase) Descriptor() ([]byte, []int) {
	return file_checker_checker_proto_rawDescGZIP(), []int{0}
}

func (x *TestCase) GetInput() string {
	if x != nil {
		return x.Input
	}
	return ""
}

func (x *TestCase) GetOutput() string {
	if x != nil {
		return x.Output
	}
	return ""
}

type CheckRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code  string      `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Tests []*TestCase `protobuf:"bytes,2,rep,name=tests,proto3" json:"tests,omitempty"`
}

func (x *CheckRequest) Reset() {
	*x = CheckRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checker_checker_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckRequest) ProtoMessage() {}

func (x *CheckRequest) ProtoReflect() protoreflect.Message {
	mi := &file_checker_checker_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckRequest.ProtoReflect.Descriptor instead.
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return file_checker_checker_proto_rawDescGZIP(), []int{1}
}

func (x *CheckRequest) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *CheckRequest) GetTests() []*TestCase {
	if x != nil {
		return x.Tests
	}
	return nil
}

type CheckResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success    bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	TestNumber int64  `protobuf:"varint,2,opt,name=test_number,json=testNumber,proto3" json:"test_number,omitempty"`
	Msg        string `protobuf:"bytes,3,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *CheckResponse) Reset() {
	*x = CheckResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_checker_checker_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckResponse) ProtoMessage() {}

func (x *CheckResponse) ProtoReflect() protoreflect.Message {
	mi := &file_checker_checker_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckResponse.ProtoReflect.Descriptor instead.
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return file_checker_checker_proto_rawDescGZIP(), []int{2}
}

func (x *CheckResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *CheckResponse) GetTestNumber() int64 {
	if x != nil {
		return x.TestNumber
	}
	return 0
}

func (x *CheckResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

var File_checker_checker_proto protoreflect.FileDescriptor

var file_checker_checker_proto_rawDesc = []byte{
	0x0a, 0x15, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x22, 0x38, 0x0a, 0x08, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x69, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x6e, 0x70,
	0x75, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x6f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x22, 0x4b, 0x0a, 0x0c, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x27,
	0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x61, 0x73, 0x65,
	0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x22, 0x5c, 0x0a, 0x0d, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x4e, 0x75, 0x6d,
	0x62, 0x65, 0x72, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x32, 0x43, 0x0a, 0x07, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72,
	0x12, 0x38, 0x0a, 0x05, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x15, 0x2e, 0x63, 0x68, 0x65, 0x63,
	0x6b, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x18, 0x5a, 0x16, 0x67, 0x65,
	0x6e, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x72, 0x3b, 0x63, 0x68, 0x65,
	0x63, 0x6b, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_checker_checker_proto_rawDescOnce sync.Once
	file_checker_checker_proto_rawDescData = file_checker_checker_proto_rawDesc
)

func file_checker_checker_proto_rawDescGZIP() []byte {
	file_checker_checker_proto_rawDescOnce.Do(func() {
		file_checker_checker_proto_rawDescData = protoimpl.X.CompressGZIP(file_checker_checker_proto_rawDescData)
	})
	return file_checker_checker_proto_rawDescData
}

var file_checker_checker_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_checker_checker_proto_goTypes = []interface{}{
	(*TestCase)(nil),      // 0: checker.TestCase
	(*CheckRequest)(nil),  // 1: checker.CheckRequest
	(*CheckResponse)(nil), // 2: checker.CheckResponse
}
var file_checker_checker_proto_depIdxs = []int32{
	0, // 0: checker.CheckRequest.tests:type_name -> checker.TestCase
	1, // 1: checker.checker.Check:input_type -> checker.CheckRequest
	2, // 2: checker.checker.Check:output_type -> checker.CheckResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_checker_checker_proto_init() }
func file_checker_checker_proto_init() {
	if File_checker_checker_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_checker_checker_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TestCase); i {
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
		file_checker_checker_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckRequest); i {
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
		file_checker_checker_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckResponse); i {
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
			RawDescriptor: file_checker_checker_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_checker_checker_proto_goTypes,
		DependencyIndexes: file_checker_checker_proto_depIdxs,
		MessageInfos:      file_checker_checker_proto_msgTypes,
	}.Build()
	File_checker_checker_proto = out.File
	file_checker_checker_proto_rawDesc = nil
	file_checker_checker_proto_goTypes = nil
	file_checker_checker_proto_depIdxs = nil
}
