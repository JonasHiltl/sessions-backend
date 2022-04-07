package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (h *storyGatewayHandler) GetStory(c *fiber.Ctx) error {
	sId := c.Params("id")

	res, err := h.storyClient.GetStory(c.Context(), &story.GetStoryRequest{SId: sId})
	if err != nil {
		return err
	}

	// TODO: if tagged friends exist get profile of friends

	return c.Status(fiber.StatusOK).JSON(res)
}
