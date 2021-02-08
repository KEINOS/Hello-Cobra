#!/bin/bash
# =============================================================================
#  Test Script to Run GolangCI-Lint
# =============================================================================
#
#  Requirements:
#    - golangci-lint: https://github.com/golangci/golangci-lint

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
PATH_DIR_PARENT="$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"
FAILURE=1

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------
set -eu
set -o pipefail

cd "${PATH_DIR_PARENT}" || {
    echo >&2 "Failed to change directory: ${PATH_DIR_PARENT}"
    exit $FAILURE
}

golangci-lint run ./...
