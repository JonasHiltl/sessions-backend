package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (h *partyGatewayHandler) GetParty(c *fiber.Ctx) error {
	id := c.Params("id")

	res, err := h.c.GetParty(c.Context(), &party.GetPartyRequest{PId: id})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
