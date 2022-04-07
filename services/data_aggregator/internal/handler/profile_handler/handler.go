package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

type profileGatewayHandler struct {
	profileClient profile.ProfileServiceClient
}

type ProfileGatewayHandler interface {
	GetMe(c *fiber.Ctx) error
	GetProfileByUsername(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateProfile(c *fiber.Ctx) error
	UsernameTaken(c *fiber.Ctx) error
}

func NewProfileGatewayHandler(profileClient profile.ProfileServiceClient) ProfileGatewayHandler {
	return &profileGatewayHandler{
		profileClient: profileClient,
	}
}
