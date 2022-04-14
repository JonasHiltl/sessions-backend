package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h *profileGatewayHandler) GetMe(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	res, err := h.profileClient.GetMe(c.Context(), &profile.GetMeRequest{Id: user.Sub})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
