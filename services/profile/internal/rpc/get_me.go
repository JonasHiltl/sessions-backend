package rpc

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s *profileServer) GetMe(c context.Context, req *pg.GetMeRequest) (*pg.Profile, error) {
	p, err := s.ps.GetById(c, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return p.ToGRPCProfile(), nil
}
