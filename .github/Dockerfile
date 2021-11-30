# =============================================================================
#  Test Container for Vaious Go Versions (ver. 2021.11.30-09:11)
# =============================================================================

# Default version
ARG VARIANT="1.15-alpine"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
FROM golang:${VARIANT}

ENV GO111MODULE=on

RUN apk add --no-cache \
    build-base \
    alpine-sdk \
    git \
    bash

WORKDIR /workspaces

ENTRYPOINT go mod download && go test -race ./...
