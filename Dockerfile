FROM ubuntu:16.04

RUN apt-get update && \
    apt-get install -y golang && \
    apt-get clean

ENV GOPATH=/go \
    pkg=github.com/dallasmarlow/composers

ADD . /go/src/${pkg}
RUN go install ${pkg}/composer

ENTRYPOINT ["/go/bin/composer"]
