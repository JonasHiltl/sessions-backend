package handler

import (
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	"github.com/jonashiltl/sessions-backend/services/party/internal/service"
)

type partyServer struct {
	ps service.PartyService
	pg.UnimplementedPartyServiceServer
}

func NewPartyServer(ps service.PartyService) pg.PartyServiceServer {
	return &partyServer{
		ps: ps,
	}
}
