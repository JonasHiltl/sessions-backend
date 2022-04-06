package partyhandler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jonashiltl/sessions-backend/packages/grpc/party"
)

func (h *partyGatewayHandler) UpdateParty(c *fiber.Ctx) error {
	req := new(party.UpdatePartyRequest)

	res, err := h.c.UpdateParty(c.Context(), req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(res)
}
