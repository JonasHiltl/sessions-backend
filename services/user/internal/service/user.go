package service

import (
	"context"
	"errors"
	"strings"

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
	CountFriends(ctx context.Context, id uuid.UUID) int
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

	if err != nil {
		if strings.Contains(err.Error(), "for key 'users.email'") {
			return &ent.User{}, errors.New("email already taken")
		}
		if strings.Contains(err.Error(), "for key 'users.username'") {
			return &ent.User{}, errors.New("username already taken")
		}
	}

	return res, err
}

func (us *userService) GetById(ctx context.Context, id uuid.UUID) (*ent.User, error) {
	res, err := us.client.
		Query().
		Where(user.ID(id)).
		Select(user.FieldID, user.FieldUsername, user.FieldFirstName, user.FieldLastName, user.FieldRole, user.FieldCreatedAt, user.FieldPicture).
		Only(ctx)

	return res, err
}

func (us *userService) Update(ctx context.Context, id uuid.UUID, u datastruct.RequestUser) (*ent.User, error) {
	builder := us.client.UpdateOneID(id)

	if u.Username != "" {
		builder.SetUsername(u.Username)
	}
	if u.FirstName != "" {
		builder.SetFirstName(u.FirstName)
	}
	if u.LastName != "" {
		builder.SetLastName(u.LastName)
	}
	if u.Email != "" {
		builder.SetEmail(u.Email)
	}
	if u.Picture != "" {
		builder.SetPicture(u.Picture)
	}
	if u.Password != "" {
		builder.SetPassword(u.Password)
	}

	res, err := builder.
		Select(user.FieldID, user.FieldUsername, user.FieldFirstName, user.FieldLastName, user.FieldRole, user.FieldCreatedAt, user.FieldPicture).
		Save(ctx)

	return res, err
}

func (us *userService) Delete(ctx context.Context, id uuid.UUID) error {
	return us.client.DeleteOneID(id).Exec(ctx)
}

func (us *userService) UsernameExists(ctx context.Context, username string) (bool, error) {
	return us.client.
		Query().
		Where(user.UsernameEQ(username)).
		Exist(ctx)
}

func (us *userService) CountFriends(ctx context.Context, id uuid.UUID) int {
	count, err := us.client.Query().Where(user.ID(id)).QueryFriends().Count(ctx)
	if err != nil {
		return 0
	}
	return count
}
