package k8smodels

import (
	"k8s.io/apimachinery/pkg/types"

	slurmv1 "nebius.ai/slurm-operator/api/v1"
	"nebius.ai/slurm-operator/internal/consts"
	k8snaming "nebius.ai/slurm-operator/internal/models/k8s/naming"
)

type StatefulSet struct {
	types.NamespacedName

	Replicas int32
}

func BuildStatefulSetFrom(
	namespace,
	clusterName string,
	componentType consts.ComponentType,
	controllerSpec slurmv1.ControllerNodeSpec,
) StatefulSet {
	return StatefulSet{
		NamespacedName: types.NamespacedName{
			Namespace: namespace,
			Name:      k8snaming.BuildStatefulSetName(clusterName, componentType),
		},
		Replicas: controllerSpec.Size,
	}
}
