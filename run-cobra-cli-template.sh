#!/usr/bin/env bash
set -eux pipefail

# Build the CLI binary.
bash build.sh

if [[ ! -x "./bin/cobra-cli-template" ]]; then
  echo "Error: ./bin/cobra-cli-template not found or not executable. Please run 'bash build.sh' first." >&2
  exit 1
else
  echo "Found ./bin/cobra-cli-template"
fi

./bin/cobra-cli-template
