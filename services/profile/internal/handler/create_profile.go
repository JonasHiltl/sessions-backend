package handler

import (
	"context"
	"log"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) CreateProfile(c context.Context, req *pg.CreateProfileRequest) (*pg.PublicProfile, error) {
	du := dto.Profile{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	log.Println(du)

	u, err := s.us.Create(c, du)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return u.ToPublicProfile(), nil
}
