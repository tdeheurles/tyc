#!/usr/bin/env bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)/.."

cd ${project_root}/infrastructure/terraform

echo "-- init"
terraform init

echo "-- fmt"
terraform fmt

echo "-- validate"
terraform validate


echo "-- apply"
terraform apply
