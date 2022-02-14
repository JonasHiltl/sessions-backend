package service

import (
	"context"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	gonanoid "github.com/matoous/go-nanoid/v2"
	"github.com/mmcloughlin/geohash"
)

var GEOHASH_PRECISION uint = 9

type PartyService interface {
	Create(ctx context.Context, p dto.Party) (datastruct.Party, error)
	Update(ctx context.Context, pId string, p dto.Party) (datastruct.Party, error)
	Delete(ctx context.Context, cId, pId string) error
	Get(ctx context.Context, pId string) (datastruct.Party, error)
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

	var id strings.Builder
	id.WriteString("P#")
	id.WriteString(nanoid)

	gHash := geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)

	t := time.Now()

	dp := datastruct.Party{
		Id:        id.String(),
		SK:        "party",
		CreatorId: p.CreatorId,
		Title:     p.Title,
		Ttl:       t.Add(time.Hour * 24),
		GHash:     gHash,
		IsPublic:  strconv.FormatBool(p.IsPublic),
	}
	return ps.dao.NewPartyQuery().Create(ctx, dp)
}

func (ps *partyService) Update(ctx context.Context, pId string, p dto.Party) (datastruct.Party, error) {
	var gHash string
	if p.Lat != 0 && p.Long != 0 {
		gHash = geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)
	}

	dp := datastruct.Party{
		Id:        pId,
		SK:        "party",
		CreatorId: p.CreatorId,
		Title:     p.Title,
		GHash:     gHash,
		IsPublic:  strconv.FormatBool(p.IsPublic),
	}

	return ps.dao.NewPartyQuery().Update(ctx, dp)
}

func (ps *partyService) Delete(ctx context.Context, cId, pId string) error {
	return ps.dao.NewPartyQuery().Delete(ctx, cId, pId)
}

func (ps *partyService) Get(ctx context.Context, pId string) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Get(ctx, pId)
}
