package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
)

func (s *partyServer) CreateParty(c context.Context, req *pg.CreatePartyRequest) (*pg.PublicParty, error) {
	d := dto.Party{
		Title:    req.Title,
		UId:      req.RequesterId,
		Lat:      float64(req.Lat),
		Long:     float64(req.Long),
		IsPublic: req.IsPublic,
	}

	p, err := s.ps.Create(c, d)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToPublicParty(), nil
}
