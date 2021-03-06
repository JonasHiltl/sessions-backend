// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: packages/grpc/story/story.proto

package story

import (
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
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

type PublicStory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	PartyId       string   `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	UserId        string   `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Lat           float32  `protobuf:"fixed32,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          float32  `protobuf:"fixed32,6,opt,name=long,proto3" json:"long,omitempty"`
	Url           string   `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	TaggedFriends []string `protobuf:"bytes,7,rep,name=tagged_friends,json=taggedFriends,proto3" json:"tagged_friends,omitempty"`
	CreatedAt     string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PublicStory) Reset() {
	*x = PublicStory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicStory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicStory) ProtoMessage() {}

func (x *PublicStory) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicStory.ProtoReflect.Descriptor instead.
func (*PublicStory) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{0}
}

func (x *PublicStory) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicStory) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *PublicStory) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *PublicStory) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *PublicStory) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *PublicStory) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *PublicStory) GetTaggedFriends() []string {
	if x != nil {
		return x.TaggedFriends
	}
	return nil
}

func (x *PublicStory) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type CreateStoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterId   string   `protobuf:"bytes,1,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	PartyId       string   `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Url           string   `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
	Lat           float32  `protobuf:"fixed32,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          float32  `protobuf:"fixed32,5,opt,name=long,proto3" json:"long,omitempty"`
	TaggedFriends []string `protobuf:"bytes,6,rep,name=tagged_friends,json=taggedFriends,proto3" json:"tagged_friends,omitempty"`
}

func (x *CreateStoryRequest) Reset() {
	*x = CreateStoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateStoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateStoryRequest) ProtoMessage() {}

func (x *CreateStoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateStoryRequest.ProtoReflect.Descriptor instead.
func (*CreateStoryRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{1}
}

func (x *CreateStoryRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *CreateStoryRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *CreateStoryRequest) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *CreateStoryRequest) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *CreateStoryRequest) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *CreateStoryRequest) GetTaggedFriends() []string {
	if x != nil {
		return x.TaggedFriends
	}
	return nil
}

type GetStoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StoryId string `protobuf:"bytes,1,opt,name=story_id,json=storyId,proto3" json:"story_id,omitempty"`
}

func (x *GetStoryRequest) Reset() {
	*x = GetStoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStoryRequest) ProtoMessage() {}

func (x *GetStoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStoryRequest.ProtoReflect.Descriptor instead.
func (*GetStoryRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{2}
}

func (x *GetStoryRequest) GetStoryId() string {
	if x != nil {
		return x.StoryId
	}
	return ""
}

type DeleteStoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterId string `protobuf:"bytes,1,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	StoryId     string `protobuf:"bytes,2,opt,name=story_id,json=storyId,proto3" json:"story_id,omitempty"`
}

func (x *DeleteStoryRequest) Reset() {
	*x = DeleteStoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteStoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteStoryRequest) ProtoMessage() {}

func (x *DeleteStoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteStoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteStoryRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteStoryRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *DeleteStoryRequest) GetStoryId() string {
	if x != nil {
		return x.StoryId
	}
	return ""
}

type GetByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	NextPage string `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	Limit    uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetByUserRequest) Reset() {
	*x = GetByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserRequest) ProtoMessage() {}

