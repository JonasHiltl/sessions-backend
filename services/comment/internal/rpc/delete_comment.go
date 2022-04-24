package rpc

import (
	"context"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
)

func (s commentServer) DeleteComment(ctx context.Context, req *cg.DeleteCommentRequest) (*common.MessageResponse, error) {
	err := s.cs.Delete(ctx, req.AuthorId, req.CommentId)
	if err != nil {
		return nil, err
	}

	return &common.MessageResponse{Message: "Comment removed"}, nil
}
