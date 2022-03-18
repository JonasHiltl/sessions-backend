// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.6.1
// source: party/party.proto

package gp

import (
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type CreatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string  `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Lat      float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Long     float32 `protobuf:"fixed32,3,opt,name=long,proto3" json:"long,omitempty"`
	IsPublic bool    `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
}

func (x *CreatePartyRequest) Reset() {
	*x = CreatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyRequest) ProtoMessage() {}

func (x *CreatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_party_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartyRequest.ProtoReflect.Descriptor instead.
func (*CreatePartyRequest) Descriptor() ([]byte, []int) {
	return file_party_party_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePartyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *CreatePartyRequest) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *CreatePartyRequest) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *CreatePartyRequest) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

type UpdatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PId   string  `protobuf:"bytes,1,opt,name=p_id,json=pId,proto3" json:"p_id,omitempty"`
	Title string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Lat   float32 `protobuf:"fixed32,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Long  float32 `protobuf:"fixed32,4,opt,name=long,proto3" json:"long,omitempty"`
}

func (x *UpdatePartyRequest) Reset() {
	*x = UpdatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartyRequest) ProtoMessage() {}

func (x *UpdatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_party_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdatePartyRequest.ProtoReflect.Descriptor instead.
func (*UpdatePartyRequest) Descriptor() ([]byte, []int) {
	return file_party_party_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePartyRequest) GetPId() string {
	if x != nil {
		return x.PId
	}
	return ""
}

func (x *UpdatePartyRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *UpdatePartyRequest) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *UpdatePartyRequest) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

type GeoSearchRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Precision int64   `protobuf:"varint,1,opt,name=precision,proto3" json:"precision,omitempty"`
	Lat       float32 `protobuf:"fixed32,2,opt,name=lat,proto3" json:"lat,omitempty"`
	Long      float32 `protobuf:"fixed32,3,opt,name=long,proto3" json:"long,omitempty"`
	NextPage  string  `protobuf:"bytes,4,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *GeoSearchRequest) Reset() {
	*x = GeoSearchRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoSearchRequest) ProtoMessage() {}

func (x *GeoSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeoSearchRequest.ProtoReflect.Descriptor instead.
func (*GeoSearchRequest) Descriptor() ([]byte, []int) {
	return file_party_party_proto_rawDescGZIP(), []int{2}
}

func (x *GeoSearchRequest) GetPrecision() int64 {
	if x != nil {
		return x.Precision
	}
	return 0
}

func (x *GeoSearchRequest) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *GeoSearchRequest) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *GeoSearchRequest) GetNextPage() string {
	if x != nil {
		return x.NextPage
	}
	return ""
}

type GetPartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PId string `protobuf:"bytes,1,opt,name=p_id,json=pId,proto3" json:"p_id,omitempty"`
}

func (x *GetPartyRequest) Reset() {
	*x = GetPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_party_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyRequest) ProtoMessage() {}

