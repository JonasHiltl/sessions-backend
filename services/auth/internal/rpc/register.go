package rpc

import (
	"context"
	"net/mail"

	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/types"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) Register(c context.Context, req *ag.RegisterRequest) (*ag.RegisterResponse, error) {
	hash, err := s.password.HashPassword(req.Password)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	du := dto.AuthUser{
		Email:         req.Email,
		EmailVerified: false,
		PasswordHash:  hash,
		Role:          types.UserRole,
	}

	u, err := s.authService.Create(c, du)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	t, err := s.token.NewJWT(u)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ag.RegisterResponse{
		Token: t,
		AuthUser: &ag.AuthUser{
			Id:            u.Id.Hex(),
			Provider:      u.Provider,
			Email:         u.Email,
			EmailVerified: u.EmailVerified,
			EmailCode:     u.EmailCode,
			Role:          u.Role,
		},
	}, nil
}
