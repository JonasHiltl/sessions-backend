package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (h *storyGatewayHandler) GetStoryByUser(c *fiber.Ctx) error {
	uId := c.Params("id")
	nextPage := c.Query("nextPage")

	res, err := h.storyClient.GetByUser(c.Context(), &story.GetByUserRequest{UId: uId, NextPage: nextPage})
	if err != nil {
		return err
	}

	// TODO: if tagged friends exist get profile of friends

	return c.Status(fiber.StatusOK).JSON(res)
}
