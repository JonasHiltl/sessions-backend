package handler

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) GetProfile(c context.Context, req *pg.GetProfileRequest) (*pg.PublicProfile, error) {
	if req.UId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid User id")
	}

	u, err := s.us.GetById(c, req.UId)
	if err != nil {
		return nil, err
	}

	return u.ToPublicProfile(), nil
}
