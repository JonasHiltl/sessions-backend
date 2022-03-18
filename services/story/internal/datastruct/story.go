package datastruct

import (
	"time"

	"github.com/gofrs/uuid"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/mmcloughlin/geohash"
)

type Story struct {
	Id            string   `json:"id"                       db:"id"             validate:"required"`
	PId           string   `json:"partyId"                  db:"party_id"       validate:"required"`
	UId           string   `json:"userId"                   db:"user_id"        validate:"required"`
	GHash         string   `json:"geohash"                  db:"geohash"        validate:"required"`
	Url           string   `json:"url"                      db:"url"            validate:"required"`
	TaggedFriends []string `json:"tagged_friends,omitempty" db:"tagged_friends"`
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

type PagedStories struct {
	Stories  []PublicStory `json:"stories,omitempty"`
	NextPage string        `json:"nextPage"`
}

func (s Story) ToPublicStory() *sg.PublicStory {
	lat, lon := geohash.DecodeCenter(s.GHash)

	uuidv1, err := uuid.FromString(s.Id)
	if err != nil {
		return &sg.PublicStory{}
	}
	timestamp, err := uuid.TimestampFromV1(uuidv1)
	if err != nil {
		return &sg.PublicStory{}
	}
	t, err := timestamp.Time()
	if err != nil {
		return &sg.PublicStory{}
	}

	return &sg.PublicStory{
		Id:            s.Id,
		PId:           s.PId,
		UId:           s.UId,
		Lat:           float32(lat),
		Long:          float32(lon),
		Url:           s.Url,
		TaggedFriends: s.TaggedFriends,
		CreatedAt:     t.String(),
	}
}
