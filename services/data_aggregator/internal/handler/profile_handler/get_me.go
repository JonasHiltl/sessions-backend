package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
)

func (h *profileGatewayHandler) GetMe(c *fiber.Ctx) error {
	// TODO read id from headers
	id := "tawtrwa"

	res, err := h.profileClient.GetMe(c.Context(), &profile.GetMeRequest{Id: id})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
