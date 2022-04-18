package rpc

import (
	"context"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) DeleteUser(ctx context.Context, req *ug.DeleteUserRequest) (*cg.MessageResponse, error) {
	if req.Id == "" {
		return nil, status.Error(codes.InvalidArgument, "empty user id")
	}
	err := s.us.Delete(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.MessageResponse{Message: "User removed"}, nil
}
