#!/bin/bash
# =============================================================================
#  Bash shell script to build/compile the CLI app
# =============================================================================
#  This script will build the app under ./bin directory.
#
#  Usage:
#    ./build-app.sh --help
#    ./build-app.sh linux         // AMD64 by default
#    ./build-app.sh linux arm 6   // Arm v6 such as RaspberryPi Zero
#    ./build-app.sh darwin
#    ./build-app.sh darwin arm64  // For M1 mac
#    ./build-app.sh windows

# -----------------------------------------------------------------------------
#  Constants
# -----------------------------------------------------------------------------
# Name of the CLI app
NAME_FILE_BIN='hello-cobra'

# App Version from git tag
VERSION_APP="$(git describe --tag)"

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
  The 3rd argument is the ARM variant/version. Such as:

    "5", "6", "7".(Default: empty)

  For supported combination see: https://github.com/golang/go/wiki/GoArm

Options:
  -l --list ... Displays available platforms and architectures to build.
  -h --help ... Displays this help.

Sample usage:

  # Display available arcitectures
  ./build-app.sh --list
  ./build-app.sh -l

  # Build Linux (Intel) binary
  ./build-app.sh linux

  # Build macOS binary
  ./build-app.sh darwin        #Equivalent to: ./build-app.sh darwin amd64
  ./build-app.sh darwin arm64

  # Build Windows10 binary
  ./build-app.sh windows

  # Build Raspberry Pi 3 binary
  ./build-app.sh linux arm 7

  # Build QNAP ARM5 binary
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
    echo 'List of available platforms to build. (GOOS/GOARCH)'
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
echo "  Ver: ${VERSION_APP}"

if CGO_ENABLED=0 \
    GOOS="$GOOS" \
    GOARCH="$GOARCH" \
    GOARM="$GOARM" \
    go build \
    -installsuffix "$NAME_FILE_BIN" \
    -ldflags="-s -w -extldflags \"-static\" -X 'main.Version=${VERSION_APP}'" \
    -o="$PATH_FILE_BIN_FINAL" \
    ./hello-cobra/; then
    exit $SUCCESS
fi
echo 'Failed to build binary.'
exit $FAILURE
