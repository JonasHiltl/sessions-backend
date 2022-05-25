package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s userServer) GetManyProfiles(ctx context.Context, req *ug.GetManyProfilesRequest) (*ug.GetManyProfilesResponse, error) {
	p, err := s.ps.GetMany(ctx, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	profiles := make([]*ug.Profile, len(p))
	for _, p := range p {
		profiles = append(profiles, p.ToGRPCProfile())
	}

	return &ug.GetManyProfilesResponse{Profiles: profiles}, nil

}
