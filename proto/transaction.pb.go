// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: transaction.proto

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

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload *ChaincodeActionPayload `protobuf:"bytes,1,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetPayload() *ChaincodeActionPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ChaincodeActionPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input                   *ChaincodeSpec `protobuf:"bytes,1,opt,name=Input,proto3" json:"Input,omitempty"`
	Endorsements            []*Endorsement `protobuf:"bytes,2,rep,name=Endorsements,proto3" json:"Endorsements,omitempty"`
	ProposalResponsePayload []byte         `protobuf:"bytes,3,opt,name=ProposalResponsePayload,proto3" json:"ProposalResponsePayload,omitempty"`
}

func (x *ChaincodeActionPayload) Reset() {
	*x = ChaincodeActionPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeActionPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeActionPayload) ProtoMessage() {}

func (x *ChaincodeActionPayload) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeActionPayload.ProtoReflect.Descriptor instead.
func (*ChaincodeActionPayload) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *ChaincodeActionPayload) GetInput() *ChaincodeSpec {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *ChaincodeActionPayload) GetEndorsements() []*Endorsement {
	if x != nil {
		return x.Endorsements
	}
	return nil
}

func (x *ChaincodeActionPayload) GetProposalResponsePayload() []byte {
	if x != nil {
		return x.ProposalResponsePayload
	}
	return nil
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x0f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x43, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x34, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0xb0, 0x01, 0x0a, 0x16, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x27, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x52, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x33, 0x0a, 0x0c, 0x45, 0x6e,
	0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x38, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x17, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b,
	0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_transaction_proto_goTypes = []interface{}{
	(*Transaction)(nil),            // 0: pb.Transaction
	(*ChaincodeActionPayload)(nil), // 1: pb.ChaincodeActionPayload
	(*ChaincodeSpec)(nil),          // 2: pb.ChaincodeSpec
	(*Endorsement)(nil),            // 3: pb.Endorsement
}
var file_transaction_proto_depIdxs = []int32{
	1, // 0: pb.Transaction.Payload:type_name -> pb.ChaincodeActionPayload
	2, // 1: pb.ChaincodeActionPayload.Input:type_name -> pb.ChaincodeSpec
	3, // 2: pb.ChaincodeActionPayload.Endorsements:type_name -> pb.Endorsement
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	file_chaincode_proto_init()
	file_proposal_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
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
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeActionPayload); i {
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
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
