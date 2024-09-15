#!/usr/bin/env bash

CURRENT_DIR=$(cd `dirname $0`; pwd)

set -e

rm -rf ${CURRENT_DIR}/_output

COMMIT_ID=$(git rev-parse --short HEAD)

# 构建镜像
docker buildx build -f Dockerfile --platform linux/amd64,linux/arm64 -t toodo/tos:latest --build-arg COMMIT_ID=${COMMIT_ID} . --push

# 构建二进制
docker buildx build -f bin.Dockerfile --output type=local,dest=_output --build-arg COMMIT_ID=${COMMIT_ID} .
