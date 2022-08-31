package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OvhNginxSpec defines the desired state of OvhNginx
type OvhNginxSpec struct {
	// Number of replicas for the Nginx Pods
	ReplicaCount int32 `json:"replicaCount"`
	// Exposed port for the Nginx server
	Port int32 `json:"port"`
}

// OvhNginxStatus defines the observed state of OvhNginx
type OvhNginxStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OvhNginx is the Schema for the ovhnginxes API
type OvhNginx struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OvhNginxSpec   `json:"spec,omitempty"`
	Status OvhNginxStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OvhNginxList contains a list of OvhNginx
type OvhNginxList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OvhNginx `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OvhNginx{}, &OvhNginxList{})
}
