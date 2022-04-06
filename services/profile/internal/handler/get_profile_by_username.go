package handler

import (
	"context"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (s *profileServer) GetProfileByUsername(c context.Context, req *pg.GetProfileByUsernameRequest) (*pg.Profile, error) {
	p, err := s.us.GetByUsername(c, req.Username)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
