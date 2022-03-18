package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s *userServer) SearchFriends(c context.Context, req *ug.FriendSearchRequest) (*ug.FriendResponse, error) {
	friends, err := s.fs.Search(c, req.UId, req.Query, req.Accepted)
	if err != nil {
		return nil, err
	}

	var publicFriends []*ug.PublicUser

	for _, user := range friends {
		publicFriends = append(publicFriends, user.ToPublicUser())
	}

	return &ug.FriendResponse{Friends: publicFriends}, nil
}
