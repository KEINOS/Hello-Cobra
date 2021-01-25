<!-- markdownlint-disable MD001 MD041 -->
[![Tests before merge](https://github.com/KEINOS/Hello-Cobra/workflows/Tests%20before%20merge/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3A%22Tests+before+merge%22 "View status on GitHub")
[![100% Coverage](https://github.com/KEINOS/Hello-Cobra/workflows/100%25%20Coverage/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3A%22100%25+Coverage%22 "View status on GitHub")
[![Go Reference](https://pkg.go.dev/badge/github.com/KEINOS/Hello-Cobra.svg)](https://pkg.go.dev/github.com/KEINOS/Hello-Cobra)

# "Hello, world!" Sample of Cobra in Golang w/ Unit Test and 100% Coverage

- [Cobra](https://cobra.dev/) is an awesome framework to help create CLI apps in Go.

This repo is a "`Hello, world!`" sample of [Cobra](https://cobra.dev/) with 100% coverage tests.

Which I couldn't find by googling with '[`"golang"` `cobra` `sample` `hello` `world` `coverage` `100%`](https://www.google.com/search?q=%22golang%22+cobra+sample+hello+world+coverage+100%)'.

This repo is for Golang and `Cobra` beginners like I am. We all know keeping 100% of coverage is a myth, but aiming to find out the **"Best practices of `Hello, world!`" sample using Cobra**, to begin with.

## How to run tests

```shellsession
$ go test -cover ./...
ok    github.com/KEINOS/Hello-Cobra      0.014s  coverage: 100.0% of statements
ok    github.com/KEINOS/Hello-Cobra/cmd  0.009s  coverage: 100.0% of statements
```

- If you find hard which/where to cover when coverage was less than 100%, then try:

  ```shellsession
  $ /bin/bash ./.github/run-tests-coverage.sh --verbose
  ...
  ```

## How to run static analysis

`vet` is a Go tool for static analysis of Go programs.

```shellsession
$ # Run all analyzers of Go vet.
$ go vet ./...

$ # View exit status (0=success, else=failure)
$ echo $?
0

$ # For more details about Go vet tool
$ # see: https://golang.org/cmd/vet/
$ # or type:
$ go tool vet help
...
```

## How to generate tests

[`gotests`](https://github.com/cweill/gotests) can generate Go tests for specific source files or an entire directory from the command line.

```shellsession
$ gotests [options] PATH [PATH] ...
...
```

- VSCode users
  - [generate-unit-tests](https://github.com/golang/vscode-go/blob/master/docs/features.md#generate-unit-tests) | Go for VS Code | Golang @ GitHub

## Pull Request (PR)

Any PR that might help Golang newbies understand is welcome.

To evolve the samples in the repo through natural selection, any PR that passes the [merge tests](https://github.com/KEINOS/Hello-Cobra/blob/main/.github/run-merge-tests.sh) will be auto-merged. (squash and merge)

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
- Don't forget to move on with the latest `main/master` branch before any commit to avoid conflict.

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
