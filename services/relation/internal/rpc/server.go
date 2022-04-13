package rpc

import (
	cg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/relation/internal/service"
)

type relationServer struct {
	frs    service.FriendRelationService
	stream stream.Stream
	cg.UnimplementedRelationServiceServer
}

func NewRelationServer(frs service.FriendRelationService, stream stream.Stream) cg.RelationServiceServer {
	return &relationServer{
		frs:    frs,
		stream: stream,
	}
}
