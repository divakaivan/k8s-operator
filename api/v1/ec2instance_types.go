/*
Copyright 2026.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// EC2InstanceSpec defines the desired state of EC2Instance
type EC2InstanceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// The following markers will use OpenAPI v3 schema to validate the value
	// More info: https://book.kubebuilder.io/reference/markers/crd-validation.html

	InstanceType      string            `json:"instanceType"`
	AMIId             string            `json:"amiId"`
	Region            string            `json:"region"`
	AvailabilityZone  string            `json:"availabilityZone,omitempty"`
	KeyPair           string            `json:"keyPair,omitempty"`
	SecurityGroups    []string          `json:"securityGroups,omitempty"`
	Subnet            string            `json:"subnet,omitempty"`
	UserData          string            `json:"userData,omitempty"`
	Tags              map[string]string `json:"tags,omitempty"`
	Storage           StorageConfig     `json:"storage,omitempty"`
	AssociatePublicIP bool              `json:"associatePublicIP,omitempty"`
}

type StorageConfig struct {
	RootVolume       VolumeConfig   `json:"rootVolume,omitempty"`
	AdditionalVolume []VolumeConfig `json:"additionalVolumes,omitempty"`
}

type VolumeConfig struct {
	Size       int32  `json:"size"`
	Type       string `json:"type,omitempty"`
	DeviceName string `json:"deviceName,omitempty"`
	Encrypted  bool   `json:"encrypted,omitempty"`
}

type Condition struct {
	Type               string      `json:"type"`
	Status             string      `json:"status"`
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	Reason             string      `json:"reason,omitempty"`
	Message            string      `json:"message,omitempty"`
}

// EC2InstanceStatus defines the observed state of EC2Instance.
type EC2InstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// For Kubernetes API conventions, see:
	// https://github.com/kubernetes/community/blob/master/contributors/devel/sig-architecture/api-conventions.md#typical-status-properties

	// conditions represent the current state of the EC2Instance resource.
	// Each condition has a unique type and reflects the status of a specific aspect of the resource.
	//
	// Standard condition types include:
	// - "Available": the resource is fully functional
	// - "Progressing": the resource is being created or updated
	// - "Degraded": the resource failed to reach or maintain its desired state
	//
	// The status of each condition is one of True, False, or Unknown.
	InstanceID string       `json:"instanceId,omitempty"`
	State      string       `json:"state,omitempty"`
	PublicIP   string       `json:"publicIP,omitempty"`
	PrivateIP  string       `json:"privateIP,omitempty"`
	PublicDNS  string       `json:"publicDNS,omitempty"`
	PrivateDNS string       `json:"privateDNS,omitempty"`
	LaunchTime *metav1.Time `json:"launchTime,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// EC2Instance is the Schema for the ec2instances API
type EC2Instance struct {
	metav1.TypeMeta `json:",inline"`

	// metadata is a standard object metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// spec defines the desired state of EC2Instance
	// +required
	Spec EC2InstanceSpec `json:"spec"`

	// status defines the observed state of EC2Instance
	// +optional
	Status EC2InstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// EC2InstanceList contains a list of EC2Instance
type EC2InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []EC2Instance `json:"items"`
}

type CreatedInstanceInfo struct {
	InstanceID string `json:"instanceId"`
	PublicIP   string `json:"publicIP"`
	PrivateIP  string `json:"privateIP"`
	PublicDNS  string `json:"publicDNS"`
	PrivateDNS string `json:"privateDNS"`
	State      string `json:"state"`
}

func init() {
	SchemeBuilder.Register(&EC2Instance{}, &EC2InstanceList{})
}
