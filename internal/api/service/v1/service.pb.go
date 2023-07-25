// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: service/v1/service.proto

package servicev1

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PingRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Env bool `protobuf:"varint,1,opt,name=env,proto3" json:"env,omitempty"`
}

func (x *PingRequest) Reset() {
	*x = PingRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_v1_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PingRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PingRequest) ProtoMessage() {}

func (x *PingRequest) ProtoReflect() protoreflect.Message {
	mi := &file_service_v1_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PingRequest.ProtoReflect.Descriptor instead.
func (*PingRequest) Descriptor() ([]byte, []int) {
	return file_service_v1_service_proto_rawDescGZIP(), []int{0}
}

func (x *PingRequest) GetEnv() bool {
	if x != nil {
		return x.Env
	}
	return false
}

type PongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Hostname   string            `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
	Ip         []string          `protobuf:"bytes,3,rep,name=ip,proto3" json:"ip,omitempty"`
	RemoteAddr string            `protobuf:"bytes,4,opt,name=remote_addr,json=remoteAddr,proto3" json:"remote_addr,omitempty"`
	Env        map[string]string `protobuf:"bytes,5,rep,name=env,proto3" json:"env,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *PongResponse) Reset() {
	*x = PongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_service_v1_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PongResponse) ProtoMessage() {}

func (x *PongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_service_v1_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PongResponse.ProtoReflect.Descriptor instead.
func (*PongResponse) Descriptor() ([]byte, []int) {
	return file_service_v1_service_proto_rawDescGZIP(), []int{1}
}

func (x *PongResponse) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PongResponse) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

func (x *PongResponse) GetIp() []string {
	if x != nil {
		return x.Ip
	}
	return nil
}

func (x *PongResponse) GetRemoteAddr() string {
	if x != nil {
		return x.RemoteAddr
	}
	return ""
}

func (x *PongResponse) GetEnv() map[string]string {
	if x != nil {
		return x.Env
	}
	return nil
}

var File_service_v1_service_proto protoreflect.FileDescriptor

var file_service_v1_service_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x12, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x1a, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61,
	0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1f,
	0x0a, 0x0b, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a,
	0x03, 0x65, 0x6e, 0x76, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x22,
	0xe4, 0x01, 0x0a, 0x0c, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70,
	0x12, 0x1f, 0x0a, 0x0b, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x41, 0x64, 0x64,
	0x72, 0x12, 0x3b, 0x0a, 0x03, 0x65, 0x6e, 0x76, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x2e, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x03, 0x65, 0x6e, 0x76, 0x1a, 0x36,
	0x0a, 0x08, 0x45, 0x6e, 0x76, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0x58, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4d, 0x0a, 0x06, 0x57, 0x68, 0x6f, 0x41, 0x6d, 0x49, 0x12, 0x1f, 0x2e, 0x73, 0x70,
	0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76, 0x31,
	0x2e, 0x50, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0xaa, 0x03, 0x92, 0x41, 0xd2, 0x01, 0x12, 0xcf, 0x01, 0x0a, 0x07, 0x53, 0x70, 0x61, 0x72,
	0x6b, 0x6c, 0x65, 0x12, 0x22, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x20, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x20, 0x41, 0x50, 0x49, 0x20, 0x53, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x1a, 0x15, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f,
	0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2f, 0x22, 0x36,
	0x0a, 0x0c, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x20, 0x74, 0x65, 0x61, 0x6d, 0x12, 0x15,
	0x68, 0x74, 0x74, 0x70, 0x73, 0x3a, 0x2f, 0x2f, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2e,
	0x77, 0x69, 0x6b, 0x69, 0x2f, 0x1a, 0x0f, 0x68, 0x69, 0x40, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x6c,
	0x65, 0x2e, 0x77, 0x69, 0x6b, 0x69, 0x2a, 0x4a, 0x0a, 0x0e, 0x41, 0x47, 0x50, 0x4c, 0x2d, 0x33,
	0x2e, 0x30, 0x2d, 0x6c, 0x61, 0x74, 0x65, 0x72, 0x12, 0x38, 0x68, 0x74, 0x74, 0x70, 0x73, 0x3a,
	0x2f, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x77, 0x69, 0x74,
	0x68, 0x73, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x2f, 0x62, 0x6c, 0x6f, 0x62, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2f, 0x4c, 0x49, 0x43, 0x45, 0x4e,
	0x53, 0x45, 0x32, 0x05, 0x30, 0x2e, 0x30, 0x2e, 0x31, 0x0a, 0x16, 0x63, 0x6f, 0x6d, 0x2e, 0x73,
	0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x42, 0x0c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x42, 0x67, 0x6f, 0x2e, 0x6f, 0x63, 0x74, 0x6f, 0x6c, 0x61, 0x62, 0x2e, 0x6f, 0x72,
	0x67, 0x2f, 0x65, 0x63, 0x6f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x73, 0x70, 0x61, 0x72,
	0x6b, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x3b, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x53, 0x53, 0x58, 0xaa, 0x02, 0x12, 0x53, 0x70,
	0x61, 0x72, 0x6b, 0x6c, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x31,
	0xca, 0x02, 0x12, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x5c, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1e, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65, 0x5c,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x14, 0x53, 0x70, 0x61, 0x72, 0x6b, 0x6c, 0x65,
	0x3a, 0x3a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_service_v1_service_proto_rawDescOnce sync.Once
	file_service_v1_service_proto_rawDescData = file_service_v1_service_proto_rawDesc
)

func file_service_v1_service_proto_rawDescGZIP() []byte {
	file_service_v1_service_proto_rawDescOnce.Do(func() {
		file_service_v1_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_service_v1_service_proto_rawDescData)
	})
	return file_service_v1_service_proto_rawDescData
}

var file_service_v1_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_service_v1_service_proto_goTypes = []interface{}{
	(*PingRequest)(nil),  // 0: sparkle.service.v1.PingRequest
	(*PongResponse)(nil), // 1: sparkle.service.v1.PongResponse
	nil,                  // 2: sparkle.service.v1.PongResponse.EnvEntry
}
var file_service_v1_service_proto_depIdxs = []int32{
	2, // 0: sparkle.service.v1.PongResponse.env:type_name -> sparkle.service.v1.PongResponse.EnvEntry
	0, // 1: sparkle.service.v1.Service.WhoAmI:input_type -> sparkle.service.v1.PingRequest
	1, // 2: sparkle.service.v1.Service.WhoAmI:output_type -> sparkle.service.v1.PongResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_service_v1_service_proto_init() }
func file_service_v1_service_proto_init() {
	if File_service_v1_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_service_v1_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PingRequest); i {
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
		file_service_v1_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PongResponse); i {
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
			RawDescriptor: file_service_v1_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_service_v1_service_proto_goTypes,
		DependencyIndexes: file_service_v1_service_proto_depIdxs,
		MessageInfos:      file_service_v1_service_proto_msgTypes,
	}.Build()
	File_service_v1_service_proto = out.File
	file_service_v1_service_proto_rawDesc = nil
	file_service_v1_service_proto_goTypes = nil
	file_service_v1_service_proto_depIdxs = nil
}
