package service

import (
	"context"
	"errors"
	"time"

	"github.com/jonashiltl/sessions-backend/services/comment/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/comment/internal/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type CommentService interface {
	Create(ctx context.Context, c dto.Comment) (datastruct.Comment, error)
	Delete(ctx context.Context, uId, pId, cId string) error
	GetByParty(ctx context.Context, pId string) ([]datastruct.Comment, error)
	GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error)
}

type commentService struct {
	dao repository.Dao
}

func NewCommentServie(dao repository.Dao) CommentService {
	return &commentService{dao: dao}
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
	return cs.dao.NewCommentQuery().Create(ctx, dc)
}

func (cs *commentService) Delete(ctx context.Context, uId, pId, cId string) error {
	return cs.dao.NewCommentQuery().Delete(ctx, uId, pId, cId)
}

func (cs *commentService) GetByParty(ctx context.Context, pId string) ([]datastruct.Comment, error) {
	return cs.dao.NewCommentQuery().GetByParty(ctx, pId)
}

func (cs *commentService) GetByPartyUser(ctx context.Context, pId, uId string) ([]datastruct.Comment, error) {
	return cs.dao.NewCommentQuery().GetByPartyUser(ctx, pId, uId)
}
