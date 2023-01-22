#!/bin/bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)/.."

if [[ ! -e "${project_root}/.aws.secret" ]]; then
    echo -e "\e[31mError: please create your .aws.secret file (see README.md#define-your-aws-secrets)\e[39m"
    exit 1
fi
source "${project_root}/.aws.secret"

cd ./hugo/public
aws s3 sync . s3://the-youtube-synopsis/ --delete
