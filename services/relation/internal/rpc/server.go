package rpc

import (
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/service"
)

type relationServer struct {
	frs service.FriendRelationService
	cg.UnimplementedRelationServiceServer
}

func NewRelationServer(frs service.FriendRelationService) cg.RelationServiceServer {
	return &relationServer{
		frs: frs,
	}
}
