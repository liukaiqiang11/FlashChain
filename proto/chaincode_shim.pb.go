// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v4.22.2
// source: chaincode_shim.proto

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

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_Type_UNDEFINED             ChaincodeMessage_Type = 0
	ChaincodeMessage_Type_REGISTER              ChaincodeMessage_Type = 1
	ChaincodeMessage_Type_REGISTERED            ChaincodeMessage_Type = 2
	ChaincodeMessage_Type_INIT                  ChaincodeMessage_Type = 3
	ChaincodeMessage_Type_READY                 ChaincodeMessage_Type = 4
	ChaincodeMessage_Type_TRANSACTION           ChaincodeMessage_Type = 5
	ChaincodeMessage_Type_COMPLETED             ChaincodeMessage_Type = 6
	ChaincodeMessage_Type_ERROR                 ChaincodeMessage_Type = 7
	ChaincodeMessage_Type_GET_STATE             ChaincodeMessage_Type = 8
	ChaincodeMessage_Type_PUT_STATE             ChaincodeMessage_Type = 9
	ChaincodeMessage_Type_DEL_STATE             ChaincodeMessage_Type = 10
	ChaincodeMessage_Type_INVOKE_CHAINCODE      ChaincodeMessage_Type = 11
	ChaincodeMessage_Type_ONLY_READ_TRANSACTION ChaincodeMessage_Type = 12
	ChaincodeMessage_Type_RESPONSE              ChaincodeMessage_Type = 13
	ChaincodeMessage_Type_GET_STATE_BY_RANGE    ChaincodeMessage_Type = 14
	ChaincodeMessage_Type_GET_QUERY_RESULT      ChaincodeMessage_Type = 15
	ChaincodeMessage_Type_QUERY_STATE_NEXT      ChaincodeMessage_Type = 16
	ChaincodeMessage_Type_QUERY_STATE_CLOSE     ChaincodeMessage_Type = 17
	ChaincodeMessage_Type_KEEPALIVE             ChaincodeMessage_Type = 18
	ChaincodeMessage_Type_GET_HISTORY_FOR_KEY   ChaincodeMessage_Type = 19
	ChaincodeMessage_Type_GET_STATE_METADATA    ChaincodeMessage_Type = 20
	ChaincodeMessage_Type_PUT_STATE_METADATA    ChaincodeMessage_Type = 21
	ChaincodeMessage_Type_GET_PRIVATE_DATA_HASH ChaincodeMessage_Type = 22
	ChaincodeMessage_Type_PURGE_PRIVATE_DATA    ChaincodeMessage_Type = 23
)

// Enum value maps for ChaincodeMessage_Type.
var (
	ChaincodeMessage_Type_name = map[int32]string{
		0:  "UNDEFINED",
		1:  "REGISTER",
		2:  "REGISTERED",
		3:  "INIT",
		4:  "READY",
		5:  "TRANSACTION",
		6:  "COMPLETED",
		7:  "ERROR",
		8:  "GET_STATE",
		9:  "PUT_STATE",
		10: "DEL_STATE",
		11: "INVOKE_CHAINCODE",
		12: "ONLY_READ_TRANSACTION",
		13: "RESPONSE",
		14: "GET_STATE_BY_RANGE",
		15: "GET_QUERY_RESULT",
		16: "QUERY_STATE_NEXT",
		17: "QUERY_STATE_CLOSE",
		18: "KEEPALIVE",
		19: "GET_HISTORY_FOR_KEY",
		20: "GET_STATE_METADATA",
		21: "PUT_STATE_METADATA",
		22: "GET_PRIVATE_DATA_HASH",
		23: "PURGE_PRIVATE_DATA",
	}
	ChaincodeMessage_Type_value = map[string]int32{
		"UNDEFINED":             0,
		"REGISTER":              1,
		"REGISTERED":            2,
		"INIT":                  3,
		"READY":                 4,
		"TRANSACTION":           5,
		"COMPLETED":             6,
		"ERROR":                 7,
		"GET_STATE":             8,
		"PUT_STATE":             9,
		"DEL_STATE":             10,
		"INVOKE_CHAINCODE":      11,
		"ONLY_READ_TRANSACTION": 12,
		"RESPONSE":              13,
		"GET_STATE_BY_RANGE":    14,
		"GET_QUERY_RESULT":      15,
		"QUERY_STATE_NEXT":      16,
		"QUERY_STATE_CLOSE":     17,
		"KEEPALIVE":             18,
		"GET_HISTORY_FOR_KEY":   19,
		"GET_STATE_METADATA":    20,
		"PUT_STATE_METADATA":    21,
		"GET_PRIVATE_DATA_HASH": 22,
		"PURGE_PRIVATE_DATA":    23,
	}
)

