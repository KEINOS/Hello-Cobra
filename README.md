<!-- markdownlint-disable MD001 MD041 -->
[![go1.15+](https://img.shields.io/badge/Go-1.15,%2016,%2017,%20latest-blue?logo=go)](https://github.com/KEINOS/dev-go/actions/workflows/go-versions.yml "Supported versions")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/Hello-Cobra.svg)](https://pkg.go.dev/github.com/KEINOS/Hello-Cobra#section-documentation "Read generated documentation of the app")

# "Hello, world!" Sample of Cobra with 100% Code Coverage

- [Cobra](https://cobra.dev/) is an awesome framework of Go to help create CLI apps.

**This repo is a sample of [Cobra](https://cobra.dev/)'s "`Hello, world!`" with 100% code coverage**.

Including [CI](./github/workflows/)s of static analysis, security scan, lint check and formatters to just say "Hello" to the world with Cobra. It is a draft forever, so feel free to PR!

## Searching for the best practices of Cobra

**We all know keeping 100% of code coverage is a myth.** But as a `Golang` and `Cobra` beginner, we wanted a "Hello-world" sample with 100% of coverage, which couldn't be found by googling for "[`golang` `cobra` `sample` `hello` `world` `coverage` `100%`](https://www.google.com/search?q=%22golang%22+cobra+sample+hello+world+coverage+100%)".

This repo aims to implement best-practices of `Cobra` but keeping the code coverage as high as possible and less complexity.

- Therefore, this is a forever-[WIP](https://en.wikipedia.org/wiki/Work_in_progress)-repo. So any PR for the better is welcome!!

## Statuses

[![Test on macOS/Win/Linux](https://github.com/KEINOS/Hello-Cobra/actions/workflows/platform-test.yaml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/platform-test.yaml)
[![go1.15+](https://github.com/KEINOS/Hello-Cobra/actions/workflows/version-tests.yaml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/version-tests.yaml)
[![golangci-lint](https://github.com/KEINOS/Hello-Cobra/actions/workflows/golangci-lint.yaml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/golangci-lint.yaml)
[![codecov](https://codecov.io/gh/KEINOS/Hello-Cobra/branch/main/graph/badge.svg?token=R2B9UBIEUI)](https://codecov.io/gh/KEINOS/Hello-Cobra "View details on CodeCov.IO")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/Hello-Cobra)](https://goreportcard.com/report/github.com/KEINOS/Hello-Cobra "View on Go Report Card")
[![CodeQL](https://github.com/KEINOS/Hello-Cobra/actions/workflows/codeQL-analysis.yml/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions/workflows/codeQL-analysis.yml "Vulnerability Scan")

## Note

- This repo is [GitHub Codespaces](https://github.com/features/codespaces) compatible. Press the `.`(dot) key to open VSCode online. (You may need to register to use Codespaces)
- This repo is [VS Code + Docker](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.vscode-remote-extensionpack) compatible as well. See the [./.devcontainer/README.md](./.devcontainer/README.md) for more details.
- This repo updates monthly the `go.mod` and `go.sum` files if all the tests succeeds to run in all Go versions (Go v1.15~latest).

## License

The repo itself is MIT License. (c) Copyright, [Hello-Cobra Contributors](https://github.com/KEINOS/Hello-Cobra/graphs/contributors).

BUT any app created from this repo as a template/boilerplate may have its willing license.
