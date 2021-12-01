#!/bin/sh
# =============================================================================
#  `go install` test (Go 1.16+)
# =============================================================================

cd /tmp || {
    echo >&2 "failed to move temp dir"
    exit 1
}

go install "github.com/KEINOS/Hello-Cobra/hello-cobra@latest" || {
    echo >&2 "failed to install hello-cobra command via 'go install'"
    exit 1
}

hello-cobra --version | grep "hello-cobra version" || {
    echo >&2 "failed to execute command. Version info is missing"
    exit 1
}
