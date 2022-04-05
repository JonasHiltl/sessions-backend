package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/dto"
)

func (s *profileServer) CreateProfile(c context.Context, req *pg.CreateProfileRequest) (*pg.PublicProfile, error) {
	du := dto.Profile{
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	u, err := s.us.Create(c, du)
	if err != nil {
		return nil, comutils.HandleError(err)
	}
	return u.ToPublicProfile(), nil
}