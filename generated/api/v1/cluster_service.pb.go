// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v4.25.3
// source: api/v1/cluster_service.proto

package v1

import (
	storage "github.com/stackrox/rox/generated/storage"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type DeploymentFormat int32

const (
	DeploymentFormat_KUBECTL     DeploymentFormat = 0
	DeploymentFormat_HELM        DeploymentFormat = 1
	DeploymentFormat_HELM_VALUES DeploymentFormat = 2
)

// Enum value maps for DeploymentFormat.
var (
	DeploymentFormat_name = map[int32]string{
		0: "KUBECTL",
		1: "HELM",
		2: "HELM_VALUES",
	}
	DeploymentFormat_value = map[string]int32{
		"KUBECTL":     0,
		"HELM":        1,
		"HELM_VALUES": 2,
	}
)

func (x DeploymentFormat) Enum() *DeploymentFormat {
	p := new(DeploymentFormat)
	*p = x
	return p
}

func (x DeploymentFormat) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DeploymentFormat) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_cluster_service_proto_enumTypes[0].Descriptor()
}

func (DeploymentFormat) Type() protoreflect.EnumType {
	return &file_api_v1_cluster_service_proto_enumTypes[0]
}

func (x DeploymentFormat) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DeploymentFormat.Descriptor instead.
func (DeploymentFormat) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{0}
}

type LoadBalancerType int32

const (
	LoadBalancerType_NONE          LoadBalancerType = 0
	LoadBalancerType_LOAD_BALANCER LoadBalancerType = 1
	LoadBalancerType_NODE_PORT     LoadBalancerType = 2
	LoadBalancerType_ROUTE         LoadBalancerType = 3
)

// Enum value maps for LoadBalancerType.
var (
	LoadBalancerType_name = map[int32]string{
		0: "NONE",
		1: "LOAD_BALANCER",
		2: "NODE_PORT",
		3: "ROUTE",
	}
	LoadBalancerType_value = map[string]int32{
		"NONE":          0,
		"LOAD_BALANCER": 1,
		"NODE_PORT":     2,
		"ROUTE":         3,
	}
)

func (x LoadBalancerType) Enum() *LoadBalancerType {
	p := new(LoadBalancerType)
	*p = x
	return p
}

func (x LoadBalancerType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoadBalancerType) Descriptor() protoreflect.EnumDescriptor {
	return file_api_v1_cluster_service_proto_enumTypes[1].Descriptor()
}

func (LoadBalancerType) Type() protoreflect.EnumType {
	return &file_api_v1_cluster_service_proto_enumTypes[1]
}

func (x LoadBalancerType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoadBalancerType.Descriptor instead.
func (LoadBalancerType) EnumDescriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{1}
}

// next available tag: 3
type DecommissionedClusterRetentionInfo struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to RetentionInfo:
	//
	//	*DecommissionedClusterRetentionInfo_IsExcluded
	//	*DecommissionedClusterRetentionInfo_DaysUntilDeletion
	RetentionInfo isDecommissionedClusterRetentionInfo_RetentionInfo `protobuf_oneof:"RetentionInfo"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DecommissionedClusterRetentionInfo) Reset() {
	*x = DecommissionedClusterRetentionInfo{}
	mi := &file_api_v1_cluster_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DecommissionedClusterRetentionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DecommissionedClusterRetentionInfo) ProtoMessage() {}

func (x *DecommissionedClusterRetentionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DecommissionedClusterRetentionInfo.ProtoReflect.Descriptor instead.
func (*DecommissionedClusterRetentionInfo) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{0}
}

func (x *DecommissionedClusterRetentionInfo) GetRetentionInfo() isDecommissionedClusterRetentionInfo_RetentionInfo {
	if x != nil {
		return x.RetentionInfo
	}
	return nil
}

func (x *DecommissionedClusterRetentionInfo) GetIsExcluded() bool {
	if x != nil {
		if x, ok := x.RetentionInfo.(*DecommissionedClusterRetentionInfo_IsExcluded); ok {
			return x.IsExcluded
		}
	}
	return false
}

func (x *DecommissionedClusterRetentionInfo) GetDaysUntilDeletion() int32 {
	if x != nil {
		if x, ok := x.RetentionInfo.(*DecommissionedClusterRetentionInfo_DaysUntilDeletion); ok {
			return x.DaysUntilDeletion
		}
	}
	return 0
}

type isDecommissionedClusterRetentionInfo_RetentionInfo interface {
	isDecommissionedClusterRetentionInfo_RetentionInfo()
}

type DecommissionedClusterRetentionInfo_IsExcluded struct {
	// indicates whether a cluster is protected from deletion
	IsExcluded bool `protobuf:"varint,1,opt,name=is_excluded,json=isExcluded,proto3,oneof"`
}

type DecommissionedClusterRetentionInfo_DaysUntilDeletion struct {
	// days after which cluster will be deleted if sensor health remains UNHEALTHY
	DaysUntilDeletion int32 `protobuf:"varint,2,opt,name=days_until_deletion,json=daysUntilDeletion,proto3,oneof"`
}

func (*DecommissionedClusterRetentionInfo_IsExcluded) isDecommissionedClusterRetentionInfo_RetentionInfo() {
}

func (*DecommissionedClusterRetentionInfo_DaysUntilDeletion) isDecommissionedClusterRetentionInfo_RetentionInfo() {
}

type ClusterResponse struct {
	state                protoimpl.MessageState              `protogen:"open.v1"`
	Cluster              *storage.Cluster                    `protobuf:"bytes,1,opt,name=cluster,proto3" json:"cluster,omitempty"`
	ClusterRetentionInfo *DecommissionedClusterRetentionInfo `protobuf:"bytes,2,opt,name=cluster_retention_info,json=clusterRetentionInfo,proto3" json:"cluster_retention_info,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ClusterResponse) Reset() {
	*x = ClusterResponse{}
	mi := &file_api_v1_cluster_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterResponse) ProtoMessage() {}

