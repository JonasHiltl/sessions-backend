package datastruct

import (
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

type AggregatedParty struct {
	Id        string            `json:"id,omitempty"`
	Creator   *pg.Profile       `json:"user_id,omitempty"`
	Title     string            `json:"title,omitempty"`
	IsPublic  bool              `json:"is_public,omitempty"`
	Lat       float32           `json:"lat,omitempty"`
	Long      float32           `json:"long,omitempty"`
	Stories   []*sg.PublicStory `json:"stories,omitempty"`
	CreatedAt string            `json:"created_at,omitempty"`
}

type PagedAggregatedParty struct {
	Parties  []AggregatedParty `json:"stories,omitempty"`
	NextPage string            `json:"nextPage"`
}
