package comment

import "google.golang.org/grpc"

func NewClient(address string) (CommentServiceClient, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := NewCommentServiceClient(conn)
	return c, nil
}
