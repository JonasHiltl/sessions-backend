package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s userServer) Login(ctx context.Context, req *ug.LoginRequest) (*ug.LoginResponse, error) {
	u, err := s.us.GetByEmailOrUsername(ctx, req.UsernameOrEmail)
	if err != nil {
		if err.Error() == "no user found" {
			return nil, status.Error(codes.InvalidArgument, "Invalid Username or Email")
		}
		return nil, utils.HandleError(err)
	}

	pwEqual := s.password.CheckPasswordHash(req.Password, u.PasswordHash)
	if !pwEqual {
		return nil, status.Error(codes.InvalidArgument, "Invalid Password")
	}

	t, err := s.token.NewJWT(u)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ug.LoginResponse{
		Token: t,
		User:  u.ToGRPCUser(),
	}, nil

}
