package service

import (
	"context"
	"errors"
	"strings"

	"github.com/jonashiltl/sessions-backend/services/user/ent"
	"github.com/jonashiltl/sessions-backend/services/user/ent/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, u datastruct.RequestUser) (string, error)
	Login(ctx context.Context, l datastruct.LoginBody) (string, error)
}

type authService struct {
	client       *ent.UserClient
	tokenManager TokenManager
}

func NewAuthService(client *ent.UserClient, tokenManager TokenManager) AuthService {
	return &authService{client: client, tokenManager: tokenManager}
}

func (as *authService) Register(ctx context.Context, u datastruct.RequestUser) (string, error) {
	res, err := as.client.
		Create().
		SetUsername(u.Username).
		SetFirstName(u.FirstName).
		SetLastName(u.LastName).
		SetEmail(u.Email).
		SetPassword(u.Password).
		Save(ctx)
	if err != nil {
		return "", err
	}

	token, err := as.tokenManager.NewJWT(*res)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (as *authService) Login(ctx context.Context, l datastruct.LoginBody) (string, error) {
	var res *ent.User
	var err error

	if strings.Contains(l.UsernameOrEmail, "@") {
		res, err = as.client.Query().Where(user.EmailEQ(l.UsernameOrEmail)).First(ctx)
	} else {
		res, err = as.client.Query().Where(user.UsernameEQ(l.UsernameOrEmail)).First(ctx)
	}
	if err != nil {
		return "", errors.New("invalid Username or Password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(l.Password))
	if err != nil {
		return "", errors.New("invalid Username or Password")
	}

	token, err := as.tokenManager.NewJWT(*res)
	if err != nil {
		return "", err
	}

	return token, nil
}
