package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
)

func (s *commentServer) DeleteComment(ctx context.Context, req *cg.GetCommentRequest) (*common.MessageResponse, error) {
	me, err := middleware.ParseUser(ctx)
	if err != nil {
		return nil, err
	}

	err = s.cs.Delete(ctx, me.Sub, req.PId, req.CId)
	if err != nil {
		return nil, err
	}

	return &common.MessageResponse{Message: "Comment removed"}, nil
}
