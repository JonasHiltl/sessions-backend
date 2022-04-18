package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
)

type UserService interface {
	Create(ctx context.Context, u datastruct.User) (datastruct.User, error)
	GetById(ctx context.Context, id string) (datastruct.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.User) (datastruct.User, error)
	UpdateVerified(ctx context.Context, email string, emailVerified bool) (datastruct.User, error)
	GetByEmail(ctx context.Context, email string) (datastruct.User, error)
	GetByEmailOrUsername(ctx context.Context, emailOrUsername string) (datastruct.User, error)
	UsernameTaken(ctx context.Context, username string) bool
}
