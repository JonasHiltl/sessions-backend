package rpc

import (
	"context"
	"time"

	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/services/party/internal/datastruct"
	"github.com/segmentio/ksuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s partyServer) FavorParty(ctx context.Context, req *pg.FavorPartyRequest) (*pg.FavoriteParty, error) {
	_, err := ksuid.Parse(req.PartyId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid Party id")
	}

	_, err = ksuid.Parse(req.UserId)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "Invalid User id")
	}

	fp, err := s.ps.FavorParty(ctx, datastruct.FavoriteParty{
		UserId:      req.UserId,
		PartyId:     req.PartyId,
		FavoritedAt: time.Now(),
	})
	if err != nil {
		return nil, utils.HandleError(err)
	}

	return fp.ToGRPCFavoriteParty(), nil
}
