// Copyright (c) 2021 Banzai Cloud Zrt. All Rights Reserved.
// nolint: maligned
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ImagePullSecretSpec defines the desired state of ImagePullSecret
type ImagePullSecretSpec struct {
	// Target specifies what should be the name of the secret created in a
	// given namespace
	Target TargetConfig `json:"target"`

	// Registry contains the details of the secret to be created in each namespace
	Registry RegistryConfig `json:"registry"`
}

type NamespaceSelectorConfiguration struct {
	ObjectSelectorConfiguration `json:",inline"`
	// Namespaces specifies additional namespaces by name to generate the secret into
	Names []string `json:"names,omitempty"`
}

type ObjectSelectorConfiguration struct {
	// Labels specify the conditions, which are matched against the namespaces labels
	// to decide if this ImagePullSecret should be applied to the given namespace, if multiple
	// selectors are specified if one is matches the secret will be managed (OR)
	Labels []metav1.LabelSelector `json:"labels,omitempty"`
	// Selectors specify the conditions, which are matched against the namespaces labels
	// to decide if this ImagePullSecret should be applied to the given namespace, if multiple
	// selectors are specified if one is matches the secret will be managed (OR)
	Annotations []AnnotationSelector `json:"annotations,omitempty"`
}

type AnnotationSelector struct {
	MatchAnnotations map[string]string                 `json:"matchAnnotations,omitempty"`
	MatchExpressions []metav1.LabelSelectorRequirement `json:"matchExpressions,omitempty"`
}

// TargetConfig describes the secret to be created and the selectors required to determine which namespaces should
// contain this secret
type TargetConfig struct {
	Secret TargetSecretConfig `json:"secret"`
	// Namespaces specify conditions on the namespaces that should have the TargetSecret generated
	Namespaces NamespaceSelectorConfiguration `json:"namespaces,omitempty"`
	// Pods specify the conditions, which are matched against the pods in each namespace
	// to decide if this ImagePullSecret should be applied to the given pod's namespace, if multiple
	// selectors are specified if one is matches the secret will be managed (OR)
	NamespacesWithPods ObjectSelectorConfiguration `json:"namespacesWithPods,omitempty"`
}

// TargetSecretConfig describes the properties of the secrets created in each selected namespace
type TargetSecretConfig struct {
	// Name specifies the name of the secret object inside all the selected namespace
	Name string `json:"name"`
	// Labels specifies additional labels to be put on the Secret object
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations specifies additional annotations to be put on the Secret object
	Annotations map[string]string `json:"annotations,omitempty"`
}

// RegistryConfig specifies what secret to be used as the basis of the pull secets
type RegistryConfig struct {
	// Credentials specifies which secret to be used as the source for docker login credentials
	Credentials NamespacedName `json:"credentials"`
}

type ReconciliationStatus string

const (
	ReconciliationReady  = "Ready"
	ReconciliationFailed = "Failed"
)

// ImagePullSecretStatus defines the observed state of ImagePullSecret
type ImagePullSecretStatus struct {
	Status ReconciliationStatus `json:"status,omitempty"`
}

type NamespacedName struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +genclient:nonNamespaced

// ImagePullSecret is the Schema for the imagepullsecrets API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=imagepullsecrets,shortName=imps,scope=Cluster
// +kubebuilder:printcolumn:name="State",type="string",JSONPath=".status.status",description="Represents if the object has been successfully reconciled",priority=0,format="byte"
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
