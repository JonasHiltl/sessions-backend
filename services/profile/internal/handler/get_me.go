package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/comutils/middleware"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (s *profileServer) GetMe(c context.Context, req *pg.Empty) (*pg.PublicProfile, error) {
	me, err := middleware.ParseUser(c)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	u, err := s.us.GetById(c, me.Sub)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return u.ToPublicProfile(), nil
}
