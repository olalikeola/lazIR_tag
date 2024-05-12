// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.25.1
// source: service.proto

package proto

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

type AddRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Shooter string `protobuf:"bytes,1,opt,name=shooter,proto3" json:"shooter,omitempty"`
	Victim  string `protobuf:"bytes,2,opt,name=victim,proto3" json:"victim,omitempty"`
}

func (x *AddRecordRequest) Reset() {
	*x = AddRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordRequest) ProtoMessage() {}

func (x *AddRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordRequest.ProtoReflect.Descriptor instead.
func (*AddRecordRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{0}
}

func (x *AddRecordRequest) GetShooter() string {
	if x != nil {
		return x.Shooter
	}
	return ""
}

func (x *AddRecordRequest) GetVictim() string {
	if x != nil {
		return x.Victim
	}
	return ""
}

type AddRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommitIndex uint64 `protobuf:"varint,1,opt,name=commit_index,json=commitIndex,proto3" json:"commit_index,omitempty"`
}

func (x *AddRecordResponse) Reset() {
	*x = AddRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRecordResponse) ProtoMessage() {}

func (x *AddRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRecordResponse.ProtoReflect.Descriptor instead.
func (*AddRecordResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{1}
}

func (x *AddRecordResponse) GetCommitIndex() uint64 {
	if x != nil {
		return x.CommitIndex
	}
	return 0
}

type GetRecordRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRecordRequest) Reset() {
	*x = GetRecordRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordRequest) ProtoMessage() {}

func (x *GetRecordRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordRequest.ProtoReflect.Descriptor instead.
func (*GetRecordRequest) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{2}
}

type GetRecordResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadAtIndex uint64           `protobuf:"varint,1,opt,name=read_at_index,json=readAtIndex,proto3" json:"read_at_index,omitempty"`
	Score       map[string]int32 `protobuf:"bytes,2,rep,name=score,proto3" json:"score,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
}

func (x *GetRecordResponse) Reset() {
	*x = GetRecordResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRecordResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRecordResponse) ProtoMessage() {}

func (x *GetRecordResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRecordResponse.ProtoReflect.Descriptor instead.
func (*GetRecordResponse) Descriptor() ([]byte, []int) {
	return file_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetRecordResponse) GetReadAtIndex() uint64 {
	if x != nil {
		return x.ReadAtIndex
	}
	return 0
}

func (x *GetRecordResponse) GetScore() map[string]int32 {
	if x != nil {
		return x.Score
	}
	return nil
}

var File_service_proto protoreflect.FileDescriptor

var file_service_proto_rawDesc = []byte{
	0x0a, 0x0d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x09, 0x6c, 0x61, 0x7a, 0x49, 0x52, 0x5f, 0x74, 0x61, 0x67, 0x22, 0x44, 0x0a, 0x10, 0x41, 0x64,
	0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x73, 0x68, 0x6f, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x69, 0x63, 0x74,
	0x69, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x76, 0x69, 0x63, 0x74, 0x69, 0x6d,
	0x22, 0x36, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x74, 0x5f,
	0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x63, 0x6f, 0x6d,
	0x6d, 0x69, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0xb0, 0x01, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x22, 0x0a, 0x0d, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x61, 0x74, 0x5f, 0x69, 0x6e,
	0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x41,
	0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x3d, 0x0a, 0x05, 0x73, 0x63, 0x6f, 0x72, 0x65, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x6c, 0x61, 0x7a, 0x49, 0x52, 0x5f, 0x74, 0x61,
	0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05,
	0x73, 0x63, 0x6f, 0x72, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32,
	0x9b, 0x01, 0x0a, 0x05, 0x53, 0x63, 0x6f, 0x72, 0x65, 0x12, 0x48, 0x0a, 0x09, 0x41, 0x64, 0x64,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x7a, 0x49, 0x52, 0x5f, 0x74,
	0x61, 0x67, 0x2e, 0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x6c, 0x61, 0x7a, 0x49, 0x52, 0x5f, 0x74, 0x61, 0x67, 0x2e,
	0x41, 0x64, 0x64, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64,
	0x12, 0x1b, 0x2e, 0x6c, 0x61, 0x7a, 0x49, 0x52, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6c, 0x61, 0x7a, 0x49, 0x52, 0x5f, 0x74, 0x61, 0x67, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63,
	0x6f, 0x72, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x61, 0x72, 0x6b,
	0x68, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x39, 0x2f, 0x6c, 0x61, 0x7a, 0x49, 0x52, 0x5f, 0x74, 0x61,
	0x67, 0x2f, 0x73, 0x72, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_service_proto_rawDescOnce sync.Once
	file_service_proto_rawDescData = file_service_proto_rawDesc
)

func file_service_proto_rawDescGZIP() []byte {
	file_service_proto_rawDescOnce.Do(func() {
		file_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_proto_rawDescData)
	})
	return file_service_proto_rawDescData
}

var file_service_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_service_proto_goTypes = []interface{}{
	(*AddRecordRequest)(nil),  // 0: lazIR_tag.AddRecordRequest
	(*AddRecordResponse)(nil), // 1: lazIR_tag.AddRecordResponse
	(*GetRecordRequest)(nil),  // 2: lazIR_tag.GetRecordRequest
	(*GetRecordResponse)(nil), // 3: lazIR_tag.GetRecordResponse
	nil,                       // 4: lazIR_tag.GetRecordResponse.ScoreEntry
}
var file_service_proto_depIdxs = []int32{
	4, // 0: lazIR_tag.GetRecordResponse.score:type_name -> lazIR_tag.GetRecordResponse.ScoreEntry
	0, // 1: lazIR_tag.Score.AddRecord:input_type -> lazIR_tag.AddRecordRequest
	2, // 2: lazIR_tag.Score.GetRecord:input_type -> lazIR_tag.GetRecordRequest
	1, // 3: lazIR_tag.Score.AddRecord:output_type -> lazIR_tag.AddRecordResponse
	3, // 4: lazIR_tag.Score.GetRecord:output_type -> lazIR_tag.GetRecordResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_proto_init() }
func file_service_proto_init() {
	if File_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecordRequest); i {
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
		file_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddRecordResponse); i {
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
		file_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordRequest); i {
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
		file_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRecordResponse); i {
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
			RawDescriptor: file_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_proto_goTypes,
		DependencyIndexes: file_service_proto_depIdxs,
		MessageInfos:      file_service_proto_msgTypes,
	}.Build()
	File_service_proto = out.File
	file_service_proto_rawDesc = nil
	file_service_proto_goTypes = nil
	file_service_proto_depIdxs = nil
}