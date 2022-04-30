package datastruct

import (
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type AggregatedParty struct {
	Id            string            `json:"id,omitempty"`
	Creator       *ug.Profile       `json:"creator,omitempty"`
	Title         string            `json:"title,omitempty"`
	IsPublic      bool              `json:"is_public,omitempty"`
	Lat           float32           `json:"lat,omitempty"`
	Long          float32           `json:"long,omitempty"`
	StreetAddress string            `json:"street_address,omitempty"`
	PostalCode    string            `json:"postal_code,omitempty"`
	State         string            `json:"state,omitempty"`
	Country       string            `json:"country,omitempty"`
	Stories       []*sg.PublicStory `json:"stories,omitempty"`
	StartDate     string            `json:"start_date,omitempty"`
	CreatedAt     string            `json:"created_at,omitempty"`
}

type PagedAggregatedParty struct {
	Parties  []AggregatedParty `json:"parties"`
	NextPage string            `json:"nextPage"`
}
