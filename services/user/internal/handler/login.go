package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
)

func (s *userServer) Login(c context.Context, req *ug.LoginRequest) (*ug.AuthResponse, error) {
	l := datastruct.LoginBody{
		UsernameOrEmail: req.UsernameOrEmail,
		Password:        req.Password,
	}
	v := validator.New()
	err := v.Struct(l)
	if err != nil {
		return nil, err
	}

	token, err := s.as.Login(c, l)
	if err != nil {
		return nil, err
	}

	res := &ug.AuthResponse{
		Token:   token,
		Message: "Successfully Logged in",
	}

	return res, nil
}
