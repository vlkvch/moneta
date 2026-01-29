#!/bin/bash
set -o errexit -o nounset -o pipefail
cd "`dirname $0`/.."

rm -f ./bin/moneta
rmdir ./bin
