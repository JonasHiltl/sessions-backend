package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (h *profileGatewayHandler) GetProfile(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := h.profileClient.GetProfile(c.Context(), &profile.GetProfileRequest{Id: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
