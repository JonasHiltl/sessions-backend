package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) GetProfile(c context.Context, req *pg.GetProfileRequest) (*pg.Profile, error) {
	if req.PId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid Profile id")
	}

	p, err := s.us.GetById(c, req.PId)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
