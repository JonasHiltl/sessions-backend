package service

import (
	"context"
	"errors"
	"time"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mmcloughlin/geohash"
)

const GEOHASH_PRECISION uint = 9

type PartyService interface {
	Create(ctx context.Context, p dto.Party) (datastruct.Party, error)
	Update(ctx context.Context, p dto.Party) (datastruct.Party, error)
	Delete(ctx context.Context, uId, pId string) error
	Get(ctx context.Context, pId string) (datastruct.Party, error)
	GeoSearch(ctx context.Context, lat float64, long float64, precision uint) ([]datastruct.Party, error)
}

type partyService struct {
	dao repository.Dao
}

func NewPartyServie(dao repository.Dao) PartyService {
	return &partyService{dao: dao}
}

func (ps *partyService) Create(ctx context.Context, p dto.Party) (datastruct.Party, error) {
	nanoid, err := gonanoid.New()
	if err != nil {
		return datastruct.Party{}, errors.New("failed generate id in hook")
	}

	gHash := geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)

	dp := datastruct.Party{
		Id:         nanoid,
		UId:        p.UId,
		Title:      p.Title,
		GHash:      gHash,
		IsPublic:   p.IsPublic,
		Created_at: time.Now(),
	}
	return ps.dao.NewPartyQuery().Create(ctx, dp, time.Hour*24)
}

func (ps *partyService) Update(ctx context.Context, p dto.Party) (datastruct.Party, error) {
	var gHash string
	if p.Lat != 0 && p.Long != 0 {
		gHash = geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)
	}

	dp := datastruct.Party{
		Id:       p.Id,
		UId:      p.UId,
		Title:    p.Title,
		GHash:    gHash,
		IsPublic: p.IsPublic,
	}

	return ps.dao.NewPartyQuery().Update(ctx, dp)
}

func (ps *partyService) Delete(ctx context.Context, uId, pId string) error {
	return ps.dao.NewPartyQuery().Delete(ctx, uId, pId)
}

func (ps *partyService) Get(ctx context.Context, pId string) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Get(ctx, pId)
}

func (ps *partyService) GeoSearch(ctx context.Context, lat float64, long float64, precision uint) ([]datastruct.Party, error) {
	if precision == 0 {
		precision = GEOHASH_PRECISION
	}

	h := geohash.Neighbors(geohash.EncodeWithPrecision(lat, long, precision))
	return ps.dao.NewPartyQuery().GeoSearch(ctx, h)
}
