// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: Replication/protobuf/communication.proto

package protobuf

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

// The request message containing the Servers's name.
type NewServerRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerName string `protobuf:"bytes,1,opt,name=ServerName,proto3" json:"ServerName,omitempty"`
}

func (x *NewServerRequest) Reset() {
	*x = NewServerRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_protobuf_communication_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewServerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewServerRequest) ProtoMessage() {}

func (x *NewServerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_protobuf_communication_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewServerRequest.ProtoReflect.Descriptor instead.
func (*NewServerRequest) Descriptor() ([]byte, []int) {
	return file_Replication_protobuf_communication_proto_rawDescGZIP(), []int{0}
}

func (x *NewServerRequest) GetServerName() string {
	if x != nil {
		return x.ServerName
	}
	return ""
}

// The response message containing if the addition was succesful
type NewServerReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewServerReply) Reset() {
	*x = NewServerReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_protobuf_communication_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewServerReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewServerReply) ProtoMessage() {}

func (x *NewServerReply) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_protobuf_communication_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewServerReply.ProtoReflect.Descriptor instead.
func (*NewServerReply) Descriptor() ([]byte, []int) {
	return file_Replication_protobuf_communication_proto_rawDescGZIP(), []int{1}
}

// The request message containing the Frontend's name.
type NewFrontEndRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FrontEndName string `protobuf:"bytes,1,opt,name=FrontEndName,proto3" json:"FrontEndName,omitempty"`
}

func (x *NewFrontEndRequest) Reset() {
	*x = NewFrontEndRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_protobuf_communication_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFrontEndRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFrontEndRequest) ProtoMessage() {}

func (x *NewFrontEndRequest) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_protobuf_communication_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFrontEndRequest.ProtoReflect.Descriptor instead.
func (*NewFrontEndRequest) Descriptor() ([]byte, []int) {
	return file_Replication_protobuf_communication_proto_rawDescGZIP(), []int{2}
}

func (x *NewFrontEndRequest) GetFrontEndName() string {
	if x != nil {
		return x.FrontEndName
	}
	return ""
}

// The response message containing if the addition was succesful
type NewFrontEndReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *NewFrontEndReply) Reset() {
	*x = NewFrontEndReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Replication_protobuf_communication_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewFrontEndReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewFrontEndReply) ProtoMessage() {}

func (x *NewFrontEndReply) ProtoReflect() protoreflect.Message {
	mi := &file_Replication_protobuf_communication_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewFrontEndReply.ProtoReflect.Descriptor instead.
func (*NewFrontEndReply) Descriptor() ([]byte, []int) {
	return file_Replication_protobuf_communication_proto_rawDescGZIP(), []int{3}
}

var File_Replication_protobuf_communication_proto protoreflect.FileDescriptor

var file_Replication_protobuf_communication_proto_rawDesc = []byte{
	0x0a, 0x28, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x32, 0x0a, 0x10, 0x4e, 0x65, 0x77,
	0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a,
	0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x10, 0x0a,
	0x0e, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x38, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x45, 0x6e,
	0x64, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x46, 0x72, 0x6f,
	0x6e, 0x74, 0x45, 0x6e, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x12, 0x0a, 0x10, 0x4e, 0x65, 0x77,
	0x46, 0x72, 0x6f, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xb1, 0x01,
	0x0a, 0x0b, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a,
	0x09, 0x4e, 0x65, 0x77, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0b,
	0x4e, 0x65, 0x77, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x12, 0x21, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e, 0x65, 0x77, 0x46,
	0x72, 0x6f, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4e,
	0x65, 0x77, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x45, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x00, 0x42, 0x36, 0x5a, 0x34, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4a, 0x61, 0x63, 0x6f, 0x62, 0x4d, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x4d, 0x69, 0x6e, 0x69, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x33,
	0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_Replication_protobuf_communication_proto_rawDescOnce sync.Once
	file_Replication_protobuf_communication_proto_rawDescData = file_Replication_protobuf_communication_proto_rawDesc
)

func file_Replication_protobuf_communication_proto_rawDescGZIP() []byte {
	file_Replication_protobuf_communication_proto_rawDescOnce.Do(func() {
		file_Replication_protobuf_communication_proto_rawDescData = protoimpl.X.CompressGZIP(file_Replication_protobuf_communication_proto_rawDescData)
	})
	return file_Replication_protobuf_communication_proto_rawDescData
}

var file_Replication_protobuf_communication_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_Replication_protobuf_communication_proto_goTypes = []interface{}{
	(*NewServerRequest)(nil),   // 0: communication.NewServerRequest
	(*NewServerReply)(nil),     // 1: communication.NewServerReply
	(*NewFrontEndRequest)(nil), // 2: communication.NewFrontEndRequest
	(*NewFrontEndReply)(nil),   // 3: communication.NewFrontEndReply
}
var file_Replication_protobuf_communication_proto_depIdxs = []int32{
	0, // 0: communication.Replication.NewServer:input_type -> communication.NewServerRequest
	2, // 1: communication.Replication.NewFrontEnd:input_type -> communication.NewFrontEndRequest
	1, // 2: communication.Replication.NewServer:output_type -> communication.NewServerReply
	3, // 3: communication.Replication.NewFrontEnd:output_type -> communication.NewFrontEndReply
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_Replication_protobuf_communication_proto_init() }
func file_Replication_protobuf_communication_proto_init() {
	if File_Replication_protobuf_communication_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Replication_protobuf_communication_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewServerRequest); i {
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
		file_Replication_protobuf_communication_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewServerReply); i {
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
		file_Replication_protobuf_communication_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewFrontEndRequest); i {
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
		file_Replication_protobuf_communication_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewFrontEndReply); i {
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
			RawDescriptor: file_Replication_protobuf_communication_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_Replication_protobuf_communication_proto_goTypes,
		DependencyIndexes: file_Replication_protobuf_communication_proto_depIdxs,
		MessageInfos:      file_Replication_protobuf_communication_proto_msgTypes,
	}.Build()
	File_Replication_protobuf_communication_proto = out.File
	file_Replication_protobuf_communication_proto_rawDesc = nil
	file_Replication_protobuf_communication_proto_goTypes = nil
	file_Replication_protobuf_communication_proto_depIdxs = nil
}
