// Package v1 is a package.
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BasicType is a basic type.
// +k8s:deepcopy-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +protobuf=true
// +k8s:openapi-gen=true
// +genclient
type BasicType struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	// Field is a field.
	Field string `json:"field" protobuf:"bytes,2,opt,name=field"`
}
