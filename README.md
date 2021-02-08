<!-- markdownlint-disable MD001 MD041 -->
[![Tests before merge](https://github.com/KEINOS/Hello-Cobra/workflows/Tests%20before%20merge/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3A%22Tests+before+merge%22 "View status of the action on GitHub")
[![100% Coverage](https://github.com/KEINOS/Hello-Cobra/workflows/100%25%20Coverage/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3A%22100%25+Coverage%22 "View status of the action on GitHub")
[![golangci-lint](https://github.com/KEINOS/Hello-Cobra/workflows/golangci-lint/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3Agolangci-lint "View status of the action on GitHub")
[![Go Report Card](https://goreportcard.com/badge/github.com/KEINOS/Hello-Cobra)](https://goreportcard.com/report/github.com/KEINOS/Hello-Cobra "View on Go Report Card")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/Hello-Cobra.svg)](https://pkg.go.dev/github.com/KEINOS/Hello-Cobra#section-documentation "Read generated documentation of the app")

# "Hello, world!" Sample of Cobra with Unit Tests and 100% Coverage

- [Cobra](https://cobra.dev/) is an awesome framework to help create CLI apps in Golang.

This repo is a draft of [Cobra](https://cobra.dev/)'s "`Hello, world!`" sample, with 100% coverage tests.

Which I couldn't find by googling with "[`golang` `cobra` `sample` `hello` `world` `coverage` `100%`](https://www.google.com/search?q=%22golang%22+cobra+sample+hello+world+coverage+100%)".

## Searching for the best practices of Cobra

**We all know keeping 100% of coverage is a myth.** But as a `Golang` and `Cobra` beginner, we wanted a "Hello-world" sample with 100% code coverage as a basis for implementing best-practices and keeping the coverage high and the complexity low as possible.

- Therefore this is a Forever-[WIP](https://en.wikipedia.org/wiki/Work_in_progress)-Repo. So any PR is welcome!!

## How to run tests

```shellsession
$ go test -cover ./...
ok      github.com/KEINOS/Hello-Cobra         0.009s  coverage: 100.0% of statements
ok      github.com/KEINOS/Hello-Cobra/cmd     0.353s  coverage: 100.0% of statements
ok      github.com/KEINOS/Hello-Cobra/util    0.015s  coverage: 100.0% of statements
```

- If you find hard which/where to cover when coverage was less than 100%, then try:

  ```shellsession
  $ /bin/bash ./.github/run-tests-coverage.sh --verbose
  ...
  ```

## How to run static analysis and lint for Golang

We use [GolangCI-Lint](https://golangci-lint.run/), an aggregator of Go linters.

```shellsession
$ golangci-lint run --config ./.github/golangci.yml ./...
...
```

- See current enabled linters: [./.github/golangci.yml](./.github/golangci.yml)

## Pull Request (PR)

Any PR that might help Golang newbies understand is welcome.

To evolve the samples in this repo through natural selection, any PR that passes the [merge tests](https://github.com/KEINOS/Hello-Cobra/blob/main/.github/run-merge-tests.sh) will be auto-merged. (squash and merge)

- [Auto-merge Conditions](https://github.com/KEINOS/Hello-Cobra/blob/main/.github/mergify.yml)
  - Files changed/added were: `*.md` and/or `*.go`
    - If other type of files were changed then at least 2 approved reviews are needed to merge.
  - Pass the test of "`./.github/run-merge-tests.sh`" which include the below.
    - Lint check. (`*.sh`, `*.go`)
    - Static analysis. (`*.sh`, `*.go`)
    - Unit tests of Go. (`go test ./...`)
  - 100% Code Coverage (`./.github/run-coverage-tests.sh --verbose`)

So, if you have any better ideas, suggestions, practice, etc., then don't hesitate to PR. And if you feel something against any PR, then feel free to counter PR.

- **We will not judge which is better**, as far as the merge-tests passes and not a prank-kind commit.
- Don't forget to move on with the latest `main/master` branch before any commitment to avoid conflict.

### Draft PR Suggested

If you have your hands on something, then we suggest to [Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/) first and continue. So the other people would know what you are dealing with.

## Questions

If you have any questions about HOW-TOs on "Cobra", then let us find how in the [Discussions](https://github.com/KEINOS/Hello-Cobra/discussions) together.

## Issue

In this repo, **only bug-kind-reports are acceptable as an issue**. Such as bug report, vulnerability, license problem, and mal-attitude committer.

For any suggestions, solutions, practices, typo-fix, bug-fix, and else to improve the sample, please [PR](https://github.com/KEINOS/Hello-Cobra/pulls)!

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
