/*
Copyright 2025.

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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// RedisSentinelSpec defines the desired state of RedisSentinel.
type RedisSentinelSpec struct {
	Sentinel   SentinelSpec   `json:"sentinel"`
	Redis      RedisSpec      `json:"redis"`
	Failover   FailoverSpec   `json:"failover"`
	Networking NetworkingSpec `json:"networking"`
	Monitoring MonitoringSpec `json:"monitoring"`
}

type SentinelSpec struct {
	Replicas  int32                         `json:"replicas"`
	Quorum    int32                         `json:"quorum"`
	Image     string                        `json:"image"`
	Resources *corev1.ResourceRequirements `json:"resources,omitempty"`
}

type RedisPersistenceSpec struct {
	Storage          string `json:"storage"`
	StorageClassName string `json:"storageClassName,omitempty"`
}

type RedisSpec struct {
	Replicas    int32                         `json:"replicas"`
	Version     string                        `json:"version"`
	Image       string                        `json:"image"`
	Resources   *corev1.ResourceRequirements `json:"resources,omitempty"`
	Persistence RedisPersistenceSpec          `json:"persistence"`
}

type ReplicaPriorityRule struct {
	Label  string `json:"label"`
	Weight int    `json:"weight"`
}

type FailoverSpec struct {
	MaxDelaySeconds int                   `json:"maxDelaySeconds"`
	ReplicaPriority []ReplicaPriorityRule `json:"replicaPriority,omitempty"`
}

type NetworkingSpec struct {
	AccessMode      string `json:"accessMode"`
	DNSUpdatePolicy string `json:"dnsUpdatePolicy"`
}

type MonitoringSpec struct {
	Prometheus  bool  `json:"prometheus"`
	MetricsPort int32 `json:"metricsPort"`
}

type RedisSentinelStatus struct {
	Phase    string         `json:"phase,omitempty"`
	Redis    RedisStatus    `json:"redis,omitempty"`
	Sentinel SentinelStatus `json:"sentinel,omitempty"`
	Failover FailoverStatus `json:"failover,omitempty"`

	ObservedGeneration int64              `json:"observedGeneration,omitempty"`
	Conditions         []metav1.Condition `json:"conditions,omitempty"`
	LastUpdated        metav1.Time        `json:"lastUpdated,omitempty"`
}

type RedisStatus struct {
	ReadyReplicas    int32  `json:"readyReplicas"`
	CurrentMaster    string `json:"currentMaster"`
	ReplicationLagOK bool   `json:"replicationLagOK"`
	ConfigHash       string `json:"configHash"`
}

type SentinelStatus struct {
	ReadyReplicas   int32  `json:"readyReplicas"`
	QuorumAchieved  bool   `json:"quorumAchieved"`
	Leader          string `json:"leader,omitempty"`
	MonitoredMaster string `json:"monitoredMaster"`
}

type FailoverStatus struct {
	LastFailoverTime     *metav1.Time `json:"lastFailoverTime,omitempty"`
	LastFailoverReason   string       `json:"lastFailoverReason,omitempty"`
	InProgress           bool         `json:"inProgress"`
	ReplicaPromoted      string       `json:"replicaPromoted,omitempty"`
	FailoverDurationSecs int          `json:"failoverDurationSecs,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// RedisSentinel is the Schema for the redissentinels API.
type RedisSentinel struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RedisSentinelSpec   `json:"spec,omitempty"`
	Status RedisSentinelStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RedisSentinelList contains a list of RedisSentinel.
type RedisSentinelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisSentinel `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RedisSentinel{}, &RedisSentinelList{})
}
