package handler

import (
	"context"
	"encoding/base64"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GeoSearchBody struct {
	Lat       float64 `json:"lat"       validate:"required,latitude"`
	Long      float64 `json:"long"      validate:"required,longitude"`
	Precision uint    `json:"precision"`
}

func (s *partyServer) GeoSearch(c context.Context, req *pg.GeoSearchRequest) (*pg.PagedParties, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	ps, p, err := s.ps.GeoSearch(c, float64(req.Lat), float64(req.Long), uint(req.Precision), p)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var pp []*pg.PublicParty

	for _, ps := range ps {
		pp = append(pp, ps.ToPublicParty())
	}

	return &pg.PagedParties{Parties: pp, NextPage: nextPage}, nil
}
