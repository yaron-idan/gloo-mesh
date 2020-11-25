// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/settings/v1alpha2/settings.proto

package v1alpha2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	v1alpha2 "github.com/solo-io/gloo-mesh/pkg/api/networking.mesh.gloo.solo.io/v1alpha2"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Configure global settings and defaults.
type SettingsSpec struct {
	// Configure default mTLS settings for TrafficTargets (MTLS declared in TrafficPolicies take precedence)
	Mtls *v1alpha2.TrafficPolicySpec_MTLS `protobuf:"bytes,1,opt,name=mtls,proto3" json:"mtls,omitempty"`
	// Configure Gloo Mesh networking to communicate with one or more external gRPC NetworkingExtensions servers.
	// Updates will be applied by the servers in the order they are listed (servers towards the end of the list take precedence).
	// Note: Extension Servers have full write access to the output objects written by Gloo Mesh.
	NetworkingExtensionServers []*NetworkingExtensionsServer `protobuf:"bytes,2,rep,name=networking_extension_servers,json=networkingExtensionServers,proto3" json:"networking_extension_servers,omitempty"`
	// Istio-specific discovery settings
	Istio                *SettingsSpec_Istio `protobuf:"bytes,3,opt,name=istio,proto3" json:"istio,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *SettingsSpec) Reset()         { *m = SettingsSpec{} }
func (m *SettingsSpec) String() string { return proto.CompactTextString(m) }
func (*SettingsSpec) ProtoMessage()    {}
func (*SettingsSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb0461e7c1b076, []int{0}
}
func (m *SettingsSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsSpec.Unmarshal(m, b)
}
func (m *SettingsSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsSpec.Marshal(b, m, deterministic)
}
func (m *SettingsSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsSpec.Merge(m, src)
}
func (m *SettingsSpec) XXX_Size() int {
	return xxx_messageInfo_SettingsSpec.Size(m)
}
func (m *SettingsSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsSpec.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsSpec proto.InternalMessageInfo

func (m *SettingsSpec) GetMtls() *v1alpha2.TrafficPolicySpec_MTLS {
	if m != nil {
		return m.Mtls
	}
	return nil
}

func (m *SettingsSpec) GetNetworkingExtensionServers() []*NetworkingExtensionsServer {
	if m != nil {
		return m.NetworkingExtensionServers
	}
	return nil
}

func (m *SettingsSpec) GetIstio() *SettingsSpec_Istio {
	if m != nil {
		return m.Istio
	}
	return nil
}

type SettingsSpec_Istio struct {
	// The workload labels used during discovery to detect ingress gateways for a mesh.
	// If not specified, will default to `{"istio": "ingressgateway"}`.
	// To override the labels for a specific cluster, use `overrideWorkloadLabels`.
	GatewayWorkloadLabels map[string]string `protobuf:"bytes,1,rep,name=gateway_workload_labels,json=gatewayWorkloadLabels,proto3" json:"gateway_workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// Override the gateway workload labels on a per-cluster basis.
	// The key to the map is a k8s cluster name, and the value is a labels map.
	// If an entry is found for a given cluster, it will be used, otherwise we will fall back to
	// `gatewayWorkloadLabels`.
	OverrideWorkloadLabels map[string]*SettingsSpec_Istio_WorkloadLabel `protobuf:"bytes,2,rep,name=override_workload_labels,json=overrideWorkloadLabels,proto3" json:"override_workload_labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// The name of the TLS port used to detect ingress gateways. Services must have a port with this name
	// in order to be recognized as an ingress gateway during discovery.
	// If not specified, will default to `tls`.
	GatewayTlsPortName   string   `protobuf:"bytes,3,opt,name=gateway_tls_port_name,json=gatewayTlsPortName,proto3" json:"gateway_tls_port_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingsSpec_Istio) Reset()         { *m = SettingsSpec_Istio{} }
func (m *SettingsSpec_Istio) String() string { return proto.CompactTextString(m) }
func (*SettingsSpec_Istio) ProtoMessage()    {}
func (*SettingsSpec_Istio) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb0461e7c1b076, []int{0, 0}
}
func (m *SettingsSpec_Istio) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsSpec_Istio.Unmarshal(m, b)
}
func (m *SettingsSpec_Istio) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsSpec_Istio.Marshal(b, m, deterministic)
}
func (m *SettingsSpec_Istio) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsSpec_Istio.Merge(m, src)
}
func (m *SettingsSpec_Istio) XXX_Size() int {
	return xxx_messageInfo_SettingsSpec_Istio.Size(m)
}
func (m *SettingsSpec_Istio) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsSpec_Istio.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsSpec_Istio proto.InternalMessageInfo

