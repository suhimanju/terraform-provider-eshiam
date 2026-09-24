#!/usr/bin/env bash
# Install a pinned version of golangci-lint into ./bin.
set -euo pipefail

VERSION="${1:-v1.61.0}"

curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh \
  | sh -s -- -b "$(pwd)/bin" "$VERSION"

./bin/golangci-lint --version
