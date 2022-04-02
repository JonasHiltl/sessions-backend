package handler

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (s *profileServer) UsernameTaken(c context.Context, req *pg.UsernameTakenRequest) (*pg.UsernameTakenResponse, error) {
	usernameTaken := s.us.UsernameTaken(c, req.Username)

	return &pg.UsernameTakenResponse{Taken: usernameTaken}, nil
}
