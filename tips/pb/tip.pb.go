// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.6
// source: pb/tip.proto

package pb

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

type Tip struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Amount float64 `protobuf:"fixed64,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Note   string  `protobuf:"bytes,3,opt,name=note,proto3" json:"note,omitempty"`
	UserId string  `protobuf:"bytes,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Tip) Reset() {
	*x = Tip{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Tip) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Tip) ProtoMessage() {}

func (x *Tip) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Tip.ProtoReflect.Descriptor instead.
func (*Tip) Descriptor() ([]byte, []int) {
	return file_pb_tip_proto_rawDescGZIP(), []int{0}
}

func (x *Tip) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Tip) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *Tip) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Tip) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type FindTipsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FindTipsReq) Reset() {
	*x = FindTipsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tip_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTipsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTipsReq) ProtoMessage() {}

func (x *FindTipsReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tip_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTipsReq.ProtoReflect.Descriptor instead.
func (*FindTipsReq) Descriptor() ([]byte, []int) {
	return file_pb_tip_proto_rawDescGZIP(), []int{1}
}

type FindTipsRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tip *Tip `protobuf:"bytes,1,opt,name=tip,proto3" json:"tip,omitempty"`
}

func (x *FindTipsRes) Reset() {
	*x = FindTipsRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tip_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindTipsRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindTipsRes) ProtoMessage() {}

func (x *FindTipsRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tip_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindTipsRes.ProtoReflect.Descriptor instead.
func (*FindTipsRes) Descriptor() ([]byte, []int) {
	return file_pb_tip_proto_rawDescGZIP(), []int{2}
}

func (x *FindTipsRes) GetTip() *Tip {
	if x != nil {
		return x.Tip
	}
	return nil
}

type FindOneTipReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *FindOneTipReq) Reset() {
	*x = FindOneTipReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tip_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneTipReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneTipReq) ProtoMessage() {}

func (x *FindOneTipReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tip_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneTipReq.ProtoReflect.Descriptor instead.
func (*FindOneTipReq) Descriptor() ([]byte, []int) {
	return file_pb_tip_proto_rawDescGZIP(), []int{3}
}

func (x *FindOneTipReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type FindOneTipRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tip *Tip `protobuf:"bytes,1,opt,name=tip,proto3" json:"tip,omitempty"`
}

func (x *FindOneTipRes) Reset() {
	*x = FindOneTipRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tip_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FindOneTipRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FindOneTipRes) ProtoMessage() {}

func (x *FindOneTipRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tip_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FindOneTipRes.ProtoReflect.Descriptor instead.
func (*FindOneTipRes) Descriptor() ([]byte, []int) {
	return file_pb_tip_proto_rawDescGZIP(), []int{4}
}

func (x *FindOneTipRes) GetTip() *Tip {
	if x != nil {
		return x.Tip
	}
	return nil
}

type CreateTipReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tip *Tip `protobuf:"bytes,1,opt,name=tip,proto3" json:"tip,omitempty"`
}

func (x *CreateTipReq) Reset() {
	*x = CreateTipReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tip_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTipReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTipReq) ProtoMessage() {}

func (x *CreateTipReq) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tip_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTipReq.ProtoReflect.Descriptor instead.
func (*CreateTipReq) Descriptor() ([]byte, []int) {
	return file_pb_tip_proto_rawDescGZIP(), []int{5}
}

func (x *CreateTipReq) GetTip() *Tip {
	if x != nil {
		return x.Tip
	}
	return nil
}

type CreateTipRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tip *Tip `protobuf:"bytes,1,opt,name=tip,proto3" json:"tip,omitempty"`
}

func (x *CreateTipRes) Reset() {
	*x = CreateTipRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_tip_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateTipRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateTipRes) ProtoMessage() {}

