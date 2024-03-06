// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: chaincode.proto

package pb

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

type ChaincodeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=Name,proto3" json:"Name,omitempty"`
	Version string `protobuf:"bytes,3,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *ChaincodeID) Reset() {
	*x = ChaincodeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeID) ProtoMessage() {}

func (x *ChaincodeID) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeID.ProtoReflect.Descriptor instead.
func (*ChaincodeID) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{0}
}

func (x *ChaincodeID) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ChaincodeID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChaincodeID) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type ChaincodeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Args   [][]byte `protobuf:"bytes,1,rep,name=Args,proto3" json:"Args,omitempty"`
	IsInit bool     `protobuf:"varint,2,opt,name=IsInit,proto3" json:"IsInit,omitempty"`
}

func (x *ChaincodeInput) Reset() {
	*x = ChaincodeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInput) ProtoMessage() {}

func (x *ChaincodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInput.ProtoReflect.Descriptor instead.
func (*ChaincodeInput) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{1}
}

func (x *ChaincodeInput) GetArgs() [][]byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ChaincodeInput) GetIsInit() bool {
	if x != nil {
		return x.IsInit
	}
	return false
}

type ChaincodeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChaincodeID *ChaincodeID    `protobuf:"bytes,1,opt,name=ChaincodeID,proto3" json:"ChaincodeID,omitempty"`
	Input       *ChaincodeInput `protobuf:"bytes,2,opt,name=Input,proto3" json:"Input,omitempty"`
	Timeout     int32           `protobuf:"varint,3,opt,name=Timeout,proto3" json:"Timeout,omitempty"`
}

func (x *ChaincodeSpec) Reset() {
	*x = ChaincodeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeSpec) ProtoMessage() {}

func (x *ChaincodeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeSpec.ProtoReflect.Descriptor instead.
func (*ChaincodeSpec) Descriptor() ([]byte, []int) {
	return file_chaincode_proto_rawDescGZIP(), []int{2}
}

func (x *ChaincodeSpec) GetChaincodeID() *ChaincodeID {
	if x != nil {
		return x.ChaincodeID
	}
	return nil
}

func (x *ChaincodeSpec) GetInput() *ChaincodeInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *ChaincodeSpec) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

var File_chaincode_proto protoreflect.FileDescriptor

var file_chaincode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x02, 0x70, 0x62, 0x22, 0x4f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x50, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x56,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x3c, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x41, 0x72, 0x67, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x41, 0x72, 0x67, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x49, 0x73, 0x49, 0x6e, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x49, 0x73,
	0x49, 0x6e, 0x69, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x31, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x0b, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x12, 0x28, 0x0a, 0x05, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x52, 0x05, 0x49, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x54, 0x69, 0x6d, 0x65, 0x6f, 0x75, 0x74, 0x42, 0x07, 0x5a,
	0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaincode_proto_rawDescOnce sync.Once
	file_chaincode_proto_rawDescData = file_chaincode_proto_rawDesc
)

func file_chaincode_proto_rawDescGZIP() []byte {
	file_chaincode_proto_rawDescOnce.Do(func() {
		file_chaincode_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaincode_proto_rawDescData)
	})
	return file_chaincode_proto_rawDescData
}

var file_chaincode_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_chaincode_proto_goTypes = []interface{}{
	(*ChaincodeID)(nil),    // 0: pb.ChaincodeID
	(*ChaincodeInput)(nil), // 1: pb.ChaincodeInput
	(*ChaincodeSpec)(nil),  // 2: pb.ChaincodeSpec
}
var file_chaincode_proto_depIdxs = []int32{
	0, // 0: pb.ChaincodeSpec.ChaincodeID:type_name -> pb.ChaincodeID
	1, // 1: pb.ChaincodeSpec.Input:type_name -> pb.ChaincodeInput
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_chaincode_proto_init() }
func file_chaincode_proto_init() {
	if File_chaincode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_chaincode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeID); i {
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
		file_chaincode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInput); i {
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
		file_chaincode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeSpec); i {
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
			RawDescriptor: file_chaincode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chaincode_proto_goTypes,
		DependencyIndexes: file_chaincode_proto_depIdxs,
		MessageInfos:      file_chaincode_proto_msgTypes,
	}.Build()
	File_chaincode_proto = out.File
	file_chaincode_proto_rawDesc = nil
	file_chaincode_proto_goTypes = nil
	file_chaincode_proto_depIdxs = nil
}