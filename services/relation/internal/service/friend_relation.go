package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/relation/internal/datastruct"
)

type FriendRelationService interface {
	CreateFriendRelation(ctx context.Context, fr datastruct.FriendRelation) (datastruct.FriendRelation, error)
	AcceptFriendRelation(ctx context.Context, uId, fId string) (datastruct.FriendRelation, error)
}
