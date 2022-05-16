package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s userServer) GetUser(ctx context.Context, req *ug.GetUserRequest) (*ug.User, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	u, err := s.us.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return u.ToGRPCUser(), nil
}
