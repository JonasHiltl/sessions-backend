package auth

import (
	"google.golang.org/grpc"
)

func NewClient(address string) (AuthServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewAuthServiceClient(conn)
	return c, nil
}
