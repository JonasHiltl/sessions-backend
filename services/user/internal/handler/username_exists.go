package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s *userServer) UsernameExists(c context.Context, req *ug.UsernameExistsRequest) (*ug.UsernameExistsResponse, error) {
	usernameExists, err := s.us.UsernameExists(c, req.Username)
	if err != nil {
		return nil, err
	}

	return &ug.UsernameExistsResponse{Exists: usernameExists}, nil
}
