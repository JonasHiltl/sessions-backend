package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s userServer) GetProfile(ctx context.Context, req *ug.GetProfileRequest) (*ug.Profile, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	p, err := s.ps.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
