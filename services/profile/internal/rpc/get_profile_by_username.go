package handler

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s *profileServer) GetProfileByUsername(c context.Context, req *pg.GetProfileByUsernameRequest) (*pg.Profile, error) {
	p, err := s.ps.GetByUsername(c, req.Username)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
