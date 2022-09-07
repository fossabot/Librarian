// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: conf/porter.proto

package conf

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Porter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Server *Porter_Server `protobuf:"bytes,1,opt,name=server,proto3" json:"server,omitempty"`
	Data   *Porter_Data   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *Porter) Reset() {
	*x = Porter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_porter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Porter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Porter) ProtoMessage() {}

func (x *Porter) ProtoReflect() protoreflect.Message {
	mi := &file_conf_porter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Porter.ProtoReflect.Descriptor instead.
func (*Porter) Descriptor() ([]byte, []int) {
	return file_conf_porter_proto_rawDescGZIP(), []int{0}
}

func (x *Porter) GetServer() *Porter_Server {
	if x != nil {
		return x.Server
	}
	return nil
}

func (x *Porter) GetData() *Porter_Data {
	if x != nil {
		return x.Data
	}
	return nil
}

type Porter_Server struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Grpc *Porter_Server_GRPC `protobuf:"bytes,1,opt,name=grpc,proto3" json:"grpc,omitempty"`
}

func (x *Porter_Server) Reset() {
	*x = Porter_Server{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_porter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Porter_Server) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Porter_Server) ProtoMessage() {}

func (x *Porter_Server) ProtoReflect() protoreflect.Message {
	mi := &file_conf_porter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Porter_Server.ProtoReflect.Descriptor instead.
func (*Porter_Server) Descriptor() ([]byte, []int) {
	return file_conf_porter_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Porter_Server) GetGrpc() *Porter_Server_GRPC {
	if x != nil {
		return x.Grpc
	}
	return nil
}

type Porter_Data struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Database *Porter_Data_Database `protobuf:"bytes,1,opt,name=database,proto3" json:"database,omitempty"`
}

func (x *Porter_Data) Reset() {
	*x = Porter_Data{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_porter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Porter_Data) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Porter_Data) ProtoMessage() {}

