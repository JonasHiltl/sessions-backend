package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/services/profile/internal/dto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *profileServer) UpdateProfile(c context.Context, req *pg.UpdateProfileRequest) (*pg.Profile, error) {
	if req.Id != req.RequesterId {
		return nil, status.Error(codes.Unauthenticated, "You can only update your own profile")
	}

	dp := dto.Profile{
		Id:        req.Id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	if dp.Avatar != "" {
		loc, err := s.uploadS.Upload(c, req.RequesterId, dp.Avatar)
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
