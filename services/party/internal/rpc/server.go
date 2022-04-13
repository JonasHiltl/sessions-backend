package rpc

import (
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/packages/stream"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
)

type partyServer struct {
	ps     service.PartyService
	stream stream.Stream
	pg.UnimplementedPartyServiceServer
}

func NewPartyServer(ps service.PartyService, stream stream.Stream) pg.PartyServiceServer {
	return &partyServer{
		ps:     ps,
		stream: stream,
	}
}
