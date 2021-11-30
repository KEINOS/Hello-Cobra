#!/bin/sh
# =============================================================================
#  Test Script for Homebrew Functionality
# =============================================================================

set -eu

git config --global init.defaultBranch main
brew install KEINOS/Hello-Cobra/hello-cobra

# Smoke test
hello-cobra --version
hello-cobra hello | grep Hi
