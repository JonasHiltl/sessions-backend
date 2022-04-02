package handler

import (
	"context"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (s *profileServer) GetProfileByUsername(c context.Context, req *pg.GetProfileByUsernameRequest) (*pg.PublicProfile, error) {
	u, err := s.us.GetByUsername(c, req.Username)
	if err != nil {
		return nil, err
	}

	return u.ToPublicProfile(), nil
}
