package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/relation/internal/datastruct"
)

type FriendRelationService interface {
	CreateFriendRelation(ctx context.Context, fr datastruct.FriendRelation) error
	AcceptFriendRelation(ctx context.Context, uId, fId string) error
	RemoveFriendRelation(ctx context.Context, uId, fId string) error
	GetFriendRelation(ctx context.Context, uId, fId string) (datastruct.FriendRelation, error)
	GetFriendsOfUser(ctx context.Context, uId string, page []byte, limit uint32) ([]datastruct.FriendRelation, []byte, error)
}
