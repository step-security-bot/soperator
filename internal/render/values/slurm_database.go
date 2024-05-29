package values

import (
	"github.com/go-logr/logr"

	slurmv1 "nebius.ai/slurm-operator/api/v1"
)

// SlurmDatabase contains the data needed for configuration of database
// TODO database configuration
type SlurmDatabase struct{}

func buildSlurmDatabaseFrom(_ logr.Logger, _ *slurmv1.SlurmCluster) (SlurmDatabase, error) {
	return SlurmDatabase{}, nil
}
