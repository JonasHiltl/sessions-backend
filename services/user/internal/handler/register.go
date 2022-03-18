package handler

import (
	"context"

	"github.com/go-playground/validator/v10"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
)

func (s *userServer) Register(c context.Context, req *ug.CreateUserRequest) (*ug.AuthResponse, error) {
	u := datastruct.RequestUser{
		Username:  req.Username,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
		Avatar:    req.Avatar,
	}
	v := validator.New()
	err := v.Struct(u)
	if err != nil {
		return nil, err
	}

	token, err := s.as.Register(c, u)
	if err != nil {
		return nil, err
	}

	res := &ug.AuthResponse{
		Token:   token,
		Message: "Successfully Registered",
	}

	return res, nil
}
