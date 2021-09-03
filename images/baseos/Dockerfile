# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER
ARG ALPINE_VER

FROM alpine:${ALPINE_VER} as base

# 换源加速镜像下载
RUN set -eux && sed -i 's/dl-cdn.alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' /etc/apk/repositories

RUN apk add --no-cache tzdata
RUN addgroup -g 500 chaincode && adduser -u 500 -D -h /home/chaincode -G chaincode chaincode
USER chaincode
