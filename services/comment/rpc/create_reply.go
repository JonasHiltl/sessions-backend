package rpc

import (
	"context"

	"github.com/go-playground/validator/v10"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/services/comment/dto"
)

func (s commentServer) CreateReply(ctx context.Context, req *cg.CreateReplyRequest) (*cg.Reply, error) {
	dr := dto.Reply{
		CommentId: req.CommentId,
		AuthorId:  req.AuthorId,
		Body:      req.Body,
	}

	v := validator.New()
	err := v.Struct(dr)
	if err != nil {
		return nil, err
	}

	r, err := s.rs.Create(ctx, dr)
	if err != nil {
		return nil, err
	}

	return r.ToGRPCReply(), nil
}
