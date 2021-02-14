<!-- markdownlint-disable MD033 -->
# Dockerfile for GitHub and VSCode Users To Develop

This directory is for [GitHub Codespaces](https://github.com/features/codespaces) and/or [VS Code + Docker](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) users for development.

It includes most of the necessary packages and tools for developing Golang app. Aiming to provide the same environment to develop the app.

## Developing Online

If GitHub detects this directory (`.devcontainer`) in the repo, then you will be able to develop online via [GitHub Codespaces](https://github.com/features/codespaces).

## VS Code + Docker User

The container contains VS Code Server as well. If you already have installed the "[Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)" extension, then press "<kbd>F1</kbd>" and select "`Remote-Containers: Open in Container`".

After a while, you'll get most of the environment needed to develop and debug.

## File Description

- [cobra.yaml](cobra.yaml): Default `cobra` command Settings. Used for `$ cobra add ..`
- [devcontainer.env](devcontainer.env): ENV variables to be loaded once when the container's created.
- [devcontainer.json](devcontainer.json): VSCode Extensions to be installed and env settings.
- [Dockerfile](Dockerfile): Debian 10 (buster) based Golang development container.
- [postCreateCommand.sh](postCreateCommand.sh): Initialization script that runs after the container and the VSCode server is up.
- [README.md](README.md): This file. ;-)
- [welcome.sh](welcome.sh): Bash script to display the basic info and TIPs to use in the first shell login.