func (x *GetByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByUserRequest.ProtoReflect.Descriptor instead.
func (*GetByUserRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{4}
}

func (x *GetByUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetByUserRequest) GetNextPage() string {
	if x != nil {
		return x.NextPage
	}
	return ""
}

func (x *GetByUserRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type GetByPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PartyId  string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	NextPage string `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
	Limit    uint32 `protobuf:"varint,3,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (x *GetByPartyRequest) Reset() {
	*x = GetByPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByPartyRequest) ProtoMessage() {}

func (x *GetByPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByPartyRequest.ProtoReflect.Descriptor instead.
func (*GetByPartyRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{5}
}

func (x *GetByPartyRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
}

func (x *GetByPartyRequest) GetNextPage() string {
	if x != nil {
		return x.NextPage
	}
	return ""
}

func (x *GetByPartyRequest) GetLimit() uint32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

type PresignURLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
}

func (x *PresignURLRequest) Reset() {
	*x = PresignURLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresignURLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresignURLRequest) ProtoMessage() {}

func (x *PresignURLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresignURLRequest.ProtoReflect.Descriptor instead.
func (*PresignURLRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{6}
}

func (x *PresignURLRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type PagedStories struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stories  []*PublicStory `protobuf:"bytes,1,rep,name=stories,proto3" json:"stories,omitempty"`
	NextPage string         `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *PagedStories) Reset() {
	*x = PagedStories{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedStories) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedStories) ProtoMessage() {}

func (x *PagedStories) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedStories.ProtoReflect.Descriptor instead.
func (*PagedStories) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{7}
}

func (x *PagedStories) GetStories() []*PublicStory {
	if x != nil {
		return x.Stories
	}
	return nil
}

func (x *PagedStories) GetNextPage() string {
	if x != nil {
		return x.NextPage
	}
	return ""
}

type PresignURLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *PresignURLResponse) Reset() {
	*x = PresignURLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_story_story_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PresignURLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PresignURLResponse) ProtoMessage() {}

func (x *PresignURLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_story_story_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PresignURLResponse.ProtoReflect.Descriptor instead.
func (*PresignURLResponse) Descriptor() ([]byte, []int) {
	return file_packages_grpc_story_story_proto_rawDescGZIP(), []int{8}
}

func (x *PresignURLResponse) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

var File_packages_grpc_story_story_proto protoreflect.FileDescriptor

var file_packages_grpc_story_story_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x1a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x01, 0x0a, 0x0b,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61, 0x67, 0x67, 0x65,
	0x64, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x0d, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0xb1, 0x01,
	0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x75, 0x72, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x61,
	0x67, 0x67, 0x65, 0x64, 0x5f, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x61, 0x67, 0x67, 0x65, 0x64, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x73, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x49, 0x64, 0x22,
	0x52, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x22, 0x61, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x61, 0x72, 0x74,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x22, 0x25, 0x0a, 0x11, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67,
	0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x59, 0x0a,
	0x0c, 0x50, 0x61, 0x67, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x2c, 0x0a,
	0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x6e,
	0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x26, 0x0a, 0x12, 0x50, 0x72, 0x65, 0x73,
	0x69, 0x67, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10,
	0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c,
	0x32, 0x8e, 0x03, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x3e, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79,
	0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x22,
	0x00, 0x12, 0x38, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x16, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x50, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x19, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x50,
	0x61, 0x67, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3d, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x18, 0x2e, 0x73, 0x74,
	0x6f, 0x72, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x42, 0x79, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x50, 0x61,
	0x67, 0x65, 0x64, 0x53, 0x74, 0x6f, 0x72, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0a,
	0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x52, 0x4c, 0x12, 0x18, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x2e, 0x50, 0x72, 0x65, 0x73, 0x69, 0x67, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x2e, 0x50, 0x72, 0x65,
	0x73, 0x69, 0x67, 0x6e, 0x55, 0x52, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6a, 0x6f, 0x6e, 0x61, 0x73, 0x68, 0x69, 0x6c, 0x74, 0x6c, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x61, 0x63, 0x6b,
	0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packages_grpc_story_story_proto_rawDescOnce sync.Once
	file_packages_grpc_story_story_proto_rawDescData = file_packages_grpc_story_story_proto_rawDesc
)

func file_packages_grpc_story_story_proto_rawDescGZIP() []byte {
	file_packages_grpc_story_story_proto_rawDescOnce.Do(func() {
		file_packages_grpc_story_story_proto_rawDescData = protoimpl.X.CompressGZIP(file_packages_grpc_story_story_proto_rawDescData)
	})
	return file_packages_grpc_story_story_proto_rawDescData
}

var file_packages_grpc_story_story_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_packages_grpc_story_story_proto_goTypes = []interface{}{
	(*PublicStory)(nil),            // 0: story.PublicStory
	(*CreateStoryRequest)(nil),     // 1: story.CreateStoryRequest
	(*GetStoryRequest)(nil),        // 2: story.GetStoryRequest
	(*DeleteStoryRequest)(nil),     // 3: story.DeleteStoryRequest
	(*GetByUserRequest)(nil),       // 4: story.GetByUserRequest
	(*GetByPartyRequest)(nil),      // 5: story.GetByPartyRequest
	(*PresignURLRequest)(nil),      // 6: story.PresignURLRequest
	(*PagedStories)(nil),           // 7: story.PagedStories
	(*PresignURLResponse)(nil),     // 8: story.PresignURLResponse
	(*common.MessageResponse)(nil), // 9: common.MessageResponse
}
var file_packages_grpc_story_story_proto_depIdxs = []int32{
	0, // 0: story.PagedStories.stories:type_name -> story.PublicStory
	1, // 1: story.StoryService.CreateStory:input_type -> story.CreateStoryRequest
	2, // 2: story.StoryService.GetStory:input_type -> story.GetStoryRequest
	3, // 3: story.StoryService.DeleteStory:input_type -> story.DeleteStoryRequest
	4, // 4: story.StoryService.GetByUser:input_type -> story.GetByUserRequest
	5, // 5: story.StoryService.GetByParty:input_type -> story.GetByPartyRequest
	6, // 6: story.StoryService.PresignURL:input_type -> story.PresignURLRequest
	0, // 7: story.StoryService.CreateStory:output_type -> story.PublicStory
	0, // 8: story.StoryService.GetStory:output_type -> story.PublicStory
	9, // 9: story.StoryService.DeleteStory:output_type -> common.MessageResponse
	7, // 10: story.StoryService.GetByUser:output_type -> story.PagedStories
	7, // 11: story.StoryService.GetByParty:output_type -> story.PagedStories
	8, // 12: story.StoryService.PresignURL:output_type -> story.PresignURLResponse
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_packages_grpc_story_story_proto_init() }
func file_packages_grpc_story_story_proto_init() {
	if File_packages_grpc_story_story_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packages_grpc_story_story_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicStory); i {
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
		file_packages_grpc_story_story_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateStoryRequest); i {
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
		file_packages_grpc_story_story_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStoryRequest); i {
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
		file_packages_grpc_story_story_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteStoryRequest); i {
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
		file_packages_grpc_story_story_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByUserRequest); i {
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
		file_packages_grpc_story_story_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByPartyRequest); i {
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
		file_packages_grpc_story_story_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresignURLRequest); i {
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
		file_packages_grpc_story_story_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedStories); i {
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
		file_packages_grpc_story_story_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PresignURLResponse); i {
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
			RawDescriptor: file_packages_grpc_story_story_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_packages_grpc_story_story_proto_goTypes,
		DependencyIndexes: file_packages_grpc_story_story_proto_depIdxs,
		MessageInfos:      file_packages_grpc_story_story_proto_msgTypes,
	}.Build()
	File_packages_grpc_story_story_proto = out.File
	file_packages_grpc_story_story_proto_rawDesc = nil
	file_packages_grpc_story_story_proto_goTypes = nil
	file_packages_grpc_story_story_proto_depIdxs = nil
}
