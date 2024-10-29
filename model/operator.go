package model

import (
	corev1 "k8s.io/api/core/v1"
)

type AppDeploymentSpec struct {
	Image        string                      `json:"image"`
	Replicas     int32                       `json:"replicas"`
	BearerToken  string                      `json:"bearerToken,omitempty"`
	RevisionId   string                      `json:"revisionId,omitempty"`
	Ports        []corev1.ContainerPort      `json:"ports,omitempty"`
	Resources    corev1.ResourceRequirements `json:"resources,omitempty"`
	VolumeMounts []corev1.VolumeMount        `json:"volumeMounts,omitempty"`
}
