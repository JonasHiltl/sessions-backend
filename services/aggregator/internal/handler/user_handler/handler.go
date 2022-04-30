package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	rg "github.com/jonashiltl/sessions-backend/packages/grpc/relation"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type userGatewayHandler struct {
	uc ug.UserServiceClient
	rc rg.RelationServiceClient
}

type UserGatewayHandler interface {
	GetMe(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	UsernameTaken(c *fiber.Ctx) error
}

func NewUserGatewayHandler(uc ug.UserServiceClient, rc rg.RelationServiceClient) UserGatewayHandler {
	return &userGatewayHandler{
		uc: uc,
		rc: rc,
	}
}
