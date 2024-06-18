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

type CronJobReconciler struct {
	*Reconciler
}

var (
	_ patcher = &CronJobReconciler{}
)

func NewCronJobReconciler(r *Reconciler) *CronJobReconciler {
	return &CronJobReconciler{
		Reconciler: r,
	}
}

func (r *CronJobReconciler) Reconcile(
	ctx context.Context,
	cluster *slurmv1.SlurmCluster,
	desired *batchv1.CronJob,
	deps ...metav1.Object,
) error {
	if err := r.reconcile(ctx, cluster, desired, r.patch, deps...); err != nil {
		log.FromContext(ctx).
			WithValues(logfield.ResourceKV(desired)...).
			Error(err, "Failed to reconcile CronJob")
		return errors.Wrap(err, "reconciling CronJob")
	}
	return nil
}

func (r *CronJobReconciler) patch(existing, desired client.Object) (client.Patch, error) {
	patchImpl := func(e, d *batchv1.CronJob) client.Patch {
		res := client.MergeFrom(e.DeepCopy())

		e.Spec.Schedule = d.Spec.Schedule
		e.Spec.Suspend = d.Spec.Suspend
		e.Spec.SuccessfulJobsHistoryLimit = d.Spec.SuccessfulJobsHistoryLimit
		e.Spec.FailedJobsHistoryLimit = d.Spec.FailedJobsHistoryLimit
		e.Spec.JobTemplate.Spec.Template.Spec = d.Spec.JobTemplate.Spec.Template.Spec

		return res
	}
	return patchImpl(existing.(*batchv1.CronJob), desired.(*batchv1.CronJob)), nil
}