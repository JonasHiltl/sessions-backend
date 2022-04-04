package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (s *partyServer) DeleteParty(c context.Context, req *pg.GetPartyRequest) (*common.MessageResponse, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	err = s.ps.Delete(c, me.Sub, req.PId)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &common.MessageResponse{Message: "Party removed"}, nil
}
