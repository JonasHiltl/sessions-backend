package service

import (
	"context"
	"errors"
	"time"

	"github.com/jonashiltl/sessions-backend/services/comment/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/nats-io/nats.go"
)

type CommentService interface {
	Create(ctx context.Context, c dto.Comment) (datastruct.Comment, error)
	Delete(ctx context.Context, uId, pId, cId string) error
	GetByParty(ctx context.Context, pId string) ([]datastruct.Comment, error)
	GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error)
}

type commentService struct {
	repo repository.CommentRepository
	nc   *nats.EncodedConn
}

func NewCommentServie(repo repository.CommentRepository, nc *nats.EncodedConn) CommentService {
	return &commentService{repo: repo, nc: nc}
}

func (cs *commentService) Create(ctx context.Context, c dto.Comment) (datastruct.Comment, error) {
	nanoid, err := gonanoid.New()
	if err != nil {
		return datastruct.Comment{}, errors.New("failed generate id in hook")
	}

	dc := datastruct.Comment{
		Id:         nanoid,
		PId:        c.PId,
		AId:        c.AId,
		Body:       c.Body,
		Created_at: time.Now(),
	}
	result, err := cs.repo.Create(ctx, dc)
	return result, err
}

func (cs *commentService) Delete(ctx context.Context, uId, pId, cId string) error {
	return cs.repo.Delete(ctx, uId, pId, cId)
}

func (cs *commentService) GetByParty(ctx context.Context, pId string) ([]datastruct.Comment, error) {
	return cs.repo.GetByParty(ctx, pId)
}

func (cs *commentService) GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error) {
	return cs.repo.GetByPartyUser(ctx, pId, uId)
}
