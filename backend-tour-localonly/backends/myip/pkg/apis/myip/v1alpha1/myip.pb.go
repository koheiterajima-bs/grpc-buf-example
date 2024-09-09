// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0-devel
// 	protoc        (unknown)
// source: myip/v1alpha1/myip.proto

package myipv1alpha1

import (
	_ "google.golang.org/genproto/googleapis/rpc/status"
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

// IP の種別
type IPType int32

const (
	IPType_IP_TYPE_UNSPECIFIED IPType = 0
	IPType_IP_TYPE_IPV4        IPType = 1
	IPType_IP_TYPE_UNIVERSAL   IPType = 2
)

// Enum value maps for IPType.
var (
	IPType_name = map[int32]string{
		0: "IP_TYPE_UNSPECIFIED",
		1: "IP_TYPE_IPV4",
		2: "IP_TYPE_UNIVERSAL",
	}
	IPType_value = map[string]int32{
		"IP_TYPE_UNSPECIFIED": 0,
		"IP_TYPE_IPV4":        1,
		"IP_TYPE_UNIVERSAL":   2,
	}
)

func (x IPType) Enum() *IPType {
	p := new(IPType)
	*p = x
	return p
}

func (x IPType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (IPType) Descriptor() protoreflect.EnumDescriptor {
	return file_myip_v1alpha1_myip_proto_enumTypes[0].Descriptor()
}

func (IPType) Type() protoreflect.EnumType {
	return &file_myip_v1alpha1_myip_proto_enumTypes[0]
}

func (x IPType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use IPType.Descriptor instead.
func (IPType) EnumDescriptor() ([]byte, []int) {
	return file_myip_v1alpha1_myip_proto_rawDescGZIP(), []int{0}
}

// メッセージの状態
// エラーを4つに集約
type MyIP_State int32

const (
	MyIP_STATE_UNSPECIFIED MyIP_State = 0
	// 処理が受け付けられた
	MyIP_STATE_ACCEPTED MyIP_State = 1
	// 送信エラー
	MyIP_STATE_FAILED MyIP_State = 2
)

// Enum value maps for MyIP_State.
var (
	MyIP_State_name = map[int32]string{
		0: "STATE_UNSPECIFIED",
		1: "STATE_ACCEPTED",
		2: "STATE_FAILED",
	}
	MyIP_State_value = map[string]int32{
		"STATE_UNSPECIFIED": 0,
		"STATE_ACCEPTED":    1,
		"STATE_FAILED":      2,
	}
)

func (x MyIP_State) Enum() *MyIP_State {
	p := new(MyIP_State)
	*p = x
	return p
}

func (x MyIP_State) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MyIP_State) Descriptor() protoreflect.EnumDescriptor {
	return file_myip_v1alpha1_myip_proto_enumTypes[1].Descriptor()
}

func (MyIP_State) Type() protoreflect.EnumType {
	return &file_myip_v1alpha1_myip_proto_enumTypes[1]
}

func (x MyIP_State) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MyIP_State.Descriptor instead.
func (MyIP_State) EnumDescriptor() ([]byte, []int) {
	return file_myip_v1alpha1_myip_proto_rawDescGZIP(), []int{0, 0}
}

// myip リソースメッセージ
type MyIP struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// myipリソース名
	// /myips/{id}
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	//　IP Type
	IpType IPType `protobuf:"varint,2,opt,name=ip_type,json=ipType,proto3,enum=myip.v1alpha1.IPType" json:"ip_type,omitempty"`
	// ip address
	IpAddress string `protobuf:"bytes,3,opt,name=ip_address,json=ipAddress,proto3" json:"ip_address,omitempty"`
	// 状態
	State MyIP_State `protobuf:"varint,4,opt,name=state,proto3,enum=myip.v1alpha1.MyIP_State" json:"state,omitempty"`
	// 作成日時
	CreateTime *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// 更新日
	UpdateTime *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=update_time,json=updateTime,proto3" json:"update_time,omitempty"`
}

func (x *MyIP) Reset() {
	*x = MyIP{}
	if protoimpl.UnsafeEnabled {
		mi := &file_myip_v1alpha1_myip_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MyIP) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MyIP) ProtoMessage() {}

func (x *MyIP) ProtoReflect() protoreflect.Message {
	mi := &file_myip_v1alpha1_myip_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MyIP.ProtoReflect.Descriptor instead.
func (*MyIP) Descriptor() ([]byte, []int) {
	return file_myip_v1alpha1_myip_proto_rawDescGZIP(), []int{0}
}

func (x *MyIP) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *MyIP) GetIpType() IPType {
	if x != nil {
		return x.IpType
	}
	return IPType_IP_TYPE_UNSPECIFIED
}

func (x *MyIP) GetIpAddress() string {
	if x != nil {
		return x.IpAddress
	}
	return ""
}

func (x *MyIP) GetState() MyIP_State {
	if x != nil {
		return x.State
	}
	return MyIP_STATE_UNSPECIFIED
}

func (x *MyIP) GetCreateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.CreateTime
	}
	return nil
}

func (x *MyIP) GetUpdateTime() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdateTime
	}
	return nil
}

var File_myip_v1alpha1_myip_proto protoreflect.FileDescriptor

