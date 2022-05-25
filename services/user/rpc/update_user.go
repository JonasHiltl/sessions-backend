package rpc

import (
	"context"
	"errors"
	"net/mail"
	"strings"

	"github.com/jonashiltl/sessions-backend/packages/events"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/user/datastruct"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s userServer) UpdateUser(ctx context.Context, req *ug.UpdateUserRequest) (*ug.User, error) {
	id, err := ksuid.Parse(req.Id)
	if err != nil {
		return nil, errors.New("invalid id")
	}

	du := datastruct.User{
		Id:        id.String(),
		Username:  strings.ToLower(req.Username),
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

	if req.Email != "" {
		_, err = mail.ParseAddress(req.Email)
		if err != nil {
			return nil, status.Error(codes.InvalidArgument, "Invalid Email")
		}
		du.Email = req.Email
	}

	err = s.us.Update(ctx, du)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	s.stream.PublishEvent(&events.ProfileUpdated{
		Profile: du.ToGRPCProfile(),
	})

	return du.ToGRPCUser(), nil
}
