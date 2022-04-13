package rpc

import (
	"context"
	"encoding/base64"

	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *relationServer) GetFriendsOfUser(ctx context.Context, req *rg.GetFriendsOfUserRequest) (*rg.PagedFriendRelations, error) {
	p, err := base64.URLEncoding.DecodeString(req.NextPage)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Next Page Param")
	}

	frs, p, err := s.frs.GetFriendsOfUser(ctx, req.UserId, p, req.Limit)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	nextPage := base64.URLEncoding.EncodeToString(p)

	var res []*rg.FriendRelation
	for _, fr := range frs {
		res = append(res, fr.ToGRPCFriendRelation())
	}

	return &rg.PagedFriendRelations{Relations: res, NextPage: nextPage}, nil
}
