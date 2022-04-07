package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
)

func (h *authGatewayHandler) Login(c *fiber.Ctx) error {
	req := new(auth.LoginRequest)

	res, err := h.authClient.Login(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
