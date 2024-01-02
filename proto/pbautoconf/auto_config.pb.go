// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: proto/pbautoconf/auto_config.proto

package pbautoconf

import (
	pbconfig "github.com/hashicorp/consul/proto/pbconfig"
	pbconnect "github.com/hashicorp/consul/proto/pbconnect"
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

// AutoConfigRequest is the data structure to be sent along with the
// AutoConfig.InitialConfiguration RPC
type AutoConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Datacenter is the local datacenter name. This wont actually be set by clients
	// but rather will be set by the servers to allow for forwarding to
	// the leader. If it ever happens to be set and differs from the local datacenters
	// name then an error should be returned.
	Datacenter string `protobuf:"bytes,1,opt,name=Datacenter,proto3" json:"Datacenter,omitempty"`
	// Node is the node name that the requester would like to assume
	// the identity of.
	Node string `protobuf:"bytes,2,opt,name=Node,proto3" json:"Node,omitempty"`
	// Segment is the network segment that the requester would like to join
	Segment string `protobuf:"bytes,4,opt,name=Segment,proto3" json:"Segment,omitempty"`
	// Partition is the partition that the requester would like to join
	Partition string `protobuf:"bytes,8,opt,name=Partition,proto3" json:"Partition,omitempty"`
	// JWT is a signed JSON Web Token used to authorize the request
	JWT string `protobuf:"bytes,5,opt,name=JWT,proto3" json:"JWT,omitempty"`
	// ConsulToken is a Consul ACL token that the agent requesting the
	// configuration already has.
	ConsulToken string `protobuf:"bytes,6,opt,name=ConsulToken,proto3" json:"ConsulToken,omitempty"`
	// CSR is a certificate signing request to be used when generating the
	// agents TLS certificate
	CSR string `protobuf:"bytes,7,opt,name=CSR,proto3" json:"CSR,omitempty"`
}

func (x *AutoConfigRequest) Reset() {
	*x = AutoConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pbautoconf_auto_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoConfigRequest) ProtoMessage() {}

func (x *AutoConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbautoconf_auto_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoConfigRequest.ProtoReflect.Descriptor instead.
func (*AutoConfigRequest) Descriptor() ([]byte, []int) {
	return file_proto_pbautoconf_auto_config_proto_rawDescGZIP(), []int{0}
}

func (x *AutoConfigRequest) GetDatacenter() string {
	if x != nil {
		return x.Datacenter
	}
	return ""
}

func (x *AutoConfigRequest) GetNode() string {
	if x != nil {
		return x.Node
	}
	return ""
}

func (x *AutoConfigRequest) GetSegment() string {
	if x != nil {
		return x.Segment
	}
	return ""
}

func (x *AutoConfigRequest) GetPartition() string {
	if x != nil {
		return x.Partition
	}
	return ""
}

func (x *AutoConfigRequest) GetJWT() string {
	if x != nil {
		return x.JWT
	}
	return ""
}

func (x *AutoConfigRequest) GetConsulToken() string {
	if x != nil {
		return x.ConsulToken
	}
	return ""
}

func (x *AutoConfigRequest) GetCSR() string {
	if x != nil {
		return x.CSR
	}
	return ""
}

// AutoConfigResponse is the data structure sent in response to a AutoConfig.InitialConfiguration request
type AutoConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Config is the partial Consul configuration to inject into the agents own configuration
	Config *pbconfig.Config `protobuf:"bytes,1,opt,name=Config,proto3" json:"Config,omitempty"`
	// CARoots is the current list of Connect CA Roots
	CARoots *pbconnect.CARoots `protobuf:"bytes,2,opt,name=CARoots,proto3" json:"CARoots,omitempty"`
	// Certificate is the TLS certificate issued for the agent
	Certificate *pbconnect.IssuedCert `protobuf:"bytes,3,opt,name=Certificate,proto3" json:"Certificate,omitempty"`
	// ExtraCACertificates holds non-Connect certificates that may be necessary
	// to verify TLS connections with the Consul servers
	ExtraCACertificates []string `protobuf:"bytes,4,rep,name=ExtraCACertificates,proto3" json:"ExtraCACertificates,omitempty"`
}

func (x *AutoConfigResponse) Reset() {
	*x = AutoConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pbautoconf_auto_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutoConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutoConfigResponse) ProtoMessage() {}

