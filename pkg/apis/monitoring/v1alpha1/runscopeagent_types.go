package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RunscopeAgentSpec defines the desired state of RunscopeAgent
// +k8s:openapi-gen=true
type RunscopeAgentSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// RunscopeAgentStatus defines the observed state of RunscopeAgent
// +k8s:openapi-gen=true
type RunscopeAgentStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book.kubebuilder.io/beyond_basics/generating_crd.html
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RunscopeAgent is the Schema for the runscopeagents API
// +k8s:openapi-gen=true
type RunscopeAgent struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RunscopeAgentSpec   `json:"spec,omitempty"`
	Status RunscopeAgentStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RunscopeAgentList contains a list of RunscopeAgent
type RunscopeAgentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RunscopeAgent `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RunscopeAgent{}, &RunscopeAgentList{})
}
