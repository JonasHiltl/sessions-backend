package rpc

import (
	"context"
	"time"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/party/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s partyServer) CreateParty(c context.Context, req *pg.CreatePartyRequest) (*pg.Party, error) {
	start, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid start date")
	}

	d := dto.Party{
		Title:         req.Title,
		UserId:        req.RequesterId,
		Lat:           req.Lat,
		Long:          req.Long,
		IsPublic:      req.IsPublic,
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
		State:         req.State,
		Country:       req.Country,
		StartDate:     start,
	}

	p, err := s.ps.Create(c, d)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToParty(), nil
}
