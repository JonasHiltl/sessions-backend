package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) DeleteUser(c context.Context, req *ug.GetUserRequest) (*common.MessageResponse, error) {
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
