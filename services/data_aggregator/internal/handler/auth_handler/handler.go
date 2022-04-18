package authhandler

import (
	"github.com/gofiber/fiber/v2"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
)

type authGatewayHandler struct {
	uc ug.UserServiceClient
}

type AuthGatewayHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error
	VerifyEmail(c *fiber.Ctx) error
}

func NewAuthGatewayHandler(uc ug.UserServiceClient) AuthGatewayHandler {
	return &authGatewayHandler{
		uc: uc,
	}
}
