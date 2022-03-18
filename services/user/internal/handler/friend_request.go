package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	common "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) FriendRequest(c context.Context, req *ug.FriendRequestRequest) (*common.MessageResponse, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, err
	}

	if req.FId == me.Sub {
		return nil, status.Error(codes.Unauthenticated, "You can't add yourself")
	}

	err = s.fs.FriendRequest(c, req.FId, me.Sub)
	if err != nil {
		return nil, err
	}

	return &common.MessageResponse{Message: "Friend request send"}, nil
}
