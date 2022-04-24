package rpc

import (
	"context"

	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s partyServer) DeleteParty(c context.Context, req *pg.DeletePartyRequest) (*common.MessageResponse, error) {
	err := s.ps.Delete(c, req.RequesterId, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &common.MessageResponse{Message: "Party removed"}, nil
}