var file_myip_v1alpha1_myip_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6d, 0x79, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x6d, 0x79, 0x69, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6d, 0x79, 0x69, 0x70,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x72, 0x70, 0x63, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0xda, 0x02, 0x0a, 0x04, 0x4d, 0x79, 0x49, 0x50, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x2e, 0x0a, 0x07, 0x69, 0x70, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x15, 0x2e, 0x6d, 0x79, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x49, 0x50, 0x54, 0x79, 0x70, 0x65, 0x52, 0x06, 0x69, 0x70, 0x54, 0x79, 0x70, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x69, 0x70, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x69, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12,
	0x2f, 0x0a, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19,
	0x2e, 0x6d, 0x79, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d,
	0x79, 0x49, 0x50, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x3b, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x3b, 0x0a,
	0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x22, 0x44, 0x0a, 0x05, 0x53, 0x74,
	0x61, 0x74, 0x65, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x12, 0x0a, 0x0e, 0x53, 0x54,
	0x41, 0x54, 0x45, 0x5f, 0x41, 0x43, 0x43, 0x45, 0x50, 0x54, 0x45, 0x44, 0x10, 0x01, 0x12, 0x10,
	0x0a, 0x0c, 0x53, 0x54, 0x41, 0x54, 0x45, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x02,
	0x2a, 0x4a, 0x0a, 0x06, 0x49, 0x50, 0x54, 0x79, 0x70, 0x65, 0x12, 0x17, 0x0a, 0x13, 0x49, 0x50,
	0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x10, 0x0a, 0x0c, 0x49, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x49,
	0x50, 0x56, 0x34, 0x10, 0x01, 0x12, 0x15, 0x0a, 0x11, 0x49, 0x50, 0x5f, 0x54, 0x59, 0x50, 0x45,
	0x5f, 0x55, 0x4e, 0x49, 0x56, 0x45, 0x52, 0x53, 0x41, 0x4c, 0x10, 0x02, 0x42, 0xc2, 0x01, 0x0a,
	0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x79, 0x69, 0x70, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x09, 0x4d, 0x79, 0x69, 0x70, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x4d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x6c, 0x75, 0x73,
	0x6d, 0x65, 0x64, 0x69, 0x2f, 0x6d, 0x68, 0x76, 0x32, 0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x73, 0x2f, 0x6d, 0x6f, 0x63, 0x2f, 0x73, 0x6d, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61,
	0x70, 0x69, 0x73, 0x2f, 0x6d, 0x79, 0x69, 0x70, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x6d, 0x79, 0x69, 0x70, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x4d, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x4d, 0x79, 0x69, 0x70, 0x2e, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x4d, 0x79, 0x69, 0x70, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x4d, 0x79, 0x69, 0x70, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x0e, 0x4d, 0x79, 0x69, 0x70, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_myip_v1alpha1_myip_proto_rawDescOnce sync.Once
	file_myip_v1alpha1_myip_proto_rawDescData = file_myip_v1alpha1_myip_proto_rawDesc
)

func file_myip_v1alpha1_myip_proto_rawDescGZIP() []byte {
	file_myip_v1alpha1_myip_proto_rawDescOnce.Do(func() {
		file_myip_v1alpha1_myip_proto_rawDescData = protoimpl.X.CompressGZIP(file_myip_v1alpha1_myip_proto_rawDescData)
	})
	return file_myip_v1alpha1_myip_proto_rawDescData
}

var file_myip_v1alpha1_myip_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_myip_v1alpha1_myip_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_myip_v1alpha1_myip_proto_goTypes = []interface{}{
	(IPType)(0),                   // 0: myip.v1alpha1.IPType
	(MyIP_State)(0),               // 1: myip.v1alpha1.MyIP.State
	(*MyIP)(nil),                  // 2: myip.v1alpha1.MyIP
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_myip_v1alpha1_myip_proto_depIdxs = []int32{
	0, // 0: myip.v1alpha1.MyIP.ip_type:type_name -> myip.v1alpha1.IPType
	1, // 1: myip.v1alpha1.MyIP.state:type_name -> myip.v1alpha1.MyIP.State
	3, // 2: myip.v1alpha1.MyIP.create_time:type_name -> google.protobuf.Timestamp
	3, // 3: myip.v1alpha1.MyIP.update_time:type_name -> google.protobuf.Timestamp
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_myip_v1alpha1_myip_proto_init() }
func file_myip_v1alpha1_myip_proto_init() {
	if File_myip_v1alpha1_myip_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_myip_v1alpha1_myip_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MyIP); i {
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
			RawDescriptor: file_myip_v1alpha1_myip_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_myip_v1alpha1_myip_proto_goTypes,
		DependencyIndexes: file_myip_v1alpha1_myip_proto_depIdxs,
		EnumInfos:         file_myip_v1alpha1_myip_proto_enumTypes,
		MessageInfos:      file_myip_v1alpha1_myip_proto_msgTypes,
	}.Build()
	File_myip_v1alpha1_myip_proto = out.File
	file_myip_v1alpha1_myip_proto_rawDesc = nil
	file_myip_v1alpha1_myip_proto_goTypes = nil
	file_myip_v1alpha1_myip_proto_depIdxs = nil
}
