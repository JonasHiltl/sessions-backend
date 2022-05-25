package rpc

import (
	"context"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
)

func (s commentServer) DeleteReply(ctx context.Context, req *cg.DeleteReplyRequest) (*common.MessageResponse, error) {
	err := s.rs.Delete(ctx, req.AuthorId, req.CommentId, req.ReplyId)
	if err != nil {
		return nil, err
	}

	return &common.MessageResponse{Message: "Reply removed"}, nil
}
