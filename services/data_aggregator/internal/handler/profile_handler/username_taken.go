package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (h *profileGatewayHandler) UsernameTaken(c *fiber.Ctx) error {
	uName := c.Params("username")

	res, err := h.profileClient.UsernameTaken(c.Context(), &profile.UsernameTakenRequest{Username: uName})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
