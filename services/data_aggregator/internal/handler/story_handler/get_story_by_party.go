package storyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/story"
)

func (h *storyGatewayHandler) GetStoryByParty(c *fiber.Ctx) error {
	pId := c.Params("id")
	nextPage := c.Query("nextPage")

	res, err := h.storyClient.GetByParty(c.Context(), &story.GetByPartyRequest{PartyId: pId, NextPage: nextPage})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