func (x *ClusterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterResponse.ProtoReflect.Descriptor instead.
func (*ClusterResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{1}
}

func (x *ClusterResponse) GetCluster() *storage.Cluster {
	if x != nil {
		return x.Cluster
	}
	return nil
}

func (x *ClusterResponse) GetClusterRetentionInfo() *DecommissionedClusterRetentionInfo {
	if x != nil {
		return x.ClusterRetentionInfo
	}
	return nil
}

type ClusterDefaultsResponse struct {
	state                    protoimpl.MessageState `protogen:"open.v1"`
	MainImageRepository      string                 `protobuf:"bytes,1,opt,name=main_image_repository,json=mainImageRepository,proto3" json:"main_image_repository,omitempty"`
	CollectorImageRepository string                 `protobuf:"bytes,2,opt,name=collector_image_repository,json=collectorImageRepository,proto3" json:"collector_image_repository,omitempty"`
	KernelSupportAvailable   bool                   `protobuf:"varint,3,opt,name=kernel_support_available,json=kernelSupportAvailable,proto3" json:"kernel_support_available,omitempty"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ClusterDefaultsResponse) Reset() {
	*x = ClusterDefaultsResponse{}
	mi := &file_api_v1_cluster_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClusterDefaultsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClusterDefaultsResponse) ProtoMessage() {}

func (x *ClusterDefaultsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClusterDefaultsResponse.ProtoReflect.Descriptor instead.
func (*ClusterDefaultsResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{2}
}

func (x *ClusterDefaultsResponse) GetMainImageRepository() string {
	if x != nil {
		return x.MainImageRepository
	}
	return ""
}

func (x *ClusterDefaultsResponse) GetCollectorImageRepository() string {
	if x != nil {
		return x.CollectorImageRepository
	}
	return ""
}

func (x *ClusterDefaultsResponse) GetKernelSupportAvailable() bool {
	if x != nil {
		return x.KernelSupportAvailable
	}
	return false
}

type ClustersList struct {
	state    protoimpl.MessageState `protogen:"open.v1"`
	Clusters []*storage.Cluster     `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	// Maps 'UNHEALTHY' clusters' IDs to their retention info
	ClusterIdToRetentionInfo map[string]*DecommissionedClusterRetentionInfo `protobuf:"bytes,2,rep,name=cluster_id_to_retention_info,json=clusterIdToRetentionInfo,proto3" json:"cluster_id_to_retention_info,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	unknownFields            protoimpl.UnknownFields
	sizeCache                protoimpl.SizeCache
}

func (x *ClustersList) Reset() {
	*x = ClustersList{}
	mi := &file_api_v1_cluster_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClustersList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClustersList) ProtoMessage() {}

func (x *ClustersList) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClustersList.ProtoReflect.Descriptor instead.
func (*ClustersList) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{3}
}

func (x *ClustersList) GetClusters() []*storage.Cluster {
	if x != nil {
		return x.Clusters
	}
	return nil
}

func (x *ClustersList) GetClusterIdToRetentionInfo() map[string]*DecommissionedClusterRetentionInfo {
	if x != nil {
		return x.ClusterIdToRetentionInfo
	}
	return nil
}

type GetClustersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Query         string                 `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetClustersRequest) Reset() {
	*x = GetClustersRequest{}
	mi := &file_api_v1_cluster_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetClustersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetClustersRequest) ProtoMessage() {}

func (x *GetClustersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetClustersRequest.ProtoReflect.Descriptor instead.
func (*GetClustersRequest) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetClustersRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

// Deprecated: Marked as deprecated in api/v1/cluster_service.proto.
type KernelSupportAvailableResponse struct {
	state                  protoimpl.MessageState `protogen:"open.v1"`
	KernelSupportAvailable bool                   `protobuf:"varint,1,opt,name=kernel_support_available,json=kernelSupportAvailable,proto3" json:"kernel_support_available,omitempty"`
	unknownFields          protoimpl.UnknownFields
	sizeCache              protoimpl.SizeCache
}

func (x *KernelSupportAvailableResponse) Reset() {
	*x = KernelSupportAvailableResponse{}
	mi := &file_api_v1_cluster_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KernelSupportAvailableResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KernelSupportAvailableResponse) ProtoMessage() {}

func (x *KernelSupportAvailableResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_v1_cluster_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KernelSupportAvailableResponse.ProtoReflect.Descriptor instead.
func (*KernelSupportAvailableResponse) Descriptor() ([]byte, []int) {
	return file_api_v1_cluster_service_proto_rawDescGZIP(), []int{5}
}

func (x *KernelSupportAvailableResponse) GetKernelSupportAvailable() bool {
	if x != nil {
		return x.KernelSupportAvailable
	}
	return false
}

var File_api_v1_cluster_service_proto protoreflect.FileDescriptor

var file_api_v1_cluster_service_proto_rawDesc = string([]byte{
	0x0a, 0x1c, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02,
	0x76, 0x31, 0x1a, 0x13, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x12, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f,
	0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x73, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x8a, 0x01, 0x0a, 0x22, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x65, 0x78,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0a,
	0x69, 0x73, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x30, 0x0a, 0x13, 0x64, 0x61,
	0x79, 0x73, 0x5f, 0x75, 0x6e, 0x74, 0x69, 0x6c, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f,
	0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x11, 0x64, 0x61, 0x79, 0x73, 0x55,
	0x6e, 0x74, 0x69, 0x6c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x0a, 0x0d,
	0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0x9b, 0x01,
	0x0a, 0x0f, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x52, 0x07, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x5c, 0x0a,
	0x16, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69,
	0x6f, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x65,
	0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x14, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65,
	0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x22, 0xc5, 0x01, 0x0a, 0x17,
	0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x15, 0x6d, 0x61, 0x69, 0x6e, 0x5f,
	0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x13, 0x6d, 0x61, 0x69, 0x6e, 0x49, 0x6d, 0x61, 0x67,
	0x65, 0x52, 0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3c, 0x0a, 0x1a, 0x63,
	0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x5f, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x72,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x18, 0x63, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x52,
	0x65, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x38, 0x0a, 0x18, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x5f, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69,
	0x6c, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x6b, 0x65, 0x72,
	0x6e, 0x65, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61,
	0x62, 0x6c, 0x65, 0x22, 0xa1, 0x02, 0x0a, 0x0c, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65,
	0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x08, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x6e, 0x0a, 0x1c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x5f, 0x74, 0x6f, 0x5f, 0x72, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x6e,
	0x66, 0x6f, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x49, 0x64, 0x54, 0x6f, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x49,
	0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x18, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65,
	0x72, 0x49, 0x64, 0x54, 0x6f, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e,
	0x66, 0x6f, 0x1a, 0x73, 0x0a, 0x1d, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x49, 0x64, 0x54,
	0x6f, 0x52, 0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x3c, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x6d, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x65, 0x64, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x74, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6c,
	0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x22, 0x5e, 0x0a, 0x1e, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x38, 0x0a, 0x18, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x5f,
	0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x61, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x16, 0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x53,
	0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x3a,
	0x02, 0x18, 0x01, 0x2a, 0x3a, 0x0a, 0x10, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e,
	0x74, 0x46, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x0b, 0x0a, 0x07, 0x4b, 0x55, 0x42, 0x45, 0x43,
	0x54, 0x4c, 0x10, 0x00, 0x12, 0x08, 0x0a, 0x04, 0x48, 0x45, 0x4c, 0x4d, 0x10, 0x01, 0x12, 0x0f,
	0x0a, 0x0b, 0x48, 0x45, 0x4c, 0x4d, 0x5f, 0x56, 0x41, 0x4c, 0x55, 0x45, 0x53, 0x10, 0x02, 0x2a,
	0x49, 0x0a, 0x10, 0x4c, 0x6f, 0x61, 0x64, 0x42, 0x61, 0x6c, 0x61, 0x6e, 0x63, 0x65, 0x72, 0x54,
	0x79, 0x70, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x00, 0x12, 0x11, 0x0a,
	0x0d, 0x4c, 0x4f, 0x41, 0x44, 0x5f, 0x42, 0x41, 0x4c, 0x41, 0x4e, 0x43, 0x45, 0x52, 0x10, 0x01,
	0x12, 0x0d, 0x0a, 0x09, 0x4e, 0x4f, 0x44, 0x45, 0x5f, 0x50, 0x4f, 0x52, 0x54, 0x10, 0x02, 0x12,
	0x09, 0x0a, 0x05, 0x52, 0x4f, 0x55, 0x54, 0x45, 0x10, 0x03, 0x32, 0xff, 0x04, 0x0a, 0x0f, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4d,
	0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12,
	0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x4e, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79, 0x49, 0x44, 0x1a, 0x13, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x12, 0x11, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x4d, 0x0a,
	0x0b, 0x50, 0x6f, 0x73, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x13,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11, 0x3a, 0x01, 0x2a, 0x22, 0x0c,
	0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x12, 0x51, 0x0a, 0x0a,
	0x50, 0x75, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x12, 0x10, 0x2e, 0x73, 0x74, 0x6f,
	0x72, 0x61, 0x67, 0x65, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x1a, 0x13, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x1a, 0x11, 0x2f, 0x76,
	0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12,
	0x47, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x12, 0x10, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x42, 0x79,
	0x49, 0x44, 0x1a, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x19, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x13, 0x2a, 0x11, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x80, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74,
	0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x53, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x41, 0x76, 0x61,
	0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x22, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x53, 0x75, 0x70,
	0x70, 0x6f, 0x72, 0x74, 0x41, 0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x34, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x2b, 0x12, 0x29, 0x2f,
	0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x73, 0x2d, 0x65, 0x6e, 0x76, 0x2f,
	0x6b, 0x65, 0x72, 0x6e, 0x65, 0x6c, 0x2d, 0x73, 0x75, 0x70, 0x70, 0x6f, 0x72, 0x74, 0x2d, 0x61,
	0x76, 0x61, 0x69, 0x6c, 0x61, 0x62, 0x6c, 0x65, 0x88, 0x02, 0x01, 0x12, 0x5f, 0x0a, 0x17, 0x47,
	0x65, 0x74, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x12, 0x09, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1b, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x44, 0x65,
	0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1c,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x12, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6c, 0x75, 0x73,
	0x74, 0x65, 0x72, 0x2d, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x73, 0x42, 0x27, 0x0a, 0x18,
	0x69, 0x6f, 0x2e, 0x73, 0x74, 0x61, 0x63, 0x6b, 0x72, 0x6f, 0x78, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x5a, 0x0b, 0x2e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x76, 0x31, 0x3b, 0x76, 0x31, 0x58, 0x02, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
})

var (
	file_api_v1_cluster_service_proto_rawDescOnce sync.Once
	file_api_v1_cluster_service_proto_rawDescData []byte
)

func file_api_v1_cluster_service_proto_rawDescGZIP() []byte {
	file_api_v1_cluster_service_proto_rawDescOnce.Do(func() {
		file_api_v1_cluster_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_v1_cluster_service_proto_rawDesc), len(file_api_v1_cluster_service_proto_rawDesc)))
	})
	return file_api_v1_cluster_service_proto_rawDescData
}

