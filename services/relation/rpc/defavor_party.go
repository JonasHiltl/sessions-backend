package rpc

import (
	"context"

	cg "github.com/jonashiltl/sessions-backend/packages/grpc/common"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s relationServer) DefavorParty(ctx context.Context, req *rg.FavorPartyRequest) (*cg.SuccessIndicator, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	_, err = ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	err = s.fp.DefavorParty(ctx, req.UserId, req.PartyId)
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return &cg.SuccessIndicator{Sucess: true}, nil
}
