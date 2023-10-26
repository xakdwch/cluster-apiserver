
/*
Copyright 2023.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
 	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource"
	"sigs.k8s.io/apiserver-runtime/pkg/builder/resource/resourcestrategy"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Cluster
// +k8s:openapi-gen=true
type Cluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterSpec   `json:"spec,omitempty"`
	Status ClusterStatus `json:"status,omitempty"`
}

// ClusterList
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type ClusterList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Cluster `json:"items"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
}

var _ resource.Object = &Cluster{}
var _ resourcestrategy.Validater = &Cluster{}

func (in *Cluster) GetObjectMeta() *metav1.ObjectMeta {
	return &in.ObjectMeta
}

func (in *Cluster) NamespaceScoped() bool {
	return false
}

func (in *Cluster) New() runtime.Object {
	return &Cluster{}
}

func (in *Cluster) NewList() runtime.Object {
	return &ClusterList{}
}

func (in *Cluster) GetGroupVersionResource() schema.GroupVersionResource {
	return schema.GroupVersionResource{
		Group:    "cat.pets.com",
		Version:  "v1alpha1",
		Resource: "clusters",
	}
}

func (in *Cluster) IsStorageVersion() bool {
	return true
}

func (in *Cluster) Validate(ctx context.Context) field.ErrorList {
	// TODO(user): Modify it, adding your API validation here.
	return nil
}

var _ resource.ObjectList = &ClusterList{}

func (in *ClusterList) GetListMeta() *metav1.ListMeta {
	return &in.ListMeta
}
// ClusterStatus defines the observed state of Cluster
type ClusterStatus struct {
}

func (in ClusterStatus) SubResourceName() string {
	return "status"
}

// Cluster implements ObjectWithStatusSubResource interface.
var _ resource.ObjectWithStatusSubResource = &Cluster{}

func (in *Cluster) GetStatus() resource.StatusSubResource {
	return in.Status
}

// ClusterStatus{} implements StatusSubResource interface.
var _ resource.StatusSubResource = &ClusterStatus{}

func (in ClusterStatus) CopyTo(parent resource.ObjectWithStatusSubResource) {
	parent.(*Cluster).Status = in
}