var file_api_v1_cluster_service_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_api_v1_cluster_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_v1_cluster_service_proto_goTypes = []any{
	(DeploymentFormat)(0),                      // 0: v1.DeploymentFormat
	(LoadBalancerType)(0),                      // 1: v1.LoadBalancerType
	(*DecommissionedClusterRetentionInfo)(nil), // 2: v1.DecommissionedClusterRetentionInfo
	(*ClusterResponse)(nil),                    // 3: v1.ClusterResponse
	(*ClusterDefaultsResponse)(nil),            // 4: v1.ClusterDefaultsResponse
	(*ClustersList)(nil),                       // 5: v1.ClustersList
	(*GetClustersRequest)(nil),                 // 6: v1.GetClustersRequest
	(*KernelSupportAvailableResponse)(nil),     // 7: v1.KernelSupportAvailableResponse
	nil,                                        // 8: v1.ClustersList.ClusterIdToRetentionInfoEntry
	(*storage.Cluster)(nil),                    // 9: storage.Cluster
	(*ResourceByID)(nil),                       // 10: v1.ResourceByID
	(*Empty)(nil),                              // 11: v1.Empty
}
var file_api_v1_cluster_service_proto_depIdxs = []int32{
	9,  // 0: v1.ClusterResponse.cluster:type_name -> storage.Cluster
	2,  // 1: v1.ClusterResponse.cluster_retention_info:type_name -> v1.DecommissionedClusterRetentionInfo
	9,  // 2: v1.ClustersList.clusters:type_name -> storage.Cluster
	8,  // 3: v1.ClustersList.cluster_id_to_retention_info:type_name -> v1.ClustersList.ClusterIdToRetentionInfoEntry
	2,  // 4: v1.ClustersList.ClusterIdToRetentionInfoEntry.value:type_name -> v1.DecommissionedClusterRetentionInfo
	6,  // 5: v1.ClustersService.GetClusters:input_type -> v1.GetClustersRequest
	10, // 6: v1.ClustersService.GetCluster:input_type -> v1.ResourceByID
	9,  // 7: v1.ClustersService.PostCluster:input_type -> storage.Cluster
	9,  // 8: v1.ClustersService.PutCluster:input_type -> storage.Cluster
	10, // 9: v1.ClustersService.DeleteCluster:input_type -> v1.ResourceByID
	11, // 10: v1.ClustersService.GetKernelSupportAvailable:input_type -> v1.Empty
	11, // 11: v1.ClustersService.GetClusterDefaultValues:input_type -> v1.Empty
	5,  // 12: v1.ClustersService.GetClusters:output_type -> v1.ClustersList
	3,  // 13: v1.ClustersService.GetCluster:output_type -> v1.ClusterResponse
	3,  // 14: v1.ClustersService.PostCluster:output_type -> v1.ClusterResponse
	3,  // 15: v1.ClustersService.PutCluster:output_type -> v1.ClusterResponse
	11, // 16: v1.ClustersService.DeleteCluster:output_type -> v1.Empty
	7,  // 17: v1.ClustersService.GetKernelSupportAvailable:output_type -> v1.KernelSupportAvailableResponse
	4,  // 18: v1.ClustersService.GetClusterDefaultValues:output_type -> v1.ClusterDefaultsResponse
	12, // [12:19] is the sub-list for method output_type
	5,  // [5:12] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_api_v1_cluster_service_proto_init() }
func file_api_v1_cluster_service_proto_init() {
	if File_api_v1_cluster_service_proto != nil {
		return
	}
	file_api_v1_common_proto_init()
	file_api_v1_empty_proto_init()
	file_api_v1_cluster_service_proto_msgTypes[0].OneofWrappers = []any{
		(*DecommissionedClusterRetentionInfo_IsExcluded)(nil),
		(*DecommissionedClusterRetentionInfo_DaysUntilDeletion)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_v1_cluster_service_proto_rawDesc), len(file_api_v1_cluster_service_proto_rawDesc)),
			NumEnums:      2,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_v1_cluster_service_proto_goTypes,
		DependencyIndexes: file_api_v1_cluster_service_proto_depIdxs,
		EnumInfos:         file_api_v1_cluster_service_proto_enumTypes,
		MessageInfos:      file_api_v1_cluster_service_proto_msgTypes,
	}.Build()
	File_api_v1_cluster_service_proto = out.File
	file_api_v1_cluster_service_proto_goTypes = nil
	file_api_v1_cluster_service_proto_depIdxs = nil
}
