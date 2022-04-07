package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (h *storyGatewayHandler) DeleteStory(c *fiber.Ctx) error {
	// TODO read u_id from headers
	u_id := "tawtrwa"

	sId := c.Params("id")

	res, err := h.storyClient.DeleteStory(c.Context(), &story.DeleteStoryRequest{RequesterId: u_id, SId: sId})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
