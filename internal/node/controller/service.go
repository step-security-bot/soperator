package controller

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"nebius.ai/slurm-operator/internal/consts"
	"nebius.ai/slurm-operator/internal/models/k8s"
	"nebius.ai/slurm-operator/internal/models/slurm"
)

// BuildService creates a new Service for Slurm controller's pods.
// It exposes the following port:
// - [consts.ServiceControllerClusterPort]: the port at which the controller listens for workers
func BuildService(cluster smodels.ClusterValues) *corev1.Service {
	labels := k8smodels.BuildClusterDefaultLabels(cluster.Name, consts.ComponentTypeController)
	matchLabels := k8smodels.BuildClusterDefaultMatchLabels(cluster.Name, consts.ComponentTypeController)

	svc := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      cluster.Controller.Service.Name,
			Namespace: cluster.Controller.Service.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceTypeClusterIP,
			Selector: matchLabels,
			Ports:    cluster.Controller.Service.Ports,
		},
	}

	return svc
}
