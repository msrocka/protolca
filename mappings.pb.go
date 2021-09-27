// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: mappings.proto

package protolca

import (
	proto "github.com/golang/protobuf/proto"
	empty "github.com/golang/protobuf/ptypes/empty"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type FlowMapInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *FlowMapInfo) Reset() {
	*x = FlowMapInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mappings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowMapInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowMapInfo) ProtoMessage() {}

func (x *FlowMapInfo) ProtoReflect() protoreflect.Message {
	mi := &file_mappings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowMapInfo.ProtoReflect.Descriptor instead.
func (*FlowMapInfo) Descriptor() ([]byte, []int) {
	return file_mappings_proto_rawDescGZIP(), []int{0}
}

func (x *FlowMapInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_mappings_proto protoreflect.FileDescriptor

var file_mappings_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x73, 0x1a, 0x0a, 0x6f, 0x6c, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x21, 0x0a, 0x0b,
	0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x32,
	0x82, 0x02, 0x0a, 0x0e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x12, 0x1e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73,
	0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x38, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x1e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6c, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x2e,
	0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f, 0x1a, 0x11, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x6c, 0x63, 0x61, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x12, 0x42,
	0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x1e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x63, 0x61, 0x2e, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x73, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x49, 0x6e, 0x66, 0x6f,
	0x30, 0x01, 0x12, 0x30, 0x0a, 0x03, 0x50, 0x75, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x6c, 0x63, 0x61, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x4d, 0x61, 0x70, 0x1a, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x42, 0x49, 0x0a, 0x16, 0x6f, 0x72, 0x67, 0x2e, 0x6f, 0x70, 0x65, 0x6e,
	0x6c, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x42, 0x0d,
	0x4d, 0x61, 0x70, 0x70, 0x69, 0x6e, 0x67, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x0a, 0x2e, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x6c, 0x63, 0x61, 0xaa, 0x02, 0x11, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x4c, 0x43, 0x41, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x73, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mappings_proto_rawDescOnce sync.Once
	file_mappings_proto_rawDescData = file_mappings_proto_rawDesc
)

func file_mappings_proto_rawDescGZIP() []byte {
	file_mappings_proto_rawDescOnce.Do(func() {
		file_mappings_proto_rawDescData = protoimpl.X.CompressGZIP(file_mappings_proto_rawDescData)
	})
	return file_mappings_proto_rawDescData
}

var file_mappings_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_mappings_proto_goTypes = []interface{}{
	(*FlowMapInfo)(nil), // 0: protolca.services.FlowMapInfo
	(*empty.Empty)(nil), // 1: google.protobuf.Empty
	(*FlowMap)(nil),     // 2: protolca.FlowMap
}
var file_mappings_proto_depIdxs = []int32{
	0, // 0: protolca.services.FlowMapService.Delete:input_type -> protolca.services.FlowMapInfo
	0, // 1: protolca.services.FlowMapService.Get:input_type -> protolca.services.FlowMapInfo
	1, // 2: protolca.services.FlowMapService.GetAll:input_type -> google.protobuf.Empty
	2, // 3: protolca.services.FlowMapService.Put:input_type -> protolca.FlowMap
	1, // 4: protolca.services.FlowMapService.Delete:output_type -> google.protobuf.Empty
	2, // 5: protolca.services.FlowMapService.Get:output_type -> protolca.FlowMap
	0, // 6: protolca.services.FlowMapService.GetAll:output_type -> protolca.services.FlowMapInfo
	1, // 7: protolca.services.FlowMapService.Put:output_type -> google.protobuf.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mappings_proto_init() }
func file_mappings_proto_init() {
	if File_mappings_proto != nil {
		return
	}
	file_olca_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_mappings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowMapInfo); i {
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
			RawDescriptor: file_mappings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mappings_proto_goTypes,
		DependencyIndexes: file_mappings_proto_depIdxs,
		MessageInfos:      file_mappings_proto_msgTypes,
	}.Build()
	File_mappings_proto = out.File
	file_mappings_proto_rawDesc = nil
	file_mappings_proto_goTypes = nil
	file_mappings_proto_depIdxs = nil
}
