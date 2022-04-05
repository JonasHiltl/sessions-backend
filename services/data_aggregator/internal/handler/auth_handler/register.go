package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
)

func (h *authGatewayHandler) Register(c *fiber.Ctx) error {
	req := new(auth.RegisterRequest)

	res, err := h.c.Register(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
