package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
)

func (s *userServer) CreateUser(c context.Context, req *ug.CreateUserRequest) (*ug.PublicUser, error) {
	du := datastruct.RequestUser{
		Username:  req.Username,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
		Avatar:    req.Avatar,
	}

	u, err := s.us.Create(c, du)
	if err != nil {
		return nil, err
	}
	return u.ToPublicUser(), nil
}
