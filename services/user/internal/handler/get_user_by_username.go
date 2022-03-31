package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s *userServer) GetUserByUsername(c context.Context, req *ug.GetUserByUsernameRequest) (*ug.PublicUser, error) {
	u, err := s.us.GetByUsername(c, req.Username)
	if err != nil {
		return nil, err
	}

	return u.ToPublicUser(), nil
}
