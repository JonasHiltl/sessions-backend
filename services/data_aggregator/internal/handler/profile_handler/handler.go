package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	pg "github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
)

type profileGatewayHandler struct {
	profileClient  pg.ProfileServiceClient
	relationClient rg.RelationServiceClient
}

type ProfileGatewayHandler interface {
	GetMe(c *fiber.Ctx) error
	GetProfileByUsername(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	UsernameTaken(c *fiber.Ctx) error
}

func NewProfileGatewayHandler(profileClient pg.ProfileServiceClient, relationClient rg.RelationServiceClient) ProfileGatewayHandler {
	return &profileGatewayHandler{
		profileClient:  profileClient,
		relationClient: relationClient,
	}
}
