package partyhandler

import "github.com/jonashiltl/sessions-backend/packages/grpc/party"

type partyGatewayHandler struct {
	c party.PartyServiceClient
}

type PartyGatewayHandler interface {
}

func NewPartyGatewayHandler(pc party.PartyServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		c: pc,
	}
}
