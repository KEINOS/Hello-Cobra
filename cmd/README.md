# Commands of Hello-Cobra

We place the command files in the same directory and level. But the actual relationships are as below.

```text
root (root.go) ..................... The main command of Hello-Cobra.
  └── hello (hello.go) ............. `hello` command is a sub-command of `root`.
      └── ext (helloExtended.go) ... `ext` command is sub-command of `hello` and grand child of `root`.
```
