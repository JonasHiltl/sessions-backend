package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s userServer) UsernameTaken(ctx context.Context, req *ug.UsernameTakenRequest) (*ug.UsernameTakenResponse, error) {
	usernameTaken := s.us.UsernameTaken(ctx, req.Username)

	return &ug.UsernameTakenResponse{Taken: usernameTaken}, nil
}
