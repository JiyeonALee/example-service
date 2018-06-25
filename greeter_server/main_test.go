package main

import (
	"context"
	"testing"

	pb "github.com/backendservice/example-service/helloworld"
	"github.com/stretchr/testify/assert"
)

func TestSayHello(t *testing.T) {
	ctx := context.Background()
	req := &pb.HelloRequest{
		Name: "Taehoon",
	}
	instance := new(server)
	res, err := instance.SayHello(ctx, req)
	assert.NoError(t, err)

	assert.Equal(t, "Hello Taehoon", res.Message)
}
