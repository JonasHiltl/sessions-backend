package auth

import "google.golang.org/grpc"

func NewClient(address string) (AuthServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := NewAuthServiceClient(conn)
	return c, nil
}
