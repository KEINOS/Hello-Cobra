#!/bin/bash
# This script will be ran after the container was up.
#
# This file will be placed under "/postCreateCommand.sh" while building the Docker
# image. (See: Dockerfile)
# Therefore if you change something here, you need to re-build the image.

set -eu # Stop the scipt on any error

echo "- Running post comands (postCreateCommand.sh)"

# Copy the VSCode's setting file which VSCode loads after connecting to the
# "VSCode server" in the container.
echo '- Copying VSCode settings file (vscode.json) ...'
cp /settings.vscode.json "${HOME}/.vscode-server/data/Machine/settings.json"

echo '- Current VSCode setting (settings.json) ...'
cat "${HOME}/.vscode-server/data/Machine/settings.json"

echo '- Copying cobra settings file (cobra.yaml) ... '
cp /cobra.yaml  "${HOME}/.cobra.yaml"

# This is a welcome message of bash. To show basic settings and TIPs to
# the terminal user.
cat /welcome.sh > "${HOME}/.welcome.sh"
echo '/bin/sh ~/.welcome.sh' >> "${HOME}/.bashrc"
