package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/datastruct"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/dto"
)

func (s *authServer) GoogleLogin(c context.Context, req *ag.GoogleLoginRequest) (*ag.TokenResponse, error) {
	claims, err := s.google.ValidateGoogleJWT(req.Token)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	du := dto.AuthUser{
		Provider:      datastruct.Google,
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
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

	return &ag.TokenResponse{Jwt: t, Message: "Successfully logged in"}, nil
}
