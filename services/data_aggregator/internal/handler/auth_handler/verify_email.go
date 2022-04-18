package authhandler

import (
	"github.com/gofiber/fiber/v2"
	ug "github.com/jonashiltl/sessions-backend/packages/grpc/user"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (h *authGatewayHandler) VerifyEmail(c *fiber.Ctx) error {
	req := new(ug.VerifyEmailRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.uc.VerifyEmail(c.Context(), req)
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
