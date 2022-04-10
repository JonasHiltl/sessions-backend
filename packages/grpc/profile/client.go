package profile

import "google.golang.org/grpc"

func NewClient(address string) (ProfileServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewProfileServiceClient(conn)
	return c, nil
}
