package handler

import (
	"context"
	"net/mail"

	"github.com/jonashiltl/sessions-backend/packages/comutils"
	ag "github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *authServer) VerifyEmail(ctx context.Context, req *ag.VerifyEmailRequest) (*cg.MessageResponse, error) {
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Email")
	}

	u, err := s.authService.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	if req.Code != u.EmailCode {
		return nil, status.Error(codes.InvalidArgument, "Invalid Code")
	}

	u, err = s.authService.SetVerified(ctx, req.Email)
	if err != nil {
		return nil, comutils.HandleError(err)
	}

	return &cg.MessageResponse{Message: "Your Email is now verified"}, nil
}
