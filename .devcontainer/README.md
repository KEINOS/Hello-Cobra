<!-- markdownlint-disable MD033 -->
# Dockerfile for GitHub and VSCode Users To Develop

This directory is for [GitHub Codespaces](https://github.com/features/codespaces) and/or [VS Code + Docker](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) users for development.

It includes:

- Latest [Alpine-based container image](https://github.com/KEINOS/VSCode-Dev-Container-Go/pkgs/container/vscode-dev-container-go) with vulnerability and security scan via Synk, Dockle and Trivy.
- Latest Go version.
- Common packages and tools for developing Golang app and shell scripts.
  - [Installed tools and Go modules](https://github.com/KEINOS/VSCode-Dev-Container-Go/blob/main/image_info.txt) | VSCode-Dev-Container-Go | KEINOS @ GitHub
- Weekly updated.

## Developing Online

Just press the `.`(dot) key on GitHub and you should be redirected to [GitHub Codespaces](https://github.com/features/codespaces). (You may need to register to use Codesspaces)

## VS Code + Docker User

The container contains VS Code Server as well. If you already have installed the "[Remote - Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)" extension, then press "<kbd>F1</kbd>" and select "`Remote-Containers: Open in Container`".

After a while, you'll get most of the environment needed to develop and debug.
