package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (h *storyGatewayHandler) PresignURL(c *fiber.Ctx) error {
	key := c.Params("key")

	res, err := h.storyClient.PresignURL(c.Context(), &story.PresignURLRequest{Key: key})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
