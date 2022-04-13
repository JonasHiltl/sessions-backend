package rpc

import (
	"context"
	"time"

	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/datastruct"
)

func (s *relationServer) FriendRequest(ctx context.Context, req *rg.FriendRequestRequest) (*rg.FriendRelation, error) {
	fr := datastruct.FriendRelation{
		UserId:    req.UserId,
		FriendId:  req.FriendId,
		Accepted:  false,
		CreatedAt: time.Now(),
	}

	fr, err := s.frs.CreateFriendRelation(ctx, fr)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return fr.ToGRPCProfile(), nil
}
