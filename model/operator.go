package model

import (
	corev1 "k8s.io/api/core/v1"
)

type ContainerPort corev1.ContainerPort
type ResourceRequirements corev1.ResourceRequirements
type VolumeMount corev1.VolumeMount

type SyncAppRequest struct {
	AppName      string `json:"app_name"`
	AppNamespace string `json:"app_namespace"`

	Replicas     int32                  `json:"replicas"`
	Ports        []interface{}          `json:"ports"`
	Resources    map[string]interface{} `json:"resources"`
	VolumeMounts map[string]string      `json:"volume_mounts"`
}

type AppDeploymentSpec struct {
	Image        string               `json:"image"`
	Replicas     int32                `json:"replicas"`
	Ports        []ContainerPort      `json:"ports,omitempty"`
	Resources    ResourceRequirements `json:"resources,omitempty"`
	VolumeMounts []VolumeMount        `json:"volumeMounts,omitempty"`
}
