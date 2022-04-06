package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) UpdateProfile(c context.Context, req *pg.UpdateProfileRequest) (*pg.Profile, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	if req.Id != me.Sub {
		return nil, status.Error(codes.Unauthenticated, "You can only update your own information")
	}

	dp := dto.Profile{
		Id:        req.Id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	if dp.Avatar != "" {
		loc, err := s.uploadS.Upload(c, me.Sub, dp.Avatar)
		if err != nil {
			return nil, comutils.HandleError(err)
		}
		dp.Avatar = loc
	}

	p, err := s.us.Update(c, dp)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
