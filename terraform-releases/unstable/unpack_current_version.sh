#!/bin/bash

set -e

# This script unpacks a terraform release tarball with the current VERSION into the working directory excluding the
# terraform variables file.

version=$(cat < ../../VERSION | tr -d '\n')
version_formatted=$(echo "${version}" | tr '-' '_' | tr '.' '_')
tarball="slurm_operator_tf_$version_formatted.tar.gz"

if [ ! -f "$tarball" ]; then
    echo "No release with the current version $version: file with name $tarball doesn't exist"
    exit 1
fi

echo "Extracting tarball $tarball"
tar -xf "${tarball}" --exclude "^terraform/terraform.tfvars$"

echo "Updating slurm_operator_version in the existing terraform.tfvars file"
sed -i.bak -E "s/(slurm_operator_version[[:space:]]*=[[:space:]]*\").*(\")/\1${version}\2/" terraform/terraform.tfvars
