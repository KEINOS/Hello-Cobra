#!/bin/sh
# =============================================================================
#  This script updates Go modules to the latest version.
# =============================================================================
#  It will remove the go.mod file and run `go mod tidy` to get the latest moule
#  versions.
#  Then it will run the tests to make sure the code is still working, and fails
#  if any errors are found during the process.
#
#  NOTE: This script is aimed to run in the container via docker-compose.
#    See "tidy" service: ./docker-compose.yml
# =============================================================================

set -eu

echo '* Backup modules ...'
mv go.mod go.mod.bak
mv go.sum go.sum.bak

echo '* Create new blank go.mod ...'
< go.mod.bak head -n 4 > go.mod

echo '* Run go tidy ...'
go mod tidy

echo '* Run tests ...'
go test ./... && {
    echo '* Testing passed. Removing old go.mod file ...'
    rm -f go.mod.bak
    rm -f go.sum.bak
    echo 'Successfully updated modules!'
}
