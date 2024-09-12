#!/usr/bin/env bash

CURRENT_DIR=$(cd `dirname $0`; pwd)

set -e

rm -rf ${CURRENT_DIR}/_output
docker buildx build -f bin.Dockerfile --output type=local,dest=_output .