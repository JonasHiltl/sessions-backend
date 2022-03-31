package handler

import (
	"context"
	"log"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) CreateUser(c context.Context, req *ug.CreateUserRequest) (*ug.PublicUser, error) {
	du := dto.User{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
		Avatar:    req.Avatar,
	}

	log.Println(du)

	u, err := s.us.Create(c, du)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return u.ToPublicUser(), nil
}
