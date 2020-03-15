// Copyright 2018 Imhotep Software LLC.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// PainterSpec defines the desired state of Painter
type PainterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// CHANGE_ME! Define color property and annotation
	// +kubebuilder:validation:<CHANGE_ME!>
	<CHANGE_ME!>
}

// PainterStatus defines the observed state of Painter
type PainterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Painter is the Schema for the painters API
// +k8s:openapi-gen=true
type Painter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PainterSpec   `json:"spec,omitempty"`
	Status PainterStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// PainterList contains a list of Painter
type PainterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Painter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Painter{}, &PainterList{})
}
