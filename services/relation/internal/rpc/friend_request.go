package rpc

import (
	"context"
	"time"

	"github.com/jonashiltl/sessions-backend/packages/events"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/datastruct"
)

func (s relationServer) FriendRequest(ctx context.Context, req *rg.FriendRequestRequest) (*cg.SuccessIndicator, error) {
	fr := datastruct.FriendRelation{
		UserId:      req.UserId,
		FriendId:    req.FriendId,
		Accepted:    false,
		RequestedAt: time.Now(),
	}

	err := s.frs.CreateFriendRequest(ctx, fr)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	s.stream.PublishEvent(&events.FriendRequested{
		UserId:   req.UserId,
		FriendId: req.FriendId,
	})

	return &cg.SuccessIndicator{Sucess: true}, nil
}
