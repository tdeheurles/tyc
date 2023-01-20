#!/bin/bash
set -euo pipefail

project_root="$(cd "$(dirname "${BASH_SOURCE[0]}")" >/dev/null 2>&1 && pwd)/.."

${project_root}/scripts/generate_content_files.sh
cd ${project_root}/hugo
hugo
