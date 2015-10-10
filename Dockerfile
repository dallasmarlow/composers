FROM golang
ENV pkg github.com/dallasmarlow/composers
ADD . /go/src/${pkg}
RUN go install ${pkg}/composer
ENTRYPOINT /go/bin/composer
EXPOSE 8080
