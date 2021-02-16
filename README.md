<!-- markdownlint-disable MD001 MD041 -->
[![Tests before merge](https://github.com/KEINOS/Hello-Cobra/workflows/Tests%20before%20merge/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3A%22Tests+before+merge%22 "View status of the action on GitHub")
[![golangci-lint](https://github.com/KEINOS/Hello-Cobra/workflows/golangci-lint/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3Agolangci-lint "View status of the action on GitHub")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/Hello-Cobra)](https://goreportcard.com/report/github.com/KEINOS/Hello-Cobra "View on Go Report Card")
[![codecov](https://codecov.io/gh/KEINOS/Hello-Cobra/branch/main/graph/badge.svg?token=R2B9UBIEUI)](https://codecov.io/gh/KEINOS/Hello-Cobra "View details on CodeCov.IO")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/Hello-Cobra.svg)](https://pkg.go.dev/github.com/KEINOS/Hello-Cobra#section-documentation "Read generated documentation of the app")

# "Hello, world!" Sample of Cobra with 100% Code Coverage

- [Cobra](https://cobra.dev/) is an awesome framework of Golang to help create CLI apps.

**This repo is a sample of [Cobra](https://cobra.dev/)'s "`Hello, world!`" with 100% code coverage**. Including [CI](./github/workflows/)s of static analysis, lint check and formatters to just say "Hello" to the world with Cobra. It is a draft forever, so feel free to PR!

## Searching for the best practices of Cobra

**We all know keeping 100% of code coverage is a myth.** But as a `Golang` and `Cobra` beginner, we wanted a "Hello-world" sample with 100% of coverage, which couldn't be found by googling for "[`golang` `cobra` `sample` `hello` `world` `coverage` `100%`](https://www.google.com/search?q=%22golang%22+cobra+sample+hello+world+coverage+100%)".

This repo aims to be the basis for implementing best-practices by keeping the code coverage as high and complexity as low as possible.

- Therefore, this is a forever-[WIP](https://en.wikipedia.org/wiki/Work_in_progress)-repo. So any PR is welcome!!

## How to run tests

```shellsession
$ go test -cover ./...
ok      github.com/KEINOS/Hello-Cobra         0.009s  coverage: 100.0% of statements
ok      github.com/KEINOS/Hello-Cobra/cmd     0.353s  coverage: 100.0% of statements
ok      github.com/KEINOS/Hello-Cobra/util    0.015s  coverage: 100.0% of statements
```

- When coverage was less than 100% and if you find hard which/where to cover and fix, then try:

  ```shellsession
  $ # Unix/Linux users (POSIX compatible)
  $ /bin/bash ./.github/run-tests-coverage.sh --verbose
  ```

## How to run static analysis and lint for Golang

We use [GolangCI-Lint](https://golangci-lint.run/) for static analysis and lint.

```shellsession
$ golangci-lint run --config ./.github/golangci.yml ./...
...
```

- See current enabled linters: [./.github/golangci.yml](./.github/golangci.yml)

## Pull Request (PR)

Any PR that might help Golang newbies understand is welcome.

To evolve the samples through natural selection, **any PR that passes the [merge tests](https://github.com/KEINOS/Hello-Cobra/blob/main/.github/run-merge-tests.sh) will be auto-merged**. (squash and merge)

- [Auto-merge Conditions](https://github.com/KEINOS/Hello-Cobra/blob/main/.github/mergify.yml)
  - Files changed/added were: `*.md` and/or `*.go`
    - If other type of files were changed then at least 2 approved reviews are needed to merge.
  - Pass the tests of "`./.github/run-merge-tests.sh`" which includes the below.
    - Lint check. (`*.sh`, `*.go`)
    - Static analysis. (`*.sh`, `*.go`)
    - Unit tests of Go.
  - 100% Code Coverage (`./.github/run-coverage-tests.sh --verbose`)

So, if you have any better ideas, suggestions, practice, etc., then don't hesitate to PR. And if you feel something against any PR, then feel free to counter PR.

- **We will not judge which is better**, as far as the merge-tests passes and not a prank-kind commit.
- To avoid conflict, don't forget to move on with the latest `main/master` branch before any commit.

### Draft PR Suggested

If you have your hands on something, then we suggest to [Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/) first and continue. So the other people would know what you are dealing with.

## Questions

If you have any questions about this repo or "`Cobra`", then let us know in the [Discussions](https://github.com/KEINOS/Hello-Cobra/discussions) and find it out together.

- https://github.com/KEINOS/Hello-Cobra/discussions

## Issue

In this repo, **only bug-kind-reports are acceptable as an issue**. Such as bug report, vulnerability, license problem, and reporting mal-attitude committer.

- https://github.com/KEINOS/Hello-Cobra/issues

## License

The repo itself is MIT License. (c) Copyright, [Hello-Cobra Contributors](https://github.com/KEINOS/Hello-Cobra/graphs/contributors).

BUT any app created from this repo as a template/boilerplate may have its willing license.

<!-- WIP
## Codespaces and VSCode Friendly

This repo works on [GitHub Codespaces](https://github.com/features/codespaces) which lets you edit/code online.

1. [Request early access](https://github.com/features/codespaces/signup) and wait to be accepted.
2. [Fork](https://docs.github.com/en/free-pro-team@latest/github/getting-started-with-github/fork-a-repo) this repo to your GitHub account.
3. Open the forked repo in GitHub and select the ”`Open with Codespaces`" dropdown menu in the upper right "`↓ Code`" button.
4. Create/add a "`New codespace`" and wait until the Docker image gets built.
-->
