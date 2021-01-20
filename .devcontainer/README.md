<!-- markdownlint-disable MD033 -->
# Dockerfile for CIs and Development

## For CIs

The [CI](https://en.wikipedia.org/wiki/Continuous_integration) will run the [Dockerfile](Dockerfile) to build the image and then run the tests in a container.

- Current CI
  - [GitHub Actions](https://docs.github.com/en/free-pro-team@latest/actions): [../.github/workflows/](https://github.com/KEINOS/Hello-Cobra/tree/main/.github/workflows)

## For DEVs

This directory is for [GitHub Codespaces](https://github.com/features/codespaces) and [VS Code + Docker](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) users to develop the cloned/forked repo.

### VS Code + Docker User

If you have already installed the [Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension, then press "<kbd>F1</kbd>" and select "`Remote-Containers: Open in Container`". After a while, you'll get most of the environment needed to develop and/or debug.

- File Description
  - [cobra.yaml](cobra.yaml): Default `cobra` command Settings.
  - [devcontainer.json](devcontainer.json): VSCode Extensions to be installed.
  - [Dockerfile](Dockerfile): Alpine based Golang development container.
  - [postCreateCommand.sh](postCreateCommand.sh): Initialization script that runs after the container and the VSCode server is up and before connection from VSCode.
  - [settings.vscode.json](settings.vscode.json): Additional VSCode Settings.
  - [welcome.sh](welcome.sh): Bash script that pre-runs to display the welcome message in the terminal. It will display basic info and TIPs to use.