func (x *GetPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_party_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartyRequest.ProtoReflect.Descriptor instead.
func (*GetPartyRequest) Descriptor() ([]byte, []int) {
	return file_party_party_proto_rawDescGZIP(), []int{3}
}

func (x *GetPartyRequest) GetPId() string {
	if x != nil {
		return x.PId
	}
	return ""
}

type GetByUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UId      string `protobuf:"bytes,1,opt,name=u_id,json=uId,proto3" json:"u_id,omitempty"`
	NextPage string `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *GetByUserRequest) Reset() {
	*x = GetByUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_party_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserRequest) ProtoMessage() {}

func (x *GetByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_party_party_proto_msgTypes[4]
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
	return file_party_party_proto_rawDescGZIP(), []int{4}
}

func (x *GetByUserRequest) GetUId() string {
	if x != nil {
		return x.UId
	}
	return ""
}

func (x *GetByUserRequest) GetNextPage() string {
	if x != nil {
		return x.NextPage
	}
	return ""
}

type PagedParties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parties  []*PublicParty `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
	NextPage string         `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *PagedParties) Reset() {
	*x = PagedParties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_party_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedParties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedParties) ProtoMessage() {}

func (x *PagedParties) ProtoReflect() protoreflect.Message {
	mi := &file_party_party_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PagedParties.ProtoReflect.Descriptor instead.
func (*PagedParties) Descriptor() ([]byte, []int) {
	return file_party_party_proto_rawDescGZIP(), []int{5}
}

func (x *PagedParties) GetParties() []*PublicParty {
	if x != nil {
		return x.Parties
	}
	return nil
}

func (x *PagedParties) GetNextPage() string {
	if x != nil {
		return x.NextPage
	}
	return ""
}

type PublicParty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UId       string   `protobuf:"bytes,2,opt,name=u_id,json=uId,proto3" json:"u_id,omitempty"`
	Title     string   `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IsPublic  bool     `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	Lat       float32  `protobuf:"fixed32,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Long      float32  `protobuf:"fixed32,6,opt,name=long,proto3" json:"long,omitempty"`
	Stories   []string `protobuf:"bytes,7,rep,name=stories,proto3" json:"stories,omitempty"`
	CreatedAt string   `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *PublicParty) Reset() {
	*x = PublicParty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_party_party_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PublicParty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PublicParty) ProtoMessage() {}

func (x *PublicParty) ProtoReflect() protoreflect.Message {
	mi := &file_party_party_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PublicParty.ProtoReflect.Descriptor instead.
func (*PublicParty) Descriptor() ([]byte, []int) {
	return file_party_party_proto_rawDescGZIP(), []int{6}
}

func (x *PublicParty) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PublicParty) GetUId() string {
	if x != nil {
		return x.UId
	}
	return ""
}

func (x *PublicParty) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *PublicParty) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *PublicParty) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *PublicParty) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *PublicParty) GetStories() []string {
	if x != nil {
		return x.Stories
	}
	return nil
}

func (x *PublicParty) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_party_party_proto protoreflect.FileDescriptor

var file_party_party_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6d, 0x0a,
	0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c,
	0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12,
	0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x22, 0x63, 0x0a, 0x12,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x70, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e,
	0x67, 0x22, 0x73, 0x0a, 0x10, 0x47, 0x65, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x24, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x11, 0x0a, 0x04, 0x70, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x70, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x10,
	0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x11, 0x0a, 0x04, 0x75, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x75, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65,
	0x22, 0x59, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x2c, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x12, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1b,
	0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0xc2, 0x01, 0x0a, 0x0b,
	0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x11, 0x0a, 0x04, 0x75,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69,
	0x63, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03,
	0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69,
	0x65, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x73, 0x74, 0x6f, 0x72, 0x69, 0x65,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x32, 0xe1, 0x03, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x4a, 0x0a, 0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79,
	0x12, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50, 0x61, 0x72, 0x74, 0x79, 0x22,
	0x0c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x06, 0x22, 0x01, 0x2f, 0x3a, 0x01, 0x2a, 0x12, 0x47, 0x0a,
	0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x12, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09, 0x12, 0x07, 0x2f,
	0x7b, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x50, 0x0a, 0x0b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x12, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x22, 0x12, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0c, 0x32, 0x07, 0x2f, 0x7b,
	0x70, 0x5f, 0x69, 0x64, 0x7d, 0x3a, 0x01, 0x2a, 0x12, 0x4f, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x17, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0f, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x09,
	0x2a, 0x07, 0x2f, 0x7b, 0x70, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x4f, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x7b, 0x75, 0x5f, 0x69, 0x64, 0x7d, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65,
	0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e,
	0x47, 0x65, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x13, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x64, 0x50, 0x61,
	0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0x0d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x07, 0x12, 0x05, 0x2f,
	0x6e, 0x65, 0x61, 0x72, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x6e, 0x61, 0x73, 0x68, 0x69, 0x6c, 0x74, 0x6c, 0x2f, 0x73, 0x65,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70,
	0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x67, 0x70, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_party_party_proto_rawDescOnce sync.Once
	file_party_party_proto_rawDescData = file_party_party_proto_rawDesc
)

func file_party_party_proto_rawDescGZIP() []byte {
	file_party_party_proto_rawDescOnce.Do(func() {
		file_party_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_party_party_proto_rawDescData)
	})
	return file_party_party_proto_rawDescData
}

var file_party_party_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_party_party_proto_goTypes = []interface{}{
	(*CreatePartyRequest)(nil),     // 0: party.CreatePartyRequest
	(*UpdatePartyRequest)(nil),     // 1: party.UpdatePartyRequest
	(*GeoSearchRequest)(nil),       // 2: party.GeoSearchRequest
	(*GetPartyRequest)(nil),        // 3: party.GetPartyRequest
	(*GetByUserRequest)(nil),       // 4: party.GetByUserRequest
	(*PagedParties)(nil),           // 5: party.PagedParties
	(*PublicParty)(nil),            // 6: party.PublicParty
	(*common.MessageResponse)(nil), // 7: common.MessageResponse
}
var file_party_party_proto_depIdxs = []int32{
	6, // 0: party.PagedParties.parties:type_name -> party.PublicParty
	0, // 1: party.PartyService.CreateParty:input_type -> party.CreatePartyRequest
	3, // 2: party.PartyService.GetParty:input_type -> party.GetPartyRequest
	1, // 3: party.PartyService.UpdateParty:input_type -> party.UpdatePartyRequest
	3, // 4: party.PartyService.DeleteParty:input_type -> party.GetPartyRequest
	4, // 5: party.PartyService.GetByUser:input_type -> party.GetByUserRequest
	2, // 6: party.PartyService.GeoSearch:input_type -> party.GeoSearchRequest
	6, // 7: party.PartyService.CreateParty:output_type -> party.PublicParty
	6, // 8: party.PartyService.GetParty:output_type -> party.PublicParty
	6, // 9: party.PartyService.UpdateParty:output_type -> party.PublicParty
	7, // 10: party.PartyService.DeleteParty:output_type -> common.MessageResponse
	5, // 11: party.PartyService.GetByUser:output_type -> party.PagedParties
	5, // 12: party.PartyService.GeoSearch:output_type -> party.PagedParties
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_party_party_proto_init() }
func file_party_party_proto_init() {
	if File_party_party_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_party_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePartyRequest); i {
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
		file_party_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdatePartyRequest); i {
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
		file_party_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GeoSearchRequest); i {
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
		file_party_party_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPartyRequest); i {
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
		file_party_party_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_party_party_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PagedParties); i {
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
		file_party_party_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PublicParty); i {
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
			RawDescriptor: file_party_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_party_party_proto_goTypes,
		DependencyIndexes: file_party_party_proto_depIdxs,
		MessageInfos:      file_party_party_proto_msgTypes,
	}.Build()
	File_party_party_proto = out.File
	file_party_party_proto_rawDesc = nil
	file_party_party_proto_goTypes = nil
	file_party_party_proto_depIdxs = nil
}
