package service

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/party/internal/repository"
	"github.com/mmcloughlin/geohash"
	"github.com/nats-io/nats.go"
)

const GEOHASH_PRECISION uint = 9

type PartyService interface {
	Create(ctx context.Context, p dto.Party) (datastruct.Party, error)
	Update(ctx context.Context, p dto.Party) (datastruct.Party, error)
	Delete(ctx context.Context, UserId, pId string) error
	Get(ctx context.Context, pId string) (datastruct.Party, error)
	GetByUser(ctx context.Context, UserId string, page []byte) ([]datastruct.Party, []byte, error)
	GeoSearch(ctx context.Context, lat float64, long float64, precision uint, page []byte) ([]datastruct.Party, []byte, error)
}

type partyService struct {
	dao repository.Dao
	nc  *nats.EncodedConn
}

func NewPartyServie(dao repository.Dao, nc *nats.EncodedConn) PartyService {
	return &partyService{dao: dao, nc: nc}
}

func (ps *partyService) Create(ctx context.Context, p dto.Party) (datastruct.Party, error) {
	uUserId, err := uuid.NewV1()
	if err != nil {
		return datastruct.Party{}, errors.New("failed generate Party id")
	}

	gHash := geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)

	dp := datastruct.Party{
		Id:       uUserId.String(),
		UserId:   p.UserId,
		Title:    p.Title,
		GHash:    gHash,
		IsPublic: p.IsPublic,
	}
	newParty, err := ps.dao.NewPartyQuery().Create(ctx, dp, time.Hour*24)
	if err == nil {
		ps.nc.Publish("notification.push.party.created", &comtypes.PartyCreatedNotification{CreatorId: p.Id, Title: p.Title})
	}
	return newParty, err
}

func (ps *partyService) Update(ctx context.Context, p dto.Party) (datastruct.Party, error) {
	var gHash string
	if p.Lat != 0 && p.Long != 0 {
		gHash = geohash.EncodeWithPrecision(p.Lat, p.Long, GEOHASH_PRECISION)
	}

	dp := datastruct.Party{
		Id:       p.Id,
		UserId:   p.UserId,
		Title:    p.Title,
		GHash:    gHash,
		IsPublic: p.IsPublic,
	}

	err := ps.dao.NewPartyQuery().Update(ctx, dp)
	if err != nil {
		return datastruct.Party{}, err
	}

	newP, err := ps.dao.NewPartyQuery().Get(ctx, p.Id)
	// make sure new value is getting returned

	if p.Title != "" && p.Title != newP.Title {
		newP.Title = p.Title
	}

	if gHash != "" && gHash != newP.GHash {
		newP.GHash = gHash
	}

	return newP, err
}

func (ps *partyService) Delete(ctx context.Context, UserId, pId string) error {
	return ps.dao.NewPartyQuery().Delete(ctx, UserId, pId)
}

func (ps *partyService) Get(ctx context.Context, pId string) (datastruct.Party, error) {
	return ps.dao.NewPartyQuery().Get(ctx, pId)
}

func (ps *partyService) GetByUser(ctx context.Context, UserId string, page []byte) ([]datastruct.Party, []byte, error) {
	return ps.dao.NewPartyQuery().GetByUser(ctx, UserId, page)
}

func (ps *partyService) GeoSearch(ctx context.Context, lat float64, long float64, precision uint, page []byte) ([]datastruct.Party, []byte, error) {
	if precision == 0 {
		precision = GEOHASH_PRECISION
	}

	h := geohash.Neighbors(geohash.EncodeWithPrecision(lat, long, precision))
	return ps.dao.NewPartyQuery().GeoSearch(ctx, h, page)
}