func (x ChaincodeMessage_Type) Enum() *ChaincodeMessage_Type {
	p := new(ChaincodeMessage_Type)
	*p = x
	return p
}

func (x ChaincodeMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChaincodeMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_chaincode_shim_proto_enumTypes[0].Descriptor()
}

func (ChaincodeMessage_Type) Type() protoreflect.EnumType {
	return &file_chaincode_shim_proto_enumTypes[0]
}

func (x ChaincodeMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChaincodeMessage_Type.Descriptor instead.
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{0}
}

type ValidateMessage_Type int32

const (
	ValidateMessage_Type_VALIDATE           ValidateMessage_Type = 0
	ValidateMessage_Type_VALIDATE_COMPLETED ValidateMessage_Type = 1
	ValidateMessage_Type_VALIDATE_SUCCESS   ValidateMessage_Type = 2
	ValidateMessage_Type_VALIDATE_FAIL      ValidateMessage_Type = 3
)

// Enum value maps for ValidateMessage_Type.
var (
	ValidateMessage_Type_name = map[int32]string{
		0: "VALIDATE",
		1: "VALIDATE_COMPLETED",
		2: "VALIDATE_SUCCESS",
		3: "VALIDATE_FAIL",
	}
	ValidateMessage_Type_value = map[string]int32{
		"VALIDATE":           0,
		"VALIDATE_COMPLETED": 1,
		"VALIDATE_SUCCESS":   2,
		"VALIDATE_FAIL":      3,
	}
)

func (x ValidateMessage_Type) Enum() *ValidateMessage_Type {
	p := new(ValidateMessage_Type)
	*p = x
	return p
}

func (x ValidateMessage_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ValidateMessage_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_chaincode_shim_proto_enumTypes[1].Descriptor()
}

func (ValidateMessage_Type) Type() protoreflect.EnumType {
	return &file_chaincode_shim_proto_enumTypes[1]
}

func (x ValidateMessage_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ValidateMessage_Type.Descriptor instead.
func (ValidateMessage_Type) EnumDescriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{1}
}

type ChaincodeMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      ChaincodeMessage_Type `protobuf:"varint,1,opt,name=Type,proto3,enum=pb.ChaincodeMessage_Type" json:"Type,omitempty"`
	Timestamp uint64                `protobuf:"varint,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Payload   []byte                `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	TxID      string                `protobuf:"bytes,4,opt,name=TxID,proto3" json:"TxID,omitempty"`
	Proposal  *SignedProposal       `protobuf:"bytes,5,opt,name=Proposal,proto3" json:"Proposal,omitempty"`
}

func (x *ChaincodeMessage) Reset() {
	*x = ChaincodeMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_shim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeMessage) ProtoMessage() {}

func (x *ChaincodeMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_shim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeMessage.ProtoReflect.Descriptor instead.
func (*ChaincodeMessage) Descriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{0}
}

func (x *ChaincodeMessage) GetType() ChaincodeMessage_Type {
	if x != nil {
		return x.Type
	}
	return ChaincodeMessage_Type_UNDEFINED
}

func (x *ChaincodeMessage) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *ChaincodeMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ChaincodeMessage) GetTxID() string {
	if x != nil {
		return x.TxID
	}
	return ""
}

func (x *ChaincodeMessage) GetProposal() *SignedProposal {
	if x != nil {
		return x.Proposal
	}
	return nil
}

type GetState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *GetState) Reset() {
	*x = GetState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_shim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetState) ProtoMessage() {}

func (x *GetState) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_shim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetState.ProtoReflect.Descriptor instead.
func (*GetState) Descriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{1}
}

func (x *GetState) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PutState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key   string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *PutState) Reset() {
	*x = PutState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_shim_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PutState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PutState) ProtoMessage() {}

func (x *PutState) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_shim_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PutState.ProtoReflect.Descriptor instead.
func (*PutState) Descriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{2}
}

