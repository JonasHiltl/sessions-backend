package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
)

func (s *partyServer) CreateParty(c context.Context, req *pg.CreatePartyRequest) (*pg.PublicParty, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, err
	}

	d := dto.Party{
		Title:    req.Title,
		UId:      me.Sub,
		Lat:      float64(req.Lat),
		Long:     float64(req.Long),
		IsPublic: req.IsPublic,
	}

	p, err := s.ps.Create(c, d)
	if err != nil {
		return nil, err
	}

	return p.ToPublicParty(), nil
}
