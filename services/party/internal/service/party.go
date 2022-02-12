package service

import (
	"context"
	"strconv"
	"strings"
	"time"

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

	var sb2 strings.Builder
	sb2.WriteString("IS_GLOBAL#")
	sb2.WriteString(strconv.FormatBool(p.IsPublic))

	gHash := geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)

	t := time.Now()

	dp := datastruct.Party{
		CId:       p.CId,
		KSUID:     sb.String(),
		Title:     p.Title,
		ExpiresAt: t.Add(time.Hour * 24),
		GHash:     gHash,
		IsPublic:  sb2.String(),
	}
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
