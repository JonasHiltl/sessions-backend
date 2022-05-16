package rpc

import (
	"context"
	"strings"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s userServer) Login(ctx context.Context, req *ug.LoginRequest) (*ug.LoginResponse, error) {
	var u datastruct.User
	var err error

	if strings.Contains(req.UsernameOrEmail, "@") {
		u, err = s.us.GetByEmail(ctx, req.UsernameOrEmail)
	} else {
		u, err = s.us.GetByUsername(ctx, req.UsernameOrEmail)
	}
	if err != nil {
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
