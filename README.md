<!-- markdownlint-disable MD001 MD041 -->
[![Tests before merge](https://github.com/KEINOS/Hello-Cobra/workflows/Tests%20before%20merge/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3A%22Tests+before+merge%22 "View status on GitHub")
[![100% Coverage](https://github.com/KEINOS/Hello-Cobra/workflows/100%25%20Coverage/badge.svg)](https://github.com/KEINOS/Hello-Cobra/actions?query=workflow%3A%22100%25+Coverage%22 "View status on GitHub")

# "Hello, world!" Sample of Cobra in Golang w/ Unit Test and 100% Coverage

- [Cobra](https://cobra.dev/) is an awesome framework to help create CLI apps in Go.

This repo is a "`Hello, world!`" sample of [Cobra](https://cobra.dev/) with 100% coverage tests.

Which I couldn't find by googling with '[`"golang"` `cobra` `sample` `hello` `world` `coverage` `100%`](https://www.google.com/search?q=%22golang%22+cobra+sample+hello+world+coverage+100%)'.

This repo is for Golang and `Cobra` beginners like I am. We all know keeping 100% of coverage is a myth, but aiming to find out the **"Best practices of `Hello, world!`" sample using Cobra** to begin with.

## How to run tests

```shellsession
$ go test -cover ./...
ok    github.com/KEINOS/Hello-Cobra      0.014s  coverage: 100.0% of statements
ok    github.com/KEINOS/Hello-Cobra/cmd  0.009s  coverage: 100.0% of statements
```

- If you find hard which/where to cover when coverage was less than 100%, then try:

  ```shellsession
  $ /bin/bash ./.github/run-coverage-tests.sh --verbose
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

## Pull Request (PR)

Any PR that might help Golang newbies understand is welcome.

To evolve the sample through natural selection, if you have any better comments, suggestions, practice, etc., then don't hesitate to PR.

- If the PR passes the tests, then **it will be merged automatically**.(squash and merge)
- If you feel something against any PR, then feel free to counter PR.
  - We won't judge, as far as the test passes and not a prank-kind commit.

### Auto-merge Conditions

- Files changed/added were: `*.md` and/or `*.go`
  - If other files were changed then at least 2 approved reviews are needed to merge.
- Pass the test of "`./.github/run-merge-tests.sh`" which include the below.
  - Lint check. (`*.sh`, `*.go`)
  - Static analysis. (`*.sh`, `*.go`)
  - Unit tests of Go. (`go test ./...`)
- 100% Code Coverage (`./.github/run-coverage-tests.sh --verbose`)

### Draft PR Suggested

If you have your hands into something, then we suggest to [Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/) first and continue. So the other people would know what you are dealing with.

## Issue

In this repo, **only reports are accepted** as an issue. Such as bug report, license problem, and mal-attitude committer.

For any better suggestions, solutions, practices, typo-fix, bug-fix, etc. you have, then please PR.

## License

The repo itself is MIT License. (c) Copyright: [Hello-Cobra Contributors](https://github.com/KEINOS/Hello-Cobra/graphs/contributors).

BUT any app created from this repo as a template/boilerplate may have their own-license.

<!-- WIP
## Codespaces and VSCode Friendly

This repo works on [GitHub Codespaces](https://github.com/features/codespaces) which lets you edit/code online.

1. [Request early access](https://github.com/features/codespaces/signup) and wait to be accepted.
2. [Fork](https://docs.github.com/en/free-pro-team@latest/github/getting-started-with-github/fork-a-repo) this repo to your GitHub account.
3. Open the forked repo in GitHub and select ”`Open with Codespaces`" dropdown menu in the upper right "`↓ Code`" button.
4. Create/add a "`New codespace`" and wait until the Docker image gets built.
-->
