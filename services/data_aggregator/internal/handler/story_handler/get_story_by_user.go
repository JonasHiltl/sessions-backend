package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/comutils"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (h *storyGatewayHandler) GetStoryByUser(c *fiber.Ctx) error {
	userId := c.Params("id")
	nextPage := c.Query("nextPage")

	res, err := h.storyClient.GetByUser(c.Context(), &story.GetByUserRequest{UserId: userId, NextPage: nextPage})
	if err != nil {
		return comutils.ToHTTPError(c, err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
