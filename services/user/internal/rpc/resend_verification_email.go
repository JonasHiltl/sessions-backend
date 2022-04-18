package rpc

import (
	"context"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) ResendVerificationEmail(ctx context.Context, req *cg.Empty) (*cg.MessageResponse, error) {
	return nil, status.Error(codes.Unavailable, "not yet implemented")
}
