{{/* Local storage class */}}
{{- define "slurm-cluster-filestore.class.local.name" -}}
    {{- default "slurm-local-pv" .Values.storageClass.local.name | trim | kebabcase | quote -}}
{{- end }}

{{/*
---
*/}}

{{/* Jail device name */}}
{{- define "slurm-cluster-filestore.volume.jail.device" -}}
    {{- default "jail" .Values.volume.jail.filestoreDeviceName | trim | kebabcase -}}
{{- end }}

{{/* Jail volume */}}
{{- define "slurm-cluster-filestore.volume.jail.name" -}}
    {{- default "jail" .Values.volume.jail.name | trim | kebabcase -}}
{{- end }}

{{/* Jail PVC name */}}
{{- define "slurm-cluster-filestore.volume.jail.pvc" -}}
    {{- cat (include "slurm-cluster-filestore.volume.jail.name" .) "pvc" | kebabcase | quote -}}
{{- end }}

{{/* Jail PV name */}}
{{- define "slurm-cluster-filestore.volume.jail.pv" -}}
    {{- cat (include "slurm-cluster-filestore.volume.jail.name" .) "pv" | kebabcase | quote -}}
{{- end }}

{{/* Jail mount name */}}
{{- define "slurm-cluster-filestore.volume.jail.mount" -}}
    {{- cat (include "slurm-cluster-filestore.volume.jail.name" .) "mount" | kebabcase | quote -}}
{{- end }}

{{/* Jail storage class name */}}
{{- define "slurm-cluster-filestore.volume.jail.storageClass" -}}
    {{- include "slurm-cluster-filestore.class.local.name" . -}}
{{- end }}

{{/* Jail size */}}
{{- define "slurm-cluster-filestore.volume.jail.size" -}}
    {{- required "Jail volume size is required." .Values.volume.jail.size -}}
{{- end }}

{{/*
---
*/}}

{{/* Spool device name */}}
{{- define "slurm-cluster-filestore.volume.spool.device" -}}
    {{- default "controller-spool" .Values.volume.spool.filestoreDeviceName | trim | kebabcase -}}
{{- end }}

{{/* Spool volume */}}
{{- define "slurm-cluster-filestore.volume.spool.name" -}}
    {{- default "spool" .Values.volume.spool.name | trim | kebabcase -}}
{{- end }}

{{/* Spool PVC name */}}
{{- define "slurm-cluster-filestore.volume.spool.pvc" -}}
    {{- cat (include "slurm-cluster-filestore.volume.spool.name" .) "pvc" | kebabcase | quote -}}
{{- end }}

{{/* Spool PV name */}}
{{- define "slurm-cluster-filestore.volume.spool.pv" -}}
    {{- cat (include "slurm-cluster-filestore.volume.spool.name" .) "pv" | kebabcase | quote -}}
{{- end }}

{{/* Spool mount name */}}
{{- define "slurm-cluster-filestore.volume.spool.mount" -}}
    {{- cat (include "slurm-cluster-filestore.volume.spool.name" .) "mount" | kebabcase | quote -}}
{{- end }}

{{/* Spool storage class name */}}
{{- define "slurm-cluster-filestore.volume.spool.storageClass" -}}
    {{- include "slurm-cluster-filestore.class.local.name" . -}}
{{- end }}

{{/* Spool size */}}
{{- define "slurm-cluster-filestore.volume.spool.size" -}}
    {{- required "Spool volume size is required." .Values.volume.spool.size -}}
{{- end }}

{{/*
---
*/}}

{{/* GPU node group */}}
{{- define "slurm-cluster-filestore.nodeGroup.gpu" -}}
    {{- required "GPU node group ID is required." .Values.nodeGroup.gpu.id | quote -}}
{{- end }}

{{/* Non-GPU node group */}}
{{- define "slurm-cluster-filestore.nodeGroup.nonGpu" -}}
    {{- required "Non-GPU node group ID is required." .Values.nodeGroup.nonGpu.id | quote -}}
{{- end }}
