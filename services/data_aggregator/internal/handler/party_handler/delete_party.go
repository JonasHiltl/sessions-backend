package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (h *partyGatewayHandler) DeleteParty(c *fiber.Ctx) error {
	// TODO read u_id from headers
	u_id := "tawtrwa"

	pId := c.Params("id")

	res, err := h.partyClient.DeleteParty(c.Context(), &party.DeletePartyRequest{RequesterId: u_id, PId: pId})
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
