package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/party/datastruct"
)

type PartyService interface {
	Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Update(ctx context.Context, p datastruct.Party) error
	Delete(ctx context.Context, uId, pId string) error
	Get(ctx context.Context, pId string) (datastruct.Party, error)
	GetMany(ctx context.Context, ids []string) ([]datastruct.Party, error)
	GetByUser(ctx context.Context, uId string, page []byte, limit uint32) ([]datastruct.Party, []byte, error)
}
