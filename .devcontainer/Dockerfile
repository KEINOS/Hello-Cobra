# =============================================================================
#  Pre Defined Variables
# =============================================================================
# User provided Go version of Docker image
ARG VARIANT
# Default Go version. Choices: 1, 1.15, 1.14
ARG VARIANT_DEFAULT="1"
# Default Language
ARG LC_ALL_DEFAULT='en_US.utf8'
ARG LANG_DEFAULT='en_US.utf8'
# Default Node.js settings
ARG INSTALL_NODE_DEFAULT="false"
ARG NODE_VERSION_DEFAULT="lts/*"

# =============================================================================
#  Define Dockerfile.
#
#  For image contents see:
#  https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/go/.devcontainer/base.Dockerfile
# =============================================================================
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT:-$VARIANT_DEFAULT}

# Declare user args to receive while building an image.
ARG LANG
ARG LC_ALL
ARG INSTALL_NODE
ARG NODE_VERSION

ENV \
    LANG="${LANG:-$LANG_DEFAULT}" \
    LC_ALL="${LC_ALL:-$LC_ALL_DEFAULT}" \
    PATH="/usr/local/go/bin:${PATH}" \
    # Enforce go module mode
    GO111MODULE='on' \
    # Fix: https://github.com/microsoft/vscode-dev-containers/issues/51
    SHELL="/bin/bash"

#RUN localedef -f UTF-8

# [Option] Install Node.js
ARG INSTALL_NODE="${INSTALL_NODE:-INSTALL_NODE_DEFAULT}"
ARG NODE_VERSION="${NODE_VERSION:-NODE_VERSION_DEFAULT}"
RUN if [ "${INSTALL_NODE}" = "true" ]; then \
    echo 'Installing Node.js'; \
    su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; \
    fi

# Bash script file to run right after the container was created.
# See: "postCreateCommand" section in "devcontainer.json" and "postCreateCommand.sh"
# file.
COPY postCreateCommand.sh /postCreateCommand.sh

# Install additional OS packages.
RUN export DEBIAN_FRONTEND=noninteractive \
    && apt-get update \
    && apt-get -y upgrade \
    && apt-get -y install --no-install-recommends \
    # The `rg` is a string search command from files. ex) rg ./cmd foo
    ripgrep \
    # Directory tree viewer for documentation.
    tree \
    # xz to untar xz compressed files
    xz-utils \
    # missing locales
    locales

# ShellCheck - Static analyzer and formatter for shell script
# Note: Install the latest shellcheck. See: https://github.com/koalaman/shellcheck/issues/704
RUN url_download="https://github.com/koalaman/shellcheck/releases/download/latest/shellcheck-latest.linux.$(uname -m).tar.xz" \
    && path_tmp_dir="$(pwd)/tmp_install_dir" \
    && wget -P "${path_tmp_dir}/" "$url_download"\
    && tar xvf ${path_tmp_dir}/shellcheck* -C "${path_tmp_dir}/" \
    && cp "${path_tmp_dir}/shellcheck-latest/shellcheck" "$(dirname $(which tree))/shellcheck" \
    # Smoke test
    && shellcheck --version \
    && rm -r "$path_tmp_dir"

# golangci-lint - The fast Go linters runner.
# binary will be $(go env GOPATH)/bin/golangci-lint
RUN curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.36.0 \
    # Smoke test
    && golangci-lint --version

# Install anything else needed to go get
RUN go get -u -v \
    # gopls (language server)
    #   `gopls` will install the packages in the URL below as well.
    #   https://github.com/golang/vscode-go/blob/master/docs/tools.md#table-of-contents
    "golang.org/x/tools/gopls@latest" \
    # go-carpet (Cover area checker)
    "github.com/msoap/go-carpet" \
    # shfmt (shellformat) is a shell script linter
    "mvdan.cc/sh/v3/cmd/shfmt" \
    # gp enables to share go files over Go Playground
    "github.com/tenntenn/goplayground/cmd/gp" \
    # gomarkdoc generates documentation in markdown
    "github.com/princjef/gomarkdoc/cmd/gomarkdoc" \
    #"github.com/rogpeppe/godef" \
    # -------------------------------------------------------------------------
    #  irregular install (Download source and build)
    # -------------------------------------------------------------------------
    && cd "${GOPATH}/pkg/mod" \
    && go get -v -d \
    # gocode and gocode-modules is no-good to install as module
    "github.com/stamblerre/gocode" \
    # cobra (sub-command creation)
    "github.com/spf13/cobra" \
    # gomock
    "github.com/golang/mock/gomock" \
    # gocode and gocode-modules
    && cd ${GOPATH}/pkg/mod/github.com/stamblerre/gocode* \
    && go build -o "${GOPATH}/bin/gocode-gomod" . \
    # cobra command
    && cd ${GOPATH}/pkg/mod/github.com/spf13/cobra*/cobra \
    && go build -o "${GOPATH}/bin/cobra" . \
    # mockgen
    && cd ${GOPATH}/pkg/mod/github.com/golang/mock*/mockgen \
    && go install \
    # List bin files
    && echo "- List dir: ${GOPATH}/bin" && ls -l "${GOPATH}/bin"

# Miscellaneous
RUN \
    # Set path for go
    echo 'export PATH="/go/bin:/usr/local/go/bin:${PATH}"' >> "${HOME}/.bashrc"

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
