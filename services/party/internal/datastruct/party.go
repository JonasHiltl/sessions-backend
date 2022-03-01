package datastruct

import (
	"time"

	"github.com/mmcloughlin/geohash"
)

type Party struct {
	Id         string    `json:"id"        db:"id"         validate:"required"`
	UId        string    `json:"userId"    db:"user_id"    validate:"required"`
	Title      string    `json:"title"     db:"title"      validate:"required"`
	IsPublic   bool      `json:"isPublic"  db:"is_public"`
	GHash      string    `json:"geohash"   db:"geohash"    validate:"required"`
	Created_at time.Time `json:"createdAt" db:"created_at" validate:"required"`
}

type PublicParty struct {
	Id       string   `json:"id"`
	UId      string   `json:"userId"`
	IsPublic bool     `json:"isPublic"`
	Lat      float64  `json:"lat"`
	Long     float64  `json:"long"`
	Stories  []string `json:"stories,omitempty"`
	Title    string   `json:"title"`
}

func (p Party) ToPublicParty() PublicParty {
	lat, lon := geohash.DecodeCenter(p.GHash)

	return PublicParty{
		Id:       p.Id,
		UId:      p.UId,
		IsPublic: p.IsPublic,
		Lat:      lat,
		Long:     lon,
		Title:    p.Title,
	}
}
