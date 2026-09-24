#!/usr/bin/env bash
# Run integration tests against a live tenant.
#
# Required environment variables:
#   EXAMPLE_HOST           API base URL (e.g. https://tenant.api.example.com)
#   EXAMPLE_CLIENT_ID      OAuth2 client id
#   EXAMPLE_CLIENT_SECRET  OAuth2 client secret
#
# Usage: ./integration_tests.sh
set -euo pipefail

: "${EXAMPLE_HOST:?EXAMPLE_HOST must be set}"
: "${EXAMPLE_CLIENT_ID:?EXAMPLE_CLIENT_ID must be set}"
: "${EXAMPLE_CLIENT_SECRET:?EXAMPLE_CLIENT_SECRET must be set}"

export TF_ACC=1

go test -v -tags=integration -timeout 120m ./internal/...
