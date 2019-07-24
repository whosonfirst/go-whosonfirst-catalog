FROM golang:1.12-alpine AS build-env

RUN apk update && apk upgrade \
    && apk add coreutils make

ADD . /go-whosonfirst-inspector

RUN cd /go-whosonfirst-inspector; make tools

FROM alpine

RUN apk update && apk upgrade \
    && apk add ca-certificates

COPY --from=build-env /go-whosonfirst-inspector/bin/wof-inspectord /usr/local/bin/wof-inspectord
