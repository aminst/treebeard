// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: shardnode.proto

package shardnode

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

type RequestBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadRequests  []*ReadRequest  `protobuf:"bytes,1,rep,name=read_requests,json=readRequests,proto3" json:"read_requests,omitempty"`
	WriteRequests []*WriteRequest `protobuf:"bytes,2,rep,name=write_requests,json=writeRequests,proto3" json:"write_requests,omitempty"`
}

func (x *RequestBatch) Reset() {
	*x = RequestBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RequestBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestBatch) ProtoMessage() {}

func (x *RequestBatch) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestBatch.ProtoReflect.Descriptor instead.
func (*RequestBatch) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{0}
}

func (x *RequestBatch) GetReadRequests() []*ReadRequest {
	if x != nil {
		return x.ReadRequests
	}
	return nil
}

func (x *RequestBatch) GetWriteRequests() []*WriteRequest {
	if x != nil {
		return x.WriteRequests
	}
	return nil
}

type ReplyBatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReadReplies  []*ReadReply  `protobuf:"bytes,1,rep,name=read_replies,json=readReplies,proto3" json:"read_replies,omitempty"`
	WriteReplies []*WriteReply `protobuf:"bytes,2,rep,name=write_replies,json=writeReplies,proto3" json:"write_replies,omitempty"`
}

func (x *ReplyBatch) Reset() {
	*x = ReplyBatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReplyBatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReplyBatch) ProtoMessage() {}

func (x *ReplyBatch) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReplyBatch.ProtoReflect.Descriptor instead.
func (*ReplyBatch) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{1}
}

func (x *ReplyBatch) GetReadReplies() []*ReadReply {
	if x != nil {
		return x.ReadReplies
	}
	return nil
}

func (x *ReplyBatch) GetWriteReplies() []*WriteReply {
	if x != nil {
		return x.WriteReplies
	}
	return nil
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Block     string `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{2}
}

func (x *ReadRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ReadRequest) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

type ReadReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Value     string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *ReadReply) Reset() {
	*x = ReadReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadReply) ProtoMessage() {}

func (x *ReadReply) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadReply.ProtoReflect.Descriptor instead.
func (*ReadReply) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{3}
}

func (x *ReadReply) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *ReadReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type WriteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Block     string `protobuf:"bytes,2,opt,name=block,proto3" json:"block,omitempty"`
	Value     string `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{4}
}

func (x *WriteRequest) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *WriteRequest) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *WriteRequest) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type WriteReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequestId string `protobuf:"bytes,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	Success   bool   `protobuf:"varint,2,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *WriteReply) Reset() {
	*x = WriteReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteReply) ProtoMessage() {}

func (x *WriteReply) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteReply.ProtoReflect.Descriptor instead.
func (*WriteReply) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{5}
}

func (x *WriteReply) GetRequestId() string {
	if x != nil {
		return x.RequestId
	}
	return ""
}

func (x *WriteReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
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
		mi := &file_shardnode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRaftVoterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRaftVoterRequest) ProtoMessage() {}

func (x *JoinRaftVoterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[6]
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
	return file_shardnode_proto_rawDescGZIP(), []int{6}
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
		mi := &file_shardnode_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *JoinRaftVoterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*JoinRaftVoterReply) ProtoMessage() {}

func (x *JoinRaftVoterReply) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[7]
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
	return file_shardnode_proto_rawDescGZIP(), []int{7}
}

func (x *JoinRaftVoterReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

type SendBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MaxBlocks int32 `protobuf:"varint,1,opt,name=maxBlocks,proto3" json:"maxBlocks,omitempty"`
	StorageId int32 `protobuf:"varint,3,opt,name=storage_id,json=storageId,proto3" json:"storage_id,omitempty"`
}

func (x *SendBlocksRequest) Reset() {
	*x = SendBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBlocksRequest) ProtoMessage() {}

func (x *SendBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBlocksRequest.ProtoReflect.Descriptor instead.
func (*SendBlocksRequest) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{8}
}

func (x *SendBlocksRequest) GetMaxBlocks() int32 {
	if x != nil {
		return x.MaxBlocks
	}
	return 0
}

func (x *SendBlocksRequest) GetStorageId() int32 {
	if x != nil {
		return x.StorageId
	}
	return 0
}

type Block struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block string `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	Path  int32  `protobuf:"varint,3,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Block) Reset() {
	*x = Block{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Block) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Block) ProtoMessage() {}

func (x *Block) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Block.ProtoReflect.Descriptor instead.
func (*Block) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{9}
}

func (x *Block) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *Block) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

func (x *Block) GetPath() int32 {
	if x != nil {
		return x.Path
	}
	return 0
}

type SendBlocksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Blocks []*Block `protobuf:"bytes,1,rep,name=blocks,proto3" json:"blocks,omitempty"`
}

func (x *SendBlocksReply) Reset() {
	*x = SendBlocksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendBlocksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendBlocksReply) ProtoMessage() {}

func (x *SendBlocksReply) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendBlocksReply.ProtoReflect.Descriptor instead.
func (*SendBlocksReply) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{10}
}

