package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) GetUser(c context.Context, req *ug.GetUserRequest) (*ug.PublicUser, error) {
	if req.UId == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid User id")
	}

	u, err := s.us.GetById(c, req.UId)
	if err != nil {
		return nil, err
	}

	return u.ToPublicUser(), nil
}
