package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s userServer) GetManyProfilesMap(ctx context.Context, req *ug.GetManyProfilesRequest) (*ug.GetManyProfilesMapResponse, error) {
	ps, err := s.ps.GetMany(ctx, req.Ids)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	profiles := make(map[string]*ug.Profile, len(ps))

	for _, p := range ps {
		profiles[p.Id] = p.ToGRPCProfile()
	}

	return &ug.GetManyProfilesMapResponse{Profiles: profiles}, nil
}
