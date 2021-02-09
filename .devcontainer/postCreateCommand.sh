#!/bin/bash
# Place any script here which you want to run after creating the container.

set -eu

# Sim-llink Welcome message for bash
ln -s "$(pwd)/.devcontainer/welcome.sh" "${HOME}/.welcome.sh"

# Single quotes are intensional. So, not to expand expressions.
# shellcheck disable=SC2016
echo '"${HOME}/.welcome.sh"' >>"${HOME}/.bashrc"

# Sim-llink Cobra configuration file to home
ln -s "$(pwd)/.devcontainer/cobra.yaml" "${HOME}/.cobra.yaml"

# Make sure go.mod matches the source code in the module.
go mod tidy
