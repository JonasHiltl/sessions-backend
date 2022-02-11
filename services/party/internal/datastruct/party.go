package datastruct

import (
	"strconv"
	"strings"
	"time"

	"github.com/mmcloughlin/geohash"
	"github.com/segmentio/ksuid"
)

type Party struct {
	CId       string    `json:"creatorId"    dynamo:"pk,hash"             validate:"required,len=21"` // PARTY#CreatorId
	KSUID     string    `json:"id,omitempty" dynamo:"sk,range"            validate:"required"`
	IsGlobal  string    `json:"isGlobal"     dynamo:"gsi_pk_isGlobal"                         index:",hash"`
	GHash     string    `json:"geohash"      dynamo:"gsi_sk_geohash"      validate:"required" index:",range"`
	Title     string    `json:"title"        dynamo:"title,string"        validate:"required"`
	ExpiresAt time.Time `json:"-"            dynamo:"expiresAt,unixtime"  validate:"required"`
	Stories   map[string]struct {
		creatorId string
		ksuid     ksuid.KSUID
	} `json:"stories,omitempty"              dynamo:"stories,set"`
}

type PublicParty struct {
	CId       string    `json:"creatorId"` // PARTY#CreatorId
	KSUID     string    `json:"id,omitempty"`
	IsGlobal  bool      `json:"isGlobal"`
	Lat       float64   `json:"lat"`
	Long      float64   `json:"long"`
	Title     string    `json:"title"`
	ExpiresAt time.Time `json:"expiresAt"`
	Stories   map[string]struct {
		creatorId string
		ksuid     ksuid.KSUID
	} `json:"stories,omitempty"`
}

func (p Party) ToPublicParty() PublicParty {
	isG, err := strconv.ParseBool(strings.TrimPrefix(p.IsGlobal, "IS_GLOBAL#"))
	if err != nil {
		return PublicParty{}
	}

	lat, lng := geohash.DecodeCenter(p.GHash)

	return PublicParty{
		CId:       p.CId,
		KSUID:     strings.TrimPrefix(p.KSUID, "PARTY#"),
		IsGlobal:  isG,
		Lat:       lat,
		Long:      lng,
		Title:     p.Title,
		ExpiresAt: p.ExpiresAt,
		Stories:   p.Stories,
	}
}
