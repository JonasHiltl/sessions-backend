package authhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/auth"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (h *authGatewayHandler) Login(c *fiber.Ctx) error {
	req := new(auth.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.authClient.Login(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
