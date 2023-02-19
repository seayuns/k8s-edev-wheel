package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

//VirtualMachine 根据 CRD 定义 的 结构体
type VirtualMachine struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VMSpec `json:"spec"`
}

// +k8s:deepcopy-gen=false

type VMSpec struct {
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Image string `json:"image"`
	Memory int `json:"memory"`
	Disk int `json:"disk"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualMachineList 资源列表
type VirtualMachineList struct {
	metav1.TypeMeta `json:",inline"`

	// 标准的 list metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []VirtualMachine `json:"items"`
}
