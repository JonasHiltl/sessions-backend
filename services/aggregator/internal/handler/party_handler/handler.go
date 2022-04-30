package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type partyGatewayHandler struct {
	pc pg.PartyServiceClient
	uc ug.UserServiceClient
	sc sg.StoryServiceClient
}

type PartyGatewayHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error
}

func NewPartyGatewayHandler(pc pg.PartyServiceClient, uc ug.UserServiceClient, sc sg.StoryServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		pc: pc,
		uc: uc,
		sc: sc,
	}
}
