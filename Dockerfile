FROM --platform=$TARGETPLATFORM golang:alpine

MAINTAINER sntshkmr60@gmail.com

COPY dist/cotu /usr/bin/cotu

ENTRYPOINT [ "cotu" ]
