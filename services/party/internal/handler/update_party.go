package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
)

func (s *partyServer) UpdateParty(c context.Context, req *pg.UpdatePartyRequest) (*pg.PublicParty, error) {
	d := dto.Party{
		Id:    req.PId,
		UId:   req.RequesterId,
		Title: req.Title,
		Lat:   float64(req.Lat),
		Long:  float64(req.Long),
	}

	p, err := s.ps.Update(c, d)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToPublicParty(), nil
}
