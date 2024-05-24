// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo/projects/gateway2/api/v1alpha1/kube/service.proto

package kube

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/solo-io/protoc-gen-ext/extproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Kubernetes Service type. If not specified, defaults to `ClusterIP`.
// See
// https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
// for details on each service type.
// Currently, only ClusterIP and LoadBalancer are supported.
type Service_ServiceType int32

const (
	// Unspecified allows us to mark a field for inheritance when overriding defaults. This is the default.
	Service_Unspecified Service_ServiceType = 0
	// Exposes the Service on a cluster-internal IP.
	Service_ClusterIP Service_ServiceType = 1
	// Exposes the Service externally using an external load balancer.
	Service_LoadBalancer Service_ServiceType = 2
)

// Enum value maps for Service_ServiceType.
var (
	Service_ServiceType_name = map[int32]string{
		0: "Unspecified",
		1: "ClusterIP",
		2: "LoadBalancer",
	}
	Service_ServiceType_value = map[string]int32{
		"Unspecified":  0,
		"ClusterIP":    1,
		"LoadBalancer": 2,
	}
)

func (x Service_ServiceType) Enum() *Service_ServiceType {
	p := new(Service_ServiceType)
	*p = x
	return p
}

func (x Service_ServiceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Service_ServiceType) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_enumTypes[0].Descriptor()
}

func (Service_ServiceType) Type() protoreflect.EnumType {
	return &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_enumTypes[0]
}

func (x Service_ServiceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Service_ServiceType.Descriptor instead.
func (Service_ServiceType) EnumDescriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescGZIP(), []int{0, 0}
}

// Configuration for a Kubernetes Service.
type Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The Kubernetes Service type.
	Type Service_ServiceType `protobuf:"varint,1,opt,name=type,proto3,enum=kube.gateway.gloo.solo.io.Service_ServiceType" json:"type,omitempty"`
	// The manually specified IP address of the service, if a randomly assigned
	// IP is not desired. See
	// https://kubernetes.io/docs/concepts/services-networking/service/#choosing-your-own-ip-address
	// and
	// https://kubernetes.io/docs/concepts/services-networking/service/#headless-services
	// on the implications of setting `clusterIP`.
	ClusterIP string `protobuf:"bytes,2,opt,name=clusterIP,proto3" json:"clusterIP,omitempty"`
	// Additional labels to add to the Service object metadata.
	ExtraLabels map[string]string `protobuf:"bytes,3,rep,name=extra_labels,json=extraLabels,proto3" json:"extra_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Additional annotations to add to the Service object metadata.
	ExtraAnnotations map[string]string `protobuf:"bytes,4,rep,name=extra_annotations,json=extraAnnotations,proto3" json:"extra_annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Service) Reset() {
	*x = Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Service) ProtoMessage() {}

func (x *Service) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Service.ProtoReflect.Descriptor instead.
func (*Service) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescGZIP(), []int{0}
}

func (x *Service) GetType() Service_ServiceType {
	if x != nil {
		return x.Type
	}
	return Service_Unspecified
}

func (x *Service) GetClusterIP() string {
	if x != nil {
		return x.ClusterIP
	}
	return ""
}

func (x *Service) GetExtraLabels() map[string]string {
	if x != nil {
		return x.ExtraLabels
	}
	return nil
}

func (x *Service) GetExtraAnnotations() map[string]string {
	if x != nil {
		return x.ExtraAnnotations
	}
	return nil
}

var File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDesc = []byte{
	0x0a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63,
	0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x32, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x2f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x19, 0x6b, 0x75, 0x62,
	0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73,
	0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a, 0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf0, 0x03, 0x0a, 0x07, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x2e, 0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f,
	0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x50, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x50, 0x12, 0x56, 0x0a, 0x0c, 0x65, 0x78, 0x74, 0x72,
	0x61, 0x5f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33,
	0x2e, 0x6b, 0x75, 0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c,
	0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0b, 0x65, 0x78, 0x74, 0x72, 0x61, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x12, 0x65, 0x0a, 0x11, 0x65, 0x78, 0x74, 0x72, 0x61, 0x5f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x38, 0x2e, 0x6b, 0x75,
	0x62, 0x65, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e,
	0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x45, 0x78, 0x74, 0x72, 0x61, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x65, 0x78, 0x74, 0x72, 0x61, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x1a, 0x3e, 0x0a, 0x10, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3f, 0x0a, 0x0b,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0f, 0x0a, 0x0b, 0x55,
	0x6e, 0x73, 0x70, 0x65, 0x63, 0x69, 0x66, 0x69, 0x65, 0x64, 0x10, 0x00, 0x12, 0x0d, 0x0a, 0x09,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x50, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x4c,
	0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x10, 0x02, 0x42, 0x5e, 0xb8,
	0xf5, 0x04, 0x01, 0xc0, 0xf5, 0x04, 0x01, 0x5a, 0x54, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x73, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x32, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x6b, 0x75, 0x62, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescData = file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDesc
)

func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDescData
}

var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_goTypes = []interface{}{
	(Service_ServiceType)(0), // 0: kube.gateway.gloo.solo.io.Service.ServiceType
	(*Service)(nil),          // 1: kube.gateway.gloo.solo.io.Service
	nil,                      // 2: kube.gateway.gloo.solo.io.Service.ExtraLabelsEntry
	nil,                      // 3: kube.gateway.gloo.solo.io.Service.ExtraAnnotationsEntry
}
var file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_depIdxs = []int32{
	0, // 0: kube.gateway.gloo.solo.io.Service.type:type_name -> kube.gateway.gloo.solo.io.Service.ServiceType
	2, // 1: kube.gateway.gloo.solo.io.Service.extra_labels:type_name -> kube.gateway.gloo.solo.io.Service.ExtraLabelsEntry
	3, // 2: kube.gateway.gloo.solo.io.Service.extra_annotations:type_name -> kube.gateway.gloo.solo.io.Service.ExtraAnnotationsEntry
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_init() }
func file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_init() {
	if File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Service); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_depIdxs,
		EnumInfos:         file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_enumTypes,
		MessageInfos:      file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto = out.File
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_rawDesc = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_goTypes = nil
	file_github_com_solo_io_gloo_projects_gateway2_api_v1alpha1_kube_service_proto_depIdxs = nil
}
