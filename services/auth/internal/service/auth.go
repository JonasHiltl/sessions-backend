package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jonashiltl/sessions-backend/services/auth/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/dto"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthService interface {
	Create(context.Context, dto.AuthUser) (datastruct.AuthUser, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, u dto.AuthUser) (datastruct.AuthUser, error)
	GetById(ctx context.Context, id string) (datastruct.AuthUser, error)
	GetByEmail(ctx context.Context, email string) (datastruct.AuthUser, error)
	SetVerified(ctx context.Context, email string) (datastruct.AuthUser, error)
}

type authService struct {
	dao repository.Dao
}

func NewAuthService(dao repository.Dao) AuthService {
	return &authService{dao: dao}
}

func (as *authService) Create(ctx context.Context, u dto.AuthUser) (datastruct.AuthUser, error) {
	id := primitive.NewObjectID()

	newU := datastruct.AuthUser{
		Id:            id,
		Provider:      u.Provider,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		PasswordHash:  u.PasswordHash,
		Role:          u.Role,
	}

	res, err := as.dao.NewAuthRepository().Create(ctx, newU)
	if err != nil {
		if strings.Contains(err.Error(), "dup key: { email:") {
			return datastruct.AuthUser{}, errors.New("email already taken")
		}
	}

	return res, err
}

func (as *authService) Delete(ctx context.Context, id string) error {
	return as.dao.NewAuthRepository().Delete(ctx, id)
}

func (as *authService) Update(ctx context.Context, u dto.AuthUser) (datastruct.AuthUser, error) {
	id, err := primitive.ObjectIDFromHex(u.Id)
	if err != nil {
		return datastruct.AuthUser{}, err
	}

	newU := datastruct.AuthUser{
		Id:            id,
		Provider:      u.Provider,
		Email:         u.Email,
		EmailVerified: u.EmailVerified,
		PasswordHash:  u.PasswordHash,
		Role:          u.Role,
	}

	return as.dao.NewAuthRepository().Update(ctx, newU)
}

func (as *authService) GetById(ctx context.Context, id string) (datastruct.AuthUser, error) {
	return as.dao.NewAuthRepository().GetById(ctx, id)
}
func (as *authService) GetByEmail(ctx context.Context, username string) (datastruct.AuthUser, error) {
	return as.dao.NewAuthRepository().GetByEmail(ctx, username)
}

func (as *authService) SetVerified(ctx context.Context, email string) (datastruct.AuthUser, error) {
	return as.dao.NewAuthRepository().SetVerified(ctx, email)
}
