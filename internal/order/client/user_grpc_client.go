package client

import (
	"context"
	"time"

	userpb "github.com/abdillahfazri/grpc-simple/proto/user/v1"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type UserClient struct {
	client userpb.UserServiceClient
}

func NewUserClient(addr string) (*UserClient, error) {
	conn, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, err
	}

	return &UserClient{
		client: userpb.NewUserServiceClient(conn),
	}, nil
}

func (c *UserClient) GetUser(ctx context.Context, id string) (*userpb.GetUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	return c.client.GetUser(ctx, &userpb.GetUserRequest{
		Id: id,
	})
}
