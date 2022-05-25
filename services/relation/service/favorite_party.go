package service

import (
	"context"

	"github.com/jonashiltl/sessions-backend/services/relation/datastruct"
)

type FavoriteParty interface {
	FavorParty(ctx context.Context, fp datastruct.FavoriteParty) (datastruct.FavoriteParty, error)
	DefavorParty(ctx context.Context, uId, pId string) error
	GetFavoritePartiesByUser(ctx context.Context, uId string, page []byte, limit uint32) ([]datastruct.FavoriteParty, []byte, error)
	GetFavorisingUsersByParty(ctx context.Context, pId string, page []byte, limit uint32) ([]datastruct.FavoriteParty, []byte, error)
}
