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

func (s *partyServer) CreateParty(c context.Context, req *pg.CreatePartyRequest) (*pg.PublicParty, error) {
	start, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid start date")
	}

	d := dto.Party{
		Title:     req.Title,
		UserId:    req.RequesterId,
		Lat:       float64(req.Lat),
		Long:      float64(req.Long),
		IsPublic:  req.IsPublic,
		StartDate: start,
	}

	p, err := s.ps.Create(c, d)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToPublicParty(), nil
}
