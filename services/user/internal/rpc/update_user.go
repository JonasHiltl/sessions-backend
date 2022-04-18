package rpc

import (
	"context"
	"errors"

	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/internal/datastruct"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *userServer) UpdateUser(ctx context.Context, req *ug.UpdateUserRequest) (*ug.User, error) {
	id, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	du := datastruct.User{
		Id:        id,
		Username:  req.Username,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Avatar:    req.Avatar,
	}

	if du.Avatar != "" {
		loc, err := s.upload.Upload(ctx, req.Id, du.Avatar)
		if err != nil {
			return nil, utils.HandleError(err)
		}
		du.Avatar = loc
	}

	if req.Password != "" {
		hash, err := s.password.HashPassword(req.Password)
		if err != nil {
			return nil, utils.HandleError(err)
		}
		du.PasswordHash = hash
	}

	u, err := s.us.Update(ctx, du)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return u.ToGRPCUser(), nil
}
