package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

type partyGatewayHandler struct {
	partyClient party.PartyServiceClient
}

type PartyGatewayHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error
}

func NewPartyGatewayHandler(partyClient party.PartyServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		partyClient: partyClient,
	}
}
