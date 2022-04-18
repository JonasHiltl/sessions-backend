package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (s *userServer) GetMe(ctx context.Context, req *ug.GetMeRequest) (*ug.User, error) {
	u, err := s.us.GetById(ctx, req.Id)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return u.ToGRPCUser(), nil
}
