FROM golang:1.10

EXPOSE 50051
WORKDIR /Users/fguy/go/src/example-service
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["greeter-server"]
