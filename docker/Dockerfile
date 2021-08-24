FROM golang:1.16-alpine as builder


WORKDIR $GOPATH/src

RUN apk update && \
    apk add --no-cache \
    alpine-sdk \
    git

ENV GO111MODULE=on