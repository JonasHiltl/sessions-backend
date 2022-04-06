package profilehandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/profile"
)

func (h *profileGatewayHandler) UpdateProfile(c *fiber.Ctx) error {
	req := new(profile.UpdateProfileRequest)

	res, err := h.c.UpdateProfile(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
