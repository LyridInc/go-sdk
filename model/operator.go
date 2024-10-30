package model

import (
	"fmt"
	"strconv"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
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

func ConvertCrdToLyridDefinition(appName string, crd AppDeploymentSpec) AppDefinition {
	volumeDefinitions := []VolumeDefinition{}
	for _, v := range crd.VolumeMounts {
		volumeDefinitions = append(volumeDefinitions, VolumeDefinition{
			Alias:     v.Name,
			MountPath: v.MountPath,
		})
	}

	portDefinitions := []PortDefinition{}
	for _, p := range crd.Ports {
		portDefinitions = append(portDefinitions, PortDefinition{
			Alias: p.Name,
			Port:  int64(p.ContainerPort),
		})
	}

	resources := ResourcesDefinition{
		Requests: ResourceList{
			Cpu:    crd.Resources.Requests.Cpu().String(),
			Memory: crd.Resources.Requests.Memory().String(),
		},
		Limits: ResourceList{
			Cpu:    crd.Resources.Limits.Cpu().String(),
			Memory: crd.Resources.Limits.Memory().String(),
		},
	}

	modules := []ModuleDefinition{
		{
			Name:         "container",
			Language:     "docker",
			Description:  "App module deployment with operator",
			Volumes:      volumeDefinitions,
			Ports:        portDefinitions,
			Resources:    resources,
			LastActivity: time.Now().UTC(),
			LastUpdate:   time.Now().UTC(),
		},
	}
	lyridDefinition := AppDefinition{
		Name:        appName,
		Modules:     modules,
		Description: "App deployment with operator",
		Resources:   []ResourcesDefinition{resources},
		Spec: []SpecDefinition{
			{Replica: fmt.Sprintf("%d", crd.Replicas)},
		},
	}
	return lyridDefinition
}

func ConvertLyridDefinitionToCrd(dockerImage string, d AppDefinition) (*AppDeploymentSpec, error) {
	if len(d.Spec) <= 0 {
		return nil, fmt.Errorf("no spec specified")
	}

	spec := d.Spec[0]
	replica, err := strconv.Atoi(spec.Replica)
	if err != nil {
		return nil, err
	}

	if len(d.Modules) <= 0 {
		return nil, fmt.Errorf("no module specified")
	}

	module := d.Modules[0]

	ports := []ContainerPort{}
	for _, p := range module.Ports {
		ports = append(ports, ContainerPort{
			ContainerPort: int32(p.Port),
			Name:          p.Alias,
		})
	}

	volumeMounts := []VolumeMount{}
	for _, v := range module.Volumes {
		volumeMounts = append(volumeMounts, VolumeMount{
			Name:      v.Alias,
			MountPath: v.MountPath,
		})
	}

	cpuLimit, err := resource.ParseQuantity(module.Resources.Limits.Cpu)
	memoryLimit, err := resource.ParseQuantity(module.Resources.Limits.Memory)
	cpuRequest, err := resource.ParseQuantity(module.Resources.Requests.Cpu)
	memoryRequest, err := resource.ParseQuantity(module.Resources.Requests.Memory)

	resources := ResourceRequirements{
		Limits: corev1.ResourceList{
			corev1.ResourceCPU:    cpuLimit,
			corev1.ResourceMemory: memoryLimit,
		},
		Requests: corev1.ResourceList{
			corev1.ResourceCPU:    cpuRequest,
			corev1.ResourceMemory: memoryRequest,
		},
	}

	crd := &AppDeploymentSpec{
		Image:        dockerImage,
		Replicas:     int32(replica),
		Ports:        ports,
		Resources:    resources,
		VolumeMounts: volumeMounts,
	}
	return crd, nil
}
