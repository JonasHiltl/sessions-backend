package repository

import (
	"time"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (p Party) ToGRPCParty() *pg.Party {
	return &pg.Party{
		Id:            p.ID,
		UserId:        p.UserID,
		Title:         p.Title,
		IsPublic:      p.IsPublic,
		Lat:           p.Lat,
		Long:          p.Long,
		StreetAddress: p.StreetAddress,
		PostalCode:    p.PostalCode,
		State:         p.State,
		Country:       p.Country,
		StartDate:     p.StartDate.UTC().Format(time.RFC3339),
		EndDate:       p.EndDate.UTC().Format(time.RFC3339),
		CreatedAt:     id.Time().UTC().Format(time.RFC3339),
	}
}
