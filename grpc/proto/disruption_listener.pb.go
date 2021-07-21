// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.17.3
// source: disruption_listener.proto

package disruption_listener

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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

type DisruptionSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EndpointAlterations []*EndpointAlteration `protobuf:"bytes,1,rep,name=endpointAlterations,proto3" json:"endpointAlterations,omitempty"`
}

func (x *DisruptionSpec) Reset() {
	*x = DisruptionSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disruption_listener_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DisruptionSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DisruptionSpec) ProtoMessage() {}

func (x *DisruptionSpec) ProtoReflect() protoreflect.Message {
	mi := &file_disruption_listener_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DisruptionSpec.ProtoReflect.Descriptor instead.
func (*DisruptionSpec) Descriptor() ([]byte, []int) {
	return file_disruption_listener_proto_rawDescGZIP(), []int{0}
}

func (x *DisruptionSpec) GetEndpointAlterations() []*EndpointAlteration {
	if x != nil {
		return x.EndpointAlterations
	}
	return nil
}

type EndpointAlteration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TargetEndpoint   string `protobuf:"bytes,1,opt,name=targetEndpoint,proto3" json:"targetEndpoint,omitempty"`
	ErrorToReturn    string `protobuf:"bytes,2,opt,name=errorToReturn,proto3" json:"errorToReturn,omitempty"`
	OverrideToReturn string `protobuf:"bytes,3,opt,name=overrideToReturn,proto3" json:"overrideToReturn,omitempty"`
}

func (x *EndpointAlteration) Reset() {
	*x = EndpointAlteration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_disruption_listener_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndpointAlteration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndpointAlteration) ProtoMessage() {}

func (x *EndpointAlteration) ProtoReflect() protoreflect.Message {
	mi := &file_disruption_listener_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndpointAlteration.ProtoReflect.Descriptor instead.
func (*EndpointAlteration) Descriptor() ([]byte, []int) {
	return file_disruption_listener_proto_rawDescGZIP(), []int{1}
}

func (x *EndpointAlteration) GetTargetEndpoint() string {
	if x != nil {
		return x.TargetEndpoint
	}
	return ""
}

func (x *EndpointAlteration) GetErrorToReturn() string {
	if x != nil {
		return x.ErrorToReturn
	}
	return ""
}

func (x *EndpointAlteration) GetOverrideToReturn() string {
	if x != nil {
		return x.OverrideToReturn
	}
	return ""
}

var File_disruption_listener_proto protoreflect.FileDescriptor

var file_disruption_listener_proto_rawDesc = []byte{
	0x0a, 0x19, 0x64, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x64, 0x69, 0x73,
	0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6b, 0x0a,
	0x0e, 0x44, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x59, 0x0a, 0x13, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x6c, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x64,
	0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x6c, 0x74, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x13, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41,
	0x6c, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x8e, 0x01, 0x0a, 0x12, 0x45,
	0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x41, 0x6c, 0x74, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65, 0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f,
	0x69, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x74, 0x61, 0x72, 0x67, 0x65,
	0x74, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x54, 0x6f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x6f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x12,
	0x2a, 0x0a, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x72, 0x69, 0x64, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x74,
	0x75, 0x72, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x6f, 0x76, 0x65, 0x72, 0x72,
	0x69, 0x64, 0x65, 0x54, 0x6f, 0x52, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x32, 0xfd, 0x01, 0x0a, 0x12,
	0x44, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x4c, 0x69, 0x73, 0x74, 0x65, 0x6e,
	0x65, 0x72, 0x12, 0x4f, 0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x44, 0x69, 0x73, 0x72, 0x75, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x23, 0x2e, 0x64, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x73, 0x72, 0x75,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65, 0x63, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x44, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a,
	0x23, 0x2e, 0x64, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x65, 0x6e, 0x65, 0x72, 0x2e, 0x44, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x70, 0x65, 0x63, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0f, 0x43, 0x6c, 0x65, 0x61, 0x6e, 0x44,
	0x69, 0x73, 0x72, 0x75, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x42, 0x6f, 0x5a, 0x6d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x74, 0x61, 0x44, 0x6f,
	0x67, 0x2f, 0x64, 0x64, 0x2d, 0x67, 0x6f, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x73, 0x2f,
	0x63, 0x68, 0x61, 0x6f, 0x73, 0x2d, 0x74, 0x6f, 0x6f, 0x6c, 0x69, 0x6e, 0x67, 0x2f, 0x61, 0x70,
	0x70, 0x73, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x64, 0x6f,
	0x67, 0x66, 0x6f, 0x6f, 0x64, 0x2f, 0x63, 0x68, 0x61, 0x6f, 0x73, 0x2d, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x63, 0x65, 0x70, 0x74, 0x6f, 0x72, 0x73, 0x2f, 0x64, 0x69, 0x73, 0x72, 0x75, 0x70, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x65, 0x6e, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_disruption_listener_proto_rawDescOnce sync.Once
	file_disruption_listener_proto_rawDescData = file_disruption_listener_proto_rawDesc
)

func file_disruption_listener_proto_rawDescGZIP() []byte {
	file_disruption_listener_proto_rawDescOnce.Do(func() {
		file_disruption_listener_proto_rawDescData = protoimpl.X.CompressGZIP(file_disruption_listener_proto_rawDescData)
	})
	return file_disruption_listener_proto_rawDescData
}

var file_disruption_listener_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_disruption_listener_proto_goTypes = []interface{}{
	(*DisruptionSpec)(nil),     // 0: disruption_listener.DisruptionSpec
	(*EndpointAlteration)(nil), // 1: disruption_listener.EndpointAlteration
	(*emptypb.Empty)(nil),      // 2: google.protobuf.Empty
}
var file_disruption_listener_proto_depIdxs = []int32{
	1, // 0: disruption_listener.DisruptionSpec.endpointAlterations:type_name -> disruption_listener.EndpointAlteration
	0, // 1: disruption_listener.DisruptionListener.SendDisruption:input_type -> disruption_listener.DisruptionSpec
	2, // 2: disruption_listener.DisruptionListener.DisruptionStatus:input_type -> google.protobuf.Empty
	2, // 3: disruption_listener.DisruptionListener.CleanDisruption:input_type -> google.protobuf.Empty
	2, // 4: disruption_listener.DisruptionListener.SendDisruption:output_type -> google.protobuf.Empty
	0, // 5: disruption_listener.DisruptionListener.DisruptionStatus:output_type -> disruption_listener.DisruptionSpec
	2, // 6: disruption_listener.DisruptionListener.CleanDisruption:output_type -> google.protobuf.Empty
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_disruption_listener_proto_init() }
func file_disruption_listener_proto_init() {
	if File_disruption_listener_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_disruption_listener_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DisruptionSpec); i {
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
		file_disruption_listener_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndpointAlteration); i {
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
			RawDescriptor: file_disruption_listener_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_disruption_listener_proto_goTypes,
		DependencyIndexes: file_disruption_listener_proto_depIdxs,
		MessageInfos:      file_disruption_listener_proto_msgTypes,
	}.Build()
	File_disruption_listener_proto = out.File
	file_disruption_listener_proto_rawDesc = nil
	file_disruption_listener_proto_goTypes = nil
	file_disruption_listener_proto_depIdxs = nil
}
