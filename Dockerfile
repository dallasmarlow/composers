FROM debian:latest AS build

RUN apt-get update && \
    apt-get install -y golang && \
    apt-get clean

ENV DEBIAN_FRONTEND=noninteractive \
    GOARCH=arm64 \
    GOPATH=/go \
    pkg=github.com/dallasmarlow/composers

COPY . /go/src/${pkg}
WORKDIR /usr/local/bin
RUN go build ${pkg}/composer

FROM debian:latest
COPY --from=build /usr/local/bin/composer /usr/local/bin/composer
ENTRYPOINT ["/usr/local/bin/composer"]
