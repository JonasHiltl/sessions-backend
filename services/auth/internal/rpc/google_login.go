package handler

import (
	"context"

	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/types"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/auth/internal/dto"
)

func (s *authServer) GoogleLogin(c context.Context, req *ag.GoogleLoginRequest) (*ag.LoginResponse, error) {
	claims, err := s.google.ValidateGoogleJWT(req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	du := dto.AuthUser{
		Provider:      types.Google,
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
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

	return &ag.LoginResponse{
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
