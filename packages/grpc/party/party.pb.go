// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: packages/grpc/party/party.proto

package party

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

type CreatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterId   string  `protobuf:"bytes,1,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	Title         string  `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Lat           float32 `protobuf:"fixed32,3,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          float32 `protobuf:"fixed32,4,opt,name=long,proto3" json:"long,omitempty"`
	IsPublic      bool    `protobuf:"varint,5,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	StreetAddress string  `protobuf:"bytes,6,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	PostalCode    string  `protobuf:"bytes,7,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State         string  `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	Country       string  `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	StartDate     string  `protobuf:"bytes,10,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
}

func (x *CreatePartyRequest) Reset() {
	*x = CreatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_party_party_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartyRequest) ProtoMessage() {}

func (x *CreatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[0]
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
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePartyRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
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

func (x *CreatePartyRequest) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *CreatePartyRequest) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *CreatePartyRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *CreatePartyRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *CreatePartyRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

type UpdatePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterId   string  `protobuf:"bytes,1,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	PartyId       string  `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
	Title         string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Lat           float32 `protobuf:"fixed32,4,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          float32 `protobuf:"fixed32,5,opt,name=long,proto3" json:"long,omitempty"`
	StreetAddress string  `protobuf:"bytes,6,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	PostalCode    string  `protobuf:"bytes,7,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State         string  `protobuf:"bytes,8,opt,name=state,proto3" json:"state,omitempty"`
	Country       string  `protobuf:"bytes,9,opt,name=country,proto3" json:"country,omitempty"`
	StartDate     string  `protobuf:"bytes,10,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
}

func (x *UpdatePartyRequest) Reset() {
	*x = UpdatePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_party_party_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdatePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdatePartyRequest) ProtoMessage() {}

func (x *UpdatePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[1]
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
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{1}
}

func (x *UpdatePartyRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *UpdatePartyRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
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

func (x *UpdatePartyRequest) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *UpdatePartyRequest) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *UpdatePartyRequest) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *UpdatePartyRequest) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *UpdatePartyRequest) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

type DeletePartyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RequesterId string `protobuf:"bytes,1,opt,name=requester_id,json=requesterId,proto3" json:"requester_id,omitempty"`
	PartyId     string `protobuf:"bytes,2,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *DeletePartyRequest) Reset() {
	*x = DeletePartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_party_party_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeletePartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeletePartyRequest) ProtoMessage() {}

func (x *DeletePartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeletePartyRequest.ProtoReflect.Descriptor instead.
func (*DeletePartyRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{2}
}

func (x *DeletePartyRequest) GetRequesterId() string {
	if x != nil {
		return x.RequesterId
	}
	return ""
}

func (x *DeletePartyRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
	}
	return ""
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
		mi := &file_packages_grpc_party_party_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GeoSearchRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeoSearchRequest) ProtoMessage() {}

func (x *GeoSearchRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[3]
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
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{3}
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

	PartyId string `protobuf:"bytes,1,opt,name=party_id,json=partyId,proto3" json:"party_id,omitempty"`
}

func (x *GetPartyRequest) Reset() {
	*x = GetPartyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_party_party_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPartyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartyRequest) ProtoMessage() {}

func (x *GetPartyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[4]
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
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{4}
}

func (x *GetPartyRequest) GetPartyId() string {
	if x != nil {
		return x.PartyId
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
		mi := &file_packages_grpc_party_party_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByUserRequest) ProtoMessage() {}

func (x *GetByUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[5]
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
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{5}
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

type PagedParties struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Parties  []*Party `protobuf:"bytes,1,rep,name=parties,proto3" json:"parties,omitempty"`
	NextPage string   `protobuf:"bytes,2,opt,name=next_page,json=nextPage,proto3" json:"next_page,omitempty"`
}

func (x *PagedParties) Reset() {
	*x = PagedParties{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_party_party_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PagedParties) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PagedParties) ProtoMessage() {}

func (x *PagedParties) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[6]
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
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{6}
}

