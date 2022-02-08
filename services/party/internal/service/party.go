package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
)

type PartyService interface {
	Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Update(ctx context.Context, p datastruct.Party) (datastruct.Party, error)
	Delete(ctx context.Context, pId string) error
	GetById(ctx context.Context, pId string) (datastruct.Party, error)
	Search(ctx context.Context, q string, p int) ([]datastruct.Party, error)
}

type partyService struct {
	dao repository.Dao
}

func NewPartyServie(dao repository.Dao) PartyService {
	return &partyService{dao: dao}
}

func (ps *partyService) Create(ctx context.Context, p datastruct.Party) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Create(ctx, p)
}

func (ps *partyService) Update(ctx context.Context, p datastruct.Party) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Update(ctx, p)
}

func (ps *partyService) Delete(ctx context.Context, partyId string) error {
	return ps.dao.NewPartyQuery().Delete(ctx, partyId)
}

func (ps *partyService) GetById(ctx context.Context, partyId string) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().GetById(ctx, partyId)
}

func (ps *partyService) Search(ctx context.Context, q string, p int) ([]datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Search(ctx, q, p)
}
