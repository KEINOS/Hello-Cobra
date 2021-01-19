#!/bin/bash

head -n1 </etc/issue

echo "- $(go version)"
echo "- Current path: $(pwd)"
echo "- Current user: $(whoami)"
echo "- Time: $(date)"
echo "- To run unit and coverage tests: ./.github/run-coverage-tests.sh"
echo "- To run full tests before merge: ./.github/run-merge-tests.sh"
echo "- To build binary: ./build-app.sh --help"
