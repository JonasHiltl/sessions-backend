package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/comment"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/dto"
)

func (s *commentServer) CreateComment(ctx context.Context, req *cg.CreateCommentRequest) (*cg.Comment, error) {
	me, err := middleware.ParseUser(ctx)
	if err != nil {
		return nil, err
	}

	dc := dto.Comment{
		PId:  req.PId,
		AId:  me.Sub,
		Body: req.Body,
	}

	v := validator.New()
	err = v.Struct(dc)
	if err != nil {
		return nil, err
	}

	c, err := s.cs.Create(ctx, dc)
	if err != nil {
		return nil, err
	}

	return &cg.Comment{
		Id:        c.Id,
		PId:       c.PId,
		AId:       c.AId,
		Body:      c.Body,
		CreatedAt: c.Created_at.String(),
	}, nil

}