func (x *SendBlocksReply) GetBlocks() []*Block {
	if x != nil {
		return x.Blocks
	}
	return nil
}

// it represents both acks and nacks
type Ack struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Block string `protobuf:"bytes,1,opt,name=block,proto3" json:"block,omitempty"`
	IsAck bool   `protobuf:"varint,2,opt,name=is_ack,json=isAck,proto3" json:"is_ack,omitempty"` // if flase: nack
}

func (x *Ack) Reset() {
	*x = Ack{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ack) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ack) ProtoMessage() {}

func (x *Ack) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ack.ProtoReflect.Descriptor instead.
func (*Ack) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{11}
}

func (x *Ack) GetBlock() string {
	if x != nil {
		return x.Block
	}
	return ""
}

func (x *Ack) GetIsAck() bool {
	if x != nil {
		return x.IsAck
	}
	return false
}

type AckSentBlocksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Acks []*Ack `protobuf:"bytes,1,rep,name=acks,proto3" json:"acks,omitempty"`
}

func (x *AckSentBlocksRequest) Reset() {
	*x = AckSentBlocksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckSentBlocksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckSentBlocksRequest) ProtoMessage() {}

func (x *AckSentBlocksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckSentBlocksRequest.ProtoReflect.Descriptor instead.
func (*AckSentBlocksRequest) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{12}
}

func (x *AckSentBlocksRequest) GetAcks() []*Ack {
	if x != nil {
		return x.Acks
	}
	return nil
}

type AckSentBlocksReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *AckSentBlocksReply) Reset() {
	*x = AckSentBlocksReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shardnode_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AckSentBlocksReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AckSentBlocksReply) ProtoMessage() {}

func (x *AckSentBlocksReply) ProtoReflect() protoreflect.Message {
	mi := &file_shardnode_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AckSentBlocksReply.ProtoReflect.Descriptor instead.
func (*AckSentBlocksReply) Descriptor() ([]byte, []int) {
	return file_shardnode_proto_rawDescGZIP(), []int{13}
}

func (x *AckSentBlocksReply) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_shardnode_proto protoreflect.FileDescriptor

var file_shardnode_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x09, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x22, 0x8b, 0x01, 0x0a,
	0x0c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x3b, 0x0a,
	0x0d, 0x72, 0x65, 0x61, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65,
	0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0c, 0x72, 0x65,
	0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x12, 0x3e, 0x0a, 0x0e, 0x77, 0x72,
	0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0d, 0x77, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x81, 0x01, 0x0a, 0x0a, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x12, 0x37, 0x0a, 0x0c, 0x72, 0x65, 0x61,
	0x64, 0x5f, 0x72, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x61, 0x64,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x0b, 0x72, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x69,
	0x65, 0x73, 0x12, 0x3a, 0x0a, 0x0d, 0x77, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6c,
	0x69, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x52, 0x0c, 0x77, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x69, 0x65, 0x73, 0x22, 0x42,
	0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f,
	0x63, 0x6b, 0x22, 0x40, 0x0a, 0x09, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x22, 0x59, 0x0a, 0x0c, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22,
	0x45, 0x0a, 0x0a, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07,
	0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73,
	0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x4c, 0x0a, 0x14, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61,
	0x66, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x06, 0x6e, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65,
	0x41, 0x64, 0x64, 0x72, 0x22, 0x2e, 0x0a, 0x12, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66, 0x74,
	0x56, 0x6f, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x22, 0x50, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x6d, 0x61, 0x78,
	0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6d, 0x61,
	0x78, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x49, 0x64, 0x22, 0x47, 0x0a, 0x05, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x12,
	0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22,
	0x3b, 0x0a, 0x0f, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x52, 0x06, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x22, 0x32, 0x0a, 0x03,
	0x41, 0x63, 0x6b, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x62, 0x6c, 0x6f, 0x63, 0x6b, 0x12, 0x15, 0x0a, 0x06, 0x69, 0x73, 0x5f,
	0x61, 0x63, 0x6b, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x69, 0x73, 0x41, 0x63, 0x6b,
	0x22, 0x3a, 0x0a, 0x14, 0x41, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x04, 0x61, 0x63, 0x6b, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x52, 0x04, 0x61, 0x63, 0x6b, 0x73, 0x22, 0x2e, 0x0a, 0x12,
	0x41, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xbb, 0x02, 0x0a,
	0x09, 0x53, 0x68, 0x61, 0x72, 0x64, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3e, 0x0a, 0x0a, 0x42, 0x61,
	0x74, 0x63, 0x68, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x42, 0x61, 0x74, 0x63,
	0x68, 0x1a, 0x15, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x42, 0x61, 0x74, 0x63, 0x68, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0a, 0x53, 0x65,
	0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1c, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x41, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x42,
	0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64,
	0x65, 0x2e, 0x41, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f,
	0x64, 0x65, 0x2e, 0x41, 0x63, 0x6b, 0x53, 0x65, 0x6e, 0x74, 0x42, 0x6c, 0x6f, 0x63, 0x6b, 0x73,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0d, 0x4a, 0x6f, 0x69, 0x6e, 0x52,
	0x61, 0x66, 0x74, 0x56, 0x6f, 0x74, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66, 0x74, 0x56, 0x6f, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x73, 0x68, 0x61, 0x72,
	0x64, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x4a, 0x6f, 0x69, 0x6e, 0x52, 0x61, 0x66, 0x74, 0x56, 0x6f,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x32, 0x5a, 0x30, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x64, 0x73, 0x67, 0x2d, 0x75, 0x77, 0x61,
	0x74, 0x65, 0x72, 0x6c, 0x6f, 0x6f, 0x2f, 0x6f, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x61, 0x72, 0x64,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x68, 0x61, 0x72, 0x64, 0x6e, 0x6f, 0x64, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shardnode_proto_rawDescOnce sync.Once
	file_shardnode_proto_rawDescData = file_shardnode_proto_rawDesc
)

