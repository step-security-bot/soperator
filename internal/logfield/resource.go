package logfield

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func ResourceKV(obj client.Object) []any {
	return []any{
		ResourceKind, obj.GetObjectKind().GroupVersionKind().String(),
		ResourceName, obj.GetName(),
	}
}
