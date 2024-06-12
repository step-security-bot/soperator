#!/bin/bash

set -e

usage() { echo "usage: ${0} -c <context_name> [-h]" >&2; exit 1; }

while getopts c:h flag
do
    case "${flag}" in
        c) context_name=${OPTARG};;
        h) usage;;
        *) usage;;
    esac
done

if [ -z "$context_name" ]; then
    usage
fi

kubectl --context="$context_name" --namespace=dstaroff apply -f common/local_pv_storageclass.yaml

kubectl --context="$context_name" --namespace=dstaroff apply -f common/jail/pvc.yaml
kubectl --context="$context_name" --namespace=dstaroff apply -f common/jail/pv.yaml

kubectl --context="$context_name" --namespace=dstaroff apply -f node/controller/spool_pvc.yaml
kubectl --context="$context_name" --namespace=dstaroff apply -f node/controller/spool_pv.yaml

kubectl --context="$context_name" --namespace=dstaroff apply -f node/controller/spool_mount_daemonset.yaml
kubectl --context="$context_name" --namespace=dstaroff apply -f common/jail/mount_daemonset.yaml
kubectl --context="$context_name" --namespace=dstaroff rollout status daemonset jail-mount --timeout=10h

kubectl --context="$context_name" --namespace=dstaroff apply -f common/jail/populate_jail_job.yaml
kubectl --context="$context_name" --namespace=dstaroff wait --for=condition=complete --timeout=10h job/populate-jail