func file_shardnode_proto_rawDescGZIP() []byte {
	file_shardnode_proto_rawDescOnce.Do(func() {
		file_shardnode_proto_rawDescData = protoimpl.X.CompressGZIP(file_shardnode_proto_rawDescData)
	})
	return file_shardnode_proto_rawDescData
}

var file_shardnode_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_shardnode_proto_goTypes = []interface{}{
	(*RequestBatch)(nil),         // 0: shardnode.RequestBatch
	(*ReplyBatch)(nil),           // 1: shardnode.ReplyBatch
	(*ReadRequest)(nil),          // 2: shardnode.ReadRequest
	(*ReadReply)(nil),            // 3: shardnode.ReadReply
	(*WriteRequest)(nil),         // 4: shardnode.WriteRequest
	(*WriteReply)(nil),           // 5: shardnode.WriteReply
	(*JoinRaftVoterRequest)(nil), // 6: shardnode.JoinRaftVoterRequest
	(*JoinRaftVoterReply)(nil),   // 7: shardnode.JoinRaftVoterReply
	(*SendBlocksRequest)(nil),    // 8: shardnode.SendBlocksRequest
	(*Block)(nil),                // 9: shardnode.Block
	(*SendBlocksReply)(nil),      // 10: shardnode.SendBlocksReply
	(*Ack)(nil),                  // 11: shardnode.Ack
	(*AckSentBlocksRequest)(nil), // 12: shardnode.AckSentBlocksRequest
	(*AckSentBlocksReply)(nil),   // 13: shardnode.AckSentBlocksReply
}
var file_shardnode_proto_depIdxs = []int32{
	2,  // 0: shardnode.RequestBatch.read_requests:type_name -> shardnode.ReadRequest
	4,  // 1: shardnode.RequestBatch.write_requests:type_name -> shardnode.WriteRequest
	3,  // 2: shardnode.ReplyBatch.read_replies:type_name -> shardnode.ReadReply
	5,  // 3: shardnode.ReplyBatch.write_replies:type_name -> shardnode.WriteReply
	9,  // 4: shardnode.SendBlocksReply.blocks:type_name -> shardnode.Block
	11, // 5: shardnode.AckSentBlocksRequest.acks:type_name -> shardnode.Ack
	0,  // 6: shardnode.ShardNode.BatchQuery:input_type -> shardnode.RequestBatch
	8,  // 7: shardnode.ShardNode.SendBlocks:input_type -> shardnode.SendBlocksRequest
	12, // 8: shardnode.ShardNode.AckSentBlocks:input_type -> shardnode.AckSentBlocksRequest
	6,  // 9: shardnode.ShardNode.JoinRaftVoter:input_type -> shardnode.JoinRaftVoterRequest
	1,  // 10: shardnode.ShardNode.BatchQuery:output_type -> shardnode.ReplyBatch
	10, // 11: shardnode.ShardNode.SendBlocks:output_type -> shardnode.SendBlocksReply
	13, // 12: shardnode.ShardNode.AckSentBlocks:output_type -> shardnode.AckSentBlocksReply
	7,  // 13: shardnode.ShardNode.JoinRaftVoter:output_type -> shardnode.JoinRaftVoterReply
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_shardnode_proto_init() }
func file_shardnode_proto_init() {
	if File_shardnode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shardnode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RequestBatch); i {
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
		file_shardnode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReplyBatch); i {
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
		file_shardnode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_shardnode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadReply); i {
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
		file_shardnode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRequest); i {
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
		file_shardnode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteReply); i {
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
		file_shardnode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_shardnode_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
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
		file_shardnode_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBlocksRequest); i {
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
		file_shardnode_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Block); i {
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
		file_shardnode_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendBlocksReply); i {
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
		file_shardnode_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ack); i {
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
		file_shardnode_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckSentBlocksRequest); i {
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
		file_shardnode_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AckSentBlocksReply); i {
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
			RawDescriptor: file_shardnode_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shardnode_proto_goTypes,
		DependencyIndexes: file_shardnode_proto_depIdxs,
		MessageInfos:      file_shardnode_proto_msgTypes,
	}.Build()
	File_shardnode_proto = out.File
	file_shardnode_proto_rawDesc = nil
	file_shardnode_proto_goTypes = nil
	file_shardnode_proto_depIdxs = nil
}
