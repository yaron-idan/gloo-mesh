// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: github.com/solo-io/gloo-mesh/api/settings/v1alpha2/settings.proto

package v1alpha2

import (
	reflect "reflect"
	sync "sync"

	proto "github.com/golang/protobuf/proto"
	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// Configure global settings and defaults.
type SettingsSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Configure default mTLS settings for TrafficTargets (MTLS declared in TrafficPolicies take precedence)
	Mtls *v1alpha2.TrafficPolicySpec_MTLS `protobuf:"bytes,1,opt,name=mtls,proto3" json:"mtls,omitempty"`
	// Configure Gloo Mesh networking to communicate with one or more external gRPC NetworkingExtensions servers.
	// Updates will be applied by the servers in the order they are listed (servers towards the end of the list take precedence).
	// Note: Extension Servers have full write access to the output objects written by Gloo Mesh.
	NetworkingExtensionServers []*NetworkingExtensionsServer `protobuf:"bytes,2,rep,name=networking_extension_servers,json=networkingExtensionServers,proto3" json:"networking_extension_servers,omitempty"`
	// Istio-specific discovery settings
	Istio *SettingsSpec_Istio `protobuf:"bytes,3,opt,name=istio,proto3" json:"istio,omitempty"`
}

func (x *SettingsSpec) Reset() {
	*x = SettingsSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsSpec) ProtoMessage() {}

func (x *SettingsSpec) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsSpec.ProtoReflect.Descriptor instead.
func (*SettingsSpec) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescGZIP(), []int{0}
}

func (x *SettingsSpec) GetMtls() *v1alpha2.TrafficPolicySpec_MTLS {
	if x != nil {
		return x.Mtls
	}
	return nil
}

func (x *SettingsSpec) GetNetworkingExtensionServers() []*NetworkingExtensionsServer {
	if x != nil {
		return x.NetworkingExtensionServers
	}
	return nil
}

func (x *SettingsSpec) GetIstio() *SettingsSpec_Istio {
	if x != nil {
		return x.Istio
	}
	return nil
}

// Options for connecting to an external gRPC NetworkingExtensions server
type NetworkingExtensionsServer struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// TCP address of the Networking Extensions Server (including port)
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Communicate over HTTP rather than HTTPS
	Insecure bool `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	// Instruct Gloo Mesh to automatically reconnect to the server on network failures
	ReconnectOnNetworkFailures bool `protobuf:"varint,3,opt,name=reconnect_on_network_failures,json=reconnectOnNetworkFailures,proto3" json:"reconnect_on_network_failures,omitempty"`
}

func (x *NetworkingExtensionsServer) Reset() {
	*x = NetworkingExtensionsServer{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NetworkingExtensionsServer) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NetworkingExtensionsServer) ProtoMessage() {}

func (x *NetworkingExtensionsServer) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NetworkingExtensionsServer.ProtoReflect.Descriptor instead.
func (*NetworkingExtensionsServer) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescGZIP(), []int{1}
}

func (x *NetworkingExtensionsServer) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *NetworkingExtensionsServer) GetInsecure() bool {
	if x != nil {
		return x.Insecure
	}
	return false
}

func (x *NetworkingExtensionsServer) GetReconnectOnNetworkFailures() bool {
	if x != nil {
		return x.ReconnectOnNetworkFailures
	}
	return false
}

type SettingsStatus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The most recent generation observed in the the Settings metadata.
	// If the observedGeneration does not match generation, the controller has not processed the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource.
	// It will only show accepted if no processing errors encountered.
	State v1alpha2.ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// Any errors encountered while processing Settings object.
	Errors []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
}

func (x *SettingsStatus) Reset() {
	*x = SettingsStatus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsStatus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsStatus) ProtoMessage() {}

func (x *SettingsStatus) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsStatus.ProtoReflect.Descriptor instead.
func (*SettingsStatus) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescGZIP(), []int{2}
}

func (x *SettingsStatus) GetObservedGeneration() int64 {
	if x != nil {
		return x.ObservedGeneration
	}
	return 0
}

func (x *SettingsStatus) GetState() v1alpha2.ApprovalState {
	if x != nil {
		return x.State
	}
	return v1alpha2.ApprovalState_PENDING
}

func (x *SettingsStatus) GetErrors() []string {
	if x != nil {
		return x.Errors
	}
	return nil
}

type SettingsSpec_Istio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Ingress gateway detectors for each cluster. The key to the map is either a k8s cluster name or the wildcard
	// `*` meaning all clusters. If an entry is found for a given cluster, it will be used. Otherwise, the
	// wildcard entry will be used if it exists. Lastly, we will fall back to the default values.
	IngressGatewayDetectors map[string]*SettingsSpec_Istio_IngressGatewayDetector `protobuf:"bytes,1,rep,name=ingress_gateway_detectors,json=ingressGatewayDetectors,proto3" json:"ingress_gateway_detectors,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *SettingsSpec_Istio) Reset() {
	*x = SettingsSpec_Istio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsSpec_Istio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsSpec_Istio) ProtoMessage() {}

