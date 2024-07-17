# syntax=docker/dockerfile:1
FROM golang:alpine AS builder


ENV CGO_ENABLED 0
ENV GOOS linux
ENV GO111MODULE on
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o /bin/main .



FROM alpine:latest
RUN apk add --no-cache tzdata

ENV TZ="Asia/Shanghai"

ENV WORKDIR     /app
COPY --from=builder /bin/main $WORKDIR/
RUN chmod +x $WORKDIR/main


WORKDIR $WORKDIR
CMD ./main