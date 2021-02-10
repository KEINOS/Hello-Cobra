#!/bin/bash
# Place any script here which you want to run after creating the container.
echo '==============================================================================='
echo ' Post Create Command'
echo '==============================================================================='

set -eu

# Sim-link Welcome message for bash
ln -s "$(pwd)/.devcontainer/welcome.sh" "${HOME}/.welcome.sh"

# Single quotes are intensional. So, not to expand expressions.
# shellcheck disable=SC2016
echo '"${HOME}/.welcome.sh"' >>"${HOME}/.bashrc"

# Sim-llink Cobra configuration file to home
ln -s "$(pwd)/.devcontainer/cobra.yaml" "${HOME}/.cobra.yaml"

# Make sure go.mod matches the source code in the module.
go mod tidy

# Set language
LANGUAGE="${LANG//\./:}" # <-- echo $LANG | sed "s/\./:/"
echo "export LANGUAGE=${LANGUAGE}" >>"${HOME}/.bashrc"

echo "${LC_ALL} UTF-8" >>/etc/locale.gen
locale-gen "$LC_ALL"
update-locale LANG="$LANG"
