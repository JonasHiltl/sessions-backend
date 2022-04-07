package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (h *profileGatewayHandler) GetProfileByUsername(c *fiber.Ctx) error {
	uName := c.Params("username")

	res, err := h.profileClient.GetProfileByUsername(c.Context(), &profile.GetProfileByUsernameRequest{Username: uName})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
