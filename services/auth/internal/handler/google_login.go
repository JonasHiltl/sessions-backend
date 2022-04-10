package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comtypes"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/dto"
)

func (s *authServer) GoogleLogin(c context.Context, req *ag.GoogleLoginRequest) (*ag.LoginResponse, error) {
	claims, err := s.google.ValidateGoogleJWT(req.Token)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	du := dto.AuthUser{
		Provider:      comtypes.Google,
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		Role:          comtypes.UserRole,
	}

	u, err := s.authService.Create(c, du)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	t, err := s.token.NewJWT(u)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &ag.LoginResponse{
		Token: t,
		AuthUser: &ag.AuthUser{
			Id:            u.Id.Hex(),
			Provider:      u.Provider.String(),
			Email:         u.Email,
			EmailVerified: u.EmailVerified,
			EmailCode:     u.EmailCode,
			Role:          u.Role.String(),
		},
	}, nil
}
