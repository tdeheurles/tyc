#!/usr/bin/env bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)/.."

if [[ "${AWS_ACCESS_KEY_ID+x}" == "" ]]; then
    if [[ ! -e "${project_root}/.aws.secret" ]]; then
        echo -e "\e[31mError: please create your .aws.secret file (see README.md#define-your-aws-secrets)\e[39m"
        exit 1
    fi
    source "${project_root}/.aws.secret"
fi

cd ${project_root}/infrastructure/terraform

echo "-- init"
terraform init

echo "-- fmt"
terraform fmt

echo "-- validate"
terraform validate


echo "-- apply"
terraform apply
