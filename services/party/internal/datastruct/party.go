package datastruct

import (
	"time"

	"github.com/segmentio/ksuid"
)

type Party struct {
	CId      string    `json:"creatorId"    dynamo:"pk,hash"      validate:"required,len=21"` // PARTY#CreatorId
	KSUID    string    `json:"id,omitempty" dynamo:"sk,range"     validate:"required"`
	IsGlobal bool      `json:"isGlobal"     index:"isGlobal,hash"`
	GHash    string    `json:"geohash"      index:"geohash,range" validate:"required"`
	Title    string    `json:"title"        dynamo:"title,string"       validate:"required"`
	Ttl      time.Time `json:"-"            dynamo:"ttl,unixtime" validate:"required"`
	Stories  map[string]struct {
		creatorId string
		ksuid     ksuid.KSUID
	} `json:"stories,omitempty"              dynamo:"stories,set"`
}
