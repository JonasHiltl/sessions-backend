package handler

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s *profileServer) GetManyProfiles(c context.Context, req *pg.GetManyProfilesRequest) (*pg.GetManyProfilesResponse, error) {
	p, err := s.ps.GetMany(c, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	var profiles []*pg.Profile
	for _, p := range p {
		profiles = append(profiles, p.ToGRPCProfile())
	}

	return &pg.GetManyProfilesResponse{Profiles: profiles}, nil
}
