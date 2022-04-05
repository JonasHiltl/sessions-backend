package party

import "google.golang.org/grpc"

func NewClient(address string) (PartyServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := NewPartyServiceClient(conn)
	return c, nil
}
