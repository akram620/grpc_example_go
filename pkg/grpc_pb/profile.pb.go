// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.12.4
// source: pkg/grpc_proto/profile.proto

package grpc_pb

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

type ContentTypes int32

const (
	ContentTypes_image ContentTypes = 0
	ContentTypes_video ContentTypes = 1
)

// Enum value maps for ContentTypes.
var (
	ContentTypes_name = map[int32]string{
		0: "image",
		1: "video",
	}
	ContentTypes_value = map[string]int32{
		"image": 0,
		"video": 1,
	}
)

func (x ContentTypes) Enum() *ContentTypes {
	p := new(ContentTypes)
	*p = x
	return p
}

func (x ContentTypes) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ContentTypes) Descriptor() protoreflect.EnumDescriptor {
	return file_pkg_grpc_proto_profile_proto_enumTypes[0].Descriptor()
}

func (ContentTypes) Type() protoreflect.EnumType {
	return &file_pkg_grpc_proto_profile_proto_enumTypes[0]
}

func (x ContentTypes) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ContentTypes.Descriptor instead.
func (ContentTypes) EnumDescriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_profile_proto_rawDescGZIP(), []int{0}
}

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId int64 `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_profile_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_profile_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_profile_proto_rawDescGZIP(), []int{0}
}

func (x *UserId) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age  int32  `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_profile_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_profile_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_profile_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

type Content struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ContentType ContentTypes `protobuf:"varint,2,opt,name=contentType,proto3,enum=service.ContentTypes" json:"contentType,omitempty"`
	Url         string       `protobuf:"bytes,3,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Content) Reset() {
	*x = Content{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_profile_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Content) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Content) ProtoMessage() {}

func (x *Content) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_profile_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Content.ProtoReflect.Descriptor instead.
func (*Content) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_profile_proto_rawDescGZIP(), []int{2}
}

func (x *Content) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Content) GetContentType() ContentTypes {
	if x != nil {
		return x.ContentType
	}
	return ContentTypes_image
}

func (x *Content) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Contents struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int64      `protobuf:"varint,1,opt,name=userId,proto3" json:"userId,omitempty"`
	Contents []*Content `protobuf:"bytes,2,rep,name=contents,proto3" json:"contents,omitempty"`
}

func (x *Contents) Reset() {
	*x = Contents{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_grpc_proto_profile_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contents) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contents) ProtoMessage() {}

func (x *Contents) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_grpc_proto_profile_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contents.ProtoReflect.Descriptor instead.
func (*Contents) Descriptor() ([]byte, []int) {
	return file_pkg_grpc_proto_profile_proto_rawDescGZIP(), []int{3}
}

func (x *Contents) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Contents) GetContents() []*Content {
	if x != nil {
		return x.Contents
	}
	return nil
}

var File_pkg_grpc_proto_profile_proto protoreflect.FileDescriptor

var file_pkg_grpc_proto_profile_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x22, 0x20, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x64, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x40, 0x0a, 0x08, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x22, 0x64, 0x0a, 0x07, 0x43,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x37, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70,
	0x65, 0x73, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72,
	0x6c, 0x22, 0x50, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x16, 0x0a,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75,
	0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x73, 0x2a, 0x24, 0x0a, 0x0c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x73, 0x12, 0x09, 0x0a, 0x05, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x10, 0x01, 0x32, 0x7a, 0x0a, 0x07, 0x50, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x12, 0x35, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x42, 0x79, 0x49, 0x64, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x38, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x0f, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x1a, 0x11, 0x2e, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x42, 0x0d, 0x5a, 0x0b, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70,
	0x63, 0x5f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_grpc_proto_profile_proto_rawDescOnce sync.Once
	file_pkg_grpc_proto_profile_proto_rawDescData = file_pkg_grpc_proto_profile_proto_rawDesc
)

func file_pkg_grpc_proto_profile_proto_rawDescGZIP() []byte {
	file_pkg_grpc_proto_profile_proto_rawDescOnce.Do(func() {
		file_pkg_grpc_proto_profile_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_grpc_proto_profile_proto_rawDescData)
	})
	return file_pkg_grpc_proto_profile_proto_rawDescData
}

var file_pkg_grpc_proto_profile_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_pkg_grpc_proto_profile_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_pkg_grpc_proto_profile_proto_goTypes = []interface{}{
	(ContentTypes)(0), // 0: service.ContentTypes
	(*UserId)(nil),    // 1: service.UserId
	(*UserInfo)(nil),  // 2: service.UserInfo
	(*Content)(nil),   // 3: service.Content
	(*Contents)(nil),  // 4: service.Contents
}
var file_pkg_grpc_proto_profile_proto_depIdxs = []int32{
	0, // 0: service.Content.contentType:type_name -> service.ContentTypes
	3, // 1: service.Contents.contents:type_name -> service.Content
	1, // 2: service.Profile.GetUserInfoById:input_type -> service.UserId
	1, // 3: service.Profile.GetUserContentById:input_type -> service.UserId
	2, // 4: service.Profile.GetUserInfoById:output_type -> service.UserInfo
	4, // 5: service.Profile.GetUserContentById:output_type -> service.Contents
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pkg_grpc_proto_profile_proto_init() }
func file_pkg_grpc_proto_profile_proto_init() {
	if File_pkg_grpc_proto_profile_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_grpc_proto_profile_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_pkg_grpc_proto_profile_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_pkg_grpc_proto_profile_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Content); i {
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
		file_pkg_grpc_proto_profile_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contents); i {
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
			RawDescriptor: file_pkg_grpc_proto_profile_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_grpc_proto_profile_proto_goTypes,
		DependencyIndexes: file_pkg_grpc_proto_profile_proto_depIdxs,
		EnumInfos:         file_pkg_grpc_proto_profile_proto_enumTypes,
		MessageInfos:      file_pkg_grpc_proto_profile_proto_msgTypes,
	}.Build()
	File_pkg_grpc_proto_profile_proto = out.File
	file_pkg_grpc_proto_profile_proto_rawDesc = nil
	file_pkg_grpc_proto_profile_proto_goTypes = nil
	file_pkg_grpc_proto_profile_proto_depIdxs = nil
}
