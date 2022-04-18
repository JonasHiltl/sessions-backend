package rpc

import (
	"context"
	"net/mail"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/types"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *userServer) Register(ctx context.Context, req *ug.RegisterRequest) (*ug.RegisterResponse, error) {
	hash, err := s.password.HashPassword(req.Password)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	_, err = mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	du := datastruct.User{
		Id:           primitive.NewObjectID(),
		Email:        req.Email,
		Username:     req.Username,
		Firstname:    req.Firstname,
		Lastname:     req.Lastname,
		PasswordHash: hash,
		Role:         types.UserRole.String(),
	}

	u, err := s.us.Create(ctx, du)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	t, err := s.token.NewJWT(u)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &ug.RegisterResponse{
		Token: t,
		User:  u.ToGRPCUser(),
	}, nil
}
