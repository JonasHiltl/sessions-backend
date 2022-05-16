package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
)

type ProfileService interface {
	GetMany(ctx context.Context, ids []string) ([]datastruct.Profile, error)
	GetById(ctx context.Context, id string) (datastruct.Profile, error)
}
