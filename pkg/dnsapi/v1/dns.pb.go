// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.29.0
// 	protoc        (unknown)
// source: dnsapi/v1/dns.proto

package dnsapi

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

// Requests
type AddDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerDomain string `protobuf:"bytes,1,opt,name=serverDomain,proto3" json:"serverDomain,omitempty"`
}

func (x *AddDomainRequest) Reset() {
	*x = AddDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsapi_v1_dns_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDomainRequest) ProtoMessage() {}

func (x *AddDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dnsapi_v1_dns_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDomainRequest.ProtoReflect.Descriptor instead.
func (*AddDomainRequest) Descriptor() ([]byte, []int) {
	return file_dnsapi_v1_dns_proto_rawDescGZIP(), []int{0}
}

func (x *AddDomainRequest) GetServerDomain() string {
	if x != nil {
		return x.ServerDomain
	}
	return ""
}

type GetDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerDomain string `protobuf:"bytes,1,opt,name=serverDomain,proto3" json:"serverDomain,omitempty"`
}

func (x *GetDomainRequest) Reset() {
	*x = GetDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsapi_v1_dns_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainRequest) ProtoMessage() {}

func (x *GetDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dnsapi_v1_dns_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainRequest.ProtoReflect.Descriptor instead.
func (*GetDomainRequest) Descriptor() ([]byte, []int) {
	return file_dnsapi_v1_dns_proto_rawDescGZIP(), []int{1}
}

func (x *GetDomainRequest) GetServerDomain() string {
	if x != nil {
		return x.ServerDomain
	}
	return ""
}

type RemoveDomainRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ServerDomain string `protobuf:"bytes,1,opt,name=serverDomain,proto3" json:"serverDomain,omitempty"`
}

func (x *RemoveDomainRequest) Reset() {
	*x = RemoveDomainRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsapi_v1_dns_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDomainRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDomainRequest) ProtoMessage() {}

func (x *RemoveDomainRequest) ProtoReflect() protoreflect.Message {
	mi := &file_dnsapi_v1_dns_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDomainRequest.ProtoReflect.Descriptor instead.
func (*RemoveDomainRequest) Descriptor() ([]byte, []int) {
	return file_dnsapi_v1_dns_proto_rawDescGZIP(), []int{2}
}

func (x *RemoveDomainRequest) GetServerDomain() string {
	if x != nil {
		return x.ServerDomain
	}
	return ""
}

// Responses
type AddDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddDomainResponse) Reset() {
	*x = AddDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsapi_v1_dns_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddDomainResponse) ProtoMessage() {}

func (x *AddDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dnsapi_v1_dns_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddDomainResponse.ProtoReflect.Descriptor instead.
func (*AddDomainResponse) Descriptor() ([]byte, []int) {
	return file_dnsapi_v1_dns_proto_rawDescGZIP(), []int{3}
}

type GetDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDomainResponse) Reset() {
	*x = GetDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsapi_v1_dns_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDomainResponse) ProtoMessage() {}

func (x *GetDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dnsapi_v1_dns_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDomainResponse.ProtoReflect.Descriptor instead.
func (*GetDomainResponse) Descriptor() ([]byte, []int) {
	return file_dnsapi_v1_dns_proto_rawDescGZIP(), []int{4}
}

type RemoveDomainResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RemoveDomainResponse) Reset() {
	*x = RemoveDomainResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_dnsapi_v1_dns_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveDomainResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveDomainResponse) ProtoMessage() {}

