package rpc

import (
	"context"
	"time"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/party/datastruct"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s partyServer) UpdateParty(c context.Context, req *pg.UpdatePartyRequest) (*pg.Party, error) {
	start, err := time.Parse(time.RFC3339, req.StartDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid start date")
	}

	end, err := time.Parse(time.RFC3339, req.EndDate)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid end date")
	}

	id, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	p := datastruct.Party{
		Id:            id.String(),
		UserId:        req.RequesterId,
		Title:         req.Title,
		Lat:           req.Lat,
		Long:          req.Long,
		StreetAddress: req.StreetAddress,
		PostalCode:    req.PostalCode,
		State:         req.State,
		Country:       req.Country,
		StartDate:     start,
		EndDate:       end,
	}

	err = s.ps.Update(c, p)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCParty(), nil
}
