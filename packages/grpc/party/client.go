package party

import "google.golang.org/grpc"

func NewClient(address string) (PartyServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewPartyServiceClient(conn)
	return c, nil
}
