package relation

import "google.golang.org/grpc"

func NewClient(address string) (RelationServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewRelationServiceClient(conn)
	return c, nil
}
