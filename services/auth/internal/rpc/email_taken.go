package rpc

import (
	"context"

	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
)

func (s *authServer) EmailTaken(ctx context.Context, req *ag.EmailTakenRequest) (*ag.EmailTakenResponse, error) {
	taken := s.authService.EmailTaken(ctx, req.Email)

	return &ag.EmailTakenResponse{Taken: taken}, nil
}
