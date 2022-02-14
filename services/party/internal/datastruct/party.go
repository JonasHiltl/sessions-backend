package datastruct

import (
	"strconv"
	"strings"
	"time"

	"github.com/mmcloughlin/geohash"
)

type Party struct {
	Id        string    `json:"id"        dynamo:"pk,hash"          validate:"required"`
	SK        string    `json:"-"         dynamo:"sk,range"         validate:"required"`
	CreatorId string    `json:"creatorId" dynamo:"gsi1_pk_userId"   validate:"required"`
	Title     string    `json:"title"     dynamo:"title"            validate:"required"`
	IsPublic  string    `json:"isPublic"  dynamo:"gsi2_pk_isPublic" validate:"required"`
	GHash     string    `json:"geohash"   dynamo:"gsi2_sk_geohash"  validate:"required"`
	Ttl       time.Time `json:"ttl"       dynamo:"ttl,unixtime"     validate:"required"`
}

type PublicParty struct {
	Id        string    `json:"id"` // PARTY#CreatorId
	CreatorId string    `json:"creatorId"`
	IsPublic  bool      `json:"isPublic"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	Title     string    `json:"title"`
	Ttl       time.Time `json:"ttl"`
}

func (p Party) ToPublicParty() PublicParty {
	isP, err := strconv.ParseBool(strings.TrimPrefix(p.IsPublic, "IS_PUBLIC#"))
	if err != nil {
		return PublicParty{}
	}

	lat, lng := geohash.DecodeCenter(p.GHash)

	return PublicParty{
		Id:        p.Id,
		CreatorId: strings.TrimPrefix(p.CreatorId, "U#"),
		IsPublic:  isP,
		Lat:       lat,
		Long:      lng,
		Title:     p.Title,
		Ttl:       p.Ttl,
	}
}
