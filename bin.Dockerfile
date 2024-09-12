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
    apk add --no-cache ca-certificates make git && \
    go env -w GOPROXY=https://goproxy.cn,direct

COPY server/Makefile server/go.mod server/go.sum ./

RUN make deps

COPY server .
COPY VERSION /
COPY --from=node-builder /build/dist ./server/static

RUN make deps && make all

FROM scratch
COPY --from=go-builder /build/bin /