package service

import (
	"context"
	"errors"
	"time"

	"github.com/gofrs/uuid"
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
	GetByUser(ctx context.Context, UserId string, page []byte, limit uint32) ([]datastruct.Party, []byte, error)
	GeoSearch(ctx context.Context, lat float64, long float64, precision uint, page []byte) ([]datastruct.Party, []byte, error)
}

type partyService struct {
	repo repository.PartyRepository
	nc   *nats.EncodedConn
}

func NewPartyServie(repo repository.PartyRepository, nc *nats.EncodedConn) PartyService {
	return &partyService{repo: repo, nc: nc}
}

func (ps *partyService) Create(ctx context.Context, p dto.Party) (datastruct.Party, error) {
	uUserId, err := uuid.NewV1()
	if err != nil {
		return datastruct.Party{}, errors.New("failed generate Party id")
	}

	gHash := geohash.EncodeWithPrecision(float64(p.Lat), float64(p.Long), GEOHASH_PRECISION)

	dp := datastruct.Party{
		Id:            uUserId.String(),
		UserId:        p.UserId,
		Title:         p.Title,
		GHash:         gHash,
		Position:      []float32{p.Lat, p.Long},
		IsPublic:      p.IsPublic,
		StreetAddress: p.StreetAddress,
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		StartDate:     p.StartDate,
	}
	newParty, err := ps.repo.Create(ctx, dp, time.Hour*24)
	return newParty, err
}

func (ps *partyService) Update(ctx context.Context, p dto.Party) (datastruct.Party, error) {
	var gHash string
	if p.Lat != 0 && p.Long != 0 {
		gHash = geohash.EncodeWithPrecision(float64(p.Lat), float64(p.Long), GEOHASH_PRECISION)
	}

	dp := datastruct.Party{
		Id:            p.Id,
		UserId:        p.UserId,
		Title:         p.Title,
		GHash:         gHash,
		Position:      []float32{p.Lat, p.Long},
		IsPublic:      p.IsPublic,
		StreetAddress: p.StreetAddress,
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		StartDate:     p.StartDate,
	}

	err := ps.repo.Update(ctx, dp)
	if err != nil {
		return datastruct.Party{}, err
	}

	newP, err := ps.repo.Get(ctx, p.Id)
	// make sure new value is getting returned

	if p.Title != "" && p.Title != newP.Title {
		newP.Title = p.Title
	}

	if gHash != "" && gHash != newP.GHash {
		newP.GHash = gHash
	}

	if !p.StartDate.IsZero() && p.StartDate != newP.StartDate {
		newP.StartDate = p.StartDate
	}

	if p.StreetAddress != "" && p.StreetAddress != newP.StreetAddress {
		newP.StreetAddress = p.StreetAddress
	}

	if p.PostalCode != "" && p.PostalCode != newP.PostalCode {
		newP.PostalCode = p.PostalCode
	}

	if p.State != "" && p.State != newP.State {
		newP.State = p.State
	}

	if p.Country != "" && p.Country != newP.Country {
		newP.Country = p.Country
	}

	return newP, err
}

func (ps *partyService) Delete(ctx context.Context, UserId, pId string) error {
	return ps.repo.Delete(ctx, UserId, pId)
}

func (ps *partyService) Get(ctx context.Context, pId string) (datastruct.Party, error) {
	return ps.repo.Get(ctx, pId)
}

func (ps *partyService) GetByUser(ctx context.Context, UserId string, page []byte, limit uint32) ([]datastruct.Party, []byte, error) {
	return ps.repo.GetByUser(ctx, UserId, page, limit)
}

func (ps *partyService) GeoSearch(ctx context.Context, lat float64, long float64, precision uint, page []byte) ([]datastruct.Party, []byte, error) {
	if precision == 0 {
		precision = GEOHASH_PRECISION
	}

	h := geohash.Neighbors(geohash.EncodeWithPrecision(lat, long, precision))
	return ps.repo.GeoSearch(ctx, h, page)
}
