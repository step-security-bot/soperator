package controller

import (
	"fmt"
	"path"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"

	"nebius.ai/slurm-operator/internal/consts"
	"nebius.ai/slurm-operator/internal/render/common"
	"nebius.ai/slurm-operator/internal/values"
)

// renderContainerSlurmctld renders Slurm controller [corev1.Container] for slurmctld
func renderContainerSlurmctld(container *values.Container) corev1.Container {
	return corev1.Container{
		Name:            consts.ContainerNameSlurmctld,
		Image:           container.Image,
		ImagePullPolicy: corev1.PullAlways, // TODO use digest and set to corev1.PullIfNotPresent
		Ports: []corev1.ContainerPort{{
			Name:          consts.ContainerNameSlurmctld,
			ContainerPort: container.Port,
			Protocol:      corev1.ProtocolTCP,
		}},
		VolumeMounts: []corev1.VolumeMount{
			common.RenderVolumeMountSlurmConfigs(),
			common.RenderVolumeMountSpool(consts.ComponentTypeController, consts.SlurmctldName),
			common.RenderVolumeMountJail(),
			common.RenderVolumeMountMungeSocket(),
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				TCPSocket: &corev1.TCPSocketAction{
					Port: intstr.FromInt32(container.Port),
				},
			},
		},
		SecurityContext: &corev1.SecurityContext{
			Capabilities: &corev1.Capabilities{
				Add: []corev1.Capability{
					consts.ContainerSecurityContextCapabilitySysAdmin,
				},
			},
		},
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:              container.Resources.CPU,
				corev1.ResourceMemory:           container.Resources.Memory,
				corev1.ResourceEphemeralStorage: container.Resources.EphemeralStorage,
			},
		},
	}
}

// renderContainerMunge renders Slurm controller [corev1.Container] for munge
func renderContainerMunge(container *values.Container) corev1.Container {
	return corev1.Container{
		Name:            consts.ContainerNameMunge,
		Image:           container.Image,
		ImagePullPolicy: corev1.PullAlways, // TODO use digest and set to corev1.PullIfNotPresent
		VolumeMounts: []corev1.VolumeMount{
			common.RenderVolumeMountMungeKey(),
			common.RenderVolumeMountJail(),
			common.RenderVolumeMountMungeSocket(),
		},
		ReadinessProbe: &corev1.Probe{
			ProbeHandler: corev1.ProbeHandler{
				Exec: &corev1.ExecAction{
					Command: []string{
						"/bin/sh",
						"-c",
						fmt.Sprintf(
							"'test -S %s'",
							path.Join(consts.VolumeMountPathMungeSocket, "munge.socket.2"),
						),
					},
				},
			},
		},
		SecurityContext: &corev1.SecurityContext{
			Capabilities: &corev1.Capabilities{
				Add: []corev1.Capability{
					consts.ContainerSecurityContextCapabilitySysAdmin,
				}}},
		Resources: corev1.ResourceRequirements{
			Limits: corev1.ResourceList{
				corev1.ResourceCPU:              container.Resources.CPU,
				corev1.ResourceMemory:           container.Resources.Memory,
				corev1.ResourceEphemeralStorage: container.Resources.EphemeralStorage,
			},
		},
	}
}