func (x *PagedParties) GetParties() []*Party {
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

type Party struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId        string  `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Title         string  `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	IsPublic      bool    `protobuf:"varint,4,opt,name=is_public,json=isPublic,proto3" json:"is_public,omitempty"`
	Lat           float32 `protobuf:"fixed32,5,opt,name=lat,proto3" json:"lat,omitempty"`
	Long          float32 `protobuf:"fixed32,6,opt,name=long,proto3" json:"long,omitempty"`
	StreetAddress string  `protobuf:"bytes,7,opt,name=street_address,json=streetAddress,proto3" json:"street_address,omitempty"`
	PostalCode    string  `protobuf:"bytes,8,opt,name=postal_code,json=postalCode,proto3" json:"postal_code,omitempty"`
	State         string  `protobuf:"bytes,9,opt,name=state,proto3" json:"state,omitempty"`
	Country       string  `protobuf:"bytes,10,opt,name=country,proto3" json:"country,omitempty"`
	StartDate     string  `protobuf:"bytes,11,opt,name=start_date,json=startDate,proto3" json:"start_date,omitempty"`
	CreatedAt     string  `protobuf:"bytes,12,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Party) Reset() {
	*x = Party{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_party_party_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Party) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Party) ProtoMessage() {}

func (x *Party) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_party_party_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Party.ProtoReflect.Descriptor instead.
func (*Party) Descriptor() ([]byte, []int) {
	return file_packages_grpc_party_party_proto_rawDescGZIP(), []int{7}
}

func (x *Party) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Party) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Party) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Party) GetIsPublic() bool {
	if x != nil {
		return x.IsPublic
	}
	return false
}

func (x *Party) GetLat() float32 {
	if x != nil {
		return x.Lat
	}
	return 0
}

func (x *Party) GetLong() float32 {
	if x != nil {
		return x.Long
	}
	return 0
}

func (x *Party) GetStreetAddress() string {
	if x != nil {
		return x.StreetAddress
	}
	return ""
}

func (x *Party) GetPostalCode() string {
	if x != nil {
		return x.PostalCode
	}
	return ""
}

func (x *Party) GetState() string {
	if x != nil {
		return x.State
	}
	return ""
}

func (x *Party) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Party) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *Party) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_packages_grpc_party_party_proto protoreflect.FileDescriptor

var file_packages_grpc_party_party_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x61, 0x72, 0x74, 0x79, 0x2f, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x70, 0x61, 0x72, 0x74, 0x79, 0x1a, 0x21, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67,
	0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa7, 0x02, 0x0a, 0x12,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6c,
	0x61, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e,
	0x67, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12, 0x25,
	0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74,
	0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18,
	0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f,
	0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0xa5, 0x02, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c,
	0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74,
	0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a,
	0x0b, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73,
	0x74, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18,
	0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d,
	0x0a, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x22, 0x52, 0x0a,
	0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x74, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x79, 0x49,
	0x64, 0x22, 0x73, 0x0a, 0x10, 0x47, 0x65, 0x6f, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x70, 0x72, 0x65, 0x63, 0x69, 0x73,
	0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02,
	0x52, 0x03, 0x6c, 0x61, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x02, 0x52, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78,
	0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65,
	0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x49, 0x64, 0x22, 0x5e, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6c,
	0x69, 0x6d, 0x69, 0x74, 0x22, 0x53, 0x0a, 0x0c, 0x50, 0x61, 0x67, 0x65, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x69, 0x65, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x65, 0x78, 0x74, 0x50, 0x61, 0x67, 0x65, 0x22, 0xbf, 0x02, 0x0a, 0x05, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x12,
	0x10, 0x0a, 0x03, 0x6c, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x02, 0x52, 0x03, 0x6c, 0x61,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x18, 0x06, 0x20, 0x01, 0x28, 0x02, 0x52,
	0x04, 0x6c, 0x6f, 0x6e, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x72, 0x65, 0x65, 0x74, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73,
	0x74, 0x72, 0x65, 0x65, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1f, 0x0a, 0x0b,
	0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x70, 0x6f, 0x73, 0x74, 0x61, 0x6c, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x0a,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1d, 0x0a,
	0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xf5, 0x02, 0x0a, 0x0c,
	0x50, 0x61, 0x72, 0x74, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x38, 0x0a, 0x0b,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x61,
	0x72, 0x74, 0x79, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50,
	0x61, 0x72, 0x74, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x08, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x12, 0x16, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x61, 0x72,
	0x74, 0x79, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x79, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x0b, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x72,
	0x74, 0x79, 0x22, 0x00, 0x12, 0x43, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x79, 0x12, 0x19, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x79, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x64, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x65, 0x73, 0x22, 0x00, 0x12, 0x3b, 0x0a, 0x09, 0x47, 0x65, 0x6f, 0x53, 0x65, 0x61,
	0x72, 0x63, 0x68, 0x12, 0x17, 0x2e, 0x70, 0x61, 0x72, 0x74, 0x79, 0x2e, 0x47, 0x65, 0x6f, 0x53,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70,
	0x61, 0x72, 0x74, 0x79, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x64, 0x50, 0x61, 0x72, 0x74, 0x69, 0x65,
	0x73, 0x22, 0x00, 0x42, 0x3c, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x6a, 0x6f, 0x6e, 0x61, 0x73, 0x68, 0x69, 0x6c, 0x74, 0x6c, 0x2f, 0x73, 0x65, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x61,
	0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x61, 0x72, 0x74,
	0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packages_grpc_party_party_proto_rawDescOnce sync.Once
	file_packages_grpc_party_party_proto_rawDescData = file_packages_grpc_party_party_proto_rawDesc
)

func file_packages_grpc_party_party_proto_rawDescGZIP() []byte {
	file_packages_grpc_party_party_proto_rawDescOnce.Do(func() {
		file_packages_grpc_party_party_proto_rawDescData = protoimpl.X.CompressGZIP(file_packages_grpc_party_party_proto_rawDescData)
	})
	return file_packages_grpc_party_party_proto_rawDescData
}

var file_packages_grpc_party_party_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_packages_grpc_party_party_proto_goTypes = []interface{}{
	(*CreatePartyRequest)(nil),     // 0: party.CreatePartyRequest
	(*UpdatePartyRequest)(nil),     // 1: party.UpdatePartyRequest
	(*DeletePartyRequest)(nil),     // 2: party.DeletePartyRequest
	(*GeoSearchRequest)(nil),       // 3: party.GeoSearchRequest
	(*GetPartyRequest)(nil),        // 4: party.GetPartyRequest
	(*GetByUserRequest)(nil),       // 5: party.GetByUserRequest
	(*PagedParties)(nil),           // 6: party.PagedParties
	(*Party)(nil),                  // 7: party.Party
	(*common.MessageResponse)(nil), // 8: common.MessageResponse
}
var file_packages_grpc_party_party_proto_depIdxs = []int32{
	7, // 0: party.PagedParties.parties:type_name -> party.Party
	0, // 1: party.PartyService.CreateParty:input_type -> party.CreatePartyRequest
	4, // 2: party.PartyService.GetParty:input_type -> party.GetPartyRequest
	1, // 3: party.PartyService.UpdateParty:input_type -> party.UpdatePartyRequest
	2, // 4: party.PartyService.DeleteParty:input_type -> party.DeletePartyRequest
	5, // 5: party.PartyService.GetByUser:input_type -> party.GetByUserRequest
	3, // 6: party.PartyService.GeoSearch:input_type -> party.GeoSearchRequest
	7, // 7: party.PartyService.CreateParty:output_type -> party.Party
	7, // 8: party.PartyService.GetParty:output_type -> party.Party
	7, // 9: party.PartyService.UpdateParty:output_type -> party.Party
	8, // 10: party.PartyService.DeleteParty:output_type -> common.MessageResponse
	6, // 11: party.PartyService.GetByUser:output_type -> party.PagedParties
	6, // 12: party.PartyService.GeoSearch:output_type -> party.PagedParties
	7, // [7:13] is the sub-list for method output_type
	1, // [1:7] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_packages_grpc_party_party_proto_init() }
func file_packages_grpc_party_party_proto_init() {
	if File_packages_grpc_party_party_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packages_grpc_party_party_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_packages_grpc_party_party_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_packages_grpc_party_party_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeletePartyRequest); i {
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
		file_packages_grpc_party_party_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_packages_grpc_party_party_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
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
		file_packages_grpc_party_party_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_packages_grpc_party_party_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_packages_grpc_party_party_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Party); i {
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
			RawDescriptor: file_packages_grpc_party_party_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_packages_grpc_party_party_proto_goTypes,
		DependencyIndexes: file_packages_grpc_party_party_proto_depIdxs,
		MessageInfos:      file_packages_grpc_party_party_proto_msgTypes,
	}.Build()
	File_packages_grpc_party_party_proto = out.File
	file_packages_grpc_party_party_proto_rawDesc = nil
	file_packages_grpc_party_party_proto_goTypes = nil
	file_packages_grpc_party_party_proto_depIdxs = nil
}
