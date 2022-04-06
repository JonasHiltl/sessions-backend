package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) DeleteProfile(c context.Context, req *pg.DeleteProfileRequest) (*common.MessageResponse, error) {
	if req.PId != req.RequesterId {
		return nil, status.Error(codes.Unauthenticated, "You can only delete your own profile")
	}

	err := s.us.Delete(c, req.RequesterId)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &common.MessageResponse{Message: "User removed"}, nil
}
