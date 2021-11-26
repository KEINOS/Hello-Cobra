# =============================================================================
#  Test Container for Vaious Go Versions
# =============================================================================
# Default version
ARG VARIANT="1.15-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

ENV GO111MODULE=on

RUN apk add --no-cache \
    git \
    alpine-sdk \
    build-base

WORKDIR /workspaces

ENTRYPOINT go mod download && go test -race ./...
