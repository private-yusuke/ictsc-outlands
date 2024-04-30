// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: contestant/v1/announcement.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Announcement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ProblemId *string                `protobuf:"bytes,2,opt,name=problem_id,json=problemId,proto3,oneof" json:"problem_id,omitempty"`
	Title     string                 `protobuf:"bytes,3,opt,name=title,proto3" json:"title,omitempty"`
	Body      string                 `protobuf:"bytes,4,opt,name=body,proto3" json:"body,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Announcement) Reset() {
	*x = Announcement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_announcement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Announcement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Announcement) ProtoMessage() {}

func (x *Announcement) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_announcement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Announcement.ProtoReflect.Descriptor instead.
func (*Announcement) Descriptor() ([]byte, []int) {
	return file_contestant_v1_announcement_proto_rawDescGZIP(), []int{0}
}

func (x *Announcement) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Announcement) GetProblemId() string {
	if x != nil && x.ProblemId != nil {
		return *x.ProblemId
	}
	return ""
}

func (x *Announcement) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Announcement) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

func (x *Announcement) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type GetAnnouncementsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProblemId *string `protobuf:"bytes,1,opt,name=problem_id,json=problemId,proto3,oneof" json:"problem_id,omitempty"`
}

func (x *GetAnnouncementsRequest) Reset() {
	*x = GetAnnouncementsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_announcement_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnouncementsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnouncementsRequest) ProtoMessage() {}

func (x *GetAnnouncementsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_announcement_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnouncementsRequest.ProtoReflect.Descriptor instead.
func (*GetAnnouncementsRequest) Descriptor() ([]byte, []int) {
	return file_contestant_v1_announcement_proto_rawDescGZIP(), []int{1}
}

func (x *GetAnnouncementsRequest) GetProblemId() string {
	if x != nil && x.ProblemId != nil {
		return *x.ProblemId
	}
	return ""
}

type GetAnnouncementsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Announcements []*Announcement `protobuf:"bytes,1,rep,name=announcements,proto3" json:"announcements,omitempty"`
}

func (x *GetAnnouncementsResponse) Reset() {
	*x = GetAnnouncementsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_announcement_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAnnouncementsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAnnouncementsResponse) ProtoMessage() {}

func (x *GetAnnouncementsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_announcement_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAnnouncementsResponse.ProtoReflect.Descriptor instead.
func (*GetAnnouncementsResponse) Descriptor() ([]byte, []int) {
	return file_contestant_v1_announcement_proto_rawDescGZIP(), []int{2}
}

func (x *GetAnnouncementsResponse) GetAnnouncements() []*Announcement {
	if x != nil {
		return x.Announcements
	}
	return nil
}

var File_contestant_v1_announcement_proto protoreflect.FileDescriptor

var file_contestant_v1_announcement_proto_rawDesc = []byte{
	0x0a, 0x20, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76,
	0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f,
	0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xe9, 0x01, 0x0a, 0x0c, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x12, 0x18, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48,
	0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08,
	0xba, 0x48, 0x05, 0x72, 0x03, 0x98, 0x01, 0x1a, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62,
	0x6c, 0x65, 0x6d, 0x49, 0x64, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba, 0x48, 0x06, 0x72, 0x04, 0x10, 0x01,
	0x18, 0x64, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01,
	0x18, 0xe8, 0x07, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x12, 0x41, 0x0a, 0x0a, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01,
	0x01, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0x0d, 0x0a, 0x0b,
	0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x5f, 0x69, 0x64, 0x22, 0x56, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65,
	0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x08, 0xba, 0x48, 0x05, 0x72,
	0x03, 0x98, 0x01, 0x1a, 0x48, 0x00, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d, 0x49,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x70, 0x72, 0x6f, 0x62, 0x6c, 0x65, 0x6d,
	0x5f, 0x69, 0x64, 0x22, 0x65, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x49, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0d, 0x61, 0x6e, 0x6e,
	0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x32, 0x7a, 0x0a, 0x13, 0x41, 0x6e,
	0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x12, 0x63, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65,
	0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x26, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63,
	0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x27, 0x2e,
	0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46, 0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63,
	0x2d, 0x6f, 0x75, 0x74, 0x6c, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contestant_v1_announcement_proto_rawDescOnce sync.Once
	file_contestant_v1_announcement_proto_rawDescData = file_contestant_v1_announcement_proto_rawDesc
)

func file_contestant_v1_announcement_proto_rawDescGZIP() []byte {
	file_contestant_v1_announcement_proto_rawDescOnce.Do(func() {
		file_contestant_v1_announcement_proto_rawDescData = protoimpl.X.CompressGZIP(file_contestant_v1_announcement_proto_rawDescData)
	})
	return file_contestant_v1_announcement_proto_rawDescData
}

var file_contestant_v1_announcement_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contestant_v1_announcement_proto_goTypes = []interface{}{
	(*Announcement)(nil),             // 0: contestant.v1.Announcement
	(*GetAnnouncementsRequest)(nil),  // 1: contestant.v1.GetAnnouncementsRequest
	(*GetAnnouncementsResponse)(nil), // 2: contestant.v1.GetAnnouncementsResponse
	(*timestamppb.Timestamp)(nil),    // 3: google.protobuf.Timestamp
}
var file_contestant_v1_announcement_proto_depIdxs = []int32{
	3, // 0: contestant.v1.Announcement.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: contestant.v1.GetAnnouncementsResponse.announcements:type_name -> contestant.v1.Announcement
	1, // 2: contestant.v1.AnnouncementService.GetAnnouncements:input_type -> contestant.v1.GetAnnouncementsRequest
	2, // 3: contestant.v1.AnnouncementService.GetAnnouncements:output_type -> contestant.v1.GetAnnouncementsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_contestant_v1_announcement_proto_init() }
func file_contestant_v1_announcement_proto_init() {
	if File_contestant_v1_announcement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contestant_v1_announcement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Announcement); i {
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
		file_contestant_v1_announcement_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnouncementsRequest); i {
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
		file_contestant_v1_announcement_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAnnouncementsResponse); i {
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
	file_contestant_v1_announcement_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_contestant_v1_announcement_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_contestant_v1_announcement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contestant_v1_announcement_proto_goTypes,
		DependencyIndexes: file_contestant_v1_announcement_proto_depIdxs,
		MessageInfos:      file_contestant_v1_announcement_proto_msgTypes,
	}.Build()
	File_contestant_v1_announcement_proto = out.File
	file_contestant_v1_announcement_proto_rawDesc = nil
	file_contestant_v1_announcement_proto_goTypes = nil
	file_contestant_v1_announcement_proto_depIdxs = nil
}
