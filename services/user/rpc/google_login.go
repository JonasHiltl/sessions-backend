package rpc

import (
	"context"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/types"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/datastruct"
)

func (s userServer) GoogleLogin(ctx context.Context, req *ug.GoogleLoginRequest) (*ug.LoginResponse, error) {
	claims, err := s.google.ValidateGoogleJWT(req.Token)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	du := datastruct.User{
		Provider:      types.Google.String(),
		Firstname:     claims.FirstName,
		Lastname:      claims.LastName,
		Email:         claims.Email,
		EmailVerified: claims.EmailVerified,
		Role:          types.UserRole.String(),
	}

	u, err := s.us.Create(ctx, du)
	if err != nil {
		return nil, utils.HandleError(err)
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
