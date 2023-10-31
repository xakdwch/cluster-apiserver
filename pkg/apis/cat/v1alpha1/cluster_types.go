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
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Cluster `json:"items"`
}

// ClusterSpec defines the desired state of Cluster
type ClusterSpec struct {
	//集群唯一标识,uuid
	Id string `json:"id,omitempty"`
	// k8s版本号
	Version string `json:"version,omitempty"`
	// 供应商,默认裸金属
	Provider string `json:"provider,omitempty"`
	// 容器网络
	NetworkType string `json:"networkType,omitempty"`
	// 支持docker和containerd
	RuntimeType string `json:"runtimeType,omitempty"`
	// 支持systemd和cgroupfs
	CgroupDriver string `json:"cgroupDriver,omitempty"`
	// docker的数据目录
	DockerStorageDir string `json:"dockerStorageDir,omitempty"`
	// containerd的数据目录
	ContainerdStorageDir string `json:"containerdStorageDir,omitempty"`
	// flannel网络模式,host-gw或vxlan
	FlannelBackend string `json:"flannelBackend,omitempty"`
	// calico网络模式,bgp或ipip
	CalicoIpv4poolIpip string `json:"calicoIpv4PoolIpip,omitempty"`
	// pod子网
	KubePodSubnet string `json:"kubePodSubnet,omitempty"`
	// service子网
	KubeServiceSubnet string `json:"kubeServiceSubnet,omitempty"`
	// LB模式下apiserver的虚IP
	LbKubeApiserverIp string `json:"lbKubeApiserverIp,omitempty"`
	// 最大pod数量
	KubeMaxPods int `json:"kubeMaxPods,omitempty"`
	// kube-proxy模式,iptables或ipvs
	KubeProxyMode string `json:"kubeProxyMode,omitempty"`
	// ingress类型,nginx-ingress或traefik-ingress
	IngressControllerType string `json:"ingressControllerType,omitempty"`
	// 架构,amd64或arm64
	Architectures string `json:"architectures,omitempty"`
	// 是否开启kubernetes日志审计
	KubernetesAudit string `json:"kubernetesAudit,omitempty"`
	// 容器子网
	DockerSubnet string `json:"dockerSubnet,omitempty"`
	// 多网卡环境需指定网卡名称,单网卡可不填
	NetworkInterface string `json:"networkInterface,omitempty"`
	// 是否安装GPU套件
	SupportGpu string `json:"supportGpu,omitempty"`
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
	// 集群状态
	Phase string `json:"phase,omitempty"`
	//// 结果信息
	//Message string `json:"message,omitempty"`
	// 集群安装步骤结果信息
	Conditions []ClusterCondition `json:"conditions,omitempty"`
	// 集群apiserver地址列表
	Apiserver []string `json:"apiserver,omitempty"`
	// 创建时间
	CreateAt *metav1.Time `json:"createAt,omitempty"`
	// 更新时间
	UpdateAt *metav1.Time `json:"updateAt,omitempty"`
	//// 集群高可用部署时的虚IP和端口等信息
	//VirtualAddr []ClusterAddress `json:"virtualAddr,omitempty"`
	//// 访问集群的认证信息
	//Credential ClusterCredential `json:"credential,omitempty"`
	// 节点的基本信息,比如内核版本等
	//Machines []ClusterMachine `json:"machines,omitempty"`
}

type ClusterCondition struct {
	// 集群安装对应的安装步骤
	Type string `json:"type,omitempty"`
	// 安装步骤的结果状态
	Status string `json:"status,omitempty"`
	// 安装步骤的成功或失败原因
	Reason string `json:"reason,omitempty"`
	// 安装步骤执行结果的描述信息
	Message string `json:"message,omitempty"`
	// 最后状态检测时间
	LastProbeTime metav1.Time `json:"lastProbeTime,omitempty"`
	// 最后状态转变时间
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`
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
