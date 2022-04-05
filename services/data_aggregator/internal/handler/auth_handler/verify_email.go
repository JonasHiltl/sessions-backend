package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
)

func (h *authGatewayHandler) VerifyEmail(c *fiber.Ctx) error {
	req := new(auth.VerifyEmailRequest)

	res, err := h.c.VerifyEmail(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
