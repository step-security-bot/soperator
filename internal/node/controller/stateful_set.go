package controller

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/yaml"

	"nebius.ai/slurm-operator/internal/consts"
	k8smodels "nebius.ai/slurm-operator/internal/models/k8s"
	smodels "nebius.ai/slurm-operator/internal/models/slurm"
)

// BuildStatefulSet create a new stateful set for the Slurm controllers
func BuildStatefulSet(cluster smodels.ClusterValues) (*appsv1.StatefulSet, error) {
	labels := k8smodels.BuildClusterDefaultLabels(cluster.Name, consts.ComponentTypeController)
	matchLabels := k8smodels.BuildClusterDefaultMatchLabels(cluster.Name, consts.ComponentTypeController)

	// TODO proper versioning
	selfStsVersion := map[string]string{
		"self-sts": "version-placeholder-000",
	}
	selfStsVersionYaml, err := yaml.Marshal(selfStsVersion)
	if err != nil {
		return nil, err
	}

	selfPodTmplVersion := map[string]string{
		"self-pod-tmpl": "version-placeholder-001",
	}
	selfPodTmplVersionString, err := yaml.Marshal(selfPodTmplVersion)
	if err != nil {
		return nil, err
	}

	maxUnavailable := intstr.FromInt32(1)

	dep := &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.Controller.StatefulSet.Name,
			Namespace: cluster.Controller.StatefulSet.Namespace,
			Labels:    labels,
			Annotations: map[string]string{
				consts.AnnotationVersions: string(selfStsVersionYaml),
			},
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: cluster.Controller.Service.Name,
			Replicas:    &cluster.Controller.StatefulSet.Replicas,
			UpdateStrategy: appsv1.StatefulSetUpdateStrategy{
				Type: appsv1.RollingUpdateStatefulSetStrategyType,
				RollingUpdate: &appsv1.RollingUpdateStatefulSetStrategy{
					MaxUnavailable: &maxUnavailable,
				},
			},
			Selector: &metav1.LabelSelector{
				MatchLabels: matchLabels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
					Annotations: map[string]string{
						consts.AnnotationVersions: string(selfPodTmplVersionString),
					},
				},
				Spec: corev1.PodSpec{
					Affinity:    cluster.Controller.Affinity,
					Tolerations: cluster.Controller.Tolerations,
					// TODO move to dedicated function
					Containers: []corev1.Container{{
						Name:  consts.ContainerControllerSlurmCtldName,
						Image: cluster.Controller.Image.String(),
						// TODO use digest and set to corev1.PullIfNotPresent
						ImagePullPolicy: corev1.PullAlways,
						// TODO move to dedicated function
						Ports: []corev1.ContainerPort{
							{
								Name:          consts.ServiceControllerClusterTargetPort,
								ContainerPort: consts.ServiceControllerClusterPort,
								Protocol:      consts.ServiceControllerClusterPortProtocol,
							},
						},
						VolumeMounts: BuildVolumeMounts(cluster),
						Resources: corev1.ResourceRequirements{
							Requests: cluster.Controller.Resources,
							Limits:   cluster.Controller.Resources,
						},
						// TODO LivenessProbe + ReadinessProbe + StartupProbe
					}},
					Volumes: BuildVolumes(cluster),
					// TODO fill node selector based on disk AZ on non-GPU node
					NodeSelector: map[string]string{},
				},
			},
		},
	}
	return dep, nil
}
