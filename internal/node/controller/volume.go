package controller

import (
	corev1 "k8s.io/api/core/v1"

	"nebius.ai/slurm-operator/internal/consts"
	"nebius.ai/slurm-operator/internal/models/slurm"
)

func BuildVolumes(cluster smodels.ClusterValues) []corev1.Volume {
	return []corev1.Volume{
		BuildKeyVolume(cluster),
		BuildConfigVolume(cluster),
		BuildSpoolVolume(cluster),
	}
}

func BuildVolumeMounts(cluster smodels.ClusterValues) []corev1.VolumeMount {
	return []corev1.VolumeMount{
		BuildKeyVolumeMount(cluster),
		BuildConfigVolumeMount(cluster),
		BuildSpoolVolumeMount(cluster),
	}
}

func BuildKeyVolume(_ smodels.ClusterValues) corev1.Volume {
	return corev1.Volume{
		Name: consts.VolumeSlurmKeyName,
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName: consts.SecretSlurmKeyName,
				Items: []corev1.KeyToPath{
					{
						Key:  consts.SecretSlurmKeySlurmKeyKey,
						Path: consts.SecretSlurmKeySlurmKeyPath,
						Mode: &consts.SecretSlurmKeySlurmKeyMode,
					},
				},
			},
		},
	}
}

func BuildKeyVolumeMount(_ smodels.ClusterValues) corev1.VolumeMount {
	return corev1.VolumeMount{
		Name:      consts.VolumeSlurmKeyName,
		MountPath: "/root/slurm-k8s-conf/key",
		ReadOnly:  true,
	}
}

func BuildConfigVolume(_ smodels.ClusterValues) corev1.Volume {
	return corev1.Volume{
		Name: consts.VolumeSlurmConfigsName,
		VolumeSource: corev1.VolumeSource{
			ConfigMap: &corev1.ConfigMapVolumeSource{
				LocalObjectReference: corev1.LocalObjectReference{Name: consts.ConfigMapSlurmConfigsName},
			},
		},
	}
}

func BuildConfigVolumeMount(_ smodels.ClusterValues) corev1.VolumeMount {
	return corev1.VolumeMount{
		Name:      consts.VolumeSlurmConfigsName,
		MountPath: "/root/slurm-k8s-conf/configs",
		ReadOnly:  true,
	}
}

func BuildSpoolVolume(_ smodels.ClusterValues) corev1.Volume {
	return corev1.Volume{
		Name: consts.VolumeSlurmConfigsName,
		VolumeSource: corev1.VolumeSource{
			PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
				ClaimName: consts.PVCControllerSpoolName,
				ReadOnly:  false,
			},
		},
	}
}

func BuildSpoolVolumeMount(_ smodels.ClusterValues) corev1.VolumeMount {
	return corev1.VolumeMount{
		Name:      consts.VolumeSlurmSpoolName,
		MountPath: "/var/spool",
	}
}
