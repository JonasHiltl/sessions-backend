package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
)

type authGatewayHandler struct {
	c auth.AuthServiceClient
}

type AuthGatewayHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	GoogleLogin(c *fiber.Ctx) error
	VerifyEmail(c *fiber.Ctx) error
}

func NewAuthGatewayHandler(ac auth.AuthServiceClient) AuthGatewayHandler {
	return &authGatewayHandler{
		c: ac,
	}
}