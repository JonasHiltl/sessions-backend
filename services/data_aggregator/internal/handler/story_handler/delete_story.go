package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
	"github.com/jonashiltl/sessions-backend/packages/utils"
	"github.com/jonashiltl/sessions-backend/packages/utils/middleware"
)

func (h *storyGatewayHandler) DeleteStory(c *fiber.Ctx) error {
	user := middleware.ParseUser(c)

	sId := c.Params("id")

	res, err := h.sc.DeleteStory(c.Context(), &story.DeleteStoryRequest{RequesterId: user.Sub, SId: sId})
	if err != nil {
		return utils.ToHTTPError(err)
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
