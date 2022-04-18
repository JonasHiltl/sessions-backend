package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
)

type ProfileService interface {
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
	GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error)
	IncrementFriendCount(ctx context.Context, id string) error
	DecrementFriendCount(ctx context.Context, id string) error
}