func (x *RemoveDomainResponse) ProtoReflect() protoreflect.Message {
	mi := &file_dnsapi_v1_dns_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveDomainResponse.ProtoReflect.Descriptor instead.
func (*RemoveDomainResponse) Descriptor() ([]byte, []int) {
	return file_dnsapi_v1_dns_proto_rawDescGZIP(), []int{5}
}

var File_dnsapi_v1_dns_proto protoreflect.FileDescriptor

var file_dnsapi_v1_dns_proto_rawDesc = []byte{
	0x0a, 0x13, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x6e, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x22, 0x36, 0x0a, 0x10, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x36, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c,
	0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x22, 0x39, 0x0a, 0x13, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x41,
	0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x16, 0x0a, 0x14, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0xf3, 0x01,
	0x0a, 0x0a, 0x44, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x48, 0x0a, 0x09,
	0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x64, 0x6e, 0x73, 0x61,
	0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d,
	0x61, 0x69, 0x6e, 0x12, 0x1b, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x51, 0x0a, 0x0c, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x12, 0x1e, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1f, 0x2e, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x6d,
	0x6f, 0x76, 0x65, 0x44, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x42, 0x43, 0x5a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x4f, 0x70, 0x65, 0x6e, 0x2d, 0x54, 0x61, 0x78, 0x69, 0x2d, 0x43, 0x6f, 0x6d, 0x6d,
	0x75, 0x6e, 0x69, 0x74, 0x79, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x2d, 0x74, 0x61, 0x78, 0x69, 0x2d,
	0x64, 0x6e, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x2f, 0x76,
	0x31, 0x3b, 0x64, 0x6e, 0x73, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_dnsapi_v1_dns_proto_rawDescOnce sync.Once
	file_dnsapi_v1_dns_proto_rawDescData = file_dnsapi_v1_dns_proto_rawDesc
)

func file_dnsapi_v1_dns_proto_rawDescGZIP() []byte {
	file_dnsapi_v1_dns_proto_rawDescOnce.Do(func() {
		file_dnsapi_v1_dns_proto_rawDescData = protoimpl.X.CompressGZIP(file_dnsapi_v1_dns_proto_rawDescData)
	})
	return file_dnsapi_v1_dns_proto_rawDescData
}

var file_dnsapi_v1_dns_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_dnsapi_v1_dns_proto_goTypes = []interface{}{
	(*AddDomainRequest)(nil),     // 0: dnsapi.v1.AddDomainRequest
	(*GetDomainRequest)(nil),     // 1: dnsapi.v1.GetDomainRequest
	(*RemoveDomainRequest)(nil),  // 2: dnsapi.v1.RemoveDomainRequest
	(*AddDomainResponse)(nil),    // 3: dnsapi.v1.AddDomainResponse
	(*GetDomainResponse)(nil),    // 4: dnsapi.v1.GetDomainResponse
	(*RemoveDomainResponse)(nil), // 5: dnsapi.v1.RemoveDomainResponse
}
var file_dnsapi_v1_dns_proto_depIdxs = []int32{
	0, // 0: dnsapi.v1.DnsService.AddDomain:input_type -> dnsapi.v1.AddDomainRequest
	1, // 1: dnsapi.v1.DnsService.GetDomain:input_type -> dnsapi.v1.GetDomainRequest
	2, // 2: dnsapi.v1.DnsService.RemoveDomain:input_type -> dnsapi.v1.RemoveDomainRequest
	3, // 3: dnsapi.v1.DnsService.AddDomain:output_type -> dnsapi.v1.AddDomainResponse
	4, // 4: dnsapi.v1.DnsService.GetDomain:output_type -> dnsapi.v1.GetDomainResponse
	5, // 5: dnsapi.v1.DnsService.RemoveDomain:output_type -> dnsapi.v1.RemoveDomainResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_dnsapi_v1_dns_proto_init() }
func file_dnsapi_v1_dns_proto_init() {
	if File_dnsapi_v1_dns_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_dnsapi_v1_dns_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDomainRequest); i {
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
		file_dnsapi_v1_dns_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDomainRequest); i {
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
		file_dnsapi_v1_dns_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDomainRequest); i {
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
		file_dnsapi_v1_dns_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddDomainResponse); i {
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
		file_dnsapi_v1_dns_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDomainResponse); i {
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
		file_dnsapi_v1_dns_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveDomainResponse); i {
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
			RawDescriptor: file_dnsapi_v1_dns_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_dnsapi_v1_dns_proto_goTypes,
		DependencyIndexes: file_dnsapi_v1_dns_proto_depIdxs,
		MessageInfos:      file_dnsapi_v1_dns_proto_msgTypes,
	}.Build()
	File_dnsapi_v1_dns_proto = out.File
	file_dnsapi_v1_dns_proto_rawDesc = nil
	file_dnsapi_v1_dns_proto_goTypes = nil
	file_dnsapi_v1_dns_proto_depIdxs = nil
}
