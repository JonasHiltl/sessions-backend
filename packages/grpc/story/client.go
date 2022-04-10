package story

import "google.golang.org/grpc"

func NewClient(address string) (StoryServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewStoryServiceClient(conn)
	return c, nil
}
