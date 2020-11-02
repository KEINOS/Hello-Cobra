# "Hello, world!" Sample of Test and 100% Coverage of Cobra in Golang

- [Cobra](https://cobra.dev/) is a framework to help create CLI apps in Go.

This repo is a "`Hello, world!`" sample of [Cobra](https://cobra.dev/) with a 100% coverage tests.

Which I couldn't find by googling with '[`"golang"` `cobra` `sample` `hello` `world` `coverage` `100%`](https://www.google.com/search?q=%22golang%22+cobra+sample+hello+world+coverage+100%)'.

This repo is for Golang and `Cobra` beginners, like I am. We all know keeping 100% of coverage is a myth, but aiming to find out the **"Best practices of `Hello, world!`" sample using Cobra** to begin with.

## Pull Request (PR)

Any PR that might help understand the Golang newbie are welcome.

To evolve the sample through natural selection, if you have any better suggestion or practice or etc., please don't issue them but PR.

If the PR pass the tests then I will merge it almost automaticaly. If you feel something against to any PR then counter PR.

### Suggested PR

If you have your hands into something, we suggest to [Draft PR](https://github.blog/2019-02-14-introducing-draft-pull-requests/) first and continue, so the other people would know.

### Conditions to merge

- Files changed were:
  - `*.md`, `*.go` and/or `go.mod`
  - If other files were changed then it will be reviewed to merge.
- Pass the "`run-merge-tests.sh`" which includes the below.
    - Lint check. (*.sh, *.go)
    - Static analysis. (*.sh, *.go)
    - Unit test of Go with 100% of coverage. (`run-tests.sh`)

## Issue

In this repo, **only reports are accepted** as an issue. Such as bug report, license problem and mal-attitude committer.

For any better suggestions, solutions, practices, typo-fix, bug-fix or etc. you have, please PR.
