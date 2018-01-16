FROM golang:alpine

MAINTAINER Austin J. Alexander <austinjalexander@nicewrk.com>

RUN apk add --update --no-cache build-base

WORKDIR /go/src/github.com/nicewrk/design-brain-api
COPY . .

RUN make

CMD design-brain-api
