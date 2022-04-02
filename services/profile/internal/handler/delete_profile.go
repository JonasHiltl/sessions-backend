package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) DeleteProfile(c context.Context, req *pg.GetProfileRequest) (*common.MessageResponse, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, err
	}

	if req.UId != me.Sub {
		return nil, status.Error(codes.Unauthenticated, "You can only delete your own account")
	}

	err = s.us.Delete(c, me.Sub)
	if err != nil {
		return nil, err
	}

	return &common.MessageResponse{Message: "User removed"}, nil
}