func (x *CreateTipRes) ProtoReflect() protoreflect.Message {
	mi := &file_pb_tip_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateTipRes.ProtoReflect.Descriptor instead.
func (*CreateTipRes) Descriptor() ([]byte, []int) {
	return file_pb_tip_proto_rawDescGZIP(), []int{6}
}

func (x *CreateTipRes) GetTip() *Tip {
	if x != nil {
		return x.Tip
	}
	return nil
}

var File_pb_tip_proto protoreflect.FileDescriptor

var file_pb_tip_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x70, 0x62, 0x2f, 0x74, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x70, 0x62, 0x22, 0x5a, 0x0a, 0x03, 0x54, 0x69, 0x70, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x0d,
	0x0a, 0x0b, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x22, 0x28, 0x0a,
	0x0b, 0x46, 0x69, 0x6e, 0x64, 0x54, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x03,
	0x74, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x54,
	0x69, 0x70, 0x52, 0x03, 0x74, 0x69, 0x70, 0x22, 0x1f, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64, 0x4f,
	0x6e, 0x65, 0x54, 0x69, 0x70, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2a, 0x0a, 0x0d, 0x46, 0x69, 0x6e, 0x64,
	0x4f, 0x6e, 0x65, 0x54, 0x69, 0x70, 0x52, 0x65, 0x73, 0x12, 0x19, 0x0a, 0x03, 0x74, 0x69, 0x70,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x70, 0x52,
	0x03, 0x74, 0x69, 0x70, 0x22, 0x29, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69,
	0x70, 0x52, 0x65, 0x71, 0x12, 0x19, 0x0a, 0x03, 0x74, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x07, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x69, 0x70, 0x52, 0x03, 0x74, 0x69, 0x70, 0x22,
	0x29, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x52, 0x65, 0x73, 0x12,
	0x19, 0x0a, 0x03, 0x74, 0x69, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x70,
	0x62, 0x2e, 0x54, 0x69, 0x70, 0x52, 0x03, 0x74, 0x69, 0x70, 0x32, 0xa7, 0x01, 0x0a, 0x0a, 0x54,
	0x69, 0x70, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30, 0x0a, 0x08, 0x46, 0x69, 0x6e,
	0x64, 0x54, 0x69, 0x70, 0x73, 0x12, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x54,
	0x69, 0x70, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64,
	0x54, 0x69, 0x70, 0x73, 0x52, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x12, 0x34, 0x0a, 0x0a, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x70, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x46,
	0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x70,
	0x62, 0x2e, 0x46, 0x69, 0x6e, 0x64, 0x4f, 0x6e, 0x65, 0x54, 0x69, 0x70, 0x52, 0x65, 0x73, 0x22,
	0x00, 0x12, 0x31, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x12, 0x10,
	0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x52, 0x65, 0x71,
	0x1a, 0x10, 0x2e, 0x70, 0x62, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x70, 0x52,
	0x65, 0x73, 0x22, 0x00, 0x42, 0x1e, 0x5a, 0x1c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x65, 0x78, 0x38, 0x2f, 0x74, 0x69, 0x70, 0x6f, 0x6e, 0x2f, 0x74, 0x69, 0x70,
	0x73, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_tip_proto_rawDescOnce sync.Once
	file_pb_tip_proto_rawDescData = file_pb_tip_proto_rawDesc
)

func file_pb_tip_proto_rawDescGZIP() []byte {
	file_pb_tip_proto_rawDescOnce.Do(func() {
		file_pb_tip_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_tip_proto_rawDescData)
	})
	return file_pb_tip_proto_rawDescData
}

var file_pb_tip_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_pb_tip_proto_goTypes = []interface{}{
	(*Tip)(nil),           // 0: pb.Tip
	(*FindTipsReq)(nil),   // 1: pb.FindTipsReq
	(*FindTipsRes)(nil),   // 2: pb.FindTipsRes
	(*FindOneTipReq)(nil), // 3: pb.FindOneTipReq
	(*FindOneTipRes)(nil), // 4: pb.FindOneTipRes
	(*CreateTipReq)(nil),  // 5: pb.CreateTipReq
	(*CreateTipRes)(nil),  // 6: pb.CreateTipRes
}
var file_pb_tip_proto_depIdxs = []int32{
	0, // 0: pb.FindTipsRes.tip:type_name -> pb.Tip
	0, // 1: pb.FindOneTipRes.tip:type_name -> pb.Tip
	0, // 2: pb.CreateTipReq.tip:type_name -> pb.Tip
	0, // 3: pb.CreateTipRes.tip:type_name -> pb.Tip
	1, // 4: pb.TipService.FindTips:input_type -> pb.FindTipsReq
	3, // 5: pb.TipService.FindOneTip:input_type -> pb.FindOneTipReq
	5, // 6: pb.TipService.CreateTip:input_type -> pb.CreateTipReq
	2, // 7: pb.TipService.FindTips:output_type -> pb.FindTipsRes
	4, // 8: pb.TipService.FindOneTip:output_type -> pb.FindOneTipRes
	6, // 9: pb.TipService.CreateTip:output_type -> pb.CreateTipRes
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_pb_tip_proto_init() }
func file_pb_tip_proto_init() {
	if File_pb_tip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_tip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Tip); i {
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
		file_pb_tip_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTipsReq); i {
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
		file_pb_tip_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindTipsRes); i {
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
		file_pb_tip_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneTipReq); i {
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
		file_pb_tip_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FindOneTipRes); i {
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
		file_pb_tip_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTipReq); i {
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
		file_pb_tip_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateTipRes); i {
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
			RawDescriptor: file_pb_tip_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pb_tip_proto_goTypes,
		DependencyIndexes: file_pb_tip_proto_depIdxs,
		MessageInfos:      file_pb_tip_proto_msgTypes,
	}.Build()
	File_pb_tip_proto = out.File
	file_pb_tip_proto_rawDesc = nil
	file_pb_tip_proto_goTypes = nil
	file_pb_tip_proto_depIdxs = nil
}
