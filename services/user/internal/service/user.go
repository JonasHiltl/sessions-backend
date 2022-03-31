package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/user/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/user/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService interface {
	Create(ctx context.Context, u dto.User) (datastruct.User, error)
	GetById(ctx context.Context, id string) (datastruct.User, error)
	GetByEmail(ctx context.Context, email string) (datastruct.User, error)
	GetByUsername(ctx context.Context, username string) (datastruct.User, error)
	Update(ctx context.Context, u dto.User) (datastruct.User, error)
	Delete(ctx context.Context, id string) error
	UsernameTaken(ctx context.Context, uName string) bool
}

type userService struct {
	dao repository.Dao
}

func NewUserService(dao repository.Dao) UserService {
	return &userService{dao: dao}
}

func (us *userService) Create(ctx context.Context, u dto.User) (datastruct.User, error) {
	id := primitive.NewObjectID()

	newU := datastruct.User{
		Id:        id,
		Provider:  u.Provider,
		Username:  u.Username,
		Email:     u.Email,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Password:  u.Password,
		Avatar:    u.Avatar,
		Role:      datastruct.UserRole.String(),
	}

	res, err := us.dao.NewUserRepository().Create(ctx, newU)
	if err != nil {
		if strings.Contains(err.Error(), "dup key: { email:") {
			return datastruct.User{}, errors.New("email already taken")
		}
		if strings.Contains(err.Error(), "dup key: { username:") {
			return datastruct.User{}, errors.New("username 2 already taken")
		}
	}

	return res, err
}

func (us *userService) GetById(ctx context.Context, id string) (datastruct.User, error) {
	return us.dao.NewUserRepository().GetById(ctx, id)
}
func (us *userService) GetByEmail(ctx context.Context, email string) (datastruct.User, error) {
	return us.dao.NewUserRepository().GetByEmail(ctx, email)
}
func (us *userService) GetByUsername(ctx context.Context, username string) (datastruct.User, error) {
	return us.dao.NewUserRepository().GetByUsername(ctx, username)
}

func (us *userService) Update(ctx context.Context, u dto.User) (datastruct.User, error) {
	id, err := primitive.ObjectIDFromHex(u.Id)
	if err != nil {
		return datastruct.User{}, err
	}

	newU := datastruct.User{
		Id:        id,
		Provider:  u.Provider,
		Username:  u.Username,
		Email:     u.Email,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
		Password:  u.Password,
		Avatar:    u.Avatar,
	}

	return us.dao.NewUserRepository().Update(ctx, newU)
}

func (us *userService) Delete(ctx context.Context, id string) error {
	return us.dao.NewUserRepository().Delete(ctx, id)
}

func (us *userService) UsernameTaken(ctx context.Context, uName string) bool {
	return us.dao.NewUserRepository().UsernameTaken(ctx, uName)
}
