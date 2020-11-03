# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT="1.15"

# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.145.1/containers/go/.devcontainer/base.Dockerfile
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# Language settings (Provide the below ENV variables via Docker env args while building)
ARG LC_ALL='en_US.utf8'
ARG LANG='en_US.utf8'
ARG LANG_CHARMAP='UTF-8'
ARG LANG_LOCALNAME
ARG LANG_LOCALDEF
ARG LANG_LOCALE
ENV \
    # Enforce module mode
    GO111MODULE=on \
    # Fix: https://github.com/microsoft/vscode-dev-containers/issues/51
    SHELL="/bin/bash" \
    LC_ALL="${LANG_LOCALE:-$LC_ALL}" \
    LANG="${LANG_LOCALE:-$LANG}" \
    PATH="/usr/local/go/bin:${PATH}"
RUN localedef -f "${LANG_CHARMAP:-'UTF-8'}" -i "${LANG_LOCALDEF:-en_US}" "${LANG_LOCALNAME:-en_US}"

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "source /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi


# The below copied files will be placed into the proper dir after the container
# is up.
# See: "postCreateCommand" section in "devcontainer.json" and "postCreateCommand.sh"
# file.
COPY postCreateCommand.sh /postCreateCommand.sh
COPY cobra.yaml /cobra.yaml
COPY welcome.sh /welcome.sh
COPY settings.vscode.json /settings.vscode.json

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
    xz-utils

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

# Install anything else needed to go get
RUN go get -u -v \
        # gopls (language server)
        #   `gopls` will install the packages in the URL below as well.
        #   https://github.com/golang/vscode-go/blob/master/docs/tools.md#table-of-contents
        "golang.org/x/tools/gopls@latest" \
        # gotests (Test generator)
        "github.com/cweill/gotests/..." \
        # go-carpet (Cover area checker)
        "github.com/msoap/go-carpet" \
        # shfmt
        "mvdan.cc/sh/v3/cmd/shfmt" \
        # Other requirements of VSCode extensions
        "golang.org/x/lint/golint" \
        "github.com/ramya-rao-a/go-outline" \
        "github.com/rogpeppe/godef" \
        "github.com/uudashr/gopkgs/v2/cmd/gopkgs" \
    # irregular install
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
    && echo "- List dir: ${GOPATH}/bin" && ls -l "${GOPATH}/bin" \
    && echo '- List dir: /go/bin' && ls -l '/go/bin'

# Miscellaneous
RUN \
    # Set path
    echo 'export PATH="/go/bin:/usr/local/go/bin:${PATH}"' >> "${HOME}/.bashrc" \
    # Welcome message of bash
    && echo '/welcome.sh' >> "${HOME}/.bashrc"

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1
