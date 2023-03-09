// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.22.0--rc2
// source: internal/handlers/hendler_gRPC/proto/contact.proto

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

type Contact struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Fio   string `protobuf:"bytes,2,opt,name=Fio,proto3" json:"Fio,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=Email,proto3" json:"Email,omitempty"`
	Phone string `protobuf:"bytes,4,opt,name=Phone,proto3" json:"Phone,omitempty"`
}

func (x *Contact) Reset() {
	*x = Contact{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Contact) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Contact) ProtoMessage() {}

func (x *Contact) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Contact.ProtoReflect.Descriptor instead.
func (*Contact) Descriptor() ([]byte, []int) {
	return file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescGZIP(), []int{0}
}

func (x *Contact) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Contact) GetFio() string {
	if x != nil {
		return x.Fio
	}
	return ""
}

func (x *Contact) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Contact) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Id struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *Id) Reset() {
	*x = Id{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Id) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Id) ProtoMessage() {}

func (x *Id) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Id.ProtoReflect.Descriptor instead.
func (*Id) Descriptor() ([]byte, []int) {
	return file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescGZIP(), []int{1}
}

func (x *Id) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UpdateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Contact *Contact `protobuf:"bytes,2,opt,name=contact,proto3" json:"contact,omitempty"`
}

func (x *UpdateRequest) Reset() {
	*x = UpdateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateRequest) ProtoMessage() {}

func (x *UpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateRequest.ProtoReflect.Descriptor instead.
func (*UpdateRequest) Descriptor() ([]byte, []int) {
	return file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *UpdateRequest) GetContact() *Contact {
	if x != nil {
		return x.Contact
	}
	return nil
}

var File_internal_handlers_hendler_gRPC_proto_contact_proto protoreflect.FileDescriptor

var file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDesc = []byte{
	0x0a, 0x32, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x68, 0x61, 0x6e, 0x64, 0x6c,
	0x65, 0x72, 0x73, 0x2f, 0x68, 0x65, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52, 0x50, 0x43,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52,
	0x50, 0x43, 0x22, 0x57, 0x0a, 0x07, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10, 0x0a,
	0x03, 0x46, 0x69, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x46, 0x69, 0x6f, 0x12,
	0x14, 0x0a, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x50, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0x14, 0x0a, 0x02, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x22, 0x50, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2f, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52,
	0x50, 0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x63, 0x74, 0x32, 0xdf, 0x01, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x15, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x1a, 0x10, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65,
	0x72, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2e, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x10, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2e,
	0x49, 0x64, 0x1a, 0x15, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52, 0x50,
	0x43, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x3c, 0x0a, 0x06, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x12, 0x1b, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52,
	0x50, 0x43, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52, 0x50, 0x43, 0x2e,
	0x43, 0x6f, 0x6e, 0x74, 0x61, 0x63, 0x74, 0x12, 0x2c, 0x0a, 0x06, 0x44, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x12, 0x10, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52, 0x50, 0x43,
	0x2e, 0x49, 0x64, 0x1a, 0x10, 0x2e, 0x68, 0x61, 0x6e, 0x64, 0x6c, 0x65, 0x72, 0x5f, 0x67, 0x52,
	0x50, 0x43, 0x2e, 0x49, 0x64, 0x42, 0x05, 0x5a, 0x03, 0x70, 0x62, 0x2f, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescOnce sync.Once
	file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescData = file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDesc
)

func file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescGZIP() []byte {
	file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescOnce.Do(func() {
		file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescData)
	})
	return file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDescData
}

var file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_internal_handlers_hendler_gRPC_proto_contact_proto_goTypes = []interface{}{
	(*Contact)(nil),       // 0: handler_gRPC.Contact
	(*Id)(nil),            // 1: handler_gRPC.Id
	(*UpdateRequest)(nil), // 2: handler_gRPC.UpdateRequest
}
var file_internal_handlers_hendler_gRPC_proto_contact_proto_depIdxs = []int32{
	0, // 0: handler_gRPC.UpdateRequest.contact:type_name -> handler_gRPC.Contact
	0, // 1: handler_gRPC.ContactService.Create:input_type -> handler_gRPC.Contact
	1, // 2: handler_gRPC.ContactService.Get:input_type -> handler_gRPC.Id
	2, // 3: handler_gRPC.ContactService.Update:input_type -> handler_gRPC.UpdateRequest
	1, // 4: handler_gRPC.ContactService.Delete:input_type -> handler_gRPC.Id
	1, // 5: handler_gRPC.ContactService.Create:output_type -> handler_gRPC.Id
	0, // 6: handler_gRPC.ContactService.Get:output_type -> handler_gRPC.Contact
	0, // 7: handler_gRPC.ContactService.Update:output_type -> handler_gRPC.Contact
	1, // 8: handler_gRPC.ContactService.Delete:output_type -> handler_gRPC.Id
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_handlers_hendler_gRPC_proto_contact_proto_init() }
func file_internal_handlers_hendler_gRPC_proto_contact_proto_init() {
	if File_internal_handlers_hendler_gRPC_proto_contact_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Contact); i {
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
		file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Id); i {
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
		file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateRequest); i {
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
			RawDescriptor: file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_handlers_hendler_gRPC_proto_contact_proto_goTypes,
		DependencyIndexes: file_internal_handlers_hendler_gRPC_proto_contact_proto_depIdxs,
		MessageInfos:      file_internal_handlers_hendler_gRPC_proto_contact_proto_msgTypes,
	}.Build()
	File_internal_handlers_hendler_gRPC_proto_contact_proto = out.File
	file_internal_handlers_hendler_gRPC_proto_contact_proto_rawDesc = nil
	file_internal_handlers_hendler_gRPC_proto_contact_proto_goTypes = nil
	file_internal_handlers_hendler_gRPC_proto_contact_proto_depIdxs = nil
}
