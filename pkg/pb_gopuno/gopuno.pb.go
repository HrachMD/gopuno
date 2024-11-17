// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.28.3
// source: gopuno.proto

package pbgopuno

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

type EmailStatus int32

const (
	EmailStatus_EMAIL_STATUS_UNSPECIFIED EmailStatus = 0
	EmailStatus_EMAIL_STATUS_QUEUED      EmailStatus = 1
	EmailStatus_EMAIL_STATUS_SENDING     EmailStatus = 2
	EmailStatus_EMAIL_STATUS_SENT        EmailStatus = 3
	EmailStatus_EMAIL_STATUS_FAILED      EmailStatus = 4
)

// Enum value maps for EmailStatus.
var (
	EmailStatus_name = map[int32]string{
		0: "EMAIL_STATUS_UNSPECIFIED",
		1: "EMAIL_STATUS_QUEUED",
		2: "EMAIL_STATUS_SENDING",
		3: "EMAIL_STATUS_SENT",
		4: "EMAIL_STATUS_FAILED",
	}
	EmailStatus_value = map[string]int32{
		"EMAIL_STATUS_UNSPECIFIED": 0,
		"EMAIL_STATUS_QUEUED":      1,
		"EMAIL_STATUS_SENDING":     2,
		"EMAIL_STATUS_SENT":        3,
		"EMAIL_STATUS_FAILED":      4,
	}
)

func (x EmailStatus) Enum() *EmailStatus {
	p := new(EmailStatus)
	*p = x
	return p
}

func (x EmailStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (EmailStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_gopuno_proto_enumTypes[0].Descriptor()
}

func (EmailStatus) Type() protoreflect.EnumType {
	return &file_gopuno_proto_enumTypes[0]
}

func (x EmailStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use EmailStatus.Descriptor instead.
func (EmailStatus) EnumDescriptor() ([]byte, []int) {
	return file_gopuno_proto_rawDescGZIP(), []int{0}
}

type Email struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      []byte `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Domain  string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	From    string `protobuf:"bytes,3,opt,name=from,proto3" json:"from,omitempty"`
	To      string `protobuf:"bytes,4,opt,name=to,proto3" json:"to,omitempty"`
	Title   string `protobuf:"bytes,5,opt,name=title,proto3" json:"title,omitempty"`
	Message string `protobuf:"bytes,6,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Email) Reset() {
	*x = Email{}
	mi := &file_gopuno_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Email) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Email) ProtoMessage() {}

