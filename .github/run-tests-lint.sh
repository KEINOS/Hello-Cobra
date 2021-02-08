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
SUCCESS=0
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

if ! golangci-lint run --config ./.github/golangci.yml ./...; then
    echo >&2
    echo >&2 'References for debugging:'
    echo >&2 '  wsl    : https://github.com/bombsimon/wsl/blob/master/doc/rules.md'
    echo >&2 '  gofumpt: https://github.com/mvdan/gofumpt#added-rules'
    exit $FAILURE
fi

exit $SUCCESS
