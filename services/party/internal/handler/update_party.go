package handler

import (
	"context"
	"time"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *partyServer) UpdateParty(c context.Context, req *pg.UpdatePartyRequest) (*pg.PublicParty, error) {
	start, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid start date")
	}

	d := dto.Party{
		Id:            req.PartyId,
		UserId:        req.RequesterId,
		Title:         req.Title,
		Lat:           req.Lat,
		Long:          req.Long,
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
		State:         req.State,
		Country:       req.Country,
		StartDate:     start,
	}

	p, err := s.ps.Update(c, d)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToPublicParty(), nil
}
