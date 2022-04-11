package handler

import (
	"context"

	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) DeleteProfile(c context.Context, req *pg.DeleteProfileRequest) (*common.MessageResponse, error) {
	if req.Id != req.RequesterId {
		return nil, status.Error(codes.Unauthenticated, "You can only delete your own profile")
	}

	err := s.ps.Delete(c, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &common.MessageResponse{Message: "User removed"}, nil
}
