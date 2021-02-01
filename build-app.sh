#!/bin/bash
# =============================================================================
#  Bash shell script to build/compile the CLI app
# =============================================================================

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
# Name of the CLI app
NAME_FILE_BIN='hello-cobra'

# Path info to export the app
PATH_DIR_SCRIPT="$(cd "$(dirname "${BASH_SOURCE:-$0}")" && pwd)"
PATH_DIR_BIN="${PATH_DIR_SCRIPT}/bin"
PATH_FILE_BIN="${PATH_DIR_BIN}/${NAME_FILE_BIN}"

# Default values
GOOS_DEFAULT='linux'
GOARCH_DEFAULT='amd64'
GOARM_DEFAULT=''

# Status
SUCCESS=0
FAILURE=1

# -----------------------------------------------------------------------------
#  Functions
# -----------------------------------------------------------------------------

echoHelp() {
    cat <<'HEREDOC'
About:
  This script builds the binary of the app under ./bin directory.

Usage:
  ./build-app.sh <GOOS>
  ./build-app.sh <GOOS> [<GOARCH>]
  ./build-app.sh <GOOS> [<GOARCH> <GOARM>]
  ./build-app.sh [-l] [--list] [-h] [--help]

GOOS:
  The fisrt argument is the OS platform. Such as:

    "linux", "darwin", "windows", etc.

  For supported platforms specify '--list' option.

GOARCH:
  The 2nd argument is the architecture (CPU type). Such as:

    "amd64"(64bit, Intel/AMD/x86_64), "arm", "arm64", "386", etc. (Default: "amd64")

  For supported architectures specify '--list' option.

GOARM:
  The 3rd argument is the ARM version. Such as:

    "5", "6", "7".(Default: empty)

  For supported combination see: https://github.com/golang/go/wiki/GoArm

Options:
  -l --list ... Displays available platforms and architectures to build.
  -h --help ... Displays this help.

Sample usage:

  # Display available arcitectures
  ./build-app.sh --list
  ./build-app.sh -l

  # Linux (Intel)
  ./build-app.sh linux

  # macOS
  ./build-app.sh darwin        #Equivalent to: ./build-app.sh darwin amd64
  ./build-app.sh darwin arm64

  # Windows10
  ./build-app.sh windows

  # Raspberry Pi 3
  ./build-app.sh linux arm 7

  # QNAP ARM5
  ./build-app linux arm 5

HEREDOC

    exit $SUCCESS
}

indentSTDIN() {
    indent='    '
    while read -r line; do
        echo "${indent}${line}"
    done
    echo
}

listPlatforms() {
    list=$(go tool dist list) || {
        echo >&2 'Failed to get supported platforms.'
        echo "$list" | indentSTDIN 1>&2
        exit $FAILURE
    }
    echo 'List of available platforms. (OS/Architecture)'
    echo "$list" | indentSTDIN
    exit $SUCCESS
}

# -----------------------------------------------------------------------------
#  Main
# -----------------------------------------------------------------------------

if [ "$#" -eq 0 ]; then
    echoHelp
fi

case "$1" in
    "--help") echoHelp ;;
    "-h") echoHelp ;;
    "--list") listPlatforms ;;
    "-l") listPlatforms ;;
esac

GOOS="${1:-"$GOOS_DEFAULT"}"
GOARCH="${2:-"$GOARCH_DEFAULT"}"
GOARM="${3:-"$GOARM_DEFAULT"}"
GOARM_SUFFIX="${GOARM:+"-${GOARM}"}"

#echo $GOOS
#echo $GOARCH
#echo $GOARM
#echo $GOARM_SUFFIX

PATH_FILE_BIN_FINAL="${PATH_FILE_BIN}-${GOOS}-${GOARCH}${GOARM_SUFFIX}"

# Build as static linked binary
echo '- Building static linked binary to ...'
echo "  ${PATH_FILE_BIN_FINAL}"

if CGO_ENABLED=0 \
    GOOS="$GOOS" \
    GOARCH="$GOARCH" \
    GOARM="$GOARM" \
    go build \
    -installsuffix "$NAME_FILE_BIN" \
    -ldflags="-s -w -extldflags \"-static\"" \
    -o="$PATH_FILE_BIN_FINAL" \
    .; then
    exit $SUCCESS
fi
echo 'Failed to build binary.'
exit $FAILURE
