#!/bin/sh
# =============================================================================
#  This script checks if the commands and packages required for merge testing
#  are installed. Run this script before running ./run-tests-merge.sh.
# =============================================================================

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
SUCCESS=0
FAILURE=1
TRUE=0
FALSE=1

# -----------------------------------------------------------------------------
#  Functions
# -----------------------------------------------------------------------------
isAvailable() {
    command_tmp="${1:?Command name missing}"
    msg_error="${2:?Error message missing}"
    url_reference="${3:?'URL for reference missing'}"

    printf -- '  %s ... ' "$command_tmp"
    if ! which "$command_tmp" 1>/dev/null 2>/dev/null; then
        flag_covered_all=$FALSE
        echo 'NG'
        echo "    - ABOUT  : ${msg_error}"
        echo "    - DETAILS: ${url_reference}"

        return $FALSE
    fi
    echo 'OK'
}

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------
flag_covered_all=$TRUE

echo 'Checking requirements for:'

echo
echo 'Shell scripts:'
isAvailable shellcheck \
    '"shellcheck" is a static analysis tool for shell scripts.' \
    'https://github.com/koalaman/shellcheck'

isAvailable shfmt \
    '"shfmt" is a linter for shell scripts to support POSIX Shell, Bash, and mksh.' \
    'https://github.com/mvdan/sh'

echo
echo 'Go programs:'
isAvailable go \
    '"go" is required as a matter of course.' \
    'https://golang.org/'

isAvailable gofmt \
    '"gofmt" is a formatter for golang.' \
    'https://golang.org/cmd/gofmt/'

isAvailable golint \
    '"golint" is a linter for Go source code. this is required for GoReportCard.com status.' \
    'go install golang.org/x/lint/golint@latest'

isAvailable golangci-lint \
    '"golangci-lint" is is a Go linters aggregator.' \
    'https://golangci-lint.run/'

if [ $flag_covered_all -ne 0 ]; then
    echo
    echo >&2 "ERROR: Some requirements were missing."
    exit $FAILURE
fi

echo
echo 'OK - All requirements were installed! You are good to Go testing!'
exit $SUCCESS
