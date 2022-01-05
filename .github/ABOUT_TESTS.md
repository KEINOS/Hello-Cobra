# Battery of tests for merging

```bash
.
├── Dockerfile ................ Alpine-base image for docker-compose.
├── SECURITY.md ............... Security policy of this repo.
├── check-requirements.sh ..... Shell script to check requirements to run
│                               `run-tests-merge.sh`.
├── docker-compose.yml ........ Docker-compose file to run the tests in various
│                               contidions.
├── mergify.yml ............... Configuration for auto-merge PRs using
│                               Mergify.com service.
├── run-tests-coverage.sh ..... Shell script to run go-carpet for displaying
│                               where to cover (if less than 100%).
├── run-tests-go-get-u.sh ..... Shell script to test app installation via
│                               `go get -u` in Go 1.15 (monthly run).
├── run-tests-go-install.sh ... Shell script to test app installation via
│                               `go install` in Go 1.16+ (monthly run).
├── run-tests-lint.sh ......... Shall script to run `golangci-lint`.
├── run-tests-merge.sh ........ Battery of tests which will run static analysis
│                               and linters for both shell script and Go source.
├── update-go-mod.sh .......... Updates the "go.mod" and "go.sum" files to the
│                               latest.
└── workflows/ ................ Directory for CIs via GitHub Actions.
                                See: ./workflows/README.md
```

## To test via Docker (docker-compose)

- All-in-one (Go: latest version, Unit Test, Lint check and static analysis)

    ```bash
    docker-compose --file ./.github/docker-compose.yml run mergeability
    ```

- Unit test on various Go versions

    ```bash
    docker-compose --file ./.github/docker-compose.yml run v1_15
    docker-compose --file ./.github/docker-compose.yml run v1_16
    docker-compose --file ./.github/docker-compose.yml run v1_17
    docker-compose --file ./.github/docker-compose.yml run latest
    ```

- Check code coverage

    ```bash
    docker-compose --file ./.github/docker-compose.yml run coverage
    ```
