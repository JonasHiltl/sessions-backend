package handler

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

func (s *userServer) GetUserByEmail(c context.Context, req *ug.GetUserByEmailRequest) (*ug.PublicUser, error) {
	u, err := s.us.GetByEmail(c, req.Email)
	if err != nil {
		return nil, err
	}

	return u.ToPublicUser(), nil
}
