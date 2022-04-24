package rpc

import (
	"context"
	"net/mail"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s userServer) VerifyEmail(ctx context.Context, req *ug.VerifyEmailRequest) (*cg.MessageResponse, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	u, err := s.us.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	if req.Code != u.EmailCode {
		return nil, status.Error(codes.InvalidArgument, "Invalid Code")
	}

	u, err = s.us.UpdateVerified(ctx, req.Email, true)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.MessageResponse{Message: "Your Email is now verified"}, nil

}
