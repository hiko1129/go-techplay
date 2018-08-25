FROM golang:1.10-alpine

RUN apk add --no-cache git \
    curl \
    make

WORKDIR /go/src/bitbucket.org/hiko1129/go-techplay
COPY . .

RUN curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
ENV PATH $PATH:/go/bin
RUN dep ensure