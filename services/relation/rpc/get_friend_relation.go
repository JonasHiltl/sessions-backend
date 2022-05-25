package rpc

import (
	"context"

	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s relationServer) GetFriendRelation(ctx context.Context, req *rg.GetFriendRelationRequest) (*rg.FriendRelation, error) {
	fr, err := s.fs.GetFriendRelation(ctx, req.UserId, req.FriendId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return fr.ToGRPCFriendRelation(), nil
}
