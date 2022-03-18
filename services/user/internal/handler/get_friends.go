package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s *userServer) GetFriends(c context.Context, req *ug.GetFriendsRequest) (*ug.FriendResponse, error) {
	friends, err := s.fs.Get(c, req.UId, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}

	var publicFriends []*ug.PublicUser

	for _, user := range friends {
		publicFriends = append(publicFriends, user.ToPublicUser())
	}

	return &ug.FriendResponse{Friends: publicFriends}, nil
}
