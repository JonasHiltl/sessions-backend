package rpc

import (
	"context"
	"encoding/base64"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *partyServer) GetFovorisingUsersByParty(ctx context.Context, req *pg.GetFovorisingUsersByPartyRequest) (*pg.PagedFavoriteParties, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	fps, p, err := s.ps.GetFavorisingUsersByParty(ctx, req.PartyId, p, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var res []*pg.FavoriteParty
	for _, fp := range fps {
		res = append(res, fp.ToGRPCFavoriteParty())
	}

	return &pg.PagedFavoriteParties{FavoriteParties: res, NextPage: nextPage}, nil

}
