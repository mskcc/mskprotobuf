// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.1
// source: smile/smile.proto

package smile

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

type TempoSampleUpdateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TempoSamples []*TempoSample `protobuf:"bytes,1,rep,name=tempoSamples,proto3" json:"tempoSamples,omitempty"`
}

func (x *TempoSampleUpdateMessage) Reset() {
	*x = TempoSampleUpdateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smile_smile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TempoSampleUpdateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TempoSampleUpdateMessage) ProtoMessage() {}

func (x *TempoSampleUpdateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_smile_smile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TempoSampleUpdateMessage.ProtoReflect.Descriptor instead.
func (*TempoSampleUpdateMessage) Descriptor() ([]byte, []int) {
	return file_smile_smile_proto_rawDescGZIP(), []int{0}
}

func (x *TempoSampleUpdateMessage) GetTempoSamples() []*TempoSample {
	if x != nil {
		return x.TempoSamples
	}
	return nil
}

type TempoSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IgoId         string `protobuf:"bytes,1,opt,name=igoId,proto3" json:"igoId,omitempty"`
	CmoId         string `protobuf:"bytes,2,opt,name=cmoId,proto3" json:"cmoId,omitempty"`
	EmbargoStatus bool   `protobuf:"varint,3,opt,name=embargoStatus,proto3" json:"embargoStatus,omitempty"`
	Custodian     string `protobuf:"bytes,4,opt,name=custodian,proto3" json:"custodian,omitempty"`
}

func (x *TempoSample) Reset() {
	*x = TempoSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_smile_smile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TempoSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TempoSample) ProtoMessage() {}

func (x *TempoSample) ProtoReflect() protoreflect.Message {
	mi := &file_smile_smile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TempoSample.ProtoReflect.Descriptor instead.
func (*TempoSample) Descriptor() ([]byte, []int) {
	return file_smile_smile_proto_rawDescGZIP(), []int{1}
}

func (x *TempoSample) GetIgoId() string {
	if x != nil {
		return x.IgoId
	}
	return ""
}

func (x *TempoSample) GetCmoId() string {
	if x != nil {
		return x.CmoId
	}
	return ""
}

func (x *TempoSample) GetEmbargoStatus() bool {
	if x != nil {
		return x.EmbargoStatus
	}
	return false
}

func (x *TempoSample) GetCustodian() string {
	if x != nil {
		return x.Custodian
	}
	return ""
}

var File_smile_smile_proto protoreflect.FileDescriptor

var file_smile_smile_proto_rawDesc = []byte{
	0x0a, 0x11, 0x73, 0x6d, 0x69, 0x6c, 0x65, 0x2f, 0x73, 0x6d, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x73, 0x6d, 0x69, 0x6c, 0x65, 0x22, 0x52, 0x0a, 0x18, 0x54, 0x65,
	0x6d, 0x70, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x73,
	0x6d, 0x69, 0x6c, 0x65, 0x2e, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65,
	0x52, 0x0c, 0x74, 0x65, 0x6d, 0x70, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x7d,
	0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70, 0x6f, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x69, 0x67, 0x6f, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x69, 0x67,
	0x6f, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6d, 0x6f, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x63, 0x6d, 0x6f, 0x49, 0x64, 0x12, 0x24, 0x0a, 0x0d, 0x65, 0x6d, 0x62,
	0x61, 0x72, 0x67, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x0d, 0x65, 0x6d, 0x62, 0x61, 0x72, 0x67, 0x6f, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x1c, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x64, 0x69, 0x61, 0x6e, 0x42, 0x2a, 0x5a,
	0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6d, 0x73, 0x6b, 0x63,
	0x63, 0x2f, 0x6d, 0x73, 0x6b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x73, 0x6d,
	0x69, 0x6c, 0x65, 0x3b, 0x73, 0x6d, 0x69, 0x6c, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_smile_smile_proto_rawDescOnce sync.Once
	file_smile_smile_proto_rawDescData = file_smile_smile_proto_rawDesc
)

func file_smile_smile_proto_rawDescGZIP() []byte {
	file_smile_smile_proto_rawDescOnce.Do(func() {
		file_smile_smile_proto_rawDescData = protoimpl.X.CompressGZIP(file_smile_smile_proto_rawDescData)
	})
	return file_smile_smile_proto_rawDescData
}

var file_smile_smile_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_smile_smile_proto_goTypes = []interface{}{
	(*TempoSampleUpdateMessage)(nil), // 0: smile.TempoSampleUpdateMessage
	(*TempoSample)(nil),              // 1: smile.TempoSample
}
var file_smile_smile_proto_depIdxs = []int32{
	1, // 0: smile.TempoSampleUpdateMessage.tempoSamples:type_name -> smile.TempoSample
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_smile_smile_proto_init() }
func file_smile_smile_proto_init() {
	if File_smile_smile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_smile_smile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TempoSampleUpdateMessage); i {
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
		file_smile_smile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TempoSample); i {
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
			RawDescriptor: file_smile_smile_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_smile_smile_proto_goTypes,
		DependencyIndexes: file_smile_smile_proto_depIdxs,
		MessageInfos:      file_smile_smile_proto_msgTypes,
	}.Build()
	File_smile_smile_proto = out.File
	file_smile_smile_proto_rawDesc = nil
	file_smile_smile_proto_goTypes = nil
	file_smile_smile_proto_depIdxs = nil
}