func (x *PutState) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *PutState) GetValue() []byte {
	if x != nil {
		return x.Value
	}
	return nil
}

type DelState struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=Key,proto3" json:"Key,omitempty"`
}

func (x *DelState) Reset() {
	*x = DelState{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_shim_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DelState) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DelState) ProtoMessage() {}

func (x *DelState) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_shim_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DelState.ProtoReflect.Descriptor instead.
func (*DelState) Descriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{3}
}

func (x *DelState) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_shim_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_shim_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{4}
}

type ValidateMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type      ValidateMessage_Type   `protobuf:"varint,1,opt,name=Type,proto3,enum=pb.ValidateMessage_Type" json:"Type,omitempty"`
	Timestamp *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`
	Payload   []byte                 `protobuf:"bytes,3,opt,name=Payload,proto3" json:"Payload,omitempty"`
	Response  *Response              `protobuf:"bytes,4,opt,name=Response,proto3" json:"Response,omitempty"`
}

func (x *ValidateMessage) Reset() {
	*x = ValidateMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_chaincode_shim_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateMessage) ProtoMessage() {}

func (x *ValidateMessage) ProtoReflect() protoreflect.Message {
	mi := &file_chaincode_shim_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateMessage.ProtoReflect.Descriptor instead.
func (*ValidateMessage) Descriptor() ([]byte, []int) {
	return file_chaincode_shim_proto_rawDescGZIP(), []int{5}
}

func (x *ValidateMessage) GetType() ValidateMessage_Type {
	if x != nil {
		return x.Type
	}
	return ValidateMessage_Type_VALIDATE
}

func (x *ValidateMessage) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *ValidateMessage) GetPayload() []byte {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *ValidateMessage) GetResponse() *Response {
	if x != nil {
		return x.Response
	}
	return nil
}

var File_chaincode_shim_proto protoreflect.FileDescriptor

var file_chaincode_shim_proto_rawDesc = []byte{
	0x0a, 0x14, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x68, 0x69, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0e, 0x70, 0x72, 0x6f,
	0x70, 0x6f, 0x73, 0x61, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0d, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x70, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x01, 0x0a, 0x10, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2d, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x54, 0x79, 0x70,
	0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x54, 0x78, 0x49, 0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x54,
	0x78, 0x49, 0x44, 0x12, 0x2e, 0x0a, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x65,
	0x64, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x52, 0x08, 0x50, 0x72, 0x6f, 0x70, 0x6f,
	0x73, 0x61, 0x6c, 0x22, 0x1c, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65,
	0x79, 0x22, 0x32, 0x0a, 0x08, 0x50, 0x75, 0x74, 0x53, 0x74, 0x61, 0x74, 0x65, 0x12, 0x10, 0x0a,
	0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x4b, 0x65, 0x79, 0x12,
	0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x44, 0x65, 0x6c, 0x53, 0x74, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x4b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x4b, 0x65, 0x79, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0xbd, 0x01, 0x0a,
	0x0f, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x2c, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x18,
	0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x38,
	0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x54,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x18, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x28, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x62, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x52, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0xd6, 0x03, 0x0a,
	0x15, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x46, 0x49,
	0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45,
	0x52, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x45, 0x47, 0x49, 0x53, 0x54, 0x45, 0x52, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x49, 0x4e, 0x49, 0x54, 0x10, 0x03, 0x12, 0x09, 0x0a,
	0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x04, 0x12, 0x0f, 0x0a, 0x0b, 0x54, 0x52, 0x41, 0x4e,
	0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10, 0x05, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x4f, 0x4d,
	0x50, 0x4c, 0x45, 0x54, 0x45, 0x44, 0x10, 0x06, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f,
	0x52, 0x10, 0x07, 0x12, 0x0d, 0x0a, 0x09, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45,
	0x10, 0x08, 0x12, 0x0d, 0x0a, 0x09, 0x50, 0x55, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10,
	0x09, 0x12, 0x0d, 0x0a, 0x09, 0x44, 0x45, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x10, 0x0a,
	0x12, 0x14, 0x0a, 0x10, 0x49, 0x4e, 0x56, 0x4f, 0x4b, 0x45, 0x5f, 0x43, 0x48, 0x41, 0x49, 0x4e,
	0x43, 0x4f, 0x44, 0x45, 0x10, 0x0b, 0x12, 0x19, 0x0a, 0x15, 0x4f, 0x4e, 0x4c, 0x59, 0x5f, 0x52,
	0x45, 0x41, 0x44, 0x5f, 0x54, 0x52, 0x41, 0x4e, 0x53, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x10,
	0x0c, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x10, 0x0d, 0x12,
	0x16, 0x0a, 0x12, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x42, 0x59, 0x5f,
	0x52, 0x41, 0x4e, 0x47, 0x45, 0x10, 0x0e, 0x12, 0x14, 0x0a, 0x10, 0x47, 0x45, 0x54, 0x5f, 0x51,
	0x55, 0x45, 0x52, 0x59, 0x5f, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x10, 0x0f, 0x12, 0x14, 0x0a,
	0x10, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4e, 0x45, 0x58,
	0x54, 0x10, 0x10, 0x12, 0x15, 0x0a, 0x11, 0x51, 0x55, 0x45, 0x52, 0x59, 0x5f, 0x53, 0x54, 0x41,
	0x54, 0x45, 0x5f, 0x43, 0x4c, 0x4f, 0x53, 0x45, 0x10, 0x11, 0x12, 0x0d, 0x0a, 0x09, 0x4b, 0x45,
	0x45, 0x50, 0x41, 0x4c, 0x49, 0x56, 0x45, 0x10, 0x12, 0x12, 0x17, 0x0a, 0x13, 0x47, 0x45, 0x54,
	0x5f, 0x48, 0x49, 0x53, 0x54, 0x4f, 0x52, 0x59, 0x5f, 0x46, 0x4f, 0x52, 0x5f, 0x4b, 0x45, 0x59,
	0x10, 0x13, 0x12, 0x16, 0x0a, 0x12, 0x47, 0x45, 0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f,
	0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41, 0x10, 0x14, 0x12, 0x16, 0x0a, 0x12, 0x50, 0x55,
	0x54, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x4d, 0x45, 0x54, 0x41, 0x44, 0x41, 0x54, 0x41,
	0x10, 0x15, 0x12, 0x19, 0x0a, 0x15, 0x47, 0x45, 0x54, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54,
	0x45, 0x5f, 0x44, 0x41, 0x54, 0x41, 0x5f, 0x48, 0x41, 0x53, 0x48, 0x10, 0x16, 0x12, 0x16, 0x0a,
	0x12, 0x50, 0x55, 0x52, 0x47, 0x45, 0x5f, 0x50, 0x52, 0x49, 0x56, 0x41, 0x54, 0x45, 0x5f, 0x44,
	0x41, 0x54, 0x41, 0x10, 0x17, 0x2a, 0x65, 0x0a, 0x14, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a,
	0x08, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x10, 0x00, 0x12, 0x16, 0x0a, 0x12, 0x56,
	0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x43, 0x4f, 0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45,
	0x44, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x41, 0x54, 0x45, 0x5f,
	0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x02, 0x12, 0x11, 0x0a, 0x0d, 0x56, 0x41, 0x4c,
	0x49, 0x44, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x10, 0x03, 0x32, 0xfc, 0x01, 0x0a,
	0x10, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x21, 0x0a, 0x09, 0x53, 0x74, 0x61, 0x72, 0x74, 0x50, 0x65, 0x65, 0x72, 0x12, 0x09,
	0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x12, 0x1f, 0x0a, 0x07, 0x45, 0x6e, 0x64, 0x50, 0x65, 0x65, 0x72, 0x12,
	0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x09, 0x2e, 0x70, 0x62, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x3a, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61,
	0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x28, 0x01, 0x30,
	0x01, 0x12, 0x2d, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x49, 0x6e, 0x66,
	0x6f, 0x12, 0x09, 0x2e, 0x70, 0x62, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x12, 0x2e, 0x70,
	0x62, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x49, 0x6e, 0x66, 0x6f,
	0x12, 0x39, 0x0a, 0x0d, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x56, 0x61, 0x6c, 0x69,
	0x64, 0x61, 0x74, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x48, 0x0a, 0x09, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x12, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x1a, 0x14, 0x2e, 0x70, 0x62, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x07, 0x5a, 0x05, 0x2e, 0x2f, 0x3b, 0x70, 0x62, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_chaincode_shim_proto_rawDescOnce sync.Once
	file_chaincode_shim_proto_rawDescData = file_chaincode_shim_proto_rawDesc
)

func file_chaincode_shim_proto_rawDescGZIP() []byte {
	file_chaincode_shim_proto_rawDescOnce.Do(func() {
		file_chaincode_shim_proto_rawDescData = protoimpl.X.CompressGZIP(file_chaincode_shim_proto_rawDescData)
	})
	return file_chaincode_shim_proto_rawDescData
}

var file_chaincode_shim_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_chaincode_shim_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_chaincode_shim_proto_goTypes = []interface{}{
	(ChaincodeMessage_Type)(0),    // 0: pb.ChaincodeMessage_Type
	(ValidateMessage_Type)(0),     // 1: pb.ValidateMessage_Type
	(*ChaincodeMessage)(nil),      // 2: pb.ChaincodeMessage
	(*GetState)(nil),              // 3: pb.GetState
	(*PutState)(nil),              // 4: pb.PutState
	(*DelState)(nil),              // 5: pb.DelState
	(*Empty)(nil),                 // 6: pb.Empty
	(*ValidateMessage)(nil),       // 7: pb.ValidateMessage
	(*SignedProposal)(nil),        // 8: pb.SignedProposal
	(*timestamppb.Timestamp)(nil), // 9: google.protobuf.Timestamp
	(*Response)(nil),              // 10: pb.Response
	(*BlockchainInfo)(nil),        // 11: pb.BlockchainInfo
}
var file_chaincode_shim_proto_depIdxs = []int32{
	0,  // 0: pb.ChaincodeMessage.Type:type_name -> pb.ChaincodeMessage_Type
	8,  // 1: pb.ChaincodeMessage.Proposal:type_name -> pb.SignedProposal
	1,  // 2: pb.ValidateMessage.Type:type_name -> pb.ValidateMessage_Type
	9,  // 3: pb.ValidateMessage.Timestamp:type_name -> google.protobuf.Timestamp
	10, // 4: pb.ValidateMessage.Response:type_name -> pb.Response
	6,  // 5: pb.ChaincodeSupport.StartPeer:input_type -> pb.Empty
	6,  // 6: pb.ChaincodeSupport.EndPeer:input_type -> pb.Empty
	2,  // 7: pb.ChaincodeSupport.Register:input_type -> pb.ChaincodeMessage
	6,  // 8: pb.ChaincodeSupport.GetBlockInfo:input_type -> pb.Empty
	7,  // 9: pb.ChaincodeSupport.ValidateBlock:input_type -> pb.ValidateMessage
	2,  // 10: pb.Chaincode.Connect:input_type -> pb.ChaincodeMessage
	6,  // 11: pb.ChaincodeSupport.StartPeer:output_type -> pb.Empty
	6,  // 12: pb.ChaincodeSupport.EndPeer:output_type -> pb.Empty
	2,  // 13: pb.ChaincodeSupport.Register:output_type -> pb.ChaincodeMessage
	11, // 14: pb.ChaincodeSupport.GetBlockInfo:output_type -> pb.BlockchainInfo
	7,  // 15: pb.ChaincodeSupport.ValidateBlock:output_type -> pb.ValidateMessage
	2,  // 16: pb.Chaincode.Connect:output_type -> pb.ChaincodeMessage
	11, // [11:17] is the sub-list for method output_type
	5,  // [5:11] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_chaincode_shim_proto_init() }
func file_chaincode_shim_proto_init() {
	if File_chaincode_shim_proto != nil {
		return
	}
	file_proposal_proto_init()
	file_storage_proto_init()
	file_proposal_response_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_chaincode_shim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeMessage); i {
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
		file_chaincode_shim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetState); i {
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
		file_chaincode_shim_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PutState); i {
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
		file_chaincode_shim_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DelState); i {
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
		file_chaincode_shim_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_chaincode_shim_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateMessage); i {
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
			RawDescriptor: file_chaincode_shim_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_chaincode_shim_proto_goTypes,
		DependencyIndexes: file_chaincode_shim_proto_depIdxs,
		EnumInfos:         file_chaincode_shim_proto_enumTypes,
		MessageInfos:      file_chaincode_shim_proto_msgTypes,
	}.Build()
	File_chaincode_shim_proto = out.File
	file_chaincode_shim_proto_rawDesc = nil
	file_chaincode_shim_proto_goTypes = nil
	file_chaincode_shim_proto_depIdxs = nil
}