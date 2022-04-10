package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (h *profileGatewayHandler) GetMe(c *fiber.Ctx) error {
	// TODO read id from headers
	id := "tawtrwa"

	res, err := h.profileClient.GetMe(c.Context(), &profile.GetMeRequest{Id: id})
	if err != nil {
		return comutils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
