FROM golang:1.17.1 as builder

MAINTAINER "Xander Guzman <xander.guzman@xanderguzman.com>"

WORKDIR /go/src/github.com/theshadow/hive
COPY . .

RUN make


