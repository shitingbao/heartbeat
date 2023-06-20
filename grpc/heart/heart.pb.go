// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: heart.proto

package heart

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

type Heart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Message []byte `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Heart) Reset() {
	*x = Heart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heart_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Heart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Heart) ProtoMessage() {}

func (x *Heart) ProtoReflect() protoreflect.Message {
	mi := &file_heart_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Heart.ProtoReflect.Descriptor instead.
func (*Heart) Descriptor() ([]byte, []int) {
	return file_heart_proto_rawDescGZIP(), []int{0}
}

func (x *Heart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Heart) GetMessage() []byte {
	if x != nil {
		return x.Message
	}
	return nil
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Content string `protobuf:"bytes,1,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_heart_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_heart_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_heart_proto_rawDescGZIP(), []int{1}
}

func (x *Result) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

var File_heart_proto protoreflect.FileDescriptor

var file_heart_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x68, 0x65, 0x61, 0x72, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x68,
	0x65, 0x61, 0x72, 0x74, 0x22, 0x31, 0x0a, 0x05, 0x48, 0x65, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x22, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x32, 0x3b, 0x0a, 0x0b, 0x48,
	0x65, 0x61, 0x72, 0x74, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x09, 0x48, 0x65,
	0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x12, 0x0c, 0x2e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x2e,
	0x48, 0x65, 0x61, 0x72, 0x74, 0x1a, 0x0d, 0x2e, 0x68, 0x65, 0x61, 0x72, 0x74, 0x2e, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x28, 0x01, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x68, 0x69, 0x74, 0x69, 0x6e, 0x67, 0x62, 0x61,
	0x6f, 0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x2f, 0x67, 0x72, 0x70, 0x63,
	0x2f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x3b, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_heart_proto_rawDescOnce sync.Once
	file_heart_proto_rawDescData = file_heart_proto_rawDesc
)

func file_heart_proto_rawDescGZIP() []byte {
	file_heart_proto_rawDescOnce.Do(func() {
		file_heart_proto_rawDescData = protoimpl.X.CompressGZIP(file_heart_proto_rawDescData)
	})
	return file_heart_proto_rawDescData
}

var file_heart_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_heart_proto_goTypes = []interface{}{
	(*Heart)(nil),  // 0: heart.Heart
	(*Result)(nil), // 1: heart.Result
}
var file_heart_proto_depIdxs = []int32{
	0, // 0: heart.HeartServer.HeartBeat:input_type -> heart.Heart
	1, // 1: heart.HeartServer.HeartBeat:output_type -> heart.Result
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_heart_proto_init() }
func file_heart_proto_init() {
	if File_heart_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_heart_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Heart); i {
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
		file_heart_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
			RawDescriptor: file_heart_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_heart_proto_goTypes,
		DependencyIndexes: file_heart_proto_depIdxs,
		MessageInfos:      file_heart_proto_msgTypes,
	}.Build()
	File_heart_proto = out.File
	file_heart_proto_rawDesc = nil
	file_heart_proto_goTypes = nil
	file_heart_proto_depIdxs = nil
}
