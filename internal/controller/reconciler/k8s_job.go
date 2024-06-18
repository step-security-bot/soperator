package reconciler

import (
	"context"

	"github.com/pkg/errors"
	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	slurmv1 "nebius.ai/slurm-operator/api/v1"
	"nebius.ai/slurm-operator/internal/logfield"
)

type JobReconciler struct {
	*Reconciler
}

var (
	_ patcher = &JobReconciler{}
)

func NewJobReconciler(r *Reconciler) *JobReconciler {
	return &JobReconciler{
		Reconciler: r,
	}
}

func (r *JobReconciler) Reconcile(
	ctx context.Context,
	cluster *slurmv1.SlurmCluster,
	desired *batchv1.Job,
	deps ...metav1.Object,
) error {
	if err := r.reconcile(ctx, cluster, desired, deps...); err != nil {
		log.FromContext(ctx).
			WithValues(logfield.ResourceKV(desired)...).
			Error(err, "Failed to reconcile Job")
		return errors.Wrap(err, "reconciling Job")
	}
	return nil
}

func (r *JobReconciler) patch(existing, desired client.Object) (client.Patch, error) {
	patchImpl := func(e, d *batchv1.Job) client.Patch {
		res := client.MergeFrom(e.DeepCopy())

		//e.Spec.Schedule = d.Spec.Schedule
		//e.Spec.Suspend = d.Spec.Suspend
		//e.Spec.JobTemplate.Spec.BackoffLimit = d.Spec.JobTemplate.Spec.BackoffLimit

		return res
	}
	return patchImpl(existing.(*batchv1.Job), desired.(*batchv1.Job)), nil
}
