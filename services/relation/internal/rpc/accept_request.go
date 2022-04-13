package rpc

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/events"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s *relationServer) AcceptFriend(ctx context.Context, req *rg.AcceptFriendRequest) (*rg.FriendRelation, error) {
	fr, err := s.frs.AcceptFriendRelation(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	s.stream.PublishEvent(&events.FriendAccepted{
		UserId:   req.UserId,
		FriendId: req.FriendId,
	})

	return fr.ToGRPCFriendRelation(), nil
}
