# Dockerfile for CIs and Development.
#
# For DEVs:
#   This is for Docker + VSCode + "Remote-Containers Extension" users.
#   From VSCode press "F1" and select "Remote-Containers: Open in Container"
#   after a while you'll get most of the environment to develop/debug.
#   Add and PR if any other
FROM golang:alpine

USER root

ENV \
    # Set module cache dir under user dir.
    GOPATH=/root/go \
    # Enforce module mode
    GO111MODULE=on

# The below copied files will be placed into the proper dir after the container
# is up.
# See: "postCreateCommand" section in "devcontainer.json" and "postCreateCommand.sh"
# file.
COPY postCreateCommand.sh /postCreateCommand.sh
COPY cobra.yaml /cobra.yaml
COPY welcome.sh /welcome.sh
COPY settings.vscode.json /settings.vscode.json

# Default timezone（Japan = JST = GMT+9:00 = UTC+0900 = Asia/Tokyo）
# Change this if it's inconvenient.
ARG NAME_ZONE='Japan'

# Common setup
RUN \
    # Update/upgrade apk packages.
    # Since VSCode installs additional packages after building the Docker image,
    # we won't use "--no-cache" option and cache them to speed up booting.
    apk update && \
    apk upgrade && \
    \
    # Set timezone.
    date && \
    apk add tzdata && \
    cp "/usr/share/zoneinfo/${NAME_ZONE}" /etc/localtime && \
    echo "$NAME_ZONE" > /etc/timezone && \
    apk del tzdata && \
    # Smoke test
    date && \
    go env && \
    go version

# Install apk packages
RUN \
    apk add \
        # As a default shell
        bash \
        # Basic packages needed to compile/build
        alpine-sdk \
        build-base \
        # The `rg` is a string search command from files. ex) rg ./cmd foo
        ripgrep \
        # Static analyzer and formatter for shell script
        shellcheck \
        shfmt \
        # Directory tree viewer for documentation.
        tree

# Install binaries of go packages.
# Mostly the requirements for VSCode's extensions and to ease dev.
RUN \
    [ -d /go/bin ] && { rm -rf /go/bin; } && \
    # Sym-link to unite bin diretory to ease check which commands were installed.
    echo "- Sym-link '${GOPATH}/bin' to '/go/bin':" && ln -s "${GOPATH}/bin" /go && \
    echo "- Go env:" && go env && \
    echo "- List tree ${GOPATH}(GOPATH):" && tree "${GOPATH}" && \
    go get -u -v \
        # gopls (language server)
        #   `gopls` will install the packages in the URL below as well.
        #   https://github.com/golang/vscode-go/blob/master/docs/tools.md#table-of-contents
        "golang.org/x/tools/gopls@latest" \
        # gotests (Test generator)
        "github.com/cweill/gotests/..." \
        # go-carpet (Cover area checker)
        "github.com/msoap/go-carpet" \
        # Other requirements of VSCode extensions
        "golang.org/x/lint/golint" \
        "github.com/ramya-rao-a/go-outline" \
        "github.com/rogpeppe/godef" \
        "github.com/uudashr/gopkgs/v2/cmd/gopkgs" \
        && \
    # irregular install
    cd "${GOPATH}/pkg/mod" && \
    go get -v -d \
        # gocode and gocode-modules is no-good to install as module
        "github.com/stamblerre/gocode" \
        # cobra (sub-command creation)
        "github.com/spf13/cobra" \
        # gomock
        "github.com/golang/mock/gomock" \
        && \
    # gocode and gocode-modules
    cd ${GOPATH}/pkg/mod/github.com/stamblerre/gocode* && \
    go build -o "${GOPATH}/bin/gocode-gomod" . && \
    # cobra command
    cd ${GOPATH}/pkg/mod/github.com/spf13/cobra*/cobra && \
    go build -o "${GOPATH}/bin/cobra" . && \
    # mockgen
    cd ${GOPATH}/pkg/mod/github.com/golang/mock*/mockgen && \
    go install && \
    # List bin files
    echo "- List dir: ${GOPATH}/bin" && ls -l "${GOPATH}/bin" && \
    echo '- List dir: /go/bin' && ls -l '/go/bin'