func (x *AutoConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pbautoconf_auto_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutoConfigResponse.ProtoReflect.Descriptor instead.
func (*AutoConfigResponse) Descriptor() ([]byte, []int) {
	return file_proto_pbautoconf_auto_config_proto_rawDescGZIP(), []int{1}
}

func (x *AutoConfigResponse) GetConfig() *pbconfig.Config {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *AutoConfigResponse) GetCARoots() *pbconnect.CARoots {
	if x != nil {
		return x.CARoots
	}
	return nil
}

func (x *AutoConfigResponse) GetCertificate() *pbconnect.IssuedCert {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *AutoConfigResponse) GetExtraCACertificates() []string {
	if x != nil {
		return x.ExtraCACertificates
	}
	return nil
}

var File_proto_pbautoconf_auto_config_proto protoreflect.FileDescriptor

var file_proto_pbautoconf_auto_config_proto_rawDesc = []byte{
	0x0a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f,
	0x6e, 0x66, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x22, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x1a, 0x1b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x63,
	0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2f, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x44, 0x61,
	0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x44, 0x61, 0x74, 0x61, 0x63, 0x65, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x6f,
	0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x53, 0x65, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x61, 0x72, 0x74,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x50, 0x61, 0x72,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x4a, 0x57, 0x54, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x03, 0x4a, 0x57, 0x54, 0x12, 0x20, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x43,
	0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x43, 0x53,
	0x52, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x43, 0x53, 0x52, 0x22, 0x9f, 0x02, 0x0a,
	0x12, 0x41, 0x75, 0x74, 0x6f, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e,
	0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x44, 0x0a, 0x07, 0x43, 0x41, 0x52, 0x6f, 0x6f, 0x74, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x2e, 0x43, 0x41, 0x52, 0x6f, 0x6f,
	0x74, 0x73, 0x52, 0x07, 0x43, 0x41, 0x52, 0x6f, 0x6f, 0x74, 0x73, 0x12, 0x4f, 0x0a, 0x0b, 0x43,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x63, 0x6f, 0x6e,
	0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x63, 0x6f, 0x6e,
	0x6e, 0x65, 0x63, 0x74, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x64, 0x43, 0x65, 0x72, 0x74, 0x52,
	0x0b, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x30, 0x0a, 0x13,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x13, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x43, 0x41, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x73, 0x42, 0x93,
	0x02, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70,
	0x2e, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0x42, 0x0f, 0x41, 0x75, 0x74, 0x6f, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x2c, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x68, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x2f, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x70, 0x62, 0x61, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0xa2, 0x02, 0x04, 0x48, 0x43, 0x49,
	0x41, 0xaa, 0x02, 0x22, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x2e, 0x43, 0x6f,
	0x6e, 0x73, 0x75, 0x6c, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x41, 0x75,
	0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0xca, 0x02, 0x22, 0x48, 0x61, 0x73, 0x68, 0x69, 0x63, 0x6f,
	0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e,
	0x61, 0x6c, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66, 0xe2, 0x02, 0x2e, 0x48, 0x61,
	0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x5c, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x5c, 0x49,
	0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5c, 0x41, 0x75, 0x74, 0x6f, 0x63, 0x6f, 0x6e, 0x66,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x25, 0x48,
	0x61, 0x73, 0x68, 0x69, 0x63, 0x6f, 0x72, 0x70, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c,
	0x3a, 0x3a, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x3a, 0x3a, 0x41, 0x75, 0x74, 0x6f,
	0x63, 0x6f, 0x6e, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pbautoconf_auto_config_proto_rawDescOnce sync.Once
	file_proto_pbautoconf_auto_config_proto_rawDescData = file_proto_pbautoconf_auto_config_proto_rawDesc
)

func file_proto_pbautoconf_auto_config_proto_rawDescGZIP() []byte {
	file_proto_pbautoconf_auto_config_proto_rawDescOnce.Do(func() {
		file_proto_pbautoconf_auto_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pbautoconf_auto_config_proto_rawDescData)
	})
	return file_proto_pbautoconf_auto_config_proto_rawDescData
}

var file_proto_pbautoconf_auto_config_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_pbautoconf_auto_config_proto_goTypes = []interface{}{
	(*AutoConfigRequest)(nil),    // 0: hashicorp.consul.internal.autoconf.AutoConfigRequest
	(*AutoConfigResponse)(nil),   // 1: hashicorp.consul.internal.autoconf.AutoConfigResponse
	(*pbconfig.Config)(nil),      // 2: hashicorp.consul.internal.config.Config
	(*pbconnect.CARoots)(nil),    // 3: hashicorp.consul.internal.connect.CARoots
	(*pbconnect.IssuedCert)(nil), // 4: hashicorp.consul.internal.connect.IssuedCert
}
var file_proto_pbautoconf_auto_config_proto_depIdxs = []int32{
	2, // 0: hashicorp.consul.internal.autoconf.AutoConfigResponse.Config:type_name -> hashicorp.consul.internal.config.Config
	3, // 1: hashicorp.consul.internal.autoconf.AutoConfigResponse.CARoots:type_name -> hashicorp.consul.internal.connect.CARoots
	4, // 2: hashicorp.consul.internal.autoconf.AutoConfigResponse.Certificate:type_name -> hashicorp.consul.internal.connect.IssuedCert
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_pbautoconf_auto_config_proto_init() }
func file_proto_pbautoconf_auto_config_proto_init() {
	if File_proto_pbautoconf_auto_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pbautoconf_auto_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoConfigRequest); i {
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
		file_proto_pbautoconf_auto_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutoConfigResponse); i {
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
			RawDescriptor: file_proto_pbautoconf_auto_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_pbautoconf_auto_config_proto_goTypes,
		DependencyIndexes: file_proto_pbautoconf_auto_config_proto_depIdxs,
		MessageInfos:      file_proto_pbautoconf_auto_config_proto_msgTypes,
	}.Build()
	File_proto_pbautoconf_auto_config_proto = out.File
	file_proto_pbautoconf_auto_config_proto_rawDesc = nil
	file_proto_pbautoconf_auto_config_proto_goTypes = nil
	file_proto_pbautoconf_auto_config_proto_depIdxs = nil
}
