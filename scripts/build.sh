#!/bin/bash

set -o errexit -o nounset -o pipefail
cd "`dirname $0`/.."

go build -ldflags "-s -w" -o moneta ./cmd/cli
