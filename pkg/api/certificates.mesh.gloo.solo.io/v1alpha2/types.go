// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// IssuedCertificate is the Schema for the issuedCertificate API
type IssuedCertificate struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IssuedCertificateSpec   `json:"spec,omitempty"`
	Status IssuedCertificateStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (IssuedCertificate) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "certificates.mesh.gloo.solo.io",
		Version: "v1alpha2",
		Kind:    "IssuedCertificate",
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IssuedCertificateList contains a list of IssuedCertificate
type IssuedCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IssuedCertificate `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// CertificateRequest is the Schema for the certificateRequest API
type CertificateRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CertificateRequestSpec   `json:"spec,omitempty"`
	Status CertificateRequestStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (CertificateRequest) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "certificates.mesh.gloo.solo.io",
		Version: "v1alpha2",
		Kind:    "CertificateRequest",
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CertificateRequestList contains a list of CertificateRequest
type CertificateRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CertificateRequest `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// PodBounceDirective is the Schema for the podBounceDirective API
type PodBounceDirective struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PodBounceDirectiveSpec   `json:"spec,omitempty"`
	Status PodBounceDirectiveStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (PodBounceDirective) GVK() schema.GroupVersionKind {
	return schema.GroupVersionKind{
		Group:   "certificates.mesh.gloo.solo.io",
		Version: "v1alpha2",
		Kind:    "PodBounceDirective",
	}
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PodBounceDirectiveList contains a list of PodBounceDirective
type PodBounceDirectiveList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PodBounceDirective `json:"items"`
}

func init() {
	SchemeBuilder.Register(&IssuedCertificate{}, &IssuedCertificateList{})
	SchemeBuilder.Register(&CertificateRequest{}, &CertificateRequestList{})
	SchemeBuilder.Register(&PodBounceDirective{}, &PodBounceDirectiveList{})
}
