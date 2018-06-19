FROM golang:1.10

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

ADD . $GOPATH/src/github.com/backendservice/example-service

WORKDIR $GOPATH/src/github.com/backendservice/example-service

RUN go get -u github.com/golang/dep/...
RUN dep ensure

EXPOSE 50051
RUN go build -o server ./greeter_server
CMD ["./server"]
