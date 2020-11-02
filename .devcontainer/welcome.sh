#!/bin/sh

head -n1 < /etc/issue

echo "- $(go version)"
echo "- Current path: $(pwd)"
echo "- Current user: $(whoami)"
echo "- Time: $(date)"
echo "- To run tests: ./run-tests.sh"
echo "- To run tests before merge: ./run-merge-tests.sh"
