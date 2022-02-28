package datastruct

import (
	"time"

	"github.com/mmcloughlin/geohash"
)

type Story struct {
	Id            string    `json:"id"                       db:"id"             validate:"required"`
	PId           string    `json:"partyId"                  db:"party_id"       validate:"required"`
	UId           string    `json:"userId"                   db:"user_id"        validate:"required"`
	GHash         string    `json:"geohash"                  db:"geohash"        validate:"required"`
	Url           string    `json:"url"                      db:"url"            validate:"required"`
	TaggedFriends []string  `json:"tagged_friends,omitempty" db:"tagged_friends"`
	Created_at    time.Time `json:"createdAt"                db:"created_at"     validate:"required"`
}

type PublicStory struct {
	Id            string    `json:"id"`
	PId           string    `json:"partyId"`
	UId           string    `json:"userId"`
	Lat           float32   `json:"lat,omitempty"`
	Long          float32   `json:"long,omitempty"`
	Url           string    `json:"url"`
	TaggedFriends []string  `json:"tagged_friends,omitempty"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (s Story) ToPublicStory() PublicStory {
	lat, lon := geohash.DecodeCenter(s.GHash)

	return PublicStory{
		Id:            s.Id,
		PId:           s.PId,
		UId:           s.UId,
		Lat:           float32(lat),
		Long:          float32(lon),
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
		CreatedAt:     s.Created_at,
	}
}
