package main

import (
	"errors"
	"testing"
	"time"

	pb "github.com/backendservice/example-service/helloworld"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"

	"github.com/backendservice/example-service/mock"
)

func TestSayHello(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock_helloworld.NewMockGreeterClient(ctrl)

	expected := &pb.HelloReply{
		Message: "Hello World",
	}
	mockClient.EXPECT().SayHello(gomock.Any(), gomock.Any()).Return(expected, nil).Times(1)

	res, err := SayHello(mockClient, context.Background(), time.Second, "World")
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", *res)
}

func TestSayHello_Error(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockClient := mock_helloworld.NewMockGreeterClient(ctrl)

	mockClient.EXPECT().SayHello(gomock.Any(), gomock.Any()).Return(nil, errors.New("")).Times(1)

	res, err := SayHello(mockClient, context.Background(), time.Second, "World")
	assert.Error(t, err)
	assert.Nil(t, res)
}
