#!/bin/bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)/.."

${project_root}/scripts/create_website.sh
${project_root}/scripts/push_website_to_s3.sh
