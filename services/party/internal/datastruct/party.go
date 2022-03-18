package datastruct

import (
	"time"

	"github.com/gofrs/uuid"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/mmcloughlin/geohash"
)

type Party struct {
	Id       string `json:"id"        db:"id"         validate:"required"`
	UId      string `json:"userId"    db:"user_id"    validate:"required"`
	Title    string `json:"title"     db:"title"      validate:"required"`
	IsPublic bool   `json:"isPublic"  db:"is_public"`
	GHash    string `json:"geohash"   db:"geohash"    validate:"required"`
}

type PublicParty struct {
	Id        string    `json:"id"`
	UId       string    `json:"userId"`
	IsPublic  bool      `json:"isPublic"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	Stories   []string  `json:"stories,omitempty"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"createdAt"`
}

type PagedParties struct {
	Parties  []PublicParty `json:"parties,omitempty"`
	NextPage string        `json:"nextPage"`
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
		UId:       p.UId,
		IsPublic:  p.IsPublic,
		Lat:       float32(lat),
		Long:      float32(lon),
		Title:     p.Title,
		CreatedAt: t.String(),
	}
}
