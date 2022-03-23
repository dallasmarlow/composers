FROM debian:bullseye-slim AS build

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -y golang && \
    apt-get clean

COPY . /opt
WORKDIR /opt
RUN go install composer/*

FROM debian:bullseye-slim
COPY --from=build /root/go/bin/composer /usr/local/bin/composer
ENTRYPOINT ["/usr/local/bin/composer"]
