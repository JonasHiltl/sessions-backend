package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (s *partyServer) DeleteParty(c context.Context, req *pg.DeletePartyRequest) (*common.MessageResponse, error) {
	err := s.ps.Delete(c, req.RequesterId, req.PartyId)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &common.MessageResponse{Message: "Party removed"}, nil
}
