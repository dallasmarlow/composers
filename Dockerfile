FROM golang
ENV pkg github.com/dallsamarlow/composers
ADD . /go/src/${pkg}
RUN go install ${pkg}
ENTRYPOINT /go/bin/composers
EXPOSE 8080