package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (s *profileServer) GetManyProfiles(c context.Context, req *pg.GetManyProfilesRequest) (*pg.GetManyProfilesResponse, error) {
	p, err := s.us.GetMany(c, req.Ids)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	var profiles []*pg.Profile
	for _, p := range p {
		profiles = append(profiles, p.ToGRPCProfile())
	}

	return &pg.GetManyProfilesResponse{Profiles: profiles}, nil
}
