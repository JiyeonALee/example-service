package main

import (
	"log"
	"os"
	"time"

	pb "github.com/backendservice/example-service/helloworld"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	//	address     = "ec2-18-191-204-27.us-east-2.compute.amazonaws.com:50051"
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	message, err := SayHello(c, context.Background(), time.Minute, name)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", *message)
}

func SayHello(c pb.GreeterClient, ctx context.Context, timeout time.Duration, name string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return nil, err
	}
	return &r.Message, nil
}
