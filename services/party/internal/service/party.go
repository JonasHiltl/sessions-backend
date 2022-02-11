package service

import (
	"context"
	"strings"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/mmcloughlin/geohash"
)

var GEOHASH_PRECISION uint = 9

type PartyService interface {
	Create(ctx context.Context, p dto.Party) (datastruct.Party, error)
	Update(ctx context.Context, cId, pId, title string) (datastruct.Party, error)
	Delete(ctx context.Context, cId, pId string) error
	Get(ctx context.Context, cId, pId string) (datastruct.Party, error)
}

type partyService struct {
	dao repository.Dao
}

func NewPartyServie(dao repository.Dao) PartyService {
	return &partyService{dao: dao}
}

func (ps *partyService) Create(ctx context.Context, p dto.Party) (datastruct.Party, error) {
	var sb strings.Builder
	sb.WriteString("PARTY#")
	sb.WriteString(p.KSUID.String())

	gHash := geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)

	dp := datastruct.Party{CId: p.CId, KSUID: sb.String(), Title: p.Title, Ttl: p.Ttl, GHash: gHash}
	return ps.dao.NewPartyQuery().Create(ctx, dp)
}

func (ps *partyService) Update(ctx context.Context, cId, pId, title string) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Update(ctx, cId, pId, title)
}

func (ps *partyService) Delete(ctx context.Context, cId, pId string) error {
	return ps.dao.NewPartyQuery().Delete(ctx, cId, pId)
}

func (ps *partyService) Get(ctx context.Context, cId, pId string) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Get(ctx, cId, pId)
}