func (m *SettingsSpec_Istio) GetGatewayWorkloadLabels() map[string]string {
	if m != nil {
		return m.GatewayWorkloadLabels
	}
	return nil
}

func (m *SettingsSpec_Istio) GetOverrideWorkloadLabels() map[string]*SettingsSpec_Istio_WorkloadLabel {
	if m != nil {
		return m.OverrideWorkloadLabels
	}
	return nil
}

func (m *SettingsSpec_Istio) GetGatewayTlsPortName() string {
	if m != nil {
		return m.GatewayTlsPortName
	}
	return ""
}

// Wrapper for a set of labels.
type SettingsSpec_Istio_WorkloadLabel struct {
	Labels               map[string]string `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *SettingsSpec_Istio_WorkloadLabel) Reset()         { *m = SettingsSpec_Istio_WorkloadLabel{} }
func (m *SettingsSpec_Istio_WorkloadLabel) String() string { return proto.CompactTextString(m) }
func (*SettingsSpec_Istio_WorkloadLabel) ProtoMessage()    {}
func (*SettingsSpec_Istio_WorkloadLabel) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb0461e7c1b076, []int{0, 0, 2}
}
func (m *SettingsSpec_Istio_WorkloadLabel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsSpec_Istio_WorkloadLabel.Unmarshal(m, b)
}
func (m *SettingsSpec_Istio_WorkloadLabel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsSpec_Istio_WorkloadLabel.Marshal(b, m, deterministic)
}
func (m *SettingsSpec_Istio_WorkloadLabel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsSpec_Istio_WorkloadLabel.Merge(m, src)
}
func (m *SettingsSpec_Istio_WorkloadLabel) XXX_Size() int {
	return xxx_messageInfo_SettingsSpec_Istio_WorkloadLabel.Size(m)
}
func (m *SettingsSpec_Istio_WorkloadLabel) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsSpec_Istio_WorkloadLabel.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsSpec_Istio_WorkloadLabel proto.InternalMessageInfo

func (m *SettingsSpec_Istio_WorkloadLabel) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

// Options for connecting to an external gRPC NetworkingExternsions server
type NetworkingExtensionsServer struct {
	// TCP address of the Networking Extensions Server (including port)
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// Communicate over HTTP rather than HTTPS
	Insecure bool `protobuf:"varint,2,opt,name=insecure,proto3" json:"insecure,omitempty"`
	// Instruct Gloo Mesh to automatically reconnect to the server on network failures
	ReconnectOnNetworkFailures bool     `protobuf:"varint,3,opt,name=reconnect_on_network_failures,json=reconnectOnNetworkFailures,proto3" json:"reconnect_on_network_failures,omitempty"`
	XXX_NoUnkeyedLiteral       struct{} `json:"-"`
	XXX_unrecognized           []byte   `json:"-"`
	XXX_sizecache              int32    `json:"-"`
}

func (m *NetworkingExtensionsServer) Reset()         { *m = NetworkingExtensionsServer{} }
func (m *NetworkingExtensionsServer) String() string { return proto.CompactTextString(m) }
func (*NetworkingExtensionsServer) ProtoMessage()    {}
func (*NetworkingExtensionsServer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb0461e7c1b076, []int{1}
}
func (m *NetworkingExtensionsServer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkingExtensionsServer.Unmarshal(m, b)
}
func (m *NetworkingExtensionsServer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkingExtensionsServer.Marshal(b, m, deterministic)
}
func (m *NetworkingExtensionsServer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkingExtensionsServer.Merge(m, src)
}
func (m *NetworkingExtensionsServer) XXX_Size() int {
	return xxx_messageInfo_NetworkingExtensionsServer.Size(m)
}
func (m *NetworkingExtensionsServer) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkingExtensionsServer.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkingExtensionsServer proto.InternalMessageInfo

func (m *NetworkingExtensionsServer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *NetworkingExtensionsServer) GetInsecure() bool {
	if m != nil {
		return m.Insecure
	}
	return false
}

func (m *NetworkingExtensionsServer) GetReconnectOnNetworkFailures() bool {
	if m != nil {
		return m.ReconnectOnNetworkFailures
	}
	return false
}

type SettingsStatus struct {
	// The most recent generation observed in the the Settings metadata.
	// If the observedGeneration does not match generation, the controller has not processed the most
	// recent version of this resource.
	ObservedGeneration int64 `protobuf:"varint,1,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	// The state of the overall resource.
	// It will only show accepted if no processing errors encountered.
	State v1alpha2.ApprovalState `protobuf:"varint,2,opt,name=state,proto3,enum=networking.mesh.gloo.solo.io.ApprovalState" json:"state,omitempty"`
	// Any errors encountered while processing Settings object.
	Errors               []string `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SettingsStatus) Reset()         { *m = SettingsStatus{} }
func (m *SettingsStatus) String() string { return proto.CompactTextString(m) }
func (*SettingsStatus) ProtoMessage()    {}
func (*SettingsStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_8dcb0461e7c1b076, []int{2}
}
func (m *SettingsStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SettingsStatus.Unmarshal(m, b)
}
func (m *SettingsStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SettingsStatus.Marshal(b, m, deterministic)
}
func (m *SettingsStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SettingsStatus.Merge(m, src)
}
func (m *SettingsStatus) XXX_Size() int {
	return xxx_messageInfo_SettingsStatus.Size(m)
}
func (m *SettingsStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_SettingsStatus.DiscardUnknown(m)
}

var xxx_messageInfo_SettingsStatus proto.InternalMessageInfo

func (m *SettingsStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *SettingsStatus) GetState() v1alpha2.ApprovalState {
	if m != nil {
		return m.State
	}
	return v1alpha2.ApprovalState_PENDING
}

func (m *SettingsStatus) GetErrors() []string {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*SettingsSpec)(nil), "settings.mesh.gloo.solo.io.SettingsSpec")
	proto.RegisterType((*SettingsSpec_Istio)(nil), "settings.mesh.gloo.solo.io.SettingsSpec.Istio")
	proto.RegisterMapType((map[string]string)(nil), "settings.mesh.gloo.solo.io.SettingsSpec.Istio.GatewayWorkloadLabelsEntry")
	proto.RegisterMapType((map[string]*SettingsSpec_Istio_WorkloadLabel)(nil), "settings.mesh.gloo.solo.io.SettingsSpec.Istio.OverrideWorkloadLabelsEntry")
	proto.RegisterType((*SettingsSpec_Istio_WorkloadLabel)(nil), "settings.mesh.gloo.solo.io.SettingsSpec.Istio.WorkloadLabel")
	proto.RegisterMapType((map[string]string)(nil), "settings.mesh.gloo.solo.io.SettingsSpec.Istio.WorkloadLabel.LabelsEntry")
	proto.RegisterType((*NetworkingExtensionsServer)(nil), "settings.mesh.gloo.solo.io.NetworkingExtensionsServer")
	proto.RegisterType((*SettingsStatus)(nil), "settings.mesh.gloo.solo.io.SettingsStatus")
}

func init() {
	proto.RegisterFile("github.com/solo-io/gloo-mesh/api/settings/v1alpha2/settings.proto", fileDescriptor_8dcb0461e7c1b076)
}

var fileDescriptor_8dcb0461e7c1b076 = []byte{
	// 653 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xcd, 0x6e, 0x13, 0x3f,
	0x10, 0xd7, 0x36, 0xff, 0xf4, 0xc3, 0xf9, 0x53, 0x21, 0x53, 0x4a, 0xb4, 0x7c, 0xa8, 0xea, 0xa9,
	0x12, 0xaa, 0xa3, 0x06, 0x54, 0x15, 0xc4, 0x25, 0x88, 0xd2, 0x14, 0x95, 0xb6, 0xda, 0x54, 0x42,
	0xe2, 0xb2, 0x38, 0x9b, 0xe9, 0xd6, 0xaa, 0xe3, 0x59, 0xd9, 0x4e, 0xda, 0x1c, 0x91, 0x90, 0x78,
	0x02, 0x6e, 0x3c, 0x00, 0x07, 0x1e, 0x84, 0xe7, 0x80, 0x17, 0x41, 0xeb, 0xdd, 0x6c, 0x13, 0x25,
	0x0d, 0x2a, 0xdc, 0x76, 0x3c, 0xfe, 0x7d, 0xcc, 0xac, 0x3d, 0x26, 0x8d, 0x58, 0xd8, 0xb3, 0x5e,
	0x9b, 0x45, 0xd8, 0xad, 0x19, 0x94, 0xb8, 0x29, 0xb0, 0x16, 0x4b, 0xc4, 0xcd, 0x2e, 0x98, 0xb3,
	0x1a, 0x4f, 0x44, 0xcd, 0x80, 0xb5, 0x42, 0xc5, 0xa6, 0xd6, 0xdf, 0xe2, 0x32, 0x39, 0xe3, 0xf5,
	0x62, 0x85, 0x25, 0x1a, 0x2d, 0x52, 0xbf, 0x88, 0x53, 0x0c, 0x4b, 0xd1, 0x2c, 0xa5, 0x62, 0x02,
	0xfd, 0x95, 0x18, 0x63, 0x74, 0xdb, 0x6a, 0xe9, 0x57, 0x86, 0xf0, 0x77, 0xc6, 0x15, 0x14, 0xd8,
	0x0b, 0xd4, 0xe7, 0x42, 0xc5, 0x57, 0x1a, 0x7d, 0x2e, 0x45, 0x87, 0x5b, 0x81, 0x2a, 0x34, 0x96,
	0x5b, 0xc8, 0x91, 0xdb, 0x7f, 0x46, 0x5a, 0xcd, 0x4f, 0x4f, 0x45, 0x14, 0x26, 0x28, 0x45, 0x34,
	0xc8, 0x70, 0xeb, 0xbf, 0x16, 0xc8, 0xff, 0xad, 0xdc, 0x66, 0x2b, 0x81, 0x88, 0x36, 0xc9, 0x7f,
	0x5d, 0x2b, 0x4d, 0xd5, 0x5b, 0xf3, 0x36, 0x2a, 0xf5, 0xa7, 0xec, 0x8a, 0x69, 0xb2, 0x0a, 0x76,
	0x92, 0x51, 0x1e, 0x3b, 0xc6, 0x14, 0xce, 0xde, 0x9e, 0x1c, 0xb4, 0x02, 0xc7, 0x40, 0x2f, 0xc9,
	0x83, 0x2b, 0x70, 0x08, 0x97, 0x16, 0x94, 0x71, 0xb6, 0x41, 0xf7, 0x41, 0x9b, 0xea, 0xdc, 0x5a,
	0x69, 0xa3, 0x52, 0xdf, 0x66, 0xd7, 0x77, 0x89, 0x1d, 0x16, 0xf8, 0xdd, 0x21, 0xdc, 0xb4, 0x1c,
	0x3c, 0xf0, 0xd5, 0x64, 0x2e, 0x4b, 0x19, 0xfa, 0x8a, 0x94, 0x85, 0xb1, 0x02, 0xab, 0x25, 0x57,
	0x04, 0x9b, 0x25, 0x31, 0x5a, 0x3c, 0xdb, 0x4f, 0x51, 0x41, 0x06, 0xf6, 0x7f, 0x94, 0x49, 0xd9,
	0x2d, 0xd0, 0x8f, 0x1e, 0xb9, 0x17, 0x73, 0x0b, 0x17, 0x7c, 0x10, 0xa6, 0x9a, 0x12, 0x79, 0x27,
	0x94, 0xbc, 0x0d, 0xae, 0x4f, 0x69, 0x15, 0xfb, 0x37, 0x93, 0x60, 0x7b, 0x19, 0xdb, 0xbb, 0x9c,
	0xec, 0xc0, 0x71, 0xed, 0x2a, 0xab, 0x07, 0xc1, 0xdd, 0x78, 0x5a, 0x8e, 0x7e, 0xf2, 0x48, 0x15,
	0xfb, 0xa0, 0xb5, 0xe8, 0xc0, 0x84, 0x89, 0xac, 0x95, 0x6f, 0x6e, 0x68, 0xe2, 0x28, 0xa7, 0x9b,
	0xe6, 0x62, 0x15, 0xa7, 0x26, 0xe9, 0x16, 0x19, 0xfa, 0x0b, 0xad, 0x34, 0x61, 0x82, 0xda, 0x86,
	0x8a, 0x77, 0xc1, 0xb5, 0x7a, 0x29, 0xa0, 0x79, 0xf2, 0x44, 0x9a, 0x63, 0xd4, 0xf6, 0x90, 0x77,
	0xc1, 0x6f, 0x12, 0xff, 0xfa, 0x72, 0xe9, 0x6d, 0x52, 0x3a, 0x87, 0x81, 0x3b, 0x6e, 0x4b, 0x41,
	0xfa, 0x49, 0x57, 0x48, 0xb9, 0xcf, 0x65, 0x0f, 0xaa, 0x73, 0x6e, 0x2d, 0x0b, 0x9e, 0xcf, 0xed,
	0x78, 0xfe, 0x67, 0x8f, 0xdc, 0x9f, 0x61, 0x7a, 0x0a, 0x57, 0x30, 0xca, 0x55, 0xa9, 0xbf, 0xb8,
	0x61, 0x87, 0xc6, 0x44, 0x46, 0x9d, 0x7c, 0xf7, 0xc8, 0xad, 0xb1, 0x24, 0xfd, 0x40, 0xe6, 0xc7,
	0x4e, 0x44, 0xf3, 0x5f, 0xa4, 0xd8, 0xe8, 0xaf, 0xc8, 0x79, 0xfd, 0x67, 0xa4, 0xf2, 0x97, 0x8d,
	0x5b, 0xff, 0xe2, 0x11, 0xff, 0xfa, 0xbb, 0x44, 0xab, 0x64, 0x81, 0x77, 0x3a, 0x1a, 0x8c, 0xc9,
	0xe9, 0x86, 0x21, 0xf5, 0xc9, 0xa2, 0x50, 0x06, 0xa2, 0x9e, 0xce, 0x58, 0x17, 0x83, 0x22, 0xa6,
	0x0d, 0xf2, 0x50, 0x43, 0x84, 0x4a, 0x41, 0x64, 0x43, 0x54, 0x61, 0x7e, 0x21, 0xc3, 0x53, 0x2e,
	0x64, 0x4f, 0x83, 0x71, 0x47, 0x62, 0x31, 0xf0, 0x8b, 0x4d, 0x47, 0x2a, 0xf7, 0xf0, 0x3a, 0xdf,
	0xb1, 0xfe, 0xd5, 0x23, 0xcb, 0x45, 0x2f, 0x2c, 0xb7, 0x3d, 0x43, 0x6b, 0xe4, 0x0e, 0xb6, 0xdd,
	0x88, 0xe8, 0x84, 0x31, 0x28, 0xd0, 0x6e, 0xd6, 0x39, 0x5f, 0xa5, 0x80, 0x0e, 0x53, 0x7b, 0x45,
	0x86, 0x36, 0x48, 0xd9, 0x0d, 0x42, 0xe7, 0x6f, 0xb9, 0xfe, 0x78, 0xf6, 0xc4, 0x6a, 0x24, 0x89,
	0xc6, 0x3e, 0x97, 0xa9, 0x1a, 0x04, 0x19, 0x92, 0xae, 0x92, 0x79, 0xd0, 0x1a, 0x75, 0x6a, 0xb9,
	0xb4, 0xb1, 0x14, 0xe4, 0xd1, 0xcb, 0xc3, 0x6f, 0x3f, 0x1f, 0x79, 0xef, 0x9b, 0x33, 0x5f, 0x82,
	0xe4, 0x3c, 0x1e, 0x7b, 0x0d, 0x26, 0x35, 0x8b, 0x09, 0xdc, 0x9e, 0x77, 0x33, 0xf7, 0xc9, 0xef,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x63, 0x98, 0x31, 0x2a, 0x5c, 0x06, 0x00, 0x00,
}

func (this *SettingsSpec) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SettingsSpec)
	if !ok {
		that2, ok := that.(SettingsSpec)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Mtls.Equal(that1.Mtls) {
		return false
	}
	if len(this.NetworkingExtensionServers) != len(that1.NetworkingExtensionServers) {
		return false
	}
	for i := range this.NetworkingExtensionServers {
		if !this.NetworkingExtensionServers[i].Equal(that1.NetworkingExtensionServers[i]) {
			return false
		}
	}
	if !this.Istio.Equal(that1.Istio) {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SettingsSpec_Istio) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SettingsSpec_Istio)
	if !ok {
		that2, ok := that.(SettingsSpec_Istio)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.GatewayWorkloadLabels) != len(that1.GatewayWorkloadLabels) {
		return false
	}
	for i := range this.GatewayWorkloadLabels {
		if this.GatewayWorkloadLabels[i] != that1.GatewayWorkloadLabels[i] {
			return false
		}
	}
	if len(this.OverrideWorkloadLabels) != len(that1.OverrideWorkloadLabels) {
		return false
	}
	for i := range this.OverrideWorkloadLabels {
		if !this.OverrideWorkloadLabels[i].Equal(that1.OverrideWorkloadLabels[i]) {
			return false
		}
	}
	if this.GatewayTlsPortName != that1.GatewayTlsPortName {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SettingsSpec_Istio_WorkloadLabel) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SettingsSpec_Istio_WorkloadLabel)
	if !ok {
		that2, ok := that.(SettingsSpec_Istio_WorkloadLabel)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *NetworkingExtensionsServer) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*NetworkingExtensionsServer)
	if !ok {
		that2, ok := that.(NetworkingExtensionsServer)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.Insecure != that1.Insecure {
		return false
	}
	if this.ReconnectOnNetworkFailures != that1.ReconnectOnNetworkFailures {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *SettingsStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*SettingsStatus)
	if !ok {
		that2, ok := that.(SettingsStatus)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if this.State != that1.State {
		return false
	}
	if len(this.Errors) != len(that1.Errors) {
		return false
	}
	for i := range this.Errors {
		if this.Errors[i] != that1.Errors[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
