package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/user/datastruct"
)

type UserService interface {
	Create(context.Context, datastruct.User) (datastruct.User, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u datastruct.User) error
	UpdateVerified(ctx context.Context, email string, emailVerified bool) error
	RotateEmailCode(ctx context.Context, email string) error
	EmailTaken(ctx context.Context, email string) bool
	UsernameTaken(ctx context.Context, username string) bool
	GetById(ctx context.Context, id string) (datastruct.User, error)
	GetByEmail(ctx context.Context, email string) (datastruct.User, error)
	GetByUsername(ctx context.Context, username string) (datastruct.User, error)
}
