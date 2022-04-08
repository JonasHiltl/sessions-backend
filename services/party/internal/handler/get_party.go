package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *partyServer) GetParty(c context.Context, req *pg.GetPartyRequest) (*pg.PublicParty, error) {
	if req.PartyId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid Party id")
	}

	p, err := s.ps.Get(c, req.PartyId)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToPublicParty(), nil
}
