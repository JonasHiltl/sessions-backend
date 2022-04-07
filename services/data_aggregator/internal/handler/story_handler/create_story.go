package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (h *storyGatewayHandler) CreateStory(c *fiber.Ctx) error {
	req := new(story.CreateStoryRequest)

	res, err := h.storyClient.CreateStory(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
