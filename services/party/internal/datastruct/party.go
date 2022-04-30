package datastruct

import (
	"time"

	"github.com/gofrs/uuid"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

type Party struct {
	Id            string    `json:"id"             db:"id"             validate:"required"`
	UserId        string    `json:"user_id"        db:"user_id"        validate:"required"`
	Title         string    `json:"title"          db:"title"          validate:"required"`
	IsPublic      bool      `json:"is_public"      db:"is_public"`
	GHash         string    `json:"geohash"        db:"geohash"        validate:"required"`
	Position      []float32 `json:"position"       db:"position"       validate:"required"` // [lat, long]
	StreetAddress string    `json:"street_address" db:"street_address" validate:"required"`
	PostalCode    string    `json:"postal_code"    db:"postal_code"    validate:"required"`
	State         string    `json:"state"          db:"state"          validate:"required"`
	Country       string    `json:"country"        db:"country"        validate:"required"`
	StartDate     time.Time `json:"start_date"     db:"start_date"     validate:"required"`
}

func (p Party) ToParty() *pg.Party {
	// get party creation date from uuid
	uuidv1, err := uuid.FromString(p.Id)
	if err != nil {
		return &pg.Party{}
	}
	timestamp, err := uuid.TimestampFromV1(uuidv1)
	if err != nil {
		return &pg.Party{}
	}
	t, err := timestamp.Time()
	if err != nil {
		return &pg.Party{}
	}

	var lat float32
	var long float32
	if len(p.Position) == 2 {
		lat = p.Position[0]
		long = p.Position[1]
	}

	return &pg.Party{
		Id:            p.Id,
		UserId:        p.UserId,
		Title:         p.Title,
		IsPublic:      p.IsPublic,
		Lat:           lat,
		Long:          long,
		StreetAddress: p.StreetAddress,
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		StartDate:     p.StartDate.UTC().Format(time.RFC3339),
		CreatedAt:     t.UTC().Format(time.RFC3339),
	}
}
