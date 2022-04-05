package profile

import "google.golang.org/grpc"

func NewClient(address string) (ProfileServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := NewProfileServiceClient(conn)
	return c, nil
}
