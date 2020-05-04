FROM ubuntu:latest AS build

ENV DEBIAN_FRONTEND=noninteractive \
    GOPATH=/go \
    pkg=github.com/dallasmarlow/composers

RUN apt-get update && \
    apt-get install -y golang && \
    apt-get clean

COPY . /go/src/${pkg}
RUN go install ${pkg}/composer

FROM ubuntu:latest
COPY --from=build /go/bin/composer /usr/local/bin/composer
ENTRYPOINT ["/usr/local/bin/composer"]
