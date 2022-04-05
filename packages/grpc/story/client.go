package story

import "google.golang.org/grpc"

func NewClient(address string) (StoryServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := NewStoryServiceClient(conn)
	return c, nil
}
