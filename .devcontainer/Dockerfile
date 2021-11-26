# [Choice] https://github.com/KEINOS/VSCode-Dev-Container-Go/pkgs/container/vscode-dev-container-go
ARG VARIANT="latest"

# -----------------------------------------------------------------------------
#  Main Stage
# -----------------------------------------------------------------------------
# ghcr.io/keinos/vscode-dev-container-go is an Alpine-based container image which
# provides a VSCode development environment over Docker. This image is weekly
# scanned for vulnerability and security with Synk, Dockle and Trivy. For detals
# see: ./README.md
FROM ghcr.io/keinos/vscode-dev-container-go:${VARIANT}

# [Optional] Uncomment this section to install additional OS packages.
# USER root
# RUN apk add --no-cache <your-package-list-here>

USER vscode
RUN cd /tmp && go install golang.org/x/lint/golint@latest