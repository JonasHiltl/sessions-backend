package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s *userServer) GetMe(c context.Context, req *ug.Empty) (*ug.PublicUser, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, err
	}

	u, err := s.us.GetById(c, me.Sub)
	if err != nil {
		return nil, err
	}

	return u.ToPublicUser(), nil
}
