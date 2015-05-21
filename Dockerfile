FROM golang
ENV pkg github.com/dallasmarlow/composers
ADD . /go/src/${pkg}
RUN go install ${pkg}
ENTRYPOINT /go/bin/composers
EXPOSE 8080
