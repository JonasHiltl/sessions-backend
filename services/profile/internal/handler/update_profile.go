package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) UpdateProfile(c context.Context, req *pg.UpdateProfileRequest) (*pg.PublicProfile, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, err
	}

	if req.Id != me.Sub {
		return nil, status.Error(codes.Unauthenticated, "You can only update your own information")
	}

	du := dto.Profile{
		Id:        req.Id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	if du.Avatar != "" {
		loc, err := s.uploadS.Upload(c, me.Sub, du.Avatar)
		if err != nil {
			return nil, err
		}
		du.Avatar = loc
	}

	u, err := s.us.Update(c, du)
	if err != nil {
		return nil, err
	}

	return u.ToPublicProfile(), nil
}
