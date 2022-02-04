package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/jonashiltl/sessions-backend/services/user/ent"
	"github.com/jonashiltl/sessions-backend/services/user/ent/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
)

type UserService interface {
	Create(ctx context.Context, u datastruct.RequestUser) (*ent.User, error)
	GetById(ctx context.Context, id uuid.UUID) (*ent.User, error)
	Update(ctx context.Context, id uuid.UUID, u datastruct.RequestUser) (*ent.User, error)
	Delete(ctx context.Context, id uuid.UUID) error
	UsernameExists(ctx context.Context, username string) (bool, error)
}

type userService struct {
	client *ent.UserClient
}

func NewUserService(client *ent.UserClient) UserService {
	return &userService{client: client}
}

func (us *userService) Create(ctx context.Context, u datastruct.RequestUser) (*ent.User, error) {
	res, err := us.client.
		Create().
		SetUsername(u.Username).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetEmail(u.Email).
		SetPassword(u.Password).
		Save(ctx)

	return res, err
}

func (us *userService) GetById(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	res, err := us.client.Get(ctx, id)
	return res, err
}

func (us *userService) Update(ctx context.Context, id uuid.UUID, u datastruct.RequestUser) (*ent.User, error) {
	res, err := us.client.
		UpdateOneID(id).
		SetUsername(u.Username).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetEmail(u.Email).
		SetPicture(u.Picture).
		Save(ctx)
	return res, err
}

func (us *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return us.client.DeleteOneID(id).Exec(ctx)
}

func (us *userService) UsernameExists(ctx context.Context, username string) (bool, error) {
	return us.client.Query().Where(user.UsernameEQ(username)).Exist(ctx)
}
