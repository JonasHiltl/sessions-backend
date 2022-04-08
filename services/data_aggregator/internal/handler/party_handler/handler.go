package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/party"
	prg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	sg "github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

type partyGatewayHandler struct {
	partyClient   pg.PartyServiceClient
	profileClient prg.ProfileServiceClient
	storyClient   sg.StoryServiceClient
}

type PartyGatewayHandler interface {
	CreateParty(c *fiber.Ctx) error
	UpdateParty(c *fiber.Ctx) error
	DeleteParty(c *fiber.Ctx) error
	GetParty(c *fiber.Ctx) error
	GetPartyByUser(c *fiber.Ctx) error
}

func NewPartyGatewayHandler(partyClient pg.PartyServiceClient, profileClient prg.ProfileServiceClient, storyClient sg.StoryServiceClient) PartyGatewayHandler {
	return &partyGatewayHandler{
		partyClient:   partyClient,
		profileClient: profileClient,
		storyClient:   storyClient,
	}
}
