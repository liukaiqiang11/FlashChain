// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: proposal_response.proto

package pb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ProposalResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version     int32                  `protobuf:"varint,1,opt,name=Version,proto3" json:"Version,omitempty"`
	Timestamp   *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Response    *Response              `protobuf:"bytes,3,opt,name=Response,proto3" json:"Response,omitempty"`
	Payload     []byte                 `protobuf:"bytes,4,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Endorsement *Endorsement           `protobuf:"bytes,5,opt,name=Endorsement,proto3" json:"Endorsement,omitempty"`
}

func (x *ProposalResponse) Reset() {
	*x = ProposalResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_response_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalResponse) ProtoMessage() {}

func (x *ProposalResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_response_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalResponse.ProtoReflect.Descriptor instead.
func (*ProposalResponse) Descriptor() ([]byte, []int) {
	return file_proposal_response_proto_rawDescGZIP(), []int{0}
}

func (x *ProposalResponse) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

func (x *ProposalResponse) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ProposalResponse) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *ProposalResponse) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ProposalResponse) GetEndorsement() *Endorsement {
	if x != nil {
		return x.Endorsement
	}
	return nil
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  int32  `protobuf:"varint,1,opt,name=Status,proto3" json:"Status,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=Message,proto3" json:"Message,omitempty"`
	Payload []byte `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_response_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_response_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_proposal_response_proto_rawDescGZIP(), []int{1}
}

func (x *Response) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *Response) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *Response) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

type ProposalResponsePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProposalHash []byte `protobuf:"bytes,1,opt,name=ProposalHash,proto3" json:"ProposalHash,omitempty"`
	Extension    []byte `protobuf:"bytes,2,opt,name=Extension,proto3" json:"Extension,omitempty"`
}

func (x *ProposalResponsePayload) Reset() {
	*x = ProposalResponsePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_response_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProposalResponsePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProposalResponsePayload) ProtoMessage() {}

func (x *ProposalResponsePayload) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_response_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProposalResponsePayload.ProtoReflect.Descriptor instead.
func (*ProposalResponsePayload) Descriptor() ([]byte, []int) {
	return file_proposal_response_proto_rawDescGZIP(), []int{2}
}

func (x *ProposalResponsePayload) GetProposalHash() []byte {
	if x != nil {
		return x.ProposalHash
	}
	return nil
}

func (x *ProposalResponsePayload) GetExtension() []byte {
	if x != nil {
		return x.Extension
	}
	return nil
}

type Endorsement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Endorser  []byte `protobuf:"bytes,1,opt,name=Endorser,proto3" json:"Endorser,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (x *Endorsement) Reset() {
	*x = Endorsement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proposal_response_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endorsement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endorsement) ProtoMessage() {}

func (x *Endorsement) ProtoReflect() protoreflect.Message {
	mi := &file_proposal_response_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endorsement.ProtoReflect.Descriptor instead.
func (*Endorsement) Descriptor() ([]byte, []int) {
	return file_proposal_response_proto_rawDescGZIP(), []int{3}
}

func (x *Endorsement) GetEndorser() []byte {
	if x != nil {
		return x.Endorser
	}
	return nil
}

func (x *Endorsement) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_proposal_response_proto protoreflect.FileDescriptor

var file_proposal_response_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdd,
	0x01, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x38, 0x0a,
	0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x31, 0x0a, 0x0b, 0x45,
	0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0b, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x56,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x22, 0x5b, 0x0a, 0x17, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73,
	0x61, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x22, 0x0a, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61,
	0x6c, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73,
	0x69, 0x6f, 0x6e, 0x22, 0x47, 0x0a, 0x0b, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x45, 0x6e, 0x64, 0x6f, 0x72, 0x73, 0x65, 0x72, 0x12, 0x1c,
	0x0a, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x42, 0x07, 0x5a, 0x05,
	0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proposal_response_proto_rawDescOnce sync.Once
	file_proposal_response_proto_rawDescData = file_proposal_response_proto_rawDesc
)

func file_proposal_response_proto_rawDescGZIP() []byte {
	file_proposal_response_proto_rawDescOnce.Do(func() {
		file_proposal_response_proto_rawDescData = protoimpl.X.CompressGZIP(file_proposal_response_proto_rawDescData)
	})
	return file_proposal_response_proto_rawDescData
}

var file_proposal_response_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_proposal_response_proto_goTypes = []interface{}{
	(*ProposalResponse)(nil),        // 0: pb.ProposalResponse
	(*Response)(nil),                // 1: pb.Response
	(*ProposalResponsePayload)(nil), // 2: pb.ProposalResponsePayload
	(*Endorsement)(nil),             // 3: pb.Endorsement
	(*timestamppb.Timestamp)(nil),   // 4: google.protobuf.Timestamp
}
var file_proposal_response_proto_depIdxs = []int32{
	4, // 0: pb.ProposalResponse.Timestamp:type_name -> google.protobuf.Timestamp
	1, // 1: pb.ProposalResponse.Response:type_name -> pb.Response
	3, // 2: pb.ProposalResponse.Endorsement:type_name -> pb.Endorsement
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proposal_response_proto_init() }
func file_proposal_response_proto_init() {
	if File_proposal_response_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proposal_response_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalResponse); i {
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
		file_proposal_response_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
		file_proposal_response_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProposalResponsePayload); i {
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
		file_proposal_response_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endorsement); i {
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
			RawDescriptor: file_proposal_response_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proposal_response_proto_goTypes,
		DependencyIndexes: file_proposal_response_proto_depIdxs,
		MessageInfos:      file_proposal_response_proto_msgTypes,
	}.Build()
	File_proposal_response_proto = out.File
	file_proposal_response_proto_rawDesc = nil
	file_proposal_response_proto_goTypes = nil
	file_proposal_response_proto_depIdxs = nil
}