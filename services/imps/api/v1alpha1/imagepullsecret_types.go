// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.
// nolint: maligned
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ImagePullSecretSpec defines the desired state of ImagePullSecret
type ImagePullSecretSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// ImagePullSecretStatus defines the observed state of ImagePullSecret
type ImagePullSecretStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ImagePullSecret is the Schema for the imagepullsecrets API
// +k8s:openapi-gen=true
type ImagePullSecret struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ImagePullSecretSpec   `json:"spec,omitempty"`
	Status ImagePullSecretStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ImagePullSecretList contains a list of ImagePullSecret
type ImagePullSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ImagePullSecret `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ImagePullSecret{}, &ImagePullSecretList{})
}
