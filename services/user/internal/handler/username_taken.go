package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s *userServer) UsernameTaken(c context.Context, req *ug.UsernameTakenRequest) (*ug.UsernameTakenResponse, error) {
	usernameTaken := s.us.UsernameTaken(c, req.Username)

	return &ug.UsernameTakenResponse{Taken: usernameTaken}, nil
}
