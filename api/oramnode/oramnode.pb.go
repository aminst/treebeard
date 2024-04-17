// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: oramnode.proto

package oramnode

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

type BlockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block string `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Path  int32  `protobuf:"varint,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *BlockRequest) Reset() {
	*x = BlockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oramnode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockRequest) ProtoMessage() {}

func (x *BlockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oramnode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockRequest.ProtoReflect.Descriptor instead.
func (*BlockRequest) Descriptor() ([]byte, []int) {
	return file_oramnode_proto_rawDescGZIP(), []int{0}
}

func (x *BlockRequest) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *BlockRequest) GetPath() int32 {
	if x != nil {
		return x.Path
	}
	return 0
}

type ReadPathRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Requests  []*BlockRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	StorageId int32           `protobuf:"varint,3,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
}

func (x *ReadPathRequest) Reset() {
	*x = ReadPathRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oramnode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPathRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPathRequest) ProtoMessage() {}

func (x *ReadPathRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oramnode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPathRequest.ProtoReflect.Descriptor instead.
func (*ReadPathRequest) Descriptor() ([]byte, []int) {
	return file_oramnode_proto_rawDescGZIP(), []int{1}
}

func (x *ReadPathRequest) GetRequests() []*BlockRequest {
	if x != nil {
		return x.Requests
	}
	return nil
}

func (x *ReadPathRequest) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

type BlockResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block string `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *BlockResponse) Reset() {
	*x = BlockResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oramnode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BlockResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BlockResponse) ProtoMessage() {}

