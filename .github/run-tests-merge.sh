#!/bin/bash
# =============================================================================
#  Test Script to Run Before merge.
# =============================================================================
#  This script will run the following tests before merge.
#    - shfmt: Shell script linter for any ".sh" file
#    - shellcheck: Shell script static analysis for any ".sh" file
#    - gofmt: Golang linter for any ".go" file
#    - go test(run-tests.sh): Golang unit tests and coverage
#
#  Requirements:
#    - go-carpet: https://github.com/msoap/go-carpet
#    - shfmt: https://github.com/mvdan/sh
#    - shellcheck: https://github.com/koalaman/shellcheck

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
PATH_DIR_PARENT="$(dirname "$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)")"
SUCCESS=0
FAILURE=1

# -----------------------------------------------------------------------------
#  Functions
# -----------------------------------------------------------------------------
function indentStdIn() {
    indent="\t"
    while IFS= read -r line; do
        echo -e "${indent}${line}"
    done
    echo
}

function runShfmt() {
    echo -n '- Shell format ... '
    result=$(find . -name '*.sh' -type f -print0 | xargs -0 shfmt -d 2>&1) || {
        echo 'NG'
        echo "$result"
        return $FAILURE
    }
    echo 'OK'
    return $SUCCESS
}

function runShellCheck() {
    echo -n '- Shell check ... '
    result=$(find . -name '*.sh' -type f -print0 | xargs -0 shellcheck 2>&1) || {
        echo 'NG'
        echo "$result"
        return $FAILURE
    }
    echo 'OK'
    return $SUCCESS
}

function runGofmt() {
    echo -n '- Go format ... '
    if [ -z "$(gofmt -l .)" ]; then
        echo 'OK'
        return $SUCCESS
    fi
    echo 'NG'
    gofmt -d -e . 2>&1 | indentStdIn
    return $FAILURE
}

function runGolangCiLint() {
    echo -n '- Lint check (golangci-lint) ... '
    result=$(./.github/run-tests-lint.sh -v 2>&1) || {
        echo 'NG'
        echo "$result" | indentStdIn
        return $FAILURE
    }
    echo 'OK'
    return $SUCCESS
}

function runGoUnitTests() {
    echo -n '- Go unit tests ... '
    result=$(./.github/run-tests-coverage.sh -v 2>&1) || {
        echo 'NG'
        echo "$result" | indentStdIn
        return $FAILURE
    }
    echo 'OK'
    return $SUCCESS
}

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------

# Do not allow undecleared variables and function failure.
set -eu

cd "$PATH_DIR_PARENT"

# Check current tree structure for CI debuging purpose.
echo '- Current tree'
tree | indentStdIn

# Run tests
runShfmt
runShellCheck
runGofmt
runGoUnitTests
runGolangCiLint
