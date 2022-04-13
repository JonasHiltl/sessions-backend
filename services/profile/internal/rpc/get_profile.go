package rpc

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) GetProfile(c context.Context, req *pg.GetProfileRequest) (*pg.Profile, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid Profile id")
	}

	p, err := s.ps.GetById(c, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
