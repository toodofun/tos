FROM --platform=${BUILDPLATFORM} node:20.12.2-bullseye AS node-builder

WORKDIR /build

COPY frontend/package.json frontend/yarn.lock ./

RUN yarn config set registry 'https://registry.npmmirror.com' && \
    yarn install

COPY frontend .

RUN yarn build-only

FROM golang:1.23.0-alpine3.20 AS go-builder

WORKDIR /build

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates make git zip && \
    go env -w GOPROXY=https://goproxy.cn,direct

COPY server/Makefile server/go.mod server/go.sum ./

RUN make deps

COPY server .
COPY VERSION /
COPY --from=node-builder /build/dist ./server/static

ARG COMMIT_ID
RUN make deps && make build COMMIT=${COMMIT_ID} 

FROM alpine:3.20
WORKDIR /app

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
    apk add --no-cache ca-certificates bash

COPY --from=go-builder /build/bin/tos /app/


EXPOSE 61296
ENTRYPOINT ["/app/tos"]