package datastruct

import (
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

type StoryAggregated struct {
	Id            string        `json:"id,omitempty"`
	PartyId       string        `json:"party_id,omitempty"`
	Creator       *pg.Profile   `json:"user,omitempty"`
	Lat           float32       `json:"lat,omitempty"`
	Long          float32       `json:"long,omitempty"`
	Url           string        `json:"url,omitempty"`
	TaggedFriends []*pg.Profile `json:"tagged_friends,omitempty"`
	CreatedAt     string        `json:"created_at,omitempty"`
}

type PagedStoryAggregated struct {
	Stories  []StoryAggregated `json:"stories,omitempty"`
	NextPage string            `json:"nextPage"`
}
