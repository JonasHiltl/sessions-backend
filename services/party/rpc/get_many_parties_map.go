package rpc

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s partyServer) GetManyPartiesMap(ctx context.Context, req *pg.GetManyPartiesRequest) (*pg.GetManyPartiesMapResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	parties := make(map[string]*pg.Party)

	for _, p := range ps {
		parties[p.Id] = p.ToGRPCParty()
	}

	return &pg.GetManyPartiesMapResponse{Parties: parties}, nil
}
