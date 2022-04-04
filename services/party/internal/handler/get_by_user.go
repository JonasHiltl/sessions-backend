package handler

import (
	"context"
	"encoding/base64"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *partyServer) GetByUser(c context.Context, req *pg.GetByUserRequest) (*pg.PagedParties, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	ps, p, err := s.ps.GetByUser(c, req.UId, p)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var pp []*pg.PublicParty
	for _, p := range ps {
		pp = append(pp, p.ToPublicParty())
	}

	return &pg.PagedParties{Parties: pp, NextPage: nextPage}, nil
}
