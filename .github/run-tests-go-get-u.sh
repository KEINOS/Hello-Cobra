#!/bin/sh
# =============================================================================
#  `go get -u` test (Go 1.15)
# =============================================================================

cd /tmp || {
    echo >&2 "failed to move temp dir"
    exit 1
}

GO111MODULE="on" go get -u "github.com/KEINOS/Hello-Cobra/hello-cobra@latest" || {
    echo >&2 "failed to install hello-cobra command via 'go get -u'"
    exit 1
}

hello-cobra --version | grep "hello-cobra version v" || {
    echo >&2 "failed to execute command. Version info is missing"
    exit 1
}
