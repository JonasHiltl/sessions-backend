package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) UpdateUser(c context.Context, req *ug.UpdateUserRequest) (*ug.PublicUser, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, err
	}

	if req.Id != me.Sub {
		return nil, status.Error(codes.Unauthenticated, "You can only update your own information")
	}

	du := datastruct.RequestUser{
		Username:  req.Username,
		FirstName: req.Firstname,
		LastName:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
		Avatar:    req.Avatar,
	}

	if du.Avatar != "" {
		loc, err := s.uploadS.Upload(c, me.Sub, du.Avatar)
		if err != nil {
			return nil, err
		}
		du.Avatar = loc
	}

	u, err := s.us.Update(c, me.Sub, du)
	if err != nil {
		return nil, err
	}

	return u.ToPublicUser(), nil
}