func (x *Email) ProtoReflect() protoreflect.Message {
	mi := &file_gopuno_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Email.ProtoReflect.Descriptor instead.
func (*Email) Descriptor() ([]byte, []int) {
	return file_gopuno_proto_rawDescGZIP(), []int{0}
}

func (x *Email) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Email) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *Email) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *Email) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *Email) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Email) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type SendEmailRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email *Email `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *SendEmailRequest) Reset() {
	*x = SendEmailRequest{}
	mi := &file_gopuno_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEmailRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailRequest) ProtoMessage() {}

func (x *SendEmailRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gopuno_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailRequest.ProtoReflect.Descriptor instead.
func (*SendEmailRequest) Descriptor() ([]byte, []int) {
	return file_gopuno_proto_rawDescGZIP(), []int{1}
}

func (x *SendEmailRequest) GetEmail() *Email {
	if x != nil {
		return x.Email
	}
	return nil
}

type SendEmailResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     []byte      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status EmailStatus `protobuf:"varint,2,opt,name=status,proto3,enum=pbgopuno.EmailStatus" json:"status,omitempty"`
}

func (x *SendEmailResponse) Reset() {
	*x = SendEmailResponse{}
	mi := &file_gopuno_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendEmailResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendEmailResponse) ProtoMessage() {}

func (x *SendEmailResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gopuno_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendEmailResponse.ProtoReflect.Descriptor instead.
func (*SendEmailResponse) Descriptor() ([]byte, []int) {
	return file_gopuno_proto_rawDescGZIP(), []int{2}
}

func (x *SendEmailResponse) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *SendEmailResponse) GetStatus() EmailStatus {
	if x != nil {
		return x.Status
	}
	return EmailStatus_EMAIL_STATUS_UNSPECIFIED
}

var File_gopuno_proto protoreflect.FileDescriptor

var file_gopuno_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x6f, 0x70, 0x75, 0x6e, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x70, 0x62, 0x67, 0x6f, 0x70, 0x75, 0x6e, 0x6f, 0x22, 0x83, 0x01, 0x0a, 0x05, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72,
	0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x69, 0x74, 0x6c, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x39,
	0x0a, 0x10, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x70, 0x75, 0x6e, 0x6f, 0x2e, 0x45, 0x6d, 0x61,
	0x69, 0x6c, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x52, 0x0a, 0x11, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2d,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x70, 0x62, 0x67, 0x6f, 0x70, 0x75, 0x6e, 0x6f, 0x2e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x2a, 0x8e, 0x01,
	0x0a, 0x0b, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c, 0x0a,
	0x18, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e,
	0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x17, 0x0a, 0x13, 0x45,
	0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x51, 0x55, 0x45, 0x55,
	0x45, 0x44, 0x10, 0x01, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x53, 0x45, 0x4e, 0x44, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x15,
	0x0a, 0x11, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x53,
	0x45, 0x4e, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x45, 0x4d, 0x41, 0x49, 0x4c, 0x5f, 0x53,
	0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x04, 0x32, 0xa2,
	0x01, 0x0a, 0x0b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x75, 0x73, 0x68, 0x12, 0x49,
	0x0a, 0x0e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41, 0x73, 0x79, 0x6e, 0x63,
	0x12, 0x1a, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x70, 0x75, 0x6e, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70,
	0x62, 0x67, 0x6f, 0x70, 0x75, 0x6e, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69,
	0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48, 0x0a, 0x0d, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x53, 0x79, 0x6e, 0x63, 0x12, 0x1a, 0x2e, 0x70, 0x62, 0x67,
	0x6f, 0x70, 0x75, 0x6e, 0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x62, 0x67, 0x6f, 0x70, 0x75, 0x6e,
	0x6f, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x1d, 0x5a, 0x1b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x48, 0x72, 0x61, 0x63, 0x68, 0x4d, 0x44, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x70, 0x75,
	0x6e, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gopuno_proto_rawDescOnce sync.Once
	file_gopuno_proto_rawDescData = file_gopuno_proto_rawDesc
)

func file_gopuno_proto_rawDescGZIP() []byte {
	file_gopuno_proto_rawDescOnce.Do(func() {
		file_gopuno_proto_rawDescData = protoimpl.X.CompressGZIP(file_gopuno_proto_rawDescData)
	})
	return file_gopuno_proto_rawDescData
}

var file_gopuno_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_gopuno_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_gopuno_proto_goTypes = []any{
	(EmailStatus)(0),          // 0: pbgopuno.EmailStatus
	(*Email)(nil),             // 1: pbgopuno.Email
	(*SendEmailRequest)(nil),  // 2: pbgopuno.SendEmailRequest
	(*SendEmailResponse)(nil), // 3: pbgopuno.SendEmailResponse
}
var file_gopuno_proto_depIdxs = []int32{
	1, // 0: pbgopuno.SendEmailRequest.email:type_name -> pbgopuno.Email
	0, // 1: pbgopuno.SendEmailResponse.status:type_name -> pbgopuno.EmailStatus
	2, // 2: pbgopuno.ServicePush.SendEmailAsync:input_type -> pbgopuno.SendEmailRequest
	2, // 3: pbgopuno.ServicePush.SendEmailSync:input_type -> pbgopuno.SendEmailRequest
	3, // 4: pbgopuno.ServicePush.SendEmailAsync:output_type -> pbgopuno.SendEmailResponse
	3, // 5: pbgopuno.ServicePush.SendEmailSync:output_type -> pbgopuno.SendEmailResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gopuno_proto_init() }
func file_gopuno_proto_init() {
	if File_gopuno_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gopuno_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gopuno_proto_goTypes,
		DependencyIndexes: file_gopuno_proto_depIdxs,
		EnumInfos:         file_gopuno_proto_enumTypes,
		MessageInfos:      file_gopuno_proto_msgTypes,
	}.Build()
	File_gopuno_proto = out.File
	file_gopuno_proto_rawDesc = nil
	file_gopuno_proto_goTypes = nil
	file_gopuno_proto_depIdxs = nil
}