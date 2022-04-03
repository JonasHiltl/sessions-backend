package handler

import (
	"context"
	"net/mail"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) Register(c context.Context, req *ag.RegisterRequest) (*ag.TokenResponse, error) {
	hash, err := s.password.HashPassword(req.Password)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	du := dto.AuthUser{
		Email:         req.Email,
		EmailVerified: false,
		PasswordHash:  hash,
		Role:          datastruct.UserRole,
	}

	u, err := s.authService.Create(c, du)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	// TODO: create profile through gRPC

	t, err := s.token.NewJWT(u)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &ag.TokenResponse{Jwt: t, Message: "Successfully registered"}, nil
}
