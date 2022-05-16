package rpc

import (
	"context"
	"net/mail"
	"strings"

	"github.com/jonashiltl/sessions-backend/packages/events"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/types"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s userServer) Register(ctx context.Context, req *ug.RegisterRequest) (*ug.RegisterResponse, error) {
	hash, err := s.password.HashPassword(req.Password)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	du := datastruct.User{
		Id:            ksuid.New().String(),
		Email:         req.Email,
		EmailVerified: false,
		Username:      strings.ToLower(req.Username),
		Firstname:     req.Firstname,
		Lastname:      req.Lastname,
		PasswordHash:  hash,
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

	s.stream.PublishEvent(&events.Registered{
		Id:        u.Id,
		Email:     u.Email,
		Username:  u.Username,
		Firstname: u.Firstname,
		Lastname:  u.Lastname,
	})

	return &ug.RegisterResponse{
		Token: t,
		User:  u.ToGRPCUser(),
	}, nil
}