func (x *Porter_Data) ProtoReflect() protoreflect.Message {
	mi := &file_conf_porter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Porter_Data.ProtoReflect.Descriptor instead.
func (*Porter_Data) Descriptor() ([]byte, []int) {
	return file_conf_porter_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Porter_Data) GetDatabase() *Porter_Data_Database {
	if x != nil {
		return x.Database
	}
	return nil
}

type Porter_Server_GRPC struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Network string               `protobuf:"bytes,1,opt,name=network,proto3" json:"network,omitempty"`
	Addr    string               `protobuf:"bytes,2,opt,name=addr,proto3" json:"addr,omitempty"`
	Timeout *durationpb.Duration `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *Porter_Server_GRPC) Reset() {
	*x = Porter_Server_GRPC{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_porter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Porter_Server_GRPC) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Porter_Server_GRPC) ProtoMessage() {}

func (x *Porter_Server_GRPC) ProtoReflect() protoreflect.Message {
	mi := &file_conf_porter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Porter_Server_GRPC.ProtoReflect.Descriptor instead.
func (*Porter_Server_GRPC) Descriptor() ([]byte, []int) {
	return file_conf_porter_proto_rawDescGZIP(), []int{0, 0, 0}
}

func (x *Porter_Server_GRPC) GetNetwork() string {
	if x != nil {
		return x.Network
	}
	return ""
}

func (x *Porter_Server_GRPC) GetAddr() string {
	if x != nil {
		return x.Addr
	}
	return ""
}

func (x *Porter_Server_GRPC) GetTimeout() *durationpb.Duration {
	if x != nil {
		return x.Timeout
	}
	return nil
}

type Porter_Data_Database struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Driver string `protobuf:"bytes,1,opt,name=driver,proto3" json:"driver,omitempty"`
	Source string `protobuf:"bytes,2,opt,name=source,proto3" json:"source,omitempty"`
}

func (x *Porter_Data_Database) Reset() {
	*x = Porter_Data_Database{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conf_porter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Porter_Data_Database) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Porter_Data_Database) ProtoMessage() {}

func (x *Porter_Data_Database) ProtoReflect() protoreflect.Message {
	mi := &file_conf_porter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Porter_Data_Database.ProtoReflect.Descriptor instead.
func (*Porter_Data_Database) Descriptor() ([]byte, []int) {
	return file_conf_porter_proto_rawDescGZIP(), []int{0, 1, 0}
}

func (x *Porter_Data_Database) GetDriver() string {
	if x != nil {
		return x.Driver
	}
	return ""
}

func (x *Porter_Data_Database) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

var File_conf_porter_proto protoreflect.FileDescriptor

var file_conf_porter_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6f, 0x6e, 0x66, 0x2f, 0x70, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x1a,
	0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x95, 0x03, 0x0a, 0x06, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x6b, 0x72, 0x61,
	0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x52, 0x06, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6b, 0x72,
	0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0xa7, 0x01, 0x0a, 0x06, 0x53,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x32, 0x0a, 0x04, 0x67, 0x72, 0x70, 0x63, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x50, 0x6f, 0x72, 0x74, 0x65, 0x72, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x47,
	0x52, 0x50, 0x43, 0x52, 0x04, 0x67, 0x72, 0x70, 0x63, 0x1a, 0x69, 0x0a, 0x04, 0x47, 0x52, 0x50,
	0x43, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x12, 0x0a, 0x04, 0x61,
	0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x61, 0x64, 0x64, 0x72, 0x12,
	0x33, 0x0a, 0x07, 0x74, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x1a, 0x80, 0x01, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3c, 0x0a,
	0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x50, 0x6f, 0x72,
	0x74, 0x65, 0x72, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73,
	0x65, 0x52, 0x08, 0x64, 0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x1a, 0x3a, 0x0a, 0x08, 0x44,
	0x61, 0x74, 0x61, 0x62, 0x61, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12,
	0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x1e, 0x5a, 0x1c, 0x4c, 0x69, 0x62, 0x72, 0x61,
	0x72, 0x69, 0x61, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x3b, 0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conf_porter_proto_rawDescOnce sync.Once
	file_conf_porter_proto_rawDescData = file_conf_porter_proto_rawDesc
)

func file_conf_porter_proto_rawDescGZIP() []byte {
	file_conf_porter_proto_rawDescOnce.Do(func() {
		file_conf_porter_proto_rawDescData = protoimpl.X.CompressGZIP(file_conf_porter_proto_rawDescData)
	})
	return file_conf_porter_proto_rawDescData
}

var file_conf_porter_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_conf_porter_proto_goTypes = []interface{}{
	(*Porter)(nil),               // 0: kratos.api.Porter
	(*Porter_Server)(nil),        // 1: kratos.api.Porter.Server
	(*Porter_Data)(nil),          // 2: kratos.api.Porter.Data
	(*Porter_Server_GRPC)(nil),   // 3: kratos.api.Porter.Server.GRPC
	(*Porter_Data_Database)(nil), // 4: kratos.api.Porter.Data.Database
	(*durationpb.Duration)(nil),  // 5: google.protobuf.Duration
}
var file_conf_porter_proto_depIdxs = []int32{
	1, // 0: kratos.api.Porter.server:type_name -> kratos.api.Porter.Server
	2, // 1: kratos.api.Porter.data:type_name -> kratos.api.Porter.Data
	3, // 2: kratos.api.Porter.Server.grpc:type_name -> kratos.api.Porter.Server.GRPC
	4, // 3: kratos.api.Porter.Data.database:type_name -> kratos.api.Porter.Data.Database
	5, // 4: kratos.api.Porter.Server.GRPC.timeout:type_name -> google.protobuf.Duration
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_conf_porter_proto_init() }
func file_conf_porter_proto_init() {
	if File_conf_porter_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conf_porter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Porter); i {
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
		file_conf_porter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Porter_Server); i {
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
		file_conf_porter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Porter_Data); i {
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
		file_conf_porter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Porter_Server_GRPC); i {
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
		file_conf_porter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Porter_Data_Database); i {
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
			RawDescriptor: file_conf_porter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_conf_porter_proto_goTypes,
		DependencyIndexes: file_conf_porter_proto_depIdxs,
		MessageInfos:      file_conf_porter_proto_msgTypes,
	}.Build()
	File_conf_porter_proto = out.File
	file_conf_porter_proto_rawDesc = nil
	file_conf_porter_proto_goTypes = nil
	file_conf_porter_proto_depIdxs = nil
}
