package datastruct

import (
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type AggregatedStory struct {
	Id            string        `json:"id,omitempty"`
	PartyId       string        `json:"party_id,omitempty"`
	Creator       *ug.Profile   `json:"creator,omitempty"`
	Lat           float32       `json:"lat,omitempty"`
	Long          float32       `json:"long,omitempty"`
	Url           string        `json:"url,omitempty"`
	TaggedFriends []*ug.Profile `json:"tagged_friends,omitempty"`
	CreatedAt     string        `json:"created_at,omitempty"`
}

type PagedAggregatedStory struct {
	Stories  []AggregatedStory `json:"stories,omitempty"`
	NextPage string            `json:"nextPage"`
}
