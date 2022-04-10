package handler

import (
	"context"
	"net/mail"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) Login(c context.Context, req *ag.LoginRequest) (*ag.LoginResponse, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	u, err := s.authService.GetByEmail(c, req.Email)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	pwEqual := s.password.CheckPasswordHash(req.Password, u.PasswordHash)
	if !pwEqual {
		return nil, status.Error(codes.InvalidArgument, "Invalid Password")
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
