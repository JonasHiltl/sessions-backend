package rpc

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s partyServer) GetManyParties(ctx context.Context, req *pg.GetManyPartiesRequest) (*pg.GetManyPartiesResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	parties := make([]*pg.Party, len(ps))

	for i, p := range ps {
		parties[i] = p.ToGRPCParty()
	}

	return &pg.GetManyPartiesResponse{Parties: parties}, nil
}