func (x *BlockResponse) ProtoReflect() protoreflect.Message {
	mi := &file_oramnode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BlockResponse.ProtoReflect.Descriptor instead.
func (*BlockResponse) Descriptor() ([]byte, []int) {
	return file_oramnode_proto_rawDescGZIP(), []int{2}
}

func (x *BlockResponse) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *BlockResponse) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ReadPathReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*BlockResponse `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *ReadPathReply) Reset() {
	*x = ReadPathReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oramnode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadPathReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadPathReply) ProtoMessage() {}

func (x *ReadPathReply) ProtoReflect() protoreflect.Message {
	mi := &file_oramnode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadPathReply.ProtoReflect.Descriptor instead.
func (*ReadPathReply) Descriptor() ([]byte, []int) {
	return file_oramnode_proto_rawDescGZIP(), []int{3}
}

func (x *ReadPathReply) GetResponses() []*BlockResponse {
	if x != nil {
		return x.Responses
	}
	return nil
}

type JoinRaftVoterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NodeId   int32  `protobuf:"varint,1,opt,name=node_id,json=nodeId,proto3" json:"node_id,omitempty"`
	NodeAddr string `protobuf:"bytes,2,opt,name=node_addr,json=nodeAddr,proto3" json:"node_addr,omitempty"`
}

func (x *JoinRaftVoterRequest) Reset() {
	*x = JoinRaftVoterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oramnode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRaftVoterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRaftVoterRequest) ProtoMessage() {}

func (x *JoinRaftVoterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_oramnode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRaftVoterRequest.ProtoReflect.Descriptor instead.
func (*JoinRaftVoterRequest) Descriptor() ([]byte, []int) {
	return file_oramnode_proto_rawDescGZIP(), []int{4}
}

func (x *JoinRaftVoterRequest) GetNodeId() int32 {
	if x != nil {
		return x.NodeId
	}
	return 0
}

func (x *JoinRaftVoterRequest) GetNodeAddr() string {
	if x != nil {
		return x.NodeAddr
	}
	return ""
}

type JoinRaftVoterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *JoinRaftVoterReply) Reset() {
	*x = JoinRaftVoterReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_oramnode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRaftVoterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRaftVoterReply) ProtoMessage() {}

func (x *JoinRaftVoterReply) ProtoReflect() protoreflect.Message {
	mi := &file_oramnode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use JoinRaftVoterReply.ProtoReflect.Descriptor instead.
func (*JoinRaftVoterReply) Descriptor() ([]byte, []int) {
	return file_oramnode_proto_rawDescGZIP(), []int{5}
}

func (x *JoinRaftVoterReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_oramnode_proto protoreflect.FileDescriptor

var file_oramnode_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x6f, 0x72, 0x61, 0x6d, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x6f, 0x72, 0x61, 0x6d, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x38, 0x0a, 0x0c, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c,
	0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x70, 0x61, 0x74, 0x68, 0x22, 0x64, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x50, 0x61, 0x74, 0x68,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x32, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6f, 0x72, 0x61, 0x6d,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x0d, 0x42, 0x6c,
	0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62,
	0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63,
	0x6b, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x64, 0x50,
	0x61, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x35, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x6f, 0x72,
	0x61, 0x6d, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x52, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22,
	0x4c, 0x0a, 0x14, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x41, 0x64, 0x64, 0x72, 0x22, 0x2e, 0x0a,
	0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x9d, 0x01,
	0x0a, 0x08, 0x4f, 0x72, 0x61, 0x6d, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x40, 0x0a, 0x08, 0x52, 0x65,
	0x61, 0x64, 0x50, 0x61, 0x74, 0x68, 0x12, 0x19, 0x2e, 0x6f, 0x72, 0x61, 0x6d, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x17, 0x2e, 0x6f, 0x72, 0x61, 0x6d, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x61,
	0x64, 0x50, 0x61, 0x74, 0x68, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x4f, 0x0a, 0x0d,
	0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x1e, 0x2e,
	0x6f, 0x72, 0x61, 0x6d, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e,
	0x6f, 0x72, 0x61, 0x6d, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66,
	0x74, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x31, 0x5a,
	0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x73, 0x67, 0x2d,
	0x75, 0x77, 0x61, 0x74, 0x65, 0x72, 0x6c, 0x6f, 0x6f, 0x2f, 0x6f, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x61, 0x72, 0x64, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6f, 0x72, 0x61, 0x6d, 0x6e, 0x6f, 0x64, 0x65,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_oramnode_proto_rawDescOnce sync.Once
	file_oramnode_proto_rawDescData = file_oramnode_proto_rawDesc
)

func file_oramnode_proto_rawDescGZIP() []byte {
	file_oramnode_proto_rawDescOnce.Do(func() {
		file_oramnode_proto_rawDescData = protoimpl.X.CompressGZIP(file_oramnode_proto_rawDescData)
	})
	return file_oramnode_proto_rawDescData
}

var file_oramnode_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_oramnode_proto_goTypes = []interface{}{
	(*BlockRequest)(nil),         // 0: oramnode.BlockRequest
	(*ReadPathRequest)(nil),      // 1: oramnode.ReadPathRequest
	(*BlockResponse)(nil),        // 2: oramnode.BlockResponse
	(*ReadPathReply)(nil),        // 3: oramnode.ReadPathReply
	(*JoinRaftVoterRequest)(nil), // 4: oramnode.JoinRaftVoterRequest
	(*JoinRaftVoterReply)(nil),   // 5: oramnode.JoinRaftVoterReply
}
var file_oramnode_proto_depIdxs = []int32{
	0, // 0: oramnode.ReadPathRequest.requests:type_name -> oramnode.BlockRequest
	2, // 1: oramnode.ReadPathReply.responses:type_name -> oramnode.BlockResponse
	1, // 2: oramnode.OramNode.ReadPath:input_type -> oramnode.ReadPathRequest
	4, // 3: oramnode.OramNode.JoinRaftVoter:input_type -> oramnode.JoinRaftVoterRequest
	3, // 4: oramnode.OramNode.ReadPath:output_type -> oramnode.ReadPathReply
	5, // 5: oramnode.OramNode.JoinRaftVoter:output_type -> oramnode.JoinRaftVoterReply
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_oramnode_proto_init() }
func file_oramnode_proto_init() {
	if File_oramnode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_oramnode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockRequest); i {
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
		file_oramnode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPathRequest); i {
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
		file_oramnode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BlockResponse); i {
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
		file_oramnode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadPathReply); i {
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
		file_oramnode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRaftVoterRequest); i {
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
		file_oramnode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*JoinRaftVoterReply); i {
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
			RawDescriptor: file_oramnode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_oramnode_proto_goTypes,
		DependencyIndexes: file_oramnode_proto_depIdxs,
		MessageInfos:      file_oramnode_proto_msgTypes,
	}.Build()
	File_oramnode_proto = out.File
	file_oramnode_proto_rawDesc = nil
	file_oramnode_proto_goTypes = nil
	file_oramnode_proto_depIdxs = nil
}