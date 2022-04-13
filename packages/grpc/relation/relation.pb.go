// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.6.1
// source: packages/grpc/relation/relation.proto

package relation

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

type FriendRelation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId   string `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
	Accepted   bool   `protobuf:"varint,3,opt,name=accepted,proto3" json:"accepted,omitempty"`
	CreatedAt  string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	AcceptedAt string `protobuf:"bytes,5,opt,name=accepted_at,json=acceptedAt,proto3" json:"accepted_at,omitempty"`
}

func (x *FriendRelation) Reset() {
	*x = FriendRelation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_relation_relation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRelation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRelation) ProtoMessage() {}

func (x *FriendRelation) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_relation_relation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRelation.ProtoReflect.Descriptor instead.
func (*FriendRelation) Descriptor() ([]byte, []int) {
	return file_packages_grpc_relation_relation_proto_rawDescGZIP(), []int{0}
}

func (x *FriendRelation) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FriendRelation) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

func (x *FriendRelation) GetAccepted() bool {
	if x != nil {
		return x.Accepted
	}
	return false
}

func (x *FriendRelation) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *FriendRelation) GetAcceptedAt() string {
	if x != nil {
		return x.AcceptedAt
	}
	return ""
}

type FriendRequestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId string `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *FriendRequestRequest) Reset() {
	*x = FriendRequestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_relation_relation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendRequestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendRequestRequest) ProtoMessage() {}

func (x *FriendRequestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_relation_relation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendRequestRequest.ProtoReflect.Descriptor instead.
func (*FriendRequestRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_relation_relation_proto_rawDescGZIP(), []int{1}
}

func (x *FriendRequestRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *FriendRequestRequest) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

type AcceptFriendRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	FriendId string `protobuf:"bytes,2,opt,name=friend_id,json=friendId,proto3" json:"friend_id,omitempty"`
}

func (x *AcceptFriendRequest) Reset() {
	*x = AcceptFriendRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_packages_grpc_relation_relation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AcceptFriendRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AcceptFriendRequest) ProtoMessage() {}

func (x *AcceptFriendRequest) ProtoReflect() protoreflect.Message {
	mi := &file_packages_grpc_relation_relation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AcceptFriendRequest.ProtoReflect.Descriptor instead.
func (*AcceptFriendRequest) Descriptor() ([]byte, []int) {
	return file_packages_grpc_relation_relation_proto_rawDescGZIP(), []int{2}
}

func (x *AcceptFriendRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AcceptFriendRequest) GetFriendId() string {
	if x != nil {
		return x.FriendId
	}
	return ""
}

var File_packages_grpc_relation_relation_proto protoreflect.FileDescriptor

var file_packages_grpc_relation_relation_proto_rawDesc = []byte{
	0x0a, 0x25, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0xa2, 0x01, 0x0a, 0x0e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a,
	0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x65,
	0x70, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x4c, 0x0a, 0x14, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x13, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x49,
	0x64, 0x32, 0xa9, 0x01, 0x0a, 0x0f, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0d, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x00, 0x12, 0x49, 0x0a, 0x0c, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x12, 0x1d, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x70, 0x74, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x18, 0x2e, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x46, 0x72, 0x69,
	0x65, 0x6e, 0x64, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x00, 0x42, 0x3f, 0x5a,
	0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x6f, 0x6e, 0x61,
	0x73, 0x68, 0x69, 0x6c, 0x74, 0x6c, 0x2f, 0x73, 0x65, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x2d,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x73,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_packages_grpc_relation_relation_proto_rawDescOnce sync.Once
	file_packages_grpc_relation_relation_proto_rawDescData = file_packages_grpc_relation_relation_proto_rawDesc
)

func file_packages_grpc_relation_relation_proto_rawDescGZIP() []byte {
	file_packages_grpc_relation_relation_proto_rawDescOnce.Do(func() {
		file_packages_grpc_relation_relation_proto_rawDescData = protoimpl.X.CompressGZIP(file_packages_grpc_relation_relation_proto_rawDescData)
	})
	return file_packages_grpc_relation_relation_proto_rawDescData
}

var file_packages_grpc_relation_relation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_packages_grpc_relation_relation_proto_goTypes = []interface{}{
	(*FriendRelation)(nil),       // 0: relation.FriendRelation
	(*FriendRequestRequest)(nil), // 1: relation.FriendRequestRequest
	(*AcceptFriendRequest)(nil),  // 2: relation.AcceptFriendRequest
}
var file_packages_grpc_relation_relation_proto_depIdxs = []int32{
	1, // 0: relation.RelationService.FriendRequest:input_type -> relation.FriendRequestRequest
	2, // 1: relation.RelationService.AcceptFriend:input_type -> relation.AcceptFriendRequest
	0, // 2: relation.RelationService.FriendRequest:output_type -> relation.FriendRelation
	0, // 3: relation.RelationService.AcceptFriend:output_type -> relation.FriendRelation
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_packages_grpc_relation_relation_proto_init() }
func file_packages_grpc_relation_relation_proto_init() {
	if File_packages_grpc_relation_relation_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_packages_grpc_relation_relation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRelation); i {
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
		file_packages_grpc_relation_relation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendRequestRequest); i {
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
		file_packages_grpc_relation_relation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AcceptFriendRequest); i {
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
			RawDescriptor: file_packages_grpc_relation_relation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_packages_grpc_relation_relation_proto_goTypes,
		DependencyIndexes: file_packages_grpc_relation_relation_proto_depIdxs,
		MessageInfos:      file_packages_grpc_relation_relation_proto_msgTypes,
	}.Build()
	File_packages_grpc_relation_relation_proto = out.File
	file_packages_grpc_relation_relation_proto_rawDesc = nil
	file_packages_grpc_relation_relation_proto_goTypes = nil
	file_packages_grpc_relation_relation_proto_depIdxs = nil
}
