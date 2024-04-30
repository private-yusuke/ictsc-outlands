// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.0
// 	protoc        (unknown)
// source: contestant/v1/rule.proto

package v1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
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

type Rule struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule           string `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
	ShortRule      string `protobuf:"bytes,2,opt,name=short_rule,json=shortRule,proto3" json:"short_rule,omitempty"`
	RecreationRule string `protobuf:"bytes,3,opt,name=recreation_rule,json=recreationRule,proto3" json:"recreation_rule,omitempty"`
}

func (x *Rule) Reset() {
	*x = Rule{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_rule_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Rule) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Rule) ProtoMessage() {}

func (x *Rule) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_rule_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Rule.ProtoReflect.Descriptor instead.
func (*Rule) Descriptor() ([]byte, []int) {
	return file_contestant_v1_rule_proto_rawDescGZIP(), []int{0}
}

func (x *Rule) GetRule() string {
	if x != nil {
		return x.Rule
	}
	return ""
}

func (x *Rule) GetShortRule() string {
	if x != nil {
		return x.ShortRule
	}
	return ""
}

func (x *Rule) GetRecreationRule() string {
	if x != nil {
		return x.RecreationRule
	}
	return ""
}

type GetRuleRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetRuleRequest) Reset() {
	*x = GetRuleRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_rule_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleRequest) ProtoMessage() {}

func (x *GetRuleRequest) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_rule_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleRequest.ProtoReflect.Descriptor instead.
func (*GetRuleRequest) Descriptor() ([]byte, []int) {
	return file_contestant_v1_rule_proto_rawDescGZIP(), []int{1}
}

type GetRuleResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rule *Rule `protobuf:"bytes,1,opt,name=rule,proto3" json:"rule,omitempty"`
}

func (x *GetRuleResponse) Reset() {
	*x = GetRuleResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_contestant_v1_rule_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRuleResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRuleResponse) ProtoMessage() {}

func (x *GetRuleResponse) ProtoReflect() protoreflect.Message {
	mi := &file_contestant_v1_rule_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRuleResponse.ProtoReflect.Descriptor instead.
func (*GetRuleResponse) Descriptor() ([]byte, []int) {
	return file_contestant_v1_rule_proto_rawDescGZIP(), []int{2}
}

func (x *GetRuleResponse) GetRule() *Rule {
	if x != nil {
		return x.Rule
	}
	return nil
}

var File_contestant_v1_rule_proto protoreflect.FileDescriptor

var file_contestant_v1_rule_proto_rawDesc = []byte{
	0x0a, 0x18, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x2f,
	0x72, 0x75, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x63, 0x6f, 0x6e, 0x74,
	0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x86, 0x01, 0x0a, 0x04, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x1e, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0a, 0xba,
	0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x12,
	0x29, 0x0a, 0x0a, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xc8, 0x01, 0x52,
	0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12, 0x33, 0x0a, 0x0f, 0x72, 0x65,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x0a, 0xba, 0x48, 0x07, 0x72, 0x05, 0x10, 0x01, 0x18, 0xe8, 0x07, 0x52,
	0x0e, 0x72, 0x65, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x75, 0x6c, 0x65, 0x22,
	0x10, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x42, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x04, 0x72, 0x75, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x13, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x75, 0x6c, 0x65, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52,
	0x04, 0x72, 0x75, 0x6c, 0x65, 0x32, 0x57, 0x0a, 0x0b, 0x52, 0x75, 0x6c, 0x65, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x12,
	0x1d, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e,
	0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74, 0x61, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x46,
	0x5a, 0x44, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x63, 0x74,
	0x73, 0x63, 0x2f, 0x69, 0x63, 0x74, 0x73, 0x63, 0x2d, 0x6f, 0x75, 0x74, 0x6c, 0x61, 0x6e, 0x64,
	0x73, 0x2f, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x73, 0x74,
	0x61, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_contestant_v1_rule_proto_rawDescOnce sync.Once
	file_contestant_v1_rule_proto_rawDescData = file_contestant_v1_rule_proto_rawDesc
)

func file_contestant_v1_rule_proto_rawDescGZIP() []byte {
	file_contestant_v1_rule_proto_rawDescOnce.Do(func() {
		file_contestant_v1_rule_proto_rawDescData = protoimpl.X.CompressGZIP(file_contestant_v1_rule_proto_rawDescData)
	})
	return file_contestant_v1_rule_proto_rawDescData
}

var file_contestant_v1_rule_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_contestant_v1_rule_proto_goTypes = []interface{}{
	(*Rule)(nil),            // 0: contestant.v1.Rule
	(*GetRuleRequest)(nil),  // 1: contestant.v1.GetRuleRequest
	(*GetRuleResponse)(nil), // 2: contestant.v1.GetRuleResponse
}
var file_contestant_v1_rule_proto_depIdxs = []int32{
	0, // 0: contestant.v1.GetRuleResponse.rule:type_name -> contestant.v1.Rule
	1, // 1: contestant.v1.RuleService.GetRule:input_type -> contestant.v1.GetRuleRequest
	2, // 2: contestant.v1.RuleService.GetRule:output_type -> contestant.v1.GetRuleResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_contestant_v1_rule_proto_init() }
func file_contestant_v1_rule_proto_init() {
	if File_contestant_v1_rule_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_contestant_v1_rule_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Rule); i {
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
		file_contestant_v1_rule_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleRequest); i {
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
		file_contestant_v1_rule_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRuleResponse); i {
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
			RawDescriptor: file_contestant_v1_rule_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_contestant_v1_rule_proto_goTypes,
		DependencyIndexes: file_contestant_v1_rule_proto_depIdxs,
		MessageInfos:      file_contestant_v1_rule_proto_msgTypes,
	}.Build()
	File_contestant_v1_rule_proto = out.File
	file_contestant_v1_rule_proto_rawDesc = nil
	file_contestant_v1_rule_proto_goTypes = nil
	file_contestant_v1_rule_proto_depIdxs = nil
}
