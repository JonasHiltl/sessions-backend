package datastruct

import (
	"github.com/gofrs/uuid"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/mmcloughlin/geohash"
)

type Party struct {
	Id        string `json:"id"         db:"id"         validate:"required"`
	UserId    string `json:"user_id"    db:"user_id"    validate:"required"`
	Title     string `json:"title"      db:"title"      validate:"required"`
	IsPublic  bool   `json:"is_public"  db:"is_public"`
	GHash     string `json:"geohash"    db:"geohash"    validate:"required"`
	StartDate string `json:"start_date" db:"start_date"  validate:"required"`
}

func (p Party) ToPublicParty() *pg.PublicParty {
	lat, lon := geohash.DecodeCenter(p.GHash)
	uuidv1, err := uuid.FromString(p.Id)
	if err != nil {
		return &pg.PublicParty{}
	}

	timestamp, err := uuid.TimestampFromV1(uuidv1)
	if err != nil {
		return &pg.PublicParty{}
	}
	t, err := timestamp.Time()
	if err != nil {
		return &pg.PublicParty{}
	}

	return &pg.PublicParty{
		Id:        p.Id,
		UserId:    p.UserId,
		IsPublic:  p.IsPublic,
		Lat:       float32(lat),
		Long:      float32(lon),
		Title:     p.Title,
		CreatedAt: t.String(),
	}
}
