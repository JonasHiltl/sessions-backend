package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (h *profileGatewayHandler) UpdateProfile(c *fiber.Ctx) error {
	req := new(profile.UpdateProfileRequest)
	if err := c.BodyParser(req); err != nil {
		return err
	}

	res, err := h.profileClient.UpdateProfile(c.Context(), req)
	if err != nil {
		return comutils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