func (x *SettingsSpec_Istio) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsSpec_Istio.ProtoReflect.Descriptor instead.
func (*SettingsSpec_Istio) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescGZIP(), []int{0, 0}
}

func (x *SettingsSpec_Istio) GetIngressGatewayDetectors() map[string]*SettingsSpec_Istio_IngressGatewayDetector {
	if x != nil {
		return x.IngressGatewayDetectors
	}
	return nil
}

// Workload labels and TLS port name used during discovery to detect ingress gateways for a mesh.
type SettingsSpec_Istio_IngressGatewayDetector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// The workload labels used during discovery to detect ingress gateways for a mesh.
	// If not specified, will default to `{"istio": "ingressgateway"}`.
	GatewayWorkloadLabels map[string]string `protobuf:"bytes,1,rep,name=gateway_workload_labels,json=gatewayWorkloadLabels,proto3" json:"gateway_workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The name of the TLS port used to detect ingress gateways. Services must have a port with this name
	// in order to be recognized as an ingress gateway during discovery.
	// If not specified, will default to `tls`.
	GatewayTlsPortName string `protobuf:"bytes,2,opt,name=gateway_tls_port_name,json=gatewayTlsPortName,proto3" json:"gateway_tls_port_name,omitempty"`
}

func (x *SettingsSpec_Istio_IngressGatewayDetector) Reset() {
	*x = SettingsSpec_Istio_IngressGatewayDetector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SettingsSpec_Istio_IngressGatewayDetector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SettingsSpec_Istio_IngressGatewayDetector) ProtoMessage() {}

func (x *SettingsSpec_Istio_IngressGatewayDetector) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SettingsSpec_Istio_IngressGatewayDetector.ProtoReflect.Descriptor instead.
func (*SettingsSpec_Istio_IngressGatewayDetector) Descriptor() ([]byte, []int) {
	return file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescGZIP(), []int{0, 0, 1}
}

func (x *SettingsSpec_Istio_IngressGatewayDetector) GetGatewayWorkloadLabels() map[string]string {
	if x != nil {
		return x.GatewayWorkloadLabels
	}
	return nil
}

func (x *SettingsSpec_Istio_IngressGatewayDetector) GetGatewayTlsPortName() string {
	if x != nil {
		return x.GatewayTlsPortName
	}
	return ""
}

var File_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto protoreflect.FileDescriptor

var file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDesc = []byte{
	0x0a, 0x41, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1a, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65,
	0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x1a,
	0x12, 0x65, 0x78, 0x74, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x65, 0x78, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x4b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73,
	0x68, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x49, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c,
	0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x32, 0x2f, 0x74, 0x72, 0x61, 0x66, 0x66, 0x69, 0x63, 0x5f, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf3, 0x06, 0x0a, 0x0c,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x70, 0x65, 0x63, 0x12, 0x48, 0x0a, 0x04,
	0x6d, 0x74, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x6e, 0x65, 0x74,
	0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x54, 0x72, 0x61, 0x66, 0x66, 0x69,
	0x63, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x4d, 0x54, 0x4c, 0x53,
	0x52, 0x04, 0x6d, 0x74, 0x6c, 0x73, 0x12, 0x78, 0x0a, 0x1c, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x5f, 0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x36, 0x2e, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x69, 0x6e, 0x67, 0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65,
	0x72, 0x76, 0x65, 0x72, 0x52, 0x1a, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x73,
	0x12, 0x44, 0x0a, 0x05, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2e, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e,
	0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x74,
	0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x52,
	0x05, 0x69, 0x73, 0x74, 0x69, 0x6f, 0x1a, 0xd8, 0x04, 0x0a, 0x05, 0x49, 0x73, 0x74, 0x69, 0x6f,
	0x12, 0x87, 0x01, 0x0a, 0x19, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x5f, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x5f, 0x64, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x4b, 0x2e, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e,
	0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69,
	0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x49,
	0x73, 0x74, 0x69, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x17, 0x69, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x1a, 0x91, 0x01, 0x0a, 0x1c, 0x49,
	0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x74,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x5b, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x45, 0x2e, 0x73,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f,
	0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e,
	0x67, 0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x49, 0x6e, 0x67,
	0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0xb0,
	0x02, 0x0a, 0x16, 0x49, 0x6e, 0x67, 0x72, 0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61,
	0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x98, 0x01, 0x0a, 0x17, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x5f, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x60, 0x2e, 0x73, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67, 0x6c, 0x6f, 0x6f,
	0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x49, 0x73, 0x74, 0x69, 0x6f, 0x2e, 0x49, 0x6e, 0x67, 0x72,
	0x65, 0x73, 0x73, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x44, 0x65, 0x74, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x2e, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f,
	0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x15, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x73, 0x12, 0x31, 0x0a, 0x15, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x5f,
	0x74, 0x6c, 0x73, 0x5f, 0x70, 0x6f, 0x72, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x12, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x54, 0x6c, 0x73, 0x50,
	0x6f, 0x72, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x48, 0x0a, 0x1a, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x57, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x95, 0x01, 0x0a, 0x1a, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67,
	0x45, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x69, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x6e,
	0x73, 0x65, 0x63, 0x75, 0x72, 0x65, 0x12, 0x41, 0x0a, 0x1d, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x5f, 0x6f, 0x6e, 0x5f, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x1a, 0x72,
	0x65, 0x63, 0x6f, 0x6e, 0x6e, 0x65, 0x63, 0x74, 0x4f, 0x6e, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72,
	0x6b, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0e, 0x53, 0x65,
	0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2f, 0x0a, 0x13,
	0x6f, 0x62, 0x73, 0x65, 0x72, 0x76, 0x65, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x12, 0x6f, 0x62, 0x73, 0x65, 0x72,
	0x76, 0x65, 0x64, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x41, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2b, 0x2e, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2e, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x65, 0x52, 0x05, 0x73, 0x74, 0x61, 0x74, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x42, 0x4e, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x73, 0x6f, 0x6c, 0x6f, 0x2d, 0x69, 0x6f, 0x2f, 0x67,
	0x6c, 0x6f, 0x6f, 0x2d, 0x6d, 0x65, 0x73, 0x68, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x73, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x2e, 0x6d, 0x65, 0x73, 0x68, 0x2e, 0x67,
	0x6c, 0x6f, 0x6f, 0x2e, 0x73, 0x6f, 0x6c, 0x6f, 0x2e, 0x69, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x32, 0xc0, 0xf5, 0x04, 0x01, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescOnce sync.Once
	file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescData = file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDesc
)

func file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescGZIP() []byte {
	file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescOnce.Do(func() {
		file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescData)
	})
	return file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDescData
}

var file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_goTypes = []interface{}{
	(*SettingsSpec)(nil),               // 0: settings.mesh.gloo.solo.io.SettingsSpec
	(*NetworkingExtensionsServer)(nil), // 1: settings.mesh.gloo.solo.io.NetworkingExtensionsServer
	(*SettingsStatus)(nil),             // 2: settings.mesh.gloo.solo.io.SettingsStatus
	(*SettingsSpec_Istio)(nil),         // 3: settings.mesh.gloo.solo.io.SettingsSpec.Istio
	nil,                                // 4: settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetectorsEntry
	(*SettingsSpec_Istio_IngressGatewayDetector)(nil), // 5: settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetector
	nil,                                     // 6: settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetector.GatewayWorkloadLabelsEntry
	(*v1alpha2.TrafficPolicySpec_MTLS)(nil), // 7: networking.mesh.gloo.solo.io.TrafficPolicySpec.MTLS
	(v1alpha2.ApprovalState)(0),             // 8: networking.mesh.gloo.solo.io.ApprovalState
}
var file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_depIdxs = []int32{
	7, // 0: settings.mesh.gloo.solo.io.SettingsSpec.mtls:type_name -> networking.mesh.gloo.solo.io.TrafficPolicySpec.MTLS
	1, // 1: settings.mesh.gloo.solo.io.SettingsSpec.networking_extension_servers:type_name -> settings.mesh.gloo.solo.io.NetworkingExtensionsServer
	3, // 2: settings.mesh.gloo.solo.io.SettingsSpec.istio:type_name -> settings.mesh.gloo.solo.io.SettingsSpec.Istio
	8, // 3: settings.mesh.gloo.solo.io.SettingsStatus.state:type_name -> networking.mesh.gloo.solo.io.ApprovalState
	4, // 4: settings.mesh.gloo.solo.io.SettingsSpec.Istio.ingress_gateway_detectors:type_name -> settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetectorsEntry
	5, // 5: settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetectorsEntry.value:type_name -> settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetector
	6, // 6: settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetector.gateway_workload_labels:type_name -> settings.mesh.gloo.solo.io.SettingsSpec.Istio.IngressGatewayDetector.GatewayWorkloadLabelsEntry
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_init() }
func file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_init() {
	if File_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsSpec); i {
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
		file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NetworkingExtensionsServer); i {
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
		file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsStatus); i {
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
		file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsSpec_Istio); i {
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
		file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SettingsSpec_Istio_IngressGatewayDetector); i {
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
			RawDescriptor: file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_goTypes,
		DependencyIndexes: file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_depIdxs,
		MessageInfos:      file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_msgTypes,
	}.Build()
	File_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto = out.File
	file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_rawDesc = nil
	file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_goTypes = nil
	file_github_com_solo_io_gloo_mesh_api_settings_v1alpha2_settings_proto_depIdxs = nil
}
