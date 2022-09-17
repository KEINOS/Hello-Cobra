<!-- markdownlint-disable MD001 MD041 -->
[![go1.15+](https://img.shields.io/badge/Go-1.15,%2016,%2017,%20latest-blue?logo=go)](https://github.com/KEINOS/dev-go/actions/workflows/go-versions.yml "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/Hello-Cobra.svg)](https://pkg.go.dev/github.com/KEINOS/Hello-Cobra#section-documentation "Read generated documentation of the app")

# "Hello, world!" example of Cobra with 100% code coverage

> __Note__ : "`Cobra`" is an awesome framework of Go to help create CLI apps.

**This repo is a example of [Cobra](https://cobra.dev/)'s "`Hello, world!`" with 100% code coverage**.

Including the below to just say "Hello" to the world!

- [CI](./github/workflows/)
    - Unit tests on Go v1.15 to the latest.
    - Static analysis, security scan, lint and format check.
    - Monthly vulnerability scan via CodeQL.
    - Automated monthly update of `go.mod` and `go.sum` on test-pass.
- Automated [Homebrew release](https://github.com/KEINOS/homebrew-Hello-Cobra) on release push.

## Searching for the best practices of Cobra

**We all know keeping 100% of code coverage is a myth.** But as a `Golang` and `Cobra` beginner, we wanted a "Hello-world" example with 100% of coverage, which couldn't be found by _googling_ for "[`golang` `cobra` `sample` `example` `hello` `world` `coverage` `100%`](https://www.google.com/search?q=%22golang%22+cobra+sample+example+hello+world+coverage+100%)".

This repo aims to find the best-practices of `Cobra` by refactoring it on a moment-to-moment basis. But keeping the code coverage as high as possible and less complexity as possible.

- Therefore, this is a forever-[WIP](https://en.wikipedia.org/wiki/Work_in_progress)-repo.<br>So any PR for the better is welcome!!  We will merge it as long as it passess the tests with 100% coverage and not a prank-kind PR.

## Statuses

[![Test on macOS/Win/Linux](https://github.com/KEINOS/Hello-Cobra/actions/workflows/platform-test.yaml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/platform-test.yaml)
[![go1.15+](https://github.com/KEINOS/Hello-Cobra/actions/workflows/version-tests.yaml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/version-tests.yaml)
[![golangci-lint](https://github.com/KEINOS/Hello-Cobra/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/golangci-lint.yaml)
[![codecov](https://codecov.io/gh/KEINOS/Hello-Cobra/branch/main/graph/badge.svg?token=R2B9UBIEUI)](https://codecov.io/gh/KEINOS/Hello-Cobra "View details on CodeCov.IO")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/Hello-Cobra)](https://goreportcard.com/report/github.com/KEINOS/Hello-Cobra "View on Go Report Card")
[![CodeQL](https://github.com/KEINOS/Hello-Cobra/actions/workflows/codeQL-analysis.yaml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/codeQL-analysis.yaml "Vulnerability Scan")

## Note

- This example app supports Homebrew installation. To know how it works see: [.goreleaser.yml](./.goreleaser.yml)
    ```bash
    brew install KEINOS/Hello-Cobra/hello-cobra
    ```
    ```shellsession
    $ brew install KEINOS/Hello-Cobra/hello-cobra
    ==> Tapping keinos/hello-cobra
    ... **snip** ...
    ==> Installing hello-cobra from keinos/hello-cobra
    ... **snip** ...

    $ hello-cobra --version
    hello-cobra version 1.3.2-alpha (c3c9eab)

    $ hello-cobra hello foo bar
    Hello, foo and bar!

    $ hello-cobra hello foo bar --reverse
    !rab dna oof ,olleH
    ```
- This package auto-detects the app version from the `git` tag if the app was installed via `go install` (on Go v1.16+) or `go get -u` (on Go 1.15). Try:
    ```bash
    # For Go 1.16 or above
    cd /tmp
    go install "github.com/KEINOS/Hello-Cobra/hello-cobra@latest"
    hello-cobra --version
    # Output: hello-cobra version v1.3.0
    ```
    ```bash
    # For Go 1.15
    cd /tmp
    GO111MODULE="on" go get -u "github.com/KEINOS/Hello-Cobra/hello-cobra@latest"
    hello-cobra --version
    # Output: hello-cobra version v1.3.0
    ```
- This repo is [GitHub Codespaces](https://github.com/features/codespaces) compatible. Press the `.`(dot) key to open VSCode online. (You may need to register to use Codespaces)
- This repo is [VS Code + Docker](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) compatible as well. See the [./.devcontainer/README.md](./.devcontainer/README.md) for more details.
- This repo updates monthly the `go.mod` and `go.sum` files if all the tests succeeds to run in all Go versions (Go v1.15~latest).

## License

The repo itself is MIT License. (c) Copyright, [Hello-Cobra Contributors](https://github.com/KEINOS/Hello-Cobra/graphs/contributors).

BUT any app created from this repo as a template/boilerplate may have its willing license.
