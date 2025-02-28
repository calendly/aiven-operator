// Copyright (c) 2022 Aiven, Helsinki, Finland. https://aiven.io/

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProjectVPCSpec defines the desired state of ProjectVPC
type ProjectVPCSpec struct {
	// +kubebuilder:validation:MaxLength=63
	// +kubebuilder:validation:Format="^[a-zA-Z0-9_-]*$"
	// The project the VPC belongs to
	Project string `json:"project"`

	// +kubebuilder:validation:MaxLength=256
	// Cloud the VPC is in
	CloudName string `json:"cloudName"`

	// +kubebuilder:validation:MaxLength=36
	// Network address range used by the VPC like 192.168.0.0/24
	NetworkCidr string `json:"networkCidr"`

	// Authentication reference to Aiven token in a secret
	AuthSecretRef AuthSecretReference `json:"authSecretRef"`
}

// ProjectVPCStatus defines the observed state of ProjectVPC
type ProjectVPCStatus struct {
	// Conditions represent the latest available observations of an ProjectVPC state
	Conditions []metav1.Condition `json:"conditions"`

	// State of VPC
	State string `json:"state"`

	// Project VPC id
	ID string `json:"id"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// ProjectVPC is the Schema for the projectvpcs API
// +kubebuilder:printcolumn:name="Project",type="string",JSONPath=".spec.project"
// +kubebuilder:printcolumn:name="Cloud",type="string",JSONPath=".spec.cloudName"
// +kubebuilder:printcolumn:name="Network CIDR",type="string",JSONPath=".spec.networkCidr"
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.state"
type ProjectVPC struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProjectVPCSpec   `json:"spec,omitempty"`
	Status ProjectVPCStatus `json:"status,omitempty"`
}

func (pvpc ProjectVPC) AuthSecretRef() AuthSecretReference {
	return pvpc.Spec.AuthSecretRef
}

// +kubebuilder:object:root=true

// ProjectVPCList contains a list of ProjectVPC
type ProjectVPCList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProjectVPC `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ProjectVPC{}, &ProjectVPCList{})
}
