// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.15.8
// source: auth_room.proto

package protocol

import (
	livekit "github.com/livekit/protocol/livekit"
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

type ActiveRoomInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomTitle               string `protobuf:"bytes,1,opt,name=room_title,json=roomTitle,proto3" json:"room_title,omitempty"`
	RoomId                  string `protobuf:"bytes,2,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	Sid                     string `protobuf:"bytes,3,opt,name=sid,proto3" json:"sid,omitempty"`
	NumOfJoinedParticipants uint64 `protobuf:"varint,4,opt,name=num_of_joined_participants,json=numOfJoinedParticipants,proto3" json:"num_of_joined_participants,omitempty"`
	WebhookUrl              string `protobuf:"bytes,5,opt,name=webhook_url,json=webhookUrl,proto3" json:"webhook_url,omitempty"`
	Metadata                string `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	CreatedAt               int64  `protobuf:"varint,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ParentRoomId            string `protobuf:"bytes,8,opt,name=parent_room_id,json=parentRoomId,proto3" json:"parent_room_id,omitempty"`
}

func (x *ActiveRoomInfo) Reset() {
	*x = ActiveRoomInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_room_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRoomInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomInfo) ProtoMessage() {}

func (x *ActiveRoomInfo) ProtoReflect() protoreflect.Message {
	mi := &file_auth_room_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRoomInfo.ProtoReflect.Descriptor instead.
func (*ActiveRoomInfo) Descriptor() ([]byte, []int) {
	return file_auth_room_proto_rawDescGZIP(), []int{0}
}

func (x *ActiveRoomInfo) GetRoomTitle() string {
	if x != nil {
		return x.RoomTitle
	}
	return ""
}

func (x *ActiveRoomInfo) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *ActiveRoomInfo) GetSid() string {
	if x != nil {
		return x.Sid
	}
	return ""
}

func (x *ActiveRoomInfo) GetNumOfJoinedParticipants() uint64 {
	if x != nil {
		return x.NumOfJoinedParticipants
	}
	return 0
}

func (x *ActiveRoomInfo) GetWebhookUrl() string {
	if x != nil {
		return x.WebhookUrl
	}
	return ""
}

func (x *ActiveRoomInfo) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

func (x *ActiveRoomInfo) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *ActiveRoomInfo) GetParentRoomId() string {
	if x != nil {
		return x.ParentRoomId
	}
	return ""
}

type ActiveRoomWithParticipants struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ActiveRoomWithParticipants) Reset() {
	*x = ActiveRoomWithParticipants{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_room_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRoomWithParticipants) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomWithParticipants) ProtoMessage() {}

func (x *ActiveRoomWithParticipants) ProtoReflect() protoreflect.Message {
	mi := &file_auth_room_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRoomWithParticipants.ProtoReflect.Descriptor instead.
func (*ActiveRoomWithParticipants) Descriptor() ([]byte, []int) {
	return file_auth_room_proto_rawDescGZIP(), []int{1}
}

type ActiveRoomInfoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *ActiveRoomInfoRequest) Reset() {
	*x = ActiveRoomInfoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_room_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRoomInfoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomInfoRequest) ProtoMessage() {}

func (x *ActiveRoomInfoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_room_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRoomInfoRequest.ProtoReflect.Descriptor instead.
func (*ActiveRoomInfoRequest) Descriptor() ([]byte, []int) {
	return file_auth_room_proto_rawDescGZIP(), []int{2}
}

func (x *ActiveRoomInfoRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type ActiveRoomInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status          bool                       `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
	Msg             string                     `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	RoomInfo        *ActiveRoomInfo            `protobuf:"bytes,3,opt,name=room_info,json=roomInfo,proto3,oneof" json:"room_info,omitempty"`
	ParticipantInfo []*livekit.ParticipantInfo `protobuf:"bytes,4,rep,name=participant_info,json=participantInfo,proto3" json:"participant_info,omitempty"`
}

func (x *ActiveRoomInfoResponse) Reset() {
	*x = ActiveRoomInfoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_room_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ActiveRoomInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ActiveRoomInfoResponse) ProtoMessage() {}

func (x *ActiveRoomInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_room_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ActiveRoomInfoResponse.ProtoReflect.Descriptor instead.
func (*ActiveRoomInfoResponse) Descriptor() ([]byte, []int) {
	return file_auth_room_proto_rawDescGZIP(), []int{3}
}

func (x *ActiveRoomInfoResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

func (x *ActiveRoomInfoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *ActiveRoomInfoResponse) GetRoomInfo() *ActiveRoomInfo {
	if x != nil {
		return x.RoomInfo
	}
	return nil
}

func (x *ActiveRoomInfoResponse) GetParticipantInfo() []*livekit.ParticipantInfo {
	if x != nil {
		return x.ParticipantInfo
	}
	return nil
}

type IsRoomActiveRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *IsRoomActiveRequest) Reset() {
	*x = IsRoomActiveRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_room_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsRoomActiveRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRoomActiveRequest) ProtoMessage() {}

func (x *IsRoomActiveRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_room_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsRoomActiveRequest.ProtoReflect.Descriptor instead.
func (*IsRoomActiveRequest) Descriptor() ([]byte, []int) {
	return file_auth_room_proto_rawDescGZIP(), []int{4}
}

func (x *IsRoomActiveRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

type IsRoomActiveResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *IsRoomActiveResponse) Reset() {
	*x = IsRoomActiveResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_room_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsRoomActiveResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsRoomActiveResponse) ProtoMessage() {}

func (x *IsRoomActiveResponse) ProtoReflect() protoreflect.Message {
	mi := &file_auth_room_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsRoomActiveResponse.ProtoReflect.Descriptor instead.
func (*IsRoomActiveResponse) Descriptor() ([]byte, []int) {
	return file_auth_room_proto_rawDescGZIP(), []int{5}
}

func (x *IsRoomActiveResponse) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type EndRoomRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
}

func (x *EndRoomRequest) Reset() {
	*x = EndRoomRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_auth_room_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EndRoomRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EndRoomRequest) ProtoMessage() {}

func (x *EndRoomRequest) ProtoReflect() protoreflect.Message {
	mi := &file_auth_room_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EndRoomRequest.ProtoReflect.Descriptor instead.
func (*EndRoomRequest) Descriptor() ([]byte, []int) {
	return file_auth_room_proto_rawDescGZIP(), []int{6}
}

func (x *EndRoomRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

var File_auth_room_proto protoreflect.FileDescriptor

var file_auth_room_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x61, 0x75, 0x74, 0x68, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x08, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x1a, 0x14, 0x6c, 0x69, 0x76,
	0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x99, 0x02, 0x0a, 0x0e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x54, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03,
	0x73, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x73, 0x69, 0x64, 0x12, 0x3b,
	0x0a, 0x1a, 0x6e, 0x75, 0x6d, 0x5f, 0x6f, 0x66, 0x5f, 0x6a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x5f,
	0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x17, 0x6e, 0x75, 0x6d, 0x4f, 0x66, 0x4a, 0x6f, 0x69, 0x6e, 0x65, 0x64, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x77,
	0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f, 0x6b, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x24, 0x0a, 0x0e, 0x70, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0c, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0x1c, 0x0a,
	0x1a, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x57, 0x69, 0x74, 0x68, 0x50,
	0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x73, 0x22, 0x30, 0x0a, 0x15, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x22, 0xd1, 0x01,
	0x0a, 0x16, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d,
	0x73, 0x67, 0x12, 0x3a, 0x0a, 0x09, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2e, 0x41, 0x63, 0x74, 0x69, 0x76, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x08, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x6e, 0x66, 0x6f, 0x88, 0x01, 0x01, 0x12, 0x43,
	0x0a, 0x10, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b,
	0x69, 0x74, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x0f, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x6e, 0x66,
	0x6f, 0x22, 0x2e, 0x0a, 0x13, 0x49, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49,
	0x64, 0x22, 0x2e, 0x0a, 0x14, 0x49, 0x73, 0x52, 0x6f, 0x6f, 0x6d, 0x41, 0x63, 0x74, 0x69, 0x76,
	0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x29, 0x0a, 0x0e, 0x45, 0x6e, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x42, 0x21, 0x5a, 0x1f,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x6c, 0x6c, 0x69,
	0x62, 0x69, 0x6c, 0x69, 0x73, 0x6d, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_auth_room_proto_rawDescOnce sync.Once
	file_auth_room_proto_rawDescData = file_auth_room_proto_rawDesc
)

func file_auth_room_proto_rawDescGZIP() []byte {
	file_auth_room_proto_rawDescOnce.Do(func() {
		file_auth_room_proto_rawDescData = protoimpl.X.CompressGZIP(file_auth_room_proto_rawDescData)
	})
	return file_auth_room_proto_rawDescData
}

var file_auth_room_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_auth_room_proto_goTypes = []interface{}{
	(*ActiveRoomInfo)(nil),             // 0: protocol.ActiveRoomInfo
	(*ActiveRoomWithParticipants)(nil), // 1: protocol.ActiveRoomWithParticipants
	(*ActiveRoomInfoRequest)(nil),      // 2: protocol.ActiveRoomInfoRequest
	(*ActiveRoomInfoResponse)(nil),     // 3: protocol.ActiveRoomInfoResponse
	(*IsRoomActiveRequest)(nil),        // 4: protocol.IsRoomActiveRequest
	(*IsRoomActiveResponse)(nil),       // 5: protocol.IsRoomActiveResponse
	(*EndRoomRequest)(nil),             // 6: protocol.EndRoomRequest
	(*livekit.ParticipantInfo)(nil),    // 7: livekit.ParticipantInfo
}
var file_auth_room_proto_depIdxs = []int32{
	0, // 0: protocol.ActiveRoomInfoResponse.room_info:type_name -> protocol.ActiveRoomInfo
	7, // 1: protocol.ActiveRoomInfoResponse.participant_info:type_name -> livekit.ParticipantInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_auth_room_proto_init() }
func file_auth_room_proto_init() {
	if File_auth_room_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_auth_room_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRoomInfo); i {
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
		file_auth_room_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRoomWithParticipants); i {
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
		file_auth_room_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRoomInfoRequest); i {
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
		file_auth_room_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ActiveRoomInfoResponse); i {
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
		file_auth_room_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsRoomActiveRequest); i {
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
		file_auth_room_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsRoomActiveResponse); i {
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
		file_auth_room_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EndRoomRequest); i {
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
	file_auth_room_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_auth_room_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_auth_room_proto_goTypes,
		DependencyIndexes: file_auth_room_proto_depIdxs,
		MessageInfos:      file_auth_room_proto_msgTypes,
	}.Build()
	File_auth_room_proto = out.File
	file_auth_room_proto_rawDesc = nil
	file_auth_room_proto_goTypes = nil
	file_auth_room_proto_depIdxs = nil
}
