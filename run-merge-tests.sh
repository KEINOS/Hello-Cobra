#!/bin/bash

set -eu

# indentStdIn indents the STDIN given to the function
function indentStdIn() {
    indent="\t"
    while read -r line; do
        echo -e "${indent}${line}"
    done
    echo
}

function runShfmt() {
    echo -n '- Shell format ... '
    result=$(find . -name '*.sh' -type f -print0 | xargs -0 shfmt -d 2>&1) || {
        echo 'NG'
        echo "$result"
        return 1
    }
    echo 'OK'
    return 0
}

function runShellCheck() {
    echo -n '- Shell check ... '
    result=$(find . -name '*.sh' -type f -print0 | xargs -0 shellcheck 2>&1) || {
        echo 'NG'
        echo "$result"
        return 1
    }
    echo 'OK'
    return 0
}

function runGofmt() {
    echo -n '- Go format ... '
    if [ -z "$(gofmt -l .)" ]; then
        echo 'OK'
        return 0
    fi
    echo 'NG'
    gofmt -d -e . 2>&1 | indentStdIn
    return 1
}

function runGoUnitTests() {
    echo -n '- Go unit tests ... '
    result=$(./run-tests.sh -v 2>&1) || {
        echo 'NG'
        echo "$result" | indentStdIn
        return 1
    }
    echo 'OK'
    return 0
}

runShfmt
runShellCheck
runGofmt
runGoUnitTests
